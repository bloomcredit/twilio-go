/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV1VoiceCountry struct for PricingV1VoiceCountry
type PricingV1VoiceCountry struct {
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
