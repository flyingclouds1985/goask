package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
)

var (
	configDir       string
	errKeyNotExists = errors.New("Key does not exist")
)

// Setup inits configs.
func Setup() {
	_, f, _, _ := runtime.Caller(0)
	configDir = path.Dir(f)

	initRouterConfig()
}

// Get returns value of key.
func Get(fileName, key string) (string, error) {
	fn := configDir + "/" + fileName + ".json"

	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return "", err
	}

	f, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", err
	}

	var m map[string]string
	err = json.Unmarshal(f, &m)
	if err != nil {
		return "", err
	}

	if v, ok := m[key]; ok {
		return v, nil
	}
	return "", errKeyNotExists
}

// Set sets a key value in specified config file.
func Set(fileName, key, value string) error {
	fn := configDir + "/" + fileName + ".json"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return err
	}

	f, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	var m map[string]string
	err = json.Unmarshal(f, &m)
	if err != nil {
		return err
	}

	m[key] = value

	b, err := json.MarshalIndent(m, "", "\t")

	err = ioutil.WriteFile(fn, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func initRouterConfig() {
	secretKey, err := Get("router", "secretKey")
	if err != nil {
		log.Println(err)
	}
	if secretKey == "" {
		err = Set("router", "secretKey", routerSecretKey(20))
		if err != nil {
			log.Println(err)
		}
	}
}
