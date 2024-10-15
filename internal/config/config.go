package config

// Config defines the general configuration
var globalConfig *Config

// Server defines the server configuration.
type Server struct {
	Host         string       `mapstructure:"host"`
	Port         int          `mapstructure:"port"`
	Certificates Certificates `mapstructure:"certificates"`
}

// Certificates defines the server credential configuration.
type Certificates struct {
	CertFile string `mapstructure:"cert"`
	KeyFile  string `mapstructure:"key"`
	CaFile   string `mapstructure:"ca"`
}

// Datastore defines the database configuration.
type Datastore struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// Metric defines the metrics server configuration.
type Metric struct {
	Status bool   `mapstructure:"status"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Token  string `mapstructure:"token"`
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// Config defines the general configuration.
type Config struct {
	Datastore Datastore `mapstructure:"datastore"`
	Logs      Logs      `mapstructure:"log"`
	Metric    Metric    `mapstructure:"metric"`
	Server    Server    `mapstructure:"server"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}

// Get returns the global configuration.
func Get() *Config {
	return globalConfig
}

// Set returns the global configuration.
func Set(config *Config) {
	globalConfig = config
}
