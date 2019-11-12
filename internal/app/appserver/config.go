package appserver

type Config struct {
	Title        string   `toml:"title"`
	BindAddr     string   `toml:"bind_addr"`
	LogLevel     string   `toml:"log_level"`
	TemplatesDir string   `toml:"templates"`
	DB           database `toml:"database"`
	Session      session  `toml:"Session"`
}
type database struct {
	Url string
}

type session struct {
	Key    string
	Name   string
	MaxAge int
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
