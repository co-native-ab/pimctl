package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomExtensionClientConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The max number of retries that Microsoft Entra ID makes to the external API. Values of 0 or 1 are supported. If null, the default for the service applies.
    maximumRetries *int32
    // The OdataType property
    odataType *string
    // The max duration in milliseconds that Microsoft Entra ID waits for a response from the external app before it shuts down the connection. The valid range is between 200 and 2000 milliseconds. Default duration is 1000.
    timeoutInMilliseconds *int32
}
// NewCustomExtensionClientConfiguration instantiates a new CustomExtensionClientConfiguration and sets the default values.
func NewCustomExtensionClientConfiguration()(*CustomExtensionClientConfiguration) {
    m := &CustomExtensionClientConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomExtensionClientConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomExtensionClientConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionClientConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomExtensionClientConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomExtensionClientConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maximumRetries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumRetries(val)
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
    res["timeoutInMilliseconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeoutInMilliseconds(val)
        }
        return nil
    }
    return res
}
// GetMaximumRetries gets the maximumRetries property value. The max number of retries that Microsoft Entra ID makes to the external API. Values of 0 or 1 are supported. If null, the default for the service applies.
// returns a *int32 when successful
func (m *CustomExtensionClientConfiguration) GetMaximumRetries()(*int32) {
    return m.maximumRetries
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CustomExtensionClientConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetTimeoutInMilliseconds gets the timeoutInMilliseconds property value. The max duration in milliseconds that Microsoft Entra ID waits for a response from the external app before it shuts down the connection. The valid range is between 200 and 2000 milliseconds. Default duration is 1000.
// returns a *int32 when successful
func (m *CustomExtensionClientConfiguration) GetTimeoutInMilliseconds()(*int32) {
    return m.timeoutInMilliseconds
}
// Serialize serializes information the current object
func (m *CustomExtensionClientConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maximumRetries", m.GetMaximumRetries())
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
        err := writer.WriteInt32Value("timeoutInMilliseconds", m.GetTimeoutInMilliseconds())
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
func (m *CustomExtensionClientConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaximumRetries sets the maximumRetries property value. The max number of retries that Microsoft Entra ID makes to the external API. Values of 0 or 1 are supported. If null, the default for the service applies.
func (m *CustomExtensionClientConfiguration) SetMaximumRetries(value *int32)() {
    m.maximumRetries = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomExtensionClientConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTimeoutInMilliseconds sets the timeoutInMilliseconds property value. The max duration in milliseconds that Microsoft Entra ID waits for a response from the external app before it shuts down the connection. The valid range is between 200 and 2000 milliseconds. Default duration is 1000.
func (m *CustomExtensionClientConfiguration) SetTimeoutInMilliseconds(value *int32)() {
    m.timeoutInMilliseconds = value
}
type CustomExtensionClientConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumRetries()(*int32)
    GetOdataType()(*string)
    GetTimeoutInMilliseconds()(*int32)
    SetMaximumRetries(value *int32)()
    SetOdataType(value *string)()
    SetTimeoutInMilliseconds(value *int32)()
}
