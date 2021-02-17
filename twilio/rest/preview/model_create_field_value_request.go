/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateFieldValueRequest struct for CreateFieldValueRequest
type CreateFieldValueRequest struct {
	// An ISO language-country string of the value.
	Language string `json:"Language"`
	// A value that indicates this field value is a synonym of. Empty if the value is not a synonym.
	SynonymOf string `json:"SynonymOf,omitempty"`
	// A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long.
	Value string `json:"Value"`
}