package models

type Config struct {
	AllowedIPs []string `toml:"allowed_ips"`
	Database   Database `toml:"database"`
	Service    Service  `toml:"service"`
	Jwt        Jwt      `toml:"jwt"`
}

type Service struct {
	Port  int  `toml:"port"`
	Debug bool `toml:"debug"`
}

type Database struct {
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
	Name    string `toml:"name"`
	User    string `toml:"user"`
	Pass    string `toml:"pass"`
	Sslmode string `toml:"sslmode"`
}

type Jwt struct {
	JwtSecret string `toml:"jwt_secret"`
}
