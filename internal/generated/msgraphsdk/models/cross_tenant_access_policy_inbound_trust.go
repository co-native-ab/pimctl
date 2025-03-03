package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CrossTenantAccessPolicyInboundTrust struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies whether compliant devices from external Microsoft Entra organizations are trusted.
    isCompliantDeviceAccepted *bool
    // Specifies whether Microsoft Entra hybrid joined devices from external Microsoft Entra organizations are trusted.
    isHybridAzureADJoinedDeviceAccepted *bool
    // Specifies whether MFA from external Microsoft Entra organizations is trusted.
    isMfaAccepted *bool
    // The OdataType property
    odataType *string
}
// NewCrossTenantAccessPolicyInboundTrust instantiates a new CrossTenantAccessPolicyInboundTrust and sets the default values.
func NewCrossTenantAccessPolicyInboundTrust()(*CrossTenantAccessPolicyInboundTrust) {
    m := &CrossTenantAccessPolicyInboundTrust{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCrossTenantAccessPolicyInboundTrustFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCrossTenantAccessPolicyInboundTrustFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyInboundTrust(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CrossTenantAccessPolicyInboundTrust) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CrossTenantAccessPolicyInboundTrust) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isCompliantDeviceAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCompliantDeviceAccepted(val)
        }
        return nil
    }
    res["isHybridAzureADJoinedDeviceAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHybridAzureADJoinedDeviceAccepted(val)
        }
        return nil
    }
    res["isMfaAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaAccepted(val)
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
    return res
}
// GetIsCompliantDeviceAccepted gets the isCompliantDeviceAccepted property value. Specifies whether compliant devices from external Microsoft Entra organizations are trusted.
// returns a *bool when successful
func (m *CrossTenantAccessPolicyInboundTrust) GetIsCompliantDeviceAccepted()(*bool) {
    return m.isCompliantDeviceAccepted
}
// GetIsHybridAzureADJoinedDeviceAccepted gets the isHybridAzureADJoinedDeviceAccepted property value. Specifies whether Microsoft Entra hybrid joined devices from external Microsoft Entra organizations are trusted.
// returns a *bool when successful
func (m *CrossTenantAccessPolicyInboundTrust) GetIsHybridAzureADJoinedDeviceAccepted()(*bool) {
    return m.isHybridAzureADJoinedDeviceAccepted
}
// GetIsMfaAccepted gets the isMfaAccepted property value. Specifies whether MFA from external Microsoft Entra organizations is trusted.
// returns a *bool when successful
func (m *CrossTenantAccessPolicyInboundTrust) GetIsMfaAccepted()(*bool) {
    return m.isMfaAccepted
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CrossTenantAccessPolicyInboundTrust) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyInboundTrust) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isCompliantDeviceAccepted", m.GetIsCompliantDeviceAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHybridAzureADJoinedDeviceAccepted", m.GetIsHybridAzureADJoinedDeviceAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMfaAccepted", m.GetIsMfaAccepted())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyInboundTrust) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsCompliantDeviceAccepted sets the isCompliantDeviceAccepted property value. Specifies whether compliant devices from external Microsoft Entra organizations are trusted.
func (m *CrossTenantAccessPolicyInboundTrust) SetIsCompliantDeviceAccepted(value *bool)() {
    m.isCompliantDeviceAccepted = value
}
// SetIsHybridAzureADJoinedDeviceAccepted sets the isHybridAzureADJoinedDeviceAccepted property value. Specifies whether Microsoft Entra hybrid joined devices from external Microsoft Entra organizations are trusted.
func (m *CrossTenantAccessPolicyInboundTrust) SetIsHybridAzureADJoinedDeviceAccepted(value *bool)() {
    m.isHybridAzureADJoinedDeviceAccepted = value
}
// SetIsMfaAccepted sets the isMfaAccepted property value. Specifies whether MFA from external Microsoft Entra organizations is trusted.
func (m *CrossTenantAccessPolicyInboundTrust) SetIsMfaAccepted(value *bool)() {
    m.isMfaAccepted = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CrossTenantAccessPolicyInboundTrust) SetOdataType(value *string)() {
    m.odataType = value
}
type CrossTenantAccessPolicyInboundTrustable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsCompliantDeviceAccepted()(*bool)
    GetIsHybridAzureADJoinedDeviceAccepted()(*bool)
    GetIsMfaAccepted()(*bool)
    GetOdataType()(*string)
    SetIsCompliantDeviceAccepted(value *bool)()
    SetIsHybridAzureADJoinedDeviceAccepted(value *bool)()
    SetIsMfaAccepted(value *bool)()
    SetOdataType(value *string)()
}
