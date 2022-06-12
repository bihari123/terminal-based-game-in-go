package helper

import (
	"fmt"
	"log"

	"github.com/bihari123/terminal-based-game-in-go/config"
	"github.com/fsnotify/fsnotify"
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
			log.Fatal("Error in reading the config file")
		}
	}

	if err:=viper.Unmarshal(&myConfig);err!=nil{
		panic(fmt.Errorf("Error while unmarshalling the config : %w",err))

	}

	/*
	Viper supports the ability to have your application live read a config file while running. Gone are the days of needing to restart a server to have a
	*/ 
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
	fmt.Println("Config file changed:", e.Name)
})
	return &myConfig
}
