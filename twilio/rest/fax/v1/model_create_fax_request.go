/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateFaxRequest struct for CreateFaxRequest
type CreateFaxRequest struct {
	// The number the fax was sent from. Can be the phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the SIP `from` value. The caller ID displayed to the recipient uses this value. If this is a phone number, it must be a Twilio number or a verified outgoing caller id from your account. If `to` is a SIP address, this can be any alphanumeric string (and also the characters `+`, `_`, `.`, and `-`), which will be used in the `from` header of the SIP request.
	From string `json:"From,omitempty"`
	// The URL of the PDF that contains the fax. See our [security](https://www.twilio.com/docs/usage/security) page for information on how to ensure the request for your media comes from Twilio.
	MediaUrl string `json:"MediaUrl"`
	// The [Fax Quality value](https://www.twilio.com/docs/fax/api/fax-resource#fax-quality-values) that describes the fax quality. Can be: `standard`, `fine`, or `superfine` and defaults to `fine`.
	Quality string `json:"Quality,omitempty"`
	// The password to use with `sip_auth_username` to authenticate faxes sent to a SIP address.
	SipAuthPassword string `json:"SipAuthPassword,omitempty"`
	// The username to use with the `sip_auth_password` to authenticate faxes sent to a SIP address.
	SipAuthUsername string `json:"SipAuthUsername,omitempty"`
	// The URL we should call using the `POST` method to send [status information](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-callback) to your application when the status of the fax changes.
	StatusCallback string `json:"StatusCallback,omitempty"`
	// Whether to store a copy of the sent media on our servers for later retrieval. Can be: `true` or `false` and the default is `true`.
	StoreMedia bool `json:"StoreMedia,omitempty"`
	// The phone number to receive the fax in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the recipient's SIP URI.
	To string `json:"To"`
	// How long in minutes from when the fax is initiated that we should try to send the fax.
	Ttl int32 `json:"Ttl,omitempty"`
}