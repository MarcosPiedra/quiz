package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	WebConfig struct {
		Host string `mapstructure:"HOST"`
		Port string `mapstructure:"PORT"`
	}

	AppConfig struct {
		Environment string    `mapstructure:"ENVIRONMENT"`
		LogLevel    string    `mapstructure:"LOGLEVEL"`
		Web         WebConfig `mapstructure:"WEB"`
	}
)

func Setup() (cfg AppConfig, err error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/quiz")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.SetDefault("LogLevel", "Debug")

	err = viper.Unmarshal(&cfg)

	return
}

func (c WebConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
