package config

// Server defines the server configuration.
type Server struct {
	Host       string `mapstructure:"host"`
	Port       uint   `mapstructure:"port"`
	PrivateKey string `mapstructure:"cert"`
	PublicKey  string `mapstructure:"key"`
}

// Datastore defines the database configuration.
type Datastore struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
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
	Auth      Auth      `mapstructure:"auth"`
	Datastore Datastore `mapstructure:"datastore"`
	Logs      Logs      `mapstructure:"log"`
	Metrics   Metrics   `mapstructure:"metrics"`
	Server    Server    `mapstructure:"server"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
