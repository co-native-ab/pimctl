package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RegistryValueEvidence struct {
    AlertEvidence
    // A unique identifier assigned to a device by Microsoft Defender for Endpoint.
    mdeDeviceId *string
    // Registry hive of the key that the recorded action was applied to.
    registryHive *string
    // Registry key that the recorded action was applied to.
    registryKey *string
    // Data of the registry value that the recorded action was applied to.
    registryValue *string
    // Name of the registry value that the recorded action was applied to.
    registryValueName *string
    // Data type, such as binary or string, of the registry value that the recorded action was applied to.
    registryValueType *string
}
// NewRegistryValueEvidence instantiates a new RegistryValueEvidence and sets the default values.
func NewRegistryValueEvidence()(*RegistryValueEvidence) {
    m := &RegistryValueEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.registryValueEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRegistryValueEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRegistryValueEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegistryValueEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RegistryValueEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["mdeDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdeDeviceId(val)
        }
        return nil
    }
    res["registryHive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryHive(val)
        }
        return nil
    }
    res["registryKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryKey(val)
        }
        return nil
    }
    res["registryValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryValue(val)
        }
        return nil
    }
    res["registryValueName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryValueName(val)
        }
        return nil
    }
    res["registryValueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryValueType(val)
        }
        return nil
    }
    return res
}
// GetMdeDeviceId gets the mdeDeviceId property value. A unique identifier assigned to a device by Microsoft Defender for Endpoint.
// returns a *string when successful
func (m *RegistryValueEvidence) GetMdeDeviceId()(*string) {
    return m.mdeDeviceId
}
// GetRegistryHive gets the registryHive property value. Registry hive of the key that the recorded action was applied to.
// returns a *string when successful
func (m *RegistryValueEvidence) GetRegistryHive()(*string) {
    return m.registryHive
}
// GetRegistryKey gets the registryKey property value. Registry key that the recorded action was applied to.
// returns a *string when successful
func (m *RegistryValueEvidence) GetRegistryKey()(*string) {
    return m.registryKey
}
// GetRegistryValue gets the registryValue property value. Data of the registry value that the recorded action was applied to.
// returns a *string when successful
func (m *RegistryValueEvidence) GetRegistryValue()(*string) {
    return m.registryValue
}
// GetRegistryValueName gets the registryValueName property value. Name of the registry value that the recorded action was applied to.
// returns a *string when successful
func (m *RegistryValueEvidence) GetRegistryValueName()(*string) {
    return m.registryValueName
}
// GetRegistryValueType gets the registryValueType property value. Data type, such as binary or string, of the registry value that the recorded action was applied to.
// returns a *string when successful
func (m *RegistryValueEvidence) GetRegistryValueType()(*string) {
    return m.registryValueType
}
// Serialize serializes information the current object
func (m *RegistryValueEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("mdeDeviceId", m.GetMdeDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registryHive", m.GetRegistryHive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registryKey", m.GetRegistryKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registryValue", m.GetRegistryValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registryValueName", m.GetRegistryValueName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registryValueType", m.GetRegistryValueType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMdeDeviceId sets the mdeDeviceId property value. A unique identifier assigned to a device by Microsoft Defender for Endpoint.
func (m *RegistryValueEvidence) SetMdeDeviceId(value *string)() {
    m.mdeDeviceId = value
}
// SetRegistryHive sets the registryHive property value. Registry hive of the key that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryHive(value *string)() {
    m.registryHive = value
}
// SetRegistryKey sets the registryKey property value. Registry key that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryKey(value *string)() {
    m.registryKey = value
}
// SetRegistryValue sets the registryValue property value. Data of the registry value that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryValue(value *string)() {
    m.registryValue = value
}
// SetRegistryValueName sets the registryValueName property value. Name of the registry value that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryValueName(value *string)() {
    m.registryValueName = value
}
// SetRegistryValueType sets the registryValueType property value. Data type, such as binary or string, of the registry value that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryValueType(value *string)() {
    m.registryValueType = value
}
type RegistryValueEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMdeDeviceId()(*string)
    GetRegistryHive()(*string)
    GetRegistryKey()(*string)
    GetRegistryValue()(*string)
    GetRegistryValueName()(*string)
    GetRegistryValueType()(*string)
    SetMdeDeviceId(value *string)()
    SetRegistryHive(value *string)()
    SetRegistryKey(value *string)()
    SetRegistryValue(value *string)()
    SetRegistryValueName(value *string)()
    SetRegistryValueType(value *string)()
}
