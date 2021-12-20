package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/BrenekH/logange"
	smartautoreboot "github.com/BrenekH/smart-auto-reboot"
	"github.com/BrenekH/smart-auto-reboot/defaults"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	conf, err := smartautoreboot.ParseConfFile("/etc/smartreboot/smartreboot.conf")
	if err != nil {
		panic(err)
	}

	mainLogger := logange.NewLogger("main")
	stdoutHandler := logange.NewStdoutHandler()
	stdoutHandler.SetLevel(conf.LogLevel)
	mainLogger.AddHandler(&stdoutHandler)

	run(conf.CheckInterval, defaults.RebootChecker{}, defaults.BlockChecker{}, defaults.Rebooter{}, mainLogger, ctx)
}

func run(waitMinutes int, rc smartautoreboot.RebootChecker, bc smartautoreboot.BlockChecker, r smartautoreboot.Rebooter, logger logange.Logger, ctx context.Context) {
	logger.Info("smartrebootd started")

mainLoop:
	for {
		if rc.IsRebootRequired() && !bc.IsRebootBlocked() {
			if err := r.Reboot(); err != nil {
				panic(err)
			}
			break
		}

		select {
		case <-ctx.Done():
			break mainLoop
		case <-time.After(time.Minute * time.Duration(waitMinutes)):
		}
	}
}
