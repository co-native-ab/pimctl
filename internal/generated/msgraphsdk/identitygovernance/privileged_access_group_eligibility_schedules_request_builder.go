package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrivilegedAccessGroupEligibilitySchedulesRequestBuilder builds and executes requests for operations under \identityGovernance\privilegedAccess\group\eligibilitySchedules
type PrivilegedAccessGroupEligibilitySchedulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByPrivilegedAccessGroupEligibilityScheduleId provides operations to manage the eligibilitySchedules property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder when successful
func (m *PrivilegedAccessGroupEligibilitySchedulesRequestBuilder) ByPrivilegedAccessGroupEligibilityScheduleId(privilegedAccessGroupEligibilityScheduleId string)(*PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if privilegedAccessGroupEligibilityScheduleId != "" {
        urlTplParams["privilegedAccessGroupEligibilitySchedule%2Did"] = privilegedAccessGroupEligibilityScheduleId
    }
    return NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupEligibilitySchedulesRequestBuilderInternal instantiates a new PrivilegedAccessGroupEligibilitySchedulesRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilitySchedulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilitySchedulesRequestBuilder) {
    m := &PrivilegedAccessGroupEligibilitySchedulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilitySchedules", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupEligibilitySchedulesRequestBuilder instantiates a new PrivilegedAccessGroupEligibilitySchedulesRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilitySchedulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilitySchedulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupEligibilitySchedulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegedAccessGroupEligibilitySchedulesCountRequestBuilder when successful
func (m *PrivilegedAccessGroupEligibilitySchedulesRequestBuilder) Count()(*PrivilegedAccessGroupEligibilitySchedulesCountRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilitySchedulesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder when successful
func (m *PrivilegedAccessGroupEligibilitySchedulesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
