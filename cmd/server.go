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
	"time"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start integrated server",
		Run:   serverAction,
		Args:  cobra.NoArgs,
	}

	defaultMetricsAddr = "0.0.0.0:8081"
	defaultServerAddr  = "0.0.0.0:8080"
	defaultServerCert  = ""
	defaultServerKey   = ""
	defaultLogLevel    = "info"
	defaultLogPretty   = true
	defaultLogColor    = true
)

func init() {
	rootCmd.AddCommand(serverCmd)

	//region Metrics
	serverCmd.PersistentFlags().String("metrics-addr", defaultMetricsAddr, "Address to bind the metrics")
	viper.SetDefault("metrics.addr", defaultMetricsAddr)
	_ = viper.BindPFlag("metrics.addr", serverCmd.PersistentFlags().Lookup("metrics-addr"))

	serverCmd.PersistentFlags().String("metrics-token", "", "Token to make metrics secure")
	viper.SetDefault("metrics.token", "")
	_ = viper.BindPFlag("metrics.token", serverCmd.PersistentFlags().Lookup("metrics-token"))
	//endregion Metrics

	//region Server
	serverCmd.PersistentFlags().String("server-addr", defaultServerAddr, "Address to bind the server")
	viper.SetDefault("server.addr", defaultServerAddr)
	_ = viper.BindPFlag("server.addr", serverCmd.PersistentFlags().Lookup("server-addr"))

	serverCmd.PersistentFlags().String("server-cert", defaultServerCert, "Path to cert for SSL encryption")
	viper.SetDefault("server.cert", defaultServerCert)
	_ = viper.BindPFlag("server.cert", serverCmd.PersistentFlags().Lookup("server-cert"))

	serverCmd.PersistentFlags().String("server-key", defaultServerKey, "Path to key for SSL encryption")
	viper.SetDefault("server.key", defaultServerKey)
	_ = viper.BindPFlag("server.key", serverCmd.PersistentFlags().Lookup("server-key"))
	//endregion Server

	//region Logger
	serverCmd.PersistentFlags().String("log-level", defaultLogLevel, "Log level (panic, fatal, error, warn, info, debug)")
	viper.SetDefault("log.level", defaultLogLevel)
	_ = viper.BindPFlag("log.level", serverCmd.PersistentFlags().Lookup("log-level"))

	serverCmd.PersistentFlags().Bool("log-pretty", defaultLogPretty, "Enable pretty logging output")
	viper.SetDefault("log.pretty", defaultLogPretty)
	_ = viper.BindPFlag("log.pretty", serverCmd.PersistentFlags().Lookup("log-pretty"))

	serverCmd.PersistentFlags().Bool("log-color", defaultLogColor, "Enable colored logging output")
	viper.SetDefault("log.color", defaultLogColor)
	_ = viper.BindPFlag("log.color", serverCmd.PersistentFlags().Lookup("log-color"))
	//endregion Logger
}

func serverAction(_ *cobra.Command, _ []string) {
	var gr run.Group

	// Setup grpc server
	lis, grpcServer := createGrpcServer()
	gr.Add(func() error {
		log.Info().
			Str("addr", cfg.Server.Addr).
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
			Str("addr", cfg.Metrics.Addr).
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
	lis, err := net.Listen("tcp", cfg.Server.Addr)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	return lis, server.NewGRPCServer(cfg)

}

func createMetricServer() http.Server {
	return http.Server{
		Addr:         cfg.Metrics.Addr,
		Handler:      server.Metrics(cfg),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
