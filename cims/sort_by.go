// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
// **Note**: Before you can create service requests with this API,
// complete user registration at My Oracle Cloud Support
// and then ask your tenancy administrator to provide you authorization for the related user groups.
//

package cims

import (
	"strings"
)

// SortByEnum Enum with underlying type: string
type SortByEnum string

// Set of constants representing the allowable values for SortByEnum
const (
	SortByDateUpdated SortByEnum = "dateUpdated"
	SortBySeverity    SortByEnum = "severity"
)

var mappingSortByEnum = map[string]SortByEnum{
	"dateUpdated": SortByDateUpdated,
	"severity":    SortBySeverity,
}

var mappingSortByEnumLowerCase = map[string]SortByEnum{
	"dateupdated": SortByDateUpdated,
	"severity":    SortBySeverity,
}

// GetSortByEnumValues Enumerates the set of values for SortByEnum
func GetSortByEnumValues() []SortByEnum {
	values := make([]SortByEnum, 0)
	for _, v := range mappingSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSortByEnumStringValues Enumerates the set of values in String for SortByEnum
func GetSortByEnumStringValues() []string {
	return []string{
		"dateUpdated",
		"severity",
	}
}

// GetMappingSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortByEnum(val string) (SortByEnum, bool) {
	enum, ok := mappingSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
