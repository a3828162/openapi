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

type TraceDataResponse struct {
	TraceData         *TraceData `json:"traceData,omitempty" yaml:"traceData" bson:"traceData,omitempty"`
	SharedTraceDataId string     `json:"sharedTraceDataId,omitempty" yaml:"sharedTraceDataId" bson:"sharedTraceDataId,omitempty"`
}