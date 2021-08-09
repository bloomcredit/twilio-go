/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListDayResponse struct for ListDayResponse
type ListDayResponse struct {
	Days []BulkexportsV1Day  `json:"days,omitempty"`
	Meta ListDayResponseMeta `json:"meta,omitempty"`
}
