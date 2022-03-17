// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

// TSIG defines the object used to provide a cryptographically secure method
// through which zone transfers may take place between a master and slave name
// server.
type TSIG struct {
	// Indicates a brief description for the TSIG key.
	Alias string `json:"Alias,omitempty"`

	// Identifies the key on the master name server and our Route name servers.
	// This name must be unique.
	KeyName string `json:"KeyName,omitempty"`

	// Identifies a hash value through which our name servers will be
	// authenticated to a master name server.
	KeyValue string `json:"KeyValue,omitempty"`

	// Identifies a cryptographic hash function by its system-defined ID.
	AlgorithmID TSIGAlgorithmType `json:"AlgorithmId,omitempty"`
}

// TSIGGetOK defines the additional parameters returned when retrieving a TSIG.
type TSIGGetOK struct {
	TSIG

	// Identifies a TSIG key by its system-defined ID.
	ID int `json:"Id,omitempty"`

	// Identifies the cryptographic hash function used to generate the key
	// value. Valid values are:
	// HMAC-MD5 | HMAC-SHA1 | HMAC-SHA256 | HMAC-SHA384 | HMAC-SHA224 |
	// HMAC-SHA512
	AlgorithmName string `json:"AlgorithmName,omitempty"`
}

//
// Enums
//

// TSIGAlgorithmType identifies the cryptographic hash function used to generate
// the key value.
type TSIGAlgorithmType int

const (
	HMAC_MD5 TSIGAlgorithmType = iota + 1
	HMAC_SHA1
	HMAC_SHA256
	HMAC_SHA384
	HMAC_SHA224
	HMAC_SHA512
)

//
// Params TSIG
//

func NewGetTSIGParams() *GetTSIGParams {
	return &GetTSIGParams{}
}

type GetTSIGParams struct {
	AccountNumber string
	TSIGID        int
}

func NewAddTSIGParams() *AddTSIGParams {
	return &AddTSIGParams{}
}

type AddTSIGParams struct {
	AccountNumber string
	TSIG          TSIG
}

func NewUpdateTSIGParams() *UpdateTSIGParams {
	return &UpdateTSIGParams{}
}

type UpdateTSIGParams struct {
	AccountNumber string
	TSIG          TSIGGetOK
}

func NewDeleteTSIGParams() *DeleteTSIGParams {
	return &DeleteTSIGParams{}
}

type DeleteTSIGParams struct {
	AccountNumber string
	TSIG          TSIGGetOK
}
