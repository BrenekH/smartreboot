package defaults

import "fmt"

type BlockChecker struct{}

// IsRebootBlocked returns true if any of the executable scripts in /etc/smartreboot/blockchecks
// returns a non-zero exit code, otherwise false is returned.
func (b BlockChecker) IsRebootBlocked() bool {
	fmt.Println("Checking for blocks")

	codes := runScriptsInDir("/etc/smartreboot/blockchecks")

	for _, exitCode := range codes {
		if exitCode != 0 {
			return true
		}
	}

	return false
}
