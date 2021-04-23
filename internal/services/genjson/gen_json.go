package genjson

import (
	"io"
	"log"
	"os"

	"github.com/Kanai-Yuki/goapp_init/config"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

var (
	jsonFileName string = "/goappinit.json"
	jsonFilePath string
)

func (_ Service) Exec() {
	jsonFilePath = config.Conf.App.CurrentDir + "/internal/services/genjson" + jsonFileName

	src, err := os.Open(jsonFilePath)
	if err != nil {
		log.Fatalln("Opne goappinit.json error: ", err)
	}
	defer src.Close()

	dst, err := os.OpenFile(config.Conf.App.CurrentDir+jsonFileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalln("Create goappinit.json error: ", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatalln("Copy goappinit.json error: ", err)
	}
}
