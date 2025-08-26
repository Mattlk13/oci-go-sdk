// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateRoverClusterDetails The information required to update a RoverCluster.
type UpdateRoverClusterDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Number of nodes desired in the cluster, in standalone clusters, between 5 and 15 inclusive. In station clusters, between 15 and 30 inclusive.
	ClusterSize *int `mandatory:"false" json:"clusterSize"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// List of existing workloads that should be provisioned on the nodes.
	ClusterWorkloads []RoverWorkload `mandatory:"false" json:"clusterWorkloads"`

	// Root password for the rover cluster.
	SuperUserPassword *string `mandatory:"false" json:"superUserPassword"`

	// The current state of the RoverCluster.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Password to unlock the rover cluster.
	UnlockPassphrase *string `mandatory:"false" json:"unlockPassphrase"`

	// The type of enclosure rover nodes in this cluster are shipped in.
	EnclosureType EnclosureTypeEnum `mandatory:"false" json:"enclosureType,omitempty"`

	// Name of point of contact for this order if customer is picking up.
	PointOfContact *string `mandatory:"false" json:"pointOfContact"`

	// Phone number of point of contact for this order if customer is picking up.
	PointOfContactPhoneNumber *string `mandatory:"false" json:"pointOfContactPhoneNumber"`

	// Preference for device delivery.
	ShippingPreference UpdateRoverClusterDetailsShippingPreferenceEnum `mandatory:"false" json:"shippingPreference,omitempty"`

	// Tracking Url for the shipped Rover Cluster.
	OracleShippingTrackingUrl *string `mandatory:"false" json:"oracleShippingTrackingUrl"`

	// ID provided to customer after successful subscription to Rover Stations.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Shipping vendor of choice for orace to customer shipping.
	ShippingVendor *string `mandatory:"false" json:"shippingVendor"`

	// Expected date when customer wants to pickup the device if they chose customer pickup.
	TimePickupExpected *common.SDKTime `mandatory:"false" json:"timePickupExpected"`

	// The flag indicating that customer requests data to be imported to OCI upon Rover cluster return.
	IsImportRequested *bool `mandatory:"false" json:"isImportRequested"`

	// An OCID of a compartment where data will be imported to upon Rover cluster return.
	ImportCompartmentId *string `mandatory:"false" json:"importCompartmentId"`

	// Name of a bucket where files from NFS share will be imported to upon Rover cluster return.
	ImportFileBucket *string `mandatory:"false" json:"importFileBucket"`

	// Validation code returned by data validation tool. Required for return shipping label generation if data import was requested.
	DataValidationCode *string `mandatory:"false" json:"dataValidationCode"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m UpdateRoverClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRoverClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEnclosureTypeEnum(string(m.EnclosureType)); !ok && m.EnclosureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnclosureType: %s. Supported values are: %s.", m.EnclosureType, strings.Join(GetEnclosureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateRoverClusterDetailsShippingPreferenceEnum(string(m.ShippingPreference)); !ok && m.ShippingPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShippingPreference: %s. Supported values are: %s.", m.ShippingPreference, strings.Join(GetUpdateRoverClusterDetailsShippingPreferenceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateRoverClusterDetailsShippingPreferenceEnum Enum with underlying type: string
type UpdateRoverClusterDetailsShippingPreferenceEnum string

// Set of constants representing the allowable values for UpdateRoverClusterDetailsShippingPreferenceEnum
const (
	UpdateRoverClusterDetailsShippingPreferenceOracleShipped  UpdateRoverClusterDetailsShippingPreferenceEnum = "ORACLE_SHIPPED"
	UpdateRoverClusterDetailsShippingPreferenceCustomerPickup UpdateRoverClusterDetailsShippingPreferenceEnum = "CUSTOMER_PICKUP"
)

var mappingUpdateRoverClusterDetailsShippingPreferenceEnum = map[string]UpdateRoverClusterDetailsShippingPreferenceEnum{
	"ORACLE_SHIPPED":  UpdateRoverClusterDetailsShippingPreferenceOracleShipped,
	"CUSTOMER_PICKUP": UpdateRoverClusterDetailsShippingPreferenceCustomerPickup,
}

var mappingUpdateRoverClusterDetailsShippingPreferenceEnumLowerCase = map[string]UpdateRoverClusterDetailsShippingPreferenceEnum{
	"oracle_shipped":  UpdateRoverClusterDetailsShippingPreferenceOracleShipped,
	"customer_pickup": UpdateRoverClusterDetailsShippingPreferenceCustomerPickup,
}

// GetUpdateRoverClusterDetailsShippingPreferenceEnumValues Enumerates the set of values for UpdateRoverClusterDetailsShippingPreferenceEnum
func GetUpdateRoverClusterDetailsShippingPreferenceEnumValues() []UpdateRoverClusterDetailsShippingPreferenceEnum {
	values := make([]UpdateRoverClusterDetailsShippingPreferenceEnum, 0)
	for _, v := range mappingUpdateRoverClusterDetailsShippingPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateRoverClusterDetailsShippingPreferenceEnumStringValues Enumerates the set of values in String for UpdateRoverClusterDetailsShippingPreferenceEnum
func GetUpdateRoverClusterDetailsShippingPreferenceEnumStringValues() []string {
	return []string{
		"ORACLE_SHIPPED",
		"CUSTOMER_PICKUP",
	}
}

// GetMappingUpdateRoverClusterDetailsShippingPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateRoverClusterDetailsShippingPreferenceEnum(val string) (UpdateRoverClusterDetailsShippingPreferenceEnum, bool) {
	enum, ok := mappingUpdateRoverClusterDetailsShippingPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
