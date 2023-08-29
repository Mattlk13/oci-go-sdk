// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddUtilityNodesDetails The information about added utility nodes.
type AddUtilityNodesDetails struct {

	// Base-64 encoded Cluster Admin Password for cluster admin user.
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// Number of additional utility nodes for the cluster.
	NumberOfUtilityNodes *int `mandatory:"true" json:"numberOfUtilityNodes"`

	// Shape of the node. It's a read-only property derived from existing Utility node.
	Shape *string `mandatory:"false" json:"shape"`

	// The size of block volume in GB to be attached to the given node. It's a read-only property.
	BlockVolumeSizeInGBs *int64 `mandatory:"false" json:"blockVolumeSizeInGBs"`

	ShapeConfig *ShapeConfigDetails `mandatory:"false" json:"shapeConfig"`
}

func (m AddUtilityNodesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddUtilityNodesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}