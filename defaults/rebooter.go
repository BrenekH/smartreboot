package defaults

type Rebooter struct{}

func (r Rebooter) Reboot() error {
	return nil
}
