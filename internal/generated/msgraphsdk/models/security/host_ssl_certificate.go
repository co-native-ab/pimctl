package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HostSslCertificate struct {
    Artifact
    // The first date and time when this hostSslCertificate was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The host for this hostSslCertificate.
    host Hostable
    // The most recent date and time when this hostSslCertificate was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ports related with this hostSslCertificate.
    ports []HostSslCertificatePortable
    // The sslCertificate for this hostSslCertificate.
    sslCertificate SslCertificateable
}
// NewHostSslCertificate instantiates a new HostSslCertificate and sets the default values.
func NewHostSslCertificate()(*HostSslCertificate) {
    m := &HostSslCertificate{
        Artifact: *NewArtifact(),
    }
    odataTypeValue := "#microsoft.graph.security.hostSslCertificate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateHostSslCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHostSslCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostSslCertificate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HostSslCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["ports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostSslCertificatePortFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostSslCertificatePortable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostSslCertificatePortable)
                }
            }
            m.SetPorts(res)
        }
        return nil
    }
    res["sslCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSslCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSslCertificate(val.(SslCertificateable))
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The first date and time when this hostSslCertificate was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HostSslCertificate) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetHost gets the host property value. The host for this hostSslCertificate.
// returns a Hostable when successful
func (m *HostSslCertificate) GetHost()(Hostable) {
    return m.host
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The most recent date and time when this hostSslCertificate was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HostSslCertificate) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetPorts gets the ports property value. The ports related with this hostSslCertificate.
// returns a []HostSslCertificatePortable when successful
func (m *HostSslCertificate) GetPorts()([]HostSslCertificatePortable) {
    return m.ports
}
// GetSslCertificate gets the sslCertificate property value. The sslCertificate for this hostSslCertificate.
// returns a SslCertificateable when successful
func (m *HostSslCertificate) GetSslCertificate()(SslCertificateable) {
    return m.sslCertificate
}
// Serialize serializes information the current object
func (m *HostSslCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPorts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPorts()))
        for i, v := range m.GetPorts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("ports", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sslCertificate", m.GetSslCertificate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The first date and time when this hostSslCertificate was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *HostSslCertificate) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetHost sets the host property value. The host for this hostSslCertificate.
func (m *HostSslCertificate) SetHost(value Hostable)() {
    m.host = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The most recent date and time when this hostSslCertificate was observed. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *HostSslCertificate) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetPorts sets the ports property value. The ports related with this hostSslCertificate.
func (m *HostSslCertificate) SetPorts(value []HostSslCertificatePortable)() {
    m.ports = value
}
// SetSslCertificate sets the sslCertificate property value. The sslCertificate for this hostSslCertificate.
func (m *HostSslCertificate) SetSslCertificate(value SslCertificateable)() {
    m.sslCertificate = value
}
type HostSslCertificateable interface {
    Artifactable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHost()(Hostable)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPorts()([]HostSslCertificatePortable)
    GetSslCertificate()(SslCertificateable)
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHost(value Hostable)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPorts(value []HostSslCertificatePortable)()
    SetSslCertificate(value SslCertificateable)()
}
