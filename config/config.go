package config

import (
	"log"
	"os"
)

var Conf *Config

func New() {
	Conf = &Config{
		App: newAppConfig(),
	}
}

func newAppConfig() *AppConfig {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Get current directory error: ", err)
	}

	return &AppConfig{
		CurrentDir: currentDir,
	}
}
