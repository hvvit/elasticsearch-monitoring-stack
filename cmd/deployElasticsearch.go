/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/hvvit/elasticsearch-monitoring-stack/helm"
	"github.com/hvvit/elasticsearch-monitoring-stack/kube"
	"github.com/spf13/cobra"
)

// deployElasticsearchCmd represents the deployElasticsearch command
var deployElasticsearchCmd = &cobra.Command{
	Use:   "deployElasticsearch",
	Short: "This Will deploy the elasticsearch-operator helm chart",
	Long: `This command will deploy the elasticsearch-operator helm chart
		by default it will deploy the elasticsearch-operator chart from the elastic helm repository
		namespace is set to elastic-system by default
		version is set to 2.16.1 by default
		chart-url is set to https://helm.elastic.co by default
		chart-path is set to eck-operator by default
		This will also try to install default elasticsearch setup`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("deployElasticsearch called")
		deploymentFile := cmd.Flag("deployment-file").Value.String()
		namespace := cmd.Flag("namespace").Value.String()
		alertRules := cmd.Flag("alert-rules").Value.String()
		dashboardFile := cmd.Flag("dashboard").Value.String()
		log.Println("deploying elasticsearch-operator")
		elasticOperatorPackage := &helm.Package{
			Name:      "elastic-operator",
			Namespace: "elastic-system",
			Version:   "2.16.1",
			ChartUrl:  "https://helm.elastic.co",
			ChartPath: "eck-operator",
			Values: map[string]interface{}{
				"serviceMonitor": map[string]interface{}{
					"enabled":   true,
					"namespace": "monitoring",
				},
				"config": map[string]interface{}{
					"metrics": map[string]interface{}{
						"port": 9090,
					},
				},
			},
		}
		_, err := helm.UpgradeOrInstallChart(elasticOperatorPackage)
		if err != nil {
			log.Println(err)
		}
		log.Println("deploying elasticsearch")
		elasticDeployment := &kube.KubeYaml{
			File:       deploymentFile,
			Namespace:  namespace,
			ApiGroup:   "elasticsearch.k8s.elastic.co",
			Resource:   "elasticsearches",
			ApiVersion: "v1",
		}
		err = kube.ApplyYamlFile(elasticDeployment)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Applied elasticsearch yaml")
		log.Println("Fetching elastic user password")
		secretName := "es-cluster-es-elastic-user"
		secret, err := kube.FetchSecret(secretName, namespace, "elastic")
		if err != nil {
			log.Println(err)
		}
		esUrl := fmt.Sprintf("https://elastic:%s@es-cluster-es-http:9200", string(secret))
		log.Println("Elasticsearch URL: ", esUrl)
		log.Println("Deploying Prometheus ElasticSearch Exporter")
		log.Println("Reading rules file for alertmanager rules")
		if err != nil {
			log.Fatal(err)
		}
		elasticExporter := &helm.Package{
			Name:      "elastic-exporter",
			Namespace: "default",
			Version:   "6.6.1",
			ChartUrl:  "https://prometheus-community.github.io/helm-charts",
			ChartPath: "prometheus-elasticsearch-exporter",
			Values: map[string]interface{}{
				"es": map[string]interface{}{
					"uri":           esUrl,
					"aliases":       true,
					"timeout":       "10s",
					"sslSkipVerify": true,
				},
				"serviceMonitor": map[string]interface{}{
					"enabled":   true,
					"namespace": "default",
					"interval":  "10s",
					"labels": map[string]interface{}{
						"release": "prom-stack",
					},
				},
			},
		}
		_, err = helm.UpgradeOrInstallChart(elasticExporter)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Deployed ElasticSearch Exporter")
		log.Println("Deploying Alertmanager rules")
		alertManagerRules := &kube.KubeYaml{
			File:       alertRules,
			Namespace:  "default",
			ApiGroup:   "monitoring.coreos.com",
			Resource:   "prometheusrules",
			ApiVersion: "v1",
		}
		err = kube.ApplyYamlFile(alertManagerRules)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Applied Alertmanager rules")
		log.Println("Deploying Dashboard")
		dashboard := &kube.KubeYaml{
			File:       dashboardFile,
			Namespace:  "monitoring",
			ApiGroup:   "",
			Resource:   "configmaps",
			ApiVersion: "v1",
		}
		err = kube.ApplyYamlFile(dashboard)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Applied Dashboard")
	},
}

func init() {
	rootCmd.AddCommand(deployElasticsearchCmd)

	deployElasticsearchCmd.Flags().StringP("deployment-file", "f", "./elasticsearch-deployment.yaml", "The path to the helm deployment file for elastic search")
	deployElasticsearchCmd.Flags().StringP("namespace", "n", "default", "The namespace to deploy elasticsearch-operator")
	deployElasticsearchCmd.Flags().StringP("alert-rules", "a", "./rules.yaml", "The path to the alertmanager rules file")
	deployElasticsearchCmd.Flags().StringP("dashboard", "d", "./dashboard.yaml", "The path to the dashboard file")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployElasticsearchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployElasticsearchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
