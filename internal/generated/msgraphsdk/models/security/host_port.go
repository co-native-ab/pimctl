package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type HostPort struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The hostPortBanners retrieved from scanning the port.
    banners []HostPortBannerable
    // The first date and time when Microsoft Defender Threat Intelligence observed the hostPort. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The host property
    host Hostable
    // The last date and time when Microsoft Defender Threat Intelligence scanned the hostPort. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
    lastScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last date and time when Microsoft Defender Threat Intelligence observed the hostPort. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The most recent sslCertificate used to communicate on the port.
    mostRecentSslCertificate SslCertificateable
    // The numerical identifier of the port which is standardized across the internet.
    port *int32
    // The general protocol used to scan the port. The possible values are: tcp, udp, unknownFutureValue.
    protocol *HostPortProtocol
    // The hostPortComponents retrieved from scanning the port.
    services []HostPortComponentable
    // The status of the port. The possible values are: open, filtered, closed, unknownFutureValue.
    status *HostPortStatus
    // The total amount of times that Microsoft Defender Threat Intelligence has observed the hostPort in all its scans.
    timesObserved *int32
}
// NewHostPort instantiates a new HostPort and sets the default values.
func NewHostPort()(*HostPort) {
    m := &HostPort{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateHostPortFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHostPortFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostPort(), nil
}
// GetBanners gets the banners property value. The hostPortBanners retrieved from scanning the port.
// returns a []HostPortBannerable when successful
func (m *HostPort) GetBanners()([]HostPortBannerable) {
    return m.banners
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HostPort) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["banners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostPortBannerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostPortBannerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostPortBannerable)
                }
            }
            m.SetBanners(res)
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
    res["lastScanDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastScanDateTime(val)
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
    res["mostRecentSslCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSslCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMostRecentSslCertificate(val.(SslCertificateable))
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHostPortProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val.(*HostPortProtocol))
        }
        return nil
    }
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostPortComponentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostPortComponentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostPortComponentable)
                }
            }
            m.SetServices(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHostPortStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*HostPortStatus))
        }
        return nil
    }
    res["timesObserved"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimesObserved(val)
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The first date and time when Microsoft Defender Threat Intelligence observed the hostPort. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HostPort) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetHost gets the host property value. The host property
// returns a Hostable when successful
func (m *HostPort) GetHost()(Hostable) {
    return m.host
}
// GetLastScanDateTime gets the lastScanDateTime property value. The last date and time when Microsoft Defender Threat Intelligence scanned the hostPort. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HostPort) GetLastScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastScanDateTime
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The last date and time when Microsoft Defender Threat Intelligence observed the hostPort. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HostPort) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetMostRecentSslCertificate gets the mostRecentSslCertificate property value. The most recent sslCertificate used to communicate on the port.
// returns a SslCertificateable when successful
func (m *HostPort) GetMostRecentSslCertificate()(SslCertificateable) {
    return m.mostRecentSslCertificate
}
// GetPort gets the port property value. The numerical identifier of the port which is standardized across the internet.
// returns a *int32 when successful
func (m *HostPort) GetPort()(*int32) {
    return m.port
}
// GetProtocol gets the protocol property value. The general protocol used to scan the port. The possible values are: tcp, udp, unknownFutureValue.
// returns a *HostPortProtocol when successful
func (m *HostPort) GetProtocol()(*HostPortProtocol) {
    return m.protocol
}
// GetServices gets the services property value. The hostPortComponents retrieved from scanning the port.
// returns a []HostPortComponentable when successful
func (m *HostPort) GetServices()([]HostPortComponentable) {
    return m.services
}
// GetStatus gets the status property value. The status of the port. The possible values are: open, filtered, closed, unknownFutureValue.
// returns a *HostPortStatus when successful
func (m *HostPort) GetStatus()(*HostPortStatus) {
    return m.status
}
// GetTimesObserved gets the timesObserved property value. The total amount of times that Microsoft Defender Threat Intelligence has observed the hostPort in all its scans.
// returns a *int32 when successful
func (m *HostPort) GetTimesObserved()(*int32) {
    return m.timesObserved
}
// Serialize serializes information the current object
func (m *HostPort) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBanners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBanners()))
        for i, v := range m.GetBanners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("banners", cast)
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
        err = writer.WriteTimeValue("lastScanDateTime", m.GetLastScanDateTime())
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
        err = writer.WriteObjectValue("mostRecentSslCertificate", m.GetMostRecentSslCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    if m.GetProtocol() != nil {
        cast := (*m.GetProtocol()).String()
        err = writer.WriteStringValue("protocol", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("timesObserved", m.GetTimesObserved())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBanners sets the banners property value. The hostPortBanners retrieved from scanning the port.
func (m *HostPort) SetBanners(value []HostPortBannerable)() {
    m.banners = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The first date and time when Microsoft Defender Threat Intelligence observed the hostPort. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *HostPort) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetHost sets the host property value. The host property
func (m *HostPort) SetHost(value Hostable)() {
    m.host = value
}
// SetLastScanDateTime sets the lastScanDateTime property value. The last date and time when Microsoft Defender Threat Intelligence scanned the hostPort. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *HostPort) SetLastScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastScanDateTime = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The last date and time when Microsoft Defender Threat Intelligence observed the hostPort. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *HostPort) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetMostRecentSslCertificate sets the mostRecentSslCertificate property value. The most recent sslCertificate used to communicate on the port.
func (m *HostPort) SetMostRecentSslCertificate(value SslCertificateable)() {
    m.mostRecentSslCertificate = value
}
// SetPort sets the port property value. The numerical identifier of the port which is standardized across the internet.
func (m *HostPort) SetPort(value *int32)() {
    m.port = value
}
// SetProtocol sets the protocol property value. The general protocol used to scan the port. The possible values are: tcp, udp, unknownFutureValue.
func (m *HostPort) SetProtocol(value *HostPortProtocol)() {
    m.protocol = value
}
// SetServices sets the services property value. The hostPortComponents retrieved from scanning the port.
func (m *HostPort) SetServices(value []HostPortComponentable)() {
    m.services = value
}
// SetStatus sets the status property value. The status of the port. The possible values are: open, filtered, closed, unknownFutureValue.
func (m *HostPort) SetStatus(value *HostPortStatus)() {
    m.status = value
}
// SetTimesObserved sets the timesObserved property value. The total amount of times that Microsoft Defender Threat Intelligence has observed the hostPort in all its scans.
func (m *HostPort) SetTimesObserved(value *int32)() {
    m.timesObserved = value
}
type HostPortable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBanners()([]HostPortBannerable)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHost()(Hostable)
    GetLastScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMostRecentSslCertificate()(SslCertificateable)
    GetPort()(*int32)
    GetProtocol()(*HostPortProtocol)
    GetServices()([]HostPortComponentable)
    GetStatus()(*HostPortStatus)
    GetTimesObserved()(*int32)
    SetBanners(value []HostPortBannerable)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHost(value Hostable)()
    SetLastScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMostRecentSslCertificate(value SslCertificateable)()
    SetPort(value *int32)()
    SetProtocol(value *HostPortProtocol)()
    SetServices(value []HostPortComponentable)()
    SetStatus(value *HostPortStatus)()
    SetTimesObserved(value *int32)()
}
