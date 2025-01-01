package rolemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DirectoryRoleAssignmentScheduleInstancesRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleAssignmentScheduleInstances
type DirectoryRoleAssignmentScheduleInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUnifiedRoleAssignmentScheduleInstanceId provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
// returns a *DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder when successful
func (m *DirectoryRoleAssignmentScheduleInstancesRequestBuilder) ByUnifiedRoleAssignmentScheduleInstanceId(unifiedRoleAssignmentScheduleInstanceId string)(*DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleAssignmentScheduleInstanceId != "" {
        urlTplParams["unifiedRoleAssignmentScheduleInstance%2Did"] = unifiedRoleAssignmentScheduleInstanceId
    }
    return NewDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleAssignmentScheduleInstancesRequestBuilderInternal instantiates a new DirectoryRoleAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentScheduleInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentScheduleInstancesRequestBuilder) {
    m := &DirectoryRoleAssignmentScheduleInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentScheduleInstances", pathParameters),
    }
    return m
}
// NewDirectoryRoleAssignmentScheduleInstancesRequestBuilder instantiates a new DirectoryRoleAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentScheduleInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentScheduleInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleAssignmentScheduleInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DirectoryRoleAssignmentScheduleInstancesCountRequestBuilder when successful
func (m *DirectoryRoleAssignmentScheduleInstancesRequestBuilder) Count()(*DirectoryRoleAssignmentScheduleInstancesCountRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleInstancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilder when successful
func (m *DirectoryRoleAssignmentScheduleInstancesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
