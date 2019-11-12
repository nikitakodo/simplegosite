package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"simplesite/internal/app/appserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	log.Println("ky ky")

	config := appserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalf("error parce file: %s", err.Error())
	}

	if err := appserver.Start(config); err != nil {
		log.Fatalf("error start server: %s ", err.Error())
	}
}
