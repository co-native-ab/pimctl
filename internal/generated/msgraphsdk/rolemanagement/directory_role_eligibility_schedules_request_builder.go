package rolemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DirectoryRoleEligibilitySchedulesRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleEligibilitySchedules
type DirectoryRoleEligibilitySchedulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUnifiedRoleEligibilityScheduleId provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
// returns a *DirectoryRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder when successful
func (m *DirectoryRoleEligibilitySchedulesRequestBuilder) ByUnifiedRoleEligibilityScheduleId(unifiedRoleEligibilityScheduleId string)(*DirectoryRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleEligibilityScheduleId != "" {
        urlTplParams["unifiedRoleEligibilitySchedule%2Did"] = unifiedRoleEligibilityScheduleId
    }
    return NewDirectoryRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleEligibilitySchedulesRequestBuilderInternal instantiates a new DirectoryRoleEligibilitySchedulesRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilitySchedulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleEligibilitySchedulesRequestBuilder) {
    m := &DirectoryRoleEligibilitySchedulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleEligibilitySchedules", pathParameters),
    }
    return m
}
// NewDirectoryRoleEligibilitySchedulesRequestBuilder instantiates a new DirectoryRoleEligibilitySchedulesRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilitySchedulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleEligibilitySchedulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleEligibilitySchedulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DirectoryRoleEligibilitySchedulesCountRequestBuilder when successful
func (m *DirectoryRoleEligibilitySchedulesRequestBuilder) Count()(*DirectoryRoleEligibilitySchedulesCountRequestBuilder) {
    return NewDirectoryRoleEligibilitySchedulesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *DirectoryRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder when successful
func (m *DirectoryRoleEligibilitySchedulesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*DirectoryRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilder) {
    return NewDirectoryRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
