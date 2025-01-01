package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type Sensor struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The date and time when the sensor was generated. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deploymentStatus property
    deploymentStatus *DeploymentStatus
    // The display name of the sensor.
    displayName *string
    // The fully qualified domain name of the sensor.
    domainName *string
    // Represents potential issues within a customer's Microsoft Defender for Identity configuration that Microsoft Defender for Identity identified related to the sensor.
    healthIssues []HealthIssueable
    // The healthStatus property
    healthStatus *SensorHealthStatus
    // This field displays the count of health issues related to this sensor.
    openHealthIssuesCount *int64
    // The sensorType property
    sensorType *SensorType
    // The settings property
    settings SensorSettingsable
    // The version of the sensor.
    version *string
}
// NewSensor instantiates a new Sensor and sets the default values.
func NewSensor()(*Sensor) {
    m := &Sensor{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateSensorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSensorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensor(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the sensor was generated. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *Sensor) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeploymentStatus gets the deploymentStatus property value. The deploymentStatus property
// returns a *DeploymentStatus when successful
func (m *Sensor) GetDeploymentStatus()(*DeploymentStatus) {
    return m.deploymentStatus
}
// GetDisplayName gets the displayName property value. The display name of the sensor.
// returns a *string when successful
func (m *Sensor) GetDisplayName()(*string) {
    return m.displayName
}
// GetDomainName gets the domainName property value. The fully qualified domain name of the sensor.
// returns a *string when successful
func (m *Sensor) GetDomainName()(*string) {
    return m.domainName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Sensor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["deploymentStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentStatus(val.(*DeploymentStatus))
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
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["healthIssues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHealthIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HealthIssueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HealthIssueable)
                }
            }
            m.SetHealthIssues(res)
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensorHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*SensorHealthStatus))
        }
        return nil
    }
    res["openHealthIssuesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenHealthIssuesCount(val)
        }
        return nil
    }
    res["sensorType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensorType(val.(*SensorType))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSensorSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(SensorSettingsable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// GetHealthIssues gets the healthIssues property value. Represents potential issues within a customer's Microsoft Defender for Identity configuration that Microsoft Defender for Identity identified related to the sensor.
// returns a []HealthIssueable when successful
func (m *Sensor) GetHealthIssues()([]HealthIssueable) {
    return m.healthIssues
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *SensorHealthStatus when successful
func (m *Sensor) GetHealthStatus()(*SensorHealthStatus) {
    return m.healthStatus
}
// GetOpenHealthIssuesCount gets the openHealthIssuesCount property value. This field displays the count of health issues related to this sensor.
// returns a *int64 when successful
func (m *Sensor) GetOpenHealthIssuesCount()(*int64) {
    return m.openHealthIssuesCount
}
// GetSensorType gets the sensorType property value. The sensorType property
// returns a *SensorType when successful
func (m *Sensor) GetSensorType()(*SensorType) {
    return m.sensorType
}
// GetSettings gets the settings property value. The settings property
// returns a SensorSettingsable when successful
func (m *Sensor) GetSettings()(SensorSettingsable) {
    return m.settings
}
// GetVersion gets the version property value. The version of the sensor.
// returns a *string when successful
func (m *Sensor) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *Sensor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDeploymentStatus() != nil {
        cast := (*m.GetDeploymentStatus()).String()
        err = writer.WriteStringValue("deploymentStatus", &cast)
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
        err = writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    if m.GetHealthIssues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHealthIssues()))
        for i, v := range m.GetHealthIssues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("healthIssues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("openHealthIssuesCount", m.GetOpenHealthIssuesCount())
        if err != nil {
            return err
        }
    }
    if m.GetSensorType() != nil {
        cast := (*m.GetSensorType()).String()
        err = writer.WriteStringValue("sensorType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the sensor was generated. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Sensor) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeploymentStatus sets the deploymentStatus property value. The deploymentStatus property
func (m *Sensor) SetDeploymentStatus(value *DeploymentStatus)() {
    m.deploymentStatus = value
}
// SetDisplayName sets the displayName property value. The display name of the sensor.
func (m *Sensor) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDomainName sets the domainName property value. The fully qualified domain name of the sensor.
func (m *Sensor) SetDomainName(value *string)() {
    m.domainName = value
}
// SetHealthIssues sets the healthIssues property value. Represents potential issues within a customer's Microsoft Defender for Identity configuration that Microsoft Defender for Identity identified related to the sensor.
func (m *Sensor) SetHealthIssues(value []HealthIssueable)() {
    m.healthIssues = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *Sensor) SetHealthStatus(value *SensorHealthStatus)() {
    m.healthStatus = value
}
// SetOpenHealthIssuesCount sets the openHealthIssuesCount property value. This field displays the count of health issues related to this sensor.
func (m *Sensor) SetOpenHealthIssuesCount(value *int64)() {
    m.openHealthIssuesCount = value
}
// SetSensorType sets the sensorType property value. The sensorType property
func (m *Sensor) SetSensorType(value *SensorType)() {
    m.sensorType = value
}
// SetSettings sets the settings property value. The settings property
func (m *Sensor) SetSettings(value SensorSettingsable)() {
    m.settings = value
}
// SetVersion sets the version property value. The version of the sensor.
func (m *Sensor) SetVersion(value *string)() {
    m.version = value
}
type Sensorable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentStatus()(*DeploymentStatus)
    GetDisplayName()(*string)
    GetDomainName()(*string)
    GetHealthIssues()([]HealthIssueable)
    GetHealthStatus()(*SensorHealthStatus)
    GetOpenHealthIssuesCount()(*int64)
    GetSensorType()(*SensorType)
    GetSettings()(SensorSettingsable)
    GetVersion()(*string)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentStatus(value *DeploymentStatus)()
    SetDisplayName(value *string)()
    SetDomainName(value *string)()
    SetHealthIssues(value []HealthIssueable)()
    SetHealthStatus(value *SensorHealthStatus)()
    SetOpenHealthIssuesCount(value *int64)()
    SetSensorType(value *SensorType)()
    SetSettings(value SensorSettingsable)()
    SetVersion(value *string)()
}
