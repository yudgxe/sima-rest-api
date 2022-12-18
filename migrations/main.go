package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/yudgxe/sima-rest-api/internal/app"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
	flag.Parse()

	config := app.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	addr :=  fmt.Sprintf("%s:%s", config.DB.Host, config.DB.Port)
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     config.DB.User,
		Password: config.DB.Password,
		Database: config.DB.Name,
	})

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}

	if err != nil {
		log.Fatal(err)
	}

}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
