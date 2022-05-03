// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ADM API
//
// Use the ADM API to create ADM knowledge base and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/adm/using/home.htm).
//

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KnowledgeBaseCollection Collection of Knowledge Base summaries.
type KnowledgeBaseCollection struct {

	// List of Knowledge Base summaries.
	Items []KnowledgeBaseSummary `mandatory:"true" json:"items"`
}

func (m KnowledgeBaseCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KnowledgeBaseCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
