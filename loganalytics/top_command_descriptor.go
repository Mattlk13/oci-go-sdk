// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v49/common"
)

// TopCommandDescriptor Command descriptor for querylanguage TOP command.
type TopCommandDescriptor struct {

	// Command fragment display string from user specified query string formatted by query builder.
	DisplayQueryString *string `mandatory:"true" json:"displayQueryString"`

	// Command fragment internal string from user specified query string formatted by query builder.
	InternalQueryString *string `mandatory:"true" json:"internalQueryString"`

	// querylanguage command designation for example; reporting vs filtering
	Category *string `mandatory:"false" json:"category"`

	// Fields referenced in command fragment from user specified query string.
	ReferencedFields []AbstractField `mandatory:"false" json:"referencedFields"`

	// Fields declared in command fragment from user specified query string.
	DeclaredFields []AbstractField `mandatory:"false" json:"declaredFields"`

	// Value from queryString for top command limit argument.
	Limit *int `mandatory:"false" json:"limit"`
}

//GetDisplayQueryString returns DisplayQueryString
func (m TopCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m TopCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m TopCommandDescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m TopCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m TopCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

func (m TopCommandDescriptor) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TopCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTopCommandDescriptor TopCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeTopCommandDescriptor
	}{
		"TOP",
		(MarshalTypeTopCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TopCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string         `json:"category"`
		ReferencedFields    []abstractfield `json:"referencedFields"`
		DeclaredFields      []abstractfield `json:"declaredFields"`
		Limit               *int            `json:"limit"`
		DisplayQueryString  *string         `json:"displayQueryString"`
		InternalQueryString *string         `json:"internalQueryString"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Category = model.Category

	m.ReferencedFields = make([]AbstractField, len(model.ReferencedFields))
	for i, n := range model.ReferencedFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ReferencedFields[i] = nn.(AbstractField)
		} else {
			m.ReferencedFields[i] = nil
		}
	}

	m.DeclaredFields = make([]AbstractField, len(model.DeclaredFields))
	for i, n := range model.DeclaredFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DeclaredFields[i] = nn.(AbstractField)
		} else {
			m.DeclaredFields[i] = nil
		}
	}

	m.Limit = model.Limit

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}
