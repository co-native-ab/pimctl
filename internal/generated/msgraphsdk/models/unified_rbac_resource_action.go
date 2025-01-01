package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UnifiedRbacResourceAction struct {
    Entity
    // The actionVerb property
    actionVerb *string
    // The authenticationContextId property
    authenticationContextId *string
    // The description property
    description *string
    // The isAuthenticationContextSettable property
    isAuthenticationContextSettable *bool
    // The name property
    name *string
    // The resourceScopeId property
    resourceScopeId *string
}
// NewUnifiedRbacResourceAction instantiates a new UnifiedRbacResourceAction and sets the default values.
func NewUnifiedRbacResourceAction()(*UnifiedRbacResourceAction) {
    m := &UnifiedRbacResourceAction{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRbacResourceActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUnifiedRbacResourceActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRbacResourceAction(), nil
}
// GetActionVerb gets the actionVerb property value. The actionVerb property
// returns a *string when successful
func (m *UnifiedRbacResourceAction) GetActionVerb()(*string) {
    return m.actionVerb
}
// GetAuthenticationContextId gets the authenticationContextId property value. The authenticationContextId property
// returns a *string when successful
func (m *UnifiedRbacResourceAction) GetAuthenticationContextId()(*string) {
    return m.authenticationContextId
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *UnifiedRbacResourceAction) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UnifiedRbacResourceAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionVerb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionVerb(val)
        }
        return nil
    }
    res["authenticationContextId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationContextId(val)
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
    res["isAuthenticationContextSettable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAuthenticationContextSettable(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["resourceScopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceScopeId(val)
        }
        return nil
    }
    return res
}
// GetIsAuthenticationContextSettable gets the isAuthenticationContextSettable property value. The isAuthenticationContextSettable property
// returns a *bool when successful
func (m *UnifiedRbacResourceAction) GetIsAuthenticationContextSettable()(*bool) {
    return m.isAuthenticationContextSettable
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *UnifiedRbacResourceAction) GetName()(*string) {
    return m.name
}
// GetResourceScopeId gets the resourceScopeId property value. The resourceScopeId property
// returns a *string when successful
func (m *UnifiedRbacResourceAction) GetResourceScopeId()(*string) {
    return m.resourceScopeId
}
// Serialize serializes information the current object
func (m *UnifiedRbacResourceAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionVerb", m.GetActionVerb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("authenticationContextId", m.GetAuthenticationContextId())
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
        err = writer.WriteBoolValue("isAuthenticationContextSettable", m.GetIsAuthenticationContextSettable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceScopeId", m.GetResourceScopeId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionVerb sets the actionVerb property value. The actionVerb property
func (m *UnifiedRbacResourceAction) SetActionVerb(value *string)() {
    m.actionVerb = value
}
// SetAuthenticationContextId sets the authenticationContextId property value. The authenticationContextId property
func (m *UnifiedRbacResourceAction) SetAuthenticationContextId(value *string)() {
    m.authenticationContextId = value
}
// SetDescription sets the description property value. The description property
func (m *UnifiedRbacResourceAction) SetDescription(value *string)() {
    m.description = value
}
// SetIsAuthenticationContextSettable sets the isAuthenticationContextSettable property value. The isAuthenticationContextSettable property
func (m *UnifiedRbacResourceAction) SetIsAuthenticationContextSettable(value *bool)() {
    m.isAuthenticationContextSettable = value
}
// SetName sets the name property value. The name property
func (m *UnifiedRbacResourceAction) SetName(value *string)() {
    m.name = value
}
// SetResourceScopeId sets the resourceScopeId property value. The resourceScopeId property
func (m *UnifiedRbacResourceAction) SetResourceScopeId(value *string)() {
    m.resourceScopeId = value
}
type UnifiedRbacResourceActionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionVerb()(*string)
    GetAuthenticationContextId()(*string)
    GetDescription()(*string)
    GetIsAuthenticationContextSettable()(*bool)
    GetName()(*string)
    GetResourceScopeId()(*string)
    SetActionVerb(value *string)()
    SetAuthenticationContextId(value *string)()
    SetDescription(value *string)()
    SetIsAuthenticationContextSettable(value *bool)()
    SetName(value *string)()
    SetResourceScopeId(value *string)()
}
