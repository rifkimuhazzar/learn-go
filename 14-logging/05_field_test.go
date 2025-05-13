package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.SetFormatter(&logrus.TextFormatter{})

	logger.Info("Hello Info")
	logger.WithField("First", 1).Info("Hello With Field")
	logger.WithField("First", 1).
		WithField("Second", 2).
		Info("Hello With Field")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"first_name": "Svelte",
		"last_name":  "Kit",
	}).Info("Hello Info")
}
