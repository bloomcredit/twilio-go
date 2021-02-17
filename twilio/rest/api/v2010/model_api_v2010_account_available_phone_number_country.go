/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ApiV2010AccountAvailablePhoneNumberCountry struct for ApiV2010AccountAvailablePhoneNumberCountry
type ApiV2010AccountAvailablePhoneNumberCountry struct {
	Beta            bool                   `json:"Beta,omitempty"`
	Country         string                 `json:"Country,omitempty"`
	CountryCode     string                 `json:"CountryCode,omitempty"`
	SubresourceUris map[string]interface{} `json:"SubresourceUris,omitempty"`
	Uri             string                 `json:"Uri,omitempty"`
}