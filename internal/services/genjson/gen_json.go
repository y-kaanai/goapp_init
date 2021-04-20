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

func (_ Service) Exec() {
	src, err := os.Open("goappinit.json")
	if err != nil {
		log.Fatalln("Opne goappinit.json error: ", err)
	}
	defer src.Close()

	dst, err := os.OpenFile(config.Conf.App.CurrentDir+"/"+"goappinit2.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalln("Create goappinit.json error: ", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatalln("Copy goappinit.json error: ", err)
	}
}
