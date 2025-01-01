package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileHash struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The algorithm property
    algorithm *FileHashAlgorithm
    // The OdataType property
    odataType *string
    // The hash value.
    value *string
}
// NewFileHash instantiates a new FileHash and sets the default values.
func NewFileHash()(*FileHash) {
    m := &FileHash{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileHashFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileHashFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileHash(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FileHash) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlgorithm gets the algorithm property value. The algorithm property
// returns a *FileHashAlgorithm when successful
func (m *FileHash) GetAlgorithm()(*FileHashAlgorithm) {
    return m.algorithm
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileHash) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["algorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileHashAlgorithm)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlgorithm(val.(*FileHashAlgorithm))
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
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *FileHash) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The hash value.
// returns a *string when successful
func (m *FileHash) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *FileHash) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAlgorithm() != nil {
        cast := (*m.GetAlgorithm()).String()
        err := writer.WriteStringValue("algorithm", &cast)
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
func (m *FileHash) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlgorithm sets the algorithm property value. The algorithm property
func (m *FileHash) SetAlgorithm(value *FileHashAlgorithm)() {
    m.algorithm = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FileHash) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The hash value.
func (m *FileHash) SetValue(value *string)() {
    m.value = value
}
type FileHashable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlgorithm()(*FileHashAlgorithm)
    GetOdataType()(*string)
    GetValue()(*string)
    SetAlgorithm(value *FileHashAlgorithm)()
    SetOdataType(value *string)()
    SetValue(value *string)()
}
