package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	smartautoreboot "github.com/BrenekH/smart-auto-reboot"
	"github.com/BrenekH/smart-auto-reboot/defaults"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	run(1, defaults.RebootChecker{}, defaults.BlockChecker{}, defaults.Rebooter{}, ctx)
}

func run(waitMinutes int, rc smartautoreboot.RebootChecker, bc smartautoreboot.BlockChecker, r smartautoreboot.Rebooter, ctx context.Context) {
	fmt.Println("smartrebootd started")

mainLoop:
	for {
		select {
		case <-ctx.Done():
			break mainLoop
		case <-time.After(time.Minute * time.Duration(waitMinutes)):
		}

		if rc.IsRebootRequired() && !bc.IsRebootBlocked() {
			if err := r.Reboot(); err != nil {
				panic(err)
			}
		}
	}
}
