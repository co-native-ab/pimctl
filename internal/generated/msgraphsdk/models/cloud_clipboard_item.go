package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudClipboardItem struct {
    Entity
    // Set by the server. DateTime in UTC when the object was created on the server.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Set by the server. DateTime in UTC when the object expires and after that the object is no longer available. The default and also maximum TTL is 12 hours after the creation, but it might change for performance optimization.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Set by the server if not provided in the client's request. DateTime in UTC when the object was modified by the client.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A cloudClipboardItem can have multiple cloudClipboardItemPayload objects in the payloads. A window can place more than one clipboard object on the clipboard. Each one represents the same information in a different clipboard format.
    payloads []CloudClipboardItemPayloadable
}
// NewCloudClipboardItem instantiates a new CloudClipboardItem and sets the default values.
func NewCloudClipboardItem()(*CloudClipboardItem) {
    m := &CloudClipboardItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudClipboardItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudClipboardItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudClipboardItem(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Set by the server. DateTime in UTC when the object was created on the server.
// returns a *Time when successful
func (m *CloudClipboardItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetExpirationDateTime gets the expirationDateTime property value. Set by the server. DateTime in UTC when the object expires and after that the object is no longer available. The default and also maximum TTL is 12 hours after the creation, but it might change for performance optimization.
// returns a *Time when successful
func (m *CloudClipboardItem) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudClipboardItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["payloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudClipboardItemPayloadFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudClipboardItemPayloadable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudClipboardItemPayloadable)
                }
            }
            m.SetPayloads(res)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Set by the server if not provided in the client's request. DateTime in UTC when the object was modified by the client.
// returns a *Time when successful
func (m *CloudClipboardItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPayloads gets the payloads property value. A cloudClipboardItem can have multiple cloudClipboardItemPayload objects in the payloads. A window can place more than one clipboard object on the clipboard. Each one represents the same information in a different clipboard format.
// returns a []CloudClipboardItemPayloadable when successful
func (m *CloudClipboardItem) GetPayloads()([]CloudClipboardItemPayloadable) {
    return m.payloads
}
// Serialize serializes information the current object
func (m *CloudClipboardItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPayloads() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPayloads()))
        for i, v := range m.GetPayloads() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("payloads", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Set by the server. DateTime in UTC when the object was created on the server.
func (m *CloudClipboardItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Set by the server. DateTime in UTC when the object expires and after that the object is no longer available. The default and also maximum TTL is 12 hours after the creation, but it might change for performance optimization.
func (m *CloudClipboardItem) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Set by the server if not provided in the client's request. DateTime in UTC when the object was modified by the client.
func (m *CloudClipboardItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPayloads sets the payloads property value. A cloudClipboardItem can have multiple cloudClipboardItemPayload objects in the payloads. A window can place more than one clipboard object on the clipboard. Each one represents the same information in a different clipboard format.
func (m *CloudClipboardItem) SetPayloads(value []CloudClipboardItemPayloadable)() {
    m.payloads = value
}
type CloudClipboardItemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPayloads()([]CloudClipboardItemPayloadable)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPayloads(value []CloudClipboardItemPayloadable)()
}
