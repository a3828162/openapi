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

// DNS message handling rule
type DnsRule struct {
	DnsRuleId string `json:"dnsRuleId,omitempty" yaml:"dnsRuleId" bson:"dnsRuleId,omitempty"`
	Label     string `json:"label,omitempty" yaml:"label" bson:"label,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	Precedence int32 `json:"precedence,omitempty" yaml:"precedence" bson:"precedence,omitempty"`
	// map of DNS query message detection templates where a valid JSON string serves as key
	DnsQueryMdtList     map[string]DnsQueryMdt    `json:"dnsQueryMdtList,omitempty" yaml:"dnsQueryMdtList" bson:"dnsQueryMdtList,omitempty"`
	BaseDnsQueryMdtList []BaselineDnsQueryMdtInfo `json:"baseDnsQueryMdtList,omitempty" yaml:"baseDnsQueryMdtList" bson:"baseDnsQueryMdtList,omitempty"`
	// map of DNS response message detection templates where a valid JSON string serves as key
	DnsRspMdtList     map[string]DnsRspMdt    `json:"dnsRspMdtList,omitempty" yaml:"dnsRspMdtList" bson:"dnsRspMdtList,omitempty"`
	BaseDnsRspMdtList []BaselineDnsRspMdtInfo `json:"baseDnsRspMdtList,omitempty" yaml:"baseDnsRspMdtList" bson:"baseDnsRspMdtList,omitempty"`
	DnsMsgId          string                  `json:"dnsMsgId,omitempty" yaml:"dnsMsgId" bson:"dnsMsgId,omitempty"`
	// map of actions where a valid JSON string serves as key
	ActionList map[string]Action `json:"actionList" yaml:"actionList" bson:"actionList,omitempty"`
}
