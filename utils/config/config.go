package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var GC Config

type Redis struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Config struct {
	Redis *Redis
}

func LoadConfig() {
	viper.AddConfigPath("/Users/zhaomeng/GolangProject/markdown/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err := viper.Unmarshal(&GC)
	if err != nil {
		panic(err)
	}
}
