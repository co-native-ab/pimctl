package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppAssignmentSettings contains properties used to assign an Win32 LOB mobile app to a group.
type Win32LobAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // The auto-update settings to apply for this app assignment.
    autoUpdateSettings Win32LobAppAutoUpdateSettingsable
    // Contains value for delivery optimization priority.
    deliveryOptimizationPriority *Win32LobAppDeliveryOptimizationPriority
    // The install time settings to apply for this app assignment.
    installTimeSettings MobileAppInstallTimeSettingsable
    // Contains value for notification status.
    notifications *Win32LobAppNotification
    // The reboot settings to apply for this app assignment.
    restartSettings Win32LobAppRestartSettingsable
}
// NewWin32LobAppAssignmentSettings instantiates a new Win32LobAppAssignmentSettings and sets the default values.
func NewWin32LobAppAssignmentSettings()(*Win32LobAppAssignmentSettings) {
    m := &Win32LobAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppAssignmentSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWin32LobAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppAssignmentSettings(), nil
}
// GetAutoUpdateSettings gets the autoUpdateSettings property value. The auto-update settings to apply for this app assignment.
// returns a Win32LobAppAutoUpdateSettingsable when successful
func (m *Win32LobAppAssignmentSettings) GetAutoUpdateSettings()(Win32LobAppAutoUpdateSettingsable) {
    return m.autoUpdateSettings
}
// GetDeliveryOptimizationPriority gets the deliveryOptimizationPriority property value. Contains value for delivery optimization priority.
// returns a *Win32LobAppDeliveryOptimizationPriority when successful
func (m *Win32LobAppAssignmentSettings) GetDeliveryOptimizationPriority()(*Win32LobAppDeliveryOptimizationPriority) {
    return m.deliveryOptimizationPriority
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Win32LobAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["autoUpdateSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWin32LobAppAutoUpdateSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoUpdateSettings(val.(Win32LobAppAutoUpdateSettingsable))
        }
        return nil
    }
    res["deliveryOptimizationPriority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppDeliveryOptimizationPriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryOptimizationPriority(val.(*Win32LobAppDeliveryOptimizationPriority))
        }
        return nil
    }
    res["installTimeSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppInstallTimeSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallTimeSettings(val.(MobileAppInstallTimeSettingsable))
        }
        return nil
    }
    res["notifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppNotification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifications(val.(*Win32LobAppNotification))
        }
        return nil
    }
    res["restartSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWin32LobAppRestartSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartSettings(val.(Win32LobAppRestartSettingsable))
        }
        return nil
    }
    return res
}
// GetInstallTimeSettings gets the installTimeSettings property value. The install time settings to apply for this app assignment.
// returns a MobileAppInstallTimeSettingsable when successful
func (m *Win32LobAppAssignmentSettings) GetInstallTimeSettings()(MobileAppInstallTimeSettingsable) {
    return m.installTimeSettings
}
// GetNotifications gets the notifications property value. Contains value for notification status.
// returns a *Win32LobAppNotification when successful
func (m *Win32LobAppAssignmentSettings) GetNotifications()(*Win32LobAppNotification) {
    return m.notifications
}
// GetRestartSettings gets the restartSettings property value. The reboot settings to apply for this app assignment.
// returns a Win32LobAppRestartSettingsable when successful
func (m *Win32LobAppAssignmentSettings) GetRestartSettings()(Win32LobAppRestartSettingsable) {
    return m.restartSettings
}
// Serialize serializes information the current object
func (m *Win32LobAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("autoUpdateSettings", m.GetAutoUpdateSettings())
        if err != nil {
            return err
        }
    }
    if m.GetDeliveryOptimizationPriority() != nil {
        cast := (*m.GetDeliveryOptimizationPriority()).String()
        err = writer.WriteStringValue("deliveryOptimizationPriority", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("installTimeSettings", m.GetInstallTimeSettings())
        if err != nil {
            return err
        }
    }
    if m.GetNotifications() != nil {
        cast := (*m.GetNotifications()).String()
        err = writer.WriteStringValue("notifications", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("restartSettings", m.GetRestartSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAutoUpdateSettings sets the autoUpdateSettings property value. The auto-update settings to apply for this app assignment.
func (m *Win32LobAppAssignmentSettings) SetAutoUpdateSettings(value Win32LobAppAutoUpdateSettingsable)() {
    m.autoUpdateSettings = value
}
// SetDeliveryOptimizationPriority sets the deliveryOptimizationPriority property value. Contains value for delivery optimization priority.
func (m *Win32LobAppAssignmentSettings) SetDeliveryOptimizationPriority(value *Win32LobAppDeliveryOptimizationPriority)() {
    m.deliveryOptimizationPriority = value
}
// SetInstallTimeSettings sets the installTimeSettings property value. The install time settings to apply for this app assignment.
func (m *Win32LobAppAssignmentSettings) SetInstallTimeSettings(value MobileAppInstallTimeSettingsable)() {
    m.installTimeSettings = value
}
// SetNotifications sets the notifications property value. Contains value for notification status.
func (m *Win32LobAppAssignmentSettings) SetNotifications(value *Win32LobAppNotification)() {
    m.notifications = value
}
// SetRestartSettings sets the restartSettings property value. The reboot settings to apply for this app assignment.
func (m *Win32LobAppAssignmentSettings) SetRestartSettings(value Win32LobAppRestartSettingsable)() {
    m.restartSettings = value
}
type Win32LobAppAssignmentSettingsable interface {
    MobileAppAssignmentSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoUpdateSettings()(Win32LobAppAutoUpdateSettingsable)
    GetDeliveryOptimizationPriority()(*Win32LobAppDeliveryOptimizationPriority)
    GetInstallTimeSettings()(MobileAppInstallTimeSettingsable)
    GetNotifications()(*Win32LobAppNotification)
    GetRestartSettings()(Win32LobAppRestartSettingsable)
    SetAutoUpdateSettings(value Win32LobAppAutoUpdateSettingsable)()
    SetDeliveryOptimizationPriority(value *Win32LobAppDeliveryOptimizationPriority)()
    SetInstallTimeSettings(value MobileAppInstallTimeSettingsable)()
    SetNotifications(value *Win32LobAppNotification)()
    SetRestartSettings(value Win32LobAppRestartSettingsable)()
}
