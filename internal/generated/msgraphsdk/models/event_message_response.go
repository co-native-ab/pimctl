package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EventMessageResponse struct {
    EventMessage
    // An alternate date/time proposed by an invitee for a meeting request to start and end. Read-only. Not filterable.
    proposedNewTime TimeSlotable
    // Specifies the type of response to a meeting request. Possible values are: tentativelyAccepted, accepted, declined. For the eventMessageResponse type, none, organizer, and notResponded are not supported. Read-only. Not filterable.
    responseType *ResponseType
}
// NewEventMessageResponse instantiates a new EventMessageResponse and sets the default values.
func NewEventMessageResponse()(*EventMessageResponse) {
    m := &EventMessageResponse{
        EventMessage: *NewEventMessage(),
    }
    odataTypeValue := "#microsoft.graph.eventMessageResponse"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEventMessageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEventMessageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEventMessageResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EventMessageResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessage.GetFieldDeserializers()
    res["proposedNewTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposedNewTime(val.(TimeSlotable))
        }
        return nil
    }
    res["responseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResponseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseType(val.(*ResponseType))
        }
        return nil
    }
    return res
}
// GetProposedNewTime gets the proposedNewTime property value. An alternate date/time proposed by an invitee for a meeting request to start and end. Read-only. Not filterable.
// returns a TimeSlotable when successful
func (m *EventMessageResponse) GetProposedNewTime()(TimeSlotable) {
    return m.proposedNewTime
}
// GetResponseType gets the responseType property value. Specifies the type of response to a meeting request. Possible values are: tentativelyAccepted, accepted, declined. For the eventMessageResponse type, none, organizer, and notResponded are not supported. Read-only. Not filterable.
// returns a *ResponseType when successful
func (m *EventMessageResponse) GetResponseType()(*ResponseType) {
    return m.responseType
}
// Serialize serializes information the current object
func (m *EventMessageResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessage.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("proposedNewTime", m.GetProposedNewTime())
        if err != nil {
            return err
        }
    }
    if m.GetResponseType() != nil {
        cast := (*m.GetResponseType()).String()
        err = writer.WriteStringValue("responseType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetProposedNewTime sets the proposedNewTime property value. An alternate date/time proposed by an invitee for a meeting request to start and end. Read-only. Not filterable.
func (m *EventMessageResponse) SetProposedNewTime(value TimeSlotable)() {
    m.proposedNewTime = value
}
// SetResponseType sets the responseType property value. Specifies the type of response to a meeting request. Possible values are: tentativelyAccepted, accepted, declined. For the eventMessageResponse type, none, organizer, and notResponded are not supported. Read-only. Not filterable.
func (m *EventMessageResponse) SetResponseType(value *ResponseType)() {
    m.responseType = value
}
type EventMessageResponseable interface {
    EventMessageable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProposedNewTime()(TimeSlotable)
    GetResponseType()(*ResponseType)
    SetProposedNewTime(value TimeSlotable)()
    SetResponseType(value *ResponseType)()
}
