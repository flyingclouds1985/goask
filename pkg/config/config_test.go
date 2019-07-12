package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Load config file
var c, _ = Load("conf_example.json")

func TestConfigFileNotExists(t *testing.T) {
	_, err := Load("not_exists.json")
	assert.Equal(t, "stat not_exists.json: no such file or directory", err.Error())
}

func TestGet(t *testing.T) {
	d := c.GetString("domain")
	u := c.GetString("database.username")
	p := c.GetInt("database.port")
	assert.Equal(t, d, "www.example.com")
	assert.Equal(t, u, "postgres")
	assert.Equal(t, p, 2345)
}

func TestNotKeyExists(t *testing.T) {
	empty := c.GetString("notExists")
	assert.Equal(t, empty, "")
}

func TestSet(t *testing.T) {
	c.Set("realm", "example.com")
	c.Set("auth.token", "jwt")
	c.Set("server.domain", "localhost")

	f := c.GetString("realm")
	s := c.GetString("auth.token")
	d := c.GetString("server.domain")

	assert.Equal(t, f, "example.com")
	assert.Equal(t, s, "jwt")
	assert.Equal(t, d, "localhost")

}
