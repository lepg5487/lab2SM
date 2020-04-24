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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{ service: s }
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"SmPoliciesPost",
			strings.ToUpper("Post"),
			"/npcf-smpolicycontrol/v1/sm-policies",
			c.SmPoliciesPost,
		},
		{
			"SmPoliciesSmPolicyIdDeletePost",
			strings.ToUpper("Post"),
			"/npcf-smpolicycontrol/v1/sm-policies/{smPolicyId}/delete",
			c.SmPoliciesSmPolicyIdDeletePost,
		},
		{
			"SmPoliciesSmPolicyIdGet",
			strings.ToUpper("Get"),
			"/npcf-smpolicycontrol/v1/sm-policies/{smPolicyId}",
			c.SmPoliciesSmPolicyIdGet,
		},
		{
			"SmPoliciesSmPolicyIdUpdatePost",
			strings.ToUpper("Post"),
			"/npcf-smpolicycontrol/v1/sm-policies/{smPolicyId}/update",
			c.SmPoliciesSmPolicyIdUpdatePost,
		},
	}
}

// SmPoliciesPost - 
func (c *DefaultApiController) SmPoliciesPost(w http.ResponseWriter, r *http.Request) { 
	smPolicyContextData := &SmPolicyContextData{}
	if err := json.NewDecoder(r.Body).Decode(&smPolicyContextData); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.SmPoliciesPost(*smPolicyContextData)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// SmPoliciesSmPolicyIdDeletePost - 
func (c *DefaultApiController) SmPoliciesSmPolicyIdDeletePost(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	smPolicyId := params["smPolicyId"]
	smPolicyDeleteData := &SmPolicyDeleteData{}
	if err := json.NewDecoder(r.Body).Decode(&smPolicyDeleteData); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.SmPoliciesSmPolicyIdDeletePost(smPolicyId, *smPolicyDeleteData)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// SmPoliciesSmPolicyIdGet - 
func (c *DefaultApiController) SmPoliciesSmPolicyIdGet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	smPolicyId := params["smPolicyId"]
	result, err := c.service.SmPoliciesSmPolicyIdGet(smPolicyId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// SmPoliciesSmPolicyIdUpdatePost - 
func (c *DefaultApiController) SmPoliciesSmPolicyIdUpdatePost(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	smPolicyId := params["smPolicyId"]
	smPolicyUpdateContextData := &SmPolicyUpdateContextData{}
	if err := json.NewDecoder(r.Body).Decode(&smPolicyUpdateContextData); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.SmPoliciesSmPolicyIdUpdatePost(smPolicyId, *smPolicyUpdateContextData)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}
