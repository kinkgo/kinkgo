package kinkgo

import (
	"testing"

	"github.com/kinkgo/kinkgo/hooks"
	kinkgotypes "github.com/kinkgo/kinkgo/types"

	"github.com/onsi/ginkgo/v2"
	ginkgotypes "github.com/onsi/ginkgo/v2/types"
	"github.com/onsi/gomega"
)

// TestSuite is a simple marker interface that is used to identify test suites.
type TestSuite interface {
	// Intentionally left blank.
}

// Run is a simple wrapper around ginkgo.RunSpecs that configures the runner for the kinkgo framework.
func Run(t *testing.T, description string, suite TestSuite, labels ...string) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	suiteCfg, reporterCfg := ginkgo.GinkgoConfiguration()

	// hook definitions

	// modify suite config
	if hook, ok := suite.(hooks.ModifySuiteConfigHook); ok {
		suiteCfg = ginkgotypes.SuiteConfig(hook.ModifySuiteConfig(kinkgotypes.SuiteConfig(suiteCfg)))
	}

	// modify reporter config
	if hook, ok := suite.(hooks.ModifyReporterConfigHook); ok {
		reporterCfg = ginkgotypes.ReporterConfig(hook.ModifyReporterConfig(kinkgotypes.ReporterConfig(reporterCfg)))
	}

	ginkgo.RunSpecs(t, description, ginkgo.Label(labels...), suiteCfg, reporterCfg)
}
