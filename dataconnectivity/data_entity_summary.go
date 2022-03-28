// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// DataEntitySummary The data entity summary object.
type DataEntitySummary interface {
	GetMetadata() *ObjectMetadata
}

type dataentitysummary struct {
	JsonData  []byte
	Metadata  *ObjectMetadata `mandatory:"false" json:"metadata"`
	ModelType string          `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dataentitysummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataentitysummary dataentitysummary
	s := struct {
		Model Unmarshalerdataentitysummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Metadata = s.Model.Metadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataentitysummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "FILE_ENTITY":
		mm := DataEntitySummaryFromFile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TABLE_ENTITY":
		mm := DataEntitySummaryFromTable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_STORE_ENTITY":
		mm := DataEntitySummaryFromDataStore{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_ENTITY":
		mm := DataEntitySummaryFromSql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIEW_ENTITY":
		mm := DataEntitySummaryFromView{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetMetadata returns Metadata
func (m dataentitysummary) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m dataentitysummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataentitysummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataEntitySummaryModelTypeEnum Enum with underlying type: string
type DataEntitySummaryModelTypeEnum string

// Set of constants representing the allowable values for DataEntitySummaryModelTypeEnum
const (
	DataEntitySummaryModelTypeViewEntity      DataEntitySummaryModelTypeEnum = "VIEW_ENTITY"
	DataEntitySummaryModelTypeTableEntity     DataEntitySummaryModelTypeEnum = "TABLE_ENTITY"
	DataEntitySummaryModelTypeFileEntity      DataEntitySummaryModelTypeEnum = "FILE_ENTITY"
	DataEntitySummaryModelTypeDataStoreEntity DataEntitySummaryModelTypeEnum = "DATA_STORE_ENTITY"
	DataEntitySummaryModelTypeSqlEntity       DataEntitySummaryModelTypeEnum = "SQL_ENTITY"
)

var mappingDataEntitySummaryModelTypeEnum = map[string]DataEntitySummaryModelTypeEnum{
	"VIEW_ENTITY":       DataEntitySummaryModelTypeViewEntity,
	"TABLE_ENTITY":      DataEntitySummaryModelTypeTableEntity,
	"FILE_ENTITY":       DataEntitySummaryModelTypeFileEntity,
	"DATA_STORE_ENTITY": DataEntitySummaryModelTypeDataStoreEntity,
	"SQL_ENTITY":        DataEntitySummaryModelTypeSqlEntity,
}

var mappingDataEntitySummaryModelTypeEnumLowerCase = map[string]DataEntitySummaryModelTypeEnum{
	"view_entity":       DataEntitySummaryModelTypeViewEntity,
	"table_entity":      DataEntitySummaryModelTypeTableEntity,
	"file_entity":       DataEntitySummaryModelTypeFileEntity,
	"data_store_entity": DataEntitySummaryModelTypeDataStoreEntity,
	"sql_entity":        DataEntitySummaryModelTypeSqlEntity,
}

// GetDataEntitySummaryModelTypeEnumValues Enumerates the set of values for DataEntitySummaryModelTypeEnum
func GetDataEntitySummaryModelTypeEnumValues() []DataEntitySummaryModelTypeEnum {
	values := make([]DataEntitySummaryModelTypeEnum, 0)
	for _, v := range mappingDataEntitySummaryModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataEntitySummaryModelTypeEnumStringValues Enumerates the set of values in String for DataEntitySummaryModelTypeEnum
func GetDataEntitySummaryModelTypeEnumStringValues() []string {
	return []string{
		"VIEW_ENTITY",
		"TABLE_ENTITY",
		"FILE_ENTITY",
		"DATA_STORE_ENTITY",
		"SQL_ENTITY",
	}
}

// GetMappingDataEntitySummaryModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataEntitySummaryModelTypeEnum(val string) (DataEntitySummaryModelTypeEnum, bool) {
	enum, ok := mappingDataEntitySummaryModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
