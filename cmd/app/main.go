package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/yudgxe/sima-rest-api/internal/app"
	
	_ "github.com/lib/pq"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := app.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	app.Start(config);
}
