package kinkgo

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	ginkgotypes "github.com/onsi/ginkgo/v2/types"
	"github.com/onsi/gomega"
)

// TestSuite is a simple marker interface that is used to identify test suites.
type TestSuite interface {
	// Intentionally left blank.
}

// Run is a simple wrapper around ginkgo.RunSpecs that configures the runner for the kinkgo framework.
func Run(t *testing.T, description string, suite TestSuite, cfg Config) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	suiteCfg, reporterCfg := ginkgo.GinkgoConfiguration()

	// hook definitions

	// modify suite config
	if hook, ok := suite.(ModifySuiteConfigHook); ok {
		suiteCfg = ginkgotypes.SuiteConfig(hook.ModifySuiteConfig(SuiteConfig(suiteCfg)))
	}

	// modify reporter config
	if hook, ok := suite.(ModifyReporterConfigHook); ok {
		reporterCfg = ginkgotypes.ReporterConfig(hook.ModifyReporterConfig(ReporterConfig(reporterCfg)))
	}

	ginkgo.RunSpecs(t, description, ginkgo.Label(cfg.Labels...), suiteCfg, reporterCfg)
}
