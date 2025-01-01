package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageResourceRole struct {
    Entity
    // A description for the resource role.
    description *string
    // The display name of the resource role such as the role defined by the application.
    displayName *string
    // The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId is the sequence number of the role in the site.
    originId *string
    // The type of the resource in the origin system, such as SharePointOnline, AadApplication, or AadGroup.
    originSystem *string
    // The resource property
    resource AccessPackageResourceable
}
// NewAccessPackageResourceRole instantiates a new AccessPackageResourceRole and sets the default values.
func NewAccessPackageResourceRole()(*AccessPackageResourceRole) {
    m := &AccessPackageResourceRole{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageResourceRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageResourceRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceRole(), nil
}
// GetDescription gets the description property value. A description for the resource role.
// returns a *string when successful
func (m *AccessPackageResourceRole) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the resource role such as the role defined by the application.
// returns a *string when successful
func (m *AccessPackageResourceRole) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageResourceRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(AccessPackageResourceable))
        }
        return nil
    }
    return res
}
// GetOriginId gets the originId property value. The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId is the sequence number of the role in the site.
// returns a *string when successful
func (m *AccessPackageResourceRole) GetOriginId()(*string) {
    return m.originId
}
// GetOriginSystem gets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication, or AadGroup.
// returns a *string when successful
func (m *AccessPackageResourceRole) GetOriginSystem()(*string) {
    return m.originSystem
}
// GetResource gets the resource property value. The resource property
// returns a AccessPackageResourceable when successful
func (m *AccessPackageResourceRole) GetResource()(AccessPackageResourceable) {
    return m.resource
}
// Serialize serializes information the current object
func (m *AccessPackageResourceRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. A description for the resource role.
func (m *AccessPackageResourceRole) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the resource role such as the role defined by the application.
func (m *AccessPackageResourceRole) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOriginId sets the originId property value. The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId is the sequence number of the role in the site.
func (m *AccessPackageResourceRole) SetOriginId(value *string)() {
    m.originId = value
}
// SetOriginSystem sets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication, or AadGroup.
func (m *AccessPackageResourceRole) SetOriginSystem(value *string)() {
    m.originSystem = value
}
// SetResource sets the resource property value. The resource property
func (m *AccessPackageResourceRole) SetResource(value AccessPackageResourceable)() {
    m.resource = value
}
type AccessPackageResourceRoleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetResource()(AccessPackageResourceable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetResource(value AccessPackageResourceable)()
}
