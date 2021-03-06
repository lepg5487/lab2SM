/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TraceData struct {

	TraceRef string `json:"traceRef"`

	TraceDepth TraceDepth `json:"traceDepth"`

	NeTypeList string `json:"neTypeList"`

	EventList string `json:"eventList"`

	CollectionEntityIpv4Addr string `json:"collectionEntityIpv4Addr,omitempty"`

	CollectionEntityIpv6Addr Ipv6Addr `json:"collectionEntityIpv6Addr,omitempty"`

	InterfaceList string `json:"interfaceList,omitempty"`
}
