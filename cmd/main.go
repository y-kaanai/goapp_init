package main

import (
	"github.com/Kanai-Yuki/goapp_init/config"
	"github.com/Kanai-Yuki/goapp_init/pkg/commands"
)

func main() {
	// 環境変数の設定
	config.New()

	commands.Handler()
}
