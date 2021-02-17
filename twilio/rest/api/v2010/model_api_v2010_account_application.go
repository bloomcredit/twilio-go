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

// ApiV2010AccountApplication struct for ApiV2010AccountApplication
type ApiV2010AccountApplication struct {
	AccountSid            string     `json:"AccountSid,omitempty"`
	ApiVersion            string     `json:"ApiVersion,omitempty"`
	DateCreated           string     `json:"DateCreated,omitempty"`
	DateUpdated           string     `json:"DateUpdated,omitempty"`
	FriendlyName          string     `json:"FriendlyName,omitempty"`
	MessageStatusCallback string     `json:"MessageStatusCallback,omitempty"`
	Sid                   string     `json:"Sid,omitempty"`
	SmsFallbackMethod     HttpMethod `json:"SmsFallbackMethod,omitempty"`
	SmsFallbackUrl        string     `json:"SmsFallbackUrl,omitempty"`
	SmsMethod             HttpMethod `json:"SmsMethod,omitempty"`
	SmsStatusCallback     string     `json:"SmsStatusCallback,omitempty"`
	SmsUrl                string     `json:"SmsUrl,omitempty"`
	StatusCallback        string     `json:"StatusCallback,omitempty"`
	StatusCallbackMethod  HttpMethod `json:"StatusCallbackMethod,omitempty"`
	Uri                   string     `json:"Uri,omitempty"`
	VoiceCallerIdLookup   bool       `json:"VoiceCallerIdLookup,omitempty"`
	VoiceFallbackMethod   HttpMethod `json:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackUrl      string     `json:"VoiceFallbackUrl,omitempty"`
	VoiceMethod           HttpMethod `json:"VoiceMethod,omitempty"`
	VoiceUrl              string     `json:"VoiceUrl,omitempty"`
}