/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListFleetResponse struct for ListFleetResponse
type ListFleetResponse struct {
	Fleets []SupersimV1Fleet `json:"Fleets,omitempty"`
	Meta ListCommandResponseMeta `json:"Meta,omitempty"`
}