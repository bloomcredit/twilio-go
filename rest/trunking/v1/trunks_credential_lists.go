/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'CreateCredentialList'
type CreateCredentialListParams struct {
	// The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list.
	CredentialListSid *string `json:"CredentialListSid,omitempty"`
}

func (params *CreateCredentialListParams) SetCredentialListSid(CredentialListSid string) *CreateCredentialListParams {
	params.CredentialListSid = &CredentialListSid
	return params
}

func (c *ApiService) CreateCredentialList(TrunkSid string, params *CreateCredentialListParams) (*TrunkingV1TrunkCredentialList, error) {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CredentialListSid != nil {
		data.Set("CredentialListSid", *params.CredentialListSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkCredentialList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteCredentialList(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchCredentialList(TrunkSid string, Sid string) (*TrunkingV1TrunkCredentialList, error) {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkCredentialList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCredentialList'
type ListCredentialListParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListCredentialListParams) SetPageSize(PageSize int) *ListCredentialListParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListCredentialList(TrunkSid string, params *ListCredentialListParams) (*ListCredentialListResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCredentialListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}