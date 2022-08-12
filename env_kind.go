package kinkgo

import (
	"fmt"

	"github.com/kinkgo/kinkgo/clusters/kind"
)

var _ Environment = new(KindEnvironment)

// KindEnvironment is holds the environment for a Kind cluster.
type KindEnvironment struct {
	cluster *kind.Cluster
}

// NewKindEnvironment returns a new KindEnvironment.
func NewKindEnvironment(clusterName string) KindEnvironment {
	return KindEnvironment{
		cluster: kind.NewCluster(
			clusterName,
			kind.CreateWithDisplayUsage(false),
			kind.CreateWithDisplaySalutation(false),
		),
	}
}

func (k KindEnvironment) Description() string {
	return "single cluster kind environment"
}

// KubeConfig returns the kubeconfig for the cluster.
func (k KindEnvironment) KubeConfig() (string, error) {
	return k.cluster.KubeConfig(false)
}

func (k KindEnvironment) Start() error {

	fmt.Printf("creating kind cluster %s\n", k.cluster.Name)

	return k.cluster.Create()
}

func (k KindEnvironment) Stop() error {

	fmt.Printf("deleting kind cluster %s\n", k.cluster.Name)

	return k.cluster.Delete()
}
