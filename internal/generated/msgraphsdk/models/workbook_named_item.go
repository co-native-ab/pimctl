package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookNamedItem struct {
    Entity
    // The comment associated with this name.
    comment *string
    // The name of the object. Read-only.
    name *string
    // Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only.
    scope *string
    // The type of reference is associated with the name. Possible values are: String, Integer, Double, Boolean, Range. Read-only.
    typeEscaped *string
    // The formula that the name is defined to refer to. For example, =Sheet14!$B$2:$H$12 and =4.75. Read-only.
    value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Indicates whether the object is visible.
    visible *bool
    // Returns the worksheet to which the named item is scoped. Available only if the item is scoped to the worksheet. Read-only.
    worksheet WorkbookWorksheetable
}
// NewWorkbookNamedItem instantiates a new WorkbookNamedItem and sets the default values.
func NewWorkbookNamedItem()(*WorkbookNamedItem) {
    m := &WorkbookNamedItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookNamedItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookNamedItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookNamedItem(), nil
}
// GetComment gets the comment property value. The comment associated with this name.
// returns a *string when successful
func (m *WorkbookNamedItem) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookNamedItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
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
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
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
    res["worksheet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookWorksheetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorksheet(val.(WorkbookWorksheetable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the object. Read-only.
// returns a *string when successful
func (m *WorkbookNamedItem) GetName()(*string) {
    return m.name
}
// GetScope gets the scope property value. Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only.
// returns a *string when successful
func (m *WorkbookNamedItem) GetScope()(*string) {
    return m.scope
}
// GetTypeEscaped gets the type property value. The type of reference is associated with the name. Possible values are: String, Integer, Double, Boolean, Range. Read-only.
// returns a *string when successful
func (m *WorkbookNamedItem) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetValue gets the value property value. The formula that the name is defined to refer to. For example, =Sheet14!$B$2:$H$12 and =4.75. Read-only.
// returns a UntypedNodeable when successful
func (m *WorkbookNamedItem) GetValue()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.value
}
// GetVisible gets the visible property value. Indicates whether the object is visible.
// returns a *bool when successful
func (m *WorkbookNamedItem) GetVisible()(*bool) {
    return m.visible
}
// GetWorksheet gets the worksheet property value. Returns the worksheet to which the named item is scoped. Available only if the item is scoped to the worksheet. Read-only.
// returns a WorkbookWorksheetable when successful
func (m *WorkbookNamedItem) GetWorksheet()(WorkbookWorksheetable) {
    return m.worksheet
}
// Serialize serializes information the current object
func (m *WorkbookNamedItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
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
        err = writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
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
    {
        err = writer.WriteObjectValue("worksheet", m.GetWorksheet())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComment sets the comment property value. The comment associated with this name.
func (m *WorkbookNamedItem) SetComment(value *string)() {
    m.comment = value
}
// SetName sets the name property value. The name of the object. Read-only.
func (m *WorkbookNamedItem) SetName(value *string)() {
    m.name = value
}
// SetScope sets the scope property value. Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only.
func (m *WorkbookNamedItem) SetScope(value *string)() {
    m.scope = value
}
// SetTypeEscaped sets the type property value. The type of reference is associated with the name. Possible values are: String, Integer, Double, Boolean, Range. Read-only.
func (m *WorkbookNamedItem) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. The formula that the name is defined to refer to. For example, =Sheet14!$B$2:$H$12 and =4.75. Read-only.
func (m *WorkbookNamedItem) SetValue(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.value = value
}
// SetVisible sets the visible property value. Indicates whether the object is visible.
func (m *WorkbookNamedItem) SetVisible(value *bool)() {
    m.visible = value
}
// SetWorksheet sets the worksheet property value. Returns the worksheet to which the named item is scoped. Available only if the item is scoped to the worksheet. Read-only.
func (m *WorkbookNamedItem) SetWorksheet(value WorkbookWorksheetable)() {
    m.worksheet = value
}
type WorkbookNamedItemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetName()(*string)
    GetScope()(*string)
    GetTypeEscaped()(*string)
    GetValue()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetVisible()(*bool)
    GetWorksheet()(WorkbookWorksheetable)
    SetComment(value *string)()
    SetName(value *string)()
    SetScope(value *string)()
    SetTypeEscaped(value *string)()
    SetValue(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetVisible(value *bool)()
    SetWorksheet(value WorkbookWorksheetable)()
}
