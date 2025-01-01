package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Participant struct {
    Entity
    // The info property
    info ParticipantInfoable
    // true if the participant is in lobby.
    isInLobby *bool
    // true if the participant is muted (client or server muted).
    isMuted *bool
    // The list of media streams.
    mediaStreams []MediaStreamable
    // A blob of data provided by the participant in the roster.
    metadata *string
    // Information about whether the participant has recording capability.
    recordingInfo RecordingInfoable
    // Indicates the reason why the participant was removed from the roster.
    removedState RemovedStateable
    // Indicates the reason or reasons media content from this participant is restricted.
    restrictedExperience OnlineMeetingRestrictedable
    // Indicates the roster sequence number in which the participant was last updated.
    rosterSequenceNumber *int64
}
// NewParticipant instantiates a new Participant and sets the default values.
func NewParticipant()(*Participant) {
    m := &Participant{
        Entity: *NewEntity(),
    }
    return m
}
// CreateParticipantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParticipantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipant(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Participant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["info"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfo(val.(ParticipantInfoable))
        }
        return nil
    }
    res["isInLobby"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInLobby(val)
        }
        return nil
    }
    res["isMuted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMuted(val)
        }
        return nil
    }
    res["mediaStreams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMediaStreamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MediaStreamable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MediaStreamable)
                }
            }
            m.SetMediaStreams(res)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val)
        }
        return nil
    }
    res["recordingInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecordingInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordingInfo(val.(RecordingInfoable))
        }
        return nil
    }
    res["removedState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRemovedStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemovedState(val.(RemovedStateable))
        }
        return nil
    }
    res["restrictedExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnlineMeetingRestrictedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictedExperience(val.(OnlineMeetingRestrictedable))
        }
        return nil
    }
    res["rosterSequenceNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRosterSequenceNumber(val)
        }
        return nil
    }
    return res
}
// GetInfo gets the info property value. The info property
// returns a ParticipantInfoable when successful
func (m *Participant) GetInfo()(ParticipantInfoable) {
    return m.info
}
// GetIsInLobby gets the isInLobby property value. true if the participant is in lobby.
// returns a *bool when successful
func (m *Participant) GetIsInLobby()(*bool) {
    return m.isInLobby
}
// GetIsMuted gets the isMuted property value. true if the participant is muted (client or server muted).
// returns a *bool when successful
func (m *Participant) GetIsMuted()(*bool) {
    return m.isMuted
}
// GetMediaStreams gets the mediaStreams property value. The list of media streams.
// returns a []MediaStreamable when successful
func (m *Participant) GetMediaStreams()([]MediaStreamable) {
    return m.mediaStreams
}
// GetMetadata gets the metadata property value. A blob of data provided by the participant in the roster.
// returns a *string when successful
func (m *Participant) GetMetadata()(*string) {
    return m.metadata
}
// GetRecordingInfo gets the recordingInfo property value. Information about whether the participant has recording capability.
// returns a RecordingInfoable when successful
func (m *Participant) GetRecordingInfo()(RecordingInfoable) {
    return m.recordingInfo
}
// GetRemovedState gets the removedState property value. Indicates the reason why the participant was removed from the roster.
// returns a RemovedStateable when successful
func (m *Participant) GetRemovedState()(RemovedStateable) {
    return m.removedState
}
// GetRestrictedExperience gets the restrictedExperience property value. Indicates the reason or reasons media content from this participant is restricted.
// returns a OnlineMeetingRestrictedable when successful
func (m *Participant) GetRestrictedExperience()(OnlineMeetingRestrictedable) {
    return m.restrictedExperience
}
// GetRosterSequenceNumber gets the rosterSequenceNumber property value. Indicates the roster sequence number in which the participant was last updated.
// returns a *int64 when successful
func (m *Participant) GetRosterSequenceNumber()(*int64) {
    return m.rosterSequenceNumber
}
// Serialize serializes information the current object
func (m *Participant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("info", m.GetInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInLobby", m.GetIsInLobby())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMuted", m.GetIsMuted())
        if err != nil {
            return err
        }
    }
    if m.GetMediaStreams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMediaStreams()))
        for i, v := range m.GetMediaStreams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mediaStreams", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recordingInfo", m.GetRecordingInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("removedState", m.GetRemovedState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("restrictedExperience", m.GetRestrictedExperience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("rosterSequenceNumber", m.GetRosterSequenceNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInfo sets the info property value. The info property
func (m *Participant) SetInfo(value ParticipantInfoable)() {
    m.info = value
}
// SetIsInLobby sets the isInLobby property value. true if the participant is in lobby.
func (m *Participant) SetIsInLobby(value *bool)() {
    m.isInLobby = value
}
// SetIsMuted sets the isMuted property value. true if the participant is muted (client or server muted).
func (m *Participant) SetIsMuted(value *bool)() {
    m.isMuted = value
}
// SetMediaStreams sets the mediaStreams property value. The list of media streams.
func (m *Participant) SetMediaStreams(value []MediaStreamable)() {
    m.mediaStreams = value
}
// SetMetadata sets the metadata property value. A blob of data provided by the participant in the roster.
func (m *Participant) SetMetadata(value *string)() {
    m.metadata = value
}
// SetRecordingInfo sets the recordingInfo property value. Information about whether the participant has recording capability.
func (m *Participant) SetRecordingInfo(value RecordingInfoable)() {
    m.recordingInfo = value
}
// SetRemovedState sets the removedState property value. Indicates the reason why the participant was removed from the roster.
func (m *Participant) SetRemovedState(value RemovedStateable)() {
    m.removedState = value
}
// SetRestrictedExperience sets the restrictedExperience property value. Indicates the reason or reasons media content from this participant is restricted.
func (m *Participant) SetRestrictedExperience(value OnlineMeetingRestrictedable)() {
    m.restrictedExperience = value
}
// SetRosterSequenceNumber sets the rosterSequenceNumber property value. Indicates the roster sequence number in which the participant was last updated.
func (m *Participant) SetRosterSequenceNumber(value *int64)() {
    m.rosterSequenceNumber = value
}
type Participantable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInfo()(ParticipantInfoable)
    GetIsInLobby()(*bool)
    GetIsMuted()(*bool)
    GetMediaStreams()([]MediaStreamable)
    GetMetadata()(*string)
    GetRecordingInfo()(RecordingInfoable)
    GetRemovedState()(RemovedStateable)
    GetRestrictedExperience()(OnlineMeetingRestrictedable)
    GetRosterSequenceNumber()(*int64)
    SetInfo(value ParticipantInfoable)()
    SetIsInLobby(value *bool)()
    SetIsMuted(value *bool)()
    SetMediaStreams(value []MediaStreamable)()
    SetMetadata(value *string)()
    SetRecordingInfo(value RecordingInfoable)()
    SetRemovedState(value RemovedStateable)()
    SetRestrictedExperience(value OnlineMeetingRestrictedable)()
    SetRosterSequenceNumber(value *int64)()
}
