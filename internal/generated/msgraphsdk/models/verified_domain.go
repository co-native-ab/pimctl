package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VerifiedDomain struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // For example, Email, OfficeCommunicationsOnline.
    capabilities *string
    // true if this is the default domain associated with the tenant; otherwise, false.
    isDefault *bool
    // true if this is the initial domain associated with the tenant; otherwise, false.
    isInitial *bool
    // The domain name; for example, contoso.com.
    name *string
    // The OdataType property
    odataType *string
    // For example, Managed.
    typeEscaped *string
}
// NewVerifiedDomain instantiates a new VerifiedDomain and sets the default values.
func NewVerifiedDomain()(*VerifiedDomain) {
    m := &VerifiedDomain{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVerifiedDomainFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVerifiedDomainFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifiedDomain(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VerifiedDomain) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapabilities gets the capabilities property value. For example, Email, OfficeCommunicationsOnline.
// returns a *string when successful
func (m *VerifiedDomain) GetCapabilities()(*string) {
    return m.capabilities
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VerifiedDomain) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilities(val)
        }
        return nil
    }
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isInitial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInitial(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. true if this is the default domain associated with the tenant; otherwise, false.
// returns a *bool when successful
func (m *VerifiedDomain) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetIsInitial gets the isInitial property value. true if this is the initial domain associated with the tenant; otherwise, false.
// returns a *bool when successful
func (m *VerifiedDomain) GetIsInitial()(*bool) {
    return m.isInitial
}
// GetName gets the name property value. The domain name; for example, contoso.com.
// returns a *string when successful
func (m *VerifiedDomain) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *VerifiedDomain) GetOdataType()(*string) {
    return m.odataType
}
// GetTypeEscaped gets the type property value. For example, Managed.
// returns a *string when successful
func (m *VerifiedDomain) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *VerifiedDomain) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("capabilities", m.GetCapabilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isInitial", m.GetIsInitial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *VerifiedDomain) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapabilities sets the capabilities property value. For example, Email, OfficeCommunicationsOnline.
func (m *VerifiedDomain) SetCapabilities(value *string)() {
    m.capabilities = value
}
// SetIsDefault sets the isDefault property value. true if this is the default domain associated with the tenant; otherwise, false.
func (m *VerifiedDomain) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetIsInitial sets the isInitial property value. true if this is the initial domain associated with the tenant; otherwise, false.
func (m *VerifiedDomain) SetIsInitial(value *bool)() {
    m.isInitial = value
}
// SetName sets the name property value. The domain name; for example, contoso.com.
func (m *VerifiedDomain) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VerifiedDomain) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTypeEscaped sets the type property value. For example, Managed.
func (m *VerifiedDomain) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type VerifiedDomainable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapabilities()(*string)
    GetIsDefault()(*bool)
    GetIsInitial()(*bool)
    GetName()(*string)
    GetOdataType()(*string)
    GetTypeEscaped()(*string)
    SetCapabilities(value *string)()
    SetIsDefault(value *bool)()
    SetIsInitial(value *bool)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetTypeEscaped(value *string)()
}
