package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Acl struct {
    // The accessType property
    accessType *AccessType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The type property
    typeEscaped *AclType
    // The unique identifer of the identity. For Microsoft Entra identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. For external groups value is set to the ID of the externalGroup
    value *string
}
// NewAcl instantiates a new Acl and sets the default values.
func NewAcl()(*Acl) {
    m := &Acl{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAclFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAclFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAcl(), nil
}
// GetAccessType gets the accessType property value. The accessType property
// returns a *AccessType when successful
func (m *Acl) GetAccessType()(*AccessType) {
    return m.accessType
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Acl) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Acl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessType(val.(*AccessType))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAclType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*AclType))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *Acl) GetOdataType()(*string) {
    return m.odataType
}
// GetTypeEscaped gets the type property value. The type property
// returns a *AclType when successful
func (m *Acl) GetTypeEscaped()(*AclType) {
    return m.typeEscaped
}
// GetValue gets the value property value. The unique identifer of the identity. For Microsoft Entra identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. For external groups value is set to the ID of the externalGroup
// returns a *string when successful
func (m *Acl) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *Acl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessType() != nil {
        cast := (*m.GetAccessType()).String()
        err := writer.WriteStringValue("accessType", &cast)
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
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
// SetAccessType sets the accessType property value. The accessType property
func (m *Acl) SetAccessType(value *AccessType)() {
    m.accessType = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Acl) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Acl) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Acl) SetTypeEscaped(value *AclType)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. The unique identifer of the identity. For Microsoft Entra identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. For external groups value is set to the ID of the externalGroup
func (m *Acl) SetValue(value *string)() {
    m.value = value
}
type Aclable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessType()(*AccessType)
    GetOdataType()(*string)
    GetTypeEscaped()(*AclType)
    GetValue()(*string)
    SetAccessType(value *AccessType)()
    SetOdataType(value *string)()
    SetTypeEscaped(value *AclType)()
    SetValue(value *string)()
}
