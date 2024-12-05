/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateHostedNumberOrder'
type CreateHostedNumberOrderParams struct {
	// The number to host in [+E.164](https://en.wikipedia.org/wiki/E.164) format
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The contact phone number of the person authorized to sign the Authorization Document.
	ContactPhoneNumber *string `json:"ContactPhoneNumber,omitempty"`
	// Optional. A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number.
	AddressSid *string `json:"AddressSid,omitempty"`
	// Optional. Email of the owner of this phone number that is being hosted.
	Email *string `json:"Email,omitempty"`
	// This defaults to the AccountSid of the authorization the user is using. This can be provided to specify a subaccount to add the HostedNumberOrder to.
	AccountSid *string `json:"AccountSid,omitempty"`
	// A 128 character string that is a human readable text that describes this resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Optional. A list of emails that the LOA document for this HostedNumberOrder will be carbon copied to.
	CcEmails *[]string `json:"CcEmails,omitempty"`
	// The URL that Twilio should request when somebody sends an SMS to the phone number. This will be copied onto the IncomingPhoneNumber resource.
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The HTTP method that should be used to request the SmsUrl. Must be either `GET` or `POST`.  This will be copied onto the IncomingPhoneNumber resource.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl. This will be copied onto the IncomingPhoneNumber resource.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// Used to specify that the SMS capability will be hosted on Twilio's platform.
	SmsCapability *bool `json:"SmsCapability,omitempty"`
	// The HTTP method that should be used to request the SmsFallbackUrl. Must be either `GET` or `POST`. This will be copied onto the IncomingPhoneNumber resource.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// Optional. The Status Callback URL attached to the IncomingPhoneNumber resource.
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty"`
	// Optional. The Status Callback Method attached to the IncomingPhoneNumber resource.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// Optional. The 34 character sid of the application Twilio should use to handle SMS messages sent to this number. If a `SmsApplicationSid` is present, Twilio will ignore all of the SMS urls above and use those set on the application.
	SmsApplicationSid *string `json:"SmsApplicationSid,omitempty"`
	// The title of the person authorized to sign the Authorization Document for this phone number.
	ContactTitle *string `json:"ContactTitle,omitempty"`
}

func (params *CreateHostedNumberOrderParams) SetPhoneNumber(PhoneNumber string) *CreateHostedNumberOrderParams {
	params.PhoneNumber = &PhoneNumber
	return params
}
func (params *CreateHostedNumberOrderParams) SetContactPhoneNumber(ContactPhoneNumber string) *CreateHostedNumberOrderParams {
	params.ContactPhoneNumber = &ContactPhoneNumber
	return params
}
func (params *CreateHostedNumberOrderParams) SetAddressSid(AddressSid string) *CreateHostedNumberOrderParams {
	params.AddressSid = &AddressSid
	return params
}
func (params *CreateHostedNumberOrderParams) SetEmail(Email string) *CreateHostedNumberOrderParams {
	params.Email = &Email
	return params
}
func (params *CreateHostedNumberOrderParams) SetAccountSid(AccountSid string) *CreateHostedNumberOrderParams {
	params.AccountSid = &AccountSid
	return params
}
func (params *CreateHostedNumberOrderParams) SetFriendlyName(FriendlyName string) *CreateHostedNumberOrderParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateHostedNumberOrderParams) SetCcEmails(CcEmails []string) *CreateHostedNumberOrderParams {
	params.CcEmails = &CcEmails
	return params
}
func (params *CreateHostedNumberOrderParams) SetSmsUrl(SmsUrl string) *CreateHostedNumberOrderParams {
	params.SmsUrl = &SmsUrl
	return params
}
func (params *CreateHostedNumberOrderParams) SetSmsMethod(SmsMethod string) *CreateHostedNumberOrderParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *CreateHostedNumberOrderParams) SetSmsFallbackUrl(SmsFallbackUrl string) *CreateHostedNumberOrderParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *CreateHostedNumberOrderParams) SetSmsCapability(SmsCapability bool) *CreateHostedNumberOrderParams {
	params.SmsCapability = &SmsCapability
	return params
}
func (params *CreateHostedNumberOrderParams) SetSmsFallbackMethod(SmsFallbackMethod string) *CreateHostedNumberOrderParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}
func (params *CreateHostedNumberOrderParams) SetStatusCallbackUrl(StatusCallbackUrl string) *CreateHostedNumberOrderParams {
	params.StatusCallbackUrl = &StatusCallbackUrl
	return params
}
func (params *CreateHostedNumberOrderParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateHostedNumberOrderParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateHostedNumberOrderParams) SetSmsApplicationSid(SmsApplicationSid string) *CreateHostedNumberOrderParams {
	params.SmsApplicationSid = &SmsApplicationSid
	return params
}
func (params *CreateHostedNumberOrderParams) SetContactTitle(ContactTitle string) *CreateHostedNumberOrderParams {
	params.ContactTitle = &ContactTitle
	return params
}

// Host a phone number's capability on Twilio's platform.
func (c *ApiService) CreateHostedNumberOrder(params *CreateHostedNumberOrderParams) (*NumbersV2HostedNumberOrder, error) {
	path := "/v2/HostedNumber/Orders"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}
	if params != nil && params.ContactPhoneNumber != nil {
		data.Set("ContactPhoneNumber", *params.ContactPhoneNumber)
	}
	if params != nil && params.AddressSid != nil {
		data.Set("AddressSid", *params.AddressSid)
	}
	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.CcEmails != nil {
		for _, item := range *params.CcEmails {
			data.Add("CcEmails", item)
		}
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsCapability != nil {
		data.Set("SmsCapability", fmt.Sprint(*params.SmsCapability))
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.StatusCallbackUrl != nil {
		data.Set("StatusCallbackUrl", *params.StatusCallbackUrl)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.SmsApplicationSid != nil {
		data.Set("SmsApplicationSid", *params.SmsApplicationSid)
	}
	if params != nil && params.ContactTitle != nil {
		data.Set("ContactTitle", *params.ContactTitle)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2HostedNumberOrder{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Cancel the HostedNumberOrder (only available when the status is in `received`).
func (c *ApiService) DeleteHostedNumberOrder(Sid string) error {
	path := "/v2/HostedNumber/Orders/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific HostedNumberOrder.
func (c *ApiService) FetchHostedNumberOrder(Sid string) (*NumbersV2HostedNumberOrder, error) {
	path := "/v2/HostedNumber/Orders/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2HostedNumberOrder{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListHostedNumberOrder'
type ListHostedNumberOrderParams struct {
	// The Status of this HostedNumberOrder. One of `received`, `pending-verification`, `verified`, `pending-loa`, `carrier-processing`, `testing`, `completed`, `failed`, or `action-required`.
	Status *string `json:"Status,omitempty"`
	// Whether the SMS capability will be hosted on our platform. Can be `true` of `false`.
	SmsCapability *bool `json:"SmsCapability,omitempty"`
	// An E164 formatted phone number hosted by this HostedNumberOrder.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder.
	IncomingPhoneNumberSid *string `json:"IncomingPhoneNumberSid,omitempty"`
	// A human readable description of this resource, up to 128 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListHostedNumberOrderParams) SetStatus(Status string) *ListHostedNumberOrderParams {
	params.Status = &Status
	return params
}
func (params *ListHostedNumberOrderParams) SetSmsCapability(SmsCapability bool) *ListHostedNumberOrderParams {
	params.SmsCapability = &SmsCapability
	return params
}
func (params *ListHostedNumberOrderParams) SetPhoneNumber(PhoneNumber string) *ListHostedNumberOrderParams {
	params.PhoneNumber = &PhoneNumber
	return params
}
func (params *ListHostedNumberOrderParams) SetIncomingPhoneNumberSid(IncomingPhoneNumberSid string) *ListHostedNumberOrderParams {
	params.IncomingPhoneNumberSid = &IncomingPhoneNumberSid
	return params
}
func (params *ListHostedNumberOrderParams) SetFriendlyName(FriendlyName string) *ListHostedNumberOrderParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListHostedNumberOrderParams) SetPageSize(PageSize int) *ListHostedNumberOrderParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListHostedNumberOrderParams) SetLimit(Limit int) *ListHostedNumberOrderParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of HostedNumberOrder records from the API. Request is executed immediately.
func (c *ApiService) PageHostedNumberOrder(params *ListHostedNumberOrderParams, pageToken, pageNumber string) (*ListHostedNumberOrderResponse, error) {
	path := "/v2/HostedNumber/Orders"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.SmsCapability != nil {
		data.Set("SmsCapability", fmt.Sprint(*params.SmsCapability))
	}
	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}
	if params != nil && params.IncomingPhoneNumberSid != nil {
		data.Set("IncomingPhoneNumberSid", *params.IncomingPhoneNumberSid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
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

	ps := &ListHostedNumberOrderResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists HostedNumberOrder records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListHostedNumberOrder(params *ListHostedNumberOrderParams) ([]NumbersV2HostedNumberOrder, error) {
	response, errors := c.StreamHostedNumberOrder(params)

	records := make([]NumbersV2HostedNumberOrder, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams HostedNumberOrder records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamHostedNumberOrder(params *ListHostedNumberOrderParams) (chan NumbersV2HostedNumberOrder, chan error) {
	if params == nil {
		params = &ListHostedNumberOrderParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2HostedNumberOrder, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageHostedNumberOrder(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamHostedNumberOrder(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamHostedNumberOrder(response *ListHostedNumberOrderResponse, params *ListHostedNumberOrderParams, recordChannel chan NumbersV2HostedNumberOrder, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Items
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListHostedNumberOrderResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListHostedNumberOrderResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListHostedNumberOrderResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListHostedNumberOrderResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateHostedNumberOrder'
type UpdateHostedNumberOrderParams struct {
	//
	Status *string `json:"Status,omitempty"`
	// The number of seconds to wait before initiating the ownership verification call. Can be a value between 0 and 60, inclusive.
	VerificationCallDelay *int `json:"VerificationCallDelay,omitempty"`
	// The numerical extension to dial when making the ownership verification call.
	VerificationCallExtension *string `json:"VerificationCallExtension,omitempty"`
}

func (params *UpdateHostedNumberOrderParams) SetStatus(Status string) *UpdateHostedNumberOrderParams {
	params.Status = &Status
	return params
}
func (params *UpdateHostedNumberOrderParams) SetVerificationCallDelay(VerificationCallDelay int) *UpdateHostedNumberOrderParams {
	params.VerificationCallDelay = &VerificationCallDelay
	return params
}
func (params *UpdateHostedNumberOrderParams) SetVerificationCallExtension(VerificationCallExtension string) *UpdateHostedNumberOrderParams {
	params.VerificationCallExtension = &VerificationCallExtension
	return params
}

// Updates a specific HostedNumberOrder.
func (c *ApiService) UpdateHostedNumberOrder(Sid string, params *UpdateHostedNumberOrderParams) (*NumbersV2HostedNumberOrder, error) {
	path := "/v2/HostedNumber/Orders/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.VerificationCallDelay != nil {
		data.Set("VerificationCallDelay", fmt.Sprint(*params.VerificationCallDelay))
	}
	if params != nil && params.VerificationCallExtension != nil {
		data.Set("VerificationCallExtension", *params.VerificationCallExtension)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2HostedNumberOrder{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
