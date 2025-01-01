package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AlternativeSecurityId struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // For internal use only.
    identityProvider *string
    // For internal use only.
    key []byte
    // The OdataType property
    odataType *string
    // For internal use only.
    typeEscaped *int32
}
// NewAlternativeSecurityId instantiates a new AlternativeSecurityId and sets the default values.
func NewAlternativeSecurityId()(*AlternativeSecurityId) {
    m := &AlternativeSecurityId{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAlternativeSecurityIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAlternativeSecurityIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlternativeSecurityId(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AlternativeSecurityId) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AlternativeSecurityId) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["identityProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityProvider(val)
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetIdentityProvider gets the identityProvider property value. For internal use only.
// returns a *string when successful
func (m *AlternativeSecurityId) GetIdentityProvider()(*string) {
    return m.identityProvider
}
// GetKey gets the key property value. For internal use only.
// returns a []byte when successful
func (m *AlternativeSecurityId) GetKey()([]byte) {
    return m.key
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AlternativeSecurityId) GetOdataType()(*string) {
    return m.odataType
}
// GetTypeEscaped gets the type property value. For internal use only.
// returns a *int32 when successful
func (m *AlternativeSecurityId) GetTypeEscaped()(*int32) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *AlternativeSecurityId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("identityProvider", m.GetIdentityProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("key", m.GetKey())
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
        err := writer.WriteInt32Value("type", m.GetTypeEscaped())
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
func (m *AlternativeSecurityId) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIdentityProvider sets the identityProvider property value. For internal use only.
func (m *AlternativeSecurityId) SetIdentityProvider(value *string)() {
    m.identityProvider = value
}
// SetKey sets the key property value. For internal use only.
func (m *AlternativeSecurityId) SetKey(value []byte)() {
    m.key = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlternativeSecurityId) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTypeEscaped sets the type property value. For internal use only.
func (m *AlternativeSecurityId) SetTypeEscaped(value *int32)() {
    m.typeEscaped = value
}
type AlternativeSecurityIdable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentityProvider()(*string)
    GetKey()([]byte)
    GetOdataType()(*string)
    GetTypeEscaped()(*int32)
    SetIdentityProvider(value *string)()
    SetKey(value []byte)()
    SetOdataType(value *string)()
    SetTypeEscaped(value *int32)()
}
