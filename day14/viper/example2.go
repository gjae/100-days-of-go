// Reading and parsing .env file
package main

import (
	"fmt"
	// https://github.com/spf13/viper
	"github.com/spf13/viper"
)

const (
	ENVFILE string = "envfile"
)

type Config struct {
	SecretKey string `mapstructure:"SECRET_KEY"`
	IsDebug bool `mapstructure:"DEBUG"`
}

func NewConfig(configName string, fileformat string) (config *Config ){
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.Unmarshal(&config)

	return 
}


func main() {
	config := NewConfig("", ENVFILE)
	fmt.Println(viper.AllKeys())
	fmt.Println(viper.GetBool("debug"))

	fmt.Println(config)
	fmt.Println(config.SecretKey)
	fmt.Println(config.IsDebug)
}