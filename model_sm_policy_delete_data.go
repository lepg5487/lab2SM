/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type SmPolicyDeleteData struct {

	UserLocationInfo UserLocation `json:"userLocationInfo,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	ServingNetwork NetworkId `json:"servingNetwork,omitempty"`

	UserLocationInfoTime time.Time `json:"userLocationInfoTime,omitempty"`

	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`

	// Contains the usage report
	AccuUsageReports []AccuUsageReport `json:"accuUsageReports,omitempty"`
}
