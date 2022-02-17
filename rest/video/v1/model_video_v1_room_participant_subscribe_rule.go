/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VideoV1RoomParticipantSubscribeRule struct for VideoV1RoomParticipantSubscribeRule
type VideoV1RoomParticipantSubscribeRule struct {
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Participant resource for the Subscribe Rules
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The SID of the Room resource for the Subscribe Rules
	RoomSid *string `json:"room_sid,omitempty"`
	// A collection of Subscribe Rules that describe how to include or exclude matching tracks
	Rules *[]map[string]interface{} `json:"rules,omitempty"`
}
