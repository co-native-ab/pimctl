package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97 "github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk/models/odataerrors"
)

// DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder provides operations to count the resources in the collection.
type DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderGetQueryParameters get the number of the resource
type DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderGetQueryParameters
}
// NewDirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderInternal instantiates a new DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder) {
    m := &DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentApprovals/{approval%2Did}/steps/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewDirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder instantiates a new DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i3c348b42299dea02992e75ce229fbe66a0442ea71b3fb1c422dfbe0720f96f97.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder) {
    return NewDirectoryRoleAssignmentApprovalsItemStepsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
