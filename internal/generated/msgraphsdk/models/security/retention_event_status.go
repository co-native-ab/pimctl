package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type RetentionEventStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error if the status isn't successful.
    error ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable
    // The OdataType property
    odataType *string
    // The status of the distribution. The possible values are: pending, error, success, notAvaliable.
    status *EventStatusType
}
// NewRetentionEventStatus instantiates a new RetentionEventStatus and sets the default values.
func NewRetentionEventStatus()(*RetentionEventStatus) {
    m := &RetentionEventStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRetentionEventStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRetentionEventStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionEventStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RetentionEventStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetError gets the error property value. The error if the status isn't successful.
// returns a PublicErrorable when successful
func (m *RetentionEventStatus) GetError()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RetentionEventStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable))
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEventStatusType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*EventStatusType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *RetentionEventStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. The status of the distribution. The possible values are: pending, error, success, notAvaliable.
// returns a *EventStatusType when successful
func (m *RetentionEventStatus) GetStatus()(*EventStatusType) {
    return m.status
}
// Serialize serializes information the current object
func (m *RetentionEventStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *RetentionEventStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetError sets the error property value. The error if the status isn't successful.
func (m *RetentionEventStatus) SetError(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable)() {
    m.error = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetentionEventStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. The status of the distribution. The possible values are: pending, error, success, notAvaliable.
func (m *RetentionEventStatus) SetStatus(value *EventStatusType)() {
    m.status = value
}
type RetentionEventStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable)
    GetOdataType()(*string)
    GetStatus()(*EventStatusType)
    SetError(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable)()
    SetOdataType(value *string)()
    SetStatus(value *EventStatusType)()
}
