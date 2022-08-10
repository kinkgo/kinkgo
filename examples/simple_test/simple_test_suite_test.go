package simple_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kinkgo/kinkgo"
	kinkgohooks "github.com/kinkgo/kinkgo/hooks"
	kinkgotypes "github.com/kinkgo/kinkgo/types"
)

var (
	// Suite
	_ kinkgo.TestSuite = new(SimpleTestSuite)

	// Hooks
	_ kinkgohooks.ModifySuiteConfigHook    = new(SimpleTestSuite)
	_ kinkgohooks.ModifyReporterConfigHook = new(SimpleTestSuite)
)

func TestSimpleTest(t *testing.T) {
	// Define test configuration
	cfg := kinkgo.Config{
		Labels: []string{"simple_test", "extra_label"},
	}

	// Run test suite
	kinkgo.Run(t, "SimpleTest Suite", new(SimpleTestSuite), cfg)
}

type SimpleTestSuite struct{}

func (s *SimpleTestSuite) ModifyReporterConfig(cfg kinkgotypes.ReporterConfig) kinkgotypes.ReporterConfig {
	// modify reporter config to use verbose output
	cfg.Verbose = true

	return cfg
}

func (s *SimpleTestSuite) ModifySuiteConfig(cfg kinkgotypes.SuiteConfig) kinkgotypes.SuiteConfig {
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
