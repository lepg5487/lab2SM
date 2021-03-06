/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type FlowInformation struct {

	// Defines a packet filter for an IP flow.Refer to subclause 5.4.2 of 3GPP TS 29.212 [23] for encoding.
	FlowDescription string `json:"flowDescription,omitempty"`

	EthFlowDescription EthFlowDescription `json:"ethFlowDescription,omitempty"`

	// An identifier of packet filter.
	PackFiltId string `json:"packFiltId,omitempty"`

	// The packet shall be sent to the UE.
	PacketFilterUsage bool `json:"packetFilterUsage,omitempty"`

	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field.
	TosTrafficClass *string `json:"tosTrafficClass,omitempty"`

	// the security parameter index of the IPSec packet.
	Spi *string `json:"spi,omitempty"`

	// the Ipv6 flow label header field.
	FlowLabel *string `json:"flowLabel,omitempty"`

	FlowDirection *FlowDirectionRm `json:"flowDirection,omitempty"`
}
