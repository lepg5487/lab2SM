/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ErrorReport struct {

	Error ProblemDetails `json:"error,omitempty"`

	// Used to report the PCC rule failure.
	RuleReports []RuleReport `json:"ruleReports,omitempty"`

	// Used to report the session rule failure.
	SessRuleReports []SessionRuleReport `json:"sessRuleReports,omitempty"`
}