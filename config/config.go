package config

type Configuration interface {
	GetString(name string) (configValue string, found bool)
	GetInt(name string) (configValue int, found bool)
	GetBool(name string) (configValue bool, found bool)
	GetFloat(name string) (configValue float64, found bool)

	GetStringDefault(name string, defaultValue string) (configValue string)
	GetIntDefault(name string, defaultValue int) (configValue int)
	GetBoolDefault(name string, defaultValue bool) (configValue bool)
	GetFloatDefault(name string, defaultValue float64) (configValue float64)
	GetSection(SectionName string) (section Configuration, found bool)
}
