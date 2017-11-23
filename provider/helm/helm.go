package helm

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"log"
	"os/exec"

	common "github.com/enablecloud/kulbe/common"
	"k8s.io/helm/pkg/downloader"
	"k8s.io/helm/pkg/getter"
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/helm/environment"
	"k8s.io/helm/pkg/helm/portforwarder"
	repo "k8s.io/helm/pkg/repo"
)

var (
	TillerAddress = "tiller-deploy:44134"
)

func InstallRelease(release string, version string, namespace string, releasename string, config *common.Default) {

	cmd := exec.Command("helm", "install", release, "--version", version, "-n", releasename, "--namespace", namespace)

	if version == "" {
		newcmd := exec.Command("helm", "install", release, "-n", releasename, "--namespace", namespace)
		cmd = newcmd
	}
	log.Printf("Command start with : %s", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	log.Printf("Command finished with : %v", out.String())
}

func DeleteRelease(releasename string, config *common.Default) {
	cmd := exec.Command("helm", "delete", "--purge", releasename)
	log.Printf("Command start with : %s", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	log.Printf("Command finished with : %v", out.String())

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
