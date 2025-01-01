package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder builds and executes requests for operations under \identityGovernance\privilegedAccess\group\assignmentScheduleInstances
type PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByPrivilegedAccessGroupAssignmentScheduleInstanceId provides operations to manage the assignmentScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder) ByPrivilegedAccessGroupAssignmentScheduleInstanceId(privilegedAccessGroupAssignmentScheduleInstanceId string)(*PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if privilegedAccessGroupAssignmentScheduleInstanceId != "" {
        urlTplParams["privilegedAccessGroupAssignmentScheduleInstance%2Did"] = privilegedAccessGroupAssignmentScheduleInstanceId
    }
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilderInternal instantiates a new PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder) {
    m := &PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleInstances", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder instantiates a new PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegedAccessGroupAssignmentScheduleInstancesCountRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder) Count()(*PrivilegedAccessGroupAssignmentScheduleInstancesCountRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *PrivilegedAccessGroupAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*PrivilegedAccessGroupAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
