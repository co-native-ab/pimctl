package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomExtensionCalloutInstance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identification of the custom extension that was triggered at this instance.
    customExtensionId *string
    // Details provided by the logic app during the callback of the request instance.
    detail *string
    // The unique run identifier for the logic app.
    externalCorrelationId *string
    // Unique identifier for the callout instance. Read-only.
    id *string
    // The OdataType property
    odataType *string
    // The status of the request to the custom extension. The possible values are: calloutSent, callbackReceived, calloutFailed, callbackTimedOut, waitingForCallback, unknownFutureValue.
    status *CustomExtensionCalloutInstanceStatus
}
// NewCustomExtensionCalloutInstance instantiates a new CustomExtensionCalloutInstance and sets the default values.
func NewCustomExtensionCalloutInstance()(*CustomExtensionCalloutInstance) {
    m := &CustomExtensionCalloutInstance{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomExtensionCalloutInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomExtensionCalloutInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionCalloutInstance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomExtensionCalloutInstance) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomExtensionId gets the customExtensionId property value. Identification of the custom extension that was triggered at this instance.
// returns a *string when successful
func (m *CustomExtensionCalloutInstance) GetCustomExtensionId()(*string) {
    return m.customExtensionId
}
// GetDetail gets the detail property value. Details provided by the logic app during the callback of the request instance.
// returns a *string when successful
func (m *CustomExtensionCalloutInstance) GetDetail()(*string) {
    return m.detail
}
// GetExternalCorrelationId gets the externalCorrelationId property value. The unique run identifier for the logic app.
// returns a *string when successful
func (m *CustomExtensionCalloutInstance) GetExternalCorrelationId()(*string) {
    return m.externalCorrelationId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomExtensionCalloutInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["customExtensionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtensionId(val)
        }
        return nil
    }
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val)
        }
        return nil
    }
    res["externalCorrelationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalCorrelationId(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
        val, err := n.GetEnumValue(ParseCustomExtensionCalloutInstanceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CustomExtensionCalloutInstanceStatus))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique identifier for the callout instance. Read-only.
// returns a *string when successful
func (m *CustomExtensionCalloutInstance) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CustomExtensionCalloutInstance) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. The status of the request to the custom extension. The possible values are: calloutSent, callbackReceived, calloutFailed, callbackTimedOut, waitingForCallback, unknownFutureValue.
// returns a *CustomExtensionCalloutInstanceStatus when successful
func (m *CustomExtensionCalloutInstance) GetStatus()(*CustomExtensionCalloutInstanceStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CustomExtensionCalloutInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("customExtensionId", m.GetCustomExtensionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalCorrelationId", m.GetExternalCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *CustomExtensionCalloutInstance) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomExtensionId sets the customExtensionId property value. Identification of the custom extension that was triggered at this instance.
func (m *CustomExtensionCalloutInstance) SetCustomExtensionId(value *string)() {
    m.customExtensionId = value
}
// SetDetail sets the detail property value. Details provided by the logic app during the callback of the request instance.
func (m *CustomExtensionCalloutInstance) SetDetail(value *string)() {
    m.detail = value
}
// SetExternalCorrelationId sets the externalCorrelationId property value. The unique run identifier for the logic app.
func (m *CustomExtensionCalloutInstance) SetExternalCorrelationId(value *string)() {
    m.externalCorrelationId = value
}
// SetId sets the id property value. Unique identifier for the callout instance. Read-only.
func (m *CustomExtensionCalloutInstance) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomExtensionCalloutInstance) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. The status of the request to the custom extension. The possible values are: calloutSent, callbackReceived, calloutFailed, callbackTimedOut, waitingForCallback, unknownFutureValue.
func (m *CustomExtensionCalloutInstance) SetStatus(value *CustomExtensionCalloutInstanceStatus)() {
    m.status = value
}
type CustomExtensionCalloutInstanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomExtensionId()(*string)
    GetDetail()(*string)
    GetExternalCorrelationId()(*string)
    GetId()(*string)
    GetOdataType()(*string)
    GetStatus()(*CustomExtensionCalloutInstanceStatus)
    SetCustomExtensionId(value *string)()
    SetDetail(value *string)()
    SetExternalCorrelationId(value *string)()
    SetId(value *string)()
    SetOdataType(value *string)()
    SetStatus(value *CustomExtensionCalloutInstanceStatus)()
}
