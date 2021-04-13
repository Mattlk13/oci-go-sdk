// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// REST API for Zero Downtime Migration (Oracle Database Migration Service --ODMS-- as customer-facing service name)
//
// Provides users the ability to perform Zero Downtime migration operations
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v39/common"
)

// UpdateVaultDetails OCI Vault details to store migration and connection credentials secrets. An empty object would result in the removal of the stored details.
type UpdateVaultDetails struct {

	// OCID of the compartment where the secret containing the credentials will be created.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// OCID of the vault
	VaultId *string `mandatory:"false" json:"vaultId"`

	// OCID of the vault encryption key
	KeyId *string `mandatory:"false" json:"keyId"`
}

func (m UpdateVaultDetails) String() string {
	return common.PointerString(m)
}