// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package appsecrets

import (
	"context"
	"strconv"

	"github.com/go-logr/logr"
	kcv1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned/scheme"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type DefaultSecrets struct {
	coreClient kubernetes.Interface
	log        logr.Logger
}

func NewDefaultSecrets(coreClient kubernetes.Interface, log logr.Logger) *DefaultSecrets {
	return &DefaultSecrets{coreClient, log}
}

func (as *DefaultSecrets) AttachAndReconcile(app *kcv1alpha1.App) error {
	const (
		defaultSecretsAnnKey = "kappctrl.k14s.io/default-secrets"
	)

	// No need to reconcile if app is being deleted
	if app.DeletionTimestamp != nil {
		return nil
	}

	if _, ok := app.Annotations[defaultSecretsAnnKey]; !ok {
		// TODO currently does not delete previously created secrets
		return nil
	}

	var imagePullSecretNames []string

	for i, fetchStep := range app.Spec.Fetch {
		ref := corev1.LocalObjectReference{Name: app.Name + "-fetch-" + strconv.Itoa(i)}

		switch {
		case fetchStep.Image != nil:
			fetchStep.Image.SecretRef = &kcv1alpha1.AppFetchLocalRef{LocalObjectReference: ref}
			imagePullSecretNames = append(imagePullSecretNames, ref.Name)
		case fetchStep.ImgpkgBundle != nil:
			fetchStep.ImgpkgBundle.SecretRef = &kcv1alpha1.AppFetchLocalRef{LocalObjectReference: ref}
			imagePullSecretNames = append(imagePullSecretNames, ref.Name)
		}

		app.Spec.Fetch[i] = fetchStep
	}

	return as.reconcile(imagePullSecretNames, app)
}

func (as *DefaultSecrets) reconcile(imagePullSecretNames []string, app *kcv1alpha1.App) error {
	for _, name := range imagePullSecretNames {
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: app.Namespace,
				Annotations: map[string]string{
					// TODO use secretgen constant?
					"secretgen.carvel.dev/image-pull-secret": "",
				},
			},
			Type: corev1.SecretTypeDockerConfigJson,
			Data: map[string][]byte{
				corev1.DockerConfigJsonKey: []byte("{}"),
			},
		}

		if len(app.GetUID()) == 0 {
			panic("Internal inconsistency: Expected app to have a UID")
		}

		controllerutil.SetControllerReference(app, secret, scheme.Scheme)

		_, err := as.coreClient.CoreV1().Secrets(secret.Namespace).Create(
			context.Background(), secret, metav1.CreateOptions{})
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				return err
			}
		}
	}

	return nil
}
