package test

import (
	"testing"

	"go_restful_api/simple"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializeConnection("Database")
	assert.NotNil(t, connection)
	cleanup()
}