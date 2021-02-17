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

// CreateKeyRequest struct for CreateKeyRequest
type CreateKeyRequest struct {
	// Provides the unique string identifier of an existing Device to become authenticated with this Key credential.
	DeviceSid string `json:"DeviceSid,omitempty"`
	// Provides a human readable descriptive text for this Key credential, up to 256 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
}