package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FormattedContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content of this formattedContent.
    content *string
    // The format of the content. The possible values are: text, html, markdown, unknownFutureValue.
    format *ContentFormat
    // The OdataType property
    odataType *string
}
// NewFormattedContent instantiates a new FormattedContent and sets the default values.
func NewFormattedContent()(*FormattedContent) {
    m := &FormattedContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFormattedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFormattedContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFormattedContent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FormattedContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content of this formattedContent.
// returns a *string when successful
func (m *FormattedContent) GetContent()(*string) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FormattedContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*ContentFormat))
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
// GetFormat gets the format property value. The format of the content. The possible values are: text, html, markdown, unknownFutureValue.
// returns a *ContentFormat when successful
func (m *FormattedContent) GetFormat()(*ContentFormat) {
    return m.format
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *FormattedContent) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *FormattedContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    if m.GetFormat() != nil {
        cast := (*m.GetFormat()).String()
        err := writer.WriteStringValue("format", &cast)
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
func (m *FormattedContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content of this formattedContent.
func (m *FormattedContent) SetContent(value *string)() {
    m.content = value
}
// SetFormat sets the format property value. The format of the content. The possible values are: text, html, markdown, unknownFutureValue.
func (m *FormattedContent) SetFormat(value *ContentFormat)() {
    m.format = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FormattedContent) SetOdataType(value *string)() {
    m.odataType = value
}
type FormattedContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetFormat()(*ContentFormat)
    GetOdataType()(*string)
    SetContent(value *string)()
    SetFormat(value *ContentFormat)()
    SetOdataType(value *string)()
}
