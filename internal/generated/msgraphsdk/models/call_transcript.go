package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CallTranscript struct {
    Entity
    // The unique identifier for the call that is related to this transcript. Read-only.
    callId *string
    // The content of the transcript. Read-only.
    content []byte
    // The unique identifier that links the transcript with its corresponding recording. Read-only.
    contentCorrelationId *string
    // Date and time at which the transcript was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date and time at which the transcription ends. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the online meeting related to this transcript. Read-only.
    meetingId *string
    // The identity information of the organizer of the onlineMeeting related to this transcript. Read-only.
    meetingOrganizer IdentitySetable
    // The time-aligned metadata of the utterances in the transcript. Read-only.
    metadataContent []byte
    // The URL that can be used to access the content of the transcript. Read-only.
    transcriptContentUrl *string
}
// NewCallTranscript instantiates a new CallTranscript and sets the default values.
func NewCallTranscript()(*CallTranscript) {
    m := &CallTranscript{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCallTranscriptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCallTranscriptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallTranscript(), nil
}
// GetCallId gets the callId property value. The unique identifier for the call that is related to this transcript. Read-only.
// returns a *string when successful
func (m *CallTranscript) GetCallId()(*string) {
    return m.callId
}
// GetContent gets the content property value. The content of the transcript. Read-only.
// returns a []byte when successful
func (m *CallTranscript) GetContent()([]byte) {
    return m.content
}
// GetContentCorrelationId gets the contentCorrelationId property value. The unique identifier that links the transcript with its corresponding recording. Read-only.
// returns a *string when successful
func (m *CallTranscript) GetContentCorrelationId()(*string) {
    return m.contentCorrelationId
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time at which the transcript was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *CallTranscript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetEndDateTime gets the endDateTime property value. Date and time at which the transcription ends. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *CallTranscript) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CallTranscript) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["callId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallId(val)
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentCorrelationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCorrelationId(val)
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
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["meetingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingId(val)
        }
        return nil
    }
    res["meetingOrganizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingOrganizer(val.(IdentitySetable))
        }
        return nil
    }
    res["metadataContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadataContent(val)
        }
        return nil
    }
    res["transcriptContentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranscriptContentUrl(val)
        }
        return nil
    }
    return res
}
// GetMeetingId gets the meetingId property value. The unique identifier of the online meeting related to this transcript. Read-only.
// returns a *string when successful
func (m *CallTranscript) GetMeetingId()(*string) {
    return m.meetingId
}
// GetMeetingOrganizer gets the meetingOrganizer property value. The identity information of the organizer of the onlineMeeting related to this transcript. Read-only.
// returns a IdentitySetable when successful
func (m *CallTranscript) GetMeetingOrganizer()(IdentitySetable) {
    return m.meetingOrganizer
}
// GetMetadataContent gets the metadataContent property value. The time-aligned metadata of the utterances in the transcript. Read-only.
// returns a []byte when successful
func (m *CallTranscript) GetMetadataContent()([]byte) {
    return m.metadataContent
}
// GetTranscriptContentUrl gets the transcriptContentUrl property value. The URL that can be used to access the content of the transcript. Read-only.
// returns a *string when successful
func (m *CallTranscript) GetTranscriptContentUrl()(*string) {
    return m.transcriptContentUrl
}
// Serialize serializes information the current object
func (m *CallTranscript) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("callId", m.GetCallId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentCorrelationId", m.GetContentCorrelationId())
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
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("meetingId", m.GetMeetingId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("meetingOrganizer", m.GetMeetingOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("metadataContent", m.GetMetadataContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("transcriptContentUrl", m.GetTranscriptContentUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallId sets the callId property value. The unique identifier for the call that is related to this transcript. Read-only.
func (m *CallTranscript) SetCallId(value *string)() {
    m.callId = value
}
// SetContent sets the content property value. The content of the transcript. Read-only.
func (m *CallTranscript) SetContent(value []byte)() {
    m.content = value
}
// SetContentCorrelationId sets the contentCorrelationId property value. The unique identifier that links the transcript with its corresponding recording. Read-only.
func (m *CallTranscript) SetContentCorrelationId(value *string)() {
    m.contentCorrelationId = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time at which the transcript was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *CallTranscript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetEndDateTime sets the endDateTime property value. Date and time at which the transcription ends. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *CallTranscript) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetMeetingId sets the meetingId property value. The unique identifier of the online meeting related to this transcript. Read-only.
func (m *CallTranscript) SetMeetingId(value *string)() {
    m.meetingId = value
}
// SetMeetingOrganizer sets the meetingOrganizer property value. The identity information of the organizer of the onlineMeeting related to this transcript. Read-only.
func (m *CallTranscript) SetMeetingOrganizer(value IdentitySetable)() {
    m.meetingOrganizer = value
}
// SetMetadataContent sets the metadataContent property value. The time-aligned metadata of the utterances in the transcript. Read-only.
func (m *CallTranscript) SetMetadataContent(value []byte)() {
    m.metadataContent = value
}
// SetTranscriptContentUrl sets the transcriptContentUrl property value. The URL that can be used to access the content of the transcript. Read-only.
func (m *CallTranscript) SetTranscriptContentUrl(value *string)() {
    m.transcriptContentUrl = value
}
type CallTranscriptable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallId()(*string)
    GetContent()([]byte)
    GetContentCorrelationId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMeetingId()(*string)
    GetMeetingOrganizer()(IdentitySetable)
    GetMetadataContent()([]byte)
    GetTranscriptContentUrl()(*string)
    SetCallId(value *string)()
    SetContent(value []byte)()
    SetContentCorrelationId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMeetingId(value *string)()
    SetMeetingOrganizer(value IdentitySetable)()
    SetMetadataContent(value []byte)()
    SetTranscriptContentUrl(value *string)()
}
