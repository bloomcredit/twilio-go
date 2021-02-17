/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  *twilio.Client
}

func NewDefaultApiService(client *twilio.Client) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: fmt.Sprintf("https://insights.twilio.com"),
	}
}

/*
FetchCall Method for FetchCall
 * @param Sid
@return InsightsV1Call
*/
func (c *DefaultApiService) FetchCall(Sid string) (*InsightsV1Call, error) {
	path := "/v1/Voice/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1Call{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchSummaryParams Optional parameters for the method 'FetchSummary'
type FetchSummaryParams struct {
	ProcessingState *string `json:"ProcessingState,omitempty"`
}

/*
FetchSummary Method for FetchSummary
 * @param CallSid
 * @param optional nil or *FetchSummaryOpts - Optional Parameters:
 * @param "ProcessingState" (string) -
@return InsightsV1CallSummary
*/
func (c *DefaultApiService) FetchSummary(CallSid string, params *FetchSummaryParams) (*InsightsV1CallSummary, error) {
	path := "/v1/Voice/{CallSid}/Summary"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.ProcessingState != nil {
		data.Set("ProcessingState", *params.ProcessingState)
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1CallSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
FetchVideoParticipantSummary Method for FetchVideoParticipantSummary
 * @param RoomSid
 * @param ParticipantSid
@return InsightsV1VideoRoomSummaryVideoParticipantSummary
*/
func (c *DefaultApiService) FetchVideoParticipantSummary(RoomSid string, ParticipantSid string) (*InsightsV1VideoRoomSummaryVideoParticipantSummary, error) {
	path := "/v1/Video/Rooms/{RoomSid}/Participants/{ParticipantSid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1VideoRoomSummaryVideoParticipantSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
FetchVideoRoomSummary Method for FetchVideoRoomSummary
 * @param RoomSid
@return InsightsV1VideoRoomSummary
*/
func (c *DefaultApiService) FetchVideoRoomSummary(RoomSid string) (*InsightsV1VideoRoomSummary, error) {
	path := "/v1/Video/Rooms/{RoomSid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1VideoRoomSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListEventParams Optional parameters for the method 'ListEvent'
type ListEventParams struct {
	Edge     *string `json:"Edge,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty"`
}

/*
ListEvent Method for ListEvent
 * @param CallSid
 * @param optional nil or *ListEventOpts - Optional Parameters:
 * @param "Edge" (string) -
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListEventResponse
*/
func (c *DefaultApiService) ListEvent(CallSid string, params *ListEventParams) (*ListEventResponse, error) {
	path := "/v1/Voice/{CallSid}/Events"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.Edge != nil {
		data.Set("Edge", *params.Edge)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListMetricParams Optional parameters for the method 'ListMetric'
type ListMetricParams struct {
	Edge      *string `json:"Edge,omitempty"`
	Direction *string `json:"Direction,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty"`
}

/*
ListMetric Method for ListMetric
 * @param CallSid
 * @param optional nil or *ListMetricOpts - Optional Parameters:
 * @param "Edge" (string) -
 * @param "Direction" (string) -
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListMetricResponse
*/
func (c *DefaultApiService) ListMetric(CallSid string, params *ListMetricParams) (*ListMetricResponse, error) {
	path := "/v1/Voice/{CallSid}/Metrics"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.Edge != nil {
		data.Set("Edge", *params.Edge)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMetricResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListVideoParticipantSummaryParams Optional parameters for the method 'ListVideoParticipantSummary'
type ListVideoParticipantSummaryParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListVideoParticipantSummary Method for ListVideoParticipantSummary
 * @param RoomSid
 * @param optional nil or *ListVideoParticipantSummaryOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListVideoParticipantSummaryResponse
*/
func (c *DefaultApiService) ListVideoParticipantSummary(RoomSid string, params *ListVideoParticipantSummaryParams) (*ListVideoParticipantSummaryResponse, error) {
	path := "/v1/Video/Rooms/{RoomSid}/Participants"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVideoParticipantSummaryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListVideoRoomSummaryParams Optional parameters for the method 'ListVideoRoomSummary'
type ListVideoRoomSummaryParams struct {
	RoomType      *[]string  `json:"RoomType,omitempty"`
	Codec         *[]string  `json:"Codec,omitempty"`
	RoomName      *string    `json:"RoomName,omitempty"`
	CreatedAfter  *time.Time `json:"CreatedAfter,omitempty"`
	CreatedBefore *time.Time `json:"CreatedBefore,omitempty"`
	PageSize      *int32     `json:"PageSize,omitempty"`
}

/*
ListVideoRoomSummary Method for ListVideoRoomSummary
 * @param optional nil or *ListVideoRoomSummaryOpts - Optional Parameters:
 * @param "RoomType" ([]string) -
 * @param "Codec" ([]string) -
 * @param "RoomName" (string) -
 * @param "CreatedAfter" (time.Time) -
 * @param "CreatedBefore" (time.Time) -
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListVideoRoomSummaryResponse
*/
func (c *DefaultApiService) ListVideoRoomSummary(params *ListVideoRoomSummaryParams) (*ListVideoRoomSummaryResponse, error) {
	path := "/v1/Video/Rooms"

	data := url.Values{}
	headers := 0

	if params != nil && params.RoomType != nil {
		data.Set("RoomType", strings.Join(*params.RoomType, ","))
	}
	if params != nil && params.Codec != nil {
		data.Set("Codec", strings.Join(*params.Codec, ","))
	}
	if params != nil && params.RoomName != nil {
		data.Set("RoomName", *params.RoomName)
	}
	if params != nil && params.CreatedAfter != nil {
		data.Set("CreatedAfter", fmt.Sprint(*params.CreatedAfter))
	}
	if params != nil && params.CreatedBefore != nil {
		data.Set("CreatedBefore", fmt.Sprint(*params.CreatedBefore))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVideoRoomSummaryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}