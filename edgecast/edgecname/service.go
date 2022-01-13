// // Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// // See LICENSE file in project root for terms.

package edgecname

// import (
// 	"fmt"

// 	"github.com/EdgeCast/ec-sdk-go/edgecast"
// 	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
// 	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
// 	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
// )

// // Edge Cname service interacts with the EdgeCast API for managing Edge Cnames
// type EdgeCnameService struct {
// 	client.Client
// 	Logger logging.Logger
// }

// // New creates a new Edge Cname service
// func New(config edgecast.SDKConfig) (*EdgeCnameService, error) {

// 	authProvider, err := auth.NewTokenAuthorizationProvider(config.APIToken)

// 	if err != nil {
// 		return nil, fmt.Errorf("edgecname.New(): %v", err)
// 	}

// 	c := client.NewClient(client.ClientConfig{
// 		AuthProvider: authProvider,
// 		BaseAPIURL:   config.BaseAPIURLLegacy,
// 		UserAgent:    config.UserAgent,
// 		Logger:       config.Logger,
// 	})

// 	return &EdgeCnameService{
// 		Client: c,
// 		Logger: config.Logger,
// 	}, nil
// }
