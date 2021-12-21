package smartreboot

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/BrenekH/logange"
)

var (
	checkIntervalRe *regexp.Regexp
	logLevelRe      *regexp.Regexp
)

func init() {
	// Create Regexes and exit (panic) if they fail to compile.
	// This allows them to be used for every regex search instead of needing to recompile everytime.
	var err error
	checkIntervalRe, err = regexp.Compile(`^*CheckInterval=(.*)`)
	if err != nil {
		panic(err)
	}

	logLevelRe, err = regexp.Compile(`^LogLevel=(.*)`)
	if err != nil {
		panic(err)
	}
}

func ParseConfFile(filename string) (c Conf, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return
	}

	contents := string(b)

	checkIntervalMatch := checkIntervalRe.FindStringSubmatch(contents)
	if len(checkIntervalMatch) > 0 {
		interval, err := strconv.Atoi(checkIntervalMatch[1])
		if err != nil {
			return c, err
		}

		c.CheckInterval = interval
	} else {
		c.CheckInterval = 1
	}

	logLevelMatch := logLevelRe.FindStringSubmatch(contents)
	if len(logLevelMatch) > 0 {
		switch strings.ToLower(logLevelMatch[1]) {
		case "trace":
			c.LogLevel = logange.LevelTrace
		case "debug":
			c.LogLevel = logange.LevelDebug
		case "info":
			c.LogLevel = logange.LevelInfo
		case "warning", "warn":
			c.LogLevel = logange.LevelWarn
		case "error":
			c.LogLevel = logange.LevelWarn
		case "critical":
			c.LogLevel = logange.LevelCritical
		default:
			return c, fmt.Errorf("unknown log level '%v'", logLevelMatch[1])
		}
	} else {
		c.LogLevel = logange.LevelWarn
	}

	return
}
