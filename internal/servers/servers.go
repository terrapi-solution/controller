package servers

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/core"
	"github.com/terrapi-solution/controller/internal/servers/grpc"
	"github.com/terrapi-solution/controller/internal/servers/metric"
	"github.com/terrapi-solution/controller/internal/servers/rest"
	"os"
	"os/signal"
	"syscall"
)

// Server represents the web services
type Server struct {
	Metric *metric.MetricServer
	Rest   *rest.RestServer
	Grpc   *grpc.GrpcServer
}

// StartServers starts the web services
func StartServers(coreService *core.Core) {
	log.Info().Msg("Initializing the web services...")
	s := &Server{
		Metric: metric.NewMetricServer(coreService.Config.Servers.Metric),
		Rest:   rest.NewRestServer(coreService.Config.Servers.Rest),
		Grpc:   grpc.NewGRPCServer(coreService.Config, coreService),
	}

	// Setup context and cancel function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Start the servers
	go func() {
		if err := s.Rest.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("Rest server encountered an error")
		}
	}()
	go func() {
		if err := s.Metric.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("Metric server encountered an error")
		}
	}()
	go func() {
		if err := s.Grpc.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("Grpc server encountered an error")
		}
	}()

	// Graceful shutdown
	defer func() {
		if err := s.Rest.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("Error shutting down the rest server")
		}
	}()
	defer func() {
		if err := s.Metric.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("Error shutting down the metric server")
		}
	}()
	defer func() {
		s.Grpc.Shutdown()
	}()

	// Wait for signals
	<-sigs
}
