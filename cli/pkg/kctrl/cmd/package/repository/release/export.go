package release

import (
	cmdcore "github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/core"
	"github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/package/repository/release/build"
)

const (
	ImgpkgBundle string = "Imgpkg"
)

type ExportStep struct {
	ui              cmdcore.AuthoringUI
	pkgRepoLocation string
	pkgRepoBuild    *build.PackageRepoBuild
}

func NewExportStep(ui cmdcore.AuthoringUI, pkgLocation string, pkgRepoBuild *build.PackageRepoBuild) *ExportStep {
	return &ExportStep{
		ui:              ui,
		pkgRepoLocation: pkgLocation,
		pkgRepoBuild:    pkgRepoBuild,
	}
}

func (export *ExportStep) PreInteract() error {
	return nil
}

func (export *ExportStep) Interact() error {
	fetchOptionSelected := ImgpkgBundle

	switch fetchOptionSelected {
	case ImgpkgBundle:
		//imgpkgStep := NewImgPkgStep(export.ui, export.pkgRepoLocation, export.pkgRepoBuild)
		//err := common.Run(imgpkgStep)
		//if err != nil {
		//	return err
		//}
	}
	return nil
}

func (export *ExportStep) PostInteract() error {
	return nil
}
