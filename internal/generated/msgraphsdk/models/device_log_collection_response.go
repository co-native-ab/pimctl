package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceLogCollectionResponse windows Log Collection request entity.
type DeviceLogCollectionResponse struct {
    Entity
    // The User Principal Name (UPN) of the user that enrolled the device.
    enrolledByUser *string
    // The DateTime of the expiration of the logs.
    expirationDateTimeUTC *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The UPN for who initiated the request.
    initiatedByUserPrincipalName *string
    // Indicates Intune device unique identifier.
    managedDeviceId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The DateTime the request was received.
    receivedDateTimeUTC *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The DateTime of the request.
    requestedDateTimeUTC *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The size of the logs in KB. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    sizeInKB *float64
    // AppLogUploadStatus
    status *AppLogUploadState
}
// NewDeviceLogCollectionResponse instantiates a new DeviceLogCollectionResponse and sets the default values.
func NewDeviceLogCollectionResponse()(*DeviceLogCollectionResponse) {
    m := &DeviceLogCollectionResponse{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceLogCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceLogCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceLogCollectionResponse(), nil
}
// GetEnrolledByUser gets the enrolledByUser property value. The User Principal Name (UPN) of the user that enrolled the device.
// returns a *string when successful
func (m *DeviceLogCollectionResponse) GetEnrolledByUser()(*string) {
    return m.enrolledByUser
}
// GetExpirationDateTimeUTC gets the expirationDateTimeUTC property value. The DateTime of the expiration of the logs.
// returns a *Time when successful
func (m *DeviceLogCollectionResponse) GetExpirationDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTimeUTC
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceLogCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enrolledByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledByUser(val)
        }
        return nil
    }
    res["expirationDateTimeUTC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTimeUTC(val)
        }
        return nil
    }
    res["initiatedByUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedByUserPrincipalName(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["receivedDateTimeUTC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTimeUTC(val)
        }
        return nil
    }
    res["requestedDateTimeUTC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedDateTimeUTC(val)
        }
        return nil
    }
    res["sizeInKB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInKB(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLogUploadState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AppLogUploadState))
        }
        return nil
    }
    return res
}
// GetInitiatedByUserPrincipalName gets the initiatedByUserPrincipalName property value. The UPN for who initiated the request.
// returns a *string when successful
func (m *DeviceLogCollectionResponse) GetInitiatedByUserPrincipalName()(*string) {
    return m.initiatedByUserPrincipalName
}
// GetManagedDeviceId gets the managedDeviceId property value. Indicates Intune device unique identifier.
// returns a *UUID when successful
func (m *DeviceLogCollectionResponse) GetManagedDeviceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.managedDeviceId
}
// GetReceivedDateTimeUTC gets the receivedDateTimeUTC property value. The DateTime the request was received.
// returns a *Time when successful
func (m *DeviceLogCollectionResponse) GetReceivedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTimeUTC
}
// GetRequestedDateTimeUTC gets the requestedDateTimeUTC property value. The DateTime of the request.
// returns a *Time when successful
func (m *DeviceLogCollectionResponse) GetRequestedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.requestedDateTimeUTC
}
// GetSizeInKB gets the sizeInKB property value. The size of the logs in KB. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a *float64 when successful
func (m *DeviceLogCollectionResponse) GetSizeInKB()(*float64) {
    return m.sizeInKB
}
// GetStatus gets the status property value. AppLogUploadStatus
// returns a *AppLogUploadState when successful
func (m *DeviceLogCollectionResponse) GetStatus()(*AppLogUploadState) {
    return m.status
}
// Serialize serializes information the current object
func (m *DeviceLogCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("enrolledByUser", m.GetEnrolledByUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTimeUTC", m.GetExpirationDateTimeUTC())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedByUserPrincipalName", m.GetInitiatedByUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTimeUTC", m.GetReceivedDateTimeUTC())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestedDateTimeUTC", m.GetRequestedDateTimeUTC())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("sizeInKB", m.GetSizeInKB())
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
    return nil
}
// SetEnrolledByUser sets the enrolledByUser property value. The User Principal Name (UPN) of the user that enrolled the device.
func (m *DeviceLogCollectionResponse) SetEnrolledByUser(value *string)() {
    m.enrolledByUser = value
}
// SetExpirationDateTimeUTC sets the expirationDateTimeUTC property value. The DateTime of the expiration of the logs.
func (m *DeviceLogCollectionResponse) SetExpirationDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTimeUTC = value
}
// SetInitiatedByUserPrincipalName sets the initiatedByUserPrincipalName property value. The UPN for who initiated the request.
func (m *DeviceLogCollectionResponse) SetInitiatedByUserPrincipalName(value *string)() {
    m.initiatedByUserPrincipalName = value
}
// SetManagedDeviceId sets the managedDeviceId property value. Indicates Intune device unique identifier.
func (m *DeviceLogCollectionResponse) SetManagedDeviceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.managedDeviceId = value
}
// SetReceivedDateTimeUTC sets the receivedDateTimeUTC property value. The DateTime the request was received.
func (m *DeviceLogCollectionResponse) SetReceivedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTimeUTC = value
}
// SetRequestedDateTimeUTC sets the requestedDateTimeUTC property value. The DateTime of the request.
func (m *DeviceLogCollectionResponse) SetRequestedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestedDateTimeUTC = value
}
// SetSizeInKB sets the sizeInKB property value. The size of the logs in KB. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *DeviceLogCollectionResponse) SetSizeInKB(value *float64)() {
    m.sizeInKB = value
}
// SetStatus sets the status property value. AppLogUploadStatus
func (m *DeviceLogCollectionResponse) SetStatus(value *AppLogUploadState)() {
    m.status = value
}
type DeviceLogCollectionResponseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnrolledByUser()(*string)
    GetExpirationDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInitiatedByUserPrincipalName()(*string)
    GetManagedDeviceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetReceivedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSizeInKB()(*float64)
    GetStatus()(*AppLogUploadState)
    SetEnrolledByUser(value *string)()
    SetExpirationDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInitiatedByUserPrincipalName(value *string)()
    SetManagedDeviceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetReceivedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSizeInKB(value *float64)()
    SetStatus(value *AppLogUploadState)()
}
