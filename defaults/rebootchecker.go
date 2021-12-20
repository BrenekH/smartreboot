package defaults

import (
	"github.com/BrenekH/logange"
)

type RebootChecker struct {
	Logger logange.Logger
}

// IsRebootRequired executes all scripts in /etc/smartreboot/rebootchecks returning
// true if at least one returned an exit code of 0, false if all were non-zero.
func (r RebootChecker) IsRebootRequired() bool {
	r.Logger.Info("Checking reboot required")

	codes := runScriptsInDir("/etc/smartreboot/rebootchecks")

	for _, exitCode := range codes {
		if exitCode == 0 {
			return true
		}
	}

	return false
}
