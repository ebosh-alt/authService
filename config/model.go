package config

type Config struct {
	Server   ServerConfig   `yaml:"Server"`
	Postgres PostgresConfig `yaml:"Postgres"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"DBName"`
	SSLMode  string `yaml:"sslMode"`
}

type ServerConfig struct {
	AppVersion string `yaml:"appVersion"`
	Host       string `yaml:"host" validate:"required"`
	Port       string `yaml:"port" validate:"required"`
}
