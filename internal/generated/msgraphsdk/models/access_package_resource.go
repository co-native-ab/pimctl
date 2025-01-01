package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageResource struct {
    Entity
    // Contains information about the attributes to be collected from the requestor and sent to the resource application.
    attributes []AccessPackageResourceAttributeable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A description for the resource.
    description *string
    // The display name of the resource, such as the application name, group name or site name.
    displayName *string
    // Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
    environment AccessPackageResourceEnvironmentable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the resource in the origin system. For a Microsoft Entra group, this is the identifier of the group.
    originId *string
    // The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
    originSystem *string
    // Read-only. Nullable. Supports $expand.
    roles []AccessPackageResourceRoleable
    // Read-only. Nullable. Supports $expand.
    scopes []AccessPackageResourceScopeable
}
// NewAccessPackageResource instantiates a new AccessPackageResource and sets the default values.
func NewAccessPackageResource()(*AccessPackageResource) {
    m := &AccessPackageResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResource(), nil
}
// GetAttributes gets the attributes property value. Contains information about the attributes to be collected from the requestor and sent to the resource application.
// returns a []AccessPackageResourceAttributeable when successful
func (m *AccessPackageResource) GetAttributes()([]AccessPackageResourceAttributeable) {
    return m.attributes
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *AccessPackageResource) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. A description for the resource.
// returns a *string when successful
func (m *AccessPackageResource) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the resource, such as the application name, group name or site name.
// returns a *string when successful
func (m *AccessPackageResource) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnvironment gets the environment property value. Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
// returns a AccessPackageResourceEnvironmentable when successful
func (m *AccessPackageResource) GetEnvironment()(AccessPackageResourceEnvironmentable) {
    return m.environment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceAttributeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceAttributeable)
                }
            }
            m.SetAttributes(res)
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
    res["environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceEnvironmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironment(val.(AccessPackageResourceEnvironmentable))
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["originId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginId(val)
        }
        return nil
    }
    res["originSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginSystem(val)
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceRoleable)
                }
            }
            m.SetRoles(res)
        }
        return nil
    }
    res["scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceScopeable)
                }
            }
            m.SetScopes(res)
        }
        return nil
    }
    return res
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *AccessPackageResource) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetOriginId gets the originId property value. The unique identifier of the resource in the origin system. For a Microsoft Entra group, this is the identifier of the group.
// returns a *string when successful
func (m *AccessPackageResource) GetOriginId()(*string) {
    return m.originId
}
// GetOriginSystem gets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
// returns a *string when successful
func (m *AccessPackageResource) GetOriginSystem()(*string) {
    return m.originSystem
}
// GetRoles gets the roles property value. Read-only. Nullable. Supports $expand.
// returns a []AccessPackageResourceRoleable when successful
func (m *AccessPackageResource) GetRoles()([]AccessPackageResourceRoleable) {
    return m.roles
}
// GetScopes gets the scopes property value. Read-only. Nullable. Supports $expand.
// returns a []AccessPackageResourceScopeable when successful
func (m *AccessPackageResource) GetScopes()([]AccessPackageResourceScopeable) {
    return m.scopes
}
// Serialize serializes information the current object
func (m *AccessPackageResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttributes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributes()))
        for i, v := range m.GetAttributes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("attributes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
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
        err = writer.WriteObjectValue("environment", m.GetEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originId", m.GetOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originSystem", m.GetOriginSystem())
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoles()))
        for i, v := range m.GetRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScopes()))
        for i, v := range m.GetScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("scopes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttributes sets the attributes property value. Contains information about the attributes to be collected from the requestor and sent to the resource application.
func (m *AccessPackageResource) SetAttributes(value []AccessPackageResourceAttributeable)() {
    m.attributes = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageResource) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. A description for the resource.
func (m *AccessPackageResource) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the resource, such as the application name, group name or site name.
func (m *AccessPackageResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnvironment sets the environment property value. Contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
func (m *AccessPackageResource) SetEnvironment(value AccessPackageResourceEnvironmentable)() {
    m.environment = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageResource) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetOriginId sets the originId property value. The unique identifier of the resource in the origin system. For a Microsoft Entra group, this is the identifier of the group.
func (m *AccessPackageResource) SetOriginId(value *string)() {
    m.originId = value
}
// SetOriginSystem sets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
func (m *AccessPackageResource) SetOriginSystem(value *string)() {
    m.originSystem = value
}
// SetRoles sets the roles property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) SetRoles(value []AccessPackageResourceRoleable)() {
    m.roles = value
}
// SetScopes sets the scopes property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResource) SetScopes(value []AccessPackageResourceScopeable)() {
    m.scopes = value
}
type AccessPackageResourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributes()([]AccessPackageResourceAttributeable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEnvironment()(AccessPackageResourceEnvironmentable)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetRoles()([]AccessPackageResourceRoleable)
    GetScopes()([]AccessPackageResourceScopeable)
    SetAttributes(value []AccessPackageResourceAttributeable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEnvironment(value AccessPackageResourceEnvironmentable)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetRoles(value []AccessPackageResourceRoleable)()
    SetScopes(value []AccessPackageResourceScopeable)()
}
