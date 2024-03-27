/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Set of parameters supported by UPF for a given DNN
type DnnUpfInfoItem struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn               string                            `json:"dnn" yaml:"dnn" bson:"dnn,omitempty"`
	DnaiList          []string                          `json:"dnaiList,omitempty" yaml:"dnaiList" bson:"dnaiList,omitempty"`
	PduSessionTypes   []PduSessionType                  `json:"pduSessionTypes,omitempty" yaml:"pduSessionTypes" bson:"pduSessionTypes,omitempty"`
	Ipv4AddressRanges []NrfNfManagementIpv4AddressRange `json:"ipv4AddressRanges,omitempty" yaml:"ipv4AddressRanges" bson:"ipv4AddressRanges,omitempty"`
	Ipv6PrefixRanges  []NrfNfManagementIpv6PrefixRange  `json:"ipv6PrefixRanges,omitempty" yaml:"ipv6PrefixRanges" bson:"ipv6PrefixRanges,omitempty"`
	Ipv4IndexList     []UdmSdmIpIndex                   `json:"ipv4IndexList,omitempty" yaml:"ipv4IndexList" bson:"ipv4IndexList,omitempty"`
	Ipv6IndexList     []UdmSdmIpIndex                   `json:"ipv6IndexList,omitempty" yaml:"ipv6IndexList" bson:"ipv6IndexList,omitempty"`
	// Map of network instance per DNAI for the DNN, where the key of the map is the DNAI. When present, the value of each entry of the map shall contain a N6 network instance that is configured for the DNAI indicated by the key.
	DnaiNwInstanceList map[string]string `json:"dnaiNwInstanceList,omitempty" yaml:"dnaiNwInstanceList" bson:"dnaiNwInstanceList,omitempty"`
}