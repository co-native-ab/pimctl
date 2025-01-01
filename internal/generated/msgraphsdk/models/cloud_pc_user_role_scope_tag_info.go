package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcUserRoleScopeTagInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Scope tag display name.
    displayName *string
    // The OdataType property
    odataType *string
    // Scope tag identity.
    roleScopeTagId *string
}
// NewCloudPcUserRoleScopeTagInfo instantiates a new CloudPcUserRoleScopeTagInfo and sets the default values.
func NewCloudPcUserRoleScopeTagInfo()(*CloudPcUserRoleScopeTagInfo) {
    m := &CloudPcUserRoleScopeTagInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcUserRoleScopeTagInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcUserRoleScopeTagInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcUserRoleScopeTagInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcUserRoleScopeTagInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Scope tag display name.
// returns a *string when successful
func (m *CloudPcUserRoleScopeTagInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcUserRoleScopeTagInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["roleScopeTagId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleScopeTagId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcUserRoleScopeTagInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetRoleScopeTagId gets the roleScopeTagId property value. Scope tag identity.
// returns a *string when successful
func (m *CloudPcUserRoleScopeTagInfo) GetRoleScopeTagId()(*string) {
    return m.roleScopeTagId
}
// Serialize serializes information the current object
func (m *CloudPcUserRoleScopeTagInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteStringValue("roleScopeTagId", m.GetRoleScopeTagId())
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
func (m *CloudPcUserRoleScopeTagInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Scope tag display name.
func (m *CloudPcUserRoleScopeTagInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcUserRoleScopeTagInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRoleScopeTagId sets the roleScopeTagId property value. Scope tag identity.
func (m *CloudPcUserRoleScopeTagInfo) SetRoleScopeTagId(value *string)() {
    m.roleScopeTagId = value
}
type CloudPcUserRoleScopeTagInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetOdataType()(*string)
    GetRoleScopeTagId()(*string)
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
    SetRoleScopeTagId(value *string)()
}
