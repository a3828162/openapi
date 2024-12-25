/*
 * Neasdf_DNSContext
 *
 * EASDF DNS Context Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.556 V17.3.0; 5G System; Edge Application Server Discovery Services; Stage3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.556/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Data within Create response
type DnsContextCreatedData struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	EasdfIpv4Addr string `json:"easdfIpv4Addr,omitempty" yaml:"easdfIpv4Addr" bson:"easdfIpv4Addr,omitempty"`
	EasdfIpv6Addr string `json:"easdfIpv6Addr,omitempty" yaml:"easdfIpv6Addr" bson:"easdfIpv6Addr,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
}
