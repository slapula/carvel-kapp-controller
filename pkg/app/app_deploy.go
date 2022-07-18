// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"fmt"
	"sigs.k8s.io/yaml"

	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	ctldep "github.com/vmware-tanzu/carvel-kapp-controller/pkg/deploy"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/exec"
)

func (a *App) deploy(tplOutput string, changedFunc func(exec.CmdRunResult)) exec.CmdRunResult {
	err := a.blockDeletion()
	if err != nil {
		return exec.NewCmdRunResultWithErr(fmt.Errorf("Blocking for deploy: %s", err))
	}

	if len(a.app.Spec.Deploy) != 1 {
		return exec.NewCmdRunResultWithErr(fmt.Errorf("Expected exactly one deploy option"))
	}

	var result exec.CmdRunResult

	for _, dep := range a.app.Spec.Deploy {
		switch {
		case dep.Kapp != nil:
			cancelCh, closeCancelCh := a.newCancelCh(onCanceled, onDeleted)
			defer closeCancelCh()

			kapp, err := a.newKapp(*dep.Kapp, cancelCh)
			if err != nil {
				return exec.NewCmdRunResultWithErr(fmt.Errorf("Preparing kapp: %s", err))
			}

			result = kapp.Deploy(tplOutput, a.startFlushingAllStatusUpdates, changedFunc)

		default:
			result.AttachErrorf("%s", fmt.Errorf("Unsupported way to deploy"))
		}

		if result.Error != nil {
			break
		}
	}

	return result
}

func (a *App) delete(changedFunc func(exec.CmdRunResult)) exec.CmdRunResult {
	if len(a.app.Spec.Deploy) != 1 {
		return exec.NewCmdRunResultWithErr(fmt.Errorf("Expected exactly one deploy option"))
	}

	var result exec.CmdRunResult
	if !a.app.Spec.NoopDelete {
		for _, dep := range a.app.Spec.Deploy {
			switch {
			case dep.Kapp != nil:
				cancelCh, closeCancelCh := a.newCancelCh(onCanceled)
				defer closeCancelCh()

				kapp, err := a.newKapp(*dep.Kapp, cancelCh)
				if err != nil {
					return exec.NewCmdRunResultWithErr(fmt.Errorf("Preparing kapp: %s", err))
				}

				result = kapp.Delete(a.startFlushingAllStatusUpdates, changedFunc)

			default:
				result.AttachErrorf("%s", fmt.Errorf("Unsupported way to delete"))
			}

			if result.Error != nil {
				break
			}
		}
	}

	if result.Error == nil {
		err := a.unblockDeletion()
		if err != nil {
			return exec.NewCmdRunResultWithErr(fmt.Errorf("Unblocking for deploy: %s", err))
		}
	}

	return result
}

func (a *App) inspect() exec.CmdRunResult {
	if len(a.app.Spec.Deploy) != 1 {
		return exec.NewCmdRunResultWithErr(fmt.Errorf("Expected exactly one deploy option"))
	}

	var result exec.CmdRunResult

	for _, dep := range a.app.Spec.Deploy {
		switch {
		case dep.Kapp != nil:
			if dep.Kapp.Inspect != nil {
				cancelCh, closeCancelCh := a.newCancelCh(onCanceled, onDeleted)
				defer closeCancelCh()

				kapp, err := a.newKapp(*dep.Kapp, cancelCh)
				if err != nil {
					return exec.NewCmdRunResultWithErr(fmt.Errorf("Preparing kapp: %s", err))
				}

				result = kapp.Inspect()
			}

		default:
			result.AttachErrorf("%s", fmt.Errorf("Unsupported way to inspect"))
		}

		if result.Error != nil {
			break
		}
	}

	return result
}

type AppMeta struct {
	LabelKey   string `yaml:"labelKey"`
	LabelValue string `yaml:"labelValue"`
	LastChange struct {
		Namespaces []string `yaml:"namespaces"`
	} `yaml:"lastChange"`
	UsedGKs []struct {
		Group string `yaml:"Group"`
		Kind  string `yaml:"Kind"`
	} `yaml:"usedGKs"`
}

func (a *App) GetMetadata() (AppMeta, error) {

	for _, dep := range a.app.Spec.Deploy {
		switch {
		case dep.Kapp != nil:
			cancelCh, closeCancelCh := a.newCancelCh(onCanceled, onDeleted)
			defer closeCancelCh()

			kapp, err := a.newKapp(*dep.Kapp, cancelCh)
			if err != nil {
				return AppMeta{}, err
			}

			cm, err := kapp.InternalAppConfigMap()
			if err != nil {
				return AppMeta{}, err
			}

			var appMetadata AppMeta
			err = yaml.Unmarshal([]byte(cm.Data["spec"]), &appMetadata)
			if err != nil {
				return AppMeta{}, err
			}

			return appMetadata, err
		default:
			return AppMeta{}, fmt.Errorf("Unsupported way to get config map")
		}

	}

	return AppMeta{}, fmt.Errorf("Unsupported way to get config map")
}

func (a *App) newKapp(kapp v1alpha1.AppDeployKapp, cancelCh chan struct{}) (*ctldep.Kapp, error) {
	genericOpts := ctldep.GenericOpts{Name: a.app.Name, Namespace: a.app.Namespace}

	return a.deployFactory.NewKapp(kapp, a.app.Spec.ServiceAccountName,
		a.app.Spec.Cluster, genericOpts, cancelCh)
}

type cancelCondition func(v1alpha1.App) bool

func onCanceled(app v1alpha1.App) bool {
	return app.Spec.Canceled
}

func onDeleted(app v1alpha1.App) bool {
	return app.DeletionTimestamp != nil
}
