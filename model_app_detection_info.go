/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AppDetectionInfo struct {

	// A reference to the application detection filter configured at the UPF
	AppId string `json:"appId"`

	// Identifier sent by the SMF in order to allow correlation of application Start and Stop events to the specific service data flow description, if service data flow descriptions are deducible.
	InstanceId string `json:"instanceId,omitempty"`

	// Contains the detected service data flow descriptions if they are deducible.
	SdfDescriptions []FlowInformation `json:"sdfDescriptions,omitempty"`
}
