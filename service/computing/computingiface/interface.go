// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package computingiface provides an interface to enable mocking the NIFCLOUD Computing service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package computingiface

import (
	"github.com/alice02/nifcloud-sdk-go-v2/service/computing"
)

// ComputingAPI provides an interface to enable mocking the
// computing.Computing service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // NIFCLOUD Computing.
//    func myFunc(svc computingiface.ComputingAPI) bool {
//        // Make svc.AllocateAddress request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := computing.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockComputingClient struct {
//        computingiface.ComputingAPI
//    }
//    func (m *mockComputingClient) AllocateAddress(input *computing.AllocateAddressInput) (*computing.AllocateAddressOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockComputingClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ComputingAPI interface {
	AllocateAddressRequest(*computing.AllocateAddressInput) computing.AllocateAddressRequest

	AssociateAddressRequest(*computing.AssociateAddressInput) computing.AssociateAddressRequest

	AssociateRouteTableRequest(*computing.AssociateRouteTableInput) computing.AssociateRouteTableRequest

	AssociateUsersRequest(*computing.AssociateUsersInput) computing.AssociateUsersRequest

	AttachNetworkInterfaceRequest(*computing.AttachNetworkInterfaceInput) computing.AttachNetworkInterfaceRequest

	AttachVolumeRequest(*computing.AttachVolumeInput) computing.AttachVolumeRequest

	AuthorizeSecurityGroupIngressRequest(*computing.AuthorizeSecurityGroupIngressInput) computing.AuthorizeSecurityGroupIngressRequest

	CancelCopyInstancesRequest(*computing.CancelCopyInstancesInput) computing.CancelCopyInstancesRequest

	CancelUploadRequest(*computing.CancelUploadInput) computing.CancelUploadRequest

	ClearLoadBalancerSessionRequest(*computing.ClearLoadBalancerSessionInput) computing.ClearLoadBalancerSessionRequest

	ConfigureHealthCheckRequest(*computing.ConfigureHealthCheckInput) computing.ConfigureHealthCheckRequest

	CopyInstancesRequest(*computing.CopyInstancesInput) computing.CopyInstancesRequest

	CreateCustomerGatewayRequest(*computing.CreateCustomerGatewayInput) computing.CreateCustomerGatewayRequest

	CreateDhcpOptionsRequest(*computing.CreateDhcpOptionsInput) computing.CreateDhcpOptionsRequest

	CreateImageRequest(*computing.CreateImageInput) computing.CreateImageRequest

	CreateKeyPairRequest(*computing.CreateKeyPairInput) computing.CreateKeyPairRequest

	CreateLoadBalancerRequest(*computing.CreateLoadBalancerInput) computing.CreateLoadBalancerRequest

	CreateNetworkInterfaceRequest(*computing.CreateNetworkInterfaceInput) computing.CreateNetworkInterfaceRequest

	CreateRouteRequest(*computing.CreateRouteInput) computing.CreateRouteRequest

	CreateRouteTableRequest(*computing.CreateRouteTableInput) computing.CreateRouteTableRequest

	CreateSecurityGroupRequest(*computing.CreateSecurityGroupInput) computing.CreateSecurityGroupRequest

	CreateSslCertificateRequest(*computing.CreateSslCertificateInput) computing.CreateSslCertificateRequest

	CreateVolumeRequest(*computing.CreateVolumeInput) computing.CreateVolumeRequest

	CreateVpnConnectionRequest(*computing.CreateVpnConnectionInput) computing.CreateVpnConnectionRequest

	CreateVpnGatewayRequest(*computing.CreateVpnGatewayInput) computing.CreateVpnGatewayRequest

	DeleteCustomerGatewayRequest(*computing.DeleteCustomerGatewayInput) computing.DeleteCustomerGatewayRequest

	DeleteDhcpOptionsRequest(*computing.DeleteDhcpOptionsInput) computing.DeleteDhcpOptionsRequest

	DeleteImageRequest(*computing.DeleteImageInput) computing.DeleteImageRequest

	DeleteKeyPairRequest(*computing.DeleteKeyPairInput) computing.DeleteKeyPairRequest

	DeleteLoadBalancerRequest(*computing.DeleteLoadBalancerInput) computing.DeleteLoadBalancerRequest

	DeleteNetworkInterfaceRequest(*computing.DeleteNetworkInterfaceInput) computing.DeleteNetworkInterfaceRequest

	DeleteRouteRequest(*computing.DeleteRouteInput) computing.DeleteRouteRequest

	DeleteRouteTableRequest(*computing.DeleteRouteTableInput) computing.DeleteRouteTableRequest

	DeleteSecurityGroupRequest(*computing.DeleteSecurityGroupInput) computing.DeleteSecurityGroupRequest

	DeleteSslCertificateRequest(*computing.DeleteSslCertificateInput) computing.DeleteSslCertificateRequest

	DeleteVolumeRequest(*computing.DeleteVolumeInput) computing.DeleteVolumeRequest

	DeleteVpnConnectionRequest(*computing.DeleteVpnConnectionInput) computing.DeleteVpnConnectionRequest

	DeleteVpnGatewayRequest(*computing.DeleteVpnGatewayInput) computing.DeleteVpnGatewayRequest

	DeregisterInstancesFromLoadBalancerRequest(*computing.DeregisterInstancesFromLoadBalancerInput) computing.DeregisterInstancesFromLoadBalancerRequest

	DeregisterInstancesFromSecurityGroupRequest(*computing.DeregisterInstancesFromSecurityGroupInput) computing.DeregisterInstancesFromSecurityGroupRequest

	DescribeAddressesRequest(*computing.DescribeAddressesInput) computing.DescribeAddressesRequest

	DescribeAssociatedUsersRequest(*computing.DescribeAssociatedUsersInput) computing.DescribeAssociatedUsersRequest

	DescribeAvailabilityZonesRequest(*computing.DescribeAvailabilityZonesInput) computing.DescribeAvailabilityZonesRequest

	DescribeCustomerGatewaysRequest(*computing.DescribeCustomerGatewaysInput) computing.DescribeCustomerGatewaysRequest

	DescribeDhcpOptionsRequest(*computing.DescribeDhcpOptionsInput) computing.DescribeDhcpOptionsRequest

	DescribeImagesRequest(*computing.DescribeImagesInput) computing.DescribeImagesRequest

	DescribeInstanceAttributeRequest(*computing.DescribeInstanceAttributeInput) computing.DescribeInstanceAttributeRequest

	DescribeInstanceHealthRequest(*computing.DescribeInstanceHealthInput) computing.DescribeInstanceHealthRequest

	DescribeInstancesRequest(*computing.DescribeInstancesInput) computing.DescribeInstancesRequest

	DescribeKeyPairsRequest(*computing.DescribeKeyPairsInput) computing.DescribeKeyPairsRequest

	DescribeLoadBalancersRequest(*computing.DescribeLoadBalancersInput) computing.DescribeLoadBalancersRequest

	DescribeNetworkInterfacesRequest(*computing.DescribeNetworkInterfacesInput) computing.DescribeNetworkInterfacesRequest

	DescribeRegionsRequest(*computing.DescribeRegionsInput) computing.DescribeRegionsRequest

	DescribeResourcesRequest(*computing.DescribeResourcesInput) computing.DescribeResourcesRequest

	DescribeRouteTablesRequest(*computing.DescribeRouteTablesInput) computing.DescribeRouteTablesRequest

	DescribeSecurityActivitiesRequest(*computing.DescribeSecurityActivitiesInput) computing.DescribeSecurityActivitiesRequest

	DescribeSecurityGroupOptionRequest(*computing.DescribeSecurityGroupOptionInput) computing.DescribeSecurityGroupOptionRequest

	DescribeSecurityGroupsRequest(*computing.DescribeSecurityGroupsInput) computing.DescribeSecurityGroupsRequest

	DescribeServiceStatusRequest(*computing.DescribeServiceStatusInput) computing.DescribeServiceStatusRequest

	DescribeSslCertificateAttributeRequest(*computing.DescribeSslCertificateAttributeInput) computing.DescribeSslCertificateAttributeRequest

	DescribeSslCertificatesRequest(*computing.DescribeSslCertificatesInput) computing.DescribeSslCertificatesRequest

	DescribeUploadsRequest(*computing.DescribeUploadsInput) computing.DescribeUploadsRequest

	DescribeUsageRequest(*computing.DescribeUsageInput) computing.DescribeUsageRequest

	DescribeUserActivitiesRequest(*computing.DescribeUserActivitiesInput) computing.DescribeUserActivitiesRequest

	DescribeVolumesRequest(*computing.DescribeVolumesInput) computing.DescribeVolumesRequest

	DescribeVpnConnectionsRequest(*computing.DescribeVpnConnectionsInput) computing.DescribeVpnConnectionsRequest

	DescribeVpnGatewaysRequest(*computing.DescribeVpnGatewaysInput) computing.DescribeVpnGatewaysRequest

	DetachNetworkInterfaceRequest(*computing.DetachNetworkInterfaceInput) computing.DetachNetworkInterfaceRequest

	DetachVolumeRequest(*computing.DetachVolumeInput) computing.DetachVolumeRequest

	DisassociateAddressRequest(*computing.DisassociateAddressInput) computing.DisassociateAddressRequest

	DisassociateRouteTableRequest(*computing.DisassociateRouteTableInput) computing.DisassociateRouteTableRequest

	DissociateUsersRequest(*computing.DissociateUsersInput) computing.DissociateUsersRequest

	DownloadSslCertificateRequest(*computing.DownloadSslCertificateInput) computing.DownloadSslCertificateRequest

	ImportInstanceRequest(*computing.ImportInstanceInput) computing.ImportInstanceRequest

	ImportKeyPairRequest(*computing.ImportKeyPairInput) computing.ImportKeyPairRequest

	ModifyImageAttributeRequest(*computing.ModifyImageAttributeInput) computing.ModifyImageAttributeRequest

	ModifyInstanceAttributeRequest(*computing.ModifyInstanceAttributeInput) computing.ModifyInstanceAttributeRequest

	ModifyNetworkInterfaceAttributeRequest(*computing.ModifyNetworkInterfaceAttributeInput) computing.ModifyNetworkInterfaceAttributeRequest

	ModifySslCertificateAttributeRequest(*computing.ModifySslCertificateAttributeInput) computing.ModifySslCertificateAttributeRequest

	ModifyVolumeAttributeRequest(*computing.ModifyVolumeAttributeInput) computing.ModifyVolumeAttributeRequest

	NiftyAssociateImageRequest(*computing.NiftyAssociateImageInput) computing.NiftyAssociateImageRequest

	NiftyAssociateNatTableRequest(*computing.NiftyAssociateNatTableInput) computing.NiftyAssociateNatTableRequest

	NiftyAssociateRouteTableWithVpnGatewayRequest(*computing.NiftyAssociateRouteTableWithVpnGatewayInput) computing.NiftyAssociateRouteTableWithVpnGatewayRequest

	NiftyConfigureElasticLoadBalancerHealthCheckRequest(*computing.NiftyConfigureElasticLoadBalancerHealthCheckInput) computing.NiftyConfigureElasticLoadBalancerHealthCheckRequest

	NiftyCreateAlarmRequest(*computing.NiftyCreateAlarmInput) computing.NiftyCreateAlarmRequest

	NiftyCreateAutoScalingGroupRequest(*computing.NiftyCreateAutoScalingGroupInput) computing.NiftyCreateAutoScalingGroupRequest

	NiftyCreateDhcpConfigRequest(*computing.NiftyCreateDhcpConfigInput) computing.NiftyCreateDhcpConfigRequest

	NiftyCreateDhcpIpAddressPoolRequest(*computing.NiftyCreateDhcpIpAddressPoolInput) computing.NiftyCreateDhcpIpAddressPoolRequest

	NiftyCreateDhcpStaticMappingRequest(*computing.NiftyCreateDhcpStaticMappingInput) computing.NiftyCreateDhcpStaticMappingRequest

	NiftyCreateElasticLoadBalancerRequest(*computing.NiftyCreateElasticLoadBalancerInput) computing.NiftyCreateElasticLoadBalancerRequest

	NiftyCreateInstanceSnapshotRequest(*computing.NiftyCreateInstanceSnapshotInput) computing.NiftyCreateInstanceSnapshotRequest

	NiftyCreateNatRuleRequest(*computing.NiftyCreateNatRuleInput) computing.NiftyCreateNatRuleRequest

	NiftyCreateNatTableRequest(*computing.NiftyCreateNatTableInput) computing.NiftyCreateNatTableRequest

	NiftyCreatePrivateLanRequest(*computing.NiftyCreatePrivateLanInput) computing.NiftyCreatePrivateLanRequest

	NiftyCreateRouterRequest(*computing.NiftyCreateRouterInput) computing.NiftyCreateRouterRequest

	NiftyCreateSeparateInstanceRuleRequest(*computing.NiftyCreateSeparateInstanceRuleInput) computing.NiftyCreateSeparateInstanceRuleRequest

	NiftyCreateWebProxyRequest(*computing.NiftyCreateWebProxyInput) computing.NiftyCreateWebProxyRequest

	NiftyDeleteAlarmRequest(*computing.NiftyDeleteAlarmInput) computing.NiftyDeleteAlarmRequest

	NiftyDeleteAutoScalingGroupRequest(*computing.NiftyDeleteAutoScalingGroupInput) computing.NiftyDeleteAutoScalingGroupRequest

	NiftyDeleteDhcpConfigRequest(*computing.NiftyDeleteDhcpConfigInput) computing.NiftyDeleteDhcpConfigRequest

	NiftyDeleteDhcpIpAddressPoolRequest(*computing.NiftyDeleteDhcpIpAddressPoolInput) computing.NiftyDeleteDhcpIpAddressPoolRequest

	NiftyDeleteDhcpStaticMappingRequest(*computing.NiftyDeleteDhcpStaticMappingInput) computing.NiftyDeleteDhcpStaticMappingRequest

	NiftyDeleteElasticLoadBalancerRequest(*computing.NiftyDeleteElasticLoadBalancerInput) computing.NiftyDeleteElasticLoadBalancerRequest

	NiftyDeleteInstanceSnapshotRequest(*computing.NiftyDeleteInstanceSnapshotInput) computing.NiftyDeleteInstanceSnapshotRequest

	NiftyDeleteNatRuleRequest(*computing.NiftyDeleteNatRuleInput) computing.NiftyDeleteNatRuleRequest

	NiftyDeleteNatTableRequest(*computing.NiftyDeleteNatTableInput) computing.NiftyDeleteNatTableRequest

	NiftyDeletePrivateLanRequest(*computing.NiftyDeletePrivateLanInput) computing.NiftyDeletePrivateLanRequest

	NiftyDeleteRouterRequest(*computing.NiftyDeleteRouterInput) computing.NiftyDeleteRouterRequest

	NiftyDeleteSeparateInstanceRuleRequest(*computing.NiftyDeleteSeparateInstanceRuleInput) computing.NiftyDeleteSeparateInstanceRuleRequest

	NiftyDeleteWebProxyRequest(*computing.NiftyDeleteWebProxyInput) computing.NiftyDeleteWebProxyRequest

	NiftyDeregisterInstancesFromElasticLoadBalancerRequest(*computing.NiftyDeregisterInstancesFromElasticLoadBalancerInput) computing.NiftyDeregisterInstancesFromElasticLoadBalancerRequest

	NiftyDeregisterInstancesFromSeparateInstanceRuleRequest(*computing.NiftyDeregisterInstancesFromSeparateInstanceRuleInput) computing.NiftyDeregisterInstancesFromSeparateInstanceRuleRequest

	NiftyDeregisterRoutersFromSecurityGroupRequest(*computing.NiftyDeregisterRoutersFromSecurityGroupInput) computing.NiftyDeregisterRoutersFromSecurityGroupRequest

	NiftyDeregisterVpnGatewaysFromSecurityGroupRequest(*computing.NiftyDeregisterVpnGatewaysFromSecurityGroupInput) computing.NiftyDeregisterVpnGatewaysFromSecurityGroupRequest

	NiftyDescribeAlarmHistoryRequest(*computing.NiftyDescribeAlarmHistoryInput) computing.NiftyDescribeAlarmHistoryRequest

	NiftyDescribeAlarmRulesActivitiesRequest(*computing.NiftyDescribeAlarmRulesActivitiesInput) computing.NiftyDescribeAlarmRulesActivitiesRequest

	NiftyDescribeAlarmsRequest(*computing.NiftyDescribeAlarmsInput) computing.NiftyDescribeAlarmsRequest

	NiftyDescribeAlarmsPartitionsRequest(*computing.NiftyDescribeAlarmsPartitionsInput) computing.NiftyDescribeAlarmsPartitionsRequest

	NiftyDescribeAutoScalingGroupsRequest(*computing.NiftyDescribeAutoScalingGroupsInput) computing.NiftyDescribeAutoScalingGroupsRequest

	NiftyDescribeCorporateInfoForCertificateRequest(*computing.NiftyDescribeCorporateInfoForCertificateInput) computing.NiftyDescribeCorporateInfoForCertificateRequest

	NiftyDescribeDhcpConfigsRequest(*computing.NiftyDescribeDhcpConfigsInput) computing.NiftyDescribeDhcpConfigsRequest

	NiftyDescribeDhcpStatusRequest(*computing.NiftyDescribeDhcpStatusInput) computing.NiftyDescribeDhcpStatusRequest

	NiftyDescribeElasticLoadBalancersRequest(*computing.NiftyDescribeElasticLoadBalancersInput) computing.NiftyDescribeElasticLoadBalancersRequest

	NiftyDescribeInstanceElasticLoadBalancerHealthRequest(*computing.NiftyDescribeInstanceElasticLoadBalancerHealthInput) computing.NiftyDescribeInstanceElasticLoadBalancerHealthRequest

	NiftyDescribeInstanceSnapshotsRequest(*computing.NiftyDescribeInstanceSnapshotsInput) computing.NiftyDescribeInstanceSnapshotsRequest

	NiftyDescribeNatTablesRequest(*computing.NiftyDescribeNatTablesInput) computing.NiftyDescribeNatTablesRequest

	NiftyDescribePerformanceChartRequest(*computing.NiftyDescribePerformanceChartInput) computing.NiftyDescribePerformanceChartRequest

	NiftyDescribePrivateLansRequest(*computing.NiftyDescribePrivateLansInput) computing.NiftyDescribePrivateLansRequest

	NiftyDescribeRoutersRequest(*computing.NiftyDescribeRoutersInput) computing.NiftyDescribeRoutersRequest

	NiftyDescribeScalingActivitiesRequest(*computing.NiftyDescribeScalingActivitiesInput) computing.NiftyDescribeScalingActivitiesRequest

	NiftyDescribeSeparateInstanceRulesRequest(*computing.NiftyDescribeSeparateInstanceRulesInput) computing.NiftyDescribeSeparateInstanceRulesRequest

	NiftyDescribeVpnGatewayActivitiesRequest(*computing.NiftyDescribeVpnGatewayActivitiesInput) computing.NiftyDescribeVpnGatewayActivitiesRequest

	NiftyDescribeWebProxiesRequest(*computing.NiftyDescribeWebProxiesInput) computing.NiftyDescribeWebProxiesRequest

	NiftyDisableDhcpRequest(*computing.NiftyDisableDhcpInput) computing.NiftyDisableDhcpRequest

	NiftyDisassociateNatTableRequest(*computing.NiftyDisassociateNatTableInput) computing.NiftyDisassociateNatTableRequest

	NiftyDisassociateRouteTableFromVpnGatewayRequest(*computing.NiftyDisassociateRouteTableFromVpnGatewayInput) computing.NiftyDisassociateRouteTableFromVpnGatewayRequest

	NiftyEnableDhcpRequest(*computing.NiftyEnableDhcpInput) computing.NiftyEnableDhcpRequest

	NiftyModifyAddressAttributeRequest(*computing.NiftyModifyAddressAttributeInput) computing.NiftyModifyAddressAttributeRequest

	NiftyModifyCustomerGatewayAttributeRequest(*computing.NiftyModifyCustomerGatewayAttributeInput) computing.NiftyModifyCustomerGatewayAttributeRequest

	NiftyModifyElasticLoadBalancerAttributesRequest(*computing.NiftyModifyElasticLoadBalancerAttributesInput) computing.NiftyModifyElasticLoadBalancerAttributesRequest

	NiftyModifyInstanceSnapshotAttributeRequest(*computing.NiftyModifyInstanceSnapshotAttributeInput) computing.NiftyModifyInstanceSnapshotAttributeRequest

	NiftyModifyKeyPairAttributeRequest(*computing.NiftyModifyKeyPairAttributeInput) computing.NiftyModifyKeyPairAttributeRequest

	NiftyModifyPrivateLanAttributeRequest(*computing.NiftyModifyPrivateLanAttributeInput) computing.NiftyModifyPrivateLanAttributeRequest

	NiftyModifyRouterAttributeRequest(*computing.NiftyModifyRouterAttributeInput) computing.NiftyModifyRouterAttributeRequest

	NiftyModifyVpnGatewayAttributeRequest(*computing.NiftyModifyVpnGatewayAttributeInput) computing.NiftyModifyVpnGatewayAttributeRequest

	NiftyModifyWebProxyAttributeRequest(*computing.NiftyModifyWebProxyAttributeInput) computing.NiftyModifyWebProxyAttributeRequest

	NiftyRebootRoutersRequest(*computing.NiftyRebootRoutersInput) computing.NiftyRebootRoutersRequest

	NiftyRebootVpnGatewaysRequest(*computing.NiftyRebootVpnGatewaysInput) computing.NiftyRebootVpnGatewaysRequest

	NiftyRegisterInstancesWithElasticLoadBalancerRequest(*computing.NiftyRegisterInstancesWithElasticLoadBalancerInput) computing.NiftyRegisterInstancesWithElasticLoadBalancerRequest

	NiftyRegisterInstancesWithSeparateInstanceRuleRequest(*computing.NiftyRegisterInstancesWithSeparateInstanceRuleInput) computing.NiftyRegisterInstancesWithSeparateInstanceRuleRequest

	NiftyRegisterPortWithElasticLoadBalancerRequest(*computing.NiftyRegisterPortWithElasticLoadBalancerInput) computing.NiftyRegisterPortWithElasticLoadBalancerRequest

	NiftyRegisterRoutersWithSecurityGroupRequest(*computing.NiftyRegisterRoutersWithSecurityGroupInput) computing.NiftyRegisterRoutersWithSecurityGroupRequest

	NiftyRegisterVpnGatewaysWithSecurityGroupRequest(*computing.NiftyRegisterVpnGatewaysWithSecurityGroupInput) computing.NiftyRegisterVpnGatewaysWithSecurityGroupRequest

	NiftyReleaseRouterBackupStateRequest(*computing.NiftyReleaseRouterBackupStateInput) computing.NiftyReleaseRouterBackupStateRequest

	NiftyReleaseVpnGatewayBackupStateRequest(*computing.NiftyReleaseVpnGatewayBackupStateInput) computing.NiftyReleaseVpnGatewayBackupStateRequest

	NiftyReplaceDhcpConfigRequest(*computing.NiftyReplaceDhcpConfigInput) computing.NiftyReplaceDhcpConfigRequest

	NiftyReplaceDhcpOptionRequest(*computing.NiftyReplaceDhcpOptionInput) computing.NiftyReplaceDhcpOptionRequest

	NiftyReplaceElasticLoadBalancerLatestVersionRequest(*computing.NiftyReplaceElasticLoadBalancerLatestVersionInput) computing.NiftyReplaceElasticLoadBalancerLatestVersionRequest

	NiftyReplaceNatRuleRequest(*computing.NiftyReplaceNatRuleInput) computing.NiftyReplaceNatRuleRequest

	NiftyReplaceNatTableAssociationRequest(*computing.NiftyReplaceNatTableAssociationInput) computing.NiftyReplaceNatTableAssociationRequest

	NiftyReplaceRouteTableAssociationWithVpnGatewayRequest(*computing.NiftyReplaceRouteTableAssociationWithVpnGatewayInput) computing.NiftyReplaceRouteTableAssociationWithVpnGatewayRequest

	NiftyReplaceRouterLatestVersionRequest(*computing.NiftyReplaceRouterLatestVersionInput) computing.NiftyReplaceRouterLatestVersionRequest

	NiftyReplaceVpnGatewayLatestVersionRequest(*computing.NiftyReplaceVpnGatewayLatestVersionInput) computing.NiftyReplaceVpnGatewayLatestVersionRequest

	NiftyRestoreInstanceSnapshotRequest(*computing.NiftyRestoreInstanceSnapshotInput) computing.NiftyRestoreInstanceSnapshotRequest

	NiftyRestoreRouterPreviousVersionRequest(*computing.NiftyRestoreRouterPreviousVersionInput) computing.NiftyRestoreRouterPreviousVersionRequest

	NiftyRestoreVpnGatewayPreviousVersionRequest(*computing.NiftyRestoreVpnGatewayPreviousVersionInput) computing.NiftyRestoreVpnGatewayPreviousVersionRequest

	NiftyRetryImportInstanceRequest(*computing.NiftyRetryImportInstanceInput) computing.NiftyRetryImportInstanceRequest

	NiftyUpdateAlarmRequest(*computing.NiftyUpdateAlarmInput) computing.NiftyUpdateAlarmRequest

	NiftyUpdateAutoScalingGroupRequest(*computing.NiftyUpdateAutoScalingGroupInput) computing.NiftyUpdateAutoScalingGroupRequest

	NiftyUpdateElasticLoadBalancerRequest(*computing.NiftyUpdateElasticLoadBalancerInput) computing.NiftyUpdateElasticLoadBalancerRequest

	NiftyUpdateInstanceNetworkInterfacesRequest(*computing.NiftyUpdateInstanceNetworkInterfacesInput) computing.NiftyUpdateInstanceNetworkInterfacesRequest

	NiftyUpdateRouterNetworkInterfacesRequest(*computing.NiftyUpdateRouterNetworkInterfacesInput) computing.NiftyUpdateRouterNetworkInterfacesRequest

	NiftyUpdateSeparateInstanceRuleRequest(*computing.NiftyUpdateSeparateInstanceRuleInput) computing.NiftyUpdateSeparateInstanceRuleRequest

	NiftyUpdateVpnGatewayNetworkInterfacesRequest(*computing.NiftyUpdateVpnGatewayNetworkInterfacesInput) computing.NiftyUpdateVpnGatewayNetworkInterfacesRequest

	RebootInstancesRequest(*computing.RebootInstancesInput) computing.RebootInstancesRequest

	RegisterCorporateInfoForCertificateRequest(*computing.RegisterCorporateInfoForCertificateInput) computing.RegisterCorporateInfoForCertificateRequest

	RegisterInstancesWithLoadBalancerRequest(*computing.RegisterInstancesWithLoadBalancerInput) computing.RegisterInstancesWithLoadBalancerRequest

	RegisterInstancesWithSecurityGroupRequest(*computing.RegisterInstancesWithSecurityGroupInput) computing.RegisterInstancesWithSecurityGroupRequest

	RegisterPortWithLoadBalancerRequest(*computing.RegisterPortWithLoadBalancerInput) computing.RegisterPortWithLoadBalancerRequest

	ReleaseAddressRequest(*computing.ReleaseAddressInput) computing.ReleaseAddressRequest

	ReplaceRouteRequest(*computing.ReplaceRouteInput) computing.ReplaceRouteRequest

	ReplaceRouteTableAssociationRequest(*computing.ReplaceRouteTableAssociationInput) computing.ReplaceRouteTableAssociationRequest

	RevokeSecurityGroupIngressRequest(*computing.RevokeSecurityGroupIngressInput) computing.RevokeSecurityGroupIngressRequest

	RunInstancesRequest(*computing.RunInstancesInput) computing.RunInstancesRequest

	SetFilterForLoadBalancerRequest(*computing.SetFilterForLoadBalancerInput) computing.SetFilterForLoadBalancerRequest

	StartInstancesRequest(*computing.StartInstancesInput) computing.StartInstancesRequest

	StopInstancesRequest(*computing.StopInstancesInput) computing.StopInstancesRequest

	TerminateInstancesRequest(*computing.TerminateInstancesInput) computing.TerminateInstancesRequest

	UpdateLoadBalancerRequest(*computing.UpdateLoadBalancerInput) computing.UpdateLoadBalancerRequest

	UpdateLoadBalancerOptionRequest(*computing.UpdateLoadBalancerOptionInput) computing.UpdateLoadBalancerOptionRequest

	UpdateSecurityGroupRequest(*computing.UpdateSecurityGroupInput) computing.UpdateSecurityGroupRequest

	UpdateSecurityGroupOptionRequest(*computing.UpdateSecurityGroupOptionInput) computing.UpdateSecurityGroupOptionRequest

	UploadSslCertificateRequest(*computing.UploadSslCertificateInput) computing.UploadSslCertificateRequest
}

var _ ComputingAPI = (*computing.Computing)(nil)
