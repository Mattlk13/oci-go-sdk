// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// AutoScalingConfiguration The information about the autoscale configuration.
type AutoScalingConfiguration struct {

	// The unique identifier for the autoscale configuration.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A node type that is managed by an autoscale configuration. The only supported type is WORKER.
	NodeType NodeNodeTypeEnum `mandatory:"true" json:"nodeType"`

	// The state of the autoscale configuration.
	LifecycleState AutoScalingConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the autoscale configuration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	Policy *AutoScalePolicy `mandatory:"true" json:"policy"`
}

func (m AutoScalingConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoScalingConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingNodeNodeTypeEnum[string(m.NodeType)]; !ok && m.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", m.NodeType, strings.Join(GetNodeNodeTypeEnumStringValues(), ",")))
	}
	if _, ok := mappingAutoScalingConfigurationLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutoScalingConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoScalingConfigurationLifecycleStateEnum Enum with underlying type: string
type AutoScalingConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for AutoScalingConfigurationLifecycleStateEnum
const (
	AutoScalingConfigurationLifecycleStateCreating AutoScalingConfigurationLifecycleStateEnum = "CREATING"
	AutoScalingConfigurationLifecycleStateActive   AutoScalingConfigurationLifecycleStateEnum = "ACTIVE"
	AutoScalingConfigurationLifecycleStateUpdating AutoScalingConfigurationLifecycleStateEnum = "UPDATING"
	AutoScalingConfigurationLifecycleStateDeleting AutoScalingConfigurationLifecycleStateEnum = "DELETING"
	AutoScalingConfigurationLifecycleStateDeleted  AutoScalingConfigurationLifecycleStateEnum = "DELETED"
	AutoScalingConfigurationLifecycleStateFailed   AutoScalingConfigurationLifecycleStateEnum = "FAILED"
)

var mappingAutoScalingConfigurationLifecycleStateEnum = map[string]AutoScalingConfigurationLifecycleStateEnum{
	"CREATING": AutoScalingConfigurationLifecycleStateCreating,
	"ACTIVE":   AutoScalingConfigurationLifecycleStateActive,
	"UPDATING": AutoScalingConfigurationLifecycleStateUpdating,
	"DELETING": AutoScalingConfigurationLifecycleStateDeleting,
	"DELETED":  AutoScalingConfigurationLifecycleStateDeleted,
	"FAILED":   AutoScalingConfigurationLifecycleStateFailed,
}

// GetAutoScalingConfigurationLifecycleStateEnumValues Enumerates the set of values for AutoScalingConfigurationLifecycleStateEnum
func GetAutoScalingConfigurationLifecycleStateEnumValues() []AutoScalingConfigurationLifecycleStateEnum {
	values := make([]AutoScalingConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingAutoScalingConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoScalingConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for AutoScalingConfigurationLifecycleStateEnum
func GetAutoScalingConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}
