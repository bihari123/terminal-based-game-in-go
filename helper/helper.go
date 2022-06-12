package helper

import (
	"log"

	"github.com/bihari123/terminal-based-game-in-go/config"
	"github.com/spf13/viper"
)

func GetConfig() *config.Configuration {
	var myConfig config.Configuration = config.Configuration{}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found")
		} else {
			log.Fatal("Error in readiung the config file")
		}
	}

	return &myConfig
}
