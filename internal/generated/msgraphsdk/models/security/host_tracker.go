package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HostTracker struct {
    Artifact
    // The first date and time when this hostTracker was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The host property
    host Hostable
    // The kind of hostTracker that was detected. For example, GoogleAnalyticsID or JarmHash.
    kind *string
    // The most recent date and time when this hostTracker was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The identification value for the hostTracker.
    value *string
}
// NewHostTracker instantiates a new HostTracker and sets the default values.
func NewHostTracker()(*HostTracker) {
    m := &HostTracker{
        Artifact: *NewArtifact(),
    }
    odataTypeValue := "#microsoft.graph.security.hostTracker"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateHostTrackerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHostTrackerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostTracker(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HostTracker) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Artifact.GetFieldDeserializers()
    res["firstSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["host"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHost(val.(Hostable))
        }
        return nil
    }
    res["kind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKind(val)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The first date and time when this hostTracker was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HostTracker) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetHost gets the host property value. The host property
// returns a Hostable when successful
func (m *HostTracker) GetHost()(Hostable) {
    return m.host
}
// GetKind gets the kind property value. The kind of hostTracker that was detected. For example, GoogleAnalyticsID or JarmHash.
// returns a *string when successful
func (m *HostTracker) GetKind()(*string) {
    return m.kind
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The most recent date and time when this hostTracker was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HostTracker) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetValue gets the value property value. The identification value for the hostTracker.
// returns a *string when successful
func (m *HostTracker) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *HostTracker) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Artifact.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("host", m.GetHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kind", m.GetKind())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The first date and time when this hostTracker was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *HostTracker) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetHost sets the host property value. The host property
func (m *HostTracker) SetHost(value Hostable)() {
    m.host = value
}
// SetKind sets the kind property value. The kind of hostTracker that was detected. For example, GoogleAnalyticsID or JarmHash.
func (m *HostTracker) SetKind(value *string)() {
    m.kind = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The most recent date and time when this hostTracker was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *HostTracker) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetValue sets the value property value. The identification value for the hostTracker.
func (m *HostTracker) SetValue(value *string)() {
    m.value = value
}
type HostTrackerable interface {
    Artifactable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHost()(Hostable)
    GetKind()(*string)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetValue()(*string)
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHost(value Hostable)()
    SetKind(value *string)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetValue(value *string)()
}
