package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PlannerPlanContainer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The identifier of the resource that contains the plan. Optional.
    containerId *string
    // The OdataType property
    odataType *string
    // The type of the resource that contains the plan. For supported types, see the previous table. Possible values are: group, unknownFutureValue, roster. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: roster. Optional.
    typeEscaped *PlannerContainerType
    // The full canonical URL of the container. Optional.
    url *string
}
// NewPlannerPlanContainer instantiates a new PlannerPlanContainer and sets the default values.
func NewPlannerPlanContainer()(*PlannerPlanContainer) {
    m := &PlannerPlanContainer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePlannerPlanContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlannerPlanContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanContainer(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PlannerPlanContainer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContainerId gets the containerId property value. The identifier of the resource that contains the plan. Optional.
// returns a *string when successful
func (m *PlannerPlanContainer) GetContainerId()(*string) {
    return m.containerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlannerPlanContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["containerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerId(val)
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
        val, err := n.GetEnumValue(ParsePlannerContainerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*PlannerContainerType))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *PlannerPlanContainer) GetOdataType()(*string) {
    return m.odataType
}
// GetTypeEscaped gets the type property value. The type of the resource that contains the plan. For supported types, see the previous table. Possible values are: group, unknownFutureValue, roster. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: roster. Optional.
// returns a *PlannerContainerType when successful
func (m *PlannerPlanContainer) GetTypeEscaped()(*PlannerContainerType) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The full canonical URL of the container. Optional.
// returns a *string when successful
func (m *PlannerPlanContainer) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PlannerPlanContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("containerId", m.GetContainerId())
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
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *PlannerPlanContainer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContainerId sets the containerId property value. The identifier of the resource that contains the plan. Optional.
func (m *PlannerPlanContainer) SetContainerId(value *string)() {
    m.containerId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerPlanContainer) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTypeEscaped sets the type property value. The type of the resource that contains the plan. For supported types, see the previous table. Possible values are: group, unknownFutureValue, roster. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: roster. Optional.
func (m *PlannerPlanContainer) SetTypeEscaped(value *PlannerContainerType)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The full canonical URL of the container. Optional.
func (m *PlannerPlanContainer) SetUrl(value *string)() {
    m.url = value
}
type PlannerPlanContainerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainerId()(*string)
    GetOdataType()(*string)
    GetTypeEscaped()(*PlannerContainerType)
    GetUrl()(*string)
    SetContainerId(value *string)()
    SetOdataType(value *string)()
    SetTypeEscaped(value *PlannerContainerType)()
    SetUrl(value *string)()
}
