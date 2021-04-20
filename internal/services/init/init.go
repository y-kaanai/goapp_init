package init

import (
	"log"
	"os"

	"github.com/Kanai-Yuki/goapp_init/config"
)

const (
	Slash    string = "/"
	initText string = `package main

import "fmt"

func main() {
	fmt.Println("Hello Go!!")
}
`
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (_ Service) Exec() {
	folders := []string{}

	for _, folder := range folders {
		err := os.MkdirAll(config.Conf.App.CurrentDir+"/"+folder, 0777)
		if err != nil {
			log.Fatalln("Create directory error: ", err)
		}

		if folder == "cmd" {
			_ = func() {
				file, err := os.OpenFile(config.Conf.App.CurrentDir+"/"+"cmd"+"/main.go", os.O_RDWR|os.O_CREATE, 0777)
				if err != nil {
					log.Fatalln("Create main.go error: ", err)
				}
				defer file.Close()

				_, err = file.WriteString(initText)
				if err != nil {
					log.Fatalln("Write main.go error: ", err)
				}
			}
		}
	}
}
