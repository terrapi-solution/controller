package service

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/config"
	rpc "github.com/terrapi-solution/protocol/health/v1"
	"net"
	"strconv"
	"time"
)

type HealthService struct {
	cfg *config.Config
}

func NewHealthService() *HealthService {
	return &HealthService{
		cfg: config.Get(),
	}
}

// CheckController checks the health of the controller service.
func (s *HealthService) CheckController() rpc.CheckResponse_ServingStatus {
	// Always return SERVING for the controller service
	// Because I'm the controller—obviously, everything revolves around me!
	return rpc.CheckResponse_SERVING_STATUS_SERVING
}

// CheckDatabase checks the health of the database service.
func (s *HealthService) CheckDatabase() rpc.CheckResponse_ServingStatus {
	// Check the database service using a TCP health check
	if s.checkTCP(s.cfg.Datastore.Host, s.cfg.Datastore.Port, 5*time.Second) {
		return rpc.CheckResponse_SERVING_STATUS_SERVING
	} else {
		return rpc.CheckResponse_SERVING_STATUS_NOT_SERVING
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
