// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package ecmodels

// PropagationStatus retrieves the propagation status for an edge CNAME
// configuration.
type PropagationStatus struct {
	// Indicates the edge CNAME 's current propagation status. Valid values are:
	// New: Indicates that the configuration has been created, but the
	// propagation process has not started.
	// propagating: Indicates that the configuration is in the process of being
	// propagated.
	// propagated: Indicates that the configuration has been propagated across
	// the entire network.
	Status string `json:"status,omitempty"`

	// Indicates the average configuration propagation percentage across all POPs.
	PercentPropagated float32 `json:"Percent_propagated,omitempty"`

	// Contains a list of POPs and their current configuration propagation
	// percentage.
	Pops []PopPropagationStatus `json:"Popsomitempty"`
}

// PopPropagationStatus represents propagation status for a POP
type PopPropagationStatus struct {
	// Identifies a POP by region and name.
	Name string `json:"name,omitempty"`

	// Indicates the percentage of servers within a POP to which the
	// configuration has been propagated.
	PercentPropagated float32 `json:"percentage_propagated,omitempty"`
}
