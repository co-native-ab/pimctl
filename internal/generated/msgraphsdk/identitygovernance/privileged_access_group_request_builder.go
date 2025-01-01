package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrivilegedAccessGroupRequestBuilder builds and executes requests for operations under \identityGovernance\privilegedAccess\group
type PrivilegedAccessGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AssignmentApprovals the assignmentApprovals property
// returns a *PrivilegedAccessGroupAssignmentApprovalsRequestBuilder when successful
func (m *PrivilegedAccessGroupRequestBuilder) AssignmentApprovals()(*PrivilegedAccessGroupAssignmentApprovalsRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentScheduleInstances the assignmentScheduleInstances property
// returns a *PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder when successful
func (m *PrivilegedAccessGroupRequestBuilder) AssignmentScheduleInstances()(*PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentScheduleRequests provides operations to manage the assignmentScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder when successful
func (m *PrivilegedAccessGroupRequestBuilder) AssignmentScheduleRequests()(*PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupRequestBuilderInternal instantiates a new PrivilegedAccessGroupRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupRequestBuilder) {
    m := &PrivilegedAccessGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupRequestBuilder instantiates a new PrivilegedAccessGroupRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// EligibilitySchedules the eligibilitySchedules property
// returns a *PrivilegedAccessGroupEligibilitySchedulesRequestBuilder when successful
func (m *PrivilegedAccessGroupRequestBuilder) EligibilitySchedules()(*PrivilegedAccessGroupEligibilitySchedulesRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilitySchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
