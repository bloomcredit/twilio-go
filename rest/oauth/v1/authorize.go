/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Oauth
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
)

// Optional parameters for the method 'FetchAuthorize'
type FetchAuthorizeParams struct {
	// Response Type
	ResponseType *string `json:"ResponseType,omitempty"`
	// The Client Identifier
	ClientId *string `json:"ClientId,omitempty"`
	// The url to which response will be redirected to
	RedirectUri *string `json:"RedirectUri,omitempty"`
	// The scope of the access request
	Scope *string `json:"Scope,omitempty"`
	// An opaque value which can be used to maintain state between the request and callback
	State *string `json:"State,omitempty"`
}

func (params *FetchAuthorizeParams) SetResponseType(ResponseType string) *FetchAuthorizeParams {
	params.ResponseType = &ResponseType
	return params
}
func (params *FetchAuthorizeParams) SetClientId(ClientId string) *FetchAuthorizeParams {
	params.ClientId = &ClientId
	return params
}
func (params *FetchAuthorizeParams) SetRedirectUri(RedirectUri string) *FetchAuthorizeParams {
	params.RedirectUri = &RedirectUri
	return params
}
func (params *FetchAuthorizeParams) SetScope(Scope string) *FetchAuthorizeParams {
	params.Scope = &Scope
	return params
}
func (params *FetchAuthorizeParams) SetState(State string) *FetchAuthorizeParams {
	params.State = &State
	return params
}

// Retrieves authorize uri
func (c *ApiService) FetchAuthorize(params *FetchAuthorizeParams) (*OauthV1Authorize, error) {
	path := "/v1/authorize"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.ResponseType != nil {
		data.Set("ResponseType", *params.ResponseType)
	}
	if params != nil && params.ClientId != nil {
		data.Set("ClientId", *params.ClientId)
	}
	if params != nil && params.RedirectUri != nil {
		data.Set("RedirectUri", *params.RedirectUri)
	}
	if params != nil && params.Scope != nil {
		data.Set("Scope", *params.Scope)
	}
	if params != nil && params.State != nil {
		data.Set("State", *params.State)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &OauthV1Authorize{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}