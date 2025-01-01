package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConditionalAccessFilter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The mode property
    mode *FilterMode
    // The OdataType property
    odataType *string
    // Rule syntax is similar to that used for membership rules for groups in Microsoft Entra ID. For details, see rules with multiple expressions
    rule *string
}
// NewConditionalAccessFilter instantiates a new ConditionalAccessFilter and sets the default values.
func NewConditionalAccessFilter()(*ConditionalAccessFilter) {
    m := &ConditionalAccessFilter{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalAccessFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessFilter(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConditionalAccessFilter) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionalAccessFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFilterMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMode(val.(*FilterMode))
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
// GetMode gets the mode property value. The mode property
// returns a *FilterMode when successful
func (m *ConditionalAccessFilter) GetMode()(*FilterMode) {
    return m.mode
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ConditionalAccessFilter) GetOdataType()(*string) {
    return m.odataType
}
// GetRule gets the rule property value. Rule syntax is similar to that used for membership rules for groups in Microsoft Entra ID. For details, see rules with multiple expressions
// returns a *string when successful
func (m *ConditionalAccessFilter) GetRule()(*string) {
    return m.rule
}
// Serialize serializes information the current object
func (m *ConditionalAccessFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ConditionalAccessFilter) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMode sets the mode property value. The mode property
func (m *ConditionalAccessFilter) SetMode(value *FilterMode)() {
    m.mode = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessFilter) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRule sets the rule property value. Rule syntax is similar to that used for membership rules for groups in Microsoft Entra ID. For details, see rules with multiple expressions
func (m *ConditionalAccessFilter) SetRule(value *string)() {
    m.rule = value
}
type ConditionalAccessFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMode()(*FilterMode)
    GetOdataType()(*string)
    GetRule()(*string)
    SetMode(value *FilterMode)()
    SetOdataType(value *string)()
    SetRule(value *string)()
}
