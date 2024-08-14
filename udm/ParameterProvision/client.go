/*
 * Nudm_PP
 *
 * Nudm Parameter Provision Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.11.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ParameterProvision

// APIClient manages communication with the Nudm_PP API v1.2.2
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	Class5GMBSGroupCreationApi                *Class5GMBSGroupCreationApiService
	Class5GMBSGroupDeletionApi                *Class5GMBSGroupDeletionApiService
	Class5GMBSGroupModificationApi            *Class5GMBSGroupModificationApiService
	Class5GVNGroupCreationApi                 *Class5GVNGroupCreationApiService
	Class5GVNGroupDeletionApi                 *Class5GVNGroupDeletionApiService
	Class5GVNGroupModificationApi             *Class5GVNGroupModificationApiService
	ParameterProvisioningDataEntryDocumentApi *ParameterProvisioningDataEntryDocumentApiService
	SubscriptionDataUpdateApi                 *SubscriptionDataUpdateApiService
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
	c.Class5GMBSGroupCreationApi = (*Class5GMBSGroupCreationApiService)(&c.common)
	c.Class5GMBSGroupDeletionApi = (*Class5GMBSGroupDeletionApiService)(&c.common)
	c.Class5GMBSGroupModificationApi = (*Class5GMBSGroupModificationApiService)(&c.common)
	c.Class5GVNGroupCreationApi = (*Class5GVNGroupCreationApiService)(&c.common)
	c.Class5GVNGroupDeletionApi = (*Class5GVNGroupDeletionApiService)(&c.common)
	c.Class5GVNGroupModificationApi = (*Class5GVNGroupModificationApiService)(&c.common)
	c.ParameterProvisioningDataEntryDocumentApi = (*ParameterProvisioningDataEntryDocumentApiService)(&c.common)
	c.SubscriptionDataUpdateApi = (*SubscriptionDataUpdateApiService)(&c.common)

	return c
}