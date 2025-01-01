package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type ExternalItem struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // An array of access control entries. Each entry specifies the access granted to a user or group. Required.
    acl []Aclable
    // Returns a list of activities performed on the item. Write-only.
    activities []ExternalActivityable
    // A plain-text  representation of the contents of the item. The text in this property is full-text indexed. Optional.
    content ExternalItemContentable
    // A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
    properties Propertiesable
}
// NewExternalItem instantiates a new ExternalItem and sets the default values.
func NewExternalItem()(*ExternalItem) {
    m := &ExternalItem{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateExternalItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalItem(), nil
}
// GetAcl gets the acl property value. An array of access control entries. Each entry specifies the access granted to a user or group. Required.
// returns a []Aclable when successful
func (m *ExternalItem) GetAcl()([]Aclable) {
    return m.acl
}
// GetActivities gets the activities property value. Returns a list of activities performed on the item. Write-only.
// returns a []ExternalActivityable when successful
func (m *ExternalItem) GetActivities()([]ExternalActivityable) {
    return m.activities
}
// GetContent gets the content property value. A plain-text  representation of the contents of the item. The text in this property is full-text indexed. Optional.
// returns a ExternalItemContentable when successful
func (m *ExternalItem) GetContent()(ExternalItemContentable) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAclFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Aclable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Aclable)
                }
            }
            m.SetAcl(res)
        }
        return nil
    }
    res["activities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalActivityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExternalActivityable)
                }
            }
            m.SetActivities(res)
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalItemContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(ExternalItemContentable))
        }
        return nil
    }
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePropertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperties(val.(Propertiesable))
        }
        return nil
    }
    return res
}
// GetProperties gets the properties property value. A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
// returns a Propertiesable when successful
func (m *ExternalItem) GetProperties()(Propertiesable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *ExternalItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAcl() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAcl()))
        for i, v := range m.GetAcl() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("acl", cast)
        if err != nil {
            return err
        }
    }
    if m.GetActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("activities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("properties", m.GetProperties())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcl sets the acl property value. An array of access control entries. Each entry specifies the access granted to a user or group. Required.
func (m *ExternalItem) SetAcl(value []Aclable)() {
    m.acl = value
}
// SetActivities sets the activities property value. Returns a list of activities performed on the item. Write-only.
func (m *ExternalItem) SetActivities(value []ExternalActivityable)() {
    m.activities = value
}
// SetContent sets the content property value. A plain-text  representation of the contents of the item. The text in this property is full-text indexed. Optional.
func (m *ExternalItem) SetContent(value ExternalItemContentable)() {
    m.content = value
}
// SetProperties sets the properties property value. A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
func (m *ExternalItem) SetProperties(value Propertiesable)() {
    m.properties = value
}
type ExternalItemable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcl()([]Aclable)
    GetActivities()([]ExternalActivityable)
    GetContent()(ExternalItemContentable)
    GetProperties()(Propertiesable)
    SetAcl(value []Aclable)()
    SetActivities(value []ExternalActivityable)()
    SetContent(value ExternalItemContentable)()
    SetProperties(value Propertiesable)()
}
