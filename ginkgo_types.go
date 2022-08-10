package kinkgo

import (
	ginkgotypes "github.com/onsi/ginkgo/v2/types"
)

// SuiteConfig is simply a wrapper around ginkgotypes.SuiteConfig. It exists to make it easier to use
// the ginkgo framework in kinkgo.
type SuiteConfig ginkgotypes.SuiteConfig

// ReporterConfig is simply a wrapper around ginkgotypes.ReporterConfig. It exists to make it easier to use
// the ginkgo framework in kinkgo.
type ReporterConfig ginkgotypes.ReporterConfig
