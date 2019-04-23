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
	"strings"
)

var (
	_, f, _, _      = runtime.Caller(0)
	configFile      = path.Dir(path.Dir(f)) + "/configuration.json"
	errKeyNotExists = errors.New("Key does not exist")
	errKeyFormat    = errors.New("Key format must be like [key] or [key1.key2] without brackets")
)

// Setup inits configs.
func Setup() {
	setSecretKey()
}

//GetString returns value of a key in string format.
func GetString(key string) (string, error) {
	v, err := Get(key)
	if err != nil {
		return "", err
	}
	if v, ok := v.(string); ok {
		return v, nil
	}

	return "", fmt.Errorf("Unable to cast %#v of type %T to string", v, v)
}

// Get returns value of a key.
func Get(key string) (interface{}, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return "", err
	}

	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		return "", err
	}

	var m map[string]interface{}
	err = json.Unmarshal(f, &m)
	if err != nil {
		return "", err
	}

	keys, err := splitKeys(key)
	if err != nil {
		return "", err
	}

	if len(keys) == 1 {
		if v, ok := m[key]; ok {
			return v, nil
		}
	} else {
		if m, ok := m[keys[0]].(map[string]interface{}); ok {
			if v, ok := m[keys[1]]; ok {
				return v, nil
			}
		}
	}

	return "", errKeyNotExists
}

// Set sets a [key/value] or [key1.key2/value] pair without brackets in configuration file.
func Set(key, value string) error {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return err
	}

	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	var m map[string]interface{}
	err = json.Unmarshal(f, &m)
	if err != nil {
		return err
	}

	keys, err := splitKeys(key)
	if err != nil {
		return err
	}
	if len(keys) == 1 {
		m[key] = value
	} else {
		// if the key exists add the value to inner map
		if innerMap, ok := m[keys[0]].(map[string]interface{}); ok {
			innerMap[keys[1]] = value
		} else {
			m[keys[0]] = map[string]interface{}{
				keys[1]: value,
			}
		}
	}

	b, err := json.MarshalIndent(m, "", "\t")

	err = ioutil.WriteFile(configFile, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func splitKeys(str string) ([]string, error) {
	k := strings.Split(str, ".")

	if len(k) == 1 {
		return k, nil
	} else if len(k) > 2 {
		return nil, errKeyFormat
	}

	return k, nil
}

func setSecretKey() {
	secretKey, err := Get("router.secretKey")
	if err != nil {
		log.Println(err)
	}
	if secretKey == "" {
		err = Set("router.secretKey", routerSecretKey(20))
		if err != nil {
			log.Println(err)
		}
	}
}
