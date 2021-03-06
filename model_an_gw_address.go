/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AnGwAddress - describes the address of the access network gateway control node
type AnGwAddress struct {

	AnGwIpv4Addr string `json:"anGwIpv4Addr,omitempty"`

	AnGwIpv6Addr Ipv6Addr `json:"anGwIpv6Addr,omitempty"`
}
