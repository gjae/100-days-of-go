// Set default values in viper
package main

import (
	"fmt"
	// https://github.com/spf13/viper
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("ContentDir", "path")
	viper.SetDefault("CookieTime", 60 * 60 * 24)	
	viper.SetDefault("Categories", []string{"Shopping", "Cart", "Movies"})

	fmt.Println(viper.AllKeys())
	fmt.Println(viper.GetString("ContentDir"))
	fmt.Println(viper.GetInt("CookieTime"))
	fmt.Println(viper.GetFloat64("CookieTime"))
}