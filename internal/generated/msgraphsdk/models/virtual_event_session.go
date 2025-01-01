package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEventSession struct {
    OnlineMeetingBase
    // The virtual event session end time.
    endDateTime DateTimeTimeZoneable
    // The virtual event session start time.
    startDateTime DateTimeTimeZoneable
}
// NewVirtualEventSession instantiates a new VirtualEventSession and sets the default values.
func NewVirtualEventSession()(*VirtualEventSession) {
    m := &VirtualEventSession{
        OnlineMeetingBase: *NewOnlineMeetingBase(),
    }
    odataTypeValue := "#microsoft.graph.virtualEventSession"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateVirtualEventSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventSession(), nil
}
// GetEndDateTime gets the endDateTime property value. The virtual event session end time.
// returns a DateTimeTimeZoneable when successful
func (m *VirtualEventSession) GetEndDateTime()(DateTimeTimeZoneable) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEventSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnlineMeetingBase.GetFieldDeserializers()
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetStartDateTime gets the startDateTime property value. The virtual event session start time.
// returns a DateTimeTimeZoneable when successful
func (m *VirtualEventSession) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *VirtualEventSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnlineMeetingBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndDateTime sets the endDateTime property value. The virtual event session end time.
func (m *VirtualEventSession) SetEndDateTime(value DateTimeTimeZoneable)() {
    m.endDateTime = value
}
// SetStartDateTime sets the startDateTime property value. The virtual event session start time.
func (m *VirtualEventSession) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
type VirtualEventSessionable interface {
    OnlineMeetingBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDateTime()(DateTimeTimeZoneable)
    GetStartDateTime()(DateTimeTimeZoneable)
    SetEndDateTime(value DateTimeTimeZoneable)()
    SetStartDateTime(value DateTimeTimeZoneable)()
}
