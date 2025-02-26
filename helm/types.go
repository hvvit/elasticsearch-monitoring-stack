package helm

type Package struct {
	Name      string
	Namespace string
	Version   string
	ChartUrl  string
	ChartPath string
	Values    map[string]interface{}
}
