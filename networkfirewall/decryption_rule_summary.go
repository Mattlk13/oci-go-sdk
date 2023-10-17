// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DecryptionRuleSummary Summary for Decryption Rule used in the firewall policy rules.
// A Decryption Rule is used to define which traffic should be decrypted by the firewall, and how it should do so.
type DecryptionRuleSummary struct {

	// Name for the decryption rule, must be unique within the policy.
	Name *string `mandatory:"true" json:"name"`

	// Action:
	// * NO_DECRYPT - Matching traffic is not decrypted.
	// * DECRYPT - Matching traffic is decrypted with the specified `secret` according to the specified `decryptionProfile`.
	Action DecryptionActionTypeEnum `mandatory:"true" json:"action"`

	// The name of the decryption profile to use.
	DecryptionProfile *string `mandatory:"true" json:"decryptionProfile"`

	// The name of a mapped secret. Its `type` must match that of the specified decryption profile.
	Secret *string `mandatory:"true" json:"secret"`

	// The priority order in which this rule should be evaluated.
	PriorityOrder *int64 `mandatory:"true" json:"priorityOrder"`

	// OCID of the Network Firewall Policy this application belongs to.
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`
}

func (m DecryptionRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DecryptionRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDecryptionActionTypeEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDecryptionActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
