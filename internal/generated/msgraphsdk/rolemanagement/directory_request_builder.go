package rolemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DirectoryRequestBuilder builds and executes requests for operations under \roleManagement\directory
type DirectoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory", pathParameters),
    }
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// RoleAssignmentScheduleInstances the roleAssignmentScheduleInstances property
// returns a *DirectoryRoleAssignmentScheduleInstancesRequestBuilder when successful
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleInstances()(*DirectoryRoleAssignmentScheduleInstancesRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignmentScheduleRequests provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
// returns a *DirectoryRoleAssignmentScheduleRequestsRequestBuilder when successful
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleRequests()(*DirectoryRoleAssignmentScheduleRequestsRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleEligibilitySchedules the roleEligibilitySchedules property
// returns a *DirectoryRoleEligibilitySchedulesRequestBuilder when successful
func (m *DirectoryRequestBuilder) RoleEligibilitySchedules()(*DirectoryRoleEligibilitySchedulesRequestBuilder) {
    return NewDirectoryRoleEligibilitySchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
