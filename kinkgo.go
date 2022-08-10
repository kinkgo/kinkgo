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
func Run[E Environment](t *testing.T, description string, suite TestSuite, cfg Config[E]) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	suiteCfg, reporterCfg := ginkgo.GinkgoConfiguration()

	// ginkgo configuration hook definitions

	// modify suite config
	if hook, ok := suite.(ModifySuiteConfigHook); ok {
		suiteCfg = ginkgotypes.SuiteConfig(hook.ModifySuiteConfig(SuiteConfig(suiteCfg)))
	}

	// modify reporter config
	if hook, ok := suite.(ModifyReporterConfigHook); ok {
		reporterCfg = ginkgotypes.ReporterConfig(hook.ModifyReporterConfig(ReporterConfig(reporterCfg)))
	}

	// Environment initialization

	env := cfg.Environment

	// Start Environment
	//
	// Since the environment operations are not part of the ginkgo framework, we need to validate the environment
	// using testing.T to make sure not panic ginkgo framework.
	if err := env.Start(); err != nil {
		t.Fatalf("failed to start environment: %v", err)
	}

	// Run ginkgo specs
	ginkgo.RunSpecs(t, description, ginkgo.Label(cfg.Labels...), suiteCfg, reporterCfg)

	// Stop Environment
	//
	// Since the environment operations are not part of the ginkgo framework, we need to validate the environment
	// using testing.T to make sure not panic ginkgo framework.
	if err := env.Stop(); err != nil {
		t.Fatalf("failed to stop environment: %v", err)
	}
}
