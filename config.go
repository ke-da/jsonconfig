// Package read map config out of json file
// It'll iterate through different possible env and replace config from env with lower priority
// with the one exist in higher priority

package jsonconfig

import (
	"encoding/json"
	"log"
	"os"
)

// Type for config value
type Value map[string]interface{}

// Get string value
func (c Value) GetStr(key string) (string, bool) {
	str, ok := c[key]
	if ok {
		return str.(string), ok
	}
	return "", ok
}

// Get string from map with key of type string
func (c Value) GetMapStr(key string) (map[string]string, bool) {
	mapStrIn, ok := c[key].(map[string]interface{})
	mapStrStr := make(map[string]string)
	for k, v := range mapStrIn {
		switch vv := v.(type) {
		case string:
			mapStrStr[k] = vv
		}
	}
	return mapStrStr, ok
}

// Get array of string from map with key of type string
func (c Value) GetStrSlice(key string) ([]string, bool) {
	sliceIntf, ok := c[key].([]interface{})
	if !ok {
		return nil, false
	}
	strSlice := make([]string, len(sliceIntf))
	for i, s := range sliceIntf {
		strSlice[i] = s.(string)
	}
	return strSlice, true
}

//GO_ENV is current application environment
var GO_ENV = os.Getenv("GO_ENV")

var envs []string

//{"dev", "staging", "test", "prod"}

func SetEnvs(customEnvs []string) {
	envs = customEnvs
}

//Load configuration file
func LoadFile(filename string) Value {
	Config := make(Value)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		panic("invalid config file provide")
	}
	//Container for config map[string]interface
	var cMap interface{}
	decoder := json.NewDecoder(file)
	decoder.Decode(&cMap)
	if cMap == nil {
		panic("Unable to load config")
	}

	//Iterate through envs and replace config according to order of the array
	hasValue := false
	for _, env := range envs {
		envMap := cMap.(map[string]interface{})
		if envConfig, ok := envMap[env]; ok {
			for k, config := range envConfig.(map[string]interface{}) {
				Config[k] = config
				hasValue = true
			}
		}
		if GO_ENV == env || GO_ENV == "" {
			break
		}
	}
	if hasValue {
		return Config
	}
	return nil
}

func init() {
	SetEnvs([]string{"dev", "staging", "test", "prod"})
}
