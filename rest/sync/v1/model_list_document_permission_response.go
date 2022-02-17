/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListDocumentPermissionResponse struct for ListDocumentPermissionResponse
type ListDocumentPermissionResponse struct {
	Meta        ListServiceResponseMeta    `json:"meta,omitempty"`
	Permissions []SyncV1DocumentPermission `json:"permissions,omitempty"`
}
