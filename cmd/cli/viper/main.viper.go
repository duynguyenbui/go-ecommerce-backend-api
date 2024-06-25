package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	}
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // name of config file (without extension)
	viper.SetConfigType("yaml")      // type of the config file

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// configure structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Config Port:::", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Println("Database User:::", db.User)
		fmt.Println("Database Password:::", db.Password)
		fmt.Println("Database Host:::", db.Host)
	}
}
