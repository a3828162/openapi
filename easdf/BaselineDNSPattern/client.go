/*
 * Neasdf_BaselineDNSPattern
 *
 * EASDF Baseline DNS Pattern Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.556 V17.3.0; 5G System; Edge Application Server Discovery Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.556/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package BaselineDNSPattern

// APIClient manages communication with the Neasdf_BaselineDNSPattern API v1.0.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
    cfg    *Configuration
    common service // Reuse a single struct instead of allocating one for each service on the heap.

    // API Services
    	IndividualBaselineDNSPatternApi *IndividualBaselineDNSPatternApiService
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
    c.IndividualBaselineDNSPatternApi = (*IndividualBaselineDNSPatternApiService)(&c.common)

    return c
}