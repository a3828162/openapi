/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.13.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UdmSdmGroupIdentifiers struct {
	ExtGroupId string `json:"extGroupId,omitempty" yaml:"extGroupId" bson:"extGroupId,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.
	IntGroupId string       `json:"intGroupId,omitempty" yaml:"intGroupId" bson:"intGroupId,omitempty"`
	UeIdList   []UdmSdmUeId `json:"ueIdList,omitempty" yaml:"ueIdList" bson:"ueIdList,omitempty"`
}