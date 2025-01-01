package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SimulationEventsContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Actual percentage of users who fell for the simulated attack in an attack simulation and training campaign.
    compromisedRate *float64
    // List of simulation events in an attack simulation and training campaign.
    events []SimulationEventable
    // The OdataType property
    odataType *string
}
// NewSimulationEventsContent instantiates a new SimulationEventsContent and sets the default values.
func NewSimulationEventsContent()(*SimulationEventsContent) {
    m := &SimulationEventsContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSimulationEventsContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSimulationEventsContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimulationEventsContent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SimulationEventsContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompromisedRate gets the compromisedRate property value. Actual percentage of users who fell for the simulated attack in an attack simulation and training campaign.
// returns a *float64 when successful
func (m *SimulationEventsContent) GetCompromisedRate()(*float64) {
    return m.compromisedRate
}
// GetEvents gets the events property value. List of simulation events in an attack simulation and training campaign.
// returns a []SimulationEventable when successful
func (m *SimulationEventsContent) GetEvents()([]SimulationEventable) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SimulationEventsContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compromisedRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompromisedRate(val)
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimulationEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimulationEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SimulationEventable)
                }
            }
            m.SetEvents(res)
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
func (m *SimulationEventsContent) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SimulationEventsContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("compromisedRate", m.GetCompromisedRate())
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("events", cast)
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
func (m *SimulationEventsContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompromisedRate sets the compromisedRate property value. Actual percentage of users who fell for the simulated attack in an attack simulation and training campaign.
func (m *SimulationEventsContent) SetCompromisedRate(value *float64)() {
    m.compromisedRate = value
}
// SetEvents sets the events property value. List of simulation events in an attack simulation and training campaign.
func (m *SimulationEventsContent) SetEvents(value []SimulationEventable)() {
    m.events = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SimulationEventsContent) SetOdataType(value *string)() {
    m.odataType = value
}
type SimulationEventsContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompromisedRate()(*float64)
    GetEvents()([]SimulationEventable)
    GetOdataType()(*string)
    SetCompromisedRate(value *float64)()
    SetEvents(value []SimulationEventable)()
    SetOdataType(value *string)()
}
