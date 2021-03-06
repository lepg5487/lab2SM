/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UeCampingRep struct {

	AccessType AccessType `json:"accessType,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	ServNfId ServingNfIdentity `json:"servNfId,omitempty"`

	ServingNetwork NetworkId `json:"servingNetwork,omitempty"`

	UserLocationInfo UserLocation `json:"userLocationInfo,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`
}
