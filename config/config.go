package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type BotConfig struct {
	BotToken            string
	BotPrefix           string
	BotCardsListEnabled bool
	CardsPath           string
	GameName            string
}

var Config BotConfig

func ReadConfig() {
	//set this up for other environments ideally based on environment vars. Super not needed right now but it helps me sleep at night :)
	if _, err := toml.DecodeFile("config/local.cfg.toml", &Config); err != nil {
		log.Fatal("Could not open config file", err)
	}
	log.Println("Successfully read config")
}
