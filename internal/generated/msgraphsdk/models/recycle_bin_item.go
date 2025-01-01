package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RecycleBinItem struct {
    BaseItem
    // Date and time when the item was deleted. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Relative URL of the list or folder that originally contained the item.
    deletedFromLocation *string
    // Size of the item in bytes.
    size *int64
}
// NewRecycleBinItem instantiates a new RecycleBinItem and sets the default values.
func NewRecycleBinItem()(*RecycleBinItem) {
    m := &RecycleBinItem{
        BaseItem: *NewBaseItem(),
    }
    odataTypeValue := "#microsoft.graph.recycleBinItem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRecycleBinItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecycleBinItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecycleBinItem(), nil
}
// GetDeletedDateTime gets the deletedDateTime property value. Date and time when the item was deleted. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *RecycleBinItem) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deletedDateTime
}
// GetDeletedFromLocation gets the deletedFromLocation property value. Relative URL of the list or folder that originally contained the item.
// returns a *string when successful
func (m *RecycleBinItem) GetDeletedFromLocation()(*string) {
    return m.deletedFromLocation
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecycleBinItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["deletedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDateTime(val)
        }
        return nil
    }
    res["deletedFromLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedFromLocation(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetSize gets the size property value. Size of the item in bytes.
// returns a *int64 when successful
func (m *RecycleBinItem) GetSize()(*int64) {
    return m.size
}
// Serialize serializes information the current object
func (m *RecycleBinItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("deletedDateTime", m.GetDeletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deletedFromLocation", m.GetDeletedFromLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedDateTime sets the deletedDateTime property value. Date and time when the item was deleted. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *RecycleBinItem) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
// SetDeletedFromLocation sets the deletedFromLocation property value. Relative URL of the list or folder that originally contained the item.
func (m *RecycleBinItem) SetDeletedFromLocation(value *string)() {
    m.deletedFromLocation = value
}
// SetSize sets the size property value. Size of the item in bytes.
func (m *RecycleBinItem) SetSize(value *int64)() {
    m.size = value
}
type RecycleBinItemable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeletedFromLocation()(*string)
    GetSize()(*int64)
    SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeletedFromLocation(value *string)()
    SetSize(value *int64)()
}
