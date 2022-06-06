package helm

import (
	"github.com/cppforlife/go-cli-ui/ui"
	"github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/package/builder/build"
	pkgui "github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/package/builder/ui"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
)

type HelmChartStep struct {
	pkgAuthoringUI pkgui.IPkgAuthoringUI
	pkgLocation    string
	pkgBuild       *build.PackageBuild
}

func NewHelmChartStep(ui pkgui.IPkgAuthoringUI, pkgLocation string, pkgBuild *build.PackageBuild) *HelmChartStep {
	helmChart := HelmChartStep{
		pkgAuthoringUI: ui,
		pkgLocation:    pkgLocation,
		pkgBuild:       pkgBuild,
	}
	return &helmChart
}

func (h HelmChartStep) PreInteract() error {
	return nil
}

func (h HelmChartStep) Interact() error {
	helmChart := h.pkgBuild.Spec.Pkg.Spec.Template.Spec.Fetch[0].HelmChart
	if helmChart == nil {
		h.pkgBuild.Spec.Pkg.Spec.Template.Spec.Fetch[0].HelmChart = &v1alpha1.AppFetchHelmChart{
			Repository: &v1alpha1.AppFetchHelmChartRepo{},
		}
	}

	err := h.configureHelmChartName()
	if err != nil {
		return err
	}

	err = h.configureHelmChartVersion()
	if err != nil {
		return err
	}

	err = h.configureHelmChartRepoURL()
	if err != nil {
		return err
	}
	return nil
}

func (h HelmChartStep) PostInteract() error {
	return nil
}

func (h HelmChartStep) configureHelmChartName() error {
	helmChartContent := h.pkgBuild.Spec.Pkg.Spec.Template.Spec.Fetch[0].HelmChart
	defaultName := helmChartContent.Name
	textOpts := ui.TextOpts{
		Label:        "Enter the helm chart name",
		Default:      defaultName,
		ValidateFunc: nil,
	}
	name, err := h.pkgAuthoringUI.AskForText(textOpts)
	if err != nil {
		return err
	}

	helmChartContent.Name = name
	h.pkgBuild.WriteToFile(h.pkgLocation)
	return nil
}

func (h HelmChartStep) configureHelmChartVersion() error {
	helmChartContent := h.pkgBuild.Spec.Pkg.Spec.Template.Spec.Fetch[0].HelmChart
	defaultVersion := helmChartContent.Version
	textOpts := ui.TextOpts{
		Label:        "Enter the helm chart version",
		Default:      defaultVersion,
		ValidateFunc: nil,
	}
	version, err := h.pkgAuthoringUI.AskForText(textOpts)
	if err != nil {
		return err
	}

	helmChartContent.Version = version
	h.pkgBuild.WriteToFile(h.pkgLocation)
	return nil
}

func (h HelmChartStep) configureHelmChartRepoURL() error {
	helmChartContent := h.pkgBuild.Spec.Pkg.Spec.Template.Spec.Fetch[0].HelmChart
	defaultURL := helmChartContent.Repository.URL
	textOpts := ui.TextOpts{
		Label:        "Enter the helm chart repository URL",
		Default:      defaultURL,
		ValidateFunc: nil,
	}
	url, err := h.pkgAuthoringUI.AskForText(textOpts)
	if err != nil {
		return err
	}

	helmChartContent.Repository.URL = url
	h.pkgBuild.WriteToFile(h.pkgLocation)
	return nil
}
