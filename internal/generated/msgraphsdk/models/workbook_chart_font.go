package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookChartFont struct {
    Entity
    // Indicates whether the fond is bold.
    bold *bool
    // The HTML color code representation of the text color. For example #FF0000 represents Red.
    color *string
    // Indicates whether the fond is italic.
    italic *bool
    // The font name. For example 'Calibri'.
    name *string
    // The size of the font. For example,  11.
    size *float64
    // The type of underlining applied to the font. The possible values are: None, Single.
    underline *string
}
// NewWorkbookChartFont instantiates a new WorkbookChartFont and sets the default values.
func NewWorkbookChartFont()(*WorkbookChartFont) {
    m := &WorkbookChartFont{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartFontFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookChartFontFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartFont(), nil
}
// GetBold gets the bold property value. Indicates whether the fond is bold.
// returns a *bool when successful
func (m *WorkbookChartFont) GetBold()(*bool) {
    return m.bold
}
// GetColor gets the color property value. The HTML color code representation of the text color. For example #FF0000 represents Red.
// returns a *string when successful
func (m *WorkbookChartFont) GetColor()(*string) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookChartFont) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBold(val)
        }
        return nil
    }
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
    res["italic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItalic(val)
        }
        return nil
    }
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
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["underline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnderline(val)
        }
        return nil
    }
    return res
}
// GetItalic gets the italic property value. Indicates whether the fond is italic.
// returns a *bool when successful
func (m *WorkbookChartFont) GetItalic()(*bool) {
    return m.italic
}
// GetName gets the name property value. The font name. For example 'Calibri'.
// returns a *string when successful
func (m *WorkbookChartFont) GetName()(*string) {
    return m.name
}
// GetSize gets the size property value. The size of the font. For example,  11.
// returns a *float64 when successful
func (m *WorkbookChartFont) GetSize()(*float64) {
    return m.size
}
// GetUnderline gets the underline property value. The type of underlining applied to the font. The possible values are: None, Single.
// returns a *string when successful
func (m *WorkbookChartFont) GetUnderline()(*string) {
    return m.underline
}
// Serialize serializes information the current object
func (m *WorkbookChartFont) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("bold", m.GetBold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("italic", m.GetItalic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("underline", m.GetUnderline())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBold sets the bold property value. Indicates whether the fond is bold.
func (m *WorkbookChartFont) SetBold(value *bool)() {
    m.bold = value
}
// SetColor sets the color property value. The HTML color code representation of the text color. For example #FF0000 represents Red.
func (m *WorkbookChartFont) SetColor(value *string)() {
    m.color = value
}
// SetItalic sets the italic property value. Indicates whether the fond is italic.
func (m *WorkbookChartFont) SetItalic(value *bool)() {
    m.italic = value
}
// SetName sets the name property value. The font name. For example 'Calibri'.
func (m *WorkbookChartFont) SetName(value *string)() {
    m.name = value
}
// SetSize sets the size property value. The size of the font. For example,  11.
func (m *WorkbookChartFont) SetSize(value *float64)() {
    m.size = value
}
// SetUnderline sets the underline property value. The type of underlining applied to the font. The possible values are: None, Single.
func (m *WorkbookChartFont) SetUnderline(value *string)() {
    m.underline = value
}
type WorkbookChartFontable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBold()(*bool)
    GetColor()(*string)
    GetItalic()(*bool)
    GetName()(*string)
    GetSize()(*float64)
    GetUnderline()(*string)
    SetBold(value *bool)()
    SetColor(value *string)()
    SetItalic(value *bool)()
    SetName(value *string)()
    SetSize(value *float64)()
    SetUnderline(value *string)()
}
