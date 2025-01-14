// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenericEntity This is an extensible entity type for users, and the base entity type for some annotation formats.
type GenericEntity struct {

	// A collection of label entities.
	Labels []Label `mandatory:"true" json:"labels"`

	DocumentEntityMetadata *DocumentEntityMetadata `mandatory:"false" json:"documentEntityMetadata"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`
}

func (m GenericEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenericEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GenericEntity) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGenericEntity GenericEntity
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeGenericEntity
	}{
		"GENERIC",
		(MarshalTypeGenericEntity)(m),
	}

	return json.Marshal(&s)
}
