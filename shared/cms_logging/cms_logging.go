package cms_logging

import (
    "fmt"
    "log"
    "log/slog"
    "os"
    "os/user"
    "strings"
)

/*
    Constants, and such
*/

// ERROR CODES 
const (
    ERR_CODE_TEMPLATE_MODULE = iota * 1000
    ERR_CODE_ADMIN_HOME
    ERR_CODE_ADMIN_PRODUCTS
    ERR_CODE_ADMIN_VARIATIONS
    ERR_CODE_ADMIN_SECTIONS
    ERR_CODE_ADMIN_INFO
    ERR_CODE_ADMIN_TEMPLATE
    ERR_CODE_ADMIN_USER
    ERR_CODE_ADMIN_WEBSITES
)

// TERMINAL COLOR
const (
    termGreen   = "\033[97;42m"
	termWhite   = "\033[90;47m"
	termYellow  = "\033[90;43m"
	termRed     = "\033[97;41m"
	termBlue    = "\033[97;44m"
	termMagenta = "\033[97;45m"
	termCyan    = "\033[97;46m"
	termReset   = "\033[0m"
)

// TERMINAL LEVEL MSGs
const (
    levelLogMsg = termCyan + "LOG" + termReset
    levelInfoMsg = termBlue + "INFO" + termReset
    levelWarningMsg = termMagenta + "WARN" + termReset
    levelErrorMsg = termRed + "ERROR" + termReset
)

// STDERR FILES' INNER PATH
// TODO: CMS Config, check for custom log stuff
const (
    logPaths = "multicms/logs"
    logFiles = "stdout_%s.log" // %s date or log number
    stderrFiless = "stderr_%s.log" // %s date or log number
)

var logDirExists bool = false

func createLogDir () {
    if logDirExists {
        return
    }

    var logDirBuilder strings.Builder
    currentUser, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }

    logDirBuilder.WriteString(currentUser.HomeDir)
    logDirBuilder.WriteString(logPaths)

    err = os.MkdirAll(logDirBuilder.String(), 0640)
    if err != nil {
        log.Fatal(err)
    }

    logDirExists = true
    return
}

func infoLogerCustomFormat (groups []string, attribute slog.Attr) slog.Attr {
    fmt.Println(attribute)
    if attribute.Key == slog.LevelKey {
        // Rename the level key from "level" to "sev".
        attribute.Key = "sev"

        // Handle custom level values.
        level := attribute.Value.Any().(slog.Level)
        fmt.Println(level)

        switch level {
        case slog.LevelDebug:
            attribute.Value = slog.StringValue(levelLogMsg)
        case slog.LevelInfo:
            attribute.Value = slog.StringValue(levelInfoMsg)
        case slog.LevelWarn:
            attribute.Value = slog.StringValue(levelWarningMsg)
        case slog.LevelError:
            attribute.Value = slog.StringValue(levelErrorMsg)
        }
    }

    return attribute
}
