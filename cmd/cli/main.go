package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cornfeedhobo/pflag"
)

var Version = "dev"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected at least 2 arguments")
		os.Exit(1)
	}

	// Create flag sets for commands
	rebootFlagSet := pflag.NewFlagSet("reboot", pflag.ExitOnError)

	switch strings.ToLower(os.Args[1]) {
	case "reboot":
		forceFlag := rebootFlagSet.BoolP("force", "f", false, "Ignores all block scripts when performing a reboot")
		rebootFlagSet.Parse(os.Args[2:])
		_ = forceFlag
		// TODO: Implement

	case "check":
		// TODO: Implement

	default:
		displayVer := pflag.BoolP("version", "v", false, "Show version")
		displayHelp := pflag.BoolP("help", "h", false, "Show this message")

		pflag.Parse()

		if *displayHelp {
			displayHelpMessage()
			break
		}

		if *displayVer {
			fmt.Printf("smartreboot %v\n", Version)
			break
		}
	}
}

func displayHelpMessage() {
	fmt.Printf(`Smart Reboot CLI %v Help

--help, -h    - Show this message
--version, -v - Show the application version

Commands:
    reboot (--force, -f) - Manually invoke the reboot process, using --force(-f) to ignore the block scripts.
    check                - Runs each runnable script and shows the exit codes.
`, Version)
	fmt.Println("") // Print an empty string to append a newline (the backtick style of strings doesn't allow for escape chars like '\n')
}
