package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type CallRecord struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // UTC time when the last user left the call. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Meeting URL associated to the call. May not be available for a peerToPeer call record type.
    joinWebUrl *string
    // UTC time when the call record was created. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of all the modalities used in the call. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
    modalities []Modality
    // The organizing party's identity. The organizer property is deprecated and will stop returning data on June 30, 2026. Going forward, use the organizer_v2 relationship.
    organizer ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable
    // Identity of the organizer of the call. This relationship is expanded by default in callRecord methods.
    organizer_v2 Organizerable
    // List of distinct identities involved in the call. Limited to 130 entries. The participants property is deprecated and will stop returning data on June 30, 2026. Going forward, use the participants_v2 relationship.
    participants []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable
    // List of distinct participants in the call.
    participants_v2 []Participantable
    // List of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
    sessions []Sessionable
    // UTC time when the first user joined the call. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type property
    typeEscaped *CallType
    // Monotonically increasing version of the call record. Higher version call records with the same id includes additional data compared to the lower version.
    version *int64
}
// NewCallRecord instantiates a new CallRecord and sets the default values.
func NewCallRecord()(*CallRecord) {
    m := &CallRecord{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateCallRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCallRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallRecord(), nil
}
// GetEndDateTime gets the endDateTime property value. UTC time when the last user left the call. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *CallRecord) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CallRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["joinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
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
    res["modalities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Modality, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*Modality))
                }
            }
            m.SetModalities(res)
        }
        return nil
    }
    res["organizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizer(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable))
        }
        return nil
    }
    res["organizer_v2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizerV2(val.(Organizerable))
        }
        return nil
    }
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)
                }
            }
            m.SetParticipants(res)
        }
        return nil
    }
    res["participants_v2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateParticipantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Participantable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Participantable)
                }
            }
            m.SetParticipantsV2(res)
        }
        return nil
    }
    res["sessions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSessionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Sessionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Sessionable)
                }
            }
            m.SetSessions(res)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*CallType))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetJoinWebUrl gets the joinWebUrl property value. Meeting URL associated to the call. May not be available for a peerToPeer call record type.
// returns a *string when successful
func (m *CallRecord) GetJoinWebUrl()(*string) {
    return m.joinWebUrl
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. UTC time when the call record was created. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *CallRecord) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetModalities gets the modalities property value. List of all the modalities used in the call. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
// returns a []Modality when successful
func (m *CallRecord) GetModalities()([]Modality) {
    return m.modalities
}
// GetOrganizer gets the organizer property value. The organizing party's identity. The organizer property is deprecated and will stop returning data on June 30, 2026. Going forward, use the organizer_v2 relationship.
// returns a IdentitySetable when successful
func (m *CallRecord) GetOrganizer()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable) {
    return m.organizer
}
// GetOrganizerV2 gets the organizer_v2 property value. Identity of the organizer of the call. This relationship is expanded by default in callRecord methods.
// returns a Organizerable when successful
func (m *CallRecord) GetOrganizerV2()(Organizerable) {
    return m.organizer_v2
}
// GetParticipants gets the participants property value. List of distinct identities involved in the call. Limited to 130 entries. The participants property is deprecated and will stop returning data on June 30, 2026. Going forward, use the participants_v2 relationship.
// returns a []IdentitySetable when successful
func (m *CallRecord) GetParticipants()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable) {
    return m.participants
}
// GetParticipantsV2 gets the participants_v2 property value. List of distinct participants in the call.
// returns a []Participantable when successful
func (m *CallRecord) GetParticipantsV2()([]Participantable) {
    return m.participants_v2
}
// GetSessions gets the sessions property value. List of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
// returns a []Sessionable when successful
func (m *CallRecord) GetSessions()([]Sessionable) {
    return m.sessions
}
// GetStartDateTime gets the startDateTime property value. UTC time when the first user joined the call. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *CallRecord) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetTypeEscaped gets the type property value. The type property
// returns a *CallType when successful
func (m *CallRecord) GetTypeEscaped()(*CallType) {
    return m.typeEscaped
}
// GetVersion gets the version property value. Monotonically increasing version of the call record. Higher version call records with the same id includes additional data compared to the lower version.
// returns a *int64 when successful
func (m *CallRecord) GetVersion()(*int64) {
    return m.version
}
// Serialize serializes information the current object
func (m *CallRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
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
    if m.GetModalities() != nil {
        err = writer.WriteCollectionOfStringValues("modalities", SerializeModality(m.GetModalities()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("organizer", m.GetOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("organizer_v2", m.GetOrganizerV2())
        if err != nil {
            return err
        }
    }
    if m.GetParticipants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParticipants()))
        for i, v := range m.GetParticipants() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("participants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetParticipantsV2() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParticipantsV2()))
        for i, v := range m.GetParticipantsV2() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("participants_v2", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSessions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSessions()))
        for i, v := range m.GetSessions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sessions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndDateTime sets the endDateTime property value. UTC time when the last user left the call. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *CallRecord) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetJoinWebUrl sets the joinWebUrl property value. Meeting URL associated to the call. May not be available for a peerToPeer call record type.
func (m *CallRecord) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. UTC time when the call record was created. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *CallRecord) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetModalities sets the modalities property value. List of all the modalities used in the call. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
func (m *CallRecord) SetModalities(value []Modality)() {
    m.modalities = value
}
// SetOrganizer sets the organizer property value. The organizing party's identity. The organizer property is deprecated and will stop returning data on June 30, 2026. Going forward, use the organizer_v2 relationship.
func (m *CallRecord) SetOrganizer(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)() {
    m.organizer = value
}
// SetOrganizerV2 sets the organizer_v2 property value. Identity of the organizer of the call. This relationship is expanded by default in callRecord methods.
func (m *CallRecord) SetOrganizerV2(value Organizerable)() {
    m.organizer_v2 = value
}
// SetParticipants sets the participants property value. List of distinct identities involved in the call. Limited to 130 entries. The participants property is deprecated and will stop returning data on June 30, 2026. Going forward, use the participants_v2 relationship.
func (m *CallRecord) SetParticipants(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)() {
    m.participants = value
}
// SetParticipantsV2 sets the participants_v2 property value. List of distinct participants in the call.
func (m *CallRecord) SetParticipantsV2(value []Participantable)() {
    m.participants_v2 = value
}
// SetSessions sets the sessions property value. List of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
func (m *CallRecord) SetSessions(value []Sessionable)() {
    m.sessions = value
}
// SetStartDateTime sets the startDateTime property value. UTC time when the first user joined the call. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CallRecord) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *CallRecord) SetTypeEscaped(value *CallType)() {
    m.typeEscaped = value
}
// SetVersion sets the version property value. Monotonically increasing version of the call record. Higher version call records with the same id includes additional data compared to the lower version.
func (m *CallRecord) SetVersion(value *int64)() {
    m.version = value
}
type CallRecordable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetJoinWebUrl()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetModalities()([]Modality)
    GetOrganizer()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)
    GetOrganizerV2()(Organizerable)
    GetParticipants()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)
    GetParticipantsV2()([]Participantable)
    GetSessions()([]Sessionable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTypeEscaped()(*CallType)
    GetVersion()(*int64)
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetJoinWebUrl(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetModalities(value []Modality)()
    SetOrganizer(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)()
    SetOrganizerV2(value Organizerable)()
    SetParticipants(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)()
    SetParticipantsV2(value []Participantable)()
    SetSessions(value []Sessionable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTypeEscaped(value *CallType)()
    SetVersion(value *int64)()
}
