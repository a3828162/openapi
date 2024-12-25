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

// Range of IPv6 prefixes
type Ipv6PrefixRange struct {
	Start string `json:"start" yaml:"start" bson:"start,omitempty"`
	End   string `json:"end" yaml:"end" bson:"end,omitempty"`
}
