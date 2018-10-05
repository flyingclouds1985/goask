package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	Setup()
}

func TestGetConfig(t *testing.T) {
	realm, _ := Get("router", "realm")
	assert.Equal(t, realm, "goask.com")
}

func TestNotExistsConfigFile(t *testing.T) {
	realm, err := Get("notexists", "realm")
	_, ok := err.(*os.PathError)
	assert.Equal(t, realm, "")
	assert.Equal(t, ok, true)
}

func TestNotKeyExists(t *testing.T) {
	_, err := Get("router", "notExists")
	assert.EqualError(t, errKeyNotExists, err.Error())
}

func TestSetConfig(t *testing.T) {
	Set("router", "3", "33")
	Set("router", "44", "444")

}
func TestSetToNotExistentFile(t *testing.T) {
	err := Set("notexists", "k", "v")
	_, ok := err.(*os.PathError)
	assert.Equal(t, ok, true)
}
