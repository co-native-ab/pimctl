package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookRange struct {
    Entity
    // Represents the range reference in A1-style. Address value contains the Sheet reference (for example, Sheet1!A1:B4). Read-only.
    address *string
    // Represents range reference for the specified range in the language of the user. Read-only.
    addressLocal *string
    // Number of cells in the range. Read-only.
    cellCount *int32
    // Represents the total number of columns in the range. Read-only.
    columnCount *int32
    // Indicates whether all columns of the current range are hidden.
    columnHidden *bool
    // Represents the column number of the first cell in the range. Zero-indexed. Read-only.
    columnIndex *int32
    // Returns a format object, encapsulating the range's font, fill, borders, alignment, and other properties. Read-only.
    format WorkbookRangeFormatable
    // Represents the formula in A1-style notation.
    formulas i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Represents the formula in A1-style notation, in the user's language and number-formatting locale.  For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
    formulasLocal i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Represents the formula in R1C1-style notation.
    formulasR1C1 i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Represents if all cells of the current range are hidden. Read-only.
    hidden *bool
    // Represents Excel's number format code for the given cell.
    numberFormat i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Returns the total number of rows in the range. Read-only.
    rowCount *int32
    // Indicates whether all rows of the current range are hidden.
    rowHidden *bool
    // Returns the row number of the first cell in the range. Zero-indexed. Read-only.
    rowIndex *int32
    // The worksheet containing the current range. Read-only.
    sort WorkbookRangeSortable
    // Text values of the specified range. The Text value doesn't depend on the cell width. The # sign substitution that happens in Excel UI doesn't affect the text value returned by the API. Read-only.
    text i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Represents the raw values of the specified range. The data returned can be of type string, number, or a Boolean. Cell that contains an error returns the error string.
    values i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Represents the type of data of each cell. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error. Read-only.
    valueTypes i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The worksheet containing the current range. Read-only.
    worksheet WorkbookWorksheetable
}
// NewWorkbookRange instantiates a new WorkbookRange and sets the default values.
func NewWorkbookRange()(*WorkbookRange) {
    m := &WorkbookRange{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRange(), nil
}
// GetAddress gets the address property value. Represents the range reference in A1-style. Address value contains the Sheet reference (for example, Sheet1!A1:B4). Read-only.
// returns a *string when successful
func (m *WorkbookRange) GetAddress()(*string) {
    return m.address
}
// GetAddressLocal gets the addressLocal property value. Represents range reference for the specified range in the language of the user. Read-only.
// returns a *string when successful
func (m *WorkbookRange) GetAddressLocal()(*string) {
    return m.addressLocal
}
// GetCellCount gets the cellCount property value. Number of cells in the range. Read-only.
// returns a *int32 when successful
func (m *WorkbookRange) GetCellCount()(*int32) {
    return m.cellCount
}
// GetColumnCount gets the columnCount property value. Represents the total number of columns in the range. Read-only.
// returns a *int32 when successful
func (m *WorkbookRange) GetColumnCount()(*int32) {
    return m.columnCount
}
// GetColumnHidden gets the columnHidden property value. Indicates whether all columns of the current range are hidden.
// returns a *bool when successful
func (m *WorkbookRange) GetColumnHidden()(*bool) {
    return m.columnHidden
}
// GetColumnIndex gets the columnIndex property value. Represents the column number of the first cell in the range. Zero-indexed. Read-only.
// returns a *int32 when successful
func (m *WorkbookRange) GetColumnIndex()(*int32) {
    return m.columnIndex
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["addressLocal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressLocal(val)
        }
        return nil
    }
    res["cellCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellCount(val)
        }
        return nil
    }
    res["columnCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnCount(val)
        }
        return nil
    }
    res["columnHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnHidden(val)
        }
        return nil
    }
    res["columnIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnIndex(val)
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookRangeFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookRangeFormatable))
        }
        return nil
    }
    res["formulas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulas(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["formulasLocal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasLocal(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["formulasR1C1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasR1C1(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["hidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidden(val)
        }
        return nil
    }
    res["numberFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberFormat(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["rowCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowCount(val)
        }
        return nil
    }
    res["rowHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowHidden(val)
        }
        return nil
    }
    res["rowIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowIndex(val)
        }
        return nil
    }
    res["sort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookRangeSortFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSort(val.(WorkbookRangeSortable))
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
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
    res["valueTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueTypes(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
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
// GetFormat gets the format property value. Returns a format object, encapsulating the range's font, fill, borders, alignment, and other properties. Read-only.
// returns a WorkbookRangeFormatable when successful
func (m *WorkbookRange) GetFormat()(WorkbookRangeFormatable) {
    return m.format
}
// GetFormulas gets the formulas property value. Represents the formula in A1-style notation.
// returns a UntypedNodeable when successful
func (m *WorkbookRange) GetFormulas()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.formulas
}
// GetFormulasLocal gets the formulasLocal property value. Represents the formula in A1-style notation, in the user's language and number-formatting locale.  For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
// returns a UntypedNodeable when successful
func (m *WorkbookRange) GetFormulasLocal()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.formulasLocal
}
// GetFormulasR1C1 gets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
// returns a UntypedNodeable when successful
func (m *WorkbookRange) GetFormulasR1C1()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.formulasR1C1
}
// GetHidden gets the hidden property value. Represents if all cells of the current range are hidden. Read-only.
// returns a *bool when successful
func (m *WorkbookRange) GetHidden()(*bool) {
    return m.hidden
}
// GetNumberFormat gets the numberFormat property value. Represents Excel's number format code for the given cell.
// returns a UntypedNodeable when successful
func (m *WorkbookRange) GetNumberFormat()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.numberFormat
}
// GetRowCount gets the rowCount property value. Returns the total number of rows in the range. Read-only.
// returns a *int32 when successful
func (m *WorkbookRange) GetRowCount()(*int32) {
    return m.rowCount
}
// GetRowHidden gets the rowHidden property value. Indicates whether all rows of the current range are hidden.
// returns a *bool when successful
func (m *WorkbookRange) GetRowHidden()(*bool) {
    return m.rowHidden
}
// GetRowIndex gets the rowIndex property value. Returns the row number of the first cell in the range. Zero-indexed. Read-only.
// returns a *int32 when successful
func (m *WorkbookRange) GetRowIndex()(*int32) {
    return m.rowIndex
}
// GetSort gets the sort property value. The worksheet containing the current range. Read-only.
// returns a WorkbookRangeSortable when successful
func (m *WorkbookRange) GetSort()(WorkbookRangeSortable) {
    return m.sort
}
// GetText gets the text property value. Text values of the specified range. The Text value doesn't depend on the cell width. The # sign substitution that happens in Excel UI doesn't affect the text value returned by the API. Read-only.
// returns a UntypedNodeable when successful
func (m *WorkbookRange) GetText()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.text
}
// GetValues gets the values property value. Represents the raw values of the specified range. The data returned can be of type string, number, or a Boolean. Cell that contains an error returns the error string.
// returns a UntypedNodeable when successful
func (m *WorkbookRange) GetValues()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.values
}
// GetValueTypes gets the valueTypes property value. Represents the type of data of each cell. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error. Read-only.
// returns a UntypedNodeable when successful
func (m *WorkbookRange) GetValueTypes()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.valueTypes
}
// GetWorksheet gets the worksheet property value. The worksheet containing the current range. Read-only.
// returns a WorkbookWorksheetable when successful
func (m *WorkbookRange) GetWorksheet()(WorkbookWorksheetable) {
    return m.worksheet
}
// Serialize serializes information the current object
func (m *WorkbookRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("addressLocal", m.GetAddressLocal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cellCount", m.GetCellCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("columnCount", m.GetColumnCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("columnHidden", m.GetColumnHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("columnIndex", m.GetColumnIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulas", m.GetFormulas())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulasLocal", m.GetFormulasLocal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulasR1C1", m.GetFormulasR1C1())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hidden", m.GetHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("numberFormat", m.GetNumberFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rowCount", m.GetRowCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("rowHidden", m.GetRowHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rowIndex", m.GetRowIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sort", m.GetSort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("valueTypes", m.GetValueTypes())
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
// SetAddress sets the address property value. Represents the range reference in A1-style. Address value contains the Sheet reference (for example, Sheet1!A1:B4). Read-only.
func (m *WorkbookRange) SetAddress(value *string)() {
    m.address = value
}
// SetAddressLocal sets the addressLocal property value. Represents range reference for the specified range in the language of the user. Read-only.
func (m *WorkbookRange) SetAddressLocal(value *string)() {
    m.addressLocal = value
}
// SetCellCount sets the cellCount property value. Number of cells in the range. Read-only.
func (m *WorkbookRange) SetCellCount(value *int32)() {
    m.cellCount = value
}
// SetColumnCount sets the columnCount property value. Represents the total number of columns in the range. Read-only.
func (m *WorkbookRange) SetColumnCount(value *int32)() {
    m.columnCount = value
}
// SetColumnHidden sets the columnHidden property value. Indicates whether all columns of the current range are hidden.
func (m *WorkbookRange) SetColumnHidden(value *bool)() {
    m.columnHidden = value
}
// SetColumnIndex sets the columnIndex property value. Represents the column number of the first cell in the range. Zero-indexed. Read-only.
func (m *WorkbookRange) SetColumnIndex(value *int32)() {
    m.columnIndex = value
}
// SetFormat sets the format property value. Returns a format object, encapsulating the range's font, fill, borders, alignment, and other properties. Read-only.
func (m *WorkbookRange) SetFormat(value WorkbookRangeFormatable)() {
    m.format = value
}
// SetFormulas sets the formulas property value. Represents the formula in A1-style notation.
func (m *WorkbookRange) SetFormulas(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.formulas = value
}
// SetFormulasLocal sets the formulasLocal property value. Represents the formula in A1-style notation, in the user's language and number-formatting locale.  For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
func (m *WorkbookRange) SetFormulasLocal(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.formulasLocal = value
}
// SetFormulasR1C1 sets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
func (m *WorkbookRange) SetFormulasR1C1(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.formulasR1C1 = value
}
// SetHidden sets the hidden property value. Represents if all cells of the current range are hidden. Read-only.
func (m *WorkbookRange) SetHidden(value *bool)() {
    m.hidden = value
}
// SetNumberFormat sets the numberFormat property value. Represents Excel's number format code for the given cell.
func (m *WorkbookRange) SetNumberFormat(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.numberFormat = value
}
// SetRowCount sets the rowCount property value. Returns the total number of rows in the range. Read-only.
func (m *WorkbookRange) SetRowCount(value *int32)() {
    m.rowCount = value
}
// SetRowHidden sets the rowHidden property value. Indicates whether all rows of the current range are hidden.
func (m *WorkbookRange) SetRowHidden(value *bool)() {
    m.rowHidden = value
}
// SetRowIndex sets the rowIndex property value. Returns the row number of the first cell in the range. Zero-indexed. Read-only.
func (m *WorkbookRange) SetRowIndex(value *int32)() {
    m.rowIndex = value
}
// SetSort sets the sort property value. The worksheet containing the current range. Read-only.
func (m *WorkbookRange) SetSort(value WorkbookRangeSortable)() {
    m.sort = value
}
// SetText sets the text property value. Text values of the specified range. The Text value doesn't depend on the cell width. The # sign substitution that happens in Excel UI doesn't affect the text value returned by the API. Read-only.
func (m *WorkbookRange) SetText(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.text = value
}
// SetValues sets the values property value. Represents the raw values of the specified range. The data returned can be of type string, number, or a Boolean. Cell that contains an error returns the error string.
func (m *WorkbookRange) SetValues(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.values = value
}
// SetValueTypes sets the valueTypes property value. Represents the type of data of each cell. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error. Read-only.
func (m *WorkbookRange) SetValueTypes(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.valueTypes = value
}
// SetWorksheet sets the worksheet property value. The worksheet containing the current range. Read-only.
func (m *WorkbookRange) SetWorksheet(value WorkbookWorksheetable)() {
    m.worksheet = value
}
type WorkbookRangeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetAddressLocal()(*string)
    GetCellCount()(*int32)
    GetColumnCount()(*int32)
    GetColumnHidden()(*bool)
    GetColumnIndex()(*int32)
    GetFormat()(WorkbookRangeFormatable)
    GetFormulas()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetFormulasLocal()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetFormulasR1C1()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetHidden()(*bool)
    GetNumberFormat()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetRowCount()(*int32)
    GetRowHidden()(*bool)
    GetRowIndex()(*int32)
    GetSort()(WorkbookRangeSortable)
    GetText()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetValues()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetValueTypes()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetWorksheet()(WorkbookWorksheetable)
    SetAddress(value *string)()
    SetAddressLocal(value *string)()
    SetCellCount(value *int32)()
    SetColumnCount(value *int32)()
    SetColumnHidden(value *bool)()
    SetColumnIndex(value *int32)()
    SetFormat(value WorkbookRangeFormatable)()
    SetFormulas(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetFormulasLocal(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetFormulasR1C1(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetHidden(value *bool)()
    SetNumberFormat(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetRowCount(value *int32)()
    SetRowHidden(value *bool)()
    SetRowIndex(value *int32)()
    SetSort(value WorkbookRangeSortable)()
    SetText(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetValues(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetValueTypes(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetWorksheet(value WorkbookWorksheetable)()
}
