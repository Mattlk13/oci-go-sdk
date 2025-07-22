// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateLeadGenListingRevisionDetails Listing revision update details for listings
type UpdateLeadGenListingRevisionDetails struct {

	// The name for the listing revision.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Single line introduction for the listing revision.
	Headline *string `mandatory:"false" json:"headline"`

	// The tagline for the listing revision.
	Tagline *string `mandatory:"false" json:"tagline"`

	// Keywords associated for the listing revision.
	Keywords *string `mandatory:"false" json:"keywords"`

	// A short description for the listing revision.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// Usage information for the listing revision.
	UsageInformation *string `mandatory:"false" json:"usageInformation"`

	// A long description for the listing revision.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	ContentLanguage *LanguageItem `mandatory:"false" json:"contentLanguage"`

	// Languages supported by the listing revision.
	Supportedlanguages []LanguageItem `mandatory:"false" json:"supportedlanguages"`

	// Contact information to use to get support from the publisher for the listing revision.
	SupportContacts []SupportContact `mandatory:"false" json:"supportContacts"`

	// Links to support resources for the listing revision.
	SupportLinks []NamedLink `mandatory:"false" json:"supportLinks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	VersionDetails *VersionDetails `mandatory:"false" json:"versionDetails"`

	// System requirements for the listing revision.
	SystemRequirements *string `mandatory:"false" json:"systemRequirements"`

	// List of Products subscribed by listing.
	Products []ListingProduct `mandatory:"false" json:"products"`

	// Url to demo of the listing
	DemoUrl *string `mandatory:"false" json:"demoUrl"`

	// Url to training resources of the listing
	SelfPacedTrainingUrl *string `mandatory:"false" json:"selfPacedTrainingUrl"`

	// OCIDs of service listings attached to lead gen listing
	RecommendedServiceProviderListingIds []string `mandatory:"false" json:"recommendedServiceProviderListingIds"`

	// Custom link to the listing
	VanityUrl *string `mandatory:"false" json:"vanityUrl"`

	DownloadInfo *DownloadInfo `mandatory:"false" json:"downloadInfo"`

	// Pricing details for lead gen listing
	PricingPlans *string `mandatory:"false" json:"pricingPlans"`

	// The pricing model for the listing revision.
	PricingType OciListingRevisionPricingTypeEnum `mandatory:"false" json:"pricingType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m UpdateLeadGenListingRevisionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetHeadline returns Headline
func (m UpdateLeadGenListingRevisionDetails) GetHeadline() *string {
	return m.Headline
}

// GetTagline returns Tagline
func (m UpdateLeadGenListingRevisionDetails) GetTagline() *string {
	return m.Tagline
}

// GetKeywords returns Keywords
func (m UpdateLeadGenListingRevisionDetails) GetKeywords() *string {
	return m.Keywords
}

// GetShortDescription returns ShortDescription
func (m UpdateLeadGenListingRevisionDetails) GetShortDescription() *string {
	return m.ShortDescription
}

// GetUsageInformation returns UsageInformation
func (m UpdateLeadGenListingRevisionDetails) GetUsageInformation() *string {
	return m.UsageInformation
}

// GetLongDescription returns LongDescription
func (m UpdateLeadGenListingRevisionDetails) GetLongDescription() *string {
	return m.LongDescription
}

// GetContentLanguage returns ContentLanguage
func (m UpdateLeadGenListingRevisionDetails) GetContentLanguage() *LanguageItem {
	return m.ContentLanguage
}

// GetSupportedlanguages returns Supportedlanguages
func (m UpdateLeadGenListingRevisionDetails) GetSupportedlanguages() []LanguageItem {
	return m.Supportedlanguages
}

// GetSupportContacts returns SupportContacts
func (m UpdateLeadGenListingRevisionDetails) GetSupportContacts() []SupportContact {
	return m.SupportContacts
}

// GetSupportLinks returns SupportLinks
func (m UpdateLeadGenListingRevisionDetails) GetSupportLinks() []NamedLink {
	return m.SupportLinks
}

// GetFreeformTags returns FreeformTags
func (m UpdateLeadGenListingRevisionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateLeadGenListingRevisionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateLeadGenListingRevisionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLeadGenListingRevisionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOciListingRevisionPricingTypeEnum(string(m.PricingType)); !ok && m.PricingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingType: %s. Supported values are: %s.", m.PricingType, strings.Join(GetOciListingRevisionPricingTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateLeadGenListingRevisionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateLeadGenListingRevisionDetails UpdateLeadGenListingRevisionDetails
	s := struct {
		DiscriminatorParam string `json:"listingType"`
		MarshalTypeUpdateLeadGenListingRevisionDetails
	}{
		"LEAD_GENERATION",
		(MarshalTypeUpdateLeadGenListingRevisionDetails)(m),
	}

	return json.Marshal(&s)
}
