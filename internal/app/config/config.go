package config

type Config struct {
	Title          string   `toml:"title"`
	AllowedOrigins string   `toml:"allowed_origins"`
	BindAddr       string   `toml:"bind_addr"`
	LogLevel       string   `toml:"log_level"`
	TemplatesDir   string   `toml:"templates"`
	DB             database `toml:"database"`
	Session        Session  `toml:"Session"`
}
type database struct {
	Url           string
	DbName        string `toml:"db_name"`
	MigrationsDir string `toml:"migrations_dir"`
}

type Session struct {
	Key    string
	Name   string
	MaxAge int
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8001",
		LogLevel: "debug",
	}
}
