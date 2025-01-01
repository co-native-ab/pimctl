package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEventPresenter struct {
    Entity
    // Email address of the presenter.
    email *string
    // Identity information of the presenter. The supported identities are: communicationsGuestIdentity and communicationsUserIdentity.
    identity Identityable
    // Other details about the presenter. This property returns null when the virtual event type is virtualEventTownhall.
    presenterDetails VirtualEventPresenterDetailsable
}
// NewVirtualEventPresenter instantiates a new VirtualEventPresenter and sets the default values.
func NewVirtualEventPresenter()(*VirtualEventPresenter) {
    m := &VirtualEventPresenter{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEventPresenterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventPresenterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventPresenter(), nil
}
// GetEmail gets the email property value. Email address of the presenter.
// returns a *string when successful
func (m *VirtualEventPresenter) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEventPresenter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(Identityable))
        }
        return nil
    }
    res["presenterDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEventPresenterDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPresenterDetails(val.(VirtualEventPresenterDetailsable))
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. Identity information of the presenter. The supported identities are: communicationsGuestIdentity and communicationsUserIdentity.
// returns a Identityable when successful
func (m *VirtualEventPresenter) GetIdentity()(Identityable) {
    return m.identity
}
// GetPresenterDetails gets the presenterDetails property value. Other details about the presenter. This property returns null when the virtual event type is virtualEventTownhall.
// returns a VirtualEventPresenterDetailsable when successful
func (m *VirtualEventPresenter) GetPresenterDetails()(VirtualEventPresenterDetailsable) {
    return m.presenterDetails
}
// Serialize serializes information the current object
func (m *VirtualEventPresenter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("presenterDetails", m.GetPresenterDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmail sets the email property value. Email address of the presenter.
func (m *VirtualEventPresenter) SetEmail(value *string)() {
    m.email = value
}
// SetIdentity sets the identity property value. Identity information of the presenter. The supported identities are: communicationsGuestIdentity and communicationsUserIdentity.
func (m *VirtualEventPresenter) SetIdentity(value Identityable)() {
    m.identity = value
}
// SetPresenterDetails sets the presenterDetails property value. Other details about the presenter. This property returns null when the virtual event type is virtualEventTownhall.
func (m *VirtualEventPresenter) SetPresenterDetails(value VirtualEventPresenterDetailsable)() {
    m.presenterDetails = value
}
type VirtualEventPresenterable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetIdentity()(Identityable)
    GetPresenterDetails()(VirtualEventPresenterDetailsable)
    SetEmail(value *string)()
    SetIdentity(value Identityable)()
    SetPresenterDetails(value VirtualEventPresenterDetailsable)()
}
