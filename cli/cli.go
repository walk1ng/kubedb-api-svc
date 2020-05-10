package cli

import (
	rest "k8s.io/client-go/rest"
	kubedb "kubedb.dev/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1"
)

var (
	config *rest.Config
	err    error
	// KubedbClientSet is client set for kubedb
	KubedbClientSet *kubedb.KubedbV1alpha1Client
)

func init() {
	config, err = rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	KubedbClientSet, err = kubedb.NewForConfig(config)
}
