package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AutonomousSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the autonomous system.
    name *string
    // The autonomous system number, assigned by IANA.
    number *int32
    // The OdataType property
    odataType *string
    // The name of the autonomous system organization.
    organization *string
    // A displayable value for these autonomous system details.
    value *string
}
// NewAutonomousSystem instantiates a new AutonomousSystem and sets the default values.
func NewAutonomousSystem()(*AutonomousSystem) {
    m := &AutonomousSystem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAutonomousSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAutonomousSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAutonomousSystem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AutonomousSystem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AutonomousSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
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
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val)
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
// GetName gets the name property value. The name of the autonomous system.
// returns a *string when successful
func (m *AutonomousSystem) GetName()(*string) {
    return m.name
}
// GetNumber gets the number property value. The autonomous system number, assigned by IANA.
// returns a *int32 when successful
func (m *AutonomousSystem) GetNumber()(*int32) {
    return m.number
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AutonomousSystem) GetOdataType()(*string) {
    return m.odataType
}
// GetOrganization gets the organization property value. The name of the autonomous system organization.
// returns a *string when successful
func (m *AutonomousSystem) GetOrganization()(*string) {
    return m.organization
}
// GetValue gets the value property value. A displayable value for these autonomous system details.
// returns a *string when successful
func (m *AutonomousSystem) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *AutonomousSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("number", m.GetNumber())
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
        err := writer.WriteStringValue("organization", m.GetOrganization())
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
func (m *AutonomousSystem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name of the autonomous system.
func (m *AutonomousSystem) SetName(value *string)() {
    m.name = value
}
// SetNumber sets the number property value. The autonomous system number, assigned by IANA.
func (m *AutonomousSystem) SetNumber(value *int32)() {
    m.number = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AutonomousSystem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrganization sets the organization property value. The name of the autonomous system organization.
func (m *AutonomousSystem) SetOrganization(value *string)() {
    m.organization = value
}
// SetValue sets the value property value. A displayable value for these autonomous system details.
func (m *AutonomousSystem) SetValue(value *string)() {
    m.value = value
}
type AutonomousSystemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetNumber()(*int32)
    GetOdataType()(*string)
    GetOrganization()(*string)
    GetValue()(*string)
    SetName(value *string)()
    SetNumber(value *int32)()
    SetOdataType(value *string)()
    SetOrganization(value *string)()
    SetValue(value *string)()
}
