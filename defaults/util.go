package defaults

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// runScriptsInDir runs all executable scripts in a directory and returns all of the exit codes for analysis
func runScriptsInDir(dir string) []int {
	allFiles := discoverFilesInDir(dir)
	filtered := filterNonExecutables(allFiles)

	codes := make([]int, 0)
	for _, v := range filtered {
		fmt.Printf("Running '%v'\n", v)

		err := exec.Command(v).Run()
		if err == nil {
			codes = append(codes, 0)
		}

		if exiterr, ok := err.(interface{ ExitCode() int }); ok {
			codes = append(codes, exiterr.ExitCode())
		} else if err != nil {
			fmt.Printf("error running script '%v': %v\n", v, err)
		}
	}

	return codes
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
