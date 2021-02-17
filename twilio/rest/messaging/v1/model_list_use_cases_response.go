/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListUseCasesResponse struct for ListUseCasesResponse
type ListUseCasesResponse struct {
	Data []MessagingV1UseCases   `json:"Data,omitempty"`
	Meta ListServiceResponseMeta `json:"Meta,omitempty"`
}