package viper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"log"
)

func GetAll() (map[string]any, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Println("[VIPER][USER_HOME] ERROR: ", err)
		return nil, err
	}

	viper.SetConfigName("settings")
	viper.SetConfigType("json")
	viper.AddConfigPath(fmt.Sprintf("%s/.config/gitlab-downloader", homedir))

	if err := viper.ReadInConfig(); err != nil {
		log.Println("[VIPER][READ] ERROR: ", err)
		return nil, err
	}

	allSettings := viper.AllSettings()
	return allSettings, nil
}

func SaveConfig() {
	if err := viper.WriteConfig(); err != nil {
		log.Println("[VIPER][WRITE] ERROR: ", err)
	}
}

func Get(key string) any {
	return viper.Get(key)
}

func Set(key string, value any) {
	viper.Set(key, value)
}
