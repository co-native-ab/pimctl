package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageResourceScope struct {
    Entity
    // The description of the scope.
    description *string
    // The display name of the scope.
    displayName *string
    // True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
    isRootScope *bool
    // The unique identifier for the scope in the resource as defined in the origin system.
    originId *string
    // The origin system for the scope.
    originSystem *string
    // The resource property
    resource AccessPackageResourceable
}
// NewAccessPackageResourceScope instantiates a new AccessPackageResourceScope and sets the default values.
func NewAccessPackageResourceScope()(*AccessPackageResourceScope) {
    m := &AccessPackageResourceScope{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageResourceScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageResourceScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceScope(), nil
}
// GetDescription gets the description property value. The description of the scope.
// returns a *string when successful
func (m *AccessPackageResourceScope) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the scope.
// returns a *string when successful
func (m *AccessPackageResourceScope) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageResourceScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["isRootScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRootScope(val)
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
// GetIsRootScope gets the isRootScope property value. True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
// returns a *bool when successful
func (m *AccessPackageResourceScope) GetIsRootScope()(*bool) {
    return m.isRootScope
}
// GetOriginId gets the originId property value. The unique identifier for the scope in the resource as defined in the origin system.
// returns a *string when successful
func (m *AccessPackageResourceScope) GetOriginId()(*string) {
    return m.originId
}
// GetOriginSystem gets the originSystem property value. The origin system for the scope.
// returns a *string when successful
func (m *AccessPackageResourceScope) GetOriginSystem()(*string) {
    return m.originSystem
}
// GetResource gets the resource property value. The resource property
// returns a AccessPackageResourceable when successful
func (m *AccessPackageResourceScope) GetResource()(AccessPackageResourceable) {
    return m.resource
}
// Serialize serializes information the current object
func (m *AccessPackageResourceScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("isRootScope", m.GetIsRootScope())
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
// SetDescription sets the description property value. The description of the scope.
func (m *AccessPackageResourceScope) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the scope.
func (m *AccessPackageResourceScope) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsRootScope sets the isRootScope property value. True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
func (m *AccessPackageResourceScope) SetIsRootScope(value *bool)() {
    m.isRootScope = value
}
// SetOriginId sets the originId property value. The unique identifier for the scope in the resource as defined in the origin system.
func (m *AccessPackageResourceScope) SetOriginId(value *string)() {
    m.originId = value
}
// SetOriginSystem sets the originSystem property value. The origin system for the scope.
func (m *AccessPackageResourceScope) SetOriginSystem(value *string)() {
    m.originSystem = value
}
// SetResource sets the resource property value. The resource property
func (m *AccessPackageResourceScope) SetResource(value AccessPackageResourceable)() {
    m.resource = value
}
type AccessPackageResourceScopeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsRootScope()(*bool)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetResource()(AccessPackageResourceable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsRootScope(value *bool)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetResource(value AccessPackageResourceable)()
}
