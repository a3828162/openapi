/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Data within Create SM Context Response
type SmfPduSessionSmContextCreatedData struct {
	// String providing an URI formatted according to RFC 3986.
	HSmfUri string `json:"hSmfUri,omitempty" yaml:"hSmfUri" bson:"hSmfUri,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SmfUri string `json:"smfUri,omitempty" yaml:"smfUri" bson:"smfUri,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	PduSessionId     int32            `json:"pduSessionId,omitempty" yaml:"pduSessionId" bson:"pduSessionId,omitempty"`
	SNssai           *Snssai          `json:"sNssai,omitempty" yaml:"sNssai" bson:"sNssai,omitempty"`
	AdditionalSnssai *Snssai          `json:"additionalSnssai,omitempty" yaml:"additionalSnssai" bson:"additionalSnssai,omitempty"`
	UpCnxState       UpCnxState       `json:"upCnxState,omitempty" yaml:"upCnxState" bson:"upCnxState,omitempty"`
	N2SmInfo         *RefToBinaryData `json:"n2SmInfo,omitempty" yaml:"n2SmInfo" bson:"n2SmInfo,omitempty"`
	N2SmInfoType     N2SmInfoType     `json:"n2SmInfoType,omitempty" yaml:"n2SmInfoType" bson:"n2SmInfoType,omitempty"`
	AllocatedEbiList []EbiArpMapping  `json:"allocatedEbiList,omitempty" yaml:"allocatedEbiList" bson:"allocatedEbiList,omitempty"`
	HoState          HoState          `json:"hoState,omitempty" yaml:"hoState" bson:"hoState,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi                 string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	SmfServiceInstanceId string `json:"smfServiceInstanceId,omitempty" yaml:"smfServiceInstanceId" bson:"smfServiceInstanceId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time `json:"recoveryTime,omitempty" yaml:"recoveryTime" bson:"recoveryTime,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SelectedSmfId string `json:"selectedSmfId,omitempty" yaml:"selectedSmfId" bson:"selectedSmfId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SelectedOldSmfId string `json:"selectedOldSmfId,omitempty" yaml:"selectedOldSmfId" bson:"selectedOldSmfId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	InterPlmnApiRoot string `json:"interPlmnApiRoot,omitempty" yaml:"interPlmnApiRoot" bson:"interPlmnApiRoot,omitempty"`
}