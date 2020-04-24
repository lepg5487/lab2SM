/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PacketFilterInfo struct {

	// An identifier of packet filter.
	PackFiltId string `json:"packFiltId,omitempty"`

	// Defines a packet filter for an IP flow.Refer to subclause 5.3.54 of 3GPP TS 29.212 [23] for encoding.
	PackFiltCont string `json:"packFiltCont,omitempty"`

	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field.
	TosTrafficClass string `json:"tosTrafficClass,omitempty"`

	// The security parameter index of the IPSec packet.
	Spi string `json:"spi,omitempty"`

	// The Ipv6 flow label header field.
	FlowLabel string `json:"flowLabel,omitempty"`

	FlowDirection FlowDirection `json:"flowDirection,omitempty"`
}
