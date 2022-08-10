package custom_env_testsuite_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kinkgo/kinkgo"
)

var (
	// Suite
	_ kinkgo.TestSuite = new(SimpleTestSuite)

	// Environment
	_ kinkgo.Environment = new(CustomEnvironment)
)

func TestCustomEnvTestsuite(t *testing.T) {
	// Define test configuration
	cfg := kinkgo.NewConfig(CustomEnvironment{}, "simple_test", "extra_label")

	// Run test suite
	kinkgo.Run(t, "SimpleTest Suite", new(SimpleTestSuite), cfg)
}

type SimpleTestSuite struct{}

type CustomEnvironment struct {
	StartTime time.Time
	StopTime  time.Time
}

func (c CustomEnvironment) Description() string {
	return "custom environment"
}

func (c CustomEnvironment) Start() error {
	c.StartTime = time.Now()

	fmt.Printf("custom environment started at %s\n", c.StartTime)

	return nil
}

func (c CustomEnvironment) Stop() error {
	c.StopTime = time.Now()

	fmt.Printf("custom environment stopped at %s\n", c.StopTime)

	return nil
}

var _ = Describe("SimpleTest", func() {
	It("should pass", Label("pass"), func() {
		Expect(1).To(Equal(1))
	})
})
