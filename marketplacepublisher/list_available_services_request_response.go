// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailableServicesRequest wrapper for the ListAvailableServices operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAvailableServices.go.html to see an example of how to use ListAvailableServicesRequest.
type ListAvailableServicesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAvailableServicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAvailableServicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableServicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableServicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableServicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableServicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableServicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailableServicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableServicesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableServicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableServicesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableServicesResponse wrapper for the ListAvailableServices operation
type ListAvailableServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AvailableServiceCollection instances
	AvailableServiceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableServicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableServicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableServicesSortOrderEnum Enum with underlying type: string
type ListAvailableServicesSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableServicesSortOrderEnum
const (
	ListAvailableServicesSortOrderAsc  ListAvailableServicesSortOrderEnum = "ASC"
	ListAvailableServicesSortOrderDesc ListAvailableServicesSortOrderEnum = "DESC"
)

var mappingListAvailableServicesSortOrderEnum = map[string]ListAvailableServicesSortOrderEnum{
	"ASC":  ListAvailableServicesSortOrderAsc,
	"DESC": ListAvailableServicesSortOrderDesc,
}

var mappingListAvailableServicesSortOrderEnumLowerCase = map[string]ListAvailableServicesSortOrderEnum{
	"asc":  ListAvailableServicesSortOrderAsc,
	"desc": ListAvailableServicesSortOrderDesc,
}

// GetListAvailableServicesSortOrderEnumValues Enumerates the set of values for ListAvailableServicesSortOrderEnum
func GetListAvailableServicesSortOrderEnumValues() []ListAvailableServicesSortOrderEnum {
	values := make([]ListAvailableServicesSortOrderEnum, 0)
	for _, v := range mappingListAvailableServicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableServicesSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableServicesSortOrderEnum
func GetListAvailableServicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableServicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableServicesSortOrderEnum(val string) (ListAvailableServicesSortOrderEnum, bool) {
	enum, ok := mappingListAvailableServicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableServicesSortByEnum Enum with underlying type: string
type ListAvailableServicesSortByEnum string

// Set of constants representing the allowable values for ListAvailableServicesSortByEnum
const (
	ListAvailableServicesSortByTimecreated ListAvailableServicesSortByEnum = "timeCreated"
	ListAvailableServicesSortByDisplayname ListAvailableServicesSortByEnum = "displayName"
)

var mappingListAvailableServicesSortByEnum = map[string]ListAvailableServicesSortByEnum{
	"timeCreated": ListAvailableServicesSortByTimecreated,
	"displayName": ListAvailableServicesSortByDisplayname,
}

var mappingListAvailableServicesSortByEnumLowerCase = map[string]ListAvailableServicesSortByEnum{
	"timecreated": ListAvailableServicesSortByTimecreated,
	"displayname": ListAvailableServicesSortByDisplayname,
}

// GetListAvailableServicesSortByEnumValues Enumerates the set of values for ListAvailableServicesSortByEnum
func GetListAvailableServicesSortByEnumValues() []ListAvailableServicesSortByEnum {
	values := make([]ListAvailableServicesSortByEnum, 0)
	for _, v := range mappingListAvailableServicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableServicesSortByEnumStringValues Enumerates the set of values in String for ListAvailableServicesSortByEnum
func GetListAvailableServicesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAvailableServicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableServicesSortByEnum(val string) (ListAvailableServicesSortByEnum, bool) {
	enum, ok := mappingListAvailableServicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
