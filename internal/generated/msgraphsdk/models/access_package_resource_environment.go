package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageResourceEnvironment struct {
    Entity
    // Connection information of an environment used to connect to a resource.
    connectionInfo ConnectionInfoable
    // The date and time that this object was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of this object.
    description *string
    // The display name of this object.
    displayName *string
    // Determines whether this is default environment or not. It is set to true for all static origin systems, such as Microsoft Entra groups and Microsoft Entra Applications.
    isDefaultEnvironment *bool
    // The date and time that this object was last modified. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of this environment in the origin system.
    originId *string
    // The type of the resource in the origin system, that is, SharePointOnline. Requires $filter (eq).
    originSystem *string
    // Read-only. Required.
    resources []AccessPackageResourceable
}
// NewAccessPackageResourceEnvironment instantiates a new AccessPackageResourceEnvironment and sets the default values.
func NewAccessPackageResourceEnvironment()(*AccessPackageResourceEnvironment) {
    m := &AccessPackageResourceEnvironment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageResourceEnvironmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageResourceEnvironmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceEnvironment(), nil
}
// GetConnectionInfo gets the connectionInfo property value. Connection information of an environment used to connect to a resource.
// returns a ConnectionInfoable when successful
func (m *AccessPackageResourceEnvironment) GetConnectionInfo()(ConnectionInfoable) {
    return m.connectionInfo
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time that this object was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *AccessPackageResourceEnvironment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of this object.
// returns a *string when successful
func (m *AccessPackageResourceEnvironment) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of this object.
// returns a *string when successful
func (m *AccessPackageResourceEnvironment) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageResourceEnvironment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConnectionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionInfo(val.(ConnectionInfoable))
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
    res["isDefaultEnvironment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultEnvironment(val)
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
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
// GetIsDefaultEnvironment gets the isDefaultEnvironment property value. Determines whether this is default environment or not. It is set to true for all static origin systems, such as Microsoft Entra groups and Microsoft Entra Applications.
// returns a *bool when successful
func (m *AccessPackageResourceEnvironment) GetIsDefaultEnvironment()(*bool) {
    return m.isDefaultEnvironment
}
// GetModifiedDateTime gets the modifiedDateTime property value. The date and time that this object was last modified. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *AccessPackageResourceEnvironment) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetOriginId gets the originId property value. The unique identifier of this environment in the origin system.
// returns a *string when successful
func (m *AccessPackageResourceEnvironment) GetOriginId()(*string) {
    return m.originId
}
// GetOriginSystem gets the originSystem property value. The type of the resource in the origin system, that is, SharePointOnline. Requires $filter (eq).
// returns a *string when successful
func (m *AccessPackageResourceEnvironment) GetOriginSystem()(*string) {
    return m.originSystem
}
// GetResources gets the resources property value. Read-only. Required.
// returns a []AccessPackageResourceable when successful
func (m *AccessPackageResourceEnvironment) GetResources()([]AccessPackageResourceable) {
    return m.resources
}
// Serialize serializes information the current object
func (m *AccessPackageResourceEnvironment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("connectionInfo", m.GetConnectionInfo())
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
        err = writer.WriteBoolValue("isDefaultEnvironment", m.GetIsDefaultEnvironment())
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
// SetConnectionInfo sets the connectionInfo property value. Connection information of an environment used to connect to a resource.
func (m *AccessPackageResourceEnvironment) SetConnectionInfo(value ConnectionInfoable)() {
    m.connectionInfo = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time that this object was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageResourceEnvironment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of this object.
func (m *AccessPackageResourceEnvironment) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of this object.
func (m *AccessPackageResourceEnvironment) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsDefaultEnvironment sets the isDefaultEnvironment property value. Determines whether this is default environment or not. It is set to true for all static origin systems, such as Microsoft Entra groups and Microsoft Entra Applications.
func (m *AccessPackageResourceEnvironment) SetIsDefaultEnvironment(value *bool)() {
    m.isDefaultEnvironment = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The date and time that this object was last modified. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageResourceEnvironment) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetOriginId sets the originId property value. The unique identifier of this environment in the origin system.
func (m *AccessPackageResourceEnvironment) SetOriginId(value *string)() {
    m.originId = value
}
// SetOriginSystem sets the originSystem property value. The type of the resource in the origin system, that is, SharePointOnline. Requires $filter (eq).
func (m *AccessPackageResourceEnvironment) SetOriginSystem(value *string)() {
    m.originSystem = value
}
// SetResources sets the resources property value. Read-only. Required.
func (m *AccessPackageResourceEnvironment) SetResources(value []AccessPackageResourceable)() {
    m.resources = value
}
type AccessPackageResourceEnvironmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectionInfo()(ConnectionInfoable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsDefaultEnvironment()(*bool)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetResources()([]AccessPackageResourceable)
    SetConnectionInfo(value ConnectionInfoable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsDefaultEnvironment(value *bool)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetResources(value []AccessPackageResourceable)()
}
