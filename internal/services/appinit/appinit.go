package appinit

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/Kanai-Yuki/goapp_init/config"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

type initJson struct {
	Folders []string `json:"folders"`
	Files   []string `json:"files"`
}

var (
	initTextFile     string = "/init.text"
	initTextFilePath string
)

func (_ Service) Exec() {
	initTextFilePath = config.Conf.App.CurrentDir + "/internal/services/appinit" + initTextFile

	data, err := ioutil.ReadFile(config.Conf.App.CurrentDir + "goappinit.json")
	if err != nil {
		log.Fatalln("Opne goappinit.json error: ", err)
	}

	var ij initJson
	err = json.Unmarshal(data, &ij)
	if err != nil {
		log.Fatalln("Unmarshal goappinit.json error: ", err)
	}

	folders := ij.Folders

	for _, folder := range folders {
		err := os.MkdirAll(config.Conf.App.CurrentDir+"/"+folder, 0777)
		if err != nil {
			log.Fatalln("Create directory error: ", err)
		}

		if folder == "cmd" {
			makeMainGo()
		}
	}
}

func makeMainGo() {
	src, err := os.Open(initTextFilePath)
	if err != nil {
		log.Fatalln("Opne init.text error: ", err)
	}
	defer src.Close()

	dst, err := os.OpenFile(config.Conf.App.CurrentDir+"/"+"cmd"+"/main.go", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalln("Create main.go error: ", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatalln("Write main.go error: ", err)
	}
}
