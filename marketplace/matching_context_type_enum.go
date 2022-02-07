// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

// MatchingContextTypeEnumEnum Enum with underlying type: string
type MatchingContextTypeEnumEnum string

// Set of constants representing the allowable values for MatchingContextTypeEnumEnum
const (
	MatchingContextTypeEnumNone       MatchingContextTypeEnumEnum = "NONE"
	MatchingContextTypeEnumHighlights MatchingContextTypeEnumEnum = "HIGHLIGHTS"
)

var mappingMatchingContextTypeEnumEnum = map[string]MatchingContextTypeEnumEnum{
	"NONE":       MatchingContextTypeEnumNone,
	"HIGHLIGHTS": MatchingContextTypeEnumHighlights,
}

// GetMatchingContextTypeEnumEnumValues Enumerates the set of values for MatchingContextTypeEnumEnum
func GetMatchingContextTypeEnumEnumValues() []MatchingContextTypeEnumEnum {
	values := make([]MatchingContextTypeEnumEnum, 0)
	for _, v := range mappingMatchingContextTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetMatchingContextTypeEnumEnumStringValues Enumerates the set of values in String for MatchingContextTypeEnumEnum
func GetMatchingContextTypeEnumEnumStringValues() []string {
	return []string{
		"NONE",
		"HIGHLIGHTS",
	}
}
