package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactFolder struct {
    Entity
    // The collection of child folders in the folder. Navigation property. Read-only. Nullable.
    childFolders []ContactFolderable
    // The contacts in the folder. Navigation property. Read-only. Nullable.
    contacts []Contactable
    // The folder's display name.
    displayName *string
    // The collection of multi-value extended properties defined for the contactFolder. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The ID of the folder's parent folder.
    parentFolderId *string
    // The collection of single-value extended properties defined for the contactFolder. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
}
// NewContactFolder instantiates a new ContactFolder and sets the default values.
func NewContactFolder()(*ContactFolder) {
    m := &ContactFolder{
        Entity: *NewEntity(),
    }
    return m
}
// CreateContactFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContactFolder(), nil
}
// GetChildFolders gets the childFolders property value. The collection of child folders in the folder. Navigation property. Read-only. Nullable.
// returns a []ContactFolderable when successful
func (m *ContactFolder) GetChildFolders()([]ContactFolderable) {
    return m.childFolders
}
// GetContacts gets the contacts property value. The contacts in the folder. Navigation property. Read-only. Nullable.
// returns a []Contactable when successful
func (m *ContactFolder) GetContacts()([]Contactable) {
    return m.contacts
}
// GetDisplayName gets the displayName property value. The folder's display name.
// returns a *string when successful
func (m *ContactFolder) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactFolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["childFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContactFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContactFolderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ContactFolderable)
                }
            }
            m.SetChildFolders(res)
        }
        return nil
    }
    res["contacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Contactable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Contactable)
                }
            }
            m.SetContacts(res)
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
    res["multiValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MultiValueLegacyExtendedPropertyable)
                }
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["parentFolderId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SingleValueLegacyExtendedPropertyable)
                }
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    return res
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contactFolder. Read-only. Nullable.
// returns a []MultiValueLegacyExtendedPropertyable when successful
func (m *ContactFolder) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetParentFolderId gets the parentFolderId property value. The ID of the folder's parent folder.
// returns a *string when successful
func (m *ContactFolder) GetParentFolderId()(*string) {
    return m.parentFolderId
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contactFolder. Read-only. Nullable.
// returns a []SingleValueLegacyExtendedPropertyable when successful
func (m *ContactFolder) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// Serialize serializes information the current object
func (m *ContactFolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildFolders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildFolders()))
        for i, v := range m.GetChildFolders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("childFolders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContacts()))
        for i, v := range m.GetContacts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("contacts", cast)
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
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildFolders sets the childFolders property value. The collection of child folders in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolder) SetChildFolders(value []ContactFolderable)() {
    m.childFolders = value
}
// SetContacts sets the contacts property value. The contacts in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolder) SetContacts(value []Contactable)() {
    m.contacts = value
}
// SetDisplayName sets the displayName property value. The folder's display name.
func (m *ContactFolder) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contactFolder. Read-only. Nullable.
func (m *ContactFolder) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetParentFolderId sets the parentFolderId property value. The ID of the folder's parent folder.
func (m *ContactFolder) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contactFolder. Read-only. Nullable.
func (m *ContactFolder) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
type ContactFolderable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildFolders()([]ContactFolderable)
    GetContacts()([]Contactable)
    GetDisplayName()(*string)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetParentFolderId()(*string)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    SetChildFolders(value []ContactFolderable)()
    SetContacts(value []Contactable)()
    SetDisplayName(value *string)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetParentFolderId(value *string)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
}
