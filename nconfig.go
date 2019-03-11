package nconfig

import (
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/spf13/viper"
)

type Config struct {
	val      map[string]interface{}
	filename string
}

func NewConfig(filename string) *Config {
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
		val:      configmap,
		filename: filename,
	}
}

func (config *Config) Get(key string) string {
	keys := strings.Split(key, ".")
	if len(keys) == 1 {
		return config.val[keys[0]].(string)
	}

	var val string
	var configmap map[string]interface{}
	for i, s := range keys {
		if i != (len(keys) - 1) {
			configmap = config.val[keys[i]].(map[string]interface{})
		}
		if i == (len(keys) - 1) {
			val = configmap[s].(string)
		}

	}
	return val
}

func merge(v1, v2 *viper.Viper) map[string]interface{} {
	v := v1.AllSettings()
	if err := mergo.Merge(&v, v2.AllSettings(), mergo.WithOverride); err != nil {
		panic(fmt.Errorf("Fatal error config file merge: %s \n", err))
	}
	return v
}
