package defaults

type BlockChecker struct{}

func (b BlockChecker) IsRebootBlocked() bool {
	codes := runScriptsInDir("/etc/smartreboot/blockchecks")

	for _, exitCode := range codes {
		if exitCode != 0 {
			return true
		}
	}

	return false
}
