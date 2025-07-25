// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudClusterInstance The details of a cloud cluster instance.
type CloudClusterInstance struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud cluster instance.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the cluster instance. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the cloud cluster instance.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud cluster that the cluster instance belongs to.
	CloudClusterId *string `mandatory:"true" json:"cloudClusterId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the cluster instance is a part of.
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`

	// The current lifecycle state of the cloud cluster instance.
	LifecycleState CloudClusterInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node.
	CloudDbNodeId *string `mandatory:"false" json:"cloudDbNodeId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
	CloudConnectorId *string `mandatory:"false" json:"cloudConnectorId"`

	// The name of the host on which the cluster instance is running.
	HostName *string `mandatory:"false" json:"hostName"`

	// The role of the cluster node.
	NodeRole CloudClusterInstanceNodeRoleEnum `mandatory:"false" json:"nodeRole,omitempty"`

	// The Oracle base location of Cluster Ready Services (CRS).
	CrsBaseDirectory *string `mandatory:"false" json:"crsBaseDirectory"`

	// The Automatic Diagnostic Repository (ADR) home directory for the cluster instance.
	AdrHomeDirectory *string `mandatory:"false" json:"adrHomeDirectory"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the cloud cluster instance was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the cloud cluster instance was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CloudClusterInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudClusterInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudClusterInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudClusterInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCloudClusterInstanceNodeRoleEnum(string(m.NodeRole)); !ok && m.NodeRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeRole: %s. Supported values are: %s.", m.NodeRole, strings.Join(GetCloudClusterInstanceNodeRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudClusterInstanceNodeRoleEnum Enum with underlying type: string
type CloudClusterInstanceNodeRoleEnum string

// Set of constants representing the allowable values for CloudClusterInstanceNodeRoleEnum
const (
	CloudClusterInstanceNodeRoleHub  CloudClusterInstanceNodeRoleEnum = "HUB"
	CloudClusterInstanceNodeRoleLeaf CloudClusterInstanceNodeRoleEnum = "LEAF"
)

var mappingCloudClusterInstanceNodeRoleEnum = map[string]CloudClusterInstanceNodeRoleEnum{
	"HUB":  CloudClusterInstanceNodeRoleHub,
	"LEAF": CloudClusterInstanceNodeRoleLeaf,
}

var mappingCloudClusterInstanceNodeRoleEnumLowerCase = map[string]CloudClusterInstanceNodeRoleEnum{
	"hub":  CloudClusterInstanceNodeRoleHub,
	"leaf": CloudClusterInstanceNodeRoleLeaf,
}

// GetCloudClusterInstanceNodeRoleEnumValues Enumerates the set of values for CloudClusterInstanceNodeRoleEnum
func GetCloudClusterInstanceNodeRoleEnumValues() []CloudClusterInstanceNodeRoleEnum {
	values := make([]CloudClusterInstanceNodeRoleEnum, 0)
	for _, v := range mappingCloudClusterInstanceNodeRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudClusterInstanceNodeRoleEnumStringValues Enumerates the set of values in String for CloudClusterInstanceNodeRoleEnum
func GetCloudClusterInstanceNodeRoleEnumStringValues() []string {
	return []string{
		"HUB",
		"LEAF",
	}
}

// GetMappingCloudClusterInstanceNodeRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudClusterInstanceNodeRoleEnum(val string) (CloudClusterInstanceNodeRoleEnum, bool) {
	enum, ok := mappingCloudClusterInstanceNodeRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudClusterInstanceLifecycleStateEnum Enum with underlying type: string
type CloudClusterInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for CloudClusterInstanceLifecycleStateEnum
const (
	CloudClusterInstanceLifecycleStateCreating     CloudClusterInstanceLifecycleStateEnum = "CREATING"
	CloudClusterInstanceLifecycleStateNotConnected CloudClusterInstanceLifecycleStateEnum = "NOT_CONNECTED"
	CloudClusterInstanceLifecycleStateActive       CloudClusterInstanceLifecycleStateEnum = "ACTIVE"
	CloudClusterInstanceLifecycleStateInactive     CloudClusterInstanceLifecycleStateEnum = "INACTIVE"
	CloudClusterInstanceLifecycleStateUpdating     CloudClusterInstanceLifecycleStateEnum = "UPDATING"
	CloudClusterInstanceLifecycleStateDeleting     CloudClusterInstanceLifecycleStateEnum = "DELETING"
	CloudClusterInstanceLifecycleStateDeleted      CloudClusterInstanceLifecycleStateEnum = "DELETED"
	CloudClusterInstanceLifecycleStateFailed       CloudClusterInstanceLifecycleStateEnum = "FAILED"
)

var mappingCloudClusterInstanceLifecycleStateEnum = map[string]CloudClusterInstanceLifecycleStateEnum{
	"CREATING":      CloudClusterInstanceLifecycleStateCreating,
	"NOT_CONNECTED": CloudClusterInstanceLifecycleStateNotConnected,
	"ACTIVE":        CloudClusterInstanceLifecycleStateActive,
	"INACTIVE":      CloudClusterInstanceLifecycleStateInactive,
	"UPDATING":      CloudClusterInstanceLifecycleStateUpdating,
	"DELETING":      CloudClusterInstanceLifecycleStateDeleting,
	"DELETED":       CloudClusterInstanceLifecycleStateDeleted,
	"FAILED":        CloudClusterInstanceLifecycleStateFailed,
}

var mappingCloudClusterInstanceLifecycleStateEnumLowerCase = map[string]CloudClusterInstanceLifecycleStateEnum{
	"creating":      CloudClusterInstanceLifecycleStateCreating,
	"not_connected": CloudClusterInstanceLifecycleStateNotConnected,
	"active":        CloudClusterInstanceLifecycleStateActive,
	"inactive":      CloudClusterInstanceLifecycleStateInactive,
	"updating":      CloudClusterInstanceLifecycleStateUpdating,
	"deleting":      CloudClusterInstanceLifecycleStateDeleting,
	"deleted":       CloudClusterInstanceLifecycleStateDeleted,
	"failed":        CloudClusterInstanceLifecycleStateFailed,
}

// GetCloudClusterInstanceLifecycleStateEnumValues Enumerates the set of values for CloudClusterInstanceLifecycleStateEnum
func GetCloudClusterInstanceLifecycleStateEnumValues() []CloudClusterInstanceLifecycleStateEnum {
	values := make([]CloudClusterInstanceLifecycleStateEnum, 0)
	for _, v := range mappingCloudClusterInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudClusterInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for CloudClusterInstanceLifecycleStateEnum
func GetCloudClusterInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NOT_CONNECTED",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCloudClusterInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudClusterInstanceLifecycleStateEnum(val string) (CloudClusterInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingCloudClusterInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
