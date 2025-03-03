package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type ExternalGroup struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The description of the external group. Optional.
    description *string
    // The friendly name of the external group. Optional.
    displayName *string
    // A member added to an externalGroup. You can add Microsoft Entra users, Microsoft Entra groups, or an externalGroup as members.
    members []Identityable
}
// NewExternalGroup instantiates a new ExternalGroup and sets the default values.
func NewExternalGroup()(*ExternalGroup) {
    m := &ExternalGroup{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateExternalGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalGroup(), nil
}
// GetDescription gets the description property value. The description of the external group. Optional.
// returns a *string when successful
func (m *ExternalGroup) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The friendly name of the external group. Optional.
// returns a *string when successful
func (m *ExternalGroup) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Identityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Identityable)
                }
            }
            m.SetMembers(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. A member added to an externalGroup. You can add Microsoft Entra users, Microsoft Entra groups, or an externalGroup as members.
// returns a []Identityable when successful
func (m *ExternalGroup) GetMembers()([]Identityable) {
    return m.members
}
// Serialize serializes information the current object
func (m *ExternalGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description of the external group. Optional.
func (m *ExternalGroup) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The friendly name of the external group. Optional.
func (m *ExternalGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetMembers sets the members property value. A member added to an externalGroup. You can add Microsoft Entra users, Microsoft Entra groups, or an externalGroup as members.
func (m *ExternalGroup) SetMembers(value []Identityable)() {
    m.members = value
}
type ExternalGroupable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetMembers()([]Identityable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetMembers(value []Identityable)()
}
