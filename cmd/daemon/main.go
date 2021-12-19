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
	fmt.Println("Hello, daemon")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	run(defaults.RebootChecker{}, defaults.BlockChecker{}, defaults.Rebooter{}, ctx)
}

func run(rc smartautoreboot.RebootChecker, bc smartautoreboot.BlockChecker, r smartautoreboot.Rebooter, ctx context.Context) {
mainLoop:
	for {
		select {
		case <-ctx.Done():
			break mainLoop
		case <-time.After(time.Minute):
		}

	}
}
