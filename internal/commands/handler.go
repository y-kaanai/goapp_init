package commands

import (
	"flag"
	"log"

	"github.com/Kanai-Yuki/goapp_init/internal/services/appinit"
	"github.com/Kanai-Yuki/goapp_init/internal/services/deljson"
	"github.com/Kanai-Yuki/goapp_init/internal/services/genjson"
)

func Handler() {
	flag.Parse()

	switch flag.Arg(0) {
	// goapp gen-json
	case string(GenJson):
		genjsonCli := genjson.New()
		genjsonCli.Exec()
	// goapp del-json
	case string(DelJson):
		deljsonCli := deljson.New()
		deljsonCli.Exec()
	// goapp init
	case string(AppInit):
		initCli := appinit.New()
		initCli.Exec()
	default:
		log.Fatalln("unexpected command: ", flag.Args())
	}
}
