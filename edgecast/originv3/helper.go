// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package originv3

// SetShieldPOPsFromEdgeNodes is a convenience function for setting a group's
// Shield POPs from a slice of OriginShieldEdgeNodes.
func (g CustomerOriginGroupHTTPRequest) SetShieldPOPsFromEdgeNodes(
	edgeNodes []OriginShieldEdgeNode,
) {
	shieldPOPs := make([]*string, 0)

	for _, n := range edgeNodes {
		for _, p := range n.Pops {
			shieldPOPs = append(shieldPOPs, p.Code)
		}
	}

	g.ShieldPops = shieldPOPs
}
