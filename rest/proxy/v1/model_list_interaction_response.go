/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListInteractionResponse struct for ListInteractionResponse
type ListInteractionResponse struct {
	Interactions []ProxyV1Interaction    `json:"interactions,omitempty"`
	Meta         ListServiceResponseMeta `json:"meta,omitempty"`
}
