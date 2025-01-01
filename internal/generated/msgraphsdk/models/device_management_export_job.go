package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementExportJob entity representing a job to export a report.
type DeviceManagementExportJob struct {
    Entity
    // Time that the exported report expires
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filters applied on the report
    filter *string
    // Possible values for the file format of a report to be exported.
    format *DeviceManagementReportFileFormat
    // Configures how the requested export job is localized.
    localizationType *DeviceManagementExportJobLocalizationType
    // Name of the report
    reportName *string
    // Time that the exported report was requested
    requestDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Columns selected from the report
    selectEscaped []string
    // A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id.
    snapshotId *string
    // Possible statuses associated with a generated report.
    status *DeviceManagementReportStatus
    // Temporary location of the exported report
    url *string
}
// NewDeviceManagementExportJob instantiates a new DeviceManagementExportJob and sets the default values.
func NewDeviceManagementExportJob()(*DeviceManagementExportJob) {
    m := &DeviceManagementExportJob{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementExportJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementExportJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementExportJob(), nil
}
// GetExpirationDateTime gets the expirationDateTime property value. Time that the exported report expires
// returns a *Time when successful
func (m *DeviceManagementExportJob) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementExportJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["filter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportFileFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*DeviceManagementReportFileFormat))
        }
        return nil
    }
    res["localizationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExportJobLocalizationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalizationType(val.(*DeviceManagementExportJobLocalizationType))
        }
        return nil
    }
    res["reportName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportName(val)
        }
        return nil
    }
    res["requestDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDateTime(val)
        }
        return nil
    }
    res["select"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSelectEscaped(res)
        }
        return nil
    }
    res["snapshotId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnapshotId(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DeviceManagementReportStatus))
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
// GetFilter gets the filter property value. Filters applied on the report
// returns a *string when successful
func (m *DeviceManagementExportJob) GetFilter()(*string) {
    return m.filter
}
// GetFormat gets the format property value. Possible values for the file format of a report to be exported.
// returns a *DeviceManagementReportFileFormat when successful
func (m *DeviceManagementExportJob) GetFormat()(*DeviceManagementReportFileFormat) {
    return m.format
}
// GetLocalizationType gets the localizationType property value. Configures how the requested export job is localized.
// returns a *DeviceManagementExportJobLocalizationType when successful
func (m *DeviceManagementExportJob) GetLocalizationType()(*DeviceManagementExportJobLocalizationType) {
    return m.localizationType
}
// GetReportName gets the reportName property value. Name of the report
// returns a *string when successful
func (m *DeviceManagementExportJob) GetReportName()(*string) {
    return m.reportName
}
// GetRequestDateTime gets the requestDateTime property value. Time that the exported report was requested
// returns a *Time when successful
func (m *DeviceManagementExportJob) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.requestDateTime
}
// GetSelectEscaped gets the select property value. Columns selected from the report
// returns a []string when successful
func (m *DeviceManagementExportJob) GetSelectEscaped()([]string) {
    return m.selectEscaped
}
// GetSnapshotId gets the snapshotId property value. A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id.
// returns a *string when successful
func (m *DeviceManagementExportJob) GetSnapshotId()(*string) {
    return m.snapshotId
}
// GetStatus gets the status property value. Possible statuses associated with a generated report.
// returns a *DeviceManagementReportStatus when successful
func (m *DeviceManagementExportJob) GetStatus()(*DeviceManagementReportStatus) {
    return m.status
}
// GetUrl gets the url property value. Temporary location of the exported report
// returns a *string when successful
func (m *DeviceManagementExportJob) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *DeviceManagementExportJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    if m.GetFormat() != nil {
        cast := (*m.GetFormat()).String()
        err = writer.WriteStringValue("format", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalizationType() != nil {
        cast := (*m.GetLocalizationType()).String()
        err = writer.WriteStringValue("localizationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportName", m.GetReportName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestDateTime", m.GetRequestDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSelectEscaped() != nil {
        err = writer.WriteCollectionOfStringValues("select", m.GetSelectEscaped())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("snapshotId", m.GetSnapshotId())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpirationDateTime sets the expirationDateTime property value. Time that the exported report expires
func (m *DeviceManagementExportJob) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetFilter sets the filter property value. Filters applied on the report
func (m *DeviceManagementExportJob) SetFilter(value *string)() {
    m.filter = value
}
// SetFormat sets the format property value. Possible values for the file format of a report to be exported.
func (m *DeviceManagementExportJob) SetFormat(value *DeviceManagementReportFileFormat)() {
    m.format = value
}
// SetLocalizationType sets the localizationType property value. Configures how the requested export job is localized.
func (m *DeviceManagementExportJob) SetLocalizationType(value *DeviceManagementExportJobLocalizationType)() {
    m.localizationType = value
}
// SetReportName sets the reportName property value. Name of the report
func (m *DeviceManagementExportJob) SetReportName(value *string)() {
    m.reportName = value
}
// SetRequestDateTime sets the requestDateTime property value. Time that the exported report was requested
func (m *DeviceManagementExportJob) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestDateTime = value
}
// SetSelectEscaped sets the select property value. Columns selected from the report
func (m *DeviceManagementExportJob) SetSelectEscaped(value []string)() {
    m.selectEscaped = value
}
// SetSnapshotId sets the snapshotId property value. A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id.
func (m *DeviceManagementExportJob) SetSnapshotId(value *string)() {
    m.snapshotId = value
}
// SetStatus sets the status property value. Possible statuses associated with a generated report.
func (m *DeviceManagementExportJob) SetStatus(value *DeviceManagementReportStatus)() {
    m.status = value
}
// SetUrl sets the url property value. Temporary location of the exported report
func (m *DeviceManagementExportJob) SetUrl(value *string)() {
    m.url = value
}
type DeviceManagementExportJobable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFilter()(*string)
    GetFormat()(*DeviceManagementReportFileFormat)
    GetLocalizationType()(*DeviceManagementExportJobLocalizationType)
    GetReportName()(*string)
    GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSelectEscaped()([]string)
    GetSnapshotId()(*string)
    GetStatus()(*DeviceManagementReportStatus)
    GetUrl()(*string)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFilter(value *string)()
    SetFormat(value *DeviceManagementReportFileFormat)()
    SetLocalizationType(value *DeviceManagementExportJobLocalizationType)()
    SetReportName(value *string)()
    SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSelectEscaped(value []string)()
    SetSnapshotId(value *string)()
    SetStatus(value *DeviceManagementReportStatus)()
    SetUrl(value *string)()
}
