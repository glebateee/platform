package config

import "strings"

type DefaultConfig struct {
	configData map[string]interface{}
}

func (c *DefaultConfig) get(path string) (result interface{}, found bool) {
	data := c.configData
	for _, key := range strings.Split(path, ":") {
		result, found = data[key]
		if newSection, ok := result.(map[string]interface{}); ok && found {
			data = newSection
		} else {
			return
		}
	}
	return
}

func (c *DefaultConfig) GetSection(path string) (section Configuration, found bool) {
	value, found := c.get(path)
	if found {
		if sectionData, ok := value.(map[string]interface{}); ok {
			section = &DefaultConfig{configData: sectionData}
		}
	}
	return
}

func (c *DefaultConfig) GetString(path string) (result string, found bool) {
	value, found := c.get(path)
	if found {
		result = value.(string)
	}
	return

}
func (c *DefaultConfig) GetInt(path string) (result int, found bool) {
	value, found := c.get(path)
	if found {
		result = int(value.(float64))
	}
	return
}
func (c *DefaultConfig) GetBool(path string) (result bool, found bool) {
	value, found := c.get(path)
	if found {
		result = value.(bool)
	}
	return
}
func (c *DefaultConfig) GetFloat(path string) (result float64, found bool) {
	value, found := c.get(path)
	if found {
		result = value.(float64)
	}
	return
}
