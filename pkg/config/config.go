package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Addr        string `mapstructure:"ADDR"`
	UserSvcPort string `mapstructure:"USER_SVC_PORT"`
	UserSvcUrl  string `mapstructure:"USER_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	fmt.Println("[CONF]: config loaded")

	return
}
