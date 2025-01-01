package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeviceManagementSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of days a device is allowed to go without checking in to remain compliant.
    deviceComplianceCheckinThresholdDays *int32
    // Is feature enabled or not for scheduled action for rule.
    isScheduledActionEnabled *bool
    // The OdataType property
    odataType *string
    // Device should be noncompliant when there is no compliance policy targeted when this is true
    secureByDefault *bool
}
// NewDeviceManagementSettings instantiates a new DeviceManagementSettings and sets the default values.
func NewDeviceManagementSettings()(*DeviceManagementSettings) {
    m := &DeviceManagementSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceManagementSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceComplianceCheckinThresholdDays gets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
// returns a *int32 when successful
func (m *DeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays()(*int32) {
    return m.deviceComplianceCheckinThresholdDays
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceComplianceCheckinThresholdDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceCheckinThresholdDays(val)
        }
        return nil
    }
    res["isScheduledActionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScheduledActionEnabled(val)
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
    res["secureByDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureByDefault(val)
        }
        return nil
    }
    return res
}
// GetIsScheduledActionEnabled gets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
// returns a *bool when successful
func (m *DeviceManagementSettings) GetIsScheduledActionEnabled()(*bool) {
    return m.isScheduledActionEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DeviceManagementSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetSecureByDefault gets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
// returns a *bool when successful
func (m *DeviceManagementSettings) GetSecureByDefault()(*bool) {
    return m.secureByDefault
}
// Serialize serializes information the current object
func (m *DeviceManagementSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("deviceComplianceCheckinThresholdDays", m.GetDeviceComplianceCheckinThresholdDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isScheduledActionEnabled", m.GetIsScheduledActionEnabled())
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
        err := writer.WriteBoolValue("secureByDefault", m.GetSecureByDefault())
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
func (m *DeviceManagementSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceComplianceCheckinThresholdDays sets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
func (m *DeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(value *int32)() {
    m.deviceComplianceCheckinThresholdDays = value
}
// SetIsScheduledActionEnabled sets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
func (m *DeviceManagementSettings) SetIsScheduledActionEnabled(value *bool)() {
    m.isScheduledActionEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSecureByDefault sets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
func (m *DeviceManagementSettings) SetSecureByDefault(value *bool)() {
    m.secureByDefault = value
}
type DeviceManagementSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceComplianceCheckinThresholdDays()(*int32)
    GetIsScheduledActionEnabled()(*bool)
    GetOdataType()(*string)
    GetSecureByDefault()(*bool)
    SetDeviceComplianceCheckinThresholdDays(value *int32)()
    SetIsScheduledActionEnabled(value *bool)()
    SetOdataType(value *string)()
    SetSecureByDefault(value *bool)()
}
