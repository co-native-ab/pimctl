package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAppX contains properties and inherited properties for Windows AppX Line Of Business apps.
type WindowsAppX struct {
    MobileLobApp
    // Contains properties for Windows architecture.
    applicableArchitectures *WindowsArchitecture
    // The identity name of the uploaded app package. For example: 'Contoso.DemoApp'.
    identityName *string
    // The identity publisher hash of the uploaded app package. This is the hash of the publisher from the manifest. For example: 'AB82CD0XYZ'.
    identityPublisherHash *string
    // The identity resource identifier of the uploaded app package. For example: 'TestResourceId'.
    identityResourceIdentifier *string
    // The identity version of the uploaded app package. For example: '1.0.0.0'.
    identityVersion *string
    // When TRUE, indicates that the app is a bundle. When FALSE, indicates that the app is not a bundle. By default, property is set to FALSE.
    isBundle *bool
    // The minimum operating system required for a Windows mobile app.
    minimumSupportedOperatingSystem WindowsMinimumOperatingSystemable
}
// NewWindowsAppX instantiates a new WindowsAppX and sets the default values.
func NewWindowsAppX()(*WindowsAppX) {
    m := &WindowsAppX{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsAppX"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsAppXFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsAppXFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAppX(), nil
}
// GetApplicableArchitectures gets the applicableArchitectures property value. Contains properties for Windows architecture.
// returns a *WindowsArchitecture when successful
func (m *WindowsAppX) GetApplicableArchitectures()(*WindowsArchitecture) {
    return m.applicableArchitectures
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsAppX) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["applicableArchitectures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableArchitectures(val.(*WindowsArchitecture))
        }
        return nil
    }
    res["identityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityName(val)
        }
        return nil
    }
    res["identityPublisherHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityPublisherHash(val)
        }
        return nil
    }
    res["identityResourceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityResourceIdentifier(val)
        }
        return nil
    }
    res["identityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityVersion(val)
        }
        return nil
    }
    res["isBundle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBundle(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(WindowsMinimumOperatingSystemable))
        }
        return nil
    }
    return res
}
// GetIdentityName gets the identityName property value. The identity name of the uploaded app package. For example: 'Contoso.DemoApp'.
// returns a *string when successful
func (m *WindowsAppX) GetIdentityName()(*string) {
    return m.identityName
}
// GetIdentityPublisherHash gets the identityPublisherHash property value. The identity publisher hash of the uploaded app package. This is the hash of the publisher from the manifest. For example: 'AB82CD0XYZ'.
// returns a *string when successful
func (m *WindowsAppX) GetIdentityPublisherHash()(*string) {
    return m.identityPublisherHash
}
// GetIdentityResourceIdentifier gets the identityResourceIdentifier property value. The identity resource identifier of the uploaded app package. For example: 'TestResourceId'.
// returns a *string when successful
func (m *WindowsAppX) GetIdentityResourceIdentifier()(*string) {
    return m.identityResourceIdentifier
}
// GetIdentityVersion gets the identityVersion property value. The identity version of the uploaded app package. For example: '1.0.0.0'.
// returns a *string when successful
func (m *WindowsAppX) GetIdentityVersion()(*string) {
    return m.identityVersion
}
// GetIsBundle gets the isBundle property value. When TRUE, indicates that the app is a bundle. When FALSE, indicates that the app is not a bundle. By default, property is set to FALSE.
// returns a *bool when successful
func (m *WindowsAppX) GetIsBundle()(*bool) {
    return m.isBundle
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
// returns a WindowsMinimumOperatingSystemable when successful
func (m *WindowsAppX) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// Serialize serializes information the current object
func (m *WindowsAppX) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableArchitectures() != nil {
        cast := (*m.GetApplicableArchitectures()).String()
        err = writer.WriteStringValue("applicableArchitectures", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityName", m.GetIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityPublisherHash", m.GetIdentityPublisherHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityResourceIdentifier", m.GetIdentityResourceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBundle", m.GetIsBundle())
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
    return nil
}
// SetApplicableArchitectures sets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *WindowsAppX) SetApplicableArchitectures(value *WindowsArchitecture)() {
    m.applicableArchitectures = value
}
// SetIdentityName sets the identityName property value. The identity name of the uploaded app package. For example: 'Contoso.DemoApp'.
func (m *WindowsAppX) SetIdentityName(value *string)() {
    m.identityName = value
}
// SetIdentityPublisherHash sets the identityPublisherHash property value. The identity publisher hash of the uploaded app package. This is the hash of the publisher from the manifest. For example: 'AB82CD0XYZ'.
func (m *WindowsAppX) SetIdentityPublisherHash(value *string)() {
    m.identityPublisherHash = value
}
// SetIdentityResourceIdentifier sets the identityResourceIdentifier property value. The identity resource identifier of the uploaded app package. For example: 'TestResourceId'.
func (m *WindowsAppX) SetIdentityResourceIdentifier(value *string)() {
    m.identityResourceIdentifier = value
}
// SetIdentityVersion sets the identityVersion property value. The identity version of the uploaded app package. For example: '1.0.0.0'.
func (m *WindowsAppX) SetIdentityVersion(value *string)() {
    m.identityVersion = value
}
// SetIsBundle sets the isBundle property value. When TRUE, indicates that the app is a bundle. When FALSE, indicates that the app is not a bundle. By default, property is set to FALSE.
func (m *WindowsAppX) SetIsBundle(value *bool)() {
    m.isBundle = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsAppX) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
type WindowsAppXable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableArchitectures()(*WindowsArchitecture)
    GetIdentityName()(*string)
    GetIdentityPublisherHash()(*string)
    GetIdentityResourceIdentifier()(*string)
    GetIdentityVersion()(*string)
    GetIsBundle()(*bool)
    GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable)
    SetApplicableArchitectures(value *WindowsArchitecture)()
    SetIdentityName(value *string)()
    SetIdentityPublisherHash(value *string)()
    SetIdentityResourceIdentifier(value *string)()
    SetIdentityVersion(value *string)()
    SetIsBundle(value *bool)()
    SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)()
}
