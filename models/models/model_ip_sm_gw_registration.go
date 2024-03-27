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

type IpSmGwRegistration struct {
	IpSmGwMapAddress      string                      `json:"ipSmGwMapAddress,omitempty" yaml:"ipSmGwMapAddress" bson:"ipSmGwMapAddress,omitempty"`
	IpSmGwDiameterAddress *NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty" yaml:"ipSmGwDiameterAddress" bson:"ipSmGwDiameterAddress,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	IpsmgwIpv4 string `json:"ipsmgwIpv4,omitempty" yaml:"ipsmgwIpv4" bson:"ipsmgwIpv4,omitempty"`
	IpsmgwIpv6 string `json:"ipsmgwIpv6,omitempty" yaml:"ipsmgwIpv6" bson:"ipsmgwIpv6,omitempty"`
	// Fully Qualified Domain Name
	IpsmgwFqdn string `json:"ipsmgwFqdn,omitempty" yaml:"ipsmgwFqdn" bson:"ipsmgwFqdn,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfInstanceId    string   `json:"nfInstanceId,omitempty" yaml:"nfInstanceId" bson:"nfInstanceId,omitempty"`
	UnriIndicator   bool     `json:"unriIndicator,omitempty" yaml:"unriIndicator" bson:"unriIndicator,omitempty"`
	ResetIds        []string `json:"resetIds,omitempty" yaml:"resetIds" bson:"resetIds,omitempty"`
	IpSmGwSbiSupInd bool     `json:"ipSmGwSbiSupInd,omitempty" yaml:"ipSmGwSbiSupInd" bson:"ipSmGwSbiSupInd,omitempty"`
}