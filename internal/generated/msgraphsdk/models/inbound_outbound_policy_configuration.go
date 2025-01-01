package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InboundOutboundPolicyConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Defines whether external users coming inbound are allowed.
    inboundAllowed *bool
    // The OdataType property
    odataType *string
    // Defines whether internal users are allowed to go outbound.
    outboundAllowed *bool
}
// NewInboundOutboundPolicyConfiguration instantiates a new InboundOutboundPolicyConfiguration and sets the default values.
func NewInboundOutboundPolicyConfiguration()(*InboundOutboundPolicyConfiguration) {
    m := &InboundOutboundPolicyConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInboundOutboundPolicyConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInboundOutboundPolicyConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInboundOutboundPolicyConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InboundOutboundPolicyConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InboundOutboundPolicyConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["inboundAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundAllowed(val)
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
    res["outboundAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutboundAllowed(val)
        }
        return nil
    }
    return res
}
// GetInboundAllowed gets the inboundAllowed property value. Defines whether external users coming inbound are allowed.
// returns a *bool when successful
func (m *InboundOutboundPolicyConfiguration) GetInboundAllowed()(*bool) {
    return m.inboundAllowed
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *InboundOutboundPolicyConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetOutboundAllowed gets the outboundAllowed property value. Defines whether internal users are allowed to go outbound.
// returns a *bool when successful
func (m *InboundOutboundPolicyConfiguration) GetOutboundAllowed()(*bool) {
    return m.outboundAllowed
}
// Serialize serializes information the current object
func (m *InboundOutboundPolicyConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("inboundAllowed", m.GetInboundAllowed())
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
        err := writer.WriteBoolValue("outboundAllowed", m.GetOutboundAllowed())
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
func (m *InboundOutboundPolicyConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInboundAllowed sets the inboundAllowed property value. Defines whether external users coming inbound are allowed.
func (m *InboundOutboundPolicyConfiguration) SetInboundAllowed(value *bool)() {
    m.inboundAllowed = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InboundOutboundPolicyConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOutboundAllowed sets the outboundAllowed property value. Defines whether internal users are allowed to go outbound.
func (m *InboundOutboundPolicyConfiguration) SetOutboundAllowed(value *bool)() {
    m.outboundAllowed = value
}
type InboundOutboundPolicyConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInboundAllowed()(*bool)
    GetOdataType()(*string)
    GetOutboundAllowed()(*bool)
    SetInboundAllowed(value *bool)()
    SetOdataType(value *string)()
    SetOutboundAllowed(value *bool)()
}
