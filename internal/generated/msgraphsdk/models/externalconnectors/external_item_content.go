package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalItemContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The type property
    typeEscaped *ExternalItemContentType
    // The content for the externalItem. Required.
    value *string
}
// NewExternalItemContent instantiates a new ExternalItemContent and sets the default values.
func NewExternalItemContent()(*ExternalItemContent) {
    m := &ExternalItemContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalItemContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalItemContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalItemContent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalItemContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalItemContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalItemContentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*ExternalItemContentType))
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
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ExternalItemContent) GetOdataType()(*string) {
    return m.odataType
}
// GetTypeEscaped gets the type property value. The type property
// returns a *ExternalItemContentType when successful
func (m *ExternalItemContent) GetTypeEscaped()(*ExternalItemContentType) {
    return m.typeEscaped
}
// GetValue gets the value property value. The content for the externalItem. Required.
// returns a *string when successful
func (m *ExternalItemContent) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ExternalItemContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *ExternalItemContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExternalItemContent) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *ExternalItemContent) SetTypeEscaped(value *ExternalItemContentType)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. The content for the externalItem. Required.
func (m *ExternalItemContent) SetValue(value *string)() {
    m.value = value
}
type ExternalItemContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetTypeEscaped()(*ExternalItemContentType)
    GetValue()(*string)
    SetOdataType(value *string)()
    SetTypeEscaped(value *ExternalItemContentType)()
    SetValue(value *string)()
}
