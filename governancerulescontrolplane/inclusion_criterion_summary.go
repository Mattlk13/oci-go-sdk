// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InclusionCriterionSummary Summary of the inclusion criterion.
type InclusionCriterionSummary struct {

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the inclusion criterion.
	Id *string `mandatory:"true" json:"id"`

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the governance rule. Every inclusion criterion is associated with a governance rule.
	GovernanceRuleId *string `mandatory:"true" json:"governanceRuleId"`

	// Type of inclusion criterion - TENANCY, ALL or TAG. We support TENANCY and ALL for now.
	Type InclusionCriterionTypeEnum `mandatory:"true" json:"type"`

	// The current state of the inclusion criterion.
	LifecycleState InclusionCriterionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Date and time the inclusion criterion was created. An RFC3339 formatted datetime string.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the inclusion criterion was updated. An RFC3339 formatted datetime string.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	Association Association `mandatory:"false" json:"association"`
}

func (m InclusionCriterionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InclusionCriterionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInclusionCriterionTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetInclusionCriterionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInclusionCriterionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInclusionCriterionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InclusionCriterionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Association      association                          `json:"association"`
		Id               *string                              `json:"id"`
		GovernanceRuleId *string                              `json:"governanceRuleId"`
		Type             InclusionCriterionTypeEnum           `json:"type"`
		LifecycleState   InclusionCriterionLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated      *common.SDKTime                      `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                      `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Association.UnmarshalPolymorphicJSON(model.Association.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Association = nn.(Association)
	} else {
		m.Association = nil
	}

	m.Id = model.Id

	m.GovernanceRuleId = model.GovernanceRuleId

	m.Type = model.Type

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}
