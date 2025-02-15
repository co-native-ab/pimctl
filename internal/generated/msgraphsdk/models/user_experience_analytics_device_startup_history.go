package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceStartupHistory the user experience analytics device startup history entity contains device boot performance history details.
type UserExperienceAnalyticsDeviceStartupHistory struct {
    Entity
    // The device core boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
    coreBootTimeInMs *int32
    // The device core login time in milliseconds. Supports: $select, $OrderBy. Read-only.
    coreLoginTimeInMs *int32
    // The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
    deviceId *string
    // The impact of device feature updates on boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
    featureUpdateBootTimeInMs *int32
    // The impact of device group policy client on boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
    groupPolicyBootTimeInMs *int32
    // The impact of device group policy client on login time in milliseconds. Supports: $select, $OrderBy. Read-only.
    groupPolicyLoginTimeInMs *int32
    // When TRUE, indicates the device boot record is associated with feature updates. When FALSE, indicates the device boot record is not associated with feature updates. Supports: $select, $OrderBy. Read-only.
    isFeatureUpdate *bool
    // When TRUE, indicates the device login is the first login after a reboot. When FALSE, indicates the device login is not the first login after a reboot. Supports: $select, $OrderBy. Read-only.
    isFirstLogin *bool
    // The user experience analytics device boot record's operating system version. Supports: $select, $OrderBy. Read-only.
    operatingSystemVersion *string
    // The time for desktop to become responsive during login process in milliseconds. Supports: $select, $OrderBy. Read-only.
    responsiveDesktopTimeInMs *int32
    // Operating System restart category.
    restartCategory *UserExperienceAnalyticsOperatingSystemRestartCategory
    // OS restart fault bucket. The fault bucket is used to find additional information about a system crash. Supports: $select, $OrderBy. Read-only.
    restartFaultBucket *string
    // OS restart stop code. This shows the bug check code which can be used to look up the blue screen reason. Supports: $select, $OrderBy. Read-only.
    restartStopCode *string
    // The device boot start time. The value cannot be modified and is automatically populated when the device performs a reboot. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2022 would look like this: '2022-01-01T00:00:00Z'. Returned by default. Read-only.
    startTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The device total boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
    totalBootTimeInMs *int32
    // The device total login time in milliseconds. Supports: $select, $OrderBy. Read-only.
    totalLoginTimeInMs *int32
}
// NewUserExperienceAnalyticsDeviceStartupHistory instantiates a new UserExperienceAnalyticsDeviceStartupHistory and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupHistory()(*UserExperienceAnalyticsDeviceStartupHistory) {
    m := &UserExperienceAnalyticsDeviceStartupHistory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceStartupHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDeviceStartupHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceStartupHistory(), nil
}
// GetCoreBootTimeInMs gets the coreBootTimeInMs property value. The device core boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetCoreBootTimeInMs()(*int32) {
    return m.coreBootTimeInMs
}
// GetCoreLoginTimeInMs gets the coreLoginTimeInMs property value. The device core login time in milliseconds. Supports: $select, $OrderBy. Read-only.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetCoreLoginTimeInMs()(*int32) {
    return m.coreLoginTimeInMs
}
// GetDeviceId gets the deviceId property value. The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFeatureUpdateBootTimeInMs gets the featureUpdateBootTimeInMs property value. The impact of device feature updates on boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetFeatureUpdateBootTimeInMs()(*int32) {
    return m.featureUpdateBootTimeInMs
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["coreBootTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreBootTimeInMs(val)
        }
        return nil
    }
    res["coreLoginTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreLoginTimeInMs(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["featureUpdateBootTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdateBootTimeInMs(val)
        }
        return nil
    }
    res["groupPolicyBootTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyBootTimeInMs(val)
        }
        return nil
    }
    res["groupPolicyLoginTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyLoginTimeInMs(val)
        }
        return nil
    }
    res["isFeatureUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFeatureUpdate(val)
        }
        return nil
    }
    res["isFirstLogin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFirstLogin(val)
        }
        return nil
    }
    res["operatingSystemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemVersion(val)
        }
        return nil
    }
    res["responsiveDesktopTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsiveDesktopTimeInMs(val)
        }
        return nil
    }
    res["restartCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsOperatingSystemRestartCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartCategory(val.(*UserExperienceAnalyticsOperatingSystemRestartCategory))
        }
        return nil
    }
    res["restartFaultBucket"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartFaultBucket(val)
        }
        return nil
    }
    res["restartStopCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartStopCode(val)
        }
        return nil
    }
    res["startTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    res["totalBootTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBootTimeInMs(val)
        }
        return nil
    }
    res["totalLoginTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLoginTimeInMs(val)
        }
        return nil
    }
    return res
}
// GetGroupPolicyBootTimeInMs gets the groupPolicyBootTimeInMs property value. The impact of device group policy client on boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetGroupPolicyBootTimeInMs()(*int32) {
    return m.groupPolicyBootTimeInMs
}
// GetGroupPolicyLoginTimeInMs gets the groupPolicyLoginTimeInMs property value. The impact of device group policy client on login time in milliseconds. Supports: $select, $OrderBy. Read-only.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetGroupPolicyLoginTimeInMs()(*int32) {
    return m.groupPolicyLoginTimeInMs
}
// GetIsFeatureUpdate gets the isFeatureUpdate property value. When TRUE, indicates the device boot record is associated with feature updates. When FALSE, indicates the device boot record is not associated with feature updates. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetIsFeatureUpdate()(*bool) {
    return m.isFeatureUpdate
}
// GetIsFirstLogin gets the isFirstLogin property value. When TRUE, indicates the device login is the first login after a reboot. When FALSE, indicates the device login is not the first login after a reboot. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetIsFirstLogin()(*bool) {
    return m.isFirstLogin
}
// GetOperatingSystemVersion gets the operatingSystemVersion property value. The user experience analytics device boot record's operating system version. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetOperatingSystemVersion()(*string) {
    return m.operatingSystemVersion
}
// GetResponsiveDesktopTimeInMs gets the responsiveDesktopTimeInMs property value. The time for desktop to become responsive during login process in milliseconds. Supports: $select, $OrderBy. Read-only.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetResponsiveDesktopTimeInMs()(*int32) {
    return m.responsiveDesktopTimeInMs
}
// GetRestartCategory gets the restartCategory property value. Operating System restart category.
// returns a *UserExperienceAnalyticsOperatingSystemRestartCategory when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetRestartCategory()(*UserExperienceAnalyticsOperatingSystemRestartCategory) {
    return m.restartCategory
}
// GetRestartFaultBucket gets the restartFaultBucket property value. OS restart fault bucket. The fault bucket is used to find additional information about a system crash. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetRestartFaultBucket()(*string) {
    return m.restartFaultBucket
}
// GetRestartStopCode gets the restartStopCode property value. OS restart stop code. This shows the bug check code which can be used to look up the blue screen reason. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetRestartStopCode()(*string) {
    return m.restartStopCode
}
// GetStartTime gets the startTime property value. The device boot start time. The value cannot be modified and is automatically populated when the device performs a reboot. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2022 would look like this: '2022-01-01T00:00:00Z'. Returned by default. Read-only.
// returns a *Time when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startTime
}
// GetTotalBootTimeInMs gets the totalBootTimeInMs property value. The device total boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetTotalBootTimeInMs()(*int32) {
    return m.totalBootTimeInMs
}
// GetTotalLoginTimeInMs gets the totalLoginTimeInMs property value. The device total login time in milliseconds. Supports: $select, $OrderBy. Read-only.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetTotalLoginTimeInMs()(*int32) {
    return m.totalLoginTimeInMs
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceStartupHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("coreBootTimeInMs", m.GetCoreBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("coreLoginTimeInMs", m.GetCoreLoginTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("featureUpdateBootTimeInMs", m.GetFeatureUpdateBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("groupPolicyBootTimeInMs", m.GetGroupPolicyBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("groupPolicyLoginTimeInMs", m.GetGroupPolicyLoginTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFeatureUpdate", m.GetIsFeatureUpdate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFirstLogin", m.GetIsFirstLogin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystemVersion", m.GetOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("responsiveDesktopTimeInMs", m.GetResponsiveDesktopTimeInMs())
        if err != nil {
            return err
        }
    }
    if m.GetRestartCategory() != nil {
        cast := (*m.GetRestartCategory()).String()
        err = writer.WriteStringValue("restartCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("restartFaultBucket", m.GetRestartFaultBucket())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("restartStopCode", m.GetRestartStopCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startTime", m.GetStartTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalBootTimeInMs", m.GetTotalBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLoginTimeInMs", m.GetTotalLoginTimeInMs())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCoreBootTimeInMs sets the coreBootTimeInMs property value. The device core boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetCoreBootTimeInMs(value *int32)() {
    m.coreBootTimeInMs = value
}
// SetCoreLoginTimeInMs sets the coreLoginTimeInMs property value. The device core login time in milliseconds. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetCoreLoginTimeInMs(value *int32)() {
    m.coreLoginTimeInMs = value
}
// SetDeviceId sets the deviceId property value. The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetFeatureUpdateBootTimeInMs sets the featureUpdateBootTimeInMs property value. The impact of device feature updates on boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetFeatureUpdateBootTimeInMs(value *int32)() {
    m.featureUpdateBootTimeInMs = value
}
// SetGroupPolicyBootTimeInMs sets the groupPolicyBootTimeInMs property value. The impact of device group policy client on boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetGroupPolicyBootTimeInMs(value *int32)() {
    m.groupPolicyBootTimeInMs = value
}
// SetGroupPolicyLoginTimeInMs sets the groupPolicyLoginTimeInMs property value. The impact of device group policy client on login time in milliseconds. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetGroupPolicyLoginTimeInMs(value *int32)() {
    m.groupPolicyLoginTimeInMs = value
}
// SetIsFeatureUpdate sets the isFeatureUpdate property value. When TRUE, indicates the device boot record is associated with feature updates. When FALSE, indicates the device boot record is not associated with feature updates. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetIsFeatureUpdate(value *bool)() {
    m.isFeatureUpdate = value
}
// SetIsFirstLogin sets the isFirstLogin property value. When TRUE, indicates the device login is the first login after a reboot. When FALSE, indicates the device login is not the first login after a reboot. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetIsFirstLogin(value *bool)() {
    m.isFirstLogin = value
}
// SetOperatingSystemVersion sets the operatingSystemVersion property value. The user experience analytics device boot record's operating system version. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetOperatingSystemVersion(value *string)() {
    m.operatingSystemVersion = value
}
// SetResponsiveDesktopTimeInMs sets the responsiveDesktopTimeInMs property value. The time for desktop to become responsive during login process in milliseconds. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetResponsiveDesktopTimeInMs(value *int32)() {
    m.responsiveDesktopTimeInMs = value
}
// SetRestartCategory sets the restartCategory property value. Operating System restart category.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetRestartCategory(value *UserExperienceAnalyticsOperatingSystemRestartCategory)() {
    m.restartCategory = value
}
// SetRestartFaultBucket sets the restartFaultBucket property value. OS restart fault bucket. The fault bucket is used to find additional information about a system crash. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetRestartFaultBucket(value *string)() {
    m.restartFaultBucket = value
}
// SetRestartStopCode sets the restartStopCode property value. OS restart stop code. This shows the bug check code which can be used to look up the blue screen reason. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetRestartStopCode(value *string)() {
    m.restartStopCode = value
}
// SetStartTime sets the startTime property value. The device boot start time. The value cannot be modified and is automatically populated when the device performs a reboot. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2022 would look like this: '2022-01-01T00:00:00Z'. Returned by default. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startTime = value
}
// SetTotalBootTimeInMs sets the totalBootTimeInMs property value. The device total boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetTotalBootTimeInMs(value *int32)() {
    m.totalBootTimeInMs = value
}
// SetTotalLoginTimeInMs sets the totalLoginTimeInMs property value. The device total login time in milliseconds. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetTotalLoginTimeInMs(value *int32)() {
    m.totalLoginTimeInMs = value
}
type UserExperienceAnalyticsDeviceStartupHistoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCoreBootTimeInMs()(*int32)
    GetCoreLoginTimeInMs()(*int32)
    GetDeviceId()(*string)
    GetFeatureUpdateBootTimeInMs()(*int32)
    GetGroupPolicyBootTimeInMs()(*int32)
    GetGroupPolicyLoginTimeInMs()(*int32)
    GetIsFeatureUpdate()(*bool)
    GetIsFirstLogin()(*bool)
    GetOperatingSystemVersion()(*string)
    GetResponsiveDesktopTimeInMs()(*int32)
    GetRestartCategory()(*UserExperienceAnalyticsOperatingSystemRestartCategory)
    GetRestartFaultBucket()(*string)
    GetRestartStopCode()(*string)
    GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTotalBootTimeInMs()(*int32)
    GetTotalLoginTimeInMs()(*int32)
    SetCoreBootTimeInMs(value *int32)()
    SetCoreLoginTimeInMs(value *int32)()
    SetDeviceId(value *string)()
    SetFeatureUpdateBootTimeInMs(value *int32)()
    SetGroupPolicyBootTimeInMs(value *int32)()
    SetGroupPolicyLoginTimeInMs(value *int32)()
    SetIsFeatureUpdate(value *bool)()
    SetIsFirstLogin(value *bool)()
    SetOperatingSystemVersion(value *string)()
    SetResponsiveDesktopTimeInMs(value *int32)()
    SetRestartCategory(value *UserExperienceAnalyticsOperatingSystemRestartCategory)()
    SetRestartFaultBucket(value *string)()
    SetRestartStopCode(value *string)()
    SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTotalBootTimeInMs(value *int32)()
    SetTotalLoginTimeInMs(value *int32)()
}
