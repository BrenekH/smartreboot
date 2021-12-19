package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Hello, daemon")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

mainLoop:
	for {
		select {
		case <-ctx.Done():
			break mainLoop
		case <-time.After(time.Minute):
		}

	}

	fmt.Println("All done")
}
