package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
)

type ConfigStructure struct {
	MacPass		string	`mapstructure:"macos"`
	LinuxPass	string	`mapstructure:"linux"`
	WindowPass	string	`mapstructure:"windows"`
	PostHost	string	`mapstructure:"postgres"`
	MySQLHost	string	`mapstructure:"mysql"`
	MongoHost	string	`mapstructure:"mongodb"`
}

var CONFIG = ".config.json"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default file", CONFIG)
	} else {
		CONFIG = os.Args[1]
	}

	viper.SetConfigType("json")
	viper.SetConfigFile(CONFIG)
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	viper.ReadInConfig()

	if viper.IsSet("macos") {
		fmt.Println("MACOS: ", viper.Get("macos"))
	} else {
		fmt.Println("macos not set!")
	}

	if viper.IsSet("active") {
		value := viper.GetBool("active")
		if value {
			postgres := viper.Get("postgres")
			mysql := viper.Get("mysql")
			mongo := viper.Get("mongodb")
			fmt.Println("P: ", postgres, "My: ", mysql, "Mo: ", mongo)
		}
	} else {
		fmt.Println("active is not set!")
	}

	var t ConfigStructure
	err := viper.Unmarshal(&t)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t)
}