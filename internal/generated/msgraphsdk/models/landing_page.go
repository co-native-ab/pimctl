package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LandingPage struct {
    Entity
    // Identity of the user who created the landing page.
    createdBy EmailIdentityable
    // Date and time when the landing page was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the landing page as defined by the user.
    description *string
    // The detail information for a landing page associated with a simulation during its creation.
    details []LandingPageDetailable
    // The display name of the landing page.
    displayName *string
    // Email identity of the user who last modified the landing page.
    lastModifiedBy EmailIdentityable
    // Date and time when the landing page was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Content locale.
    locale *string
    // The source of the content. Possible values are: unknown, global, tenant, unknownFutureValue.
    source *SimulationContentSource
    // The status of the simulation. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue.
    status *SimulationContentStatus
    // Supported locales.
    supportedLocales []string
}
// NewLandingPage instantiates a new LandingPage and sets the default values.
func NewLandingPage()(*LandingPage) {
    m := &LandingPage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateLandingPageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLandingPageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLandingPage(), nil
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the landing page.
// returns a EmailIdentityable when successful
func (m *LandingPage) GetCreatedBy()(EmailIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time when the landing page was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *LandingPage) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the landing page as defined by the user.
// returns a *string when successful
func (m *LandingPage) GetDescription()(*string) {
    return m.description
}
// GetDetails gets the details property value. The detail information for a landing page associated with a simulation during its creation.
// returns a []LandingPageDetailable when successful
func (m *LandingPage) GetDetails()([]LandingPageDetailable) {
    return m.details
}
// GetDisplayName gets the displayName property value. The display name of the landing page.
// returns a *string when successful
func (m *LandingPage) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LandingPage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailIdentityable))
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
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLandingPageDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LandingPageDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LandingPageDetailable)
                }
            }
            m.SetDetails(res)
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
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(EmailIdentityable))
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
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationContentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*SimulationContentSource))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationContentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SimulationContentStatus))
        }
        return nil
    }
    res["supportedLocales"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSupportedLocales(res)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. Email identity of the user who last modified the landing page.
// returns a EmailIdentityable when successful
func (m *LandingPage) GetLastModifiedBy()(EmailIdentityable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time when the landing page was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *LandingPage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLocale gets the locale property value. Content locale.
// returns a *string when successful
func (m *LandingPage) GetLocale()(*string) {
    return m.locale
}
// GetSource gets the source property value. The source of the content. Possible values are: unknown, global, tenant, unknownFutureValue.
// returns a *SimulationContentSource when successful
func (m *LandingPage) GetSource()(*SimulationContentSource) {
    return m.source
}
// GetStatus gets the status property value. The status of the simulation. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue.
// returns a *SimulationContentStatus when successful
func (m *LandingPage) GetStatus()(*SimulationContentStatus) {
    return m.status
}
// GetSupportedLocales gets the supportedLocales property value. Supported locales.
// returns a []string when successful
func (m *LandingPage) GetSupportedLocales()([]string) {
    return m.supportedLocales
}
// Serialize serializes information the current object
func (m *LandingPage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("details", cast)
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
        err = writer.WriteStringValue("locale", m.GetLocale())
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
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
    if m.GetSupportedLocales() != nil {
        err = writer.WriteCollectionOfStringValues("supportedLocales", m.GetSupportedLocales())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the landing page.
func (m *LandingPage) SetCreatedBy(value EmailIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time when the landing page was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LandingPage) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the landing page as defined by the user.
func (m *LandingPage) SetDescription(value *string)() {
    m.description = value
}
// SetDetails sets the details property value. The detail information for a landing page associated with a simulation during its creation.
func (m *LandingPage) SetDetails(value []LandingPageDetailable)() {
    m.details = value
}
// SetDisplayName sets the displayName property value. The display name of the landing page.
func (m *LandingPage) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Email identity of the user who last modified the landing page.
func (m *LandingPage) SetLastModifiedBy(value EmailIdentityable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time when the landing page was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LandingPage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLocale sets the locale property value. Content locale.
func (m *LandingPage) SetLocale(value *string)() {
    m.locale = value
}
// SetSource sets the source property value. The source of the content. Possible values are: unknown, global, tenant, unknownFutureValue.
func (m *LandingPage) SetSource(value *SimulationContentSource)() {
    m.source = value
}
// SetStatus sets the status property value. The status of the simulation. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue.
func (m *LandingPage) SetStatus(value *SimulationContentStatus)() {
    m.status = value
}
// SetSupportedLocales sets the supportedLocales property value. Supported locales.
func (m *LandingPage) SetSupportedLocales(value []string)() {
    m.supportedLocales = value
}
type LandingPageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(EmailIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDetails()([]LandingPageDetailable)
    GetDisplayName()(*string)
    GetLastModifiedBy()(EmailIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocale()(*string)
    GetSource()(*SimulationContentSource)
    GetStatus()(*SimulationContentStatus)
    GetSupportedLocales()([]string)
    SetCreatedBy(value EmailIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDetails(value []LandingPageDetailable)()
    SetDisplayName(value *string)()
    SetLastModifiedBy(value EmailIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocale(value *string)()
    SetSource(value *SimulationContentSource)()
    SetStatus(value *SimulationContentStatus)()
    SetSupportedLocales(value []string)()
}
