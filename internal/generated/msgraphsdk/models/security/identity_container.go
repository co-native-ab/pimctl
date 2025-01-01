package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type IdentityContainer struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Represents potential issues identified by Microsoft Defender for Identity within a customer's Microsoft Defender for Identity configuration.
    healthIssues []HealthIssueable
    // Represents a customer's Microsoft Defender for Identity sensors.
    sensors []Sensorable
}
// NewIdentityContainer instantiates a new IdentityContainer and sets the default values.
func NewIdentityContainer()(*IdentityContainer) {
    m := &IdentityContainer{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateIdentityContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityContainer(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentityContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["healthIssues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHealthIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HealthIssueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HealthIssueable)
                }
            }
            m.SetHealthIssues(res)
        }
        return nil
    }
    res["sensors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Sensorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Sensorable)
                }
            }
            m.SetSensors(res)
        }
        return nil
    }
    return res
}
// GetHealthIssues gets the healthIssues property value. Represents potential issues identified by Microsoft Defender for Identity within a customer's Microsoft Defender for Identity configuration.
// returns a []HealthIssueable when successful
func (m *IdentityContainer) GetHealthIssues()([]HealthIssueable) {
    return m.healthIssues
}
// GetSensors gets the sensors property value. Represents a customer's Microsoft Defender for Identity sensors.
// returns a []Sensorable when successful
func (m *IdentityContainer) GetSensors()([]Sensorable) {
    return m.sensors
}
// Serialize serializes information the current object
func (m *IdentityContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetHealthIssues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHealthIssues()))
        for i, v := range m.GetHealthIssues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("healthIssues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSensors()))
        for i, v := range m.GetSensors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sensors", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHealthIssues sets the healthIssues property value. Represents potential issues identified by Microsoft Defender for Identity within a customer's Microsoft Defender for Identity configuration.
func (m *IdentityContainer) SetHealthIssues(value []HealthIssueable)() {
    m.healthIssues = value
}
// SetSensors sets the sensors property value. Represents a customer's Microsoft Defender for Identity sensors.
func (m *IdentityContainer) SetSensors(value []Sensorable)() {
    m.sensors = value
}
type IdentityContainerable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHealthIssues()([]HealthIssueable)
    GetSensors()([]Sensorable)
    SetHealthIssues(value []HealthIssueable)()
    SetSensors(value []Sensorable)()
}
