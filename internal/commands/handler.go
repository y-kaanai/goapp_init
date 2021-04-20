package commands

import (
	"flag"
	"log"

	"github.com/Kanai-Yuki/goapp_init/internal/services/deljson"
	"github.com/Kanai-Yuki/goapp_init/internal/services/genjson"
	"github.com/Kanai-Yuki/goapp_init/internal/services/init"
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
	case string(AppInit):
		initCli := init.New()
		initCli.Exec()
	default:
		log.Fatalln("unexpected command: ", flag.Args())
	}
}
