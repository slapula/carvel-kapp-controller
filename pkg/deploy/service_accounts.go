// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package deploy

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/satoken"
	authenticationv1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ServiceAccounts struct {
	coreClient   kubernetes.Interface
	log          logr.Logger
	tokenManager *satoken.Manager
	caCert       []byte
	caCertMutex  sync.Mutex
}

// NewServiceAccounts provides access to the ServiceAccount Resource in kubernetes
func NewServiceAccounts(coreClient kubernetes.Interface, log logr.Logger) *ServiceAccounts {
	tokenMgr := satoken.NewManager(coreClient, log)
	return &ServiceAccounts{coreClient: coreClient, log: log, tokenManager: tokenMgr}
}

func (s *ServiceAccounts) Find(genericOpts GenericOpts, saName string) (ProcessedGenericOpts, error) {
	kubeconfigYAML, err := s.fetchServiceAccount(genericOpts.Namespace, saName)
	if err != nil {
		return ProcessedGenericOpts{}, err
	}

	kubeconfigRestricted, err := NewKubeconfigRestricted(kubeconfigYAML)
	if err != nil {
		return ProcessedGenericOpts{}, err
	}

	pgoForSA := ProcessedGenericOpts{
		Name:       genericOpts.Name,
		Namespace:  "", // Assume kubeconfig contains preferred namespace from SA
		Kubeconfig: kubeconfigRestricted,
	}

	return pgoForSA, nil
}

// GetClusterVersion returns the kubernetes API version for the cluster which has been supplied to kapp-controller via a kubeconfig
func GetClusterVersion(cc kubernetes.Interface, saName string, specCluster *v1alpha1.AppCluster, objMeta *metav1.ObjectMeta, log logr.Logger) (*version.Info, error) {
	// this logic is duplicated here and in ProcessOpts: if the serviceAccount name is present, we will be deploying to this cluster
	if len(saName) > 0 {
		return cc.Discovery().ServerVersion()
	}

	processedGenericOpts, err := processOpts(
		saName,
		specCluster,
		GenericOpts{Name: objMeta.Name, Namespace: objMeta.Namespace},
		NewServiceAccounts(cc, log),
		NewKubeconfigSecrets(cc))
	if err != nil {
		return nil, err
	}

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(processedGenericOpts.Kubeconfig.AsYAML()))
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	vi, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return nil, err
	}

	return vi, nil
}

// processOpts takes generic opts and a ServiceAccount Name, and returns a populated kubeconfig that can connect to a cluster.
// if the saName is empty then you'll connect to a cluster via the clusterOpts inside the genericOpts, otherwise you'll use the specified SA.
func processOpts(saName string, clusterOpts *v1alpha1.AppCluster, genericOpts GenericOpts, serviceAccounts *ServiceAccounts, kubeconfigSecrets *KubeconfigSecrets) (*ProcessedGenericOpts, error) {
	var err error
	var processedGenericOpts ProcessedGenericOpts

	switch {
	case len(saName) > 0:
		processedGenericOpts, err = serviceAccounts.Find(genericOpts, saName)
		if err != nil {
			return nil, err
		}

	case clusterOpts != nil:
		processedGenericOpts, err = kubeconfigSecrets.Find(genericOpts, clusterOpts)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Expected service account or cluster specified")
	}
	return &processedGenericOpts, nil
}

func (s *ServiceAccounts) fetchServiceAccount(nsName string, saName string) (string, error) {
	const (
		caCertPath      = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
		tokenExpiration = time.Hour * 2
	)

	if len(nsName) == 0 {
		return "", fmt.Errorf("Internal inconsistency: Expected namespace name to not be empty")
	}
	if len(saName) == 0 {
		return "", fmt.Errorf("Internal inconsistency: Expected service account name to not be empty")
	}

	sa, err := s.coreClient.CoreV1().ServiceAccounts(nsName).Get(context.Background(), saName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("Getting service account: %s", err)
	}

	expiration := int64(tokenExpiration.Seconds())
	t, err := s.tokenManager.GetServiceAccountToken(sa, &authenticationv1.TokenRequest{
		Spec: authenticationv1.TokenRequestSpec{
			ExpirationSeconds: &expiration,
		},
	})
	if err != nil {
		return "", fmt.Errorf("Get service account token: %s", err)
	}

	s.caCertMutex.Lock()
	defer s.caCertMutex.Unlock()
	if len(s.caCert) == 0 {
		s.caCert, err = os.ReadFile(caCertPath)
		if err != nil {
			return "", fmt.Errorf("Read ca cert from %s: %s", caCertPath, err)
		}
	}

	return s.buildKubeconfig(t.Status.Token, nsName, s.caCert)
}

func (s *ServiceAccounts) buildKubeconfig(token string, nsBytes string, caCert []byte) (string, error) {
	const kubeconfigYAMLTpl = `
apiVersion: v1
kind: Config
clusters:
- name: dst-cluster
  cluster:
    certificate-authority-data: "%[1]s"
    server: https://${KAPP_KUBERNETES_SERVICE_HOST_PORT}
users:
- name: dst-user
  user:
    token: "%[2]s"
contexts:
- name: dst-ctx
  context:
    cluster: dst-cluster
    namespace: "%[3]s"
    user: dst-user
current-context: dst-ctx
`
	caB64Encoded := base64.StdEncoding.EncodeToString(caCert)

	return fmt.Sprintf(kubeconfigYAMLTpl, caB64Encoded, []byte(token), []byte(nsBytes)), nil
}

/*

Example SA + secret:

---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: app1-sa
  namespace: app1
secrets:
- name: app1-sa-token-grr7z
---
apiVersion: v1
kind: Secret
metadata:
  name: app1-sa-token-grr7z
  namespace: app1
  annotations:
    kubernetes.io/service-account.name: app1-sa
    kubernetes.io/service-account.uid: 26675b19-769a-4145-a386-7ca2b3ab3435
type: kubernetes.io/service-account-token
data:
  ca.crt: LS0tLS...
  namespace: a2FwcC1jb250cm9sbGVy
  token: ZXlKaGJ...

*/
