package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamsAppPermissionSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // A collection of resource-specific permissions.
    resourceSpecificPermissions []TeamsAppResourceSpecificPermissionable
}
// NewTeamsAppPermissionSet instantiates a new TeamsAppPermissionSet and sets the default values.
func NewTeamsAppPermissionSet()(*TeamsAppPermissionSet) {
    m := &TeamsAppPermissionSet{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamsAppPermissionSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamsAppPermissionSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppPermissionSet(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamsAppPermissionSet) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamsAppPermissionSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["resourceSpecificPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamsAppResourceSpecificPermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsAppResourceSpecificPermissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TeamsAppResourceSpecificPermissionable)
                }
            }
            m.SetResourceSpecificPermissions(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TeamsAppPermissionSet) GetOdataType()(*string) {
    return m.odataType
}
// GetResourceSpecificPermissions gets the resourceSpecificPermissions property value. A collection of resource-specific permissions.
// returns a []TeamsAppResourceSpecificPermissionable when successful
func (m *TeamsAppPermissionSet) GetResourceSpecificPermissions()([]TeamsAppResourceSpecificPermissionable) {
    return m.resourceSpecificPermissions
}
// Serialize serializes information the current object
func (m *TeamsAppPermissionSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetResourceSpecificPermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceSpecificPermissions()))
        for i, v := range m.GetResourceSpecificPermissions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("resourceSpecificPermissions", cast)
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
func (m *TeamsAppPermissionSet) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamsAppPermissionSet) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResourceSpecificPermissions sets the resourceSpecificPermissions property value. A collection of resource-specific permissions.
func (m *TeamsAppPermissionSet) SetResourceSpecificPermissions(value []TeamsAppResourceSpecificPermissionable)() {
    m.resourceSpecificPermissions = value
}
type TeamsAppPermissionSetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetResourceSpecificPermissions()([]TeamsAppResourceSpecificPermissionable)
    SetOdataType(value *string)()
    SetResourceSpecificPermissions(value []TeamsAppResourceSpecificPermissionable)()
}
