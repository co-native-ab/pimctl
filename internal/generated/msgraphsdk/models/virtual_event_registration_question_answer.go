package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEventRegistrationQuestionAnswer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Boolean answer of the virtual event registration question. Only appears when answerInputType is boolean.
    booleanValue *bool
    // Display name of the registration question.
    displayName *string
    // Collection of text answer of the virtual event registration question. Only appears when answerInputType is multiChoice.
    multiChoiceValues []string
    // The OdataType property
    odataType *string
    // id of the virtual event registration question.
    questionId *string
    // Text answer of the virtual event registration question. Appears when answerInputType is text, multilineText or singleChoice.
    value *string
}
// NewVirtualEventRegistrationQuestionAnswer instantiates a new VirtualEventRegistrationQuestionAnswer and sets the default values.
func NewVirtualEventRegistrationQuestionAnswer()(*VirtualEventRegistrationQuestionAnswer) {
    m := &VirtualEventRegistrationQuestionAnswer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVirtualEventRegistrationQuestionAnswerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventRegistrationQuestionAnswerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventRegistrationQuestionAnswer(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VirtualEventRegistrationQuestionAnswer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBooleanValue gets the booleanValue property value. Boolean answer of the virtual event registration question. Only appears when answerInputType is boolean.
// returns a *bool when successful
func (m *VirtualEventRegistrationQuestionAnswer) GetBooleanValue()(*bool) {
    return m.booleanValue
}
// GetDisplayName gets the displayName property value. Display name of the registration question.
// returns a *string when successful
func (m *VirtualEventRegistrationQuestionAnswer) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEventRegistrationQuestionAnswer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["booleanValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBooleanValue(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["multiChoiceValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetMultiChoiceValues(res)
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
    res["questionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuestionId(val)
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
// GetMultiChoiceValues gets the multiChoiceValues property value. Collection of text answer of the virtual event registration question. Only appears when answerInputType is multiChoice.
// returns a []string when successful
func (m *VirtualEventRegistrationQuestionAnswer) GetMultiChoiceValues()([]string) {
    return m.multiChoiceValues
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *VirtualEventRegistrationQuestionAnswer) GetOdataType()(*string) {
    return m.odataType
}
// GetQuestionId gets the questionId property value. id of the virtual event registration question.
// returns a *string when successful
func (m *VirtualEventRegistrationQuestionAnswer) GetQuestionId()(*string) {
    return m.questionId
}
// GetValue gets the value property value. Text answer of the virtual event registration question. Appears when answerInputType is text, multilineText or singleChoice.
// returns a *string when successful
func (m *VirtualEventRegistrationQuestionAnswer) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *VirtualEventRegistrationQuestionAnswer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("booleanValue", m.GetBooleanValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetMultiChoiceValues() != nil {
        err := writer.WriteCollectionOfStringValues("multiChoiceValues", m.GetMultiChoiceValues())
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
        err := writer.WriteStringValue("questionId", m.GetQuestionId())
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
func (m *VirtualEventRegistrationQuestionAnswer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBooleanValue sets the booleanValue property value. Boolean answer of the virtual event registration question. Only appears when answerInputType is boolean.
func (m *VirtualEventRegistrationQuestionAnswer) SetBooleanValue(value *bool)() {
    m.booleanValue = value
}
// SetDisplayName sets the displayName property value. Display name of the registration question.
func (m *VirtualEventRegistrationQuestionAnswer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetMultiChoiceValues sets the multiChoiceValues property value. Collection of text answer of the virtual event registration question. Only appears when answerInputType is multiChoice.
func (m *VirtualEventRegistrationQuestionAnswer) SetMultiChoiceValues(value []string)() {
    m.multiChoiceValues = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VirtualEventRegistrationQuestionAnswer) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuestionId sets the questionId property value. id of the virtual event registration question.
func (m *VirtualEventRegistrationQuestionAnswer) SetQuestionId(value *string)() {
    m.questionId = value
}
// SetValue sets the value property value. Text answer of the virtual event registration question. Appears when answerInputType is text, multilineText or singleChoice.
func (m *VirtualEventRegistrationQuestionAnswer) SetValue(value *string)() {
    m.value = value
}
type VirtualEventRegistrationQuestionAnswerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBooleanValue()(*bool)
    GetDisplayName()(*string)
    GetMultiChoiceValues()([]string)
    GetOdataType()(*string)
    GetQuestionId()(*string)
    GetValue()(*string)
    SetBooleanValue(value *bool)()
    SetDisplayName(value *string)()
    SetMultiChoiceValues(value []string)()
    SetOdataType(value *string)()
    SetQuestionId(value *string)()
    SetValue(value *string)()
}
