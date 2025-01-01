package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PrivilegedAccessRequestBuilder builds and executes requests for operations under \identityGovernance\privilegedAccess
type PrivilegedAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewPrivilegedAccessRequestBuilderInternal instantiates a new PrivilegedAccessRequestBuilder and sets the default values.
func NewPrivilegedAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessRequestBuilder) {
    m := &PrivilegedAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess", pathParameters),
    }
    return m
}
// NewPrivilegedAccessRequestBuilder instantiates a new PrivilegedAccessRequestBuilder and sets the default values.
func NewPrivilegedAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Group the group property
// returns a *PrivilegedAccessGroupRequestBuilder when successful
func (m *PrivilegedAccessRequestBuilder) Group()(*PrivilegedAccessGroupRequestBuilder) {
    return NewPrivilegedAccessGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
