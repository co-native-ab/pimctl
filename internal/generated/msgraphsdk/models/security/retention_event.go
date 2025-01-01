package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type RetentionEvent struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The user who created the retentionEvent.
    createdBy ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable
    // The date time when the retentionEvent was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Optional information about the event.
    description *string
    // Name of the event.
    displayName *string
    // Represents the success status of a created event and additional information.
    eventPropagationResults []EventPropagationResultable
    // Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information associated with a retention event.
    eventQueries []EventQueryable
    // Status of event propogation to the scoped locations after the event has been created.
    eventStatus RetentionEventStatusable
    // Optional time when the event should be triggered.
    eventTriggerDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user who last modified the retentionEvent.
    lastModifiedBy ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable
    // The latest date time when the retentionEvent was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last time the status of the event was updated.
    lastStatusUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the event that will start the retention period for labels that use this event type when an event is created.
    retentionEventType RetentionEventTypeable
}
// NewRetentionEvent instantiates a new RetentionEvent and sets the default values.
func NewRetentionEvent()(*RetentionEvent) {
    m := &RetentionEvent{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateRetentionEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRetentionEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionEvent(), nil
}
// GetCreatedBy gets the createdBy property value. The user who created the retentionEvent.
// returns a IdentitySetable when successful
func (m *RetentionEvent) GetCreatedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date time when the retentionEvent was created.
// returns a *Time when successful
func (m *RetentionEvent) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Optional information about the event.
// returns a *string when successful
func (m *RetentionEvent) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the event.
// returns a *string when successful
func (m *RetentionEvent) GetDisplayName()(*string) {
    return m.displayName
}
// GetEventPropagationResults gets the eventPropagationResults property value. Represents the success status of a created event and additional information.
// returns a []EventPropagationResultable when successful
func (m *RetentionEvent) GetEventPropagationResults()([]EventPropagationResultable) {
    return m.eventPropagationResults
}
// GetEventQueries gets the eventQueries property value. Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information associated with a retention event.
// returns a []EventQueryable when successful
func (m *RetentionEvent) GetEventQueries()([]EventQueryable) {
    return m.eventQueries
}
// GetEventStatus gets the eventStatus property value. Status of event propogation to the scoped locations after the event has been created.
// returns a RetentionEventStatusable when successful
func (m *RetentionEvent) GetEventStatus()(RetentionEventStatusable) {
    return m.eventStatus
}
// GetEventTriggerDateTime gets the eventTriggerDateTime property value. Optional time when the event should be triggered.
// returns a *Time when successful
func (m *RetentionEvent) GetEventTriggerDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventTriggerDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RetentionEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable))
        }
        return nil
    }
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["eventPropagationResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventPropagationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EventPropagationResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EventPropagationResultable)
                }
            }
            m.SetEventPropagationResults(res)
        }
        return nil
    }
    res["eventQueries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EventQueryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EventQueryable)
                }
            }
            m.SetEventQueries(res)
        }
        return nil
    }
    res["eventStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRetentionEventStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventStatus(val.(RetentionEventStatusable))
        }
        return nil
    }
    res["eventTriggerDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventTriggerDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable))
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
    res["lastStatusUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastStatusUpdateDateTime(val)
        }
        return nil
    }
    res["retentionEventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRetentionEventTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetentionEventType(val.(RetentionEventTypeable))
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the retentionEvent.
// returns a IdentitySetable when successful
func (m *RetentionEvent) GetLastModifiedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The latest date time when the retentionEvent was modified.
// returns a *Time when successful
func (m *RetentionEvent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLastStatusUpdateDateTime gets the lastStatusUpdateDateTime property value. Last time the status of the event was updated.
// returns a *Time when successful
func (m *RetentionEvent) GetLastStatusUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastStatusUpdateDateTime
}
// GetRetentionEventType gets the retentionEventType property value. Specifies the event that will start the retention period for labels that use this event type when an event is created.
// returns a RetentionEventTypeable when successful
func (m *RetentionEvent) GetRetentionEventType()(RetentionEventTypeable) {
    return m.retentionEventType
}
// Serialize serializes information the current object
func (m *RetentionEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetEventPropagationResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEventPropagationResults()))
        for i, v := range m.GetEventPropagationResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("eventPropagationResults", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEventQueries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEventQueries()))
        for i, v := range m.GetEventQueries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("eventQueries", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("eventStatus", m.GetEventStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventTriggerDateTime", m.GetEventTriggerDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
    {
        err = writer.WriteTimeValue("lastStatusUpdateDateTime", m.GetLastStatusUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("retentionEventType", m.GetRetentionEventType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The user who created the retentionEvent.
func (m *RetentionEvent) SetCreatedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date time when the retentionEvent was created.
func (m *RetentionEvent) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Optional information about the event.
func (m *RetentionEvent) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the event.
func (m *RetentionEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEventPropagationResults sets the eventPropagationResults property value. Represents the success status of a created event and additional information.
func (m *RetentionEvent) SetEventPropagationResults(value []EventPropagationResultable)() {
    m.eventPropagationResults = value
}
// SetEventQueries sets the eventQueries property value. Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information associated with a retention event.
func (m *RetentionEvent) SetEventQueries(value []EventQueryable)() {
    m.eventQueries = value
}
// SetEventStatus sets the eventStatus property value. Status of event propogation to the scoped locations after the event has been created.
func (m *RetentionEvent) SetEventStatus(value RetentionEventStatusable)() {
    m.eventStatus = value
}
// SetEventTriggerDateTime sets the eventTriggerDateTime property value. Optional time when the event should be triggered.
func (m *RetentionEvent) SetEventTriggerDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventTriggerDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the retentionEvent.
func (m *RetentionEvent) SetLastModifiedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The latest date time when the retentionEvent was modified.
func (m *RetentionEvent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLastStatusUpdateDateTime sets the lastStatusUpdateDateTime property value. Last time the status of the event was updated.
func (m *RetentionEvent) SetLastStatusUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastStatusUpdateDateTime = value
}
// SetRetentionEventType sets the retentionEventType property value. Specifies the event that will start the retention period for labels that use this event type when an event is created.
func (m *RetentionEvent) SetRetentionEventType(value RetentionEventTypeable)() {
    m.retentionEventType = value
}
type RetentionEventable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEventPropagationResults()([]EventPropagationResultable)
    GetEventQueries()([]EventQueryable)
    GetEventStatus()(RetentionEventStatusable)
    GetEventTriggerDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastStatusUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRetentionEventType()(RetentionEventTypeable)
    SetCreatedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEventPropagationResults(value []EventPropagationResultable)()
    SetEventQueries(value []EventQueryable)()
    SetEventStatus(value RetentionEventStatusable)()
    SetEventTriggerDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastStatusUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRetentionEventType(value RetentionEventTypeable)()
}
