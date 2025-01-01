package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcDeviceImage struct {
    Entity
    // The display name of the associated device image. The device image display name and the version are used to uniquely identify the Cloud PC device image. Read-only.
    displayName *string
    // The error code of the status of the image that indicates why the upload failed, if applicable. Possible values are: internalServerError, sourceImageNotFound, osVersionNotSupported, sourceImageInvalid, sourceImageNotGeneralized, unknownFutureValue, vmAlreadyAzureAdJoined, paidSourceImageNotSupport, sourceImageNotSupportCustomizeVMName, sourceImageSizeExceedsLimitation. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: vmAlreadyAzureAdJoined, paidSourceImageNotSupport, sourceImageNotSupportCustomizeVMName, sourceImageSizeExceedsLimitation. Read-only.
    errorCode *CloudPcDeviceImageErrorCode
    // The date when the image became unavailable. Read-only.
    expirationDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The data and time when the image was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The operating system (OS) of the image. For example, Windows 10 Enterprise. Read-only.
    operatingSystem *string
    // The OS build version of the image. For example, 1909. Read-only.
    osBuildNumber *string
    // The OS status of this image. Possible values are: supported, supportedWithWarning, unknown, unknownFutureValue. The default value is unknown. Read-only.
    osStatus *CloudPcDeviceImageOsStatus
    // The unique identifier (ID) of the source image resource on Azure. The required ID format is: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'. Read-only.
    sourceImageResourceId *string
    // The status of the image on the Cloud PC. Possible values are: pending, ready, failed, unknownFutureValue. Read-only.
    status *CloudPcDeviceImageStatus
    // The image version. For example, 0.0.1 and 1.5.13. Read-only.
    version *string
}
// NewCloudPcDeviceImage instantiates a new CloudPcDeviceImage and sets the default values.
func NewCloudPcDeviceImage()(*CloudPcDeviceImage) {
    m := &CloudPcDeviceImage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcDeviceImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcDeviceImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcDeviceImage(), nil
}
// GetDisplayName gets the displayName property value. The display name of the associated device image. The device image display name and the version are used to uniquely identify the Cloud PC device image. Read-only.
// returns a *string when successful
func (m *CloudPcDeviceImage) GetDisplayName()(*string) {
    return m.displayName
}
// GetErrorCode gets the errorCode property value. The error code of the status of the image that indicates why the upload failed, if applicable. Possible values are: internalServerError, sourceImageNotFound, osVersionNotSupported, sourceImageInvalid, sourceImageNotGeneralized, unknownFutureValue, vmAlreadyAzureAdJoined, paidSourceImageNotSupport, sourceImageNotSupportCustomizeVMName, sourceImageSizeExceedsLimitation. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: vmAlreadyAzureAdJoined, paidSourceImageNotSupport, sourceImageNotSupportCustomizeVMName, sourceImageSizeExceedsLimitation. Read-only.
// returns a *CloudPcDeviceImageErrorCode when successful
func (m *CloudPcDeviceImage) GetErrorCode()(*CloudPcDeviceImageErrorCode) {
    return m.errorCode
}
// GetExpirationDate gets the expirationDate property value. The date when the image became unavailable. Read-only.
// returns a *DateOnly when successful
func (m *CloudPcDeviceImage) GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.expirationDate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcDeviceImage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDeviceImageErrorCode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val.(*CloudPcDeviceImageErrorCode))
        }
        return nil
    }
    res["expirationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDate(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osBuildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
        }
        return nil
    }
    res["osStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDeviceImageOsStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsStatus(val.(*CloudPcDeviceImageOsStatus))
        }
        return nil
    }
    res["sourceImageResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceImageResourceId(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDeviceImageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcDeviceImageStatus))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The data and time when the image was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *CloudPcDeviceImage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOperatingSystem gets the operatingSystem property value. The operating system (OS) of the image. For example, Windows 10 Enterprise. Read-only.
// returns a *string when successful
func (m *CloudPcDeviceImage) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOsBuildNumber gets the osBuildNumber property value. The OS build version of the image. For example, 1909. Read-only.
// returns a *string when successful
func (m *CloudPcDeviceImage) GetOsBuildNumber()(*string) {
    return m.osBuildNumber
}
// GetOsStatus gets the osStatus property value. The OS status of this image. Possible values are: supported, supportedWithWarning, unknown, unknownFutureValue. The default value is unknown. Read-only.
// returns a *CloudPcDeviceImageOsStatus when successful
func (m *CloudPcDeviceImage) GetOsStatus()(*CloudPcDeviceImageOsStatus) {
    return m.osStatus
}
// GetSourceImageResourceId gets the sourceImageResourceId property value. The unique identifier (ID) of the source image resource on Azure. The required ID format is: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'. Read-only.
// returns a *string when successful
func (m *CloudPcDeviceImage) GetSourceImageResourceId()(*string) {
    return m.sourceImageResourceId
}
// GetStatus gets the status property value. The status of the image on the Cloud PC. Possible values are: pending, ready, failed, unknownFutureValue. Read-only.
// returns a *CloudPcDeviceImageStatus when successful
func (m *CloudPcDeviceImage) GetStatus()(*CloudPcDeviceImageStatus) {
    return m.status
}
// GetVersion gets the version property value. The image version. For example, 0.0.1 and 1.5.13. Read-only.
// returns a *string when successful
func (m *CloudPcDeviceImage) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *CloudPcDeviceImage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetErrorCode() != nil {
        cast := (*m.GetErrorCode()).String()
        err = writer.WriteStringValue("errorCode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("expirationDate", m.GetExpirationDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osBuildNumber", m.GetOsBuildNumber())
        if err != nil {
            return err
        }
    }
    if m.GetOsStatus() != nil {
        cast := (*m.GetOsStatus()).String()
        err = writer.WriteStringValue("osStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceImageResourceId", m.GetSourceImageResourceId())
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
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the associated device image. The device image display name and the version are used to uniquely identify the Cloud PC device image. Read-only.
func (m *CloudPcDeviceImage) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetErrorCode sets the errorCode property value. The error code of the status of the image that indicates why the upload failed, if applicable. Possible values are: internalServerError, sourceImageNotFound, osVersionNotSupported, sourceImageInvalid, sourceImageNotGeneralized, unknownFutureValue, vmAlreadyAzureAdJoined, paidSourceImageNotSupport, sourceImageNotSupportCustomizeVMName, sourceImageSizeExceedsLimitation. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: vmAlreadyAzureAdJoined, paidSourceImageNotSupport, sourceImageNotSupportCustomizeVMName, sourceImageSizeExceedsLimitation. Read-only.
func (m *CloudPcDeviceImage) SetErrorCode(value *CloudPcDeviceImageErrorCode)() {
    m.errorCode = value
}
// SetExpirationDate sets the expirationDate property value. The date when the image became unavailable. Read-only.
func (m *CloudPcDeviceImage) SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.expirationDate = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The data and time when the image was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *CloudPcDeviceImage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOperatingSystem sets the operatingSystem property value. The operating system (OS) of the image. For example, Windows 10 Enterprise. Read-only.
func (m *CloudPcDeviceImage) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOsBuildNumber sets the osBuildNumber property value. The OS build version of the image. For example, 1909. Read-only.
func (m *CloudPcDeviceImage) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// SetOsStatus sets the osStatus property value. The OS status of this image. Possible values are: supported, supportedWithWarning, unknown, unknownFutureValue. The default value is unknown. Read-only.
func (m *CloudPcDeviceImage) SetOsStatus(value *CloudPcDeviceImageOsStatus)() {
    m.osStatus = value
}
// SetSourceImageResourceId sets the sourceImageResourceId property value. The unique identifier (ID) of the source image resource on Azure. The required ID format is: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'. Read-only.
func (m *CloudPcDeviceImage) SetSourceImageResourceId(value *string)() {
    m.sourceImageResourceId = value
}
// SetStatus sets the status property value. The status of the image on the Cloud PC. Possible values are: pending, ready, failed, unknownFutureValue. Read-only.
func (m *CloudPcDeviceImage) SetStatus(value *CloudPcDeviceImageStatus)() {
    m.status = value
}
// SetVersion sets the version property value. The image version. For example, 0.0.1 and 1.5.13. Read-only.
func (m *CloudPcDeviceImage) SetVersion(value *string)() {
    m.version = value
}
type CloudPcDeviceImageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetErrorCode()(*CloudPcDeviceImageErrorCode)
    GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOperatingSystem()(*string)
    GetOsBuildNumber()(*string)
    GetOsStatus()(*CloudPcDeviceImageOsStatus)
    GetSourceImageResourceId()(*string)
    GetStatus()(*CloudPcDeviceImageStatus)
    GetVersion()(*string)
    SetDisplayName(value *string)()
    SetErrorCode(value *CloudPcDeviceImageErrorCode)()
    SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOperatingSystem(value *string)()
    SetOsBuildNumber(value *string)()
    SetOsStatus(value *CloudPcDeviceImageOsStatus)()
    SetSourceImageResourceId(value *string)()
    SetStatus(value *CloudPcDeviceImageStatus)()
    SetVersion(value *string)()
}
