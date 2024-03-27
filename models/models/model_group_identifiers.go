/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// External or Internal Group Identifier with a list of group members
type GroupIdentifiers struct {
	ExtGroupId string `json:"extGroupId,omitempty" yaml:"extGroupId" bson:"extGroupId,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.
	IntGroupId   string       `json:"intGroupId,omitempty" yaml:"intGroupId" bson:"intGroupId,omitempty"`
	UeIdList     []UdmSdmUeId `json:"ueIdList,omitempty" yaml:"ueIdList" bson:"ueIdList,omitempty"`
	AllowedAfIds []string     `json:"allowedAfIds,omitempty" yaml:"allowedAfIds" bson:"allowedAfIds,omitempty"`
}