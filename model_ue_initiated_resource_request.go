/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UeInitiatedResourceRequest struct {

	PccRuleId string `json:"pccRuleId,omitempty"`

	RuleOp RuleOperation `json:"ruleOp"`

	Precedence int32 `json:"precedence,omitempty"`

	PackFiltInfo []PacketFilterInfo `json:"packFiltInfo"`

	ReqQos RequestedQos `json:"reqQos,omitempty"`
}