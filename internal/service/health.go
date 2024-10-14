package service

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/config"
	rpc "github.com/terrapi-solution/protocol/health"
	"net"
	"strconv"
	"time"
)

type HealthService struct {
	cfg *config.Config
}

func NewHealthService() *HealthService {
	return &HealthService{
		cfg: config.Load(),
	}
}

// CheckController checks the health of the controller service.
func (s *HealthService) CheckController() rpc.HealthCheck_ServingStatus {
	return rpc.HealthCheck_SERVING
}

// CheckDatabase checks the health of the database service.
func (s *HealthService) CheckDatabase() rpc.HealthCheck_ServingStatus {
	// Check the database service using a TCP health check
	if s.checkTCP(s.cfg.Datastore.Host, s.cfg.Datastore.Port, 5*time.Second) {
		return rpc.HealthCheck_SERVING
	} else {
		return rpc.HealthCheck_NOT_SERVING
	}
}

// CheckState checks the health of the state service.
func (s *HealthService) CheckState() rpc.HealthCheck_ServingStatus {
	// Return true if the state service is disabled,
	// In order to avoid send error to the client.
	if !s.cfg.State.Status {
		return rpc.HealthCheck_SERVING
	}

	// Check the state service using a TCP health check
	if s.checkTCP(s.cfg.State.Host, s.cfg.State.Port, 5*time.Second) {
		return rpc.HealthCheck_SERVING
	} else {
		return rpc.HealthCheck_NOT_SERVING
	}
}

// checkTCP performs a health check on a service using TCP
func (s *HealthService) checkTCP(host string, port int, timeout time.Duration) bool {
	// Create the address for the TCP connection
	address := net.JoinHostPort(host, strconv.Itoa(port))

	// Try to establish a TCP connection with the given timeout
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Error().Err(err).Msgf("unable to connect to service at %s:%d", host, port)
		return false
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msgf("unable to close connection to service at %s:%d", host, port)
		}
	}(conn)

	// Connection was successful
	return true
}
