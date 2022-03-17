package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"

	logging "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// LoadYaml load properties from yaml and convert to dot properties, then set into map
func LoadYaml(content []byte) (map[string]string, error) {
	m := make(map[string]interface{})
	err := yaml.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}
	converted := make(map[string]string)
	for k, v := range m {
		flatValue(v, k, converted)
	}
	return converted, nil
}

func flatValue(v interface{}, parent string, m map[string]string) {
	typ := reflect.TypeOf(v).Kind()
	if typ == reflect.Map {
		for k, sv := range v.(map[interface{}]interface{}) {
			key := fmt.Sprintf("%s", k)
			if parent != "" {
				key = fmt.Sprintf("%s.%s", parent, k)
			}
			flatValue(sv, key, m)
		}
	} else if typ == reflect.Int {
		value := fmt.Sprintf("%v", v)
		m[parent] = value
	} else if typ == reflect.String {
		value := fmt.Sprintf("%v", v)
		m[parent] = value
	} else if typ == reflect.Slice {
		for i, sv := range v.([]interface{}) {
			key := fmt.Sprintf("[%d]", i)
			if parent != "" {
				key = fmt.Sprintf("%s[%d]", parent, i)
			}
			flatValue(sv, key, m)
		}
	}
}

// Load load config file and unmarshall
func Load(filename string, v interface{}) error {
	content, err := ioutil.ReadFile(filepath.Clean(filename)) // fix gosec G304
	if err != nil {
		logging.WithError(err).Error("Cannot load file")
		return err
	}
	if strings.HasSuffix(filename, ".json") {
		err = json.Unmarshal(content, v)
		if err != nil {
			logging.WithError(err).Error("Cannot unmarshall json file")
			return err
		}
	} else if strings.HasSuffix(filename, ".yaml") || strings.HasSuffix(filename, ".yml") {
		err = yaml.Unmarshal(content, v)
		if err != nil {
			logging.WithError(err).Error("failed to unmarshall yaml file")
			return err
		}
	} else {
		logging.Warnf("Unsupported file type %v, neither json, nor yaml(yml), ignore it", filename)
	}
	return err
}
