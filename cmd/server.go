package cmd

import (
	"context"
	"github.com/oklog/run"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/terrapi-solution/controller/internal/server"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start integrated server",
		Run:   serverAction,
		Args:  cobra.NoArgs,
	}

	// Server configuration
	defaultServerHost = "0.0.0.0"
	defaultServerPort = 8080
	defaultServerMode = ""
	defaultServerCert = ""
	defaultServerKey  = ""

	// Datastore configuration
	defaultDatastoreHost     = "localhost"
	defaultDatastorePort     = 5432
	defaultDatastoreDatabase = "terrapi"
	defaultDatastoreUser     = ""
	defaultDatastorePassword = ""

	// Metric configuration
	defaultMetricStatus = true
	defaultMetricHost   = "localhost"
	defaultMetricPort   = 8081
	defaultMetricToken  = ""

	// Logger configuration
	defaultLogLevel  = "info"
	defaultLogPretty = true
	defaultLogColor  = true

	// Auth configuration
	defaultAuthAuthority = ""
)

// Initialization of CLI flags and viper config binding
func init() {
	rootCmd.AddCommand(serverCmd)

	// Server configuration
	serverCmd.PersistentFlags().String("server-host", defaultServerHost, "Host address to bind the server to")
	viper.SetDefault("server.host", defaultServerHost)
	_ = viper.BindPFlag("server.host", serverCmd.PersistentFlags().Lookup("server-host"))

	serverCmd.PersistentFlags().Int("server-port", defaultServerPort, "Port number for the server to listen on")
	viper.SetDefault("server.port", defaultServerPort)
	_ = viper.BindPFlag("server.port", serverCmd.PersistentFlags().Lookup("server-port"))

	serverCmd.PersistentFlags().String("server-mode", defaultServerMode, "Server mode (e.g., production, development)")
	viper.SetDefault("server.mode", defaultServerMode)
	_ = viper.BindPFlag("server.mode", serverCmd.PersistentFlags().Lookup("server-mode"))

	serverCmd.PersistentFlags().String("server-cert", defaultServerCert, "Path to SSL certificate file for secure connections")
	viper.SetDefault("server.cert", defaultServerCert)
	_ = viper.BindPFlag("server.cert", serverCmd.PersistentFlags().Lookup("server-cert"))

	serverCmd.PersistentFlags().String("server-key", defaultServerKey, "Path to SSL key file for secure connections")
	viper.SetDefault("server.key", defaultServerKey)
	_ = viper.BindPFlag("server.key", serverCmd.PersistentFlags().Lookup("server-key"))

	// Datastore configuration
	serverCmd.PersistentFlags().String("datastore-host", defaultDatastoreHost, "Host address of the datastore")
	viper.SetDefault("datastore.host", defaultDatastoreHost)
	_ = viper.BindPFlag("datastore.host", serverCmd.PersistentFlags().Lookup("datastore-host"))

	serverCmd.PersistentFlags().Int("datastore-port", defaultDatastorePort, "Port number to connect to the datastore")
	viper.SetDefault("datastore.port", defaultDatastorePort)
	_ = viper.BindPFlag("datastore.port", serverCmd.PersistentFlags().Lookup("datastore-port"))

	serverCmd.PersistentFlags().String("datastore-database", defaultDatastoreDatabase, "Name of the database to use in the datastore")
	viper.SetDefault("datastore.database", defaultDatastoreDatabase)
	_ = viper.BindPFlag("datastore.database", serverCmd.PersistentFlags().Lookup("datastore-database"))

	serverCmd.PersistentFlags().String("datastore-username", defaultDatastoreUser, "Username for connecting to the datastore")
	viper.SetDefault("datastore.username", defaultDatastoreUser)
	_ = viper.BindPFlag("datastore.username", serverCmd.PersistentFlags().Lookup("datastore-username"))

	serverCmd.PersistentFlags().String("datastore-password", defaultDatastorePassword, "Password for connecting to the datastore")
	viper.SetDefault("datastore.password", defaultDatastorePassword)
	_ = viper.BindPFlag("datastore.password", serverCmd.PersistentFlags().Lookup("datastore-password"))

	// Metric configuration
	serverCmd.PersistentFlags().Bool("metric-status", defaultMetricStatus, "Enable or disable metric collection")
	viper.SetDefault("metric.status", defaultMetricStatus)
	_ = viper.BindPFlag("metric.status", serverCmd.PersistentFlags().Lookup("metric-status"))

	serverCmd.PersistentFlags().String("metric-host", defaultMetricHost, "Host address for the metric service")
	viper.SetDefault("metric.host", defaultMetricHost)
	_ = viper.BindPFlag("metric.host", serverCmd.PersistentFlags().Lookup("metric-host"))

	serverCmd.PersistentFlags().Int("metric-port", defaultMetricPort, "Port number for the metric service")
	viper.SetDefault("metric.port", defaultMetricPort)
	_ = viper.BindPFlag("metric.port", serverCmd.PersistentFlags().Lookup("metric-port"))

	serverCmd.PersistentFlags().String("metric-token", defaultMetricToken, "Authentication token for accessing the metric service")
	viper.SetDefault("metric.token", defaultMetricToken)
	_ = viper.BindPFlag("metric.token", serverCmd.PersistentFlags().Lookup("metric-token"))

	// Logger configuration
	serverCmd.PersistentFlags().String("log-level", defaultLogLevel, "Log verbosity level (options: panic, fatal, error, warn, info, debug)")
	viper.SetDefault("log.level", defaultLogLevel)
	_ = viper.BindPFlag("log.level", serverCmd.PersistentFlags().Lookup("log-level"))

	serverCmd.PersistentFlags().Bool("log-pretty", defaultLogPretty, "Enable pretty-printed logging output")
	viper.SetDefault("log.pretty", defaultLogPretty)
	_ = viper.BindPFlag("log.pretty", serverCmd.PersistentFlags().Lookup("log-pretty"))

	serverCmd.PersistentFlags().Bool("log-color", defaultLogColor, "Enable colorized logging output")
	viper.SetDefault("log.color", defaultLogColor)
	_ = viper.BindPFlag("log.color", serverCmd.PersistentFlags().Lookup("log-color"))

	// Auth configuration
	serverCmd.PersistentFlags().String("auth-authority", defaultAuthAuthority, "Authority URL for authentication service")
	viper.SetDefault("auth.authority", defaultAuthAuthority)
	_ = viper.BindPFlag("auth.authority", serverCmd.PersistentFlags().Lookup("auth-authority"))
}

func serverAction(_ *cobra.Command, _ []string) {
	// Create a run group to manage the lifecycle of the application.
	var gr run.Group

	// Setup grpc server
	lis, grpcServer := createGrpcServer()
	gr.Add(func() error {
		log.Info().
			Str("host", cfg.Server.Host).
			Int("port", cfg.Server.Port).
			Msg("Starting grpc server")

		// Start the grpc server
		return grpcServer.Serve(lis)
	}, func(error) {
		log.Info().Msg("Shutting down grpc server")
		grpcServer.GracefulStop()
	})

	metricServer := createMetricServer()
	gr.Add(func() error {
		log.Info().
			Str("host", cfg.Metric.Host).
			Int("port", cfg.Metric.Port).
			Msg("Starting metrics server")

		return metricServer.ListenAndServe()
	}, func(error) {
		log.Info().Msg("Shutting down metric server")
		metricServer.Shutdown(context.Background())
	})

	// Start the run group
	if err := gr.Run(); err != nil {
		log.Fatal().Err(err).Msg("Error running the server")
		os.Exit(1)
	}
}

func createGrpcServer() (net.Listener, *grpc.Server) {
	address := net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.Port))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	return lis, server.NewGRPCServer(cfg)

}

func createMetricServer() http.Server {
	address := net.JoinHostPort(cfg.Metric.Host, strconv.Itoa(cfg.Metric.Port))
	return http.Server{
		Addr:         address,
		Handler:      server.Metrics(cfg),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
