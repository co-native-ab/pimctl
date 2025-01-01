package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MailboxProtectionUnit struct {
    ProtectionUnitBase
    // The ID of the directory object.
    directoryObjectId *string
    // Display name of the directory object.
    displayName *string
    // Email address associated with the directory object.
    email *string
}
// NewMailboxProtectionUnit instantiates a new MailboxProtectionUnit and sets the default values.
func NewMailboxProtectionUnit()(*MailboxProtectionUnit) {
    m := &MailboxProtectionUnit{
        ProtectionUnitBase: *NewProtectionUnitBase(),
    }
    odataTypeValue := "#microsoft.graph.mailboxProtectionUnit"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMailboxProtectionUnitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMailboxProtectionUnitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxProtectionUnit(), nil
}
// GetDirectoryObjectId gets the directoryObjectId property value. The ID of the directory object.
// returns a *string when successful
func (m *MailboxProtectionUnit) GetDirectoryObjectId()(*string) {
    return m.directoryObjectId
}
// GetDisplayName gets the displayName property value. Display name of the directory object.
// returns a *string when successful
func (m *MailboxProtectionUnit) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. Email address associated with the directory object.
// returns a *string when successful
func (m *MailboxProtectionUnit) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MailboxProtectionUnit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectionUnitBase.GetFieldDeserializers()
    res["directoryObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryObjectId(val)
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
    return res
}
// Serialize serializes information the current object
func (m *MailboxProtectionUnit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectionUnitBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("directoryObjectId", m.GetDirectoryObjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectoryObjectId sets the directoryObjectId property value. The ID of the directory object.
func (m *MailboxProtectionUnit) SetDirectoryObjectId(value *string)() {
    m.directoryObjectId = value
}
// SetDisplayName sets the displayName property value. Display name of the directory object.
func (m *MailboxProtectionUnit) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. Email address associated with the directory object.
func (m *MailboxProtectionUnit) SetEmail(value *string)() {
    m.email = value
}
type MailboxProtectionUnitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectionUnitBaseable
    GetDirectoryObjectId()(*string)
    GetDisplayName()(*string)
    GetEmail()(*string)
    SetDirectoryObjectId(value *string)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
}
