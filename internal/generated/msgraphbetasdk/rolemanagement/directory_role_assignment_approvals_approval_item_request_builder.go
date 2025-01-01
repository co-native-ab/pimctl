package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc "github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk/models"
    i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97 "github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk/models/odataerrors"
)

// DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder provides operations to manage the roleAssignmentApprovals property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderGetQueryParameters get roleAssignmentApprovals from roleManagement
type DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderGetQueryParameters
}
// DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderInternal instantiates a new DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) {
    m := &DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentApprovals/{approval%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder instantiates a new DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentApprovals for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get roleAssignmentApprovals from roleManagement
// returns a Approvalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderGetRequestConfiguration)(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.CreateApprovalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable), nil
}
// Patch update the navigation property roleAssignmentApprovals in roleManagement
// returns a Approvalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) Patch(ctx context.Context, body ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable, requestConfiguration *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderPatchRequestConfiguration)(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.CreateApprovalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable), nil
}
// Steps provides operations to manage the steps property of the microsoft.graph.approval entity.
// returns a *DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) Steps()(*DirectoryRoleAssignmentApprovalsItemStepsRequestBuilder) {
    return NewDirectoryRoleAssignmentApprovalsItemStepsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentApprovals for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get roleAssignmentApprovals from roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property roleAssignmentApprovals in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable, requestConfiguration *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) {
    return NewDirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
