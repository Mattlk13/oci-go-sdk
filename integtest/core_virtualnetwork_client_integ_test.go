// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package integtest

/*
import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/core"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testRegionForVirtualNetwork = common.REGION_PHX
)

func TestVirtualNetworkClient_CreateCpe(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateCpeRequest{}
	r, err := c.CreateCpe(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateCrossConnect(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateCrossConnectRequest{}
	r, err := c.CreateCrossConnect(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateCrossConnectGroup(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateCrossConnectGroupRequest{}
	r, err := c.CreateCrossConnectGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateDhcpOptions(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateDhcpOptionsRequest{}
	r, err := c.CreateDhcpOptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateDrg(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateDrgRequest{}
	r, err := c.CreateDrg(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateDrgAttachment(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateDrgAttachmentRequest{}
	r, err := c.CreateDrgAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateIPSecConnection(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateIPSecConnectionRequest{}
	r, err := c.CreateIPSecConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateInternetGateway(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateInternetGatewayRequest{}
	r, err := c.CreateInternetGateway(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreatePrivateIp(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreatePrivateIpRequest{}
	r, err := c.CreatePrivateIp(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateRouteTable(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateRouteTableRequest{}
	r, err := c.CreateRouteTable(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateSecurityList(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateSecurityListRequest{}
	r, err := c.CreateSecurityList(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateSubnet(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateSubnetRequest{}
	r, err := c.CreateSubnet(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateVcn(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateVcnRequest{}
	r, err := c.CreateVcn(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateVirtualCircuit(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.CreateVirtualCircuitRequest{}
	r, err := c.CreateVirtualCircuit(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteCpe(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteCpeRequest{}
	err := c.DeleteCpe(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteCrossConnect(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteCrossConnectRequest{}
	err := c.DeleteCrossConnect(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteCrossConnectGroup(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteCrossConnectGroupRequest{}
	err := c.DeleteCrossConnectGroup(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteDhcpOptions(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteDhcpOptionsRequest{}
	err := c.DeleteDhcpOptions(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteDrg(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteDrgRequest{}
	err := c.DeleteDrg(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteDrgAttachment(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteDrgAttachmentRequest{}
	err := c.DeleteDrgAttachment(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteIPSecConnection(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteIPSecConnectionRequest{}
	err := c.DeleteIPSecConnection(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteInternetGateway(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteInternetGatewayRequest{}
	err := c.DeleteInternetGateway(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeletePrivateIp(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeletePrivateIpRequest{}
	err := c.DeletePrivateIp(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteRouteTable(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteRouteTableRequest{}
	err := c.DeleteRouteTable(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteSecurityList(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteSecurityListRequest{}
	err := c.DeleteSecurityList(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteSubnet(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteSubnetRequest{}
	err := c.DeleteSubnet(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteVcn(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteVcnRequest{}
	err := c.DeleteVcn(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteVirtualCircuit(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.DeleteVirtualCircuitRequest{}
	err := c.DeleteVirtualCircuit(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetCpe(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetCpeRequest{}
	r, err := c.GetCpe(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetCrossConnect(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetCrossConnectRequest{}
	r, err := c.GetCrossConnect(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetCrossConnectGroup(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetCrossConnectGroupRequest{}
	r, err := c.GetCrossConnectGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetCrossConnectLetterOfAuthority(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetCrossConnectLetterOfAuthorityRequest{}
	r, err := c.GetCrossConnectLetterOfAuthority(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetCrossConnectStatus(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetCrossConnectStatusRequest{}
	r, err := c.GetCrossConnectStatus(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetDhcpOptions(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetDhcpOptionsRequest{}
	r, err := c.GetDhcpOptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetDrg(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetDrgRequest{}
	r, err := c.GetDrg(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetDrgAttachment(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetDrgAttachmentRequest{}
	r, err := c.GetDrgAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetFastConnectProviderService(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetFastConnectProviderServiceRequest{}
	r, err := c.GetFastConnectProviderService(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetIPSecConnection(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetIPSecConnectionRequest{}
	r, err := c.GetIPSecConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetIPSecConnectionDeviceConfig(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetIPSecConnectionDeviceConfigRequest{}
	r, err := c.GetIPSecConnectionDeviceConfig(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetIPSecConnectionDeviceStatus(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetIPSecConnectionDeviceStatusRequest{}
	r, err := c.GetIPSecConnectionDeviceStatus(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetInternetGateway(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetInternetGatewayRequest{}
	r, err := c.GetInternetGateway(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetPrivateIp(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetPrivateIpRequest{}
	r, err := c.GetPrivateIp(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetRouteTable(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetRouteTableRequest{}
	r, err := c.GetRouteTable(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetSecurityList(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetSecurityListRequest{}
	r, err := c.GetSecurityList(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetSubnet(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetSubnetRequest{}
	r, err := c.GetSubnet(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetVcn(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetVcnRequest{}
	r, err := c.GetVcn(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetVirtualCircuit(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetVirtualCircuitRequest{}
	r, err := c.GetVirtualCircuit(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetVnic(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.GetVnicRequest{}
	r, err := c.GetVnic(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListCpes(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListCpesRequest{}
	r, err := c.ListCpes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListCrossConnectGroups(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListCrossConnectGroupsRequest{}
	r, err := c.ListCrossConnectGroups(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListCrossConnectLocations(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListCrossConnectLocationsRequest{}
	r, err := c.ListCrossConnectLocations(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListCrossConnects(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListCrossConnectsRequest{}
	r, err := c.ListCrossConnects(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListCrossconnectPortSpeedShapes(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListCrossconnectPortSpeedShapesRequest{}
	r, err := c.ListCrossconnectPortSpeedShapes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListDhcpOptions(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListDhcpOptionsRequest{}
	r, err := c.ListDhcpOptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListDrgAttachments(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListDrgAttachmentsRequest{}
	r, err := c.ListDrgAttachments(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListDrgs(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListDrgsRequest{}
	r, err := c.ListDrgs(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListFastConnectProviderServices(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListFastConnectProviderServicesRequest{}
	r, err := c.ListFastConnectProviderServices(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListFastConnectProviderVirtualCircuitBandwidthShapes(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListFastConnectProviderVirtualCircuitBandwidthShapesRequest{}
	r, err := c.ListFastConnectProviderVirtualCircuitBandwidthShapes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListIPSecConnections(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListIPSecConnectionsRequest{}
	r, err := c.ListIPSecConnections(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListInternetGateways(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListInternetGatewaysRequest{}
	r, err := c.ListInternetGateways(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListPrivateIps(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListPrivateIpsRequest{}
	r, err := c.ListPrivateIps(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListRouteTables(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListRouteTablesRequest{}
	r, err := c.ListRouteTables(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListSecurityLists(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListSecurityListsRequest{}
	r, err := c.ListSecurityLists(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListSubnets(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListSubnetsRequest{}
	r, err := c.ListSubnets(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListVcns(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListVcnsRequest{}
	r, err := c.ListVcns(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListVirtualCircuitBandwidthShapes(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListVirtualCircuitBandwidthShapesRequest{}
	r, err := c.ListVirtualCircuitBandwidthShapes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListVirtualCircuits(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.ListVirtualCircuitsRequest{}
	r, err := c.ListVirtualCircuits(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateCpe(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateCpeRequest{}
	r, err := c.UpdateCpe(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateCrossConnect(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateCrossConnectRequest{}
	r, err := c.UpdateCrossConnect(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateCrossConnectGroup(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateCrossConnectGroupRequest{}
	r, err := c.UpdateCrossConnectGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateDhcpOptions(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateDhcpOptionsRequest{}
	r, err := c.UpdateDhcpOptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateDrg(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateDrgRequest{}
	r, err := c.UpdateDrg(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateDrgAttachment(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateDrgAttachmentRequest{}
	r, err := c.UpdateDrgAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateIPSecConnection(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateIPSecConnectionRequest{}
	r, err := c.UpdateIPSecConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateInternetGateway(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateInternetGatewayRequest{}
	r, err := c.UpdateInternetGateway(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdatePrivateIp(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdatePrivateIpRequest{}
	r, err := c.UpdatePrivateIp(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateRouteTable(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateRouteTableRequest{}
	r, err := c.UpdateRouteTable(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateSecurityList(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateSecurityListRequest{}
	r, err := c.UpdateSecurityList(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateSubnet(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateSubnetRequest{}
	r, err := c.UpdateSubnet(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateVcn(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateVcnRequest{}
	r, err := c.UpdateVcn(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateVirtualCircuit(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateVirtualCircuitRequest{}
	r, err := c.UpdateVirtualCircuit(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateVnic(t *testing.T) {
	c := core.NewVirtualNetworkClientForRegion(testRegionForVirtualNetwork)
	request := core.UpdateVnicRequest{}
	r, err := c.UpdateVnic(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}
*/
