package defaults

type RebootChecker struct{}

func (r RebootChecker) IsRebootRequired() bool {
	return false
}
