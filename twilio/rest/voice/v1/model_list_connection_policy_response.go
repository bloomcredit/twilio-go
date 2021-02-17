/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListConnectionPolicyResponse struct for ListConnectionPolicyResponse
type ListConnectionPolicyResponse struct {
	ConnectionPolicies []VoiceV1ConnectionPolicy `json:"ConnectionPolicies,omitempty"`
	Meta               ListByocTrunkResponseMeta `json:"Meta,omitempty"`
}