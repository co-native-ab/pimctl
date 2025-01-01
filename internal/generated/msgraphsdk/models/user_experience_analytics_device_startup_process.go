package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceStartupProcess the user experience analytics device startup process details.
type UserExperienceAnalyticsDeviceStartupProcess struct {
    Entity
    // The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
    managedDeviceId *string
    // The name of the process. Examples: outlook, excel. Supports: $select, $OrderBy. Read-only.
    processName *string
    // The product name of the process. Examples: Microsoft Outlook, Microsoft Excel. Supports: $select, $OrderBy. Read-only.
    productName *string
    // The publisher of the process. Examples: Microsoft Corporation, Contoso Corp. Supports: $select, $OrderBy. Read-only.
    publisher *string
    // The impact of startup process on device boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
    startupImpactInMs *int32
}
// NewUserExperienceAnalyticsDeviceStartupProcess instantiates a new UserExperienceAnalyticsDeviceStartupProcess and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupProcess()(*UserExperienceAnalyticsDeviceStartupProcess) {
    m := &UserExperienceAnalyticsDeviceStartupProcess{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceStartupProcess(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["processName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessName(val)
        }
        return nil
    }
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["startupImpactInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupImpactInMs(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceId gets the managedDeviceId property value. The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetProcessName gets the processName property value. The name of the process. Examples: outlook, excel. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetProcessName()(*string) {
    return m.processName
}
// GetProductName gets the productName property value. The product name of the process. Examples: Microsoft Outlook, Microsoft Excel. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetProductName()(*string) {
    return m.productName
}
// GetPublisher gets the publisher property value. The publisher of the process. Examples: Microsoft Corporation, Contoso Corp. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetPublisher()(*string) {
    return m.publisher
}
// GetStartupImpactInMs gets the startupImpactInMs property value. The impact of startup process on device boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetStartupImpactInMs()(*int32) {
    return m.startupImpactInMs
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceStartupProcess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("processName", m.GetProcessName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("startupImpactInMs", m.GetStartupImpactInMs())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManagedDeviceId sets the managedDeviceId property value. The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetProcessName sets the processName property value. The name of the process. Examples: outlook, excel. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetProcessName(value *string)() {
    m.processName = value
}
// SetProductName sets the productName property value. The product name of the process. Examples: Microsoft Outlook, Microsoft Excel. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetProductName(value *string)() {
    m.productName = value
}
// SetPublisher sets the publisher property value. The publisher of the process. Examples: Microsoft Corporation, Contoso Corp. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetPublisher(value *string)() {
    m.publisher = value
}
// SetStartupImpactInMs sets the startupImpactInMs property value. The impact of startup process on device boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetStartupImpactInMs(value *int32)() {
    m.startupImpactInMs = value
}
type UserExperienceAnalyticsDeviceStartupProcessable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManagedDeviceId()(*string)
    GetProcessName()(*string)
    GetProductName()(*string)
    GetPublisher()(*string)
    GetStartupImpactInMs()(*int32)
    SetManagedDeviceId(value *string)()
    SetProcessName(value *string)()
    SetProductName(value *string)()
    SetPublisher(value *string)()
    SetStartupImpactInMs(value *int32)()
}
