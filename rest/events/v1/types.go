/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
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

// Fetch a specific Event Type.
func (c *ApiService) FetchEventType(Type string) (*EventsV1EventType, error) {
	path := "/v1/Types/{Type}"
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1EventType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEventType'
type ListEventTypeParams struct {
	// A string parameter filtering the results to return only the Event Types using a given schema.
	SchemaId *string `json:"SchemaId,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEventTypeParams) SetSchemaId(SchemaId string) *ListEventTypeParams {
	params.SchemaId = &SchemaId
	return params
}
func (params *ListEventTypeParams) SetPageSize(PageSize int) *ListEventTypeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEventTypeParams) SetLimit(Limit int) *ListEventTypeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of EventType records from the API. Request is executed immediately.
func (c *ApiService) PageEventType(params *ListEventTypeParams, pageToken, pageNumber string) (*ListEventTypeResponse, error) {
	path := "/v1/Types"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.SchemaId != nil {
		data.Set("SchemaId", *params.SchemaId)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEventTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists EventType records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEventType(params *ListEventTypeParams) ([]EventsV1EventType, error) {
	if params == nil {
		params = &ListEventTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEventType(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []EventsV1EventType

	for response != nil {
		records = append(records, response.Types...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListEventTypeResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListEventTypeResponse)
	}

	return records, err
}

// Streams EventType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEventType(params *ListEventTypeParams) (chan EventsV1EventType, error) {
	if params == nil {
		params = &ListEventTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEventType(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan EventsV1EventType, 1)

	go func() {
		for response != nil {
			for item := range response.Types {
				channel <- response.Types[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListEventTypeResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListEventTypeResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListEventTypeResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEventTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
