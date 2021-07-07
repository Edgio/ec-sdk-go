package waf

import "github.com/VerizonDigital/ec-sdk-go/edgecast"

// Config holds the configuration for the WAF service
type WAFConfig struct {
	Logger   edgecast.Logger
	APIToken string
}

// NewConfig creates the default configuration for the WAF service
func NewConfig(apiToken string) WAFConfig {
	return WAFConfig{
		Logger:   edgecast.NullLogger{},
		APIToken: apiToken,
	}
}
