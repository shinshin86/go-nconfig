package nconfig

import (
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/spf13/viper"
)

// Config store on config name and config values
type Config struct {
	val  map[string]interface{}
	Name string
}

// New is returns the config loaded with the specified configuration name
func New(filename string) *Config {
	var configmap map[string]interface{}
	DefaultDir := "./config"
	DefaultConfig := "default"

	if filename == "" {
		filename = DefaultConfig
	}

	v := viper.New()
	v.AddConfigPath(DefaultDir)
	v.SetConfigName(DefaultConfig)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	configmap = v.AllSettings()

	if filename != DefaultConfig {
		v2 := viper.New()
		v2.AddConfigPath(DefaultDir)
		v2.SetConfigName(filename)

		err2 := v2.ReadInConfig()
		if err2 != nil {
			panic(fmt.Errorf("Fatal error config file: %s", err2))
		}
		configmap = merge(v, v2)
	}

	return &Config{
		val:  configmap,
		Name: filename,
	}
}

// Get the value of the configuration file with specified key
func (config *Config) Get(key string) string {
	keys := strings.Split(key, ".")

	v := config.val
	var val string
	for i, key := range keys {
		if i == len(keys)-1 {
			val = v[key].(string)
		} else {
			v = v[key].(map[string]interface{})
		}
	}

	return val
}

func merge(v1, v2 *viper.Viper) map[string]interface{} {
	v := v1.AllSettings()
	if err := mergo.Merge(&v, v2.AllSettings(), mergo.WithOverride); err != nil {
		panic(fmt.Errorf("Fatal error config file merge: %s", err))
	}
	return v
}
