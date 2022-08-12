package kind

import (
	"sigs.k8s.io/kind/pkg/cluster"
)

type options struct {
	// extra kind create options to pass to the cluster provider
	clusterCreateOptions []cluster.CreateOption
}

// CreateOption custom type for configuring optional parameters.
type CreateOption func(o *options)

// CreateWithDisplayUsage enables displaying usage if displayUsage is true
func CreateWithDisplayUsage(displayUsage bool) CreateOption {
	return func(o *options) {
		o.clusterCreateOptions = append(o.clusterCreateOptions, cluster.CreateWithDisplayUsage(displayUsage))
	}
}

// CreateWithDisplaySalutation enables display a salutation at the end of create
// cluster if displaySalutation is true
func CreateWithDisplaySalutation(displaySalutation bool) CreateOption {
	return func(o *options) {
		o.clusterCreateOptions = append(o.clusterCreateOptions, cluster.CreateWithDisplaySalutation(displaySalutation))
	}
}
