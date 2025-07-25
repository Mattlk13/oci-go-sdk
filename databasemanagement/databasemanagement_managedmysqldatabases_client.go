// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ManagedMySqlDatabasesClient a client for ManagedMySqlDatabases
type ManagedMySqlDatabasesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewManagedMySqlDatabasesClientWithConfigurationProvider Creates a new default ManagedMySqlDatabases client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewManagedMySqlDatabasesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ManagedMySqlDatabasesClient, err error) {
	if enabled := common.CheckForEnabledServices("databasemanagement"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newManagedMySqlDatabasesClientFromBaseClient(baseClient, provider)
}

// NewManagedMySqlDatabasesClientWithOboToken Creates a new default ManagedMySqlDatabases client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewManagedMySqlDatabasesClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ManagedMySqlDatabasesClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newManagedMySqlDatabasesClientFromBaseClient(baseClient, configProvider)
}

func newManagedMySqlDatabasesClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ManagedMySqlDatabasesClient, err error) {
	// ManagedMySqlDatabases service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ManagedMySqlDatabases"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ManagedMySqlDatabasesClient{BaseClient: baseClient}
	client.BasePath = "20201101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ManagedMySqlDatabasesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasemanagement", "https://dbmgmt.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ManagedMySqlDatabasesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ManagedMySqlDatabasesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeMysqlDatabaseManagementType Changes the management type for a HeatWave MySQL instance, from BASIC to FULL and vice versa.
// It can also be used to enable or disable database management.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangeMysqlDatabaseManagementType.go.html to see an example of how to use ChangeMysqlDatabaseManagementType API.
// A default retry strategy applies to this operation ChangeMysqlDatabaseManagementType()
func (client ManagedMySqlDatabasesClient) ChangeMysqlDatabaseManagementType(ctx context.Context, request ChangeMysqlDatabaseManagementTypeRequest) (response ChangeMysqlDatabaseManagementTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeMysqlDatabaseManagementType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMysqlDatabaseManagementTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMysqlDatabaseManagementTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMysqlDatabaseManagementTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMysqlDatabaseManagementTypeResponse")
	}
	return
}

// changeMysqlDatabaseManagementType implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) changeMysqlDatabaseManagementType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedMySqlDatabases/{managedMySqlDatabaseId}/actions/changeManagementType", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMysqlDatabaseManagementTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/ChangeMysqlDatabaseManagementType"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "ChangeMysqlDatabaseManagementType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBinaryLogInformation Retrieves information pertaining to the binary log of a specific MySQL server.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetBinaryLogInformation.go.html to see an example of how to use GetBinaryLogInformation API.
// A default retry strategy applies to this operation GetBinaryLogInformation()
func (client ManagedMySqlDatabasesClient) GetBinaryLogInformation(ctx context.Context, request GetBinaryLogInformationRequest) (response GetBinaryLogInformationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBinaryLogInformation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBinaryLogInformationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBinaryLogInformationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBinaryLogInformationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBinaryLogInformationResponse")
	}
	return
}

// getBinaryLogInformation implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) getBinaryLogInformation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/binaryLogInformation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBinaryLogInformationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/GetBinaryLogInformation"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "GetBinaryLogInformation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGeneralReplicationInformation Retrieves general information regarding the replication of a specific MySQL server.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetGeneralReplicationInformation.go.html to see an example of how to use GetGeneralReplicationInformation API.
// A default retry strategy applies to this operation GetGeneralReplicationInformation()
func (client ManagedMySqlDatabasesClient) GetGeneralReplicationInformation(ctx context.Context, request GetGeneralReplicationInformationRequest) (response GetGeneralReplicationInformationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGeneralReplicationInformation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGeneralReplicationInformationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGeneralReplicationInformationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGeneralReplicationInformationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGeneralReplicationInformationResponse")
	}
	return
}

// getGeneralReplicationInformation implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) getGeneralReplicationInformation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/generalReplicationInformation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGeneralReplicationInformationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/GetGeneralReplicationInformation"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "GetGeneralReplicationInformation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetHeatWaveFleetMetric Gets the health metrics for a fleet of HeatWave clusters in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetHeatWaveFleetMetric.go.html to see an example of how to use GetHeatWaveFleetMetric API.
func (client ManagedMySqlDatabasesClient) GetHeatWaveFleetMetric(ctx context.Context, request GetHeatWaveFleetMetricRequest) (response GetHeatWaveFleetMetricResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getHeatWaveFleetMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHeatWaveFleetMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHeatWaveFleetMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHeatWaveFleetMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHeatWaveFleetMetricResponse")
	}
	return
}

// getHeatWaveFleetMetric implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) getHeatWaveFleetMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/heatWaveFleetMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetHeatWaveFleetMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/HeatWaveFleetMetrics/GetHeatWaveFleetMetric"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "GetHeatWaveFleetMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedMySqlDatabase Retrieves the general information for a specific MySQL Database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetManagedMySqlDatabase.go.html to see an example of how to use GetManagedMySqlDatabase API.
func (client ManagedMySqlDatabasesClient) GetManagedMySqlDatabase(ctx context.Context, request GetManagedMySqlDatabaseRequest) (response GetManagedMySqlDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagedMySqlDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagedMySqlDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagedMySqlDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagedMySqlDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagedMySqlDatabaseResponse")
	}
	return
}

// getManagedMySqlDatabase implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) getManagedMySqlDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagedMySqlDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/GetManagedMySqlDatabase"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "GetManagedMySqlDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMySqlFleetMetric Gets the health metrics for a fleet of MySQL Databases in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetMySqlFleetMetric.go.html to see an example of how to use GetMySqlFleetMetric API.
func (client ManagedMySqlDatabasesClient) GetMySqlFleetMetric(ctx context.Context, request GetMySqlFleetMetricRequest) (response GetMySqlFleetMetricResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMySqlFleetMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMySqlFleetMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMySqlFleetMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMySqlFleetMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMySqlFleetMetricResponse")
	}
	return
}

// getMySqlFleetMetric implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) getMySqlFleetMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mysqlFleetMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMySqlFleetMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/MySqlFleetMetrics/GetMySqlFleetMetric"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "GetMySqlFleetMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMySqlQueryDetails Retrieves query sample details, explain plan and potential warnings for a given digest.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetMySqlQueryDetails.go.html to see an example of how to use GetMySqlQueryDetails API.
// A default retry strategy applies to this operation GetMySqlQueryDetails()
func (client ManagedMySqlDatabasesClient) GetMySqlQueryDetails(ctx context.Context, request GetMySqlQueryDetailsRequest) (response GetMySqlQueryDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMySqlQueryDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMySqlQueryDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMySqlQueryDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMySqlQueryDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMySqlQueryDetailsResponse")
	}
	return
}

// getMySqlQueryDetails implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) getMySqlQueryDetails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/queryDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMySqlQueryDetailsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/GetMySqlQueryDetails"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "GetMySqlQueryDetails", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHighAvailabilityMembers Retrieves information about the high availability members of a specific MySQL server's replication group.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListHighAvailabilityMembers.go.html to see an example of how to use ListHighAvailabilityMembers API.
// A default retry strategy applies to this operation ListHighAvailabilityMembers()
func (client ManagedMySqlDatabasesClient) ListHighAvailabilityMembers(ctx context.Context, request ListHighAvailabilityMembersRequest) (response ListHighAvailabilityMembersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHighAvailabilityMembers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHighAvailabilityMembersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHighAvailabilityMembersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHighAvailabilityMembersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHighAvailabilityMembersResponse")
	}
	return
}

// listHighAvailabilityMembers implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) listHighAvailabilityMembers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/highAvailabilityMembers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHighAvailabilityMembersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/ListHighAvailabilityMembers"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "ListHighAvailabilityMembers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInboundReplications Retrieves information about the inbound replications of a specific MySQL server.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListInboundReplications.go.html to see an example of how to use ListInboundReplications API.
// A default retry strategy applies to this operation ListInboundReplications()
func (client ManagedMySqlDatabasesClient) ListInboundReplications(ctx context.Context, request ListInboundReplicationsRequest) (response ListInboundReplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInboundReplications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInboundReplicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInboundReplicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInboundReplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInboundReplicationsResponse")
	}
	return
}

// listInboundReplications implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) listInboundReplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/inboundReplications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInboundReplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/ListInboundReplications"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "ListInboundReplications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedMySqlDatabaseConfigurationData Retrieves configuration data for a specific MySQL database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedMySqlDatabaseConfigurationData.go.html to see an example of how to use ListManagedMySqlDatabaseConfigurationData API.
// A default retry strategy applies to this operation ListManagedMySqlDatabaseConfigurationData()
func (client ManagedMySqlDatabasesClient) ListManagedMySqlDatabaseConfigurationData(ctx context.Context, request ListManagedMySqlDatabaseConfigurationDataRequest) (response ListManagedMySqlDatabaseConfigurationDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedMySqlDatabaseConfigurationData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedMySqlDatabaseConfigurationDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedMySqlDatabaseConfigurationDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedMySqlDatabaseConfigurationDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedMySqlDatabaseConfigurationDataResponse")
	}
	return
}

// listManagedMySqlDatabaseConfigurationData implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) listManagedMySqlDatabaseConfigurationData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/configurationData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedMySqlDatabaseConfigurationDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/ListManagedMySqlDatabaseConfigurationData"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "ListManagedMySqlDatabaseConfigurationData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedMySqlDatabaseSqlData Retrieves the SQL performance data for a specific MySQL database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedMySqlDatabaseSqlData.go.html to see an example of how to use ListManagedMySqlDatabaseSqlData API.
// A default retry strategy applies to this operation ListManagedMySqlDatabaseSqlData()
func (client ManagedMySqlDatabasesClient) ListManagedMySqlDatabaseSqlData(ctx context.Context, request ListManagedMySqlDatabaseSqlDataRequest) (response ListManagedMySqlDatabaseSqlDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedMySqlDatabaseSqlData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedMySqlDatabaseSqlDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedMySqlDatabaseSqlDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedMySqlDatabaseSqlDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedMySqlDatabaseSqlDataResponse")
	}
	return
}

// listManagedMySqlDatabaseSqlData implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) listManagedMySqlDatabaseSqlData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/sqlData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedMySqlDatabaseSqlDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/ListManagedMySqlDatabaseSqlData"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "ListManagedMySqlDatabaseSqlData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedMySqlDatabases Gets the list of Managed MySQL Databases in a specific compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedMySqlDatabases.go.html to see an example of how to use ListManagedMySqlDatabases API.
func (client ManagedMySqlDatabasesClient) ListManagedMySqlDatabases(ctx context.Context, request ListManagedMySqlDatabasesRequest) (response ListManagedMySqlDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedMySqlDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedMySqlDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedMySqlDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedMySqlDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedMySqlDatabasesResponse")
	}
	return
}

// listManagedMySqlDatabases implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) listManagedMySqlDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedMySqlDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabaseCollection/ListManagedMySqlDatabases"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "ListManagedMySqlDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMySqlDigestErrors Retrieves any potential errors for a given digest.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListMySqlDigestErrors.go.html to see an example of how to use ListMySqlDigestErrors API.
// A default retry strategy applies to this operation ListMySqlDigestErrors()
func (client ManagedMySqlDatabasesClient) ListMySqlDigestErrors(ctx context.Context, request ListMySqlDigestErrorsRequest) (response ListMySqlDigestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMySqlDigestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMySqlDigestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMySqlDigestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMySqlDigestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMySqlDigestErrorsResponse")
	}
	return
}

// listMySqlDigestErrors implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) listMySqlDigestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/digestErrors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMySqlDigestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/ListMySqlDigestErrors"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "ListMySqlDigestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOutboundReplications Retrieves information pertaining to the outbound replications of a specific MySQL server.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListOutboundReplications.go.html to see an example of how to use ListOutboundReplications API.
// A default retry strategy applies to this operation ListOutboundReplications()
func (client ManagedMySqlDatabasesClient) ListOutboundReplications(ctx context.Context, request ListOutboundReplicationsRequest) (response ListOutboundReplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOutboundReplications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOutboundReplicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOutboundReplicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOutboundReplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOutboundReplicationsResponse")
	}
	return
}

// listOutboundReplications implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) listOutboundReplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/outboundReplications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOutboundReplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/ListOutboundReplications"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "ListOutboundReplications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeManagedMySqlDatabaseAvailabilityMetrics Gets the availability metrics for the MySQL Database specified by managedMySqlDatabaseId.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeManagedMySqlDatabaseAvailabilityMetrics.go.html to see an example of how to use SummarizeManagedMySqlDatabaseAvailabilityMetrics API.
// A default retry strategy applies to this operation SummarizeManagedMySqlDatabaseAvailabilityMetrics()
func (client ManagedMySqlDatabasesClient) SummarizeManagedMySqlDatabaseAvailabilityMetrics(ctx context.Context, request SummarizeManagedMySqlDatabaseAvailabilityMetricsRequest) (response SummarizeManagedMySqlDatabaseAvailabilityMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeManagedMySqlDatabaseAvailabilityMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeManagedMySqlDatabaseAvailabilityMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeManagedMySqlDatabaseAvailabilityMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeManagedMySqlDatabaseAvailabilityMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeManagedMySqlDatabaseAvailabilityMetricsResponse")
	}
	return
}

// summarizeManagedMySqlDatabaseAvailabilityMetrics implements the OCIOperation interface (enables retrying operations)
func (client ManagedMySqlDatabasesClient) summarizeManagedMySqlDatabaseAvailabilityMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedMySqlDatabases/{managedMySqlDatabaseId}/availabilityMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeManagedMySqlDatabaseAvailabilityMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedMySqlDatabase/SummarizeManagedMySqlDatabaseAvailabilityMetrics"
		err = common.PostProcessServiceError(err, "ManagedMySqlDatabases", "SummarizeManagedMySqlDatabaseAvailabilityMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
