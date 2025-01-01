package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsModelScores the user experience analytics model scores entity consolidates the various Endpoint Analytics scores.
type UserExperienceAnalyticsModelScores struct {
    Entity
    // Indicates a score calculated from application health data to indicate when a device is having problems running one or more applications. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    appReliabilityScore *float64
    // Indicates a calulated score indicating the health of the device's battery. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    batteryHealthScore *float64
    // Indicates a weighted average of the various scores. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    endpointAnalyticsScore *float64
    // The healthStatus property
    healthStatus *UserExperienceAnalyticsHealthState
    // The manufacturer name of the device. Examples: Microsoft Corporation, HP, Lenovo. Supports: $select, $OrderBy. Read-only.
    manufacturer *string
    // The model name of the device. Supports: $select, $OrderBy. Read-only.
    model *string
    // Indicates unique devices count of given model in a consolidated report. Supports: $select, $OrderBy. Read-only. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
    modelDeviceCount *int64
    // Indicates a weighted average of boot score and logon score used for measuring startup performance. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    startupPerformanceScore *float64
    // Indicates a weighted score of the work from anywhere on a device level. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    workFromAnywhereScore *float64
}
// NewUserExperienceAnalyticsModelScores instantiates a new UserExperienceAnalyticsModelScores and sets the default values.
func NewUserExperienceAnalyticsModelScores()(*UserExperienceAnalyticsModelScores) {
    m := &UserExperienceAnalyticsModelScores{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsModelScores(), nil
}
// GetAppReliabilityScore gets the appReliabilityScore property value. Indicates a score calculated from application health data to indicate when a device is having problems running one or more applications. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores) GetAppReliabilityScore()(*float64) {
    return m.appReliabilityScore
}
// GetBatteryHealthScore gets the batteryHealthScore property value. Indicates a calulated score indicating the health of the device's battery. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores) GetBatteryHealthScore()(*float64) {
    return m.batteryHealthScore
}
// GetEndpointAnalyticsScore gets the endpointAnalyticsScore property value. Indicates a weighted average of the various scores. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores) GetEndpointAnalyticsScore()(*float64) {
    return m.endpointAnalyticsScore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsModelScores) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appReliabilityScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppReliabilityScore(val)
        }
        return nil
    }
    res["batteryHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryHealthScore(val)
        }
        return nil
    }
    res["endpointAnalyticsScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointAnalyticsScore(val)
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*UserExperienceAnalyticsHealthState))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["modelDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelDeviceCount(val)
        }
        return nil
    }
    res["startupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupPerformanceScore(val)
        }
        return nil
    }
    res["workFromAnywhereScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkFromAnywhereScore(val)
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *UserExperienceAnalyticsHealthState when successful
func (m *UserExperienceAnalyticsModelScores) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    return m.healthStatus
}
// GetManufacturer gets the manufacturer property value. The manufacturer name of the device. Examples: Microsoft Corporation, HP, Lenovo. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsModelScores) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The model name of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsModelScores) GetModel()(*string) {
    return m.model
}
// GetModelDeviceCount gets the modelDeviceCount property value. Indicates unique devices count of given model in a consolidated report. Supports: $select, $OrderBy. Read-only. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
// returns a *int64 when successful
func (m *UserExperienceAnalyticsModelScores) GetModelDeviceCount()(*int64) {
    return m.modelDeviceCount
}
// GetStartupPerformanceScore gets the startupPerformanceScore property value. Indicates a weighted average of boot score and logon score used for measuring startup performance. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores) GetStartupPerformanceScore()(*float64) {
    return m.startupPerformanceScore
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. Indicates a weighted score of the work from anywhere on a device level. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores) GetWorkFromAnywhereScore()(*float64) {
    return m.workFromAnywhereScore
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsModelScores) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("appReliabilityScore", m.GetAppReliabilityScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("batteryHealthScore", m.GetBatteryHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("endpointAnalyticsScore", m.GetEndpointAnalyticsScore())
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
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("modelDeviceCount", m.GetModelDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("startupPerformanceScore", m.GetStartupPerformanceScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("workFromAnywhereScore", m.GetWorkFromAnywhereScore())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppReliabilityScore sets the appReliabilityScore property value. Indicates a score calculated from application health data to indicate when a device is having problems running one or more applications. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetAppReliabilityScore(value *float64)() {
    m.appReliabilityScore = value
}
// SetBatteryHealthScore sets the batteryHealthScore property value. Indicates a calulated score indicating the health of the device's battery. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetBatteryHealthScore(value *float64)() {
    m.batteryHealthScore = value
}
// SetEndpointAnalyticsScore sets the endpointAnalyticsScore property value. Indicates a weighted average of the various scores. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetEndpointAnalyticsScore(value *float64)() {
    m.endpointAnalyticsScore = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsModelScores) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// SetManufacturer sets the manufacturer property value. The manufacturer name of the device. Examples: Microsoft Corporation, HP, Lenovo. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsModelScores) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The model name of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsModelScores) SetModel(value *string)() {
    m.model = value
}
// SetModelDeviceCount sets the modelDeviceCount property value. Indicates unique devices count of given model in a consolidated report. Supports: $select, $OrderBy. Read-only. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
func (m *UserExperienceAnalyticsModelScores) SetModelDeviceCount(value *int64)() {
    m.modelDeviceCount = value
}
// SetStartupPerformanceScore sets the startupPerformanceScore property value. Indicates a weighted average of boot score and logon score used for measuring startup performance. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetStartupPerformanceScore(value *float64)() {
    m.startupPerformanceScore = value
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. Indicates a weighted score of the work from anywhere on a device level. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetWorkFromAnywhereScore(value *float64)() {
    m.workFromAnywhereScore = value
}
type UserExperienceAnalyticsModelScoresable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppReliabilityScore()(*float64)
    GetBatteryHealthScore()(*float64)
    GetEndpointAnalyticsScore()(*float64)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetModelDeviceCount()(*int64)
    GetStartupPerformanceScore()(*float64)
    GetWorkFromAnywhereScore()(*float64)
    SetAppReliabilityScore(value *float64)()
    SetBatteryHealthScore(value *float64)()
    SetEndpointAnalyticsScore(value *float64)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetModelDeviceCount(value *int64)()
    SetStartupPerformanceScore(value *float64)()
    SetWorkFromAnywhereScore(value *float64)()
}
