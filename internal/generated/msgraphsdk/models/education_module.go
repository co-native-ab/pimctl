package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EducationModule struct {
    Entity
    // The display name of the user that created the module.
    createdBy IdentitySetable
    // Date time the module was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the module.
    description *string
    // Name of the module.
    displayName *string
    // Indicates whether the module is pinned or not.
    isPinned *bool
    // The last user that modified the module.
    lastModifiedBy IdentitySetable
    // Date time the module was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Learning objects that are associated with this module. Only teachers can modify this list. Nullable.
    resources []EducationModuleResourceable
    // Folder URL where all the file resources for this module are stored.
    resourcesFolderUrl *string
    // Status of the module. You can't use a PATCH operation to update this value. Possible values are: draft and published.
    status *EducationModuleStatus
}
// NewEducationModule instantiates a new EducationModule and sets the default values.
func NewEducationModule()(*EducationModule) {
    m := &EducationModule{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationModuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationModuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationModule(), nil
}
// GetCreatedBy gets the createdBy property value. The display name of the user that created the module.
// returns a IdentitySetable when successful
func (m *EducationModule) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Date time the module was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *EducationModule) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the module.
// returns a *string when successful
func (m *EducationModule) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the module.
// returns a *string when successful
func (m *EducationModule) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationModule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
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
    res["isPinned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPinned(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationModuleResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationModuleResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationModuleResourceable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    res["resourcesFolderUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourcesFolderUrl(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationModuleStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*EducationModuleStatus))
        }
        return nil
    }
    return res
}
// GetIsPinned gets the isPinned property value. Indicates whether the module is pinned or not.
// returns a *bool when successful
func (m *EducationModule) GetIsPinned()(*bool) {
    return m.isPinned
}
// GetLastModifiedBy gets the lastModifiedBy property value. The last user that modified the module.
// returns a IdentitySetable when successful
func (m *EducationModule) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date time the module was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *EducationModule) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetResources gets the resources property value. Learning objects that are associated with this module. Only teachers can modify this list. Nullable.
// returns a []EducationModuleResourceable when successful
func (m *EducationModule) GetResources()([]EducationModuleResourceable) {
    return m.resources
}
// GetResourcesFolderUrl gets the resourcesFolderUrl property value. Folder URL where all the file resources for this module are stored.
// returns a *string when successful
func (m *EducationModule) GetResourcesFolderUrl()(*string) {
    return m.resourcesFolderUrl
}
// GetStatus gets the status property value. Status of the module. You can't use a PATCH operation to update this value. Possible values are: draft and published.
// returns a *EducationModuleStatus when successful
func (m *EducationModule) GetStatus()(*EducationModuleStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *EducationModule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteBoolValue("isPinned", m.GetIsPinned())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The display name of the user that created the module.
func (m *EducationModule) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date time the module was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z
func (m *EducationModule) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the module.
func (m *EducationModule) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the module.
func (m *EducationModule) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsPinned sets the isPinned property value. Indicates whether the module is pinned or not.
func (m *EducationModule) SetIsPinned(value *bool)() {
    m.isPinned = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The last user that modified the module.
func (m *EducationModule) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date time the module was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z
func (m *EducationModule) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetResources sets the resources property value. Learning objects that are associated with this module. Only teachers can modify this list. Nullable.
func (m *EducationModule) SetResources(value []EducationModuleResourceable)() {
    m.resources = value
}
// SetResourcesFolderUrl sets the resourcesFolderUrl property value. Folder URL where all the file resources for this module are stored.
func (m *EducationModule) SetResourcesFolderUrl(value *string)() {
    m.resourcesFolderUrl = value
}
// SetStatus sets the status property value. Status of the module. You can't use a PATCH operation to update this value. Possible values are: draft and published.
func (m *EducationModule) SetStatus(value *EducationModuleStatus)() {
    m.status = value
}
type EducationModuleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsPinned()(*bool)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResources()([]EducationModuleResourceable)
    GetResourcesFolderUrl()(*string)
    GetStatus()(*EducationModuleStatus)
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsPinned(value *bool)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResources(value []EducationModuleResourceable)()
    SetResourcesFolderUrl(value *string)()
    SetStatus(value *EducationModuleStatus)()
}
