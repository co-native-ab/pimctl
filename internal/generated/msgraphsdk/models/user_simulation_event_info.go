package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserSimulationEventInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Browser information from where the simulation event was initiated by a user in an attack simulation and training campaign.
    browser *string
    // The clickSource property
    clickSource *ClickSource
    // Date and time of the simulation event by a user in an attack simulation and training campaign.
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the simulation event by a user in an attack simulation and training campaign.
    eventName *string
    // IP address from where the simulation event was initiated by a user in an attack simulation and training campaign.
    ipAddress *string
    // The OdataType property
    odataType *string
    // The operating system, platform, and device details from where the simulation event was initiated by a user in an attack simulation and training campaign.
    osPlatformDeviceDetails *string
}
// NewUserSimulationEventInfo instantiates a new UserSimulationEventInfo and sets the default values.
func NewUserSimulationEventInfo()(*UserSimulationEventInfo) {
    m := &UserSimulationEventInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserSimulationEventInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserSimulationEventInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSimulationEventInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserSimulationEventInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBrowser gets the browser property value. Browser information from where the simulation event was initiated by a user in an attack simulation and training campaign.
// returns a *string when successful
func (m *UserSimulationEventInfo) GetBrowser()(*string) {
    return m.browser
}
// GetClickSource gets the clickSource property value. The clickSource property
// returns a *ClickSource when successful
func (m *UserSimulationEventInfo) GetClickSource()(*ClickSource) {
    return m.clickSource
}
// GetEventDateTime gets the eventDateTime property value. Date and time of the simulation event by a user in an attack simulation and training campaign.
// returns a *Time when successful
func (m *UserSimulationEventInfo) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventDateTime
}
// GetEventName gets the eventName property value. Name of the simulation event by a user in an attack simulation and training campaign.
// returns a *string when successful
func (m *UserSimulationEventInfo) GetEventName()(*string) {
    return m.eventName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserSimulationEventInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["browser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowser(val)
        }
        return nil
    }
    res["clickSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClickSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClickSource(val.(*ClickSource))
        }
        return nil
    }
    res["eventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["eventName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventName(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["osPlatformDeviceDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsPlatformDeviceDetails(val)
        }
        return nil
    }
    return res
}
// GetIpAddress gets the ipAddress property value. IP address from where the simulation event was initiated by a user in an attack simulation and training campaign.
// returns a *string when successful
func (m *UserSimulationEventInfo) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *UserSimulationEventInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetOsPlatformDeviceDetails gets the osPlatformDeviceDetails property value. The operating system, platform, and device details from where the simulation event was initiated by a user in an attack simulation and training campaign.
// returns a *string when successful
func (m *UserSimulationEventInfo) GetOsPlatformDeviceDetails()(*string) {
    return m.osPlatformDeviceDetails
}
// Serialize serializes information the current object
func (m *UserSimulationEventInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("browser", m.GetBrowser())
        if err != nil {
            return err
        }
    }
    if m.GetClickSource() != nil {
        cast := (*m.GetClickSource()).String()
        err := writer.WriteStringValue("clickSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventName", m.GetEventName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("osPlatformDeviceDetails", m.GetOsPlatformDeviceDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSimulationEventInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBrowser sets the browser property value. Browser information from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetBrowser(value *string)() {
    m.browser = value
}
// SetClickSource sets the clickSource property value. The clickSource property
func (m *UserSimulationEventInfo) SetClickSource(value *ClickSource)() {
    m.clickSource = value
}
// SetEventDateTime sets the eventDateTime property value. Date and time of the simulation event by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// SetEventName sets the eventName property value. Name of the simulation event by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetEventName(value *string)() {
    m.eventName = value
}
// SetIpAddress sets the ipAddress property value. IP address from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserSimulationEventInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOsPlatformDeviceDetails sets the osPlatformDeviceDetails property value. The operating system, platform, and device details from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetOsPlatformDeviceDetails(value *string)() {
    m.osPlatformDeviceDetails = value
}
type UserSimulationEventInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrowser()(*string)
    GetClickSource()(*ClickSource)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEventName()(*string)
    GetIpAddress()(*string)
    GetOdataType()(*string)
    GetOsPlatformDeviceDetails()(*string)
    SetBrowser(value *string)()
    SetClickSource(value *ClickSource)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEventName(value *string)()
    SetIpAddress(value *string)()
    SetOdataType(value *string)()
    SetOsPlatformDeviceDetails(value *string)()
}
