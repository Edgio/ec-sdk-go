// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package auth

type LegacyAuthProvider struct {
	APIToken string
}

func (lp *LegacyAuthProvider) GetAuthorization() (string, error) {
	return "TOK:" + lp.APIToken, nil
}

func NewLegacyAuthProvider(apiToken string) *LegacyAuthProvider {
	return &LegacyAuthProvider{APIToken: apiToken}
}
