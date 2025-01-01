package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomExtensionOverwriteConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Configuration regarding properties of the custom extension which can be overwritten per event listener. If no values are provided, the properties on the custom extension are used.
    clientConfiguration CustomExtensionClientConfigurationable
    // The OdataType property
    odataType *string
}
// NewCustomExtensionOverwriteConfiguration instantiates a new CustomExtensionOverwriteConfiguration and sets the default values.
func NewCustomExtensionOverwriteConfiguration()(*CustomExtensionOverwriteConfiguration) {
    m := &CustomExtensionOverwriteConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomExtensionOverwriteConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomExtensionOverwriteConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionOverwriteConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomExtensionOverwriteConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientConfiguration gets the clientConfiguration property value. Configuration regarding properties of the custom extension which can be overwritten per event listener. If no values are provided, the properties on the custom extension are used.
// returns a CustomExtensionClientConfigurationable when successful
func (m *CustomExtensionOverwriteConfiguration) GetClientConfiguration()(CustomExtensionClientConfigurationable) {
    return m.clientConfiguration
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomExtensionOverwriteConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomExtensionClientConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientConfiguration(val.(CustomExtensionClientConfigurationable))
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
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CustomExtensionOverwriteConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CustomExtensionOverwriteConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("clientConfiguration", m.GetClientConfiguration())
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
func (m *CustomExtensionOverwriteConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientConfiguration sets the clientConfiguration property value. Configuration regarding properties of the custom extension which can be overwritten per event listener. If no values are provided, the properties on the custom extension are used.
func (m *CustomExtensionOverwriteConfiguration) SetClientConfiguration(value CustomExtensionClientConfigurationable)() {
    m.clientConfiguration = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomExtensionOverwriteConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
type CustomExtensionOverwriteConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientConfiguration()(CustomExtensionClientConfigurationable)
    GetOdataType()(*string)
    SetClientConfiguration(value CustomExtensionClientConfigurationable)()
    SetOdataType(value *string)()
}
