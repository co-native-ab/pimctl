package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookWorksheetProtectionOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether the worksheet protection option to allow the use of the autofilter feature is enabled.
    allowAutoFilter *bool
    // Indicates whether the worksheet protection option to allow deleting columns is enabled.
    allowDeleteColumns *bool
    // Indicates whether the worksheet protection option to allow deleting rows is enabled.
    allowDeleteRows *bool
    // Indicates whether the worksheet protection option to allow formatting cells is enabled.
    allowFormatCells *bool
    // Indicates whether the worksheet protection option to allow formatting columns is enabled.
    allowFormatColumns *bool
    // Indicates whether the worksheet protection option to allow formatting rows is enabled.
    allowFormatRows *bool
    // Indicates whether the worksheet protection option to allow inserting columns is enabled.
    allowInsertColumns *bool
    // Indicates whether the worksheet protection option to allow inserting hyperlinks is enabled.
    allowInsertHyperlinks *bool
    // Indicates whether the worksheet protection option to allow inserting rows is enabled.
    allowInsertRows *bool
    // Indicates whether the worksheet protection option to allow the use of the pivot table feature is enabled.
    allowPivotTables *bool
    // Indicates whether the worksheet protection option to allow the use of the sort feature is enabled.
    allowSort *bool
    // The OdataType property
    odataType *string
}
// NewWorkbookWorksheetProtectionOptions instantiates a new WorkbookWorksheetProtectionOptions and sets the default values.
func NewWorkbookWorksheetProtectionOptions()(*WorkbookWorksheetProtectionOptions) {
    m := &WorkbookWorksheetProtectionOptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkbookWorksheetProtectionOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookWorksheetProtectionOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookWorksheetProtectionOptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkbookWorksheetProtectionOptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowAutoFilter gets the allowAutoFilter property value. Indicates whether the worksheet protection option to allow the use of the autofilter feature is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowAutoFilter()(*bool) {
    return m.allowAutoFilter
}
// GetAllowDeleteColumns gets the allowDeleteColumns property value. Indicates whether the worksheet protection option to allow deleting columns is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowDeleteColumns()(*bool) {
    return m.allowDeleteColumns
}
// GetAllowDeleteRows gets the allowDeleteRows property value. Indicates whether the worksheet protection option to allow deleting rows is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowDeleteRows()(*bool) {
    return m.allowDeleteRows
}
// GetAllowFormatCells gets the allowFormatCells property value. Indicates whether the worksheet protection option to allow formatting cells is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatCells()(*bool) {
    return m.allowFormatCells
}
// GetAllowFormatColumns gets the allowFormatColumns property value. Indicates whether the worksheet protection option to allow formatting columns is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatColumns()(*bool) {
    return m.allowFormatColumns
}
// GetAllowFormatRows gets the allowFormatRows property value. Indicates whether the worksheet protection option to allow formatting rows is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatRows()(*bool) {
    return m.allowFormatRows
}
// GetAllowInsertColumns gets the allowInsertColumns property value. Indicates whether the worksheet protection option to allow inserting columns is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertColumns()(*bool) {
    return m.allowInsertColumns
}
// GetAllowInsertHyperlinks gets the allowInsertHyperlinks property value. Indicates whether the worksheet protection option to allow inserting hyperlinks is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertHyperlinks()(*bool) {
    return m.allowInsertHyperlinks
}
// GetAllowInsertRows gets the allowInsertRows property value. Indicates whether the worksheet protection option to allow inserting rows is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertRows()(*bool) {
    return m.allowInsertRows
}
// GetAllowPivotTables gets the allowPivotTables property value. Indicates whether the worksheet protection option to allow the use of the pivot table feature is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowPivotTables()(*bool) {
    return m.allowPivotTables
}
// GetAllowSort gets the allowSort property value. Indicates whether the worksheet protection option to allow the use of the sort feature is enabled.
// returns a *bool when successful
func (m *WorkbookWorksheetProtectionOptions) GetAllowSort()(*bool) {
    return m.allowSort
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookWorksheetProtectionOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowAutoFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAutoFilter(val)
        }
        return nil
    }
    res["allowDeleteColumns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeleteColumns(val)
        }
        return nil
    }
    res["allowDeleteRows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeleteRows(val)
        }
        return nil
    }
    res["allowFormatCells"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowFormatCells(val)
        }
        return nil
    }
    res["allowFormatColumns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowFormatColumns(val)
        }
        return nil
    }
    res["allowFormatRows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowFormatRows(val)
        }
        return nil
    }
    res["allowInsertColumns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowInsertColumns(val)
        }
        return nil
    }
    res["allowInsertHyperlinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowInsertHyperlinks(val)
        }
        return nil
    }
    res["allowInsertRows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowInsertRows(val)
        }
        return nil
    }
    res["allowPivotTables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPivotTables(val)
        }
        return nil
    }
    res["allowSort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSort(val)
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
func (m *WorkbookWorksheetProtectionOptions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *WorkbookWorksheetProtectionOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowAutoFilter", m.GetAllowAutoFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeleteColumns", m.GetAllowDeleteColumns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeleteRows", m.GetAllowDeleteRows())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowFormatCells", m.GetAllowFormatCells())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowFormatColumns", m.GetAllowFormatColumns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowFormatRows", m.GetAllowFormatRows())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowInsertColumns", m.GetAllowInsertColumns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowInsertHyperlinks", m.GetAllowInsertHyperlinks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowInsertRows", m.GetAllowInsertRows())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowPivotTables", m.GetAllowPivotTables())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowSort", m.GetAllowSort())
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
func (m *WorkbookWorksheetProtectionOptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowAutoFilter sets the allowAutoFilter property value. Indicates whether the worksheet protection option to allow the use of the autofilter feature is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowAutoFilter(value *bool)() {
    m.allowAutoFilter = value
}
// SetAllowDeleteColumns sets the allowDeleteColumns property value. Indicates whether the worksheet protection option to allow deleting columns is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowDeleteColumns(value *bool)() {
    m.allowDeleteColumns = value
}
// SetAllowDeleteRows sets the allowDeleteRows property value. Indicates whether the worksheet protection option to allow deleting rows is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowDeleteRows(value *bool)() {
    m.allowDeleteRows = value
}
// SetAllowFormatCells sets the allowFormatCells property value. Indicates whether the worksheet protection option to allow formatting cells is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatCells(value *bool)() {
    m.allowFormatCells = value
}
// SetAllowFormatColumns sets the allowFormatColumns property value. Indicates whether the worksheet protection option to allow formatting columns is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatColumns(value *bool)() {
    m.allowFormatColumns = value
}
// SetAllowFormatRows sets the allowFormatRows property value. Indicates whether the worksheet protection option to allow formatting rows is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatRows(value *bool)() {
    m.allowFormatRows = value
}
// SetAllowInsertColumns sets the allowInsertColumns property value. Indicates whether the worksheet protection option to allow inserting columns is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertColumns(value *bool)() {
    m.allowInsertColumns = value
}
// SetAllowInsertHyperlinks sets the allowInsertHyperlinks property value. Indicates whether the worksheet protection option to allow inserting hyperlinks is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertHyperlinks(value *bool)() {
    m.allowInsertHyperlinks = value
}
// SetAllowInsertRows sets the allowInsertRows property value. Indicates whether the worksheet protection option to allow inserting rows is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertRows(value *bool)() {
    m.allowInsertRows = value
}
// SetAllowPivotTables sets the allowPivotTables property value. Indicates whether the worksheet protection option to allow the use of the pivot table feature is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowPivotTables(value *bool)() {
    m.allowPivotTables = value
}
// SetAllowSort sets the allowSort property value. Indicates whether the worksheet protection option to allow the use of the sort feature is enabled.
func (m *WorkbookWorksheetProtectionOptions) SetAllowSort(value *bool)() {
    m.allowSort = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkbookWorksheetProtectionOptions) SetOdataType(value *string)() {
    m.odataType = value
}
type WorkbookWorksheetProtectionOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAutoFilter()(*bool)
    GetAllowDeleteColumns()(*bool)
    GetAllowDeleteRows()(*bool)
    GetAllowFormatCells()(*bool)
    GetAllowFormatColumns()(*bool)
    GetAllowFormatRows()(*bool)
    GetAllowInsertColumns()(*bool)
    GetAllowInsertHyperlinks()(*bool)
    GetAllowInsertRows()(*bool)
    GetAllowPivotTables()(*bool)
    GetAllowSort()(*bool)
    GetOdataType()(*string)
    SetAllowAutoFilter(value *bool)()
    SetAllowDeleteColumns(value *bool)()
    SetAllowDeleteRows(value *bool)()
    SetAllowFormatCells(value *bool)()
    SetAllowFormatColumns(value *bool)()
    SetAllowFormatRows(value *bool)()
    SetAllowInsertColumns(value *bool)()
    SetAllowInsertHyperlinks(value *bool)()
    SetAllowInsertRows(value *bool)()
    SetAllowPivotTables(value *bool)()
    SetAllowSort(value *bool)()
    SetOdataType(value *string)()
}
