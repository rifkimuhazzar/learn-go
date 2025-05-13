package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Trace("Hello Trace")

	logrus.SetLevel(logrus.TraceLevel)

	logrus.Debug("Hello Debug")
	logrus.Info("Hello Info")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Warn("Hello Warn")
	logrus.Error("Hello Error")
}
