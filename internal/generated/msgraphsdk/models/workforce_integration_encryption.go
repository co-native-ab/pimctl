package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkforceIntegrationEncryption struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Possible values are: sharedSecret, unknownFutureValue.
    protocol *WorkforceIntegrationEncryptionProtocol
    // Encryption shared secret.
    secret *string
}
// NewWorkforceIntegrationEncryption instantiates a new WorkforceIntegrationEncryption and sets the default values.
func NewWorkforceIntegrationEncryption()(*WorkforceIntegrationEncryption) {
    m := &WorkforceIntegrationEncryption{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkforceIntegrationEncryptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkforceIntegrationEncryptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkforceIntegrationEncryption(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkforceIntegrationEncryption) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkforceIntegrationEncryption) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkforceIntegrationEncryptionProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val.(*WorkforceIntegrationEncryptionProtocol))
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WorkforceIntegrationEncryption) GetOdataType()(*string) {
    return m.odataType
}
// GetProtocol gets the protocol property value. Possible values are: sharedSecret, unknownFutureValue.
// returns a *WorkforceIntegrationEncryptionProtocol when successful
func (m *WorkforceIntegrationEncryption) GetProtocol()(*WorkforceIntegrationEncryptionProtocol) {
    return m.protocol
}
// GetSecret gets the secret property value. Encryption shared secret.
// returns a *string when successful
func (m *WorkforceIntegrationEncryption) GetSecret()(*string) {
    return m.secret
}
// Serialize serializes information the current object
func (m *WorkforceIntegrationEncryption) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetProtocol() != nil {
        cast := (*m.GetProtocol()).String()
        err := writer.WriteStringValue("protocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret", m.GetSecret())
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
func (m *WorkforceIntegrationEncryption) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkforceIntegrationEncryption) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProtocol sets the protocol property value. Possible values are: sharedSecret, unknownFutureValue.
func (m *WorkforceIntegrationEncryption) SetProtocol(value *WorkforceIntegrationEncryptionProtocol)() {
    m.protocol = value
}
// SetSecret sets the secret property value. Encryption shared secret.
func (m *WorkforceIntegrationEncryption) SetSecret(value *string)() {
    m.secret = value
}
type WorkforceIntegrationEncryptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetProtocol()(*WorkforceIntegrationEncryptionProtocol)
    GetSecret()(*string)
    SetOdataType(value *string)()
    SetProtocol(value *WorkforceIntegrationEncryptionProtocol)()
    SetSecret(value *string)()
}
