// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Access Governance APIs
//
// Use the Oracle Access Governance API to create, view, and manage GovernanceInstances.
//

package accessgovernancecp

import (
	"strings"
)

// LicenseTypeEnum Enum with underlying type: string
type LicenseTypeEnum string

// Set of constants representing the allowable values for LicenseTypeEnum
const (
	LicenseTypeNewLicense          LicenseTypeEnum = "NEW_LICENSE"
	LicenseTypeBringYourOwnLicense LicenseTypeEnum = "BRING_YOUR_OWN_LICENSE"
	LicenseTypeAgOracleWorkloads   LicenseTypeEnum = "AG_ORACLE_WORKLOADS"
	LicenseTypeAgOci               LicenseTypeEnum = "AG_OCI"
)

var mappingLicenseTypeEnum = map[string]LicenseTypeEnum{
	"NEW_LICENSE":            LicenseTypeNewLicense,
	"BRING_YOUR_OWN_LICENSE": LicenseTypeBringYourOwnLicense,
	"AG_ORACLE_WORKLOADS":    LicenseTypeAgOracleWorkloads,
	"AG_OCI":                 LicenseTypeAgOci,
}

var mappingLicenseTypeEnumLowerCase = map[string]LicenseTypeEnum{
	"new_license":            LicenseTypeNewLicense,
	"bring_your_own_license": LicenseTypeBringYourOwnLicense,
	"ag_oracle_workloads":    LicenseTypeAgOracleWorkloads,
	"ag_oci":                 LicenseTypeAgOci,
}

// GetLicenseTypeEnumValues Enumerates the set of values for LicenseTypeEnum
func GetLicenseTypeEnumValues() []LicenseTypeEnum {
	values := make([]LicenseTypeEnum, 0)
	for _, v := range mappingLicenseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLicenseTypeEnumStringValues Enumerates the set of values in String for LicenseTypeEnum
func GetLicenseTypeEnumStringValues() []string {
	return []string{
		"NEW_LICENSE",
		"BRING_YOUR_OWN_LICENSE",
		"AG_ORACLE_WORKLOADS",
		"AG_OCI",
	}
}

// GetMappingLicenseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLicenseTypeEnum(val string) (LicenseTypeEnum, bool) {
	enum, ok := mappingLicenseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
