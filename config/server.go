package config

import (
	"strings"
)

// ServerConfig exposes the server public method without exposing properties
type ServerConfig interface {
	GetProtocol() string
	GetDomain() string
	GetPort() string
	GetURI() string
}

// GetServerConfig returns the server config
func GetServerConfig() ServerConfig {
	return &s.Server
}

// public properties needed for json.Unmarshal
type server struct {
	Protocol string `json:"protocol"`
	Domain   string `json:"domain"`
	Port     string `json:"port"`
}

// GetProtocol exposes server.Protocol
func (s *server) GetProtocol() string {
	return s.Protocol
}

// GetDomain exposes server.Domain
func (s *server) GetDomain() string {
	return s.Domain
}

// GetPort exposes server.Port
func (s *server) GetPort() string {
	return s.Port
}

// GetURI returns server URL by combination of Protocol, Domain and Port
func (s *server) GetURI() string {
	var b strings.Builder
	b.WriteString(s.Protocol)
	b.WriteString(s.Domain)
	b.WriteString(":")
	b.WriteString(s.Port)

	return b.String()
}
