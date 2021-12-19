package smartautoreboot

type RebootChecker interface {
	IsRebootRequired() bool
}

type BlockChecker interface {
	IsRebootBlocked() bool
}

type Rebooter interface {
	Reboot() error
}
