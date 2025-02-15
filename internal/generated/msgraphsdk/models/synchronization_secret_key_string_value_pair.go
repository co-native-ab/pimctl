package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SynchronizationSecretKeyStringValuePair struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The key property
    key *SynchronizationSecret
    // The OdataType property
    odataType *string
    // The value of the secret.
    value *string
}
// NewSynchronizationSecretKeyStringValuePair instantiates a new SynchronizationSecretKeyStringValuePair and sets the default values.
func NewSynchronizationSecretKeyStringValuePair()(*SynchronizationSecretKeyStringValuePair) {
    m := &SynchronizationSecretKeyStringValuePair{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationSecretKeyStringValuePair(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SynchronizationSecretKeyStringValuePair) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SynchronizationSecretKeyStringValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationSecret)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(*SynchronizationSecret))
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
// GetKey gets the key property value. The key property
// returns a *SynchronizationSecret when successful
func (m *SynchronizationSecretKeyStringValuePair) GetKey()(*SynchronizationSecret) {
    return m.key
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SynchronizationSecretKeyStringValuePair) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The value of the secret.
// returns a *string when successful
func (m *SynchronizationSecretKeyStringValuePair) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *SynchronizationSecretKeyStringValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKey() != nil {
        cast := (*m.GetKey()).String()
        err := writer.WriteStringValue("key", &cast)
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
func (m *SynchronizationSecretKeyStringValuePair) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. The key property
func (m *SynchronizationSecretKeyStringValuePair) SetKey(value *SynchronizationSecret)() {
    m.key = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationSecretKeyStringValuePair) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The value of the secret.
func (m *SynchronizationSecretKeyStringValuePair) SetValue(value *string)() {
    m.value = value
}
type SynchronizationSecretKeyStringValuePairable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*SynchronizationSecret)
    GetOdataType()(*string)
    GetValue()(*string)
    SetKey(value *SynchronizationSecret)()
    SetOdataType(value *string)()
    SetValue(value *string)()
}
