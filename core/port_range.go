// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// PortRange The representation of PortRange
type PortRange struct {

	// The maximum port number, which must not be less than the minimum port number. To specify
	// a single port number, set both the min and max to the same value.
	Max *int `mandatory:"true" json:"max"`

	// The minimum port number, which must not be greater than the maximum port number.
	Min *int `mandatory:"true" json:"min"`
}

func (m PortRange) String() string {
	return common.PointerString(m)
}
