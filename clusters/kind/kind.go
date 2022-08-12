package kind

import (
	"fmt"
	"os"

	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cmd"

	"github.com/anandvarma/namegen"
)

type Cluster struct {
	Name     string
	options  *options
	provider *cluster.Provider
}

func NewCluster(name string, opts ...CreateOption) *Cluster {
	// Use the default options if none are provided
	if name == "" {
		ngen := namegen.New()
		name = ngen.Get()
	}

	// Apply the options
	o := new(options)

	for _, opt := range opts {
		opt(o)
	}

	// Initialize the cluster object
	return &Cluster{
		Name:     name,
		options:  o,
		provider: cluster.NewProvider(cluster.ProviderWithLogger(cmd.NewLogger())),
	}
}

// Create creates a new kind cluster
func (c *Cluster) Create() error {
	return c.provider.Create(c.Name, c.options.clusterCreateOptions...)
}

// KubeConfig returns the kubeconfig for the cluster
func (c *Cluster) KubeConfig(internal bool) (string, error) {
	return c.provider.KubeConfig(c.Name, internal)
}

// Delete deletes an existing kind cluster
func (c *Cluster) Delete() error {
	// Create temporary file to export KubeConfig
	// This is necessary to not make sure we're using the right KubeConfig
	tmp, err := os.CreateTemp("", fmt.Sprintf("%s-%s", c.Name, "kubeconfig"))
	if err != nil {
		return err
	}

	// Make sure remove the temporary file after we're done
	defer os.Remove(tmp.Name())

	// Export the KubeConfig
	if err := c.provider.ExportKubeConfig(c.Name, tmp.Name(), false); err != nil {
		return err
	}

	// delete the cluster
	return c.provider.Delete(c.Name, tmp.Name())
}
