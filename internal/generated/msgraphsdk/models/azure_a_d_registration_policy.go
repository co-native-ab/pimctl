package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AzureADRegistrationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allowedToRegister property
    allowedToRegister DeviceRegistrationMembershipable
    // The isAdminConfigurable property
    isAdminConfigurable *bool
    // The OdataType property
    odataType *string
}
// NewAzureADRegistrationPolicy instantiates a new AzureADRegistrationPolicy and sets the default values.
func NewAzureADRegistrationPolicy()(*AzureADRegistrationPolicy) {
    m := &AzureADRegistrationPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAzureADRegistrationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureADRegistrationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADRegistrationPolicy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AzureADRegistrationPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedToRegister gets the allowedToRegister property value. The allowedToRegister property
// returns a DeviceRegistrationMembershipable when successful
func (m *AzureADRegistrationPolicy) GetAllowedToRegister()(DeviceRegistrationMembershipable) {
    return m.allowedToRegister
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureADRegistrationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedToRegister"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceRegistrationMembershipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToRegister(val.(DeviceRegistrationMembershipable))
        }
        return nil
    }
    res["isAdminConfigurable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAdminConfigurable(val)
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
// GetIsAdminConfigurable gets the isAdminConfigurable property value. The isAdminConfigurable property
// returns a *bool when successful
func (m *AzureADRegistrationPolicy) GetIsAdminConfigurable()(*bool) {
    return m.isAdminConfigurable
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AzureADRegistrationPolicy) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AzureADRegistrationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("allowedToRegister", m.GetAllowedToRegister())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAdminConfigurable", m.GetIsAdminConfigurable())
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
func (m *AzureADRegistrationPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedToRegister sets the allowedToRegister property value. The allowedToRegister property
func (m *AzureADRegistrationPolicy) SetAllowedToRegister(value DeviceRegistrationMembershipable)() {
    m.allowedToRegister = value
}
// SetIsAdminConfigurable sets the isAdminConfigurable property value. The isAdminConfigurable property
func (m *AzureADRegistrationPolicy) SetIsAdminConfigurable(value *bool)() {
    m.isAdminConfigurable = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AzureADRegistrationPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
type AzureADRegistrationPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedToRegister()(DeviceRegistrationMembershipable)
    GetIsAdminConfigurable()(*bool)
    GetOdataType()(*string)
    SetAllowedToRegister(value DeviceRegistrationMembershipable)()
    SetIsAdminConfigurable(value *bool)()
    SetOdataType(value *string)()
}
