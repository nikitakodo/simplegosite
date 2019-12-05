package config

type Config struct {
	AllowedOrigins   string   `toml:"allowed_origins"`
	BindAddr         string   `toml:"bind_addr"`
	LogLevel         string   `toml:"log_level"`
	TemplatesDir     string   `toml:"templates"`
	AssetsDir        string   `toml:"assets"`
	AssetsUrl        string   `toml:"assets_url"`
	AdminAssetsUrl   string   `toml:"admin_assets_url"`
	CabinetAssetsUrl string   `toml:"cabinet_assets_url"`
	UploadsDir       string   `toml:"uploads"`
	UploadsUrl       string   `toml:"uploads_url"`
	DB               database `toml:"database"`
	Session          session  `toml:"session"`
	Cache            cache    `toml:"cache"`
}
type database struct {
	Url           string
	DbName        string `toml:"db_name"`
	DbHost        string `toml:"db_host"`
	DbPort        string `toml:"db_port"`
	DbUser        string `toml:"db_user"`
	DbPassword    string `toml:"db_password"`
	DbSsl         string `toml:"db_ssl"`
	MigrationsDir string `toml:"migrations_dir"`
}

type cache struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Password string `toml:"password"`
	Prefix   string `toml:"prefix"`
}

type session struct {
	Key    string
	Name   string
	MaxAge int
}

func NewConfig() *Config {
	return &Config{}
}
