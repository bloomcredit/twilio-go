/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ServiceScanMessageContent the model 'ServiceScanMessageContent'
type ServiceScanMessageContent string

// List of service_scan_message_content
const (
	SERVICESCANMESSAGECONTENT_INHERIT ServiceScanMessageContent = "inherit"
	SERVICESCANMESSAGECONTENT_ENABLE  ServiceScanMessageContent = "enable"
	SERVICESCANMESSAGECONTENT_DISABLE ServiceScanMessageContent = "disable"
)