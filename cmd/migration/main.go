package main

import (
	"database/sql"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
	"os"
	"simplesite/internal/app/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	configPath string
	mode       string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
	flag.StringVar(&mode, "mode", "", "migration mode (up/down)")
}

func main() {
	flag.Parse()
	if mode != "up" && mode != "down" {
		log.Fatalf("Migration error: invalid mode - %s", mode)
	}

	conf := config.NewConfig()
	_, err := toml.DecodeFile(configPath, conf)
	if err != nil {
		log.Fatalf("error parce file: %s", err.Error())
	}

	db, err := sql.Open("postgres", conf.DB.Url)
	if err != nil {
		log.Fatalf("error open connection to db: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error ping db server: %s", err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+conf.DB.MigrationsDir,
		conf.DB.DbName,
		driver,
	)
	if err != nil {
		log.Fatalf("Migrations error %s", err.Error())
	}

	if mode == "up" {
		err = m.Up()
	} else {
		err = m.Down()
	}
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Printf("Migration done : %s", err.Error())
			os.Exit(0)
		} else {
			log.Fatalf("Migrations error %s", err.Error())
		}
	}
}
