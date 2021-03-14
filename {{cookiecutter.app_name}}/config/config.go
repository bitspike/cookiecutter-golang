package config

import (
	"bytes"
	"time"

	"github.com/spf13/viper"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper // nolint:gochecknoglobals

// Config returns a default config providers.
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance.
func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

// Initialize must be called at the very beginning of the application
// and before any other function of this package.
func Initialize() {
	defaultConfig = readViperConfig("{{cookiecutter.main_name|upper}}")
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	// Load default configuration
	err := v.ReadConfig(bytes.NewBuffer([]byte(getDefaultConfig())))
	if err != nil {
		panic(err)
	}

	return v
}
