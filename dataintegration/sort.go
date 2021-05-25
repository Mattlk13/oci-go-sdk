// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
)

// Sort The information about the sort object.
type Sort struct {

	// The sort clause.
	SortClauses []SortClause `mandatory:"false" json:"sortClauses"`
}

func (m Sort) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Sort) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSort Sort
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeSort
	}{
		"SORT",
		(MarshalTypeSort)(m),
	}

	return json.Marshal(&s)
}
