package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileStorageContainerViewpoint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The current user's effective role. Read-only.
    effectiveRole *string
    // The OdataType property
    odataType *string
}
// NewFileStorageContainerViewpoint instantiates a new FileStorageContainerViewpoint and sets the default values.
func NewFileStorageContainerViewpoint()(*FileStorageContainerViewpoint) {
    m := &FileStorageContainerViewpoint{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileStorageContainerViewpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainerViewpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainerViewpoint(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FileStorageContainerViewpoint) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEffectiveRole gets the effectiveRole property value. The current user's effective role. Read-only.
// returns a *string when successful
func (m *FileStorageContainerViewpoint) GetEffectiveRole()(*string) {
    return m.effectiveRole
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileStorageContainerViewpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["effectiveRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEffectiveRole(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *FileStorageContainerViewpoint) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *FileStorageContainerViewpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("effectiveRole", m.GetEffectiveRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileStorageContainerViewpoint) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEffectiveRole sets the effectiveRole property value. The current user's effective role. Read-only.
func (m *FileStorageContainerViewpoint) SetEffectiveRole(value *string)() {
    m.effectiveRole = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FileStorageContainerViewpoint) SetOdataType(value *string)() {
    m.odataType = value
}
type FileStorageContainerViewpointable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEffectiveRole()(*string)
    GetOdataType()(*string)
    SetEffectiveRole(value *string)()
    SetOdataType(value *string)()
}
