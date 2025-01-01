package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConditionalAccessAuthenticationFlows struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The transferMethods property
    transferMethods *ConditionalAccessTransferMethods
}
// NewConditionalAccessAuthenticationFlows instantiates a new ConditionalAccessAuthenticationFlows and sets the default values.
func NewConditionalAccessAuthenticationFlows()(*ConditionalAccessAuthenticationFlows) {
    m := &ConditionalAccessAuthenticationFlows{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessAuthenticationFlowsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalAccessAuthenticationFlowsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessAuthenticationFlows(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConditionalAccessAuthenticationFlows) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionalAccessAuthenticationFlows) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["transferMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessTransferMethods)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferMethods(val.(*ConditionalAccessTransferMethods))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ConditionalAccessAuthenticationFlows) GetOdataType()(*string) {
    return m.odataType
}
// GetTransferMethods gets the transferMethods property value. The transferMethods property
// returns a *ConditionalAccessTransferMethods when successful
func (m *ConditionalAccessAuthenticationFlows) GetTransferMethods()(*ConditionalAccessTransferMethods) {
    return m.transferMethods
}
// Serialize serializes information the current object
func (m *ConditionalAccessAuthenticationFlows) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetTransferMethods() != nil {
        cast := (*m.GetTransferMethods()).String()
        err := writer.WriteStringValue("transferMethods", &cast)
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
func (m *ConditionalAccessAuthenticationFlows) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessAuthenticationFlows) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTransferMethods sets the transferMethods property value. The transferMethods property
func (m *ConditionalAccessAuthenticationFlows) SetTransferMethods(value *ConditionalAccessTransferMethods)() {
    m.transferMethods = value
}
type ConditionalAccessAuthenticationFlowsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetTransferMethods()(*ConditionalAccessTransferMethods)
    SetOdataType(value *string)()
    SetTransferMethods(value *ConditionalAccessTransferMethods)()
}
