/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TwilioEdge the model 'TwilioEdge'
type TwilioEdge string

// List of twilio_edge
const (
	TWILIOEDGE_UNKNOWN_EDGE TwilioEdge = "unknown_edge"
	TWILIOEDGE_CARRIER_EDGE TwilioEdge = "carrier_edge"
	TWILIOEDGE_SIP_EDGE     TwilioEdge = "sip_edge"
	TWILIOEDGE_SDK_EDGE     TwilioEdge = "sdk_edge"
	TWILIOEDGE_CLIENT_EDGE  TwilioEdge = "client_edge"
)