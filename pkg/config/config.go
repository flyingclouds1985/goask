package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

var (
	errKeyFormat    = errors.New("Key format must be like [key] or [key1.key2] without brackets")
)

//Config type.
type Config struct {
	config map[string]interface{}
	file   string
}

//Load constructs a config type based on config file.
//Basically it searches for the file in project root directory.
func Load(file string) (*Config, error) {
	c := new(Config)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, err
	}

	c.file = file
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(f, &c.config)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Setup inits configs.
// func Setup() {
// 	setSecretKey()
// }

//GetString returns value of a key in string format. 
//If it's not a valid string value it returns an empty string.
func (c *Config) GetString(key string) string {
	v := c.Get(key)
	if v, ok := v.(string); ok {
		return v
	}

	return ""
}

//GetInt returns value of a key in int format.
//If it's not a valid int value it returns zero.
func (c *Config) GetInt(key string) int {
	v := c.Get(key)
	if v, ok := v.(float64); ok {
		return int(v)
	}

	return 0
}

// Get returns value of a key.
func (c *Config) Get(key string) interface{} {
	keys, err := splitKeys(key)
	if err != nil {
		panic(err)
	}

	if len(keys) == 1 {
		if v, ok := c.config[key]; ok {
			return v
		}
	} else {
		if m, ok := c.config[keys[0]].(map[string]interface{}); ok {
			if v, ok := m[keys[1]]; ok {
				return v
			}
		}
	}

	return nil
}

// Set sets a [key/value] or [key1.key2/value] pair without brackets in configuration file.
func (c *Config) Set(key, value string) error {
	keys, err := splitKeys(key)
	if err != nil {
		return err
	}
	if len(keys) == 1 {
		c.config[key] = value
	} else {
		// if key exists add the value to nested map
		if nestedMap, ok := c.config[keys[0]].(map[string]interface{}); ok {
			nestedMap[keys[1]] = value
		} else {
			c.config[keys[0]] = map[string]interface{}{
				keys[1]: value,
			}
		}
	}

	b, err := json.MarshalIndent(c.config, "", "\t")

	err = ioutil.WriteFile(c.file, b, 0644)
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

// func setSecretKey() {
// 	secretKey, err := Get("router.secretKey")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	if secretKey == "" {
// 		err = Set("router.secretKey", routerSecretKey(20))
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// }
