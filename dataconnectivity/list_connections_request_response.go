// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"net/http"
	"strings"
)

// ListConnectionsRequest wrapper for the ListConnections operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataconnectivity/ListConnections.go.html to see an example of how to use ListConnectionsRequest.
type ListConnectionsRequest struct {

	// The registry Ocid.
	RegistryId *string `mandatory:"true" contributesTo:"path" name:"registryId"`

	// Used to filter by the data asset key of the object.
	DataAssetKey *string `mandatory:"true" contributesTo:"query" name:"dataAssetKey"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Type of the object to filter the results with.
	Type *string `mandatory:"false" contributesTo:"query" name:"type"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListConnectionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListConnectionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If value is FAVORITES_ONLY, then only objects marked as favorite by the requesting user will be included in result. If value is NON_FAVORITES_ONLY, then objects marked as favorites by the requesting user will be skipped. If value is ALL or if not specified, all objects, irrespective of favorites or not will be returned. Default is ALL.
	FavoritesQueryParam ListConnectionsFavoritesQueryParamEnum `mandatory:"false" contributesTo:"query" name:"favoritesQueryParam" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConnectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConnectionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConnectionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListConnectionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConnectionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConnectionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionsFavoritesQueryParamEnum(string(request.FavoritesQueryParam)); !ok && request.FavoritesQueryParam != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FavoritesQueryParam: %s. Supported values are: %s.", request.FavoritesQueryParam, strings.Join(GetListConnectionsFavoritesQueryParamEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConnectionsResponse wrapper for the ListConnections operation
type ListConnectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConnectionSummaryCollection instances
	ConnectionSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListConnectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectionsSortByEnum Enum with underlying type: string
type ListConnectionsSortByEnum string

// Set of constants representing the allowable values for ListConnectionsSortByEnum
const (
	ListConnectionsSortById          ListConnectionsSortByEnum = "id"
	ListConnectionsSortByTimecreated ListConnectionsSortByEnum = "timeCreated"
	ListConnectionsSortByDisplayname ListConnectionsSortByEnum = "displayName"
)

var mappingListConnectionsSortByEnum = map[string]ListConnectionsSortByEnum{
	"id":          ListConnectionsSortById,
	"timeCreated": ListConnectionsSortByTimecreated,
	"displayName": ListConnectionsSortByDisplayname,
}

var mappingListConnectionsSortByEnumLowerCase = map[string]ListConnectionsSortByEnum{
	"id":          ListConnectionsSortById,
	"timecreated": ListConnectionsSortByTimecreated,
	"displayname": ListConnectionsSortByDisplayname,
}

// GetListConnectionsSortByEnumValues Enumerates the set of values for ListConnectionsSortByEnum
func GetListConnectionsSortByEnumValues() []ListConnectionsSortByEnum {
	values := make([]ListConnectionsSortByEnum, 0)
	for _, v := range mappingListConnectionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsSortByEnumStringValues Enumerates the set of values in String for ListConnectionsSortByEnum
func GetListConnectionsSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"displayName",
	}
}

// GetMappingListConnectionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsSortByEnum(val string) (ListConnectionsSortByEnum, bool) {
	enum, ok := mappingListConnectionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionsSortOrderEnum Enum with underlying type: string
type ListConnectionsSortOrderEnum string

// Set of constants representing the allowable values for ListConnectionsSortOrderEnum
const (
	ListConnectionsSortOrderAsc  ListConnectionsSortOrderEnum = "ASC"
	ListConnectionsSortOrderDesc ListConnectionsSortOrderEnum = "DESC"
)

var mappingListConnectionsSortOrderEnum = map[string]ListConnectionsSortOrderEnum{
	"ASC":  ListConnectionsSortOrderAsc,
	"DESC": ListConnectionsSortOrderDesc,
}

var mappingListConnectionsSortOrderEnumLowerCase = map[string]ListConnectionsSortOrderEnum{
	"asc":  ListConnectionsSortOrderAsc,
	"desc": ListConnectionsSortOrderDesc,
}

// GetListConnectionsSortOrderEnumValues Enumerates the set of values for ListConnectionsSortOrderEnum
func GetListConnectionsSortOrderEnumValues() []ListConnectionsSortOrderEnum {
	values := make([]ListConnectionsSortOrderEnum, 0)
	for _, v := range mappingListConnectionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsSortOrderEnumStringValues Enumerates the set of values in String for ListConnectionsSortOrderEnum
func GetListConnectionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConnectionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsSortOrderEnum(val string) (ListConnectionsSortOrderEnum, bool) {
	enum, ok := mappingListConnectionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionsFavoritesQueryParamEnum Enum with underlying type: string
type ListConnectionsFavoritesQueryParamEnum string

// Set of constants representing the allowable values for ListConnectionsFavoritesQueryParamEnum
const (
	ListConnectionsFavoritesQueryParamFavoritesOnly    ListConnectionsFavoritesQueryParamEnum = "FAVORITES_ONLY"
	ListConnectionsFavoritesQueryParamNonFavoritesOnly ListConnectionsFavoritesQueryParamEnum = "NON_FAVORITES_ONLY"
	ListConnectionsFavoritesQueryParamAll              ListConnectionsFavoritesQueryParamEnum = "ALL"
)

var mappingListConnectionsFavoritesQueryParamEnum = map[string]ListConnectionsFavoritesQueryParamEnum{
	"FAVORITES_ONLY":     ListConnectionsFavoritesQueryParamFavoritesOnly,
	"NON_FAVORITES_ONLY": ListConnectionsFavoritesQueryParamNonFavoritesOnly,
	"ALL":                ListConnectionsFavoritesQueryParamAll,
}

var mappingListConnectionsFavoritesQueryParamEnumLowerCase = map[string]ListConnectionsFavoritesQueryParamEnum{
	"favorites_only":     ListConnectionsFavoritesQueryParamFavoritesOnly,
	"non_favorites_only": ListConnectionsFavoritesQueryParamNonFavoritesOnly,
	"all":                ListConnectionsFavoritesQueryParamAll,
}

// GetListConnectionsFavoritesQueryParamEnumValues Enumerates the set of values for ListConnectionsFavoritesQueryParamEnum
func GetListConnectionsFavoritesQueryParamEnumValues() []ListConnectionsFavoritesQueryParamEnum {
	values := make([]ListConnectionsFavoritesQueryParamEnum, 0)
	for _, v := range mappingListConnectionsFavoritesQueryParamEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsFavoritesQueryParamEnumStringValues Enumerates the set of values in String for ListConnectionsFavoritesQueryParamEnum
func GetListConnectionsFavoritesQueryParamEnumStringValues() []string {
	return []string{
		"FAVORITES_ONLY",
		"NON_FAVORITES_ONLY",
		"ALL",
	}
}

// GetMappingListConnectionsFavoritesQueryParamEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsFavoritesQueryParamEnum(val string) (ListConnectionsFavoritesQueryParamEnum, bool) {
	enum, ok := mappingListConnectionsFavoritesQueryParamEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
