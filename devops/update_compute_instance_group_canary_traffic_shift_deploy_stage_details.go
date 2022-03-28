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
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails Specifies load balancer traffic shift stage.
type UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails struct {

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	RolloutPolicy *LoadBalancerTrafficShiftRolloutPolicy `mandatory:"false" json:"rolloutPolicy"`
}

//GetDescription returns Description
func (m UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeUpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails
	}{
		"COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT",
		(MarshalTypeUpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}
