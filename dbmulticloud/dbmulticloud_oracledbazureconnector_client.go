// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
//

package dbmulticloud

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OracleDBAzureConnectorClient a client for OracleDBAzureConnector
type OracleDBAzureConnectorClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOracleDBAzureConnectorClientWithConfigurationProvider Creates a new default OracleDBAzureConnector client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOracleDBAzureConnectorClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OracleDBAzureConnectorClient, err error) {
	if enabled := common.CheckForEnabledServices("dbmulticloud"); !enabled {
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
	return newOracleDBAzureConnectorClientFromBaseClient(baseClient, provider)
}

// NewOracleDBAzureConnectorClientWithOboToken Creates a new default OracleDBAzureConnector client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOracleDBAzureConnectorClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OracleDBAzureConnectorClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOracleDBAzureConnectorClientFromBaseClient(baseClient, configProvider)
}

func newOracleDBAzureConnectorClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OracleDBAzureConnectorClient, err error) {
	// OracleDBAzureConnector service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OracleDBAzureConnector"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OracleDBAzureConnectorClient{BaseClient: baseClient}
	client.BasePath = "20240501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OracleDBAzureConnectorClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dbmulticloud", "https://dbmulticloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OracleDBAzureConnectorClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OracleDBAzureConnectorClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOracleDbAzureConnectorCompartment Moves the Oracle DB Azure Connector Resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ChangeOracleDbAzureConnectorCompartment.go.html to see an example of how to use ChangeOracleDbAzureConnectorCompartment API.
// A default retry strategy applies to this operation ChangeOracleDbAzureConnectorCompartment()
func (client OracleDBAzureConnectorClient) ChangeOracleDbAzureConnectorCompartment(ctx context.Context, request ChangeOracleDbAzureConnectorCompartmentRequest) (response ChangeOracleDbAzureConnectorCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbAzureConnectorCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbAzureConnectorCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbAzureConnectorCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbAzureConnectorCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbAzureConnectorCompartmentResponse")
	}
	return
}

// changeOracleDbAzureConnectorCompartment implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureConnectorClient) changeOracleDbAzureConnectorCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureConnector/{oracleDbAzureConnectorId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbAzureConnectorCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureConnector/ChangeOracleDbAzureConnectorCompartment"
		err = common.PostProcessServiceError(err, "OracleDBAzureConnector", "ChangeOracleDbAzureConnectorCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbAzureConnector Creates Oracle DB Azure Connector Resource and configured Azure Identity in OCI Database Resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/CreateOracleDbAzureConnector.go.html to see an example of how to use CreateOracleDbAzureConnector API.
// A default retry strategy applies to this operation CreateOracleDbAzureConnector()
func (client OracleDBAzureConnectorClient) CreateOracleDbAzureConnector(ctx context.Context, request CreateOracleDbAzureConnectorRequest) (response CreateOracleDbAzureConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbAzureConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbAzureConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbAzureConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbAzureConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbAzureConnectorResponse")
	}
	return
}

// createOracleDbAzureConnector implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureConnectorClient) createOracleDbAzureConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureConnector", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbAzureConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureConnector/CreateOracleDbAzureConnector"
		err = common.PostProcessServiceError(err, "OracleDBAzureConnector", "CreateOracleDbAzureConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbAzureConnector Delete Oracle DB Azure Connector Resource and delete Azure Arc Identity too from Database Resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/DeleteOracleDbAzureConnector.go.html to see an example of how to use DeleteOracleDbAzureConnector API.
// A default retry strategy applies to this operation DeleteOracleDbAzureConnector()
func (client OracleDBAzureConnectorClient) DeleteOracleDbAzureConnector(ctx context.Context, request DeleteOracleDbAzureConnectorRequest) (response DeleteOracleDbAzureConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbAzureConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbAzureConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbAzureConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbAzureConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbAzureConnectorResponse")
	}
	return
}

// deleteOracleDbAzureConnector implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureConnectorClient) deleteOracleDbAzureConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbAzureConnector/{oracleDbAzureConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbAzureConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureConnector/DeleteOracleDbAzureConnector"
		err = common.PostProcessServiceError(err, "OracleDBAzureConnector", "DeleteOracleDbAzureConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbAzureConnector Get Oracle DB Azure Connector Resource form a particular Resource ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureConnector.go.html to see an example of how to use GetOracleDbAzureConnector API.
// A default retry strategy applies to this operation GetOracleDbAzureConnector()
func (client OracleDBAzureConnectorClient) GetOracleDbAzureConnector(ctx context.Context, request GetOracleDbAzureConnectorRequest) (response GetOracleDbAzureConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbAzureConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbAzureConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbAzureConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbAzureConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbAzureConnectorResponse")
	}
	return
}

// getOracleDbAzureConnector implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureConnectorClient) getOracleDbAzureConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureConnector/{oracleDbAzureConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbAzureConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureConnector/GetOracleDbAzureConnector"
		err = common.PostProcessServiceError(err, "OracleDBAzureConnector", "GetOracleDbAzureConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbAzureConnectors Lists the all Oracle DB Azure Connector Resource based on filters.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureConnectors.go.html to see an example of how to use ListOracleDbAzureConnectors API.
// A default retry strategy applies to this operation ListOracleDbAzureConnectors()
func (client OracleDBAzureConnectorClient) ListOracleDbAzureConnectors(ctx context.Context, request ListOracleDbAzureConnectorsRequest) (response ListOracleDbAzureConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbAzureConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbAzureConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbAzureConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbAzureConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbAzureConnectorsResponse")
	}
	return
}

// listOracleDbAzureConnectors implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureConnectorClient) listOracleDbAzureConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureConnector", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbAzureConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureConnector/ListOracleDbAzureConnectors"
		err = common.PostProcessServiceError(err, "OracleDBAzureConnector", "ListOracleDbAzureConnectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchOracleDbAzureConnector Patch Azure Arc Agent on VM Cluster with new version.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/PatchOracleDbAzureConnector.go.html to see an example of how to use PatchOracleDbAzureConnector API.
// A default retry strategy applies to this operation PatchOracleDbAzureConnector()
func (client OracleDBAzureConnectorClient) PatchOracleDbAzureConnector(ctx context.Context, request PatchOracleDbAzureConnectorRequest) (response PatchOracleDbAzureConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchOracleDbAzureConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchOracleDbAzureConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchOracleDbAzureConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchOracleDbAzureConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchOracleDbAzureConnectorResponse")
	}
	return
}

// patchOracleDbAzureConnector implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureConnectorClient) patchOracleDbAzureConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/oracleDbAzureConnector/{oracleDbAzureConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchOracleDbAzureConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureConnector/PatchOracleDbAzureConnector"
		err = common.PostProcessServiceError(err, "OracleDBAzureConnector", "PatchOracleDbAzureConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbAzureConnector Modifies the existing Oracle DB Azure Connector Resource for a given ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/UpdateOracleDbAzureConnector.go.html to see an example of how to use UpdateOracleDbAzureConnector API.
// A default retry strategy applies to this operation UpdateOracleDbAzureConnector()
func (client OracleDBAzureConnectorClient) UpdateOracleDbAzureConnector(ctx context.Context, request UpdateOracleDbAzureConnectorRequest) (response UpdateOracleDbAzureConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbAzureConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbAzureConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbAzureConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbAzureConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbAzureConnectorResponse")
	}
	return
}

// updateOracleDbAzureConnector implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureConnectorClient) updateOracleDbAzureConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbAzureConnector/{oracleDbAzureConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbAzureConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureConnector/UpdateOracleDbAzureConnector"
		err = common.PostProcessServiceError(err, "OracleDBAzureConnector", "UpdateOracleDbAzureConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
