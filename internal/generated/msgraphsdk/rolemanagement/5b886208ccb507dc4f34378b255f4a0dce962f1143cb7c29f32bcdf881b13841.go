package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
    i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0 "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models/odataerrors"
)

// DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters get the instance of an active role assignment.
type DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters
}
// DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal instantiates a new DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    m := &DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentScheduleInstances/{unifiedRoleAssignmentScheduleInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder instantiates a new DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentScheduleInstances for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the instance of an active role assignment.
// returns a UnifiedRoleAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentscheduleinstance-get?view=graph-rest-1.0
func (m *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleInstanceable), nil
}
// Patch update the navigation property roleAssignmentScheduleInstances in roleManagement
// returns a UnifiedRoleAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Patch(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleInstanceable, requestConfiguration *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleInstanceable), nil
}
// ToDeleteRequestInformation delete navigation property roleAssignmentScheduleInstances for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the instance of an active role assignment.
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignmentScheduleInstances in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleInstanceable, requestConfiguration *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder when successful
func (m *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
