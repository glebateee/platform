package config

func (c *DefaultConfig) GetStringDefault(path string, def string) (result string) {
	result, found := c.GetString(path)
	if !found {
		result = def
	}
	return
}
func (c *DefaultConfig) GetIntDefault(path string, def int) (result int) {
	result, found := c.GetInt(path)
	if !found {
		result = def
	}
	return
}
func (c *DefaultConfig) GetBoolDefault(path string, def bool) (result bool) {
	result, found := c.GetBool(path)
	if !found {
		result = def
	}
	return
}
func (c *DefaultConfig) GetFloatDefault(path string, def float64) (result float64) {
	result, found := c.GetFloat(path)
	if !found {
		result = def
	}
	return
}
