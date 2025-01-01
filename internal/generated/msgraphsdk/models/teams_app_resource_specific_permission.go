package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamsAppResourceSpecificPermission struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The type of resource-specific permission.
    permissionType *TeamsAppResourceSpecificPermissionType
    // The name of the resource-specific permission.
    permissionValue *string
}
// NewTeamsAppResourceSpecificPermission instantiates a new TeamsAppResourceSpecificPermission and sets the default values.
func NewTeamsAppResourceSpecificPermission()(*TeamsAppResourceSpecificPermission) {
    m := &TeamsAppResourceSpecificPermission{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamsAppResourceSpecificPermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamsAppResourceSpecificPermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppResourceSpecificPermission(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamsAppResourceSpecificPermission) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamsAppResourceSpecificPermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["permissionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppResourceSpecificPermissionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionType(val.(*TeamsAppResourceSpecificPermissionType))
        }
        return nil
    }
    res["permissionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionValue(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TeamsAppResourceSpecificPermission) GetOdataType()(*string) {
    return m.odataType
}
// GetPermissionType gets the permissionType property value. The type of resource-specific permission.
// returns a *TeamsAppResourceSpecificPermissionType when successful
func (m *TeamsAppResourceSpecificPermission) GetPermissionType()(*TeamsAppResourceSpecificPermissionType) {
    return m.permissionType
}
// GetPermissionValue gets the permissionValue property value. The name of the resource-specific permission.
// returns a *string when successful
func (m *TeamsAppResourceSpecificPermission) GetPermissionValue()(*string) {
    return m.permissionValue
}
// Serialize serializes information the current object
func (m *TeamsAppResourceSpecificPermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionType() != nil {
        cast := (*m.GetPermissionType()).String()
        err := writer.WriteStringValue("permissionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("permissionValue", m.GetPermissionValue())
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
func (m *TeamsAppResourceSpecificPermission) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamsAppResourceSpecificPermission) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPermissionType sets the permissionType property value. The type of resource-specific permission.
func (m *TeamsAppResourceSpecificPermission) SetPermissionType(value *TeamsAppResourceSpecificPermissionType)() {
    m.permissionType = value
}
// SetPermissionValue sets the permissionValue property value. The name of the resource-specific permission.
func (m *TeamsAppResourceSpecificPermission) SetPermissionValue(value *string)() {
    m.permissionValue = value
}
type TeamsAppResourceSpecificPermissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPermissionType()(*TeamsAppResourceSpecificPermissionType)
    GetPermissionValue()(*string)
    SetOdataType(value *string)()
    SetPermissionType(value *TeamsAppResourceSpecificPermissionType)()
    SetPermissionValue(value *string)()
}
