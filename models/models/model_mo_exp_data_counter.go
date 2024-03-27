/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Contain the MO Exception Data Counter.
type MoExpDataCounter struct {
	// Unsigned integer identifying the MO Exception Data Counter, as specified in clause 5.31.14.3 of 3GPP TS 23.501.
	Counter int32 `json:"counter" yaml:"counter" bson:"counter,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp *time.Time `json:"timeStamp,omitempty" yaml:"timeStamp" bson:"timeStamp,omitempty"`
}