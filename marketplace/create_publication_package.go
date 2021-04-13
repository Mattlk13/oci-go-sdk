// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v39/common"
)

// CreatePublicationPackage A base object for the properties of the package
type CreatePublicationPackage interface {

	// The version of the package
	GetPackageVersion() *string

	GetOperatingSystem() *OperatingSystem

	// End User License Agreeement that a consumer of this listing has to accept
	GetEula() []Eula
}

type createpublicationpackage struct {
	JsonData        []byte
	PackageVersion  *string          `mandatory:"true" json:"packageVersion"`
	OperatingSystem *OperatingSystem `mandatory:"true" json:"operatingSystem"`
	Eula            []Eula           `mandatory:"true" json:"eula"`
	PackageType     string           `json:"packageType"`
}

// UnmarshalJSON unmarshals json
func (m *createpublicationpackage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatepublicationpackage createpublicationpackage
	s := struct {
		Model Unmarshalercreatepublicationpackage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PackageVersion = s.Model.PackageVersion
	m.OperatingSystem = s.Model.OperatingSystem
	m.Eula = s.Model.Eula
	m.PackageType = s.Model.PackageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createpublicationpackage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PackageType {
	case "IMAGE":
		mm := CreateImagePublicationPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetPackageVersion returns PackageVersion
func (m createpublicationpackage) GetPackageVersion() *string {
	return m.PackageVersion
}

//GetOperatingSystem returns OperatingSystem
func (m createpublicationpackage) GetOperatingSystem() *OperatingSystem {
	return m.OperatingSystem
}

//GetEula returns Eula
func (m createpublicationpackage) GetEula() []Eula {
	return m.Eula
}

func (m createpublicationpackage) String() string {
	return common.PointerString(m)
}