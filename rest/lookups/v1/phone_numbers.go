/*
 * Twilio - Lookups
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
	"net/url"
	"strings"
)

// Optional parameters for the method 'FetchPhoneNumber'
type FetchPhoneNumberParams struct {
	// The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the phone number to fetch. This is used to specify the country when the phone number is provided in a national format.
	CountryCode *string `json:"CountryCode,omitempty"`
	// The type of information to return. Can be: `carrier` or `caller-name`. The default is null.  Carrier information costs $0.005 per phone number looked up.  Caller Name information is currently available only in the US and costs $0.01 per phone number looked up.  To retrieve both types on information, specify this parameter twice; once with `carrier` and once with `caller-name` as the value.
	Type *[]string `json:"Type,omitempty"`
	// The `unique_name` of an Add-on you would like to invoke. Can be the `unique_name` of an Add-on that is installed on your account. You can specify multiple instances of this parameter to invoke multiple Add-ons. For more information about  Add-ons, see the [Add-ons documentation](https://www.twilio.com/docs/add-ons).
	AddOns *[]string `json:"AddOns,omitempty"`
	// Data specific to the add-on you would like to invoke. The content and format of this value depends on the add-on.
	AddOnsData *map[string]interface{} `json:"AddOnsData,omitempty"`
}

func (params *FetchPhoneNumberParams) SetCountryCode(CountryCode string) *FetchPhoneNumberParams {
	params.CountryCode = &CountryCode
	return params
}
func (params *FetchPhoneNumberParams) SetType(Type []string) *FetchPhoneNumberParams {
	params.Type = &Type
	return params
}
func (params *FetchPhoneNumberParams) SetAddOns(AddOns []string) *FetchPhoneNumberParams {
	params.AddOns = &AddOns
	return params
}
func (params *FetchPhoneNumberParams) SetAddOnsData(AddOnsData map[string]interface{}) *FetchPhoneNumberParams {
	params.AddOnsData = &AddOnsData
	return params
}

func (c *ApiService) FetchPhoneNumber(PhoneNumber string, params *FetchPhoneNumberParams) (*LookupsV1PhoneNumber, error) {
	path := "/v1/PhoneNumbers/{PhoneNumber}"
	path = strings.Replace(path, "{"+"PhoneNumber"+"}", PhoneNumber, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CountryCode != nil {
		data.Set("CountryCode", *params.CountryCode)
	}
	if params != nil && params.Type != nil {
		for _, item := range *params.Type {
			data.Add("Type", item)
		}
	}
	if params != nil && params.AddOns != nil {
		for _, item := range *params.AddOns {
			data.Add("AddOns", item)
		}
	}
	if params != nil && params.AddOnsData != nil {
		v, err := json.Marshal(params.AddOnsData)

		if err != nil {
			return nil, err
		}

		data.Set("AddOnsData", string(v))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &LookupsV1PhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
