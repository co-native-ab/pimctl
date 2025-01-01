package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookPivotTable struct {
    Entity
    // The name of the pivot table.
    name *string
    // The worksheet that contains the current pivot table. Read-only.
    worksheet WorkbookWorksheetable
}
// NewWorkbookPivotTable instantiates a new WorkbookPivotTable and sets the default values.
func NewWorkbookPivotTable()(*WorkbookPivotTable) {
    m := &WorkbookPivotTable{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookPivotTableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookPivotTableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookPivotTable(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookPivotTable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
// GetName gets the name property value. The name of the pivot table.
// returns a *string when successful
func (m *WorkbookPivotTable) GetName()(*string) {
    return m.name
}
// GetWorksheet gets the worksheet property value. The worksheet that contains the current pivot table. Read-only.
// returns a WorkbookWorksheetable when successful
func (m *WorkbookPivotTable) GetWorksheet()(WorkbookWorksheetable) {
    return m.worksheet
}
// Serialize serializes information the current object
func (m *WorkbookPivotTable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
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
// SetName sets the name property value. The name of the pivot table.
func (m *WorkbookPivotTable) SetName(value *string)() {
    m.name = value
}
// SetWorksheet sets the worksheet property value. The worksheet that contains the current pivot table. Read-only.
func (m *WorkbookPivotTable) SetWorksheet(value WorkbookWorksheetable)() {
    m.worksheet = value
}
type WorkbookPivotTableable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetWorksheet()(WorkbookWorksheetable)
    SetName(value *string)()
    SetWorksheet(value WorkbookWorksheetable)()
}
