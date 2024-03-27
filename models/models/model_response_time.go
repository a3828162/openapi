/*
 * LMF Location
 *
 * LMF Location Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.572 V17.9.0; 5G System; Location Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.572/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ResponseTime string

// List of ResponseTime
const (
	ResponseTime_LOW_DELAY      ResponseTime = "LOW_DELAY"
	ResponseTime_DELAY_TOLERANT ResponseTime = "DELAY_TOLERANT"
	ResponseTime_NO_DELAY       ResponseTime = "NO_DELAY"
)