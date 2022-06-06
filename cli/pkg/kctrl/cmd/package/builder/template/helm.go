package template

import (
	"strings"

	"github.com/cppforlife/go-cli-ui/ui"
	pkgui "github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/package/builder/ui"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
)

type HelmTemplateStep struct {
	pkgAuthoringUI pkgui.IPkgAuthoringUI
	appTemplateHelm v1alpha1.AppTemplateHelmTemplate
}

func NewHelmTemplateStep(ui pkgui.IPkgAuthoringUI) *HelmTemplateStep {
	return &HelmTemplateStep{
		pkgAuthoringUI: ui,
	}
}

func (h HelmTemplateStep) PreInteract() error {
	return nil
}

func (h HelmTemplateStep) Interact() error {
    err := h.configureHelmChartPath()
	return nil
}

func (h HelmTemplateStep) PostInteract() error {
	return nil
}

func (h HelmTemplateStep) configureHelmChartPath() error {
	path, err := y.pkgAuthoringUI.AskForText(ui.TextOpts{
		Label:        "Enter the path of where helm chart is present",
		Default:      "",
		ValidateFunc: nil,
	})
	if err != nil {
		return err
	}
	h.appTemplateHelm = v1alpha1.AppTemplateHelmTemplate{Path: path}
	return nil
}


