package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrivilegedAccessGroupAssignmentApprovalsRequestBuilder builds and executes requests for operations under \identityGovernance\privilegedAccess\group\assignmentApprovals
type PrivilegedAccessGroupAssignmentApprovalsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByApprovalId provides operations to manage the assignmentApprovals property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedAccessGroupAssignmentApprovalsApprovalItemRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsRequestBuilder) ByApprovalId(approvalId string)(*PrivilegedAccessGroupAssignmentApprovalsApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if approvalId != "" {
        urlTplParams["approval%2Did"] = approvalId
    }
    return NewPrivilegedAccessGroupAssignmentApprovalsApprovalItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupAssignmentApprovalsRequestBuilderInternal instantiates a new PrivilegedAccessGroupAssignmentApprovalsRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentApprovalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentApprovalsRequestBuilder) {
    m := &PrivilegedAccessGroupAssignmentApprovalsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentApprovals", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupAssignmentApprovalsRequestBuilder instantiates a new PrivilegedAccessGroupAssignmentApprovalsRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentApprovalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentApprovalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupAssignmentApprovalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegedAccessGroupAssignmentApprovalsCountRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsRequestBuilder) Count()(*PrivilegedAccessGroupAssignmentApprovalsCountRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentApprovalsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *PrivilegedAccessGroupAssignmentApprovalsFilterByCurrentUserWithOnRequestBuilder when successful
func (m *PrivilegedAccessGroupAssignmentApprovalsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*PrivilegedAccessGroupAssignmentApprovalsFilterByCurrentUserWithOnRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentApprovalsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
