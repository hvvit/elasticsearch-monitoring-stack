package kube

type KubeYaml struct {
	ApiGroup   string
	ApiVersion string
	Resource   string
	Namespace  string
	File       string
}
