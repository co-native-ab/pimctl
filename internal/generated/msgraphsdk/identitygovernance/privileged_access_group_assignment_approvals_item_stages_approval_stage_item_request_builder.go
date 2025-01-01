package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
    i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0 "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models/odataerrors"
)

// PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder provides operations to manage the stages property of the microsoft.graph.approval entity.
type PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderGetQueryParameters a collection of stages in the approval decision.
type PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderInternal instantiates a new PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) {
    m := &PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentApprovals/{approval%2Did}/stages/{approvalStage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder instantiates a new PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property stages for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of stages in the approval decision.
// returns a ApprovalStageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderGetRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateApprovalStageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageable), nil
}
// Patch update the navigation property stages in identityGovernance
// returns a ApprovalStageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) Patch(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageable, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderPatchRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateApprovalStageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageable), nil
}
// ToDeleteRequestInformation delete navigation property stages for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of stages in the approval decision.
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property stages in identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageable, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
