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

// CreateSipCredentialListRequest struct for CreateSipCredentialListRequest
type CreateSipCredentialListRequest struct {
	// A human readable descriptive text that describes the CredentialList, up to 64 characters long.
	FriendlyName string `json:"FriendlyName"`
}