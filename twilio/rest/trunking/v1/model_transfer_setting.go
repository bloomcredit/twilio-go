/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TransferSetting the model 'TransferSetting'
type TransferSetting string

// List of transfer_setting
const (
	TRANSFERSETTING_DISABLE_ALL TransferSetting = "disable-all"
	TRANSFERSETTING_ENABLE_ALL  TransferSetting = "enable-all"
	TRANSFERSETTING_SIP_ONLY    TransferSetting = "sip-only"
)