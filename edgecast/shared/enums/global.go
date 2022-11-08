// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package enums

import "strings"

type Platform int

const (
	HttpLarge Platform = 3
	HttpSmall          = 8
	ADN                = 14
)

func (p Platform) String() string {
	switch p {
	case HttpLarge:
		return "http-large"
	case HttpSmall:
		return "http-small"
	case ADN:
		return "adn"
	}
	return "unknown"
}

func (p Platform) StringWithoutHyphen() string {
	return strings.Replace(p.String(), "-", "", 1)
}
