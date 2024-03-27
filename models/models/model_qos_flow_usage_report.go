/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Contains QoS flows usage data information.
type QosFlowUsageReport struct {
	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	Qfi int32 `json:"qfi" yaml:"qfi" bson:"qfi,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StartTimeStamp *time.Time `json:"startTimeStamp" yaml:"startTimeStamp" bson:"startTimeStamp,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTimeStamp *time.Time `json:"endTimeStamp" yaml:"endTimeStamp" bson:"endTimeStamp,omitempty"`
	// string with format 'int64' as defined in OpenAPI.
	DownlinkVolume int64 `json:"downlinkVolume" yaml:"downlinkVolume" bson:"downlinkVolume,omitempty"`
	// string with format 'int64' as defined in OpenAPI.
	UplinkVolume int64 `json:"uplinkVolume" yaml:"uplinkVolume" bson:"uplinkVolume,omitempty"`
}