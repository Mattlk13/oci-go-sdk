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

// TenancyAssociation Tenancy association represents the tenancy id to which the governance rule will be applied.
type TenancyAssociation struct {

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the tenancy to which the governance rule will be applied as part of this tenancy inclusion criterion.
	TenancyId *string `mandatory:"true" json:"tenancyId"`
}

func (m TenancyAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TenancyAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TenancyAssociation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTenancyAssociation TenancyAssociation
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTenancyAssociation
	}{
		"TENANCY",
		(MarshalTypeTenancyAssociation)(m),
	}

	return json.Marshal(&s)
}
