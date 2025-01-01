package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Teamwork struct {
    Entity
    // A collection of deleted chats.
    deletedChats []DeletedChatable
    // The deleted team.
    deletedTeams []DeletedTeamable
    // Indicates whether Microsoft Teams is enabled for the organization.
    isTeamsEnabled *bool
    // Represents the region of the organization or the tenant. The region value can be any region supported by the Teams payload. The possible values are: Americas, Europe and MiddleEast, Asia Pacific, UAE, Australia, Brazil, Canada, Switzerland, Germany, France, India, Japan, South Korea, Norway, Singapore, United Kingdom, South Africa, Sweden, Qatar, Poland, Italy, Israel, Spain, Mexico, USGov Community Cloud, USGov Community Cloud High, USGov Department of Defense, and China.
    region *string
    // Represents tenant-wide settings for all Teams apps in the tenant.
    teamsAppSettings TeamsAppSettingsable
    // The workforceIntegrations property
    workforceIntegrations []WorkforceIntegrationable
}
// NewTeamwork instantiates a new Teamwork and sets the default values.
func NewTeamwork()(*Teamwork) {
    m := &Teamwork{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamwork(), nil
}
// GetDeletedChats gets the deletedChats property value. A collection of deleted chats.
// returns a []DeletedChatable when successful
func (m *Teamwork) GetDeletedChats()([]DeletedChatable) {
    return m.deletedChats
}
// GetDeletedTeams gets the deletedTeams property value. The deleted team.
// returns a []DeletedTeamable when successful
func (m *Teamwork) GetDeletedTeams()([]DeletedTeamable) {
    return m.deletedTeams
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Teamwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deletedChats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeletedChatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeletedChatable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeletedChatable)
                }
            }
            m.SetDeletedChats(res)
        }
        return nil
    }
    res["deletedTeams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeletedTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeletedTeamable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeletedTeamable)
                }
            }
            m.SetDeletedTeams(res)
        }
        return nil
    }
    res["isTeamsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTeamsEnabled(val)
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["teamsAppSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppSettings(val.(TeamsAppSettingsable))
        }
        return nil
    }
    res["workforceIntegrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkforceIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkforceIntegrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkforceIntegrationable)
                }
            }
            m.SetWorkforceIntegrations(res)
        }
        return nil
    }
    return res
}
// GetIsTeamsEnabled gets the isTeamsEnabled property value. Indicates whether Microsoft Teams is enabled for the organization.
// returns a *bool when successful
func (m *Teamwork) GetIsTeamsEnabled()(*bool) {
    return m.isTeamsEnabled
}
// GetRegion gets the region property value. Represents the region of the organization or the tenant. The region value can be any region supported by the Teams payload. The possible values are: Americas, Europe and MiddleEast, Asia Pacific, UAE, Australia, Brazil, Canada, Switzerland, Germany, France, India, Japan, South Korea, Norway, Singapore, United Kingdom, South Africa, Sweden, Qatar, Poland, Italy, Israel, Spain, Mexico, USGov Community Cloud, USGov Community Cloud High, USGov Department of Defense, and China.
// returns a *string when successful
func (m *Teamwork) GetRegion()(*string) {
    return m.region
}
// GetTeamsAppSettings gets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
// returns a TeamsAppSettingsable when successful
func (m *Teamwork) GetTeamsAppSettings()(TeamsAppSettingsable) {
    return m.teamsAppSettings
}
// GetWorkforceIntegrations gets the workforceIntegrations property value. The workforceIntegrations property
// returns a []WorkforceIntegrationable when successful
func (m *Teamwork) GetWorkforceIntegrations()([]WorkforceIntegrationable) {
    return m.workforceIntegrations
}
// Serialize serializes information the current object
func (m *Teamwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeletedChats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedChats()))
        for i, v := range m.GetDeletedChats() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deletedChats", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedTeams()))
        for i, v := range m.GetDeletedTeams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deletedTeams", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTeamsEnabled", m.GetIsTeamsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsAppSettings", m.GetTeamsAppSettings())
        if err != nil {
            return err
        }
    }
    if m.GetWorkforceIntegrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkforceIntegrations()))
        for i, v := range m.GetWorkforceIntegrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("workforceIntegrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedChats sets the deletedChats property value. A collection of deleted chats.
func (m *Teamwork) SetDeletedChats(value []DeletedChatable)() {
    m.deletedChats = value
}
// SetDeletedTeams sets the deletedTeams property value. The deleted team.
func (m *Teamwork) SetDeletedTeams(value []DeletedTeamable)() {
    m.deletedTeams = value
}
// SetIsTeamsEnabled sets the isTeamsEnabled property value. Indicates whether Microsoft Teams is enabled for the organization.
func (m *Teamwork) SetIsTeamsEnabled(value *bool)() {
    m.isTeamsEnabled = value
}
// SetRegion sets the region property value. Represents the region of the organization or the tenant. The region value can be any region supported by the Teams payload. The possible values are: Americas, Europe and MiddleEast, Asia Pacific, UAE, Australia, Brazil, Canada, Switzerland, Germany, France, India, Japan, South Korea, Norway, Singapore, United Kingdom, South Africa, Sweden, Qatar, Poland, Italy, Israel, Spain, Mexico, USGov Community Cloud, USGov Community Cloud High, USGov Department of Defense, and China.
func (m *Teamwork) SetRegion(value *string)() {
    m.region = value
}
// SetTeamsAppSettings sets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
func (m *Teamwork) SetTeamsAppSettings(value TeamsAppSettingsable)() {
    m.teamsAppSettings = value
}
// SetWorkforceIntegrations sets the workforceIntegrations property value. The workforceIntegrations property
func (m *Teamwork) SetWorkforceIntegrations(value []WorkforceIntegrationable)() {
    m.workforceIntegrations = value
}
type Teamworkable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeletedChats()([]DeletedChatable)
    GetDeletedTeams()([]DeletedTeamable)
    GetIsTeamsEnabled()(*bool)
    GetRegion()(*string)
    GetTeamsAppSettings()(TeamsAppSettingsable)
    GetWorkforceIntegrations()([]WorkforceIntegrationable)
    SetDeletedChats(value []DeletedChatable)()
    SetDeletedTeams(value []DeletedTeamable)()
    SetIsTeamsEnabled(value *bool)()
    SetRegion(value *string)()
    SetTeamsAppSettings(value TeamsAppSettingsable)()
    SetWorkforceIntegrations(value []WorkforceIntegrationable)()
}
