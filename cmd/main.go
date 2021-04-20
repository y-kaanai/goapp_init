package main

import (
	"github.com/Kanai-Yuki/goapp_init/config"
	"github.com/Kanai-Yuki/goapp_init/internal/commands"
)

func main() {
	config.New()

	commands.Handler()
}
