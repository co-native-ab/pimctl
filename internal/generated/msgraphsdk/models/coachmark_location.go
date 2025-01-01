package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CoachmarkLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Length of coachmark.
    length *int32
    // The OdataType property
    odataType *string
    // Offset of coachmark.
    offset *int32
    // Type of coachmark location. The possible values are: unknown, fromEmail, subject, externalTag, displayName, messageBody, unknownFutureValue.
    typeEscaped *CoachmarkLocationType
}
// NewCoachmarkLocation instantiates a new CoachmarkLocation and sets the default values.
func NewCoachmarkLocation()(*CoachmarkLocation) {
    m := &CoachmarkLocation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCoachmarkLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCoachmarkLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCoachmarkLocation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CoachmarkLocation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CoachmarkLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLength(val)
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
    res["offset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffset(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCoachmarkLocationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*CoachmarkLocationType))
        }
        return nil
    }
    return res
}
// GetLength gets the length property value. Length of coachmark.
// returns a *int32 when successful
func (m *CoachmarkLocation) GetLength()(*int32) {
    return m.length
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CoachmarkLocation) GetOdataType()(*string) {
    return m.odataType
}
// GetOffset gets the offset property value. Offset of coachmark.
// returns a *int32 when successful
func (m *CoachmarkLocation) GetOffset()(*int32) {
    return m.offset
}
// GetTypeEscaped gets the type property value. Type of coachmark location. The possible values are: unknown, fromEmail, subject, externalTag, displayName, messageBody, unknownFutureValue.
// returns a *CoachmarkLocationType when successful
func (m *CoachmarkLocation) GetTypeEscaped()(*CoachmarkLocationType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *CoachmarkLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("length", m.GetLength())
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
        err := writer.WriteInt32Value("offset", m.GetOffset())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CoachmarkLocation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLength sets the length property value. Length of coachmark.
func (m *CoachmarkLocation) SetLength(value *int32)() {
    m.length = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CoachmarkLocation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOffset sets the offset property value. Offset of coachmark.
func (m *CoachmarkLocation) SetOffset(value *int32)() {
    m.offset = value
}
// SetTypeEscaped sets the type property value. Type of coachmark location. The possible values are: unknown, fromEmail, subject, externalTag, displayName, messageBody, unknownFutureValue.
func (m *CoachmarkLocation) SetTypeEscaped(value *CoachmarkLocationType)() {
    m.typeEscaped = value
}
type CoachmarkLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLength()(*int32)
    GetOdataType()(*string)
    GetOffset()(*int32)
    GetTypeEscaped()(*CoachmarkLocationType)
    SetLength(value *int32)()
    SetOdataType(value *string)()
    SetOffset(value *int32)()
    SetTypeEscaped(value *CoachmarkLocationType)()
}
