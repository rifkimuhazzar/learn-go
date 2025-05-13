package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Info")
	logger.Warn("Hello Warn")
	logger.Error("Hello Error")
}
