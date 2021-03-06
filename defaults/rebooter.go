package defaults

import (
	"os/exec"

	"github.com/BrenekH/logange"
)

type Rebooter struct {
	Logger logange.Logger
}

// Reboot "safely" restarts the system.
func (r Rebooter) Reboot() error {
	// So this could be done to work with all linux systems by using syscall.Reboot(),
	// but that makes us responsible for ensuring that all programs are shutdown properly
	// (using sync(2) apparently?). Instead of dealing with that, I'd rather, at least for
	// now, just call the shutdown command which should take care of all of that for us.

	r.Logger.Info("Rebooting")

	return exec.Command("shutdown", "-r", "now").Run()
}
