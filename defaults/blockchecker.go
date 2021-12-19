package defaults

type BlockChecker struct{}

func (b BlockChecker) IsRebootBlocked() bool {
	return false
}
