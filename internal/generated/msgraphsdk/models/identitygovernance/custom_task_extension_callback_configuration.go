package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type CustomTaskExtensionCallbackConfiguration struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomExtensionCallbackConfiguration
    // The authorizedApps property
    authorizedApps []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Applicationable
}
// NewCustomTaskExtensionCallbackConfiguration instantiates a new CustomTaskExtensionCallbackConfiguration and sets the default values.
func NewCustomTaskExtensionCallbackConfiguration()(*CustomTaskExtensionCallbackConfiguration) {
    m := &CustomTaskExtensionCallbackConfiguration{
        CustomExtensionCallbackConfiguration: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewCustomExtensionCallbackConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.customTaskExtensionCallbackConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomTaskExtensionCallbackConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomTaskExtensionCallbackConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTaskExtensionCallbackConfiguration(), nil
}
// GetAuthorizedApps gets the authorizedApps property value. The authorizedApps property
// returns a []Applicationable when successful
func (m *CustomTaskExtensionCallbackConfiguration) GetAuthorizedApps()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Applicationable) {
    return m.authorizedApps
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomTaskExtensionCallbackConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionCallbackConfiguration.GetFieldDeserializers()
    res["authorizedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Applicationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Applicationable)
                }
            }
            m.SetAuthorizedApps(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CustomTaskExtensionCallbackConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionCallbackConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthorizedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthorizedApps()))
        for i, v := range m.GetAuthorizedApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("authorizedApps", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorizedApps sets the authorizedApps property value. The authorizedApps property
func (m *CustomTaskExtensionCallbackConfiguration) SetAuthorizedApps(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Applicationable)() {
    m.authorizedApps = value
}
type CustomTaskExtensionCallbackConfigurationable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomExtensionCallbackConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizedApps()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Applicationable)
    SetAuthorizedApps(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Applicationable)()
}
