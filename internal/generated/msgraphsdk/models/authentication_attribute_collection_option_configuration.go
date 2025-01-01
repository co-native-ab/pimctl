package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationAttributeCollectionOptionConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The label of the option that will be displayed to user, unless overridden.
    label *string
    // The OdataType property
    odataType *string
    // The value of the option that will be stored.
    value *string
}
// NewAuthenticationAttributeCollectionOptionConfiguration instantiates a new AuthenticationAttributeCollectionOptionConfiguration and sets the default values.
func NewAuthenticationAttributeCollectionOptionConfiguration()(*AuthenticationAttributeCollectionOptionConfiguration) {
    m := &AuthenticationAttributeCollectionOptionConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationAttributeCollectionOptionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationAttributeCollectionOptionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationAttributeCollectionOptionConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthenticationAttributeCollectionOptionConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationAttributeCollectionOptionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetLabel gets the label property value. The label of the option that will be displayed to user, unless overridden.
// returns a *string when successful
func (m *AuthenticationAttributeCollectionOptionConfiguration) GetLabel()(*string) {
    return m.label
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AuthenticationAttributeCollectionOptionConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The value of the option that will be stored.
// returns a *string when successful
func (m *AuthenticationAttributeCollectionOptionConfiguration) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *AuthenticationAttributeCollectionOptionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("label", m.GetLabel())
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
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *AuthenticationAttributeCollectionOptionConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabel sets the label property value. The label of the option that will be displayed to user, unless overridden.
func (m *AuthenticationAttributeCollectionOptionConfiguration) SetLabel(value *string)() {
    m.label = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationAttributeCollectionOptionConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The value of the option that will be stored.
func (m *AuthenticationAttributeCollectionOptionConfiguration) SetValue(value *string)() {
    m.value = value
}
type AuthenticationAttributeCollectionOptionConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabel()(*string)
    GetOdataType()(*string)
    GetValue()(*string)
    SetLabel(value *string)()
    SetOdataType(value *string)()
    SetValue(value *string)()
}
