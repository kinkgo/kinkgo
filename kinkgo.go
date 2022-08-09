package kinkgo

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

// Run is a simple wrapper around ginkgo.RunSpecs that configures the runner for the kinkgo framework.
func Run(t *testing.T, description string) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	ginkgo.RunSpecs(t, description)
}
