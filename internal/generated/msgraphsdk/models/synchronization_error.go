package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SynchronizationError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error code. For example, AzureDirectoryB2BManagementPolicyCheckFailure.
    code *string
    // The error message. For example, Policy permitting auto-redemption of invitations not configured.
    message *string
    // The OdataType property
    odataType *string
    // The action to take to resolve the error. For example, false.
    tenantActionable *bool
}
// NewSynchronizationError instantiates a new SynchronizationError and sets the default values.
func NewSynchronizationError()(*SynchronizationError) {
    m := &SynchronizationError{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSynchronizationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSynchronizationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationError(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SynchronizationError) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The error code. For example, AzureDirectoryB2BManagementPolicyCheckFailure.
// returns a *string when successful
func (m *SynchronizationError) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SynchronizationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
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
    res["tenantActionable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantActionable(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The error message. For example, Policy permitting auto-redemption of invitations not configured.
// returns a *string when successful
func (m *SynchronizationError) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SynchronizationError) GetOdataType()(*string) {
    return m.odataType
}
// GetTenantActionable gets the tenantActionable property value. The action to take to resolve the error. For example, false.
// returns a *bool when successful
func (m *SynchronizationError) GetTenantActionable()(*bool) {
    return m.tenantActionable
}
// Serialize serializes information the current object
func (m *SynchronizationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
        err := writer.WriteBoolValue("tenantActionable", m.GetTenantActionable())
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
func (m *SynchronizationError) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The error code. For example, AzureDirectoryB2BManagementPolicyCheckFailure.
func (m *SynchronizationError) SetCode(value *string)() {
    m.code = value
}
// SetMessage sets the message property value. The error message. For example, Policy permitting auto-redemption of invitations not configured.
func (m *SynchronizationError) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationError) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTenantActionable sets the tenantActionable property value. The action to take to resolve the error. For example, false.
func (m *SynchronizationError) SetTenantActionable(value *bool)() {
    m.tenantActionable = value
}
type SynchronizationErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetMessage()(*string)
    GetOdataType()(*string)
    GetTenantActionable()(*bool)
    SetCode(value *string)()
    SetMessage(value *string)()
    SetOdataType(value *string)()
    SetTenantActionable(value *bool)()
}
