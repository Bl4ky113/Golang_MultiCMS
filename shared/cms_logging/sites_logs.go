package cms_logging

import (
    "fmt"
    "os"
    "log/slog"
)

var logLoggerStarted bool = false

var infoLogger *slog.Logger
var errorLogger *slog.Logger

func initLogger () {
    if logLoggerStarted {
        return
    }

    // TODO: CMS Config, check for stdout or output logfile
    // TODO: CMS Config, check for custom logging types and options
    infoLogger = slog.New(slog.NewTextHandler(
        os.Stdout, 
        &slog.HandlerOptions{
            Level: slog.LevelInfo,
            ReplaceAttr: infoLogerCustomFormat,
        },
    ))

    slog.SetDefault(infoLogger)
    logLoggerStarted = true
    return
}

func InfoLog (sectionCode int, info string) {
    // TODO: slog into info log file

    initLogger()
        
    infoLogger.Info(info, "section_code", sectionCode)

    return
}

func WarnLog (sectionCode int, warn string) {

    return
}

func ErrorLog (sectionCode int, err error) string {
    outputError := fmt.Sprintf("Internal Error Code: %d", sectionCode)

    return outputError
}
