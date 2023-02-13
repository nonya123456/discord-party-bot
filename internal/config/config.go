package config

import (
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configOnce sync.Once
var config *Config

type Config struct {
	Token string `mapstructure:"token"`
}

func NewConfig() *Config {
	configOnce.Do(func() {
		viper.SetConfigName("config")            // name of config file without extension
		viper.AddConfigPath("./internal/config") // path to look for config file, relative to working directory
		viper.AddConfigPath("/config")           // production config mount path

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic(errors.Wrap(err, "Config file not found"))
		}
		viper.AutomaticEnv()

		viper.WatchConfig() // Watch for changes to the configuration file and recompile
		if err := viper.Unmarshal(&config); err != nil {
			panic(errors.Wrap(err, "Cannot unmarshal config"))
		}
	})
	return config
}
