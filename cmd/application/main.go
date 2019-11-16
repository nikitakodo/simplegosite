package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"simplesite/internal/app/appserver"
	"simplesite/internal/app/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()
	conf := config.NewConfig()
	_, err := toml.DecodeFile(configPath, conf)
	if err != nil {
		log.Fatalf("error parce file: %s", err.Error())
	}

	if err := appserver.Start(conf); err != nil {
		log.Fatalf("error start server: %s ", err.Error())
	}
}
