package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type Identity struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The type of identity. Possible values are: user or group for Microsoft Entra identities and externalgroup for groups in an external system.
    typeEscaped *IdentityType
}
// NewIdentity instantiates a new Identity and sets the default values.
func NewIdentity()(*Identity) {
    m := &Identity{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Identity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*IdentityType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type of identity. Possible values are: user or group for Microsoft Entra identities and externalgroup for groups in an external system.
// returns a *IdentityType when successful
func (m *Identity) GetTypeEscaped()(*IdentityType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Identity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTypeEscaped sets the type property value. The type of identity. Possible values are: user or group for Microsoft Entra identities and externalgroup for groups in an external system.
func (m *Identity) SetTypeEscaped(value *IdentityType)() {
    m.typeEscaped = value
}
type Identityable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*IdentityType)
    SetTypeEscaped(value *IdentityType)()
}
