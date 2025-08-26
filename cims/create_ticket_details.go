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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateTicketDetails Details relevant to the support request.
type CreateTicketDetails struct {

	// The severity of the support request.
	Severity CreateTicketDetailsSeverityEnum `mandatory:"true" json:"severity"`

	// The title of the support request. Avoid entering confidential information.
	Title *string `mandatory:"true" json:"title"`

	// <b>Important</b>: On January 27, 2026, the <b>Max Length</b> value will change to 1500.
	// The description of the support request. Avoid entering confidential information.
	Description *string `mandatory:"true" json:"description"`

	// The list of resources.
	ResourceList []CreateResourceDetails `mandatory:"false" json:"resourceList"`

	ContextualData *ContextualData `mandatory:"false" json:"contextualData"`
}

func (m CreateTicketDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTicketDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateTicketDetailsSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetCreateTicketDetailsSeverityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateTicketDetailsSeverityEnum Enum with underlying type: string
type CreateTicketDetailsSeverityEnum string

// Set of constants representing the allowable values for CreateTicketDetailsSeverityEnum
const (
	CreateTicketDetailsSeverityHighest CreateTicketDetailsSeverityEnum = "HIGHEST"
	CreateTicketDetailsSeverityHigh    CreateTicketDetailsSeverityEnum = "HIGH"
	CreateTicketDetailsSeverityMedium  CreateTicketDetailsSeverityEnum = "MEDIUM"
	CreateTicketDetailsSeverityLow     CreateTicketDetailsSeverityEnum = "LOW"
)

var mappingCreateTicketDetailsSeverityEnum = map[string]CreateTicketDetailsSeverityEnum{
	"HIGHEST": CreateTicketDetailsSeverityHighest,
	"HIGH":    CreateTicketDetailsSeverityHigh,
	"MEDIUM":  CreateTicketDetailsSeverityMedium,
	"LOW":     CreateTicketDetailsSeverityLow,
}

var mappingCreateTicketDetailsSeverityEnumLowerCase = map[string]CreateTicketDetailsSeverityEnum{
	"highest": CreateTicketDetailsSeverityHighest,
	"high":    CreateTicketDetailsSeverityHigh,
	"medium":  CreateTicketDetailsSeverityMedium,
	"low":     CreateTicketDetailsSeverityLow,
}

// GetCreateTicketDetailsSeverityEnumValues Enumerates the set of values for CreateTicketDetailsSeverityEnum
func GetCreateTicketDetailsSeverityEnumValues() []CreateTicketDetailsSeverityEnum {
	values := make([]CreateTicketDetailsSeverityEnum, 0)
	for _, v := range mappingCreateTicketDetailsSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTicketDetailsSeverityEnumStringValues Enumerates the set of values in String for CreateTicketDetailsSeverityEnum
func GetCreateTicketDetailsSeverityEnumStringValues() []string {
	return []string{
		"HIGHEST",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingCreateTicketDetailsSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTicketDetailsSeverityEnum(val string) (CreateTicketDetailsSeverityEnum, bool) {
	enum, ok := mappingCreateTicketDetailsSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
