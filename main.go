package main

import (
	"log"
	"os"
)

var files []string = []string{
	"build",
	"cmd",
	"development",
	"internal",
	"pkg",
	"test",
}

const initText string = `package main

import "fmt"

func main() {
	fmt.Println("Hello Go!!")
}
`

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Get pwd error: ", err)
	}

	for _, file := range files {
		err := os.MkdirAll(pwd+"/"+file, 0777)
		if err != nil {
			log.Fatalln("Create directory error: ", err)
		}

		if file == "cmd" {
			_ = func() {
				file, err := os.OpenFile(pwd+"/"+"cmd"+"/main.go", os.O_RDWR|os.O_CREATE, 0777)
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
