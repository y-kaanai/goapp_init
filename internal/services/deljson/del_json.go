package deljson

import (
	"log"
	"os"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (_ Service) Exec() {
	err := os.Remove("goappinit2.json")
	if err != nil {
		log.Fatalln("Remove goappinit.json error: ", err)
	}
}
