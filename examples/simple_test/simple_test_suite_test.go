package simple_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kinkgo/kinkgo"
)

var (
	// Suite
	_ kinkgo.TestSuite = new(SimpleTestSuite)

	// Hooks
	_ kinkgo.ModifySuiteConfigHook    = new(SimpleTestSuite)
	_ kinkgo.ModifyReporterConfigHook = new(SimpleTestSuite)
)

func TestSimpleTest(t *testing.T) {
	// Define test configuration
	cfg := kinkgo.NewNopEnvironmentConfig("simple_test", "extra_label")

	// Run test suite
	kinkgo.Run(t, "SimpleTest Suite", new(SimpleTestSuite), cfg)
}

type SimpleTestSuite struct{}

func (s *SimpleTestSuite) ModifyReporterConfig(cfg kinkgo.ReporterConfig) kinkgo.ReporterConfig {
	// modify reporter config to use verbose output
	cfg.Verbose = true

	return cfg
}

func (s *SimpleTestSuite) ModifySuiteConfig(cfg kinkgo.SuiteConfig) kinkgo.SuiteConfig {
	// modify suite config to skip test has "skip" label
	cfg.LabelFilter = "!skip"

	return cfg
}

var _ = Describe("SimpleTest", func() {
	It("should pass", Label("pass"), func() {
		Expect(1).To(Equal(1))
	})

	// This test is skipped by ModifySuiteConfigHook change above.
	It("should fail", Label("skip"), func() {
		Expect(1).To(Equal(2))
	})
})
