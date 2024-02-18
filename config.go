package main

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig(fileName string) *viper.Viper {
	config := viper.New()

	config.SetConfigName(fileName)
	config.SetConfigType("toml")

	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")

	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Error while parsing configuration file", err)
	}

	return config
}
