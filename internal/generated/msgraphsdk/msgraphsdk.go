package msgraphsdk

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i52fa9b9ecd827f8ae5045560d8e44d42ea264e729313aa5cf3fbd5ded6ca6f6f "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/me"
    i97e62deba8acb92aa3bfdc812640fa33c616771ad08af022f7de03b008423e71 "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/policies"
    ie594ffa2655a81bfcb56be3cd0c0a0cf8c9808c32435f6366c30f7992bbe383d "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/identitygovernance"
    if7327de36188d1fe2b58c9b828971d417a32e1ba0da4d7df3e18fd75dd70cf29 "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/rolemanagement"
)

// Msgraphsdk the main entry point of the SDK, exposes the configuration and the fluent API.
type Msgraphsdk struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewMsgraphsdk instantiates a new Msgraphsdk and sets the default values.
func NewMsgraphsdk(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Msgraphsdk) {
    m := &Msgraphsdk{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://graph.microsoft.com/v1.0")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// IdentityGovernance the identityGovernance property
// returns a *IdentityGovernanceRequestBuilder when successful
func (m *Msgraphsdk) IdentityGovernance()(*ie594ffa2655a81bfcb56be3cd0c0a0cf8c9808c32435f6366c30f7992bbe383d.IdentityGovernanceRequestBuilder) {
    return ie594ffa2655a81bfcb56be3cd0c0a0cf8c9808c32435f6366c30f7992bbe383d.NewIdentityGovernanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Me provides operations to manage the user singleton.
// returns a *MeRequestBuilder when successful
func (m *Msgraphsdk) Me()(*i52fa9b9ecd827f8ae5045560d8e44d42ea264e729313aa5cf3fbd5ded6ca6f6f.MeRequestBuilder) {
    return i52fa9b9ecd827f8ae5045560d8e44d42ea264e729313aa5cf3fbd5ded6ca6f6f.NewMeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Policies the policies property
// returns a *PoliciesRequestBuilder when successful
func (m *Msgraphsdk) Policies()(*i97e62deba8acb92aa3bfdc812640fa33c616771ad08af022f7de03b008423e71.PoliciesRequestBuilder) {
    return i97e62deba8acb92aa3bfdc812640fa33c616771ad08af022f7de03b008423e71.NewPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleManagement the roleManagement property
// returns a *RoleManagementRequestBuilder when successful
func (m *Msgraphsdk) RoleManagement()(*if7327de36188d1fe2b58c9b828971d417a32e1ba0da4d7df3e18fd75dd70cf29.RoleManagementRequestBuilder) {
    return if7327de36188d1fe2b58c9b828971d417a32e1ba0da4d7df3e18fd75dd70cf29.NewRoleManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
