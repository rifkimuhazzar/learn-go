package logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello Logrus")
	fmt.Println("Hello Logrus")
}