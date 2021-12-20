package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/BrenekH/smart-auto-reboot/defaults"
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

		rebootCommand(*forceFlag)

	case "check":
		checkCommand()

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

func rebootCommand(force bool) {
	if !(defaults.RebootChecker{}.IsRebootRequired()) {
		return
	}

	if !force {
		if (defaults.BlockChecker{}.IsRebootBlocked()) {
			return
		}
	}

	err := defaults.Rebooter{}.Reboot()
	if err != nil {
		panic(err)
	}
}

func checkCommand() {
	fmt.Println("Reboot Checks:")
	printCodes("/etc/smartreboot/rebootchecks")

	fmt.Println("\nBlock Checks:")
	printCodes("/etc/smartreboot/blockchecks")
}

func printCodes(dir string) {
	allFiles := discoverFilesInDir(dir)
	filtered := filterNonExecutables(allFiles)

	for _, v := range filtered {
		err := exec.Command(v).Run()
		if err == nil {
			fmt.Printf("\t%v\t0\n", v)
		}

		if exiterr, ok := err.(interface{ ExitCode() int }); ok {
			fmt.Printf("\t%v\t%v\n", v, exiterr.ExitCode())
		} else if err != nil {
			fmt.Printf("\terror running script '%v': %v\n", v, err)
		}
	}
}

func discoverFilesInDir(dir string) []string {
	cleanSlashedPath := filepath.ToSlash(filepath.Clean(dir))
	files := make([]string, 0)

	filepath.Walk(cleanSlashedPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files
}

func filterNonExecutables(toFilter []string) []string {
	filtered := make([]string, 0)

	for _, v := range toFilter {
		if isExecutable(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func isExecutable(file string) bool {
	fInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return false // Silently swallow error
	}

	return fInfo.Mode()&0111 != 0
}
