package kinkgo

import "fmt"

type Environment interface {
	// Description returns a description of the environment.
	Description() string

	// Start the environment
	Start() error

	// Stop the environment
	Stop() error
}

var _ Environment = new(NopEnvironment)

// NopEnvironment is a no-op environment.
type NopEnvironment struct{}

func (n NopEnvironment) Description() string {
	return "no-op environment"
}

func (n NopEnvironment) Start() error {
	// no-op environment. no need to start anything.
	fmt.Println("no-op environment started")
	return nil
}

func (n NopEnvironment) Stop() error {
	// no-op environment. no need to stop anything.
	fmt.Println("no-op environment stopped")
	return nil
}
