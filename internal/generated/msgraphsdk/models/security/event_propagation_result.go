package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EventPropagationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the specific location in the workload associated with the event.
    location *string
    // The OdataType property
    odataType *string
    // The name of the workload associated with the event.
    serviceName *string
    // Indicates the status of the event creation request. The possible values are: none, inProcessing, failed, success, unknownFutureValue.
    status *EventPropagationStatus
    // Additional information about the status of the event creation request.
    statusInformation *string
}
// NewEventPropagationResult instantiates a new EventPropagationResult and sets the default values.
func NewEventPropagationResult()(*EventPropagationResult) {
    m := &EventPropagationResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEventPropagationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEventPropagationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEventPropagationResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EventPropagationResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EventPropagationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
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
    res["serviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEventPropagationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*EventPropagationStatus))
        }
        return nil
    }
    res["statusInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusInformation(val)
        }
        return nil
    }
    return res
}
// GetLocation gets the location property value. The name of the specific location in the workload associated with the event.
// returns a *string when successful
func (m *EventPropagationResult) GetLocation()(*string) {
    return m.location
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *EventPropagationResult) GetOdataType()(*string) {
    return m.odataType
}
// GetServiceName gets the serviceName property value. The name of the workload associated with the event.
// returns a *string when successful
func (m *EventPropagationResult) GetServiceName()(*string) {
    return m.serviceName
}
// GetStatus gets the status property value. Indicates the status of the event creation request. The possible values are: none, inProcessing, failed, success, unknownFutureValue.
// returns a *EventPropagationStatus when successful
func (m *EventPropagationResult) GetStatus()(*EventPropagationStatus) {
    return m.status
}
// GetStatusInformation gets the statusInformation property value. Additional information about the status of the event creation request.
// returns a *string when successful
func (m *EventPropagationResult) GetStatusInformation()(*string) {
    return m.statusInformation
}
// Serialize serializes information the current object
func (m *EventPropagationResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("location", m.GetLocation())
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
        err := writer.WriteStringValue("serviceName", m.GetServiceName())
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
        err := writer.WriteStringValue("statusInformation", m.GetStatusInformation())
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
func (m *EventPropagationResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocation sets the location property value. The name of the specific location in the workload associated with the event.
func (m *EventPropagationResult) SetLocation(value *string)() {
    m.location = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EventPropagationResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetServiceName sets the serviceName property value. The name of the workload associated with the event.
func (m *EventPropagationResult) SetServiceName(value *string)() {
    m.serviceName = value
}
// SetStatus sets the status property value. Indicates the status of the event creation request. The possible values are: none, inProcessing, failed, success, unknownFutureValue.
func (m *EventPropagationResult) SetStatus(value *EventPropagationStatus)() {
    m.status = value
}
// SetStatusInformation sets the statusInformation property value. Additional information about the status of the event creation request.
func (m *EventPropagationResult) SetStatusInformation(value *string)() {
    m.statusInformation = value
}
type EventPropagationResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocation()(*string)
    GetOdataType()(*string)
    GetServiceName()(*string)
    GetStatus()(*EventPropagationStatus)
    GetStatusInformation()(*string)
    SetLocation(value *string)()
    SetOdataType(value *string)()
    SetServiceName(value *string)()
    SetStatus(value *EventPropagationStatus)()
    SetStatusInformation(value *string)()
}
