// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package auth

// Provider defines structs that can provide API Authorization credentials
type Provider interface {
	GetAuthorization() (string, error)
}
