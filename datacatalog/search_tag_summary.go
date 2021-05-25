// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// SearchTagSummary Represents the association of an object to a term. Returned as part of search result.
type SearchTagSummary struct {

	// Name of the tag that matches the term name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique tag key that is immutable.
	Key *string `mandatory:"false" json:"key"`
}

func (m SearchTagSummary) String() string {
	return common.PointerString(m)
}
