package kinkgo

// EnvironmentPreStartHook is a hook that executed before the kinkgo framework is run to start the environment.
type EnvironmentPreStartHook[E Environment] interface {
	EnvironmentPreStart(env *E)
}

// EnvironmentPostStartHook is a hook that executed after the kinkgo framework is run to stop the environment.
type EnvironmentPostStartHook[E Environment] interface {
	EnvironmentPostStart(env *E)
}

// EnvironmentPreStopHook is a hook that executed before the kinkgo framework is run to stop the environment.
type EnvironmentPreStopHook[E Environment] interface {
	EnvironmentPreStop(env *E)
}

// EnvironmentPostStopHook is a hook that executed after the kinkgo framework is run to stop the environment.
type EnvironmentPostStopHook[E Environment] interface {
	EnvironmentPostStop(env *E)
}
