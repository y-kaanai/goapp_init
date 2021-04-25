package deljson

import (
	"log"
	"os"

	"github.com/Kanai-Yuki/goapp_init/config"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

var jsonFileName string = "/goappinit.json"

func (_ Service) Exec() {
	err := os.Remove(config.Conf.App.CurrentDir + jsonFileName)
	if err != nil {
		log.Fatalln("Remove goappinit.json error: ", err)
	}
}
