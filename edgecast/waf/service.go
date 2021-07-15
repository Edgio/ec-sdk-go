// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package waf

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
)

// WAF service interacts with the EdgeCast API for WAF
type WAFService struct {
	*client.Client
	Logger edgecast.Logger
}

// New creates a new WAF service
func New(config WAFConfig) (*WAFService, error) {
	clientConfig, err := client.NewLegacyAPIClientConfig(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("error creating new WAF Service: %v", err)
	}

	// Inject the logger into the client config
	clientConfig.Logger = config.Logger

	return &WAFService{
		Client: client.NewClient(*clientConfig),
	}, nil
}
