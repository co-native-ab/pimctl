package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSDmgApp contains properties and inherited properties for the MacOS DMG (Apple Disk Image) App.
type MacOSDmgApp struct {
    MobileLobApp
    // When TRUE, indicates that the app's version will NOT be used to detect if the app is installed on a device. When FALSE, indicates that the app's version will be used to detect if the app is installed on a device. Set this to true for apps that use a self update feature. The default value is FALSE.
    ignoreVersionDetection *bool
    // The list of .apps expected to be installed by the DMG (Apple Disk Image). This collection can contain a maximum of 500 elements.
    includedApps []MacOSIncludedAppable
    // ComplexType macOSMinimumOperatingSystem that indicates the minimum operating system applicable for the application.
    minimumSupportedOperatingSystem MacOSMinimumOperatingSystemable
    // The bundleId of the primary .app in the DMG (Apple Disk Image). This maps to the CFBundleIdentifier in the app's bundle configuration.
    primaryBundleId *string
    // The version of the primary .app in the DMG (Apple Disk Image). This maps to the CFBundleShortVersion in the app's bundle configuration.
    primaryBundleVersion *string
}
// NewMacOSDmgApp instantiates a new MacOSDmgApp and sets the default values.
func NewMacOSDmgApp()(*MacOSDmgApp) {
    m := &MacOSDmgApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.macOSDmgApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSDmgAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMacOSDmgAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSDmgApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MacOSDmgApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["ignoreVersionDetection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreVersionDetection(val)
        }
        return nil
    }
    res["includedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSIncludedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSIncludedAppable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOSIncludedAppable)
                }
            }
            m.SetIncludedApps(res)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(MacOSMinimumOperatingSystemable))
        }
        return nil
    }
    res["primaryBundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryBundleId(val)
        }
        return nil
    }
    res["primaryBundleVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryBundleVersion(val)
        }
        return nil
    }
    return res
}
// GetIgnoreVersionDetection gets the ignoreVersionDetection property value. When TRUE, indicates that the app's version will NOT be used to detect if the app is installed on a device. When FALSE, indicates that the app's version will be used to detect if the app is installed on a device. Set this to true for apps that use a self update feature. The default value is FALSE.
// returns a *bool when successful
func (m *MacOSDmgApp) GetIgnoreVersionDetection()(*bool) {
    return m.ignoreVersionDetection
}
// GetIncludedApps gets the includedApps property value. The list of .apps expected to be installed by the DMG (Apple Disk Image). This collection can contain a maximum of 500 elements.
// returns a []MacOSIncludedAppable when successful
func (m *MacOSDmgApp) GetIncludedApps()([]MacOSIncludedAppable) {
    return m.includedApps
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. ComplexType macOSMinimumOperatingSystem that indicates the minimum operating system applicable for the application.
// returns a MacOSMinimumOperatingSystemable when successful
func (m *MacOSDmgApp) GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// GetPrimaryBundleId gets the primaryBundleId property value. The bundleId of the primary .app in the DMG (Apple Disk Image). This maps to the CFBundleIdentifier in the app's bundle configuration.
// returns a *string when successful
func (m *MacOSDmgApp) GetPrimaryBundleId()(*string) {
    return m.primaryBundleId
}
// GetPrimaryBundleVersion gets the primaryBundleVersion property value. The version of the primary .app in the DMG (Apple Disk Image). This maps to the CFBundleShortVersion in the app's bundle configuration.
// returns a *string when successful
func (m *MacOSDmgApp) GetPrimaryBundleVersion()(*string) {
    return m.primaryBundleVersion
}
// Serialize serializes information the current object
func (m *MacOSDmgApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("ignoreVersionDetection", m.GetIgnoreVersionDetection())
        if err != nil {
            return err
        }
    }
    if m.GetIncludedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludedApps()))
        for i, v := range m.GetIncludedApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("includedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryBundleId", m.GetPrimaryBundleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryBundleVersion", m.GetPrimaryBundleVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIgnoreVersionDetection sets the ignoreVersionDetection property value. When TRUE, indicates that the app's version will NOT be used to detect if the app is installed on a device. When FALSE, indicates that the app's version will be used to detect if the app is installed on a device. Set this to true for apps that use a self update feature. The default value is FALSE.
func (m *MacOSDmgApp) SetIgnoreVersionDetection(value *bool)() {
    m.ignoreVersionDetection = value
}
// SetIncludedApps sets the includedApps property value. The list of .apps expected to be installed by the DMG (Apple Disk Image). This collection can contain a maximum of 500 elements.
func (m *MacOSDmgApp) SetIncludedApps(value []MacOSIncludedAppable)() {
    m.includedApps = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. ComplexType macOSMinimumOperatingSystem that indicates the minimum operating system applicable for the application.
func (m *MacOSDmgApp) SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
// SetPrimaryBundleId sets the primaryBundleId property value. The bundleId of the primary .app in the DMG (Apple Disk Image). This maps to the CFBundleIdentifier in the app's bundle configuration.
func (m *MacOSDmgApp) SetPrimaryBundleId(value *string)() {
    m.primaryBundleId = value
}
// SetPrimaryBundleVersion sets the primaryBundleVersion property value. The version of the primary .app in the DMG (Apple Disk Image). This maps to the CFBundleShortVersion in the app's bundle configuration.
func (m *MacOSDmgApp) SetPrimaryBundleVersion(value *string)() {
    m.primaryBundleVersion = value
}
type MacOSDmgAppable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIgnoreVersionDetection()(*bool)
    GetIncludedApps()([]MacOSIncludedAppable)
    GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable)
    GetPrimaryBundleId()(*string)
    GetPrimaryBundleVersion()(*string)
    SetIgnoreVersionDetection(value *bool)()
    SetIncludedApps(value []MacOSIncludedAppable)()
    SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)()
    SetPrimaryBundleId(value *string)()
    SetPrimaryBundleVersion(value *string)()
}
