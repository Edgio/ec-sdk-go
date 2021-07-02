// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package waf

import (
	"github.com/VerizonDigital/ec-sdk-go/edgecast"
	"github.com/VerizonDigital/ec-sdk-go/edgecast/client"
)

// WAF service interacts with the EdgeCast API for WAF
type WAFService struct {
	*client.Client
}

// New creates a new WAF service
func New(apiToken string) *WAFService {
	return &WAFService{
		Client: client.DefaultLegacyClient(apiToken),
	}
}

// WithLogger can be used to specify a custom logger
func (w *WAFService) WithLogger(logger edgecast.Logger) *WAFService {
	w.Logger = logger
	return w
}
