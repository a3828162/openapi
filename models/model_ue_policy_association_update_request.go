/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.525 V16.9.0; 5G System; UE Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.525/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UePolicyAssociationUpdateRequest struct {
	NotificationUri string `json:"notificationUri,omitempty" yaml:"notificationUri" bson:"notificationUri,omitempty"`
	// Alternate or backup IPv4 Address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty" yaml:"altNotifIpv4Addrs" bson:"altNotifIpv4Addrs,omitempty"`
	// Alternate or backup IPv6 Address(es) where to send Notifications.
	AltNotifIpv6Addrs []string `json:"altNotifIpv6Addrs,omitempty" yaml:"altNotifIpv6Addrs" bson:"altNotifIpv6Addrs,omitempty"`
	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty" yaml:"altNotifFqdns" bson:"altNotifFqdns,omitempty"`
	// Request Triggers that the NF service consumer observes.
	Triggers []UePolicyRequestTrigger `json:"triggers,omitempty" yaml:"triggers" bson:"triggers,omitempty"`
	// Map of PRA status information.
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty" yaml:"praStatuses" bson:"praStatuses,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty" yaml:"userLoc" bson:"userLoc,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty" yaml:"uePolDelResult" bson:"uePolDelResult,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty" yaml:"uePolTransFailNotif" bson:"uePolTransFailNotif,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty" yaml:"uePolReq" bson:"uePolReq,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty" yaml:"guami" bson:"guami,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty" yaml:"servingNfId" bson:"servingNfId,omitempty"`
	PlmnId              *PlmnId                              `json:"plmnId,omitempty" yaml:"plmnId" bson:"plmnId,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty" yaml:"connectState" bson:"connectState,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty" yaml:"groupIds" bson:"groupIds,omitempty"`
}