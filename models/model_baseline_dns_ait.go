/*
 * Neasdf_BaselineDNSPattern
 *
 * EASDF Baseline DNS Pattern Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.556 V17.3.0; 5G System; Edge Application Server Discovery Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.556/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Baseline DNS action information Template
type BaselineDnsAit struct {
	AitId                string     `json:"aitId" yaml:"aitId" bson:"aitId,omitempty"`
	Label                string     `json:"label,omitempty" yaml:"label" bson:"label,omitempty"`
	EcsOption            *EcsOption `json:"ecsOption,omitempty" yaml:"ecsOption" bson:"ecsOption,omitempty"`
	DnsServerAddressList []IpAddr   `json:"dnsServerAddressList,omitempty" yaml:"dnsServerAddressList" bson:"dnsServerAddressList,omitempty"`
}