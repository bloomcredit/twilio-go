/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ServerlessV1ServiceAsset struct for ServerlessV1ServiceAsset
type ServerlessV1ServiceAsset struct {
	AccountSid   string                 `json:"AccountSid,omitempty"`
	DateCreated  time.Time              `json:"DateCreated,omitempty"`
	DateUpdated  time.Time              `json:"DateUpdated,omitempty"`
	FriendlyName string                 `json:"FriendlyName,omitempty"`
	Links        map[string]interface{} `json:"Links,omitempty"`
	ServiceSid   string                 `json:"ServiceSid,omitempty"`
	Sid          string                 `json:"Sid,omitempty"`
	Url          string                 `json:"Url,omitempty"`
}