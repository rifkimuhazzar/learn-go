package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "helloworld")
	entry.Info("Hello World")

	logger.Info("Hello Info")
	logger.WithField("username", "helloworld").Info("Hello World")
}
