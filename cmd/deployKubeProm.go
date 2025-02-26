/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/hvvit/elasticsearch-monitoring-stack/helm"
	"github.com/spf13/cobra"
)

// deployKubePromCmd represents the deployKubeProm command
var deployKubePromCmd = &cobra.Command{
	Use:   "deployKubeProm",
	Short: "command to deploy kube-prometheus-stack",
	Long: `This command will deploy the kube-prometheus-stack helm chart
		by default it will deploy the kube-prometheus-stack chart from the prometheus-community helm repository
		namespace is set to monitoring by default
		version is set to 69.5.1 by default
		chart-url is set to https://prometheus-community.github.io/helm-charts by default
		chart-path is set to prometheus-community/kube-prometheus-stack by default`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("release").Value.String()
		namespace := cmd.Flag("namespace").Value.String()
		version := cmd.Flag("version").Value.String()
		chartUrl := cmd.Flag("chart-url").Value.String()
		chartPath := cmd.Flag("chart-path").Value.String()
		fmt.Println("deployKubeProm called")
		helmPackage := &helm.Package{
			Name:      name,
			Namespace: namespace,
			Version:   version,
			ChartUrl:  chartUrl,
			ChartPath: chartPath,
			Values:    map[string]interface{}{},
		}
		_, err := helm.UpgradeOrInstallChart(helmPackage)
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deployKubePromCmd)
	deployKubePromCmd.Flags().StringP("release", "r", "prom-stack", "The name of the kube-prometheus-stack release")
	deployKubePromCmd.Flags().StringP("namespace", "n", "monitoring", "The namespace to deploy kube-prometheus-stack")
	deployKubePromCmd.Flags().StringP("version", "v", "69.5.1", "The version of kube-prometheus-stack to deploy")
	deployKubePromCmd.Flags().StringP("chart-url", "c", "https://prometheus-community.github.io/helm-charts", "The URL of the helm chart")
	deployKubePromCmd.Flags().StringP("chart-path", "p", "kube-prometheus-stack", "The path to the helm chart")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployKubePromCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployKubePromCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
