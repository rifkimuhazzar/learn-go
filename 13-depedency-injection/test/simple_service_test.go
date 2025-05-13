package test

import (
	"fmt"
	"testing"

	"go_restful_api/simple"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
    if err != nil {
        fmt.Println(err)
        fmt.Println(simpleService)
    } else {
        fmt.Println(err)
        fmt.Println(simpleService.SimpleRepository)
    }
    assert.NotNil(t, err)
    assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
    if err != nil {
        fmt.Println(err)
        fmt.Println(simpleService)
    } else {
        fmt.Println(err)
        fmt.Println(simpleService)
        fmt.Println(simpleService.SimpleRepository)
    }
    assert.NotNil(t, simpleService)
    assert.NotNil(t, simpleService.SimpleRepository)
    assert.Nil(t, err)
}
