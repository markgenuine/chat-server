package env

import (
	"errors"
	"net"
	"os"

	"github.com/markgenuine/chat-server/internal/config"
)

var _ config.HTTPConfig = (*swaggerConfig)(nil)

const (
	swaggerHostEnvName = "SWAGGER_HOST"
	swaggerPortEnvName = "SWAGGER_PORT"
)

type swaggerConfig struct {
	host string
	port string
}

// NewSwaggerConfig ...
func NewSwaggerConfig() (*swaggerConfig, error) {
	host := os.Getenv(swaggerHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	port := os.Getenv(swaggerPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &swaggerConfig{
		host: host,
		port: port,
	}, nil
}

// Address ...
func (cfg *swaggerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
