package logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogrusOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("aplication.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Hello Info")
	logger.Warn("Hello Warn")
	logger.Error("Hello Error")
}
