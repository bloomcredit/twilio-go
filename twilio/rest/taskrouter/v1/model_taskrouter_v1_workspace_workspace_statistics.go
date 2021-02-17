/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TaskrouterV1WorkspaceWorkspaceStatistics struct for TaskrouterV1WorkspaceWorkspaceStatistics
type TaskrouterV1WorkspaceWorkspaceStatistics struct {
	AccountSid   string                 `json:"AccountSid,omitempty"`
	Cumulative   map[string]interface{} `json:"Cumulative,omitempty"`
	Realtime     map[string]interface{} `json:"Realtime,omitempty"`
	Url          string                 `json:"Url,omitempty"`
	WorkspaceSid string                 `json:"WorkspaceSid,omitempty"`
}