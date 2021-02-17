/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListFieldResponse struct for ListFieldResponse
type ListFieldResponse struct {
	Fields []AutopilotV1AssistantTaskField `json:"Fields,omitempty"`
	Meta   ListAssistantResponseMeta       `json:"Meta,omitempty"`
}