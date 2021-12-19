package defaults

type RebootChecker struct{}

// IsRebootRequired executes all scripts in /etc/smartreboot/rebootchecks returning
// true if at least one returned an exit code of 0, false if all were non-zero.
func (r RebootChecker) IsRebootRequired() bool {
	codes := runScriptsInDir("/etc/smartreboot/rebootchecks")

	for _, exitCode := range codes {
		if exitCode == 0 {
			return true
		}
	}

	return false
}
