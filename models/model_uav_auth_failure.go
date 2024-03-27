/*
 * Nnef_Authentication
 *
 * NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.256 V17.3.0; 5G System;Uncrewed Aerial Systems Network Function (UAS-NF); Aerial Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.256/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// UAV auth failure
type UavAuthFailure struct {
	Error              *ProblemDetails `json:"error" yaml:"error" bson:"error,omitempty"`
	UasResourceRelease bool            `json:"uasResourceRelease,omitempty" yaml:"uasResourceRelease" bson:"uasResourceRelease,omitempty"`
}