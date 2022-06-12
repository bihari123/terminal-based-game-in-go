package loghelper

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func Log(logVal ...string) {
	joinedInfo := strings.Join(logVal, "--")
	fmt.Println(joinedInfo)
	logrus.Info(joinedInfo)
}

func LogError(logVal ...string) {
	joinedInfo := strings.Join(logVal, "--")
	fmt.Println(joinedInfo)
	logrus.Error(joinedInfo)
}

func LogWarning(logVal ...string) {
	joinedInfo := strings.Join(logVal, "--")
	fmt.Println(joinedInfo)
	logrus.Warn(joinedInfo)
}
