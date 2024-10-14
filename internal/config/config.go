package config

// Config defines the general configuration
var globalConfig *Config

// Server defines the server configuration.
type Server struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
	Cert string `mapstructure:"cert"`
	Key  string `mapstructure:"key"`
}

// Datastore defines the database configuration.
type Datastore struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
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

// State defines the state server configuration.
type State struct {
	Status bool   `mapstructure:"status"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
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
	State     State     `mapstructure:"state"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}

func Set(config *Config) {
	// Set the global configuration
	globalConfig = config
}
