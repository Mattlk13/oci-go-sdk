// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v30/common"
)

// DirectNamedFieldMap A named field map.
type DirectNamedFieldMap struct {

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// Reference to a typed object.
	SourceTypedObject *string `mandatory:"false" json:"sourceTypedObject"`

	// Reference to a typed object
	TargetTypedObject *string `mandatory:"false" json:"targetTypedObject"`

	// The source field name.
	SourceFieldName *string `mandatory:"false" json:"sourceFieldName"`

	// The target field name.
	TargetFieldName *string `mandatory:"false" json:"targetFieldName"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

//GetDescription returns Description
func (m DirectNamedFieldMap) GetDescription() *string {
	return m.Description
}

func (m DirectNamedFieldMap) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DirectNamedFieldMap) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDirectNamedFieldMap DirectNamedFieldMap
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDirectNamedFieldMap
	}{
		"DIRECT_NAMED_FIELD_MAP",
		(MarshalTypeDirectNamedFieldMap)(m),
	}

	return json.Marshal(&s)
}
