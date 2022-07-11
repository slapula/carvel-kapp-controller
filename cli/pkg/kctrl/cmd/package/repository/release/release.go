package release

import (
	"fmt"
	"github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/package/repository/release/build"
	cmdlocal "github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/local"
	kcv1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"path/filepath"
	"strings"
	"time"

	"sigs.k8s.io/yaml"

	"github.com/cppforlife/go-cli-ui/ui"
	"github.com/spf13/cobra"
	"github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/app/init/common"
	cmdcore "github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/core"
	"github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/logger"
)

type ReleaseOptions struct {
	ui          cmdcore.AuthoringUI
	depsFactory cmdcore.DepsFactory
	logger      logger.Logger

	pkgRepoVersion string
	chdir          string
	outputLocation string
	debug          bool
}

const (
	PkgRepoBuildFileName = "pkgrepo-build.yml"
	PkgRepoLocation      = "bundle"
	defaultVersion       = "0.0.0-%d"
	lockOutputFolder     = ".imgpkg"
)

func NewReleaseOptions(ui ui.UI, depsFactory cmdcore.DepsFactory, logger logger.Logger) *ReleaseOptions {
	return &ReleaseOptions{ui: cmdcore.NewAuthoringUIImpl(ui), depsFactory: depsFactory, logger: logger}
}

func NewReleaseCmd(o *ReleaseOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release",
		Short: "Build and create a package repository",
		RunE:  func(_ *cobra.Command, args []string) error { return o.Run() },
	}

	cmd.Flags().StringVarP(&o.pkgRepoVersion, "version", "v", "", "Version to be released")
	cmd.Flags().StringVar(&o.chdir, "chdir", "", "Working directory with repo which needs to be bundles")
	cmd.Flags().StringVar(&o.outputLocation, "copy-to", "", "Output location for pkgrepo-build.yml")
	cmd.Flags().BoolVar(&o.debug, "debug", false, "Include debug output")

	return cmd
}

func (o *ReleaseOptions) Run() error {
	o.ui.PrintHeaderText("\nPre-requisite")
	o.ui.PrintInformationalText("Welcome! Before we start on the creating package  repository makes sure you have ")

	if o.pkgRepoVersion == "" {
		o.pkgRepoVersion = fmt.Sprintf(defaultVersion, time.Now().Unix())
	}

	if o.chdir != "" {
		err := os.Chdir(o.chdir)
		if err != nil {
			return err
		}
	}

	//pkgRepoLocation, err := GetPkgRepoLocation()
	//if err != nil {
	//	return err
	//}
	pkgRepoBuild, err := GetPackageRepositoryBuild(PkgRepoBuildFileName)
	if err != nil {
		return err
	}

	o.ui.PrintHeaderText("\nBasic Information(Step 1/3)")
	o.ui.PrintInformationalText("\nA package repository name is the name with which it will be referenced while deploying on the cluster.")
	defaultPkgRepoName := pkgRepoBuild.Name
	textOpts := ui.TextOpts{
		Label:        "Enter the package repository name",
		Default:      defaultPkgRepoName,
		ValidateFunc: nil,
	}
	pkgRepoName, err := o.ui.AskForText(textOpts)
	if err != nil {
		return err
	}
	pkgRepoBuild.Name = pkgRepoName

	o.ui.PrintInformationalText("To create package repository, we will create an imgpkg bundle first. Imgpkg, a Carvel tool, allows users to package, distribute, and relocate a set of files as one OCI artifact: a bundle. Imgpkg bundles are identified with a unique sha256 digest based on the file contents. Imgpkg uses that digest to ensure that the copied contents are identical to those originally pushed.")
	o.ui.PrintInformationalText("\nA package repository bundle is an imgpkg bundle that holds PackageMetadata and Package CRs. Later on, this bundle can be mentioned in the package repository CR to fetch the package and packageMetadata CRs.")
	o.ui.PrintActionableText("\nCreating the required directory structure for imgpkg bundle\n")

	o.ui.PrintHeaderText("\nRegistry URL(Step 2/3)")
	defaultRegistryURL := pkgRepoBuild.Spec.Export.ImgpkgBundle.Image
	textOpts = ui.TextOpts{
		Label:        "Enter the registry url to push the package repository bundle",
		Default:      defaultRegistryURL,
		ValidateFunc: nil,
	}
	registryURL, err := o.ui.AskForText(textOpts)
	if err != nil {
		return err
	}

	pkgRepoBuild.Spec.Export.ImgpkgBundle.Image = strings.TrimSpace(registryURL)
	pkgRepoBuild.WriteToFile()

	//exportStep := NewExportStep(o.ui, pkgRepoLocation, pkgRepoBuild)
	//err = common.Run(exportStep)
	//if err != nil {
	//	return err
	//}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	o.ui.PrintInformationalText("\nLet's use `kbld` to create immutable image reference. Kbld scans all the files in bundle configuration for any references of images and creates a mapping of image tags to a URL with sha256 digest.")
	o.ui.PrintActionableText("Lock image references using Kbld")

	// In-memory app for building and pushing images
	builderApp := kcv1alpha1.App{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kctrl-builder",
			Namespace: "in-memory",
			Annotations: map[string]string{
				"kctrl.carvel.dev/local-fetch-0": ".",
			},
		},
		Spec: kcv1alpha1.AppSpec{
			Fetch: []kcv1alpha1.AppFetch{
				{
					// To be replaced by local fetch
					Git: &kcv1alpha1.AppFetchGit{},
				},
			},
			Template: []kcv1alpha1.AppTemplate{
				{
					Ytt: &kcv1alpha1.AppTemplateYtt{
						Paths: []string{"packages"},
					},
				}, {
					Kbld: &kcv1alpha1.AppTemplateKbld{
						Paths: []string{},
					},
				},
			},
			Deploy: []kcv1alpha1.AppDeploy{
				{
					Kapp: &kcv1alpha1.AppDeployKapp{},
				},
			},
		},
	}

	buildConfigs := cmdlocal.Configs{
		Apps: []kcv1alpha1.App{builderApp},
	}

	// Create temporary directory for imgpkg lock file
	err = os.Mkdir(filepath.Join(wd, lockOutputFolder), os.ModePerm)
	if err != nil {
		return err
	}
	defer os.RemoveAll(filepath.Join(wd, lockOutputFolder))

	imgpkgLockPath := filepath.Join(wd, lockOutputFolder, "images.yml")
	cmdRunner := NewReleaseCmdRunner(os.Stdout, o.debug, imgpkgLockPath)
	reconciler := cmdlocal.NewReconciler(o.depsFactory, cmdRunner, o.logger)

	err = reconciler.Reconcile(buildConfigs, cmdlocal.ReconcileOpts{
		Local:     true,
		KbldBuild: true,
	})

	if err != nil {
		return err
	}

	var imgpkgBundleURL string

	switch {
	case pkgRepoBuild.Spec.Export.ImgpkgBundle != nil:
		imgpkgOutput, err := ImgpkgRunner{
			Image:             pkgRepoBuild.Spec.Export.ImgpkgBundle.Image,
			Version:           o.pkgRepoVersion,
			Paths:             []string{"packages"},
			UseKbldImagesLock: true,
			ImgLockFilepath:   imgpkgLockPath,
		}.Run()
		if err != nil {
			return err
		}
		imgpkgBundleURL, err = o.imgpkgBundleURLFromStdout(imgpkgOutput)
		if err != nil {
			return err
		}
	}
	fmt.Println(imgpkgBundleURL)

	//releaseStep := NewReleaseStep(o.ui, pkgRepoLocation, pkgRepoBuild)
	//err = common.Run(releaseStep)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (o *ReleaseOptions) imgpkgBundleURLFromStdout(imgpkgStdout string) (string, error) {
	lines := strings.Split(imgpkgStdout, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Pushed") {
			line = strings.TrimPrefix(line, "Pushed")
			line = strings.Replace(line, "'", "", -1)
			line = strings.Replace(line, " ", "", -1)
			return line, nil
		}
	}
	return "", fmt.Errorf("Could not get imgpkg bundle location")
}

type ReleaseStep struct {
	ui              cmdcore.AuthoringUI
	pkgRepoLocation string
	pkgRepoBuild    *build.PackageRepoBuild
}

func NewReleaseStep(ui cmdcore.AuthoringUI, pkgRepoLocation string, pkgRepoBuild *build.PackageRepoBuild) *ReleaseStep {
	return &ReleaseStep{
		ui:              ui,
		pkgRepoLocation: pkgRepoLocation,
		pkgRepoBuild:    pkgRepoBuild,
	}
}

func (release ReleaseStep) PreInteract() error {
	//err := release.createDirectory()
	//if err != nil {
	//	return err
	//}
	return nil
}

func (release ReleaseStep) Interact() error {
	return nil
}

func (release ReleaseStep) PostInteract() error {
	return nil
}

func (release ReleaseStep) createDirectory() error {
	err := os.Mkdir(release.pkgRepoLocation, os.ModePerm)
	if err != nil {
		release.ui.PrintCmdExecutionText(fmt.Sprintf("Error creating package directory.Error is: %s",
			err))
		return err
	}
	return nil
}

func GetPackageRepositoryBuild(pkgRepoBuildFilePath string) (*build.PackageRepoBuild, error) {
	var packageRepoBuild *build.PackageRepoBuild
	exists, err := common.IsFileExists(pkgRepoBuildFilePath)
	if err != nil {
		return nil, err
	}

	if exists {
		packageRepoBuild, err = NewPackageRepoBuildFromFile(pkgRepoBuildFilePath)
		if err != nil {
			return nil, err
		}
	} else {
		packageRepoBuild = &build.PackageRepoBuild{
			TypeMeta: metav1.TypeMeta{
				Kind:       build.PkgRepoBuildKind,
				APIVersion: build.PkgRepoBuildAPIVersion,
			},
			ObjectMeta: metav1.ObjectMeta{},
			Spec: build.PackageRepoBuildSpec{
				Export: &build.PackageRepoBuildExport{
					ImgpkgBundle: &build.PackageRepoBuildExportImgpkgBundle{},
				},
			},
		}
	}
	return packageRepoBuild, nil
}

func NewPackageRepoBuildFromFile(filePath string) (*build.PackageRepoBuild, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var packageRepoBuild build.PackageRepoBuild
	err = yaml.Unmarshal(content, &packageRepoBuild)
	if err != nil {
		return nil, err
	}
	return &packageRepoBuild, nil
}

func GetPkgRepoLocation() (string, error) {
	//pwd, _ := os.Getwd()
	//pkgRepoLocation, err := filepath.Rel(pwd, filepath.Join(pwd, "bundle"))
	//if err != nil {
	//	return "", err
	//}
	//return pkgRepoLocation, nil

	exists, err := common.IsFileExists("packages")
	if err != nil {
		return "", err
	}

	if !exists {
		err := os.Mkdir("packages", os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return "packages", nil
}
