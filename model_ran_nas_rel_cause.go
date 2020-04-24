/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type RanNasRelCause struct {

	NgApCause NgApCause `json:"ngApCause,omitempty"`

	Var5gMmCause int32 `json:"5gMmCause,omitempty"`

	Var5gSmCause int32 `json:"5gSmCause,omitempty"`
}
