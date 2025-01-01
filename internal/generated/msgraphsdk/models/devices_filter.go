package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DevicesFilter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Determines whether devices that satisfy the rule should be allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
    mode *CrossTenantAccessPolicyTargetConfigurationAccessType
    // The OdataType property
    odataType *string
    // Defines the rule to filter the devices. For example, device.deviceAttribute2 -eq 'PrivilegedAccessWorkstation'.
    rule *string
}
// NewDevicesFilter instantiates a new DevicesFilter and sets the default values.
func NewDevicesFilter()(*DevicesFilter) {
    m := &DevicesFilter{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDevicesFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDevicesFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDevicesFilter(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DevicesFilter) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DevicesFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCrossTenantAccessPolicyTargetConfigurationAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMode(val.(*CrossTenantAccessPolicyTargetConfigurationAccessType))
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
    res["rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRule(val)
        }
        return nil
    }
    return res
}
// GetMode gets the mode property value. Determines whether devices that satisfy the rule should be allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
// returns a *CrossTenantAccessPolicyTargetConfigurationAccessType when successful
func (m *DevicesFilter) GetMode()(*CrossTenantAccessPolicyTargetConfigurationAccessType) {
    return m.mode
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DevicesFilter) GetOdataType()(*string) {
    return m.odataType
}
// GetRule gets the rule property value. Defines the rule to filter the devices. For example, device.deviceAttribute2 -eq 'PrivilegedAccessWorkstation'.
// returns a *string when successful
func (m *DevicesFilter) GetRule()(*string) {
    return m.rule
}
// Serialize serializes information the current object
func (m *DevicesFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMode() != nil {
        cast := (*m.GetMode()).String()
        err := writer.WriteStringValue("mode", &cast)
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
        err := writer.WriteStringValue("rule", m.GetRule())
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
func (m *DevicesFilter) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMode sets the mode property value. Determines whether devices that satisfy the rule should be allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
func (m *DevicesFilter) SetMode(value *CrossTenantAccessPolicyTargetConfigurationAccessType)() {
    m.mode = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DevicesFilter) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRule sets the rule property value. Defines the rule to filter the devices. For example, device.deviceAttribute2 -eq 'PrivilegedAccessWorkstation'.
func (m *DevicesFilter) SetRule(value *string)() {
    m.rule = value
}
type DevicesFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMode()(*CrossTenantAccessPolicyTargetConfigurationAccessType)
    GetOdataType()(*string)
    GetRule()(*string)
    SetMode(value *CrossTenantAccessPolicyTargetConfigurationAccessType)()
    SetOdataType(value *string)()
    SetRule(value *string)()
}
