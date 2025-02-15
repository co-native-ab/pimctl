package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BookingQuestionAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether it's mandatory to answer the custom question.
    isRequired *bool
    // The OdataType property
    odataType *string
    // The ID of the custom question.
    questionId *string
}
// NewBookingQuestionAssignment instantiates a new BookingQuestionAssignment and sets the default values.
func NewBookingQuestionAssignment()(*BookingQuestionAssignment) {
    m := &BookingQuestionAssignment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBookingQuestionAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBookingQuestionAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingQuestionAssignment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BookingQuestionAssignment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BookingQuestionAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
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
    return res
}
// GetIsRequired gets the isRequired property value. Indicates whether it's mandatory to answer the custom question.
// returns a *bool when successful
func (m *BookingQuestionAssignment) GetIsRequired()(*bool) {
    return m.isRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *BookingQuestionAssignment) GetOdataType()(*string) {
    return m.odataType
}
// GetQuestionId gets the questionId property value. The ID of the custom question.
// returns a *string when successful
func (m *BookingQuestionAssignment) GetQuestionId()(*string) {
    return m.questionId
}
// Serialize serializes information the current object
func (m *BookingQuestionAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingQuestionAssignment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsRequired sets the isRequired property value. Indicates whether it's mandatory to answer the custom question.
func (m *BookingQuestionAssignment) SetIsRequired(value *bool)() {
    m.isRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BookingQuestionAssignment) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuestionId sets the questionId property value. The ID of the custom question.
func (m *BookingQuestionAssignment) SetQuestionId(value *string)() {
    m.questionId = value
}
type BookingQuestionAssignmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsRequired()(*bool)
    GetOdataType()(*string)
    GetQuestionId()(*string)
    SetIsRequired(value *bool)()
    SetOdataType(value *string)()
    SetQuestionId(value *string)()
}
