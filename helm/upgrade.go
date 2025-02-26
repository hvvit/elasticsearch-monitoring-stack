package helm

import (
	"fmt"
	"log"
	"os"
	"time"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

func UpgradeOrInstallChart(helmPackage *Package) (*release.Release, error) {
	settings := cli.New()
	deployed, err := checkIfDeployed(helmPackage)
	if err != nil {
		return nil, fmt.Errorf("failed to check if release is deployed: %w", err)
	}
	if !deployed {
		log.Println("Release not deployed, installing")
		log.Println("Installing chart", helmPackage.Name)
		log.Println("Namespace", helmPackage.Namespace)
		log.Println("Version", helmPackage.Version)
		log.Println("Chart URL", helmPackage.ChartUrl)
		log.Println("Chart Path", helmPackage.ChartPath)
		return installChart(helmPackage)
	}
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), helmPackage.Namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		return nil, fmt.Errorf("failed to initialize helm action configuration: %w", err)
	}

	upgradeClient := action.NewUpgrade(actionConfig)
	upgradeClient.Namespace = helmPackage.Namespace
	upgradeClient.Install = true
	upgradeClient.Wait = true
	upgradeClient.RepoURL = helmPackage.ChartUrl
	upgradeClient.Timeout = 10 * time.Minute
	chartPath, err := upgradeClient.ChartPathOptions.LocateChart(helmPackage.ChartPath, settings)
	if err != nil {
		return nil, fmt.Errorf("failed to locate chart: %w", err)
	}
	chartRequested, err := loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart: %w", err)
	}
	// Validate the chart
	if chartRequested.Metadata == nil || chartRequested.Metadata.Name == "" {
		return nil, fmt.Errorf("chart metadata is missing or invalid")
	}

	// Execute the upgrade or install
	rel, err := upgradeClient.Run(helmPackage.Name, chartRequested, helmPackage.Values)
	if err != nil {
		return nil, fmt.Errorf("failed to upgrade or install release: %w", err)
	}
	printReleaseInfo(rel)
	return rel, nil
}

func checkIfDeployed(helmPackage *Package) (bool, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), helmPackage.Namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		return false, fmt.Errorf("failed to initialize helm action configuration: %w", err)
	}
	statusClient := action.NewStatus(actionConfig)
	rel, err := statusClient.Run(helmPackage.Name)
	if err != nil {
		return false, nil
	}
	printReleaseInfo(rel)
	if rel.Info.Status == release.StatusDeployed {

		return true, nil
	}
	if rel.Info.Status == release.StatusFailed {
		return true, nil
	}
	return rel.Info.Status == release.StatusDeployed, nil
}

func installChart(helmPackage *Package) (*release.Release, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), helmPackage.Namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		return nil, fmt.Errorf("failed to initialize helm action configuration: %w", err)
	}
	installClient := action.NewInstall(actionConfig)
	installClient.Namespace = helmPackage.Namespace
	installClient.ReleaseName = helmPackage.Name
	installClient.CreateNamespace = true
	installClient.RepoURL = helmPackage.ChartUrl
	installClient.Wait = true
	installClient.Timeout = 10 * time.Minute
	chartPath, err := installClient.ChartPathOptions.LocateChart(helmPackage.ChartPath, settings)
	if err != nil {
		return nil, fmt.Errorf("failed to locate chart: %w", err)
	}
	chartRequested, err := loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart: %w", err)
	}
	// Validate the chart
	if chartRequested.Metadata == nil || chartRequested.Metadata.Name == "" {
		return nil, fmt.Errorf("chart metadata is missing or invalid")
	}
	rel, err := installClient.Run(chartRequested, helmPackage.Values)
	if err != nil {
		return nil, fmt.Errorf("failed to install release: %w", err)
	}
	printReleaseInfo(rel)
	return rel, nil
}

func printReleaseInfo(rel *release.Release) {
	log.Println("<---------------------------------------------------------->")
	log.Printf("NAME: %s\n", rel.Name)
	log.Printf("LAST DEPLOYED: %s\n", rel.Info.LastDeployed.String())
	log.Printf("NAMESPACE: %s\n", rel.Namespace)
	log.Printf("Chart: %s\n", rel.Chart.Metadata.Name)
	log.Printf("Chart VERSION: %s\n", rel.Chart.Metadata.Version)
	log.Printf("Home: %s\n", rel.Chart.Metadata.Home)
	log.Printf("STATUS: %s\n", rel.Info.Status)
	log.Printf("REVISION: %d\n", rel.Version)
	log.Println("<---------------------------------------------------------->")
}
