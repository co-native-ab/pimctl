package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IdentityGovernanceRequestBuilder builds and executes requests for operations under \identityGovernance
type IdentityGovernanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewIdentityGovernanceRequestBuilderInternal instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    m := &IdentityGovernanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance", pathParameters),
    }
    return m
}
// NewIdentityGovernanceRequestBuilder instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceRequestBuilderInternal(urlParams, requestAdapter)
}
// PrivilegedAccess the privilegedAccess property
// returns a *PrivilegedAccessRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) PrivilegedAccess()(*PrivilegedAccessRequestBuilder) {
    return NewPrivilegedAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
