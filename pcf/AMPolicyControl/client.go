/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.507 V17.9.0; 5G System; Access and Mobility Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.507/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AMPolicyControl

// APIClient manages communication with the Npcf_AMPolicyControl API v1.2.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AMPolicyAssociationsCollectionApi        *AMPolicyAssociationsCollectionApiService
	IndividualAMPolicyAssociationDocumentApi *IndividualAMPolicyAssociationDocumentApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AMPolicyAssociationsCollectionApi = (*AMPolicyAssociationsCollectionApiService)(&c.common)
	c.IndividualAMPolicyAssociationDocumentApi = (*IndividualAMPolicyAssociationDocumentApiService)(&c.common)

	return c
}