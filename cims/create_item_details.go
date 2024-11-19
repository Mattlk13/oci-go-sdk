// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
// **Note**: Before you can create service requests with this API,
// complete user registration at My Oracle Cloud Support
// and then ask your tenancy administrator to provide you authorization for the related user groups.
//

package cims

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateItemDetails Details gathered during ticket creation.
type CreateItemDetails interface {
	GetCategory() *CreateCategoryDetails

	GetSubCategory() *CreateSubCategoryDetails

	GetIssueType() *CreateIssueTypeDetails

	// The display name of the ticket. Avoid entering confidential information.
	GetName() *string
}

type createitemdetails struct {
	JsonData    []byte
	Category    *CreateCategoryDetails    `mandatory:"false" json:"category"`
	SubCategory *CreateSubCategoryDetails `mandatory:"false" json:"subCategory"`
	IssueType   *CreateIssueTypeDetails   `mandatory:"false" json:"issueType"`
	Name        *string                   `mandatory:"false" json:"name"`
	Type        string                    `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createitemdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateitemdetails createitemdetails
	s := struct {
		Model Unmarshalercreateitemdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Category = s.Model.Category
	m.SubCategory = s.Model.SubCategory
	m.IssueType = s.Model.IssueType
	m.Name = s.Model.Name
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createitemdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "tech":
		mm := CreateTechSupportItemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "limit":
		mm := CreateLimitItemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "account":
		mm := CreateAccountItemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateItemDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetCategory returns Category
func (m createitemdetails) GetCategory() *CreateCategoryDetails {
	return m.Category
}

// GetSubCategory returns SubCategory
func (m createitemdetails) GetSubCategory() *CreateSubCategoryDetails {
	return m.SubCategory
}

// GetIssueType returns IssueType
func (m createitemdetails) GetIssueType() *CreateIssueTypeDetails {
	return m.IssueType
}

// GetName returns Name
func (m createitemdetails) GetName() *string {
	return m.Name
}

func (m createitemdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createitemdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
