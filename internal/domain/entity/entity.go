package entity

import (
	"loader/internal/infrastructure/viper"
	"log"
)

var (
	CFG   map[string]any
	TOKEN string
)

func init() {
	cfg, err := viper.GetAll()
	if err != nil {
		log.Fatal("[DOMAIN][CONFIG] ERROR: ", err)
	}

	CFG = cfg

	TOKEN = CFG["token"].(string)
}

type Proj struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Links struct {
		Self     string `json:"self"`
		Branches string `json:"repo_branches"`
	} `json:"_links"`
	Branches []string
}
