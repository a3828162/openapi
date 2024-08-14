/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.7.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents other DN performance analytics requirements.
type DnPerformanceReq struct {
	DnPerfOrderCriter DnPerfOrderingCriterion `json:"dnPerfOrderCriter,omitempty" yaml:"dnPerfOrderCriter" bson:"dnPerfOrderCriter,omitempty"`
	Order             MatchingDirection       `json:"order,omitempty" yaml:"order" bson:"order,omitempty"`
	ReportThresholds  []ThresholdLevel        `json:"reportThresholds,omitempty" yaml:"reportThresholds" bson:"reportThresholds,omitempty"`
}