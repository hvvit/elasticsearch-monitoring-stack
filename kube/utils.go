package kube

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func createGVR(group string, version string, resource string) schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}
}

func homeDir() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	if h := homedir; h != "" {
		return h
	}
	return "/"
}

func createConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}
	kubeconfigPath := filepath.Join(homeDir(), ".kube", "config")
	if envKubeconfig := os.Getenv("KUBECONFIG"); envKubeconfig != "" {
		kubeconfigPath = envKubeconfig
	}
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}
	return config, nil
}

func FetchSecret(secretName, namespace, key string) (string, error) {
	config, err := createConfig()
	if err != nil {
		return "", err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	deadline := time.Now().Add(5 * time.Minute)
	pollInterval := 5 * time.Second
	log.Println("Waiting for secret to be available")
	for {
		if time.Now().After(deadline) {
			return "", fmt.Errorf("timed out waiting for secret %s in namespace %s", secretName, namespace)
		}

		// 4. Try to get the Secret
		secret, err := clientset.CoreV1().Secrets(namespace).Get(ctx, secretName, metav1.GetOptions{})
		if err == nil {
			// Secret found
			log.Printf("Secret %q in namespace %q is now present!\n", secretName, namespace)
			return string(secret.Data[key]), nil
		}

		if errors.IsNotFound(err) {
			// Secret not found yet, keep polling
			fmt.Printf("Secret %q not found in namespace %q. Retrying in %v...\n",
				secretName, namespace, pollInterval)
		} else {
			// Some other error occurred
			return "", fmt.Errorf("error checking for Secret %q: %w", secretName, err)
		}

		time.Sleep(pollInterval)
	}

}

// func ReadRulesYAMLFile(path string) ([]map[string]interface{}, error) {
// 	// Read the file contents
// 	fileBytes, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read file: %w", err)
// 	}

// 	// Prepare a variable to hold the unmarshaled data
// 	var data []map[string]interface{}

// 	// Unmarshal the YAML into our slice of maps
// 	if err := goyaml.Unmarshal(fileBytes, &data); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
// 	}
// 	return data, nil
// }
