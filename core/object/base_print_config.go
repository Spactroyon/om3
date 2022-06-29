package object

import (
	"opensvc.com/opensvc/core/rawconfig"
)

// OptsPrintConfig is the options of the PrintConfig object method.
type OptsPrintConfig struct {
	OptsLock
	Eval        bool   `flag:"eval"`
	Impersonate string `flag:"impersonate"`
}

// PrintConfig gets a keyword value
func (t *core) PrintConfig(options OptsPrintConfig) (rawconfig.T, error) {
	if options.Eval {
		return t.config.RawEvaluatedAs(options.Impersonate)
	}
	return t.config.Raw(), nil
}
