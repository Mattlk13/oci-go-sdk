// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// CreateAbsoluteWaitCriteriaDetails Specifies the absolute wait criteria. You can specify fixed length of wait duration.
type CreateAbsoluteWaitCriteriaDetails struct {

	// The absolute wait duration.
	// Minimum wait duration must be 5 seconds.
	// Maximum wait duration can be up to 2 days.
	WaitDuration *string `mandatory:"true" json:"waitDuration"`
}

func (m CreateAbsoluteWaitCriteriaDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAbsoluteWaitCriteriaDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAbsoluteWaitCriteriaDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAbsoluteWaitCriteriaDetails CreateAbsoluteWaitCriteriaDetails
	s := struct {
		DiscriminatorParam string `json:"waitType"`
		MarshalTypeCreateAbsoluteWaitCriteriaDetails
	}{
		"ABSOLUTE_WAIT",
		(MarshalTypeCreateAbsoluteWaitCriteriaDetails)(m),
	}

	return json.Marshal(&s)
}
