package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LocalizedDescription struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description in the localized language.
    description *string
    // The language tag for the label.
    languageTag *string
    // The OdataType property
    odataType *string
}
// NewLocalizedDescription instantiates a new LocalizedDescription and sets the default values.
func NewLocalizedDescription()(*LocalizedDescription) {
    m := &LocalizedDescription{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLocalizedDescriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalizedDescriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocalizedDescription(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LocalizedDescription) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description in the localized language.
// returns a *string when successful
func (m *LocalizedDescription) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalizedDescription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
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
// GetLanguageTag gets the languageTag property value. The language tag for the label.
// returns a *string when successful
func (m *LocalizedDescription) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *LocalizedDescription) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *LocalizedDescription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
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
func (m *LocalizedDescription) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description in the localized language.
func (m *LocalizedDescription) SetDescription(value *string)() {
    m.description = value
}
// SetLanguageTag sets the languageTag property value. The language tag for the label.
func (m *LocalizedDescription) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LocalizedDescription) SetOdataType(value *string)() {
    m.odataType = value
}
type LocalizedDescriptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetLanguageTag()(*string)
    GetOdataType()(*string)
    SetDescription(value *string)()
    SetLanguageTag(value *string)()
    SetOdataType(value *string)()
}
