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
	DEFAULT_DIR := "./config"
	DEFAULT_CONFIG := "default"

	v := viper.New()
	v.AddConfigPath(DEFAULT_DIR)
	v.SetConfigName(DEFAULT_CONFIG)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	configmap = v.AllSettings()

	if filename != DEFAULT_CONFIG {
		v2 := viper.New()
		v2.AddConfigPath(DEFAULT_DIR)
		v2.SetConfigName(filename)

		err2 := v2.ReadInConfig()
		if err2 != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err2))
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
	if len(keys) == 1 {
		return config.val[keys[0]].(string)
	}

	switch len(keys) {
	case 1:
		return config.val[keys[0]].(string)
	case 2:
		return config.val[keys[0]].(map[string]interface{})[keys[1]].(string)
	case 3:
		return config.val[keys[0]].(map[string]interface{})[keys[1]].(map[string]interface{})[keys[2]].(string)
	case 4:
		return config.val[keys[0]].(map[string]interface{})[keys[1]].(map[string]interface{})[keys[2]].(map[string]interface{})[keys[3]].(string)
	case 5:
		return config.val[keys[0]].(map[string]interface{})[keys[1]].(map[string]interface{})[keys[2]].(map[string]interface{})[keys[3]].(map[string]interface{})[keys[4]].(string)
	case 6:
		return config.val[keys[0]].(map[string]interface{})[keys[1]].(map[string]interface{})[keys[2]].(map[string]interface{})[keys[3]].(map[string]interface{})[keys[4]].(map[string]interface{})[keys[5]].(string)
	default:
		return "This depth is not supported"
	}
}

func merge(v1, v2 *viper.Viper) map[string]interface{} {
	v := v1.AllSettings()
	if err := mergo.Merge(&v, v2.AllSettings(), mergo.WithOverride); err != nil {
		panic(fmt.Errorf("Fatal error config file merge: %s \n", err))
	}
	return v
}
