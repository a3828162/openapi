/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V17.7.0; 5G System; Network Function Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// The claims data structure for the access token
type NrfAccessTokenAccessTokenClaims struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	Iss string `json:"iss" yaml:"iss" bson:"iss,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	Sub                string      `json:"sub" yaml:"sub" bson:"sub,omitempty"`
	Aud                interface{} `json:"aud" yaml:"aud" bson:"aud,omitempty"`
	Scope              string      `json:"scope" yaml:"scope" bson:"scope,omitempty"`
	Exp                int32       `json:"exp" yaml:"exp" bson:"exp,omitempty"`
	ConsumerPlmnId     *PlmnId     `json:"consumerPlmnId,omitempty" yaml:"consumerPlmnId" bson:"consumerPlmnId,omitempty"`
	ConsumerSnpnId     *PlmnIdNid  `json:"consumerSnpnId,omitempty" yaml:"consumerSnpnId" bson:"consumerSnpnId,omitempty"`
	ProducerPlmnId     *PlmnId     `json:"producerPlmnId,omitempty" yaml:"producerPlmnId" bson:"producerPlmnId,omitempty"`
	ProducerSnpnId     *PlmnIdNid  `json:"producerSnpnId,omitempty" yaml:"producerSnpnId" bson:"producerSnpnId,omitempty"`
	ProducerSnssaiList []Snssai    `json:"producerSnssaiList,omitempty" yaml:"producerSnssaiList" bson:"producerSnssaiList,omitempty"`
	ProducerNsiList    []string    `json:"producerNsiList,omitempty" yaml:"producerNsiList" bson:"producerNsiList,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	ProducerNfSetId string `json:"producerNfSetId,omitempty" yaml:"producerNfSetId" bson:"producerNfSetId,omitempty"`
	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.
	ProducerNfServiceSetId string `json:"producerNfServiceSetId,omitempty" yaml:"producerNfServiceSetId" bson:"producerNfServiceSetId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SourceNfInstanceId string `json:"sourceNfInstanceId,omitempty" yaml:"sourceNfInstanceId" bson:"sourceNfInstanceId,omitempty"`
}