package loghelper

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bihari123/terminal-based-game-in-go/config"
	"github.com/icza/gog"
	"github.com/sirupsen/logrus"
)

var shouldPrint bool

func Log(logVal ...string) {
	joinedInfo := strings.Join(logVal, "--")
	if shouldPrint {
		fmt.Println(joinedInfo)
	}
	logrus.Info(joinedInfo)
}

func LogError(logVal ...string) {
	joinedInfo := strings.Join(logVal, "--")
	if shouldPrint {
		fmt.Println(joinedInfo)
	}

	logrus.Error(joinedInfo)
}

func LogWarning(logVal ...string) {
	joinedInfo := strings.Join(logVal, "--")
	if shouldPrint {
		fmt.Println(joinedInfo)
	}
	logrus.Warn(joinedInfo)
}

func ConfigLogurus(myConfig *config.Configuration) {
	logFile := myConfig.AppConfig.LogFile
	logFolder := myConfig.AppConfig.LogFolder
	loglevel := myConfig.AppConfig.LogLevel

	shouldPrint = gog.If(myConfig.AppConfig.Mode != "prod", true, false)

	// SetFormatter sets the standard logger formatter.

	logrus.SetFormatter(&logrus.JSONFormatter{})

	// creating the log folder anf file if not exists
	if logFile == "stdout" {
		logrus.SetOutput(os.Stdout)
	} else {
		// create a new folder and set the output file

		logPath := filepath.Join(logFolder, fmt.Sprintf("%s.log", logFile))

		if err := os.MkdirAll(logFolder, 0700); err != nil {
			errMsg := fmt.Sprintf("error in making the log folder: %s", err)
			LogError(errMsg)
			panic(errMsg)
		}

		if file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
			LogError(fmt.Sprintf("failed to create a log file: %s", err))
			panic("failed to create a log file")
		} else {
			logrus.SetOutput(file)
			LogWarning("Output set to the logFile")
		}
		cmdLoggingLevel := flag.String("loglevel", "", "command line option to control logging level of the service")
		flag.Parse()

		if len(*cmdLoggingLevel) > 0 {
			logrus.SetLevel(LoggingLevelConfigMapping[*cmdLoggingLevel])
			LogWarning("Logging level from comand line")
		} else {
			logrus.SetLevel(LoggingLevelConfigMapping[loglevel])
			LogWarning("Logging leve from property")
		}

	}
}
