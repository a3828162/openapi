/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291  V17.0.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contain the TNAP Identifier see clause5.6.2 of 3GPP TS 23.501.
type TnapId struct {
	// This IE shall be present if the UE is accessing the 5GC via a trusted WLAN access network.When present, it shall contain the SSID of the access point to which the UE is attached, that is received over NGAP,  see IEEE Std 802.11-2012.
	SsId string `json:"ssId,omitempty" yaml:"ssId" bson:"ssId,omitempty"`
	// When present, it shall contain the BSSID of the access point to which the UE is attached, that is received over NGAP, see IEEE Std 802.11-2012.
	BssId string `json:"bssId,omitempty" yaml:"bssId" bson:"bssId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	CivicAddress string `json:"civicAddress,omitempty" yaml:"civicAddress" bson:"civicAddress,omitempty"`
}