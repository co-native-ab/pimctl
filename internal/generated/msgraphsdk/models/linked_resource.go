package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LinkedResource struct {
    Entity
    // The app name of the source that sends the linkedResource.
    applicationName *string
    // The title of the linkedResource.
    displayName *string
    // ID of the object that is associated with this task on the third-party/partner system.
    externalId *string
    // Deep link to the linkedResource.
    webUrl *string
}
// NewLinkedResource instantiates a new LinkedResource and sets the default values.
func NewLinkedResource()(*LinkedResource) {
    m := &LinkedResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateLinkedResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLinkedResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLinkedResource(), nil
}
// GetApplicationName gets the applicationName property value. The app name of the source that sends the linkedResource.
// returns a *string when successful
func (m *LinkedResource) GetApplicationName()(*string) {
    return m.applicationName
}
// GetDisplayName gets the displayName property value. The title of the linkedResource.
// returns a *string when successful
func (m *LinkedResource) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalId gets the externalId property value. ID of the object that is associated with this task on the third-party/partner system.
// returns a *string when successful
func (m *LinkedResource) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LinkedResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationName(val)
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
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetWebUrl gets the webUrl property value. Deep link to the linkedResource.
// returns a *string when successful
func (m *LinkedResource) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *LinkedResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationName", m.GetApplicationName())
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationName sets the applicationName property value. The app name of the source that sends the linkedResource.
func (m *LinkedResource) SetApplicationName(value *string)() {
    m.applicationName = value
}
// SetDisplayName sets the displayName property value. The title of the linkedResource.
func (m *LinkedResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalId sets the externalId property value. ID of the object that is associated with this task on the third-party/partner system.
func (m *LinkedResource) SetExternalId(value *string)() {
    m.externalId = value
}
// SetWebUrl sets the webUrl property value. Deep link to the linkedResource.
func (m *LinkedResource) SetWebUrl(value *string)() {
    m.webUrl = value
}
type LinkedResourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationName()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetWebUrl()(*string)
    SetApplicationName(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetWebUrl(value *string)()
}
