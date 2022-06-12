package loghelper

import "github.com/sirupsen/logrus"

var LoggingLevelConfigMapping = make(map[string]logrus.Level)

func Init() {
	LoggingLevelConfigMapping["info"] = logrus.InfoLevel
	LoggingLevelConfigMapping["warn"] = logrus.WarnLevel
	LoggingLevelConfigMapping["error"] = logrus.ErrorLevel
	LoggingLevelConfigMapping["fatal"] = logrus.FatalLevel
	LoggingLevelConfigMapping["debug"] = logrus.DebugLevel
	LoggingLevelConfigMapping["panic"] = logrus.PanicLevel

}
