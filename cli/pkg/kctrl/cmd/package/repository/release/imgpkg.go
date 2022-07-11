package release

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	goexec "os/exec"
	"path/filepath"
	"strings"
)

type ImgpkgRunner struct {
	Image             string
	Version           string
	Paths             []string
	UseKbldImagesLock bool
	ImgLockFilepath   string
}

func (r ImgpkgRunner) Run() (string, error) {
	dir, err := os.MkdirTemp(".", fmt.Sprintf("bundle-%s-*", strings.Replace(r.Image, "/", "-", 1)))
	if err != nil {
		return "", err
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for _, path := range r.Paths {
		var stderrBuf bytes.Buffer
		cmd := goexec.Command("cp", "-r", filepath.Join(wd, path), dir)
		cmd.Stderr = &stderrBuf
		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("%s", stderrBuf.String())
		}
	}
	if r.UseKbldImagesLock {
		err = goexec.Command("cp", "-r", r.ImgLockFilepath, filepath.Join(dir, lockOutputFolder)).Run()
		if err != nil {
			return "", err
		}
	}
	defer os.RemoveAll(dir)

	pushLocation := fmt.Sprintf("%s:%s", r.Image, r.Version)
	var stdoutBuf, stderrBuf bytes.Buffer
	inMemoryStdoutWriter := bufio.NewWriter(&stdoutBuf)
	cmd := goexec.Command("imgpkg", "push", "-b", pushLocation, "-f", dir, "--tty=true")
	// TODO: Switch to using Authoring UI
	fmt.Printf("Running: %s", strings.Join(cmd.Args, " "))
	cmd.Stdout = io.MultiWriter(os.Stdout, inMemoryStdoutWriter)
	cmd.Stderr = &stderrBuf
	err = cmd.Run()
	if err != nil {
		return stdoutBuf.String(), fmt.Errorf("%s", stderrBuf.String())
	}
	err = inMemoryStdoutWriter.Flush()
	if err != nil {
		return "", err
	}

	return stdoutBuf.String(), nil
}

//type ImgPkgStep struct {
//	ui              cmdcore.AuthoringUI
//	pkgRepoLocation string
//	pkgRepoBuild    *build.PackageRepoBuild
//}
//
//const (
//	lockOutputFolder = ".imgpkg"
//)
//
//func NewImgPkgStep(ui cmdcore.AuthoringUI, pkgLocation string, pkgRepoBuild *build.PackageRepoBuild) *ImgPkgStep {
//	return &ImgPkgStep{
//		ui:              ui,
//		pkgRepoLocation: pkgLocation,
//		pkgRepoBuild:    pkgRepoBuild,
//	}
//}
//
//func (imgPkg ImgPkgStep) PreInteract() error {
//	imgPkg.ui.PrintInformationalText("To create package repository, we will create an imgpkg bundle first. Imgpkg, a Carvel tool, allows users to package, distribute, and relocate a set of files as one OCI artifact: a bundle. Imgpkg bundles are identified with a unique sha256 digest based on the file contents. Imgpkg uses that digest to ensure that the copied contents are identical to those originally pushed.")
//	imgPkg.ui.PrintInformationalText("\nA package repository bundle is an imgpkg bundle that holds PackageMetadata and Package CRs. Later on, this bundle can be mentioned in the package repository CR to fetch the package and packageMetadata CRs.")
//	imgPkg.ui.PrintActionableText("\nCreating the required directory structure for imgpkg bundle\n")
//
//	return nil
//}
//
//func (imgPkg ImgPkgStep) Interact() error {
//	imgPkg.ui.PrintHeaderText("\nRegistry URL(Step 2/3)")
//	defaultRegistryURL := imgPkg.pkgRepoBuild.Spec.Export.ImgpkgBundle.Image
//	textOpts := ui.TextOpts{
//		Label:        "Enter the registry url to push the package repository bundle",
//		Default:      defaultRegistryURL,
//		ValidateFunc: nil,
//	}
//	registryURL, err := imgPkg.ui.AskForText(textOpts)
//	if err != nil {
//		return err
//	}
//
//	imgPkg.pkgRepoBuild.Spec.Export.ImgpkgBundle.Image = registryURL
//	imgPkg.pkgRepoBuild.WriteToFile()
//	return nil
//}

//func (imgPkg ImgPkgStep) PostInteract() error {
//	wd, err := os.Getwd()
//	if err != nil {
//		return err
//	}
//
//	// Create temporary directory for imgpkg lock file
//	err = os.Mkdir(filepath.Join(wd, lockOutputFolder), os.ModePerm)
//	if err != nil {
//		return err
//	}
//	defer os.RemoveAll(filepath.Join(wd, lockOutputFolder))
//
//	imgpkgLockPath := filepath.Join(wd, lockOutputFolder, "images.yml")
//
//	imgPkg.ui.PrintInformationalText("\nLet's use `kbld` to create immutable image reference. Kbld scans all the files in bundle configuration for any references of images and creates a mapping of image tags to a URL with sha256 digest.")
//	imgPkg.ui.PrintActionableText("Lock image references using Kbld")
//	imgPkg.ui.PrintCmdExecutionText(fmt.Sprintf("kbld --file %s --imgpkg-lock-output %s", imgPkg.pkgRepoLocation, imgpkgLockPath))
//
//	//cmdRunner := NewReleaseCmdRunner(os.Stdout, o.debug, imgpkgLockPath)
//	//reconciler := cmdlocal.NewReconciler(imgPkg.depsFactory, cmdRunner, o.logger)
//	//err = runningKbld(imgPkg.pkgRepoLocation, imgpkgLockPath)
//	return nil
//}

func runningKbld(bundleLocation, imagesFileLocation string) error {
	//result := util.Execute("kbld", []string{"--file", bundleLocation, "--imgpkg-lock-output", imagesFileLocation})
	//if result.Error != nil {
	//	return fmt.Errorf("Running kbld.\n %s", result.Stderr)
	//}
	return nil
}
