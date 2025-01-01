package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBaseline the user experience analytics baseline entity contains baseline values against which to compare the user experience analytics scores.
type UserExperienceAnalyticsBaseline struct {
    Entity
    // The scores and insights for the application health metrics.
    appHealthMetrics UserExperienceAnalyticsCategoryable
    // The scores and insights for the battery health metrics.
    batteryHealthMetrics UserExperienceAnalyticsCategoryable
    // The scores and insights for the best practices metrics.
    bestPracticesMetrics UserExperienceAnalyticsCategoryable
    // The date the custom baseline was created. The value cannot be modified and is automatically populated when the baseline is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The scores and insights for the device boot performance metrics.
    deviceBootPerformanceMetrics UserExperienceAnalyticsCategoryable
    // The name of the baseline.
    displayName *string
    // When TRUE, indicates the current baseline is the commercial median baseline. When FALSE, indicates it is a custom baseline. FALSE by default.
    isBuiltIn *bool
    // The scores and insights for the reboot analytics metrics.
    rebootAnalyticsMetrics UserExperienceAnalyticsCategoryable
    // The scores and insights for the resource performance metrics.
    resourcePerformanceMetrics UserExperienceAnalyticsCategoryable
    // The scores and insights for the work from anywhere metrics.
    workFromAnywhereMetrics UserExperienceAnalyticsCategoryable
}
// NewUserExperienceAnalyticsBaseline instantiates a new UserExperienceAnalyticsBaseline and sets the default values.
func NewUserExperienceAnalyticsBaseline()(*UserExperienceAnalyticsBaseline) {
    m := &UserExperienceAnalyticsBaseline{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBaseline(), nil
}
// GetAppHealthMetrics gets the appHealthMetrics property value. The scores and insights for the application health metrics.
// returns a UserExperienceAnalyticsCategoryable when successful
func (m *UserExperienceAnalyticsBaseline) GetAppHealthMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.appHealthMetrics
}
// GetBatteryHealthMetrics gets the batteryHealthMetrics property value. The scores and insights for the battery health metrics.
// returns a UserExperienceAnalyticsCategoryable when successful
func (m *UserExperienceAnalyticsBaseline) GetBatteryHealthMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.batteryHealthMetrics
}
// GetBestPracticesMetrics gets the bestPracticesMetrics property value. The scores and insights for the best practices metrics.
// returns a UserExperienceAnalyticsCategoryable when successful
func (m *UserExperienceAnalyticsBaseline) GetBestPracticesMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.bestPracticesMetrics
}
// GetCreatedDateTime gets the createdDateTime property value. The date the custom baseline was created. The value cannot be modified and is automatically populated when the baseline is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default.
// returns a *Time when successful
func (m *UserExperienceAnalyticsBaseline) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeviceBootPerformanceMetrics gets the deviceBootPerformanceMetrics property value. The scores and insights for the device boot performance metrics.
// returns a UserExperienceAnalyticsCategoryable when successful
func (m *UserExperienceAnalyticsBaseline) GetDeviceBootPerformanceMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.deviceBootPerformanceMetrics
}
// GetDisplayName gets the displayName property value. The name of the baseline.
// returns a *string when successful
func (m *UserExperienceAnalyticsBaseline) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsBaseline) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appHealthMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppHealthMetrics(val.(UserExperienceAnalyticsCategoryable))
        }
        return nil
    }
    res["batteryHealthMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryHealthMetrics(val.(UserExperienceAnalyticsCategoryable))
        }
        return nil
    }
    res["bestPracticesMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBestPracticesMetrics(val.(UserExperienceAnalyticsCategoryable))
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
    res["deviceBootPerformanceMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBootPerformanceMetrics(val.(UserExperienceAnalyticsCategoryable))
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
    res["isBuiltIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    res["rebootAnalyticsMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRebootAnalyticsMetrics(val.(UserExperienceAnalyticsCategoryable))
        }
        return nil
    }
    res["resourcePerformanceMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourcePerformanceMetrics(val.(UserExperienceAnalyticsCategoryable))
        }
        return nil
    }
    res["workFromAnywhereMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkFromAnywhereMetrics(val.(UserExperienceAnalyticsCategoryable))
        }
        return nil
    }
    return res
}
// GetIsBuiltIn gets the isBuiltIn property value. When TRUE, indicates the current baseline is the commercial median baseline. When FALSE, indicates it is a custom baseline. FALSE by default.
// returns a *bool when successful
func (m *UserExperienceAnalyticsBaseline) GetIsBuiltIn()(*bool) {
    return m.isBuiltIn
}
// GetRebootAnalyticsMetrics gets the rebootAnalyticsMetrics property value. The scores and insights for the reboot analytics metrics.
// returns a UserExperienceAnalyticsCategoryable when successful
func (m *UserExperienceAnalyticsBaseline) GetRebootAnalyticsMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.rebootAnalyticsMetrics
}
// GetResourcePerformanceMetrics gets the resourcePerformanceMetrics property value. The scores and insights for the resource performance metrics.
// returns a UserExperienceAnalyticsCategoryable when successful
func (m *UserExperienceAnalyticsBaseline) GetResourcePerformanceMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.resourcePerformanceMetrics
}
// GetWorkFromAnywhereMetrics gets the workFromAnywhereMetrics property value. The scores and insights for the work from anywhere metrics.
// returns a UserExperienceAnalyticsCategoryable when successful
func (m *UserExperienceAnalyticsBaseline) GetWorkFromAnywhereMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.workFromAnywhereMetrics
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBaseline) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appHealthMetrics", m.GetAppHealthMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("batteryHealthMetrics", m.GetBatteryHealthMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bestPracticesMetrics", m.GetBestPracticesMetrics())
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
        err = writer.WriteObjectValue("deviceBootPerformanceMetrics", m.GetDeviceBootPerformanceMetrics())
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
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rebootAnalyticsMetrics", m.GetRebootAnalyticsMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resourcePerformanceMetrics", m.GetResourcePerformanceMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("workFromAnywhereMetrics", m.GetWorkFromAnywhereMetrics())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppHealthMetrics sets the appHealthMetrics property value. The scores and insights for the application health metrics.
func (m *UserExperienceAnalyticsBaseline) SetAppHealthMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.appHealthMetrics = value
}
// SetBatteryHealthMetrics sets the batteryHealthMetrics property value. The scores and insights for the battery health metrics.
func (m *UserExperienceAnalyticsBaseline) SetBatteryHealthMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.batteryHealthMetrics = value
}
// SetBestPracticesMetrics sets the bestPracticesMetrics property value. The scores and insights for the best practices metrics.
func (m *UserExperienceAnalyticsBaseline) SetBestPracticesMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.bestPracticesMetrics = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date the custom baseline was created. The value cannot be modified and is automatically populated when the baseline is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default.
func (m *UserExperienceAnalyticsBaseline) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeviceBootPerformanceMetrics sets the deviceBootPerformanceMetrics property value. The scores and insights for the device boot performance metrics.
func (m *UserExperienceAnalyticsBaseline) SetDeviceBootPerformanceMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.deviceBootPerformanceMetrics = value
}
// SetDisplayName sets the displayName property value. The name of the baseline.
func (m *UserExperienceAnalyticsBaseline) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsBuiltIn sets the isBuiltIn property value. When TRUE, indicates the current baseline is the commercial median baseline. When FALSE, indicates it is a custom baseline. FALSE by default.
func (m *UserExperienceAnalyticsBaseline) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
// SetRebootAnalyticsMetrics sets the rebootAnalyticsMetrics property value. The scores and insights for the reboot analytics metrics.
func (m *UserExperienceAnalyticsBaseline) SetRebootAnalyticsMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.rebootAnalyticsMetrics = value
}
// SetResourcePerformanceMetrics sets the resourcePerformanceMetrics property value. The scores and insights for the resource performance metrics.
func (m *UserExperienceAnalyticsBaseline) SetResourcePerformanceMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.resourcePerformanceMetrics = value
}
// SetWorkFromAnywhereMetrics sets the workFromAnywhereMetrics property value. The scores and insights for the work from anywhere metrics.
func (m *UserExperienceAnalyticsBaseline) SetWorkFromAnywhereMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.workFromAnywhereMetrics = value
}
type UserExperienceAnalyticsBaselineable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppHealthMetrics()(UserExperienceAnalyticsCategoryable)
    GetBatteryHealthMetrics()(UserExperienceAnalyticsCategoryable)
    GetBestPracticesMetrics()(UserExperienceAnalyticsCategoryable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceBootPerformanceMetrics()(UserExperienceAnalyticsCategoryable)
    GetDisplayName()(*string)
    GetIsBuiltIn()(*bool)
    GetRebootAnalyticsMetrics()(UserExperienceAnalyticsCategoryable)
    GetResourcePerformanceMetrics()(UserExperienceAnalyticsCategoryable)
    GetWorkFromAnywhereMetrics()(UserExperienceAnalyticsCategoryable)
    SetAppHealthMetrics(value UserExperienceAnalyticsCategoryable)()
    SetBatteryHealthMetrics(value UserExperienceAnalyticsCategoryable)()
    SetBestPracticesMetrics(value UserExperienceAnalyticsCategoryable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceBootPerformanceMetrics(value UserExperienceAnalyticsCategoryable)()
    SetDisplayName(value *string)()
    SetIsBuiltIn(value *bool)()
    SetRebootAnalyticsMetrics(value UserExperienceAnalyticsCategoryable)()
    SetResourcePerformanceMetrics(value UserExperienceAnalyticsCategoryable)()
    SetWorkFromAnywhereMetrics(value UserExperienceAnalyticsCategoryable)()
}
