package rolemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DirectoryRoleAssignmentApprovalsRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleAssignmentApprovals
type DirectoryRoleAssignmentApprovalsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByApprovalId provides operations to manage the roleAssignmentApprovals property of the microsoft.graph.rbacApplication entity.
// returns a *DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsRequestBuilder) ByApprovalId(approvalId string)(*DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if approvalId != "" {
        urlTplParams["approval%2Did"] = approvalId
    }
    return NewDirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleAssignmentApprovalsRequestBuilderInternal instantiates a new DirectoryRoleAssignmentApprovalsRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsRequestBuilder) {
    m := &DirectoryRoleAssignmentApprovalsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentApprovals", pathParameters),
    }
    return m
}
// NewDirectoryRoleAssignmentApprovalsRequestBuilder instantiates a new DirectoryRoleAssignmentApprovalsRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentApprovalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentApprovalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleAssignmentApprovalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DirectoryRoleAssignmentApprovalsCountRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsRequestBuilder) Count()(*DirectoryRoleAssignmentApprovalsCountRequestBuilder) {
    return NewDirectoryRoleAssignmentApprovalsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnRequestBuilder when successful
func (m *DirectoryRoleAssignmentApprovalsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnRequestBuilder) {
    return NewDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
