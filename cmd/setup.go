package cmd

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/terrapi-solution/controller/internal/config"
	"os"
	"strings"
)

func setupLogger() error {
	switch strings.ToLower(viper.GetString("log.level")) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Override the default logger with the custom logger.
	if viper.GetBool("log.pretty") {
		log.Logger = log.With().
			Timestamp().Caller().
			Logger().Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: !viper.GetBool("log.color"),
			},
		)
	} else {
		log.Logger = log.With().
			Timestamp().Caller().Logger()
	}

	return nil
}

func setupConfig() {
	if viper.GetString("config.file") != "" {
		viper.SetConfigFile(viper.GetString("config.file"))
	} else {
		viper.SetConfigName("controller")
		viper.AddConfigPath("/etc/terrapi")
		viper.AddConfigPath("$HOME/.terrapi")
		viper.AddConfigPath("./config")
	}

	viper.SetEnvPrefix("terrapi")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := readConfig(); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to read config file")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to parse config file")
	}

	// Set the global configuration.
	config.Set(cfg)
}

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
