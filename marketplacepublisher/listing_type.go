// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"strings"
)

// ListingTypeEnum Enum with underlying type: string
type ListingTypeEnum string

// Set of constants representing the allowable values for ListingTypeEnum
const (
	ListingTypeOciApplication ListingTypeEnum = "OCI_APPLICATION"
	ListingTypeLeadGeneration ListingTypeEnum = "LEAD_GENERATION"
	ListingTypeService        ListingTypeEnum = "SERVICE"
)

var mappingListingTypeEnum = map[string]ListingTypeEnum{
	"OCI_APPLICATION": ListingTypeOciApplication,
	"LEAD_GENERATION": ListingTypeLeadGeneration,
	"SERVICE":         ListingTypeService,
}

var mappingListingTypeEnumLowerCase = map[string]ListingTypeEnum{
	"oci_application": ListingTypeOciApplication,
	"lead_generation": ListingTypeLeadGeneration,
	"service":         ListingTypeService,
}

// GetListingTypeEnumValues Enumerates the set of values for ListingTypeEnum
func GetListingTypeEnumValues() []ListingTypeEnum {
	values := make([]ListingTypeEnum, 0)
	for _, v := range mappingListingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListingTypeEnumStringValues Enumerates the set of values in String for ListingTypeEnum
func GetListingTypeEnumStringValues() []string {
	return []string{
		"OCI_APPLICATION",
		"LEAD_GENERATION",
		"SERVICE",
	}
}

// GetMappingListingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingTypeEnum(val string) (ListingTypeEnum, bool) {
	enum, ok := mappingListingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
