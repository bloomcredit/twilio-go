/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Retrieve a specific Asset Version.
func (c *ApiService) FetchAssetVersion(ServiceSid string, AssetSid string, Sid string) (*ServerlessV1AssetVersion, error) {
	path := "/v1/Services/{ServiceSid}/Assets/{AssetSid}/Versions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"AssetSid"+"}", AssetSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1AssetVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAssetVersion'
type ListAssetVersionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAssetVersionParams) SetPageSize(PageSize int) *ListAssetVersionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAssetVersionParams) SetLimit(Limit int) *ListAssetVersionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AssetVersion records from the API. Request is executed immediately.
func (c *ApiService) PageAssetVersion(ServiceSid string, AssetSid string, params *ListAssetVersionParams, pageToken string, pageNumber string) (*ListAssetVersionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Assets/{AssetSid}/Versions"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"AssetSid"+"}", AssetSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAssetVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AssetVersion records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAssetVersion(ServiceSid string, AssetSid string, params *ListAssetVersionParams) ([]ServerlessV1AssetVersion, error) {
	if params == nil {
		params = &ListAssetVersionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAssetVersion(ServiceSid, AssetSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ServerlessV1AssetVersion

	for response != nil {
		records = append(records, response.AssetVersions...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListAssetVersionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListAssetVersionResponse)
	}

	return records, err
}

// Streams AssetVersion records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAssetVersion(ServiceSid string, AssetSid string, params *ListAssetVersionParams) (chan ServerlessV1AssetVersion, error) {
	if params == nil {
		params = &ListAssetVersionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAssetVersion(ServiceSid, AssetSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ServerlessV1AssetVersion, 1)

	go func() {
		for response != nil {
			for item := range response.AssetVersions {
				channel <- response.AssetVersions[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListAssetVersionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAssetVersionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAssetVersionResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAssetVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
