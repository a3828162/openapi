/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.10.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains information about subscriptions that are requested to be transferred.
type SubscriptionTransferInfo struct {
	TransReqType TransferRequestType       `json:"transReqType" yaml:"transReqType" bson:"transReqType,omitempty"`
	NwdafEvSub   *NnwdafEventsSubscription `json:"nwdafEvSub" yaml:"nwdafEvSub" bson:"nwdafEvSub,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	ConsumerId   string                      `json:"consumerId" yaml:"consumerId" bson:"consumerId,omitempty"`
	ContextId    *AnalyticsContextIdentifier `json:"contextId,omitempty" yaml:"contextId" bson:"contextId,omitempty"`
	SourceNfIds  []string                    `json:"sourceNfIds,omitempty" yaml:"sourceNfIds" bson:"sourceNfIds,omitempty"`
	SourceSetIds []string                    `json:"sourceSetIds,omitempty" yaml:"sourceSetIds" bson:"sourceSetIds,omitempty"`
	ModelInfo    []ModelInfo                 `json:"modelInfo,omitempty" yaml:"modelInfo" bson:"modelInfo,omitempty"`
}