package kube

import (
	"context"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/yaml"
)

func ApplyYamlFile(kubeyaml *KubeYaml) error {
	config, err := createConfig()
	if err != nil {
		return err
	}
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return err
	}
	gvr := createGVR(kubeyaml.ApiGroup, kubeyaml.ApiVersion, kubeyaml.Resource)
	err = applyResourceFromFile(context.TODO(), dynClient, gvr, kubeyaml.Namespace, kubeyaml.File)
	if err != nil {
		return err
	}
	return nil
}

func applyResourceFromFile(ctx context.Context, dynClient dynamic.Interface, gvr schema.GroupVersionResource, namespace string, file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	unstructured := &unstructured.Unstructured{}
	if err := unstructured.UnmarshalJSON(jsonData); err != nil {
		return err
	}
	if namespace != "" {
		unstructured.SetNamespace(namespace)
	}
	patchData, err := unstructured.MarshalJSON()
	if err != nil {
		return err
	}
	resourceClient := dynClient.Resource(gvr).Namespace(namespace)
	_, err = resourceClient.Patch(ctx, unstructured.GetName(), types.ApplyPatchType, patchData, metav1.PatchOptions{
		FieldManager: "binary-apply",
		Force:        boolPtr(true),
	})
	if err != nil {
		return err
	}
	return nil
}

func boolPtr(b bool) *bool {
	return &b
}
