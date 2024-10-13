package config

// Server defines the server configuration.
type Server struct {
	Host       string `mapstructure:"host"`
	Port       uint   `mapstructure:"port"`
	PrivateKey string `mapstructure:"cert"`
	PublicKey  string `mapstructure:"key"`
}

// Metrics defines the metrics server configuration.
type Metrics struct {
	Addr  string `mapstructure:"addr"`
	Token string `mapstructure:"token"`
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// Auth defines the OpenID Connect configuration.
type Auth struct {
	Authority string
}

// Config defines the general configuration.
type Config struct {
	Server  Server  `mapstructure:"server"`
	Metrics Metrics `mapstructure:"metrics"`
	Logs    Logs    `mapstructure:"log"`
	Auth    Auth    `mapstructure:"auth"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
