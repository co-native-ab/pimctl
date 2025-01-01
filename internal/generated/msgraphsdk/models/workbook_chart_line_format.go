package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookChartLineFormat struct {
    Entity
    // The HTML color code that represents the color of lines in the chart.
    color *string
}
// NewWorkbookChartLineFormat instantiates a new WorkbookChartLineFormat and sets the default values.
func NewWorkbookChartLineFormat()(*WorkbookChartLineFormat) {
    m := &WorkbookChartLineFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartLineFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookChartLineFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartLineFormat(), nil
}
// GetColor gets the color property value. The HTML color code that represents the color of lines in the chart.
// returns a *string when successful
func (m *WorkbookChartLineFormat) GetColor()(*string) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookChartLineFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *WorkbookChartLineFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetColor sets the color property value. The HTML color code that represents the color of lines in the chart.
func (m *WorkbookChartLineFormat) SetColor(value *string)() {
    m.color = value
}
type WorkbookChartLineFormatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    SetColor(value *string)()
}
