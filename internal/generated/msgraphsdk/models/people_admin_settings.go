package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PeopleAdminSettings struct {
    Entity
    // Represents administrator settings that manage the support for item insights in an organization.
    itemInsights InsightsSettingsable
    // Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card.
    profileCardProperties []ProfileCardPropertyable
    // Represents administrator settings that manage the support of pronouns in an organization.
    pronouns PronounsSettingsable
}
// NewPeopleAdminSettings instantiates a new PeopleAdminSettings and sets the default values.
func NewPeopleAdminSettings()(*PeopleAdminSettings) {
    m := &PeopleAdminSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePeopleAdminSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePeopleAdminSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPeopleAdminSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PeopleAdminSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["itemInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemInsights(val.(InsightsSettingsable))
        }
        return nil
    }
    res["profileCardProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProfileCardPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfileCardPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProfileCardPropertyable)
                }
            }
            m.SetProfileCardProperties(res)
        }
        return nil
    }
    res["pronouns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePronounsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPronouns(val.(PronounsSettingsable))
        }
        return nil
    }
    return res
}
// GetItemInsights gets the itemInsights property value. Represents administrator settings that manage the support for item insights in an organization.
// returns a InsightsSettingsable when successful
func (m *PeopleAdminSettings) GetItemInsights()(InsightsSettingsable) {
    return m.itemInsights
}
// GetProfileCardProperties gets the profileCardProperties property value. Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card.
// returns a []ProfileCardPropertyable when successful
func (m *PeopleAdminSettings) GetProfileCardProperties()([]ProfileCardPropertyable) {
    return m.profileCardProperties
}
// GetPronouns gets the pronouns property value. Represents administrator settings that manage the support of pronouns in an organization.
// returns a PronounsSettingsable when successful
func (m *PeopleAdminSettings) GetPronouns()(PronounsSettingsable) {
    return m.pronouns
}
// Serialize serializes information the current object
func (m *PeopleAdminSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("itemInsights", m.GetItemInsights())
        if err != nil {
            return err
        }
    }
    if m.GetProfileCardProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProfileCardProperties()))
        for i, v := range m.GetProfileCardProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("profileCardProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("pronouns", m.GetPronouns())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemInsights sets the itemInsights property value. Represents administrator settings that manage the support for item insights in an organization.
func (m *PeopleAdminSettings) SetItemInsights(value InsightsSettingsable)() {
    m.itemInsights = value
}
// SetProfileCardProperties sets the profileCardProperties property value. Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card.
func (m *PeopleAdminSettings) SetProfileCardProperties(value []ProfileCardPropertyable)() {
    m.profileCardProperties = value
}
// SetPronouns sets the pronouns property value. Represents administrator settings that manage the support of pronouns in an organization.
func (m *PeopleAdminSettings) SetPronouns(value PronounsSettingsable)() {
    m.pronouns = value
}
type PeopleAdminSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemInsights()(InsightsSettingsable)
    GetProfileCardProperties()([]ProfileCardPropertyable)
    GetPronouns()(PronounsSettingsable)
    SetItemInsights(value InsightsSettingsable)()
    SetProfileCardProperties(value []ProfileCardPropertyable)()
    SetPronouns(value PronounsSettingsable)()
}
