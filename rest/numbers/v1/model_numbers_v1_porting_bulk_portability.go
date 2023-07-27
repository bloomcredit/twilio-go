/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// NumbersV1PortingBulkPortability struct for NumbersV1PortingBulkPortability
type NumbersV1PortingBulkPortability struct {
	// A 34 character string that uniquely identifies this Portability check.
	Sid    *string `json:"sid,omitempty"`
	Status *string `json:"status,omitempty"`
	// The date that the Portability check was created, given in ISO 8601 format.
	DatetimeCreated *time.Time `json:"datetime_created,omitempty"`
	// Contains a list with all the information of the requested phone numbers. Each phone number contains the following properties: `phone_number`: The phone number which portability is to be checked. `portable`: Boolean flag specifying if phone number is portable or not. `not_portable_reason`: Reason why the phone number cannot be ported into Twilio, `null` otherwise. `not_portable_reason_code`: The Portability Reason Code for the phone number if it cannot be ported in Twilio, `null` otherwise. `pin_and_account_number_required`: Boolean flag specifying if PIN and account number is required for the phone number. `number_type`: The type of the requested phone number. `country` Country the phone number belongs to. `messaging_carrier` Current messaging carrier of the phone number. `voice_carrier` Current voice carrier of the phone number.
	PhoneNumbers *[]interface{} `json:"phone_numbers,omitempty"`
	// This is the url of the request that you're trying to reach out to locate the resource.
	Url *string `json:"url,omitempty"`
}