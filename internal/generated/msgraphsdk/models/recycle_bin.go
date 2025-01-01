package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RecycleBin struct {
    BaseItem
    // List of the recycleBinItems deleted by a user.
    items []RecycleBinItemable
    // The settings property
    settings RecycleBinSettingsable
}
// NewRecycleBin instantiates a new RecycleBin and sets the default values.
func NewRecycleBin()(*RecycleBin) {
    m := &RecycleBin{
        BaseItem: *NewBaseItem(),
    }
    odataTypeValue := "#microsoft.graph.recycleBin"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRecycleBinFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecycleBinFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecycleBin(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecycleBin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecycleBinItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecycleBinItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RecycleBinItemable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecycleBinSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(RecycleBinSettingsable))
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. List of the recycleBinItems deleted by a user.
// returns a []RecycleBinItemable when successful
func (m *RecycleBin) GetItems()([]RecycleBinItemable) {
    return m.items
}
// GetSettings gets the settings property value. The settings property
// returns a RecycleBinSettingsable when successful
func (m *RecycleBin) GetSettings()(RecycleBinSettingsable) {
    return m.settings
}
// Serialize serializes information the current object
func (m *RecycleBin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItems sets the items property value. List of the recycleBinItems deleted by a user.
func (m *RecycleBin) SetItems(value []RecycleBinItemable)() {
    m.items = value
}
// SetSettings sets the settings property value. The settings property
func (m *RecycleBin) SetSettings(value RecycleBinSettingsable)() {
    m.settings = value
}
type RecycleBinable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]RecycleBinItemable)
    GetSettings()(RecycleBinSettingsable)
    SetItems(value []RecycleBinItemable)()
    SetSettings(value RecycleBinSettingsable)()
}
