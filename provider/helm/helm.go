package helm

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	common "github.com/enablecloud/kulbe/common"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/downloader"
	"k8s.io/helm/pkg/getter"
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/helm/environment"
	helm_env "k8s.io/helm/pkg/helm/environment"
	"k8s.io/helm/pkg/helm/portforwarder"
	"k8s.io/helm/pkg/proto/hapi/services"
	repo "k8s.io/helm/pkg/repo"
)

var (
	TillerAddress = "tiller-deploy:44134"
)

func InstallRelease(release string, version string, namespace string, releasename string, config *common.Default) *services.InstallReleaseResponse {
	var client = NewHelmImplementer(config)
	// Check chart requirements to make sure all dependencies are present in /charts
	var settings = helm_env.EnvSettings{TillerHost: config.TillerAddress, TillerNamespace: config.TillerNamespace, Home: config.Home, KubeContext: config.KubeContext}

	f, err := repo.LoadRepositoriesFile(config.Home.RepositoryFile())
	var chartURL = ""
	for _, v := range f.Repositories {
		chartURL, err := repo.FindChartInRepoURL(v.URL, release, "", "", "", "", getter.All(settings))
		if err != nil {
			fmt.Printf("Chart error: '%s'\n", err)
		}
		if chartURL != "" {

			break
		}
	}
	path, err := locateChartPath(chartURL, release, version, false, "", "", "", "")
	if err != nil {
		fmt.Printf("Created locateChartPath: '%s'\n", err)
	}
	chartRequested, err := chartutil.Load(path)
	if err != nil {
		fmt.Printf("Created error: '%s'\n", err)
	}
	fmt.Printf("Created error: '%s'\n", err)
	//releaseContent, _ := client.ReleaseContent(release, helm.ContentReleaseVersion(version))
	//releaseContent.Release.Chart
	//chart.Metadata{Version: version, Name: release}
	//&chart.Chart{Metadata: &meta
	response, err := client.InstallReleaseFromChart(chartRequested, namespace, helm.ReleaseName(releasename))
	if err != nil {
		fmt.Printf("Created error: '%s'\n", err)
	}
	return response
}

func DeleteRelease(releasename string, config *common.Default) *services.UninstallReleaseResponse {
	var client = NewHelmImplementer(config)
	response, _ := client.DeleteRelease(releasename)
	return response
}

func NewHelmImplementer(config *common.Default) *helm.Client {
	tunnel, err := portforwarder.New(config.TillerNamespace, config.Clientkub, config.RestConfig)
	if err != nil {
		return nil
	}

	tillerHost := fmt.Sprintf("127.0.0.1:%d", tunnel.Local)
	fmt.Printf("Created tunnel using local port: '%d'\n", tunnel.Local)
	if config.TillerAddress == "" {
		config.TillerAddress = TillerAddress
	} else {
		fmt.Printf("provider.helm: tiller address '%s' supplied", config.TillerAddress)
	}
	if config.TillerTunnel {
		return helm.NewClient(helm.Host(tillerHost))
	}
	return helm.NewClient(helm.Host(config.TillerAddress))
}

// locateChartPath looks for a chart directory in known places, and returns either the full path or an error.
//
// This does not ensure that the chart is well-formed; only that the requested filename exists.
//
// Order of resolution:
// - current working directory
// - if path is absolute or begins with '.', error out here
// - chart repos in $HELM_HOME
// - URL
//
// If 'verify' is true, this will attempt to also verify the chart.
func locateChartPath(repoURL, name, version string, verify bool, keyring,
	certFile, keyFile, caFile string) (string, error) {
	fmt.Println("Search ", name, version)
	fmt.Printf("Search for: '%s' '%s'\n", name, version)
	name = strings.TrimSpace(name)
	version = strings.TrimSpace(version)

	if fi, err := os.Stat(name); err == nil {
		abs, err := filepath.Abs(name)
		if err != nil {
			return abs, err
		}
		if verify {
			if fi.IsDir() {
				return "", errors.New("cannot verify a directory")
			}
			if _, err := downloader.VerifyChart(abs, keyring); err != nil {
				return "", err
			}
		}
		return abs, nil
	}
	if filepath.IsAbs(name) || strings.HasPrefix(name, ".") {
		return name, fmt.Errorf("path %q not found", name)
	}
	var settings = environment.EnvSettings{}
	getter.All(settings)
	crepo := filepath.Join(settings.Home.Repository(), name)
	if _, err := os.Stat(crepo); err == nil {
		return filepath.Abs(crepo)
	}

	dl := downloader.ChartDownloader{
		HelmHome: settings.Home,
		Out:      os.Stdout,
		Keyring:  keyring,
		Getters:  getter.All(settings),
	}
	if verify {
		dl.Verify = downloader.VerifyAlways
	}
	if repoURL != "" {
		chartURL, err := repo.FindChartInRepoURL(repoURL, name, version,
			certFile, keyFile, caFile, getter.All(settings))
		if err != nil {
			return "", err
		}
		name = chartURL
	}

	if _, err := os.Stat(settings.Home.Archive()); os.IsNotExist(err) {
		os.MkdirAll(settings.Home.Archive(), 0744)
	}

	filename, _, err := dl.DownloadTo(name, version, settings.Home.Archive())
	if err == nil {
		lname, err := filepath.Abs(filename)
		if err != nil {
			return filename, err
		}
		fmt.Printf("Fetched %s to %s\n", name, filename)
		return lname, nil
	} else if settings.Debug {
		return filename, err
	}

	return filename, fmt.Errorf("file %q not found", name)
}
