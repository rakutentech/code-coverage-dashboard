package services

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
    Setup()
}
func Setup() {
    err := godotenv.Load("../.env.testing")
    if err != nil {
        panic(err)
    }
}


func TestMe(t *testing.T) {
    assert.True(t, true)
}