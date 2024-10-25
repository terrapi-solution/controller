package config

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strings"
	"sync"
)

var (
	globalConfig *Config
	once         sync.Once
)

// GetInstance returns the configuration.
func GetInstance() *Config {
	// Make sure the configuration is only loaded once
	once.Do(func() {
		globalConfig = getConfig()
	})

	// Return the configuration
	return globalConfig
}

// getConfig gets the configuration from the config file.
func getConfig() *Config {
	cfg := &Config{}

	// Sets name for the config file.
	viper.SetConfigName("controller")

	// Adds a path for Viper to search for the config file in.
	viper.AddConfigPath("/etc/terrapi")
	viper.AddConfigPath("$HOME/.terrapi")
	viper.AddConfigPath("./config")

	// Set the prefix for environment variables
	viper.SetEnvPrefix("TERRAPI")

	// Replaces the "." in the environment variables with "_"
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Automatically read in environment variables that match
	viper.AutomaticEnv()

	// Read the configuration file
	if err := readConfig(); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to read config file")
	}

	// Unmarshal the configuration file into the config struct
	if err := viper.Unmarshal(cfg); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to parse config file")
	}

	// Return the configuration
	return cfg
}

// readConfig reads the configuration file.
func readConfig() error {
	err := viper.ReadInConfig()

	if err == nil {
		return nil
	}

	var configFileNotFoundError viper.ConfigFileNotFoundError
	if errors.As(err, &configFileNotFoundError) {
		return nil
	}

	var pathError *os.PathError
	if errors.As(err, &pathError) {
		return nil
	}

	return err
}
