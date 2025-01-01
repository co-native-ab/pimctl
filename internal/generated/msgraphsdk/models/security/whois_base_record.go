package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type WhoisBaseRecord struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The contact information for the abuse contact.
    abuse WhoisContactable
    // The contact information for the admin contact.
    admin WhoisContactable
    // The contact information for the billing contact.
    billing WhoisContactable
    // The domain status for this WHOIS object.
    domainStatus *string
    // The date and time when this WHOIS record expires with the registrar. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The first seen date and time of this WHOIS record. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The host property
    host Hostable
    // The last seen date and time of this WHOIS record. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time when this WHOIS record was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The nameservers for this WHOIS object.
    nameservers []WhoisNameserverable
    // The contact information for the noc contact.
    noc WhoisContactable
    // The raw WHOIS details for this WHOIS object.
    rawWhoisText *string
    // The contact information for the registrant contact.
    registrant WhoisContactable
    // The contact information for the registrar contact.
    registrar WhoisContactable
    // The date and time when this WHOIS record was registered with a registrar. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    registrationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The contact information for the technical contact.
    technical WhoisContactable
    // The WHOIS server that provides the details.
    whoisServer *string
    // The contact information for the zone contact.
    zone WhoisContactable
}
// NewWhoisBaseRecord instantiates a new WhoisBaseRecord and sets the default values.
func NewWhoisBaseRecord()(*WhoisBaseRecord) {
    m := &WhoisBaseRecord{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateWhoisBaseRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWhoisBaseRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.whoisHistoryRecord":
                        return NewWhoisHistoryRecord(), nil
                    case "#microsoft.graph.security.whoisRecord":
                        return NewWhoisRecord(), nil
                }
            }
        }
    }
    return NewWhoisBaseRecord(), nil
}
// GetAbuse gets the abuse property value. The contact information for the abuse contact.
// returns a WhoisContactable when successful
func (m *WhoisBaseRecord) GetAbuse()(WhoisContactable) {
    return m.abuse
}
// GetAdmin gets the admin property value. The contact information for the admin contact.
// returns a WhoisContactable when successful
func (m *WhoisBaseRecord) GetAdmin()(WhoisContactable) {
    return m.admin
}
// GetBilling gets the billing property value. The contact information for the billing contact.
// returns a WhoisContactable when successful
func (m *WhoisBaseRecord) GetBilling()(WhoisContactable) {
    return m.billing
}
// GetDomainStatus gets the domainStatus property value. The domain status for this WHOIS object.
// returns a *string when successful
func (m *WhoisBaseRecord) GetDomainStatus()(*string) {
    return m.domainStatus
}
// GetExpirationDateTime gets the expirationDateTime property value. The date and time when this WHOIS record expires with the registrar. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WhoisBaseRecord) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WhoisBaseRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["abuse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWhoisContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAbuse(val.(WhoisContactable))
        }
        return nil
    }
    res["admin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWhoisContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdmin(val.(WhoisContactable))
        }
        return nil
    }
    res["billing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWhoisContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBilling(val.(WhoisContactable))
        }
        return nil
    }
    res["domainStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainStatus(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
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
    res["lastUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
        }
        return nil
    }
    res["nameservers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWhoisNameserverFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WhoisNameserverable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WhoisNameserverable)
                }
            }
            m.SetNameservers(res)
        }
        return nil
    }
    res["noc"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWhoisContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoc(val.(WhoisContactable))
        }
        return nil
    }
    res["rawWhoisText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRawWhoisText(val)
        }
        return nil
    }
    res["registrant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWhoisContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrant(val.(WhoisContactable))
        }
        return nil
    }
    res["registrar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWhoisContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrar(val.(WhoisContactable))
        }
        return nil
    }
    res["registrationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDateTime(val)
        }
        return nil
    }
    res["technical"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWhoisContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnical(val.(WhoisContactable))
        }
        return nil
    }
    res["whoisServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWhoisServer(val)
        }
        return nil
    }
    res["zone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWhoisContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZone(val.(WhoisContactable))
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The first seen date and time of this WHOIS record. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WhoisBaseRecord) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetHost gets the host property value. The host property
// returns a Hostable when successful
func (m *WhoisBaseRecord) GetHost()(Hostable) {
    return m.host
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The last seen date and time of this WHOIS record. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WhoisBaseRecord) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. The date and time when this WHOIS record was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WhoisBaseRecord) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdateDateTime
}
// GetNameservers gets the nameservers property value. The nameservers for this WHOIS object.
// returns a []WhoisNameserverable when successful
func (m *WhoisBaseRecord) GetNameservers()([]WhoisNameserverable) {
    return m.nameservers
}
// GetNoc gets the noc property value. The contact information for the noc contact.
// returns a WhoisContactable when successful
func (m *WhoisBaseRecord) GetNoc()(WhoisContactable) {
    return m.noc
}
// GetRawWhoisText gets the rawWhoisText property value. The raw WHOIS details for this WHOIS object.
// returns a *string when successful
func (m *WhoisBaseRecord) GetRawWhoisText()(*string) {
    return m.rawWhoisText
}
// GetRegistrant gets the registrant property value. The contact information for the registrant contact.
// returns a WhoisContactable when successful
func (m *WhoisBaseRecord) GetRegistrant()(WhoisContactable) {
    return m.registrant
}
// GetRegistrar gets the registrar property value. The contact information for the registrar contact.
// returns a WhoisContactable when successful
func (m *WhoisBaseRecord) GetRegistrar()(WhoisContactable) {
    return m.registrar
}
// GetRegistrationDateTime gets the registrationDateTime property value. The date and time when this WHOIS record was registered with a registrar. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WhoisBaseRecord) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.registrationDateTime
}
// GetTechnical gets the technical property value. The contact information for the technical contact.
// returns a WhoisContactable when successful
func (m *WhoisBaseRecord) GetTechnical()(WhoisContactable) {
    return m.technical
}
// GetWhoisServer gets the whoisServer property value. The WHOIS server that provides the details.
// returns a *string when successful
func (m *WhoisBaseRecord) GetWhoisServer()(*string) {
    return m.whoisServer
}
// GetZone gets the zone property value. The contact information for the zone contact.
// returns a WhoisContactable when successful
func (m *WhoisBaseRecord) GetZone()(WhoisContactable) {
    return m.zone
}
// Serialize serializes information the current object
func (m *WhoisBaseRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("abuse", m.GetAbuse())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("admin", m.GetAdmin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("billing", m.GetBilling())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domainStatus", m.GetDomainStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
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
        err = writer.WriteTimeValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetNameservers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNameservers()))
        for i, v := range m.GetNameservers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("nameservers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("noc", m.GetNoc())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rawWhoisText", m.GetRawWhoisText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("registrant", m.GetRegistrant())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("registrar", m.GetRegistrar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registrationDateTime", m.GetRegistrationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("technical", m.GetTechnical())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("whoisServer", m.GetWhoisServer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("zone", m.GetZone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAbuse sets the abuse property value. The contact information for the abuse contact.
func (m *WhoisBaseRecord) SetAbuse(value WhoisContactable)() {
    m.abuse = value
}
// SetAdmin sets the admin property value. The contact information for the admin contact.
func (m *WhoisBaseRecord) SetAdmin(value WhoisContactable)() {
    m.admin = value
}
// SetBilling sets the billing property value. The contact information for the billing contact.
func (m *WhoisBaseRecord) SetBilling(value WhoisContactable)() {
    m.billing = value
}
// SetDomainStatus sets the domainStatus property value. The domain status for this WHOIS object.
func (m *WhoisBaseRecord) SetDomainStatus(value *string)() {
    m.domainStatus = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The date and time when this WHOIS record expires with the registrar. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WhoisBaseRecord) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The first seen date and time of this WHOIS record. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WhoisBaseRecord) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetHost sets the host property value. The host property
func (m *WhoisBaseRecord) SetHost(value Hostable)() {
    m.host = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The last seen date and time of this WHOIS record. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WhoisBaseRecord) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. The date and time when this WHOIS record was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WhoisBaseRecord) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateDateTime = value
}
// SetNameservers sets the nameservers property value. The nameservers for this WHOIS object.
func (m *WhoisBaseRecord) SetNameservers(value []WhoisNameserverable)() {
    m.nameservers = value
}
// SetNoc sets the noc property value. The contact information for the noc contact.
func (m *WhoisBaseRecord) SetNoc(value WhoisContactable)() {
    m.noc = value
}
// SetRawWhoisText sets the rawWhoisText property value. The raw WHOIS details for this WHOIS object.
func (m *WhoisBaseRecord) SetRawWhoisText(value *string)() {
    m.rawWhoisText = value
}
// SetRegistrant sets the registrant property value. The contact information for the registrant contact.
func (m *WhoisBaseRecord) SetRegistrant(value WhoisContactable)() {
    m.registrant = value
}
// SetRegistrar sets the registrar property value. The contact information for the registrar contact.
func (m *WhoisBaseRecord) SetRegistrar(value WhoisContactable)() {
    m.registrar = value
}
// SetRegistrationDateTime sets the registrationDateTime property value. The date and time when this WHOIS record was registered with a registrar. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WhoisBaseRecord) SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registrationDateTime = value
}
// SetTechnical sets the technical property value. The contact information for the technical contact.
func (m *WhoisBaseRecord) SetTechnical(value WhoisContactable)() {
    m.technical = value
}
// SetWhoisServer sets the whoisServer property value. The WHOIS server that provides the details.
func (m *WhoisBaseRecord) SetWhoisServer(value *string)() {
    m.whoisServer = value
}
// SetZone sets the zone property value. The contact information for the zone contact.
func (m *WhoisBaseRecord) SetZone(value WhoisContactable)() {
    m.zone = value
}
type WhoisBaseRecordable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAbuse()(WhoisContactable)
    GetAdmin()(WhoisContactable)
    GetBilling()(WhoisContactable)
    GetDomainStatus()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHost()(Hostable)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNameservers()([]WhoisNameserverable)
    GetNoc()(WhoisContactable)
    GetRawWhoisText()(*string)
    GetRegistrant()(WhoisContactable)
    GetRegistrar()(WhoisContactable)
    GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTechnical()(WhoisContactable)
    GetWhoisServer()(*string)
    GetZone()(WhoisContactable)
    SetAbuse(value WhoisContactable)()
    SetAdmin(value WhoisContactable)()
    SetBilling(value WhoisContactable)()
    SetDomainStatus(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHost(value Hostable)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNameservers(value []WhoisNameserverable)()
    SetNoc(value WhoisContactable)()
    SetRawWhoisText(value *string)()
    SetRegistrant(value WhoisContactable)()
    SetRegistrar(value WhoisContactable)()
    SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTechnical(value WhoisContactable)()
    SetWhoisServer(value *string)()
    SetZone(value WhoisContactable)()
}
