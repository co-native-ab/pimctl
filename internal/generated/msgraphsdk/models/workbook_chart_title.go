package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookChartTitle struct {
    Entity
    // The formatting of a chart title, which includes fill and font formatting. Read-only.
    format WorkbookChartTitleFormatable
    // Indicates whether the chart title will overlay the chart or not.
    overlay *bool
    // The title text of the chart.
    text *string
    // Indicates whether the chart title is visible.
    visible *bool
}
// NewWorkbookChartTitle instantiates a new WorkbookChartTitle and sets the default values.
func NewWorkbookChartTitle()(*WorkbookChartTitle) {
    m := &WorkbookChartTitle{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartTitleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookChartTitleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartTitle(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookChartTitle) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartTitleFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartTitleFormatable))
        }
        return nil
    }
    res["overlay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverlay(val)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    res["visible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisible(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. The formatting of a chart title, which includes fill and font formatting. Read-only.
// returns a WorkbookChartTitleFormatable when successful
func (m *WorkbookChartTitle) GetFormat()(WorkbookChartTitleFormatable) {
    return m.format
}
// GetOverlay gets the overlay property value. Indicates whether the chart title will overlay the chart or not.
// returns a *bool when successful
func (m *WorkbookChartTitle) GetOverlay()(*bool) {
    return m.overlay
}
// GetText gets the text property value. The title text of the chart.
// returns a *string when successful
func (m *WorkbookChartTitle) GetText()(*string) {
    return m.text
}
// GetVisible gets the visible property value. Indicates whether the chart title is visible.
// returns a *bool when successful
func (m *WorkbookChartTitle) GetVisible()(*bool) {
    return m.visible
}
// Serialize serializes information the current object
func (m *WorkbookChartTitle) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("overlay", m.GetOverlay())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("visible", m.GetVisible())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. The formatting of a chart title, which includes fill and font formatting. Read-only.
func (m *WorkbookChartTitle) SetFormat(value WorkbookChartTitleFormatable)() {
    m.format = value
}
// SetOverlay sets the overlay property value. Indicates whether the chart title will overlay the chart or not.
func (m *WorkbookChartTitle) SetOverlay(value *bool)() {
    m.overlay = value
}
// SetText sets the text property value. The title text of the chart.
func (m *WorkbookChartTitle) SetText(value *string)() {
    m.text = value
}
// SetVisible sets the visible property value. Indicates whether the chart title is visible.
func (m *WorkbookChartTitle) SetVisible(value *bool)() {
    m.visible = value
}
type WorkbookChartTitleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(WorkbookChartTitleFormatable)
    GetOverlay()(*bool)
    GetText()(*string)
    GetVisible()(*bool)
    SetFormat(value WorkbookChartTitleFormatable)()
    SetOverlay(value *bool)()
    SetText(value *string)()
    SetVisible(value *bool)()
}
