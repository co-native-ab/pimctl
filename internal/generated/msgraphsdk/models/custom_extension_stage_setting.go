package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomExtensionStageSetting struct {
    Entity
    // Indicates the custom workflow extension that will be executed at this stage. Nullable. Supports $expand.
    customExtension CustomCalloutExtensionable
    // The stage property
    stage *AccessPackageCustomExtensionStage
}
// NewCustomExtensionStageSetting instantiates a new CustomExtensionStageSetting and sets the default values.
func NewCustomExtensionStageSetting()(*CustomExtensionStageSetting) {
    m := &CustomExtensionStageSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomExtensionStageSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomExtensionStageSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionStageSetting(), nil
}
// GetCustomExtension gets the customExtension property value. Indicates the custom workflow extension that will be executed at this stage. Nullable. Supports $expand.
// returns a CustomCalloutExtensionable when successful
func (m *CustomExtensionStageSetting) GetCustomExtension()(CustomCalloutExtensionable) {
    return m.customExtension
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomExtensionStageSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomCalloutExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(CustomCalloutExtensionable))
        }
        return nil
    }
    res["stage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCustomExtensionStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*AccessPackageCustomExtensionStage))
        }
        return nil
    }
    return res
}
// GetStage gets the stage property value. The stage property
// returns a *AccessPackageCustomExtensionStage when successful
func (m *CustomExtensionStageSetting) GetStage()(*AccessPackageCustomExtensionStage) {
    return m.stage
}
// Serialize serializes information the current object
func (m *CustomExtensionStageSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customExtension", m.GetCustomExtension())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err = writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomExtension sets the customExtension property value. Indicates the custom workflow extension that will be executed at this stage. Nullable. Supports $expand.
func (m *CustomExtensionStageSetting) SetCustomExtension(value CustomCalloutExtensionable)() {
    m.customExtension = value
}
// SetStage sets the stage property value. The stage property
func (m *CustomExtensionStageSetting) SetStage(value *AccessPackageCustomExtensionStage)() {
    m.stage = value
}
type CustomExtensionStageSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomExtension()(CustomCalloutExtensionable)
    GetStage()(*AccessPackageCustomExtensionStage)
    SetCustomExtension(value CustomCalloutExtensionable)()
    SetStage(value *AccessPackageCustomExtensionStage)()
}
