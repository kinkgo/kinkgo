package simple_kindenv_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kinkgo/kinkgo"
)

var (
	// Suite
	_ kinkgo.TestSuite = new(TestSuite)

	// Environment
	env kinkgo.KindEnvironment
)

func TestSimpleKindEnvTest(t *testing.T) {
	env = kinkgo.NewKindEnvironment("kind-test")

	// Define test configuration
	cfg := kinkgo.NewConfig(env, "simple_test", "extra_label")

	// Run test suite
	kinkgo.Run(t, "SimpleTest Suite", new(TestSuite), cfg)
}

type TestSuite struct{}

var _ = Describe("Kind environment", func() {
	It("should provide a raw kubeconfig", Label("pass"), func() {
		Expect(env.KubeConfig()).ShouldNot(BeEmpty())
	})
})
