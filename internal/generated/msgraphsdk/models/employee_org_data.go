package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmployeeOrgData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cost center associated with the user. Returned only on $select. Supports $filter.
    costCenter *string
    // The name of the division in which the user works. Returned only on $select. Supports $filter.
    division *string
    // The OdataType property
    odataType *string
}
// NewEmployeeOrgData instantiates a new EmployeeOrgData and sets the default values.
func NewEmployeeOrgData()(*EmployeeOrgData) {
    m := &EmployeeOrgData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEmployeeOrgDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmployeeOrgDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmployeeOrgData(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmployeeOrgData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCostCenter gets the costCenter property value. The cost center associated with the user. Returned only on $select. Supports $filter.
// returns a *string when successful
func (m *EmployeeOrgData) GetCostCenter()(*string) {
    return m.costCenter
}
// GetDivision gets the division property value. The name of the division in which the user works. Returned only on $select. Supports $filter.
// returns a *string when successful
func (m *EmployeeOrgData) GetDivision()(*string) {
    return m.division
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmployeeOrgData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["costCenter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCostCenter(val)
        }
        return nil
    }
    res["division"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDivision(val)
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
func (m *EmployeeOrgData) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *EmployeeOrgData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("costCenter", m.GetCostCenter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("division", m.GetDivision())
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
func (m *EmployeeOrgData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCostCenter sets the costCenter property value. The cost center associated with the user. Returned only on $select. Supports $filter.
func (m *EmployeeOrgData) SetCostCenter(value *string)() {
    m.costCenter = value
}
// SetDivision sets the division property value. The name of the division in which the user works. Returned only on $select. Supports $filter.
func (m *EmployeeOrgData) SetDivision(value *string)() {
    m.division = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EmployeeOrgData) SetOdataType(value *string)() {
    m.odataType = value
}
type EmployeeOrgDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCostCenter()(*string)
    GetDivision()(*string)
    GetOdataType()(*string)
    SetCostCenter(value *string)()
    SetDivision(value *string)()
    SetOdataType(value *string)()
}
