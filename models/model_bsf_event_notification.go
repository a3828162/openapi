/*
 * Nbsf_Management
 *
 * Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.521 V17.7.0; 5G System; Binding Support Management Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.521/
 *
 * API version: 1.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains an event notification.
type BsfEventNotification struct {
	Event        BsfEvent      `json:"event" yaml:"event" bson:"event,omitempty"`
	PcfForUeInfo *PcfForUeInfo `json:"pcfForUeInfo,omitempty" yaml:"pcfForUeInfo" bson:"pcfForUeInfo,omitempty"`
	// The information of the PCF for a PDU session.
	PcfForPduSessInfos []PcfForPduSessionInfo `json:"pcfForPduSessInfos,omitempty" yaml:"pcfForPduSessInfos" bson:"pcfForPduSessInfos,omitempty"`
	// Matching S-NSSAI and DNN pairs.
	MatchSnssaiDnns []SnssaiDnnPair `json:"matchSnssaiDnns,omitempty" yaml:"matchSnssaiDnns" bson:"matchSnssaiDnns,omitempty"`
}