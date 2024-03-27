/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents a Termination Notification.
type PcfSmPolicyControlTerminationNotification struct {
	// String providing an URI formatted according to RFC 3986.
	ResourceUri string                          `json:"resourceUri" yaml:"resourceUri" bson:"resourceUri,omitempty"`
	Cause       SmPolicyAssociationReleaseCause `json:"cause" yaml:"cause" bson:"cause,omitempty"`
}