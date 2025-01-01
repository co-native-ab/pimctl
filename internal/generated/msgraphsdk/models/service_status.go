package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type of consumer. The possible values are: unknown, firstparty, thirdparty, unknownFutureValue.
    backupServiceConsumer *BackupServiceConsumer
    // The reason the service is disabled. The possible values are: none, controllerServiceAppDeleted, invalidBillingProfile, userRequested, unknownFutureValue.
    disableReason *DisableReason
    // The expiration time of the grace period.
    gracePeriodDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identity of the person who last modified the entity.
    lastModifiedBy IdentitySetable
    // Timestamp of the last modification of the entity.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The expiration time of the restoration allowed period.
    restoreAllowedTillDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Status of the service. This value indicates what capabilities can be used. The possible values are: disabled, enabled, protectionChangeLocked, restoreLocked, unknownFutureValue.
    status *BackupServiceStatus
}
// NewServiceStatus instantiates a new ServiceStatus and sets the default values.
func NewServiceStatus()(*ServiceStatus) {
    m := &ServiceStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServiceStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBackupServiceConsumer gets the backupServiceConsumer property value. The type of consumer. The possible values are: unknown, firstparty, thirdparty, unknownFutureValue.
// returns a *BackupServiceConsumer when successful
func (m *ServiceStatus) GetBackupServiceConsumer()(*BackupServiceConsumer) {
    return m.backupServiceConsumer
}
// GetDisableReason gets the disableReason property value. The reason the service is disabled. The possible values are: none, controllerServiceAppDeleted, invalidBillingProfile, userRequested, unknownFutureValue.
// returns a *DisableReason when successful
func (m *ServiceStatus) GetDisableReason()(*DisableReason) {
    return m.disableReason
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["backupServiceConsumer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBackupServiceConsumer)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackupServiceConsumer(val.(*BackupServiceConsumer))
        }
        return nil
    }
    res["disableReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDisableReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableReason(val.(*DisableReason))
        }
        return nil
    }
    res["gracePeriodDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
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
    res["restoreAllowedTillDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoreAllowedTillDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBackupServiceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*BackupServiceStatus))
        }
        return nil
    }
    return res
}
// GetGracePeriodDateTime gets the gracePeriodDateTime property value. The expiration time of the grace period.
// returns a *Time when successful
func (m *ServiceStatus) GetGracePeriodDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.gracePeriodDateTime
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the person who last modified the entity.
// returns a IdentitySetable when successful
func (m *ServiceStatus) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Timestamp of the last modification of the entity.
// returns a *Time when successful
func (m *ServiceStatus) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ServiceStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetRestoreAllowedTillDateTime gets the restoreAllowedTillDateTime property value. The expiration time of the restoration allowed period.
// returns a *Time when successful
func (m *ServiceStatus) GetRestoreAllowedTillDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.restoreAllowedTillDateTime
}
// GetStatus gets the status property value. Status of the service. This value indicates what capabilities can be used. The possible values are: disabled, enabled, protectionChangeLocked, restoreLocked, unknownFutureValue.
// returns a *BackupServiceStatus when successful
func (m *ServiceStatus) GetStatus()(*BackupServiceStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ServiceStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBackupServiceConsumer() != nil {
        cast := (*m.GetBackupServiceConsumer()).String()
        err := writer.WriteStringValue("backupServiceConsumer", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDisableReason() != nil {
        cast := (*m.GetDisableReason()).String()
        err := writer.WriteStringValue("disableReason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("gracePeriodDateTime", m.GetGracePeriodDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
        err := writer.WriteTimeValue("restoreAllowedTillDateTime", m.GetRestoreAllowedTillDateTime())
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
func (m *ServiceStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBackupServiceConsumer sets the backupServiceConsumer property value. The type of consumer. The possible values are: unknown, firstparty, thirdparty, unknownFutureValue.
func (m *ServiceStatus) SetBackupServiceConsumer(value *BackupServiceConsumer)() {
    m.backupServiceConsumer = value
}
// SetDisableReason sets the disableReason property value. The reason the service is disabled. The possible values are: none, controllerServiceAppDeleted, invalidBillingProfile, userRequested, unknownFutureValue.
func (m *ServiceStatus) SetDisableReason(value *DisableReason)() {
    m.disableReason = value
}
// SetGracePeriodDateTime sets the gracePeriodDateTime property value. The expiration time of the grace period.
func (m *ServiceStatus) SetGracePeriodDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.gracePeriodDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the person who last modified the entity.
func (m *ServiceStatus) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Timestamp of the last modification of the entity.
func (m *ServiceStatus) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ServiceStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRestoreAllowedTillDateTime sets the restoreAllowedTillDateTime property value. The expiration time of the restoration allowed period.
func (m *ServiceStatus) SetRestoreAllowedTillDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.restoreAllowedTillDateTime = value
}
// SetStatus sets the status property value. Status of the service. This value indicates what capabilities can be used. The possible values are: disabled, enabled, protectionChangeLocked, restoreLocked, unknownFutureValue.
func (m *ServiceStatus) SetStatus(value *BackupServiceStatus)() {
    m.status = value
}
type ServiceStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackupServiceConsumer()(*BackupServiceConsumer)
    GetDisableReason()(*DisableReason)
    GetGracePeriodDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetRestoreAllowedTillDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*BackupServiceStatus)
    SetBackupServiceConsumer(value *BackupServiceConsumer)()
    SetDisableReason(value *DisableReason)()
    SetGracePeriodDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetRestoreAllowedTillDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *BackupServiceStatus)()
}
