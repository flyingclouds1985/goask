package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	Setup()
}

func TestGet(t *testing.T) {
	realm, _ := Get("router.realm")
	assert.Equal(t, realm, "goask.com")
}

func TestNotKeyExists(t *testing.T) {
	_, err := Get("notExists")
	assert.EqualError(t, errKeyNotExists, err.Error())
}

func TestSet(t *testing.T) {
	Set("first", "1st")
	Set("second.third", "3rd")

	f, _ := Get("first")
	s, _ := Get("second.third")

	assert.Equal(t, f, "1st")
	assert.Equal(t, s, "3rd")
}
