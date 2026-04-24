package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func config() (api string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	viper.SetConfigName("specdriver")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)
	viper.AddConfigPath(home + "/.config")
	viper.AddConfigPath(home + "/AppData/Local") // Windows specific
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	api_viper := viper.GetString("API")
	return api_viper
}
