package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("we cant find your config file")
		} else {
			fmt.Printf("samething went wrong in your config file")
		}
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}
}
