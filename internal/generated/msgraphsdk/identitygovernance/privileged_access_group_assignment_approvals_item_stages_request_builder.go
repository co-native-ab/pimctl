package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
    i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0 "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models/odataerrors"
)

// PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder provides operations to manage the stages property of the microsoft.graph.approval entity.
type PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderGetQueryParameters a collection of stages in the approval decision.
type PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByApprovalStageId provides operations to manage the stages property of the microsoft.graph.approval entity.
// returns a *PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) ByApprovalStageId(approvalStageId string)(*PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if approvalStageId != "" {
        urlTplParams["approvalStage%2Did"] = approvalStageId
    }
    return NewPrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderInternal instantiates a new PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) {
    m := &PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentApprovals/{approval%2Did}/stages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder instantiates a new PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegedAccessGroupAssignmentApprovalsItemStagesCountRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) Count()(*PrivilegedAccessGroupAssignmentApprovalsItemStagesCountRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentApprovalsItemStagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of stages in the approval decision.
// returns a ApprovalStageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderGetRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateApprovalStageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageCollectionResponseable), nil
}
// Post create new navigation property to stages for identityGovernance
// returns a ApprovalStageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) Post(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageable, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderPostRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation a collection of stages in the approval decision.
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to stages for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ApprovalStageable, requestConfiguration *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentApprovalsItemStagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
