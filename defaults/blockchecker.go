package defaults

import (
	"github.com/BrenekH/logange"
)

type BlockChecker struct {
	Logger logange.Logger
}

// IsRebootBlocked returns true if any of the executable scripts in /etc/smartreboot/blockchecks
// returns a non-zero exit code, otherwise false is returned.
func (b BlockChecker) IsRebootBlocked() bool {
	b.Logger.Info("Checking for blocks")

	codes := runScriptsInDir("/etc/smartreboot/blockchecks")

	for _, exitCode := range codes {
		if exitCode != 0 {
			return true
		}
	}

	return false
}
