package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageAnswerChoice struct {
    // The actual value of the selected choice. This is typically a string value which is understandable by applications. Required.
    actualValue *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The text of the answer choice represented in a format for a specific locale.
    localizations []AccessPackageLocalizedTextable
    // The OdataType property
    odataType *string
    // The string to display for this answer; if an Accept-Language header is provided, and there is a matching localization in localizations, this string will be the matching localized string; otherwise, this string remains as the default non-localized string. Required.
    text *string
}
// NewAccessPackageAnswerChoice instantiates a new AccessPackageAnswerChoice and sets the default values.
func NewAccessPackageAnswerChoice()(*AccessPackageAnswerChoice) {
    m := &AccessPackageAnswerChoice{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessPackageAnswerChoiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageAnswerChoiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAnswerChoice(), nil
}
// GetActualValue gets the actualValue property value. The actual value of the selected choice. This is typically a string value which is understandable by applications. Required.
// returns a *string when successful
func (m *AccessPackageAnswerChoice) GetActualValue()(*string) {
    return m.actualValue
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccessPackageAnswerChoice) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageAnswerChoice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actualValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActualValue(val)
        }
        return nil
    }
    res["localizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageLocalizedTextFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageLocalizedTextable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageLocalizedTextable)
                }
            }
            m.SetLocalizations(res)
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
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetLocalizations gets the localizations property value. The text of the answer choice represented in a format for a specific locale.
// returns a []AccessPackageLocalizedTextable when successful
func (m *AccessPackageAnswerChoice) GetLocalizations()([]AccessPackageLocalizedTextable) {
    return m.localizations
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AccessPackageAnswerChoice) GetOdataType()(*string) {
    return m.odataType
}
// GetText gets the text property value. The string to display for this answer; if an Accept-Language header is provided, and there is a matching localization in localizations, this string will be the matching localized string; otherwise, this string remains as the default non-localized string. Required.
// returns a *string when successful
func (m *AccessPackageAnswerChoice) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *AccessPackageAnswerChoice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actualValue", m.GetActualValue())
        if err != nil {
            return err
        }
    }
    if m.GetLocalizations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("localizations", cast)
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
        err := writer.WriteStringValue("text", m.GetText())
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
// SetActualValue sets the actualValue property value. The actual value of the selected choice. This is typically a string value which is understandable by applications. Required.
func (m *AccessPackageAnswerChoice) SetActualValue(value *string)() {
    m.actualValue = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAnswerChoice) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocalizations sets the localizations property value. The text of the answer choice represented in a format for a specific locale.
func (m *AccessPackageAnswerChoice) SetLocalizations(value []AccessPackageLocalizedTextable)() {
    m.localizations = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageAnswerChoice) SetOdataType(value *string)() {
    m.odataType = value
}
// SetText sets the text property value. The string to display for this answer; if an Accept-Language header is provided, and there is a matching localization in localizations, this string will be the matching localized string; otherwise, this string remains as the default non-localized string. Required.
func (m *AccessPackageAnswerChoice) SetText(value *string)() {
    m.text = value
}
type AccessPackageAnswerChoiceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActualValue()(*string)
    GetLocalizations()([]AccessPackageLocalizedTextable)
    GetOdataType()(*string)
    GetText()(*string)
    SetActualValue(value *string)()
    SetLocalizations(value []AccessPackageLocalizedTextable)()
    SetOdataType(value *string)()
    SetText(value *string)()
}
