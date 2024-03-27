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

// Indicates the NAS Security Mode
type NasSecurityMode struct {
	IntegrityAlgorithm IntegrityAlgorithm `json:"integrityAlgorithm" yaml:"integrityAlgorithm" bson:"integrityAlgorithm,omitempty"`
	CipheringAlgorithm CipheringAlgorithm `json:"cipheringAlgorithm" yaml:"cipheringAlgorithm" bson:"cipheringAlgorithm,omitempty"`
}