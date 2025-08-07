package config

type Confuguration interface {
	GetString(string) (string, bool)
	GetInt(string) (int, bool)
	GetBool(string) (bool, bool)
	GetFloat(string) (float64, bool)

	GetStringDefault(string, string) string
	GetIntDefault(string, int) int
	GetBoolDefault(string, bool) bool
	GetFloatDefault(string, float64) float64
	GetSection(string) (Confuguration, bool)
}
