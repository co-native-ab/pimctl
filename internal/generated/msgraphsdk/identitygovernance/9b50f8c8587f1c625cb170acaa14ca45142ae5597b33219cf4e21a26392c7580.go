package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
    i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0 "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models/odataerrors"
)

// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder provides operations to manage the eligibilitySchedules property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetQueryParameters read the properties and relationships of a privilegedAccessGroupEligibilitySchedule object.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal instantiates a new PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    m := &PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilitySchedules/{privilegedAccessGroupEligibilitySchedule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder instantiates a new PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property eligibilitySchedules for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a privilegedAccessGroupEligibilitySchedule object.
// returns a PrivilegedAccessGroupEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupeligibilityschedule-get?view=graph-rest-1.0
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable), nil
}
// Patch update the navigation property eligibilitySchedules in identityGovernance
// returns a PrivilegedAccessGroupEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) Patch(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i38ea8ad6c3d53ce4417f2096109e2dec471b474c5b7ac95fe447862225d89bc0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable), nil
}
// ToDeleteRequestInformation delete navigation property eligibilitySchedules for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a privilegedAccessGroupEligibilitySchedule object.
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property eligibilitySchedules in identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable, requestConfiguration *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder when successful
func (m *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
