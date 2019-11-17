package config

type Config struct {
	AllowedOrigins string   `toml:"allowed_origins"`
	BindAddr       string   `toml:"bind_addr"`
	LogLevel       string   `toml:"log_level"`
	TemplatesDir   string   `toml:"templates"`
	DB             database `toml:"database"`
	Session        session  `toml:"session"`
	Cache          cache    `toml:"cache"`
}
type database struct {
	Url           string
	DbName        string `toml:"db_name"`
	MigrationsDir string `toml:"migrations_dir"`
}

type cache struct {
	Addr     string `toml:"url"`
	Password string `toml:"password"`
	Prefix   string `toml:"prefix"`
}

type session struct {
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
