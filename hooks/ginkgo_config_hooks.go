package hooks

import "github.com/kinkgo/kinkgo/types"

// ModifySuiteConfigHook is a hook that executed before the ginkgo framework is run to update the ginkgo
// suite configuration.
type ModifySuiteConfigHook interface {
	// ModifySuiteConfig is accepts current suite config and returns a modified suite config to run the test suite.
	ModifySuiteConfig(cfg types.SuiteConfig) types.SuiteConfig
}

// ModifyReporterConfigHook is a hook that executed before the ginkgo framework is run to update the ginkgo
// reporter configuration.
type ModifyReporterConfigHook interface {
	// ModifyReporterConfig is accepts current reporter config and returns a modified reporter config to run the test suite.
	ModifyReporterConfig(cfg types.ReporterConfig) types.ReporterConfig
}
