package jsonconfig

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Config struct {
	data  map[string]interface{}
}

func LoadConfig(uri string) (*Config, error) {
	if !string.Contains("://") || strings.HasPrefix(strings.ToLower(uri), "file://") {
		return LoadConfigFromFile(uri)
	}
	return LoadConfigFromUrl(uri)
}

func LoadConfigFromFile(fileName string) (*Config, error) {

	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return loadConfig(buf)
}

func LoadConfigFromString(s string) (*Config, error) {
	return loadConfig([]byte(s))
}

func loadConfig(buf []byte) (*Config, error) {

	var f interface{}
	err := json.Unmarshal(buf, &f)
	if err != nil {
		return nil, err
	}
	c := Config{
		data: f.(map[string]interface{}),
	}
	return &c, nil
}

func LoadConfigFromUrl(url string) (*Config, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return loadConfig(buf)
}

func (this Config) GetString(key string) (string, bool) {
	val, ok := this.data[key]
	if !ok {
		return "", false
	}
	return val.(string), true
}

func (this Config) GetInt(key string) (int, bool) {
	val, ok := this.data[key]
	if !ok {
		return 0, false
	}
	return int(val.(float64)), true
}

func (this Config) GetFloat(key string) (float64, bool) {
	val, ok := this.data[key]
	if !ok {
		return 0, false
	}
	return val.(float64), true
}

func (this Config) GetBool(key string) (bool, bool) {
	val, ok := this.data[key]
	if !ok {
		return false, false
	}
	return val.(bool), true
}

func (this Config) GetArray(key string) ([]interface{}, bool) {
	val, ok := this.data[key]
	if !ok {
		return []interface{}(nil), false
	}
	return val.([]interface{}), true
}

func (this Config) GetStringArray(key string) ([]string, bool) {
	val, ok := this.GetArray(key)
	if !ok {
		return []string(nil), false
	}
	ret := make([]string, len(val))
	for i, v := range val {
		ret[i] = v.(string)
	}
	return ret, true
}

func (this Config) GetIntArray(key string) ([]int, bool) {
	val, ok := this.GetArray(key)
	if !ok {
		return []int(nil), false
	}
	ret := make([]int, len(val))
	for i, v := range val {
		ret[i] = int(v.(float64))
	}
	return ret, true
}

func (this Config) GetFloat64Array(key string) ([]float64, bool) {
	val, ok := this.GetArray(key)
	if !ok {
		return []float64(nil), false
	}
	ret := make([]float64, len(val))
	for i, v := range val {
		ret[i] = v.(float64)
	}
	return ret, true
}

func (this Config) GetBoolArray(key string) ([]bool, bool) {
	val, ok := this.GetArray(key)
	if !ok {
		return []bool(nil), false
	}
	ret := make([]bool, len(val))
	for i, v := range val {
		ret[i] = v.(bool)
	}
	return ret, true
}

func (this Config) GetObject(key string) (*Config, bool) {
	val, ok := this.data[key]
	if !ok {
		return nil, false
	}
	c := Config{
		data:  val.(map[string]interface{}),
	}
	return &c, true
}
