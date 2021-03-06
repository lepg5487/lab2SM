/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type N3gaLocation struct {

	N3gppTai Tai `json:"n3gppTai,omitempty"`

	N3IwfId string `json:"n3IwfId,omitempty"`

	UeIpv4Addr string `json:"ueIpv4Addr,omitempty"`

	UeIpv6Addr Ipv6Addr `json:"ueIpv6Addr,omitempty"`

	PortNumber int32 `json:"portNumber,omitempty"`
}
