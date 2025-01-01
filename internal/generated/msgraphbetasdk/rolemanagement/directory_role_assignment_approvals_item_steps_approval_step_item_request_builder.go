package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc "github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk/models"
    i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97 "github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk/models/odataerrors"
)

// DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder provides operations to manage the steps property of the microsoft.graph.approval entity.
type DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderGetQueryParameters used to represent the decision associated with a single step in the approval process configured in approvalStage.
type DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderGetQueryParameters
}
// DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderInternal instantiates a new DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) {
    m := &DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentApprovals/{approval%2Did}/steps/{approvalStep%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder instantiates a new DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property steps for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get used to represent the decision associated with a single step in the approval process configured in approvalStage.
// returns a ApprovalStepable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderGetRequestConfiguration)(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.CreateApprovalStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable), nil
}
// Patch update the navigation property steps in roleManagement
// returns a ApprovalStepable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) Patch(ctx context.Context, body ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderPatchRequestConfiguration)(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.CreateApprovalStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable), nil
}
// ToDeleteRequestInformation delete navigation property steps for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation used to represent the decision associated with a single step in the approval process configured in approvalStage.
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property steps in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.ApprovalStepable, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder) {
    return NewDirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
