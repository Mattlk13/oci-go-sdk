// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTransferDevicesRequest wrapper for the ListTransferDevices operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/ListTransferDevices.go.html to see an example of how to use ListTransferDevicesRequest.
type ListTransferDevicesRequest struct {

	// ID of the Transfer Job
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// filtering by lifecycleState
	LifecycleState ListTransferDevicesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// filtering by displayName
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTransferDevicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTransferDevicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTransferDevicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTransferDevicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTransferDevicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTransferDevicesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTransferDevicesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTransferDevicesResponse wrapper for the ListTransferDevices operation
type ListTransferDevicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The MultipleTransferDevices instance
	MultipleTransferDevices `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTransferDevicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTransferDevicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTransferDevicesLifecycleStateEnum Enum with underlying type: string
type ListTransferDevicesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTransferDevicesLifecycleStateEnum
const (
	ListTransferDevicesLifecycleStatePreparing  ListTransferDevicesLifecycleStateEnum = "PREPARING"
	ListTransferDevicesLifecycleStateReady      ListTransferDevicesLifecycleStateEnum = "READY"
	ListTransferDevicesLifecycleStatePackaged   ListTransferDevicesLifecycleStateEnum = "PACKAGED"
	ListTransferDevicesLifecycleStateActive     ListTransferDevicesLifecycleStateEnum = "ACTIVE"
	ListTransferDevicesLifecycleStateProcessing ListTransferDevicesLifecycleStateEnum = "PROCESSING"
	ListTransferDevicesLifecycleStateComplete   ListTransferDevicesLifecycleStateEnum = "COMPLETE"
	ListTransferDevicesLifecycleStateMissing    ListTransferDevicesLifecycleStateEnum = "MISSING"
	ListTransferDevicesLifecycleStateError      ListTransferDevicesLifecycleStateEnum = "ERROR"
	ListTransferDevicesLifecycleStateDeleted    ListTransferDevicesLifecycleStateEnum = "DELETED"
	ListTransferDevicesLifecycleStateCancelled  ListTransferDevicesLifecycleStateEnum = "CANCELLED"
)

var mappingListTransferDevicesLifecycleStateEnum = map[string]ListTransferDevicesLifecycleStateEnum{
	"PREPARING":  ListTransferDevicesLifecycleStatePreparing,
	"READY":      ListTransferDevicesLifecycleStateReady,
	"PACKAGED":   ListTransferDevicesLifecycleStatePackaged,
	"ACTIVE":     ListTransferDevicesLifecycleStateActive,
	"PROCESSING": ListTransferDevicesLifecycleStateProcessing,
	"COMPLETE":   ListTransferDevicesLifecycleStateComplete,
	"MISSING":    ListTransferDevicesLifecycleStateMissing,
	"ERROR":      ListTransferDevicesLifecycleStateError,
	"DELETED":    ListTransferDevicesLifecycleStateDeleted,
	"CANCELLED":  ListTransferDevicesLifecycleStateCancelled,
}

var mappingListTransferDevicesLifecycleStateEnumLowerCase = map[string]ListTransferDevicesLifecycleStateEnum{
	"preparing":  ListTransferDevicesLifecycleStatePreparing,
	"ready":      ListTransferDevicesLifecycleStateReady,
	"packaged":   ListTransferDevicesLifecycleStatePackaged,
	"active":     ListTransferDevicesLifecycleStateActive,
	"processing": ListTransferDevicesLifecycleStateProcessing,
	"complete":   ListTransferDevicesLifecycleStateComplete,
	"missing":    ListTransferDevicesLifecycleStateMissing,
	"error":      ListTransferDevicesLifecycleStateError,
	"deleted":    ListTransferDevicesLifecycleStateDeleted,
	"cancelled":  ListTransferDevicesLifecycleStateCancelled,
}

// GetListTransferDevicesLifecycleStateEnumValues Enumerates the set of values for ListTransferDevicesLifecycleStateEnum
func GetListTransferDevicesLifecycleStateEnumValues() []ListTransferDevicesLifecycleStateEnum {
	values := make([]ListTransferDevicesLifecycleStateEnum, 0)
	for _, v := range mappingListTransferDevicesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTransferDevicesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTransferDevicesLifecycleStateEnum
func GetListTransferDevicesLifecycleStateEnumStringValues() []string {
	return []string{
		"PREPARING",
		"READY",
		"PACKAGED",
		"ACTIVE",
		"PROCESSING",
		"COMPLETE",
		"MISSING",
		"ERROR",
		"DELETED",
		"CANCELLED",
	}
}

// GetMappingListTransferDevicesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTransferDevicesLifecycleStateEnum(val string) (ListTransferDevicesLifecycleStateEnum, bool) {
	enum, ok := mappingListTransferDevicesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
