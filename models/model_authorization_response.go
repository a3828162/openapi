/*
 * Nhss_imsUECM
 *
 * Nhss UE Context Management Service for IMS. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * Source file: 3GPP TS 29.562 HSS Services for interworking with IMS, version 16.5.0
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models



// Ims Registration authorization information result
type AuthorizationResponse struct {
	AuthorizationResult AuthorizationResult `json:"authorizationResult" yaml:"authorizationResult" bson:"authorizationResult"`
	CscfServerName string `json:"cscfServerName,omitempty" yaml:"cscfServerName" bson:"cscfServerName"`
	ScscfSelectionAssistanceInfo *ScscfSelectionAssistanceInformation `json:"scscfSelectionAssistanceInfo,omitempty" yaml:"scscfSelectionAssistanceInfo" bson:"scscfSelectionAssistanceInfo"`
}