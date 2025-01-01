package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HostCookie struct {
    Artifact
    // The URI for which the cookie is valid.
    domain *string
    // The first date and time when this hostCookie was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The host property
    host Hostable
    // The most recent date and time when this hostCookie was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the cookie, for example, JSESSIONID or SEARCH_NAMESITE.
    name *string
}
// NewHostCookie instantiates a new HostCookie and sets the default values.
func NewHostCookie()(*HostCookie) {
    m := &HostCookie{
        Artifact: *NewArtifact(),
    }
    odataTypeValue := "#microsoft.graph.security.hostCookie"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateHostCookieFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHostCookieFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostCookie(), nil
}
// GetDomain gets the domain property value. The URI for which the cookie is valid.
// returns a *string when successful
func (m *HostCookie) GetDomain()(*string) {
    return m.domain
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HostCookie) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Artifact.GetFieldDeserializers()
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The first date and time when this hostCookie was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HostCookie) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetHost gets the host property value. The host property
// returns a Hostable when successful
func (m *HostCookie) GetHost()(Hostable) {
    return m.host
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The most recent date and time when this hostCookie was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HostCookie) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetName gets the name property value. The name of the cookie, for example, JSESSIONID or SEARCH_NAMESITE.
// returns a *string when successful
func (m *HostCookie) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *HostCookie) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Artifact.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
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
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDomain sets the domain property value. The URI for which the cookie is valid.
func (m *HostCookie) SetDomain(value *string)() {
    m.domain = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The first date and time when this hostCookie was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *HostCookie) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetHost sets the host property value. The host property
func (m *HostCookie) SetHost(value Hostable)() {
    m.host = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The most recent date and time when this hostCookie was observed by Microsoft Defender Threat Intelligence. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *HostCookie) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetName sets the name property value. The name of the cookie, for example, JSESSIONID or SEARCH_NAMESITE.
func (m *HostCookie) SetName(value *string)() {
    m.name = value
}
type HostCookieable interface {
    Artifactable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDomain()(*string)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHost()(Hostable)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    SetDomain(value *string)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHost(value Hostable)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
}
