package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CalendarSharingMessage struct {
    Message
    // The canAccept property
    canAccept *bool
    // The sharingMessageAction property
    sharingMessageAction CalendarSharingMessageActionable
    // The sharingMessageActions property
    sharingMessageActions []CalendarSharingMessageActionable
    // The suggestedCalendarName property
    suggestedCalendarName *string
}
// NewCalendarSharingMessage instantiates a new CalendarSharingMessage and sets the default values.
func NewCalendarSharingMessage()(*CalendarSharingMessage) {
    m := &CalendarSharingMessage{
        Message: *NewMessage(),
    }
    odataTypeValue := "#microsoft.graph.calendarSharingMessage"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCalendarSharingMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCalendarSharingMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalendarSharingMessage(), nil
}
// GetCanAccept gets the canAccept property value. The canAccept property
// returns a *bool when successful
func (m *CalendarSharingMessage) GetCanAccept()(*bool) {
    return m.canAccept
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CalendarSharingMessage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Message.GetFieldDeserializers()
    res["canAccept"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanAccept(val)
        }
        return nil
    }
    res["sharingMessageAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCalendarSharingMessageActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharingMessageAction(val.(CalendarSharingMessageActionable))
        }
        return nil
    }
    res["sharingMessageActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCalendarSharingMessageActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CalendarSharingMessageActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CalendarSharingMessageActionable)
                }
            }
            m.SetSharingMessageActions(res)
        }
        return nil
    }
    res["suggestedCalendarName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestedCalendarName(val)
        }
        return nil
    }
    return res
}
// GetSharingMessageAction gets the sharingMessageAction property value. The sharingMessageAction property
// returns a CalendarSharingMessageActionable when successful
func (m *CalendarSharingMessage) GetSharingMessageAction()(CalendarSharingMessageActionable) {
    return m.sharingMessageAction
}
// GetSharingMessageActions gets the sharingMessageActions property value. The sharingMessageActions property
// returns a []CalendarSharingMessageActionable when successful
func (m *CalendarSharingMessage) GetSharingMessageActions()([]CalendarSharingMessageActionable) {
    return m.sharingMessageActions
}
// GetSuggestedCalendarName gets the suggestedCalendarName property value. The suggestedCalendarName property
// returns a *string when successful
func (m *CalendarSharingMessage) GetSuggestedCalendarName()(*string) {
    return m.suggestedCalendarName
}
// Serialize serializes information the current object
func (m *CalendarSharingMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Message.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("canAccept", m.GetCanAccept())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharingMessageAction", m.GetSharingMessageAction())
        if err != nil {
            return err
        }
    }
    if m.GetSharingMessageActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharingMessageActions()))
        for i, v := range m.GetSharingMessageActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sharingMessageActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("suggestedCalendarName", m.GetSuggestedCalendarName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCanAccept sets the canAccept property value. The canAccept property
func (m *CalendarSharingMessage) SetCanAccept(value *bool)() {
    m.canAccept = value
}
// SetSharingMessageAction sets the sharingMessageAction property value. The sharingMessageAction property
func (m *CalendarSharingMessage) SetSharingMessageAction(value CalendarSharingMessageActionable)() {
    m.sharingMessageAction = value
}
// SetSharingMessageActions sets the sharingMessageActions property value. The sharingMessageActions property
func (m *CalendarSharingMessage) SetSharingMessageActions(value []CalendarSharingMessageActionable)() {
    m.sharingMessageActions = value
}
// SetSuggestedCalendarName sets the suggestedCalendarName property value. The suggestedCalendarName property
func (m *CalendarSharingMessage) SetSuggestedCalendarName(value *string)() {
    m.suggestedCalendarName = value
}
type CalendarSharingMessageable interface {
    Messageable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanAccept()(*bool)
    GetSharingMessageAction()(CalendarSharingMessageActionable)
    GetSharingMessageActions()([]CalendarSharingMessageActionable)
    GetSuggestedCalendarName()(*string)
    SetCanAccept(value *bool)()
    SetSharingMessageAction(value CalendarSharingMessageActionable)()
    SetSharingMessageActions(value []CalendarSharingMessageActionable)()
    SetSuggestedCalendarName(value *string)()
}
