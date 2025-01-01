package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookFilterCriteria struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The color applied to the cell.
    color *string
    // A custom criterion.
    criterion1 *string
    // A custom criterion.
    criterion2 *string
    // A dynamic formula specified in a custom filter.
    dynamicCriteria *string
    // Indicates whether a filter is applied to a column.
    filterOn *string
    // An icon applied to a cell via conditional formatting.
    icon WorkbookIconable
    // The OdataType property
    odataType *string
    // An operator in a cell; for example, =, >, <, <=, or <>.
    operator *string
    // The values that appear in the cell.
    values i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
}
// NewWorkbookFilterCriteria instantiates a new WorkbookFilterCriteria and sets the default values.
func NewWorkbookFilterCriteria()(*WorkbookFilterCriteria) {
    m := &WorkbookFilterCriteria{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkbookFilterCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookFilterCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookFilterCriteria(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkbookFilterCriteria) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The color applied to the cell.
// returns a *string when successful
func (m *WorkbookFilterCriteria) GetColor()(*string) {
    return m.color
}
// GetCriterion1 gets the criterion1 property value. A custom criterion.
// returns a *string when successful
func (m *WorkbookFilterCriteria) GetCriterion1()(*string) {
    return m.criterion1
}
// GetCriterion2 gets the criterion2 property value. A custom criterion.
// returns a *string when successful
func (m *WorkbookFilterCriteria) GetCriterion2()(*string) {
    return m.criterion2
}
// GetDynamicCriteria gets the dynamicCriteria property value. A dynamic formula specified in a custom filter.
// returns a *string when successful
func (m *WorkbookFilterCriteria) GetDynamicCriteria()(*string) {
    return m.dynamicCriteria
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookFilterCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["criterion1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriterion1(val)
        }
        return nil
    }
    res["criterion2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriterion2(val)
        }
        return nil
    }
    res["dynamicCriteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDynamicCriteria(val)
        }
        return nil
    }
    res["filterOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilterOn(val)
        }
        return nil
    }
    res["icon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookIconFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIcon(val.(WorkbookIconable))
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
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    return res
}
// GetFilterOn gets the filterOn property value. Indicates whether a filter is applied to a column.
// returns a *string when successful
func (m *WorkbookFilterCriteria) GetFilterOn()(*string) {
    return m.filterOn
}
// GetIcon gets the icon property value. An icon applied to a cell via conditional formatting.
// returns a WorkbookIconable when successful
func (m *WorkbookFilterCriteria) GetIcon()(WorkbookIconable) {
    return m.icon
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WorkbookFilterCriteria) GetOdataType()(*string) {
    return m.odataType
}
// GetOperator gets the operator property value. An operator in a cell; for example, =, >, <, <=, or <>.
// returns a *string when successful
func (m *WorkbookFilterCriteria) GetOperator()(*string) {
    return m.operator
}
// GetValues gets the values property value. The values that appear in the cell.
// returns a UntypedNodeable when successful
func (m *WorkbookFilterCriteria) GetValues()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.values
}
// Serialize serializes information the current object
func (m *WorkbookFilterCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criterion1", m.GetCriterion1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criterion2", m.GetCriterion2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dynamicCriteria", m.GetDynamicCriteria())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filterOn", m.GetFilterOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("icon", m.GetIcon())
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
        err := writer.WriteStringValue("operator", m.GetOperator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("values", m.GetValues())
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
func (m *WorkbookFilterCriteria) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The color applied to the cell.
func (m *WorkbookFilterCriteria) SetColor(value *string)() {
    m.color = value
}
// SetCriterion1 sets the criterion1 property value. A custom criterion.
func (m *WorkbookFilterCriteria) SetCriterion1(value *string)() {
    m.criterion1 = value
}
// SetCriterion2 sets the criterion2 property value. A custom criterion.
func (m *WorkbookFilterCriteria) SetCriterion2(value *string)() {
    m.criterion2 = value
}
// SetDynamicCriteria sets the dynamicCriteria property value. A dynamic formula specified in a custom filter.
func (m *WorkbookFilterCriteria) SetDynamicCriteria(value *string)() {
    m.dynamicCriteria = value
}
// SetFilterOn sets the filterOn property value. Indicates whether a filter is applied to a column.
func (m *WorkbookFilterCriteria) SetFilterOn(value *string)() {
    m.filterOn = value
}
// SetIcon sets the icon property value. An icon applied to a cell via conditional formatting.
func (m *WorkbookFilterCriteria) SetIcon(value WorkbookIconable)() {
    m.icon = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkbookFilterCriteria) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperator sets the operator property value. An operator in a cell; for example, =, >, <, <=, or <>.
func (m *WorkbookFilterCriteria) SetOperator(value *string)() {
    m.operator = value
}
// SetValues sets the values property value. The values that appear in the cell.
func (m *WorkbookFilterCriteria) SetValues(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.values = value
}
type WorkbookFilterCriteriaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetCriterion1()(*string)
    GetCriterion2()(*string)
    GetDynamicCriteria()(*string)
    GetFilterOn()(*string)
    GetIcon()(WorkbookIconable)
    GetOdataType()(*string)
    GetOperator()(*string)
    GetValues()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    SetColor(value *string)()
    SetCriterion1(value *string)()
    SetCriterion2(value *string)()
    SetDynamicCriteria(value *string)()
    SetFilterOn(value *string)()
    SetIcon(value WorkbookIconable)()
    SetOdataType(value *string)()
    SetOperator(value *string)()
    SetValues(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
}
