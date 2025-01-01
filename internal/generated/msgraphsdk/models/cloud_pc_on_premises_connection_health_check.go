package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcOnPremisesConnectionHealthCheck struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Additional details about the health check or the recommended action. For exmaple, the string value can be download.microsoft.com:443;software-download.microsoft.com:443; Read-only.
    additionalDetail *string
    // The unique identifier of the health check item-related activities. This identifier can be useful in troubleshooting.
    correlationId *string
    // The display name for this health check item.
    displayName *string
    // The value cannot be modified and is automatically populated when the health check ends. The Timestamp type represents date and time information using ISO 8601 format and is in Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2024 would look like this: '2024-01-01T00:00:00Z'. Returned by default. Read-only.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of error that occurred during this health check. Possible values are: endpointConnectivityCheckCloudPcUrlNotAllowListed, endpointConnectivityCheckWVDUrlNotAllowListed, etc. (The all possible values can refer to cloudPcOnPremisesConnectionHealthCheckErrorType) Read-Only.
    errorType *CloudPcOnPremisesConnectionHealthCheckErrorType
    // The OdataType property
    odataType *string
    // The recommended action to fix the corresponding error. For example, The Active Directory domain join check failed because the password of the domain join user has expired. Read-Only.
    recommendedAction *string
    // The value cannot be modified and is automatically populated when the health check starts. The Timestamp type represents date and time information using ISO 8601 format and is in  Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2024 would look like this: '2024-01-01T00:00:00Z'. Returned by default. Read-only.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *CloudPcOnPremisesConnectionStatus
}
// NewCloudPcOnPremisesConnectionHealthCheck instantiates a new CloudPcOnPremisesConnectionHealthCheck and sets the default values.
func NewCloudPcOnPremisesConnectionHealthCheck()(*CloudPcOnPremisesConnectionHealthCheck) {
    m := &CloudPcOnPremisesConnectionHealthCheck{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcOnPremisesConnectionHealthCheckFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcOnPremisesConnectionHealthCheckFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcOnPremisesConnectionHealthCheck(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdditionalDetail gets the additionalDetail property value. Additional details about the health check or the recommended action. For exmaple, the string value can be download.microsoft.com:443;software-download.microsoft.com:443; Read-only.
// returns a *string when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetAdditionalDetail()(*string) {
    return m.additionalDetail
}
// GetCorrelationId gets the correlationId property value. The unique identifier of the health check item-related activities. This identifier can be useful in troubleshooting.
// returns a *string when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetCorrelationId()(*string) {
    return m.correlationId
}
// GetDisplayName gets the displayName property value. The display name for this health check item.
// returns a *string when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDateTime gets the endDateTime property value. The value cannot be modified and is automatically populated when the health check ends. The Timestamp type represents date and time information using ISO 8601 format and is in Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2024 would look like this: '2024-01-01T00:00:00Z'. Returned by default. Read-only.
// returns a *Time when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetErrorType gets the errorType property value. The type of error that occurred during this health check. Possible values are: endpointConnectivityCheckCloudPcUrlNotAllowListed, endpointConnectivityCheckWVDUrlNotAllowListed, etc. (The all possible values can refer to cloudPcOnPremisesConnectionHealthCheckErrorType) Read-Only.
// returns a *CloudPcOnPremisesConnectionHealthCheckErrorType when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetErrorType()(*CloudPcOnPremisesConnectionHealthCheckErrorType) {
    return m.errorType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDetail(val)
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
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
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["errorType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionHealthCheckErrorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorType(val.(*CloudPcOnPremisesConnectionHealthCheckErrorType))
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
    res["recommendedAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedAction(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcOnPremisesConnectionStatus))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetOdataType()(*string) {
    return m.odataType
}
// GetRecommendedAction gets the recommendedAction property value. The recommended action to fix the corresponding error. For example, The Active Directory domain join check failed because the password of the domain join user has expired. Read-Only.
// returns a *string when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetRecommendedAction()(*string) {
    return m.recommendedAction
}
// GetStartDateTime gets the startDateTime property value. The value cannot be modified and is automatically populated when the health check starts. The Timestamp type represents date and time information using ISO 8601 format and is in  Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2024 would look like this: '2024-01-01T00:00:00Z'. Returned by default. Read-only.
// returns a *Time when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. The status property
// returns a *CloudPcOnPremisesConnectionStatus when successful
func (m *CloudPcOnPremisesConnectionHealthCheck) GetStatus()(*CloudPcOnPremisesConnectionStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CloudPcOnPremisesConnectionHealthCheck) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("additionalDetail", m.GetAdditionalDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("correlationId", m.GetCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetErrorType() != nil {
        cast := (*m.GetErrorType()).String()
        err := writer.WriteStringValue("errorType", &cast)
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
        err := writer.WriteStringValue("recommendedAction", m.GetRecommendedAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *CloudPcOnPremisesConnectionHealthCheck) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdditionalDetail sets the additionalDetail property value. Additional details about the health check or the recommended action. For exmaple, the string value can be download.microsoft.com:443;software-download.microsoft.com:443; Read-only.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetAdditionalDetail(value *string)() {
    m.additionalDetail = value
}
// SetCorrelationId sets the correlationId property value. The unique identifier of the health check item-related activities. This identifier can be useful in troubleshooting.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetDisplayName sets the displayName property value. The display name for this health check item.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDateTime sets the endDateTime property value. The value cannot be modified and is automatically populated when the health check ends. The Timestamp type represents date and time information using ISO 8601 format and is in Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2024 would look like this: '2024-01-01T00:00:00Z'. Returned by default. Read-only.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetErrorType sets the errorType property value. The type of error that occurred during this health check. Possible values are: endpointConnectivityCheckCloudPcUrlNotAllowListed, endpointConnectivityCheckWVDUrlNotAllowListed, etc. (The all possible values can refer to cloudPcOnPremisesConnectionHealthCheckErrorType) Read-Only.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetErrorType(value *CloudPcOnPremisesConnectionHealthCheckErrorType)() {
    m.errorType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcOnPremisesConnectionHealthCheck) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecommendedAction sets the recommendedAction property value. The recommended action to fix the corresponding error. For example, The Active Directory domain join check failed because the password of the domain join user has expired. Read-Only.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetRecommendedAction(value *string)() {
    m.recommendedAction = value
}
// SetStartDateTime sets the startDateTime property value. The value cannot be modified and is automatically populated when the health check starts. The Timestamp type represents date and time information using ISO 8601 format and is in  Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2024 would look like this: '2024-01-01T00:00:00Z'. Returned by default. Read-only.
func (m *CloudPcOnPremisesConnectionHealthCheck) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *CloudPcOnPremisesConnectionHealthCheck) SetStatus(value *CloudPcOnPremisesConnectionStatus)() {
    m.status = value
}
type CloudPcOnPremisesConnectionHealthCheckable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalDetail()(*string)
    GetCorrelationId()(*string)
    GetDisplayName()(*string)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetErrorType()(*CloudPcOnPremisesConnectionHealthCheckErrorType)
    GetOdataType()(*string)
    GetRecommendedAction()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*CloudPcOnPremisesConnectionStatus)
    SetAdditionalDetail(value *string)()
    SetCorrelationId(value *string)()
    SetDisplayName(value *string)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetErrorType(value *CloudPcOnPremisesConnectionHealthCheckErrorType)()
    SetOdataType(value *string)()
    SetRecommendedAction(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *CloudPcOnPremisesConnectionStatus)()
}
