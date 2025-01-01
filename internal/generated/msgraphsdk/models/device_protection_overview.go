package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceProtectionOverview hardware information of a given device.
type DeviceProtectionOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates number of devices reporting as clean
    cleanDeviceCount *int32
    // Indicates number of devices with critical failures
    criticalFailuresDeviceCount *int32
    // Indicates number of devices with inactive threat agent
    inactiveThreatAgentDeviceCount *int32
    // The OdataType property
    odataType *string
    // Indicates number of devices pending full scan
    pendingFullScanDeviceCount *int32
    // Indicates number of devices with pending manual steps
    pendingManualStepsDeviceCount *int32
    // Indicates number of pending offline scan devices
    pendingOfflineScanDeviceCount *int32
    // Indicates the number of devices that have a pending full scan. Valid values -2147483648 to 2147483647
    pendingQuickScanDeviceCount *int32
    // Indicates number of devices pending restart
    pendingRestartDeviceCount *int32
    // Indicates number of devices with an old signature
    pendingSignatureUpdateDeviceCount *int32
    // Total device count.
    totalReportedDeviceCount *int32
    // Indicates number of devices with threat agent state as unknown
    unknownStateThreatAgentDeviceCount *int32
}
// NewDeviceProtectionOverview instantiates a new DeviceProtectionOverview and sets the default values.
func NewDeviceProtectionOverview()(*DeviceProtectionOverview) {
    m := &DeviceProtectionOverview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceProtectionOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceProtectionOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceProtectionOverview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceProtectionOverview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCleanDeviceCount gets the cleanDeviceCount property value. Indicates number of devices reporting as clean
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetCleanDeviceCount()(*int32) {
    return m.cleanDeviceCount
}
// GetCriticalFailuresDeviceCount gets the criticalFailuresDeviceCount property value. Indicates number of devices with critical failures
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetCriticalFailuresDeviceCount()(*int32) {
    return m.criticalFailuresDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceProtectionOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cleanDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCleanDeviceCount(val)
        }
        return nil
    }
    res["criticalFailuresDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriticalFailuresDeviceCount(val)
        }
        return nil
    }
    res["inactiveThreatAgentDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactiveThreatAgentDeviceCount(val)
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
    res["pendingFullScanDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingFullScanDeviceCount(val)
        }
        return nil
    }
    res["pendingManualStepsDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingManualStepsDeviceCount(val)
        }
        return nil
    }
    res["pendingOfflineScanDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingOfflineScanDeviceCount(val)
        }
        return nil
    }
    res["pendingQuickScanDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingQuickScanDeviceCount(val)
        }
        return nil
    }
    res["pendingRestartDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingRestartDeviceCount(val)
        }
        return nil
    }
    res["pendingSignatureUpdateDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingSignatureUpdateDeviceCount(val)
        }
        return nil
    }
    res["totalReportedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalReportedDeviceCount(val)
        }
        return nil
    }
    res["unknownStateThreatAgentDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownStateThreatAgentDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetInactiveThreatAgentDeviceCount gets the inactiveThreatAgentDeviceCount property value. Indicates number of devices with inactive threat agent
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetInactiveThreatAgentDeviceCount()(*int32) {
    return m.inactiveThreatAgentDeviceCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DeviceProtectionOverview) GetOdataType()(*string) {
    return m.odataType
}
// GetPendingFullScanDeviceCount gets the pendingFullScanDeviceCount property value. Indicates number of devices pending full scan
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetPendingFullScanDeviceCount()(*int32) {
    return m.pendingFullScanDeviceCount
}
// GetPendingManualStepsDeviceCount gets the pendingManualStepsDeviceCount property value. Indicates number of devices with pending manual steps
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetPendingManualStepsDeviceCount()(*int32) {
    return m.pendingManualStepsDeviceCount
}
// GetPendingOfflineScanDeviceCount gets the pendingOfflineScanDeviceCount property value. Indicates number of pending offline scan devices
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetPendingOfflineScanDeviceCount()(*int32) {
    return m.pendingOfflineScanDeviceCount
}
// GetPendingQuickScanDeviceCount gets the pendingQuickScanDeviceCount property value. Indicates the number of devices that have a pending full scan. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetPendingQuickScanDeviceCount()(*int32) {
    return m.pendingQuickScanDeviceCount
}
// GetPendingRestartDeviceCount gets the pendingRestartDeviceCount property value. Indicates number of devices pending restart
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetPendingRestartDeviceCount()(*int32) {
    return m.pendingRestartDeviceCount
}
// GetPendingSignatureUpdateDeviceCount gets the pendingSignatureUpdateDeviceCount property value. Indicates number of devices with an old signature
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetPendingSignatureUpdateDeviceCount()(*int32) {
    return m.pendingSignatureUpdateDeviceCount
}
// GetTotalReportedDeviceCount gets the totalReportedDeviceCount property value. Total device count.
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetTotalReportedDeviceCount()(*int32) {
    return m.totalReportedDeviceCount
}
// GetUnknownStateThreatAgentDeviceCount gets the unknownStateThreatAgentDeviceCount property value. Indicates number of devices with threat agent state as unknown
// returns a *int32 when successful
func (m *DeviceProtectionOverview) GetUnknownStateThreatAgentDeviceCount()(*int32) {
    return m.unknownStateThreatAgentDeviceCount
}
// Serialize serializes information the current object
func (m *DeviceProtectionOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("cleanDeviceCount", m.GetCleanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("criticalFailuresDeviceCount", m.GetCriticalFailuresDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inactiveThreatAgentDeviceCount", m.GetInactiveThreatAgentDeviceCount())
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
        err := writer.WriteInt32Value("pendingFullScanDeviceCount", m.GetPendingFullScanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingManualStepsDeviceCount", m.GetPendingManualStepsDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingOfflineScanDeviceCount", m.GetPendingOfflineScanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingQuickScanDeviceCount", m.GetPendingQuickScanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingRestartDeviceCount", m.GetPendingRestartDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingSignatureUpdateDeviceCount", m.GetPendingSignatureUpdateDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalReportedDeviceCount", m.GetTotalReportedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownStateThreatAgentDeviceCount", m.GetUnknownStateThreatAgentDeviceCount())
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
func (m *DeviceProtectionOverview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCleanDeviceCount sets the cleanDeviceCount property value. Indicates number of devices reporting as clean
func (m *DeviceProtectionOverview) SetCleanDeviceCount(value *int32)() {
    m.cleanDeviceCount = value
}
// SetCriticalFailuresDeviceCount sets the criticalFailuresDeviceCount property value. Indicates number of devices with critical failures
func (m *DeviceProtectionOverview) SetCriticalFailuresDeviceCount(value *int32)() {
    m.criticalFailuresDeviceCount = value
}
// SetInactiveThreatAgentDeviceCount sets the inactiveThreatAgentDeviceCount property value. Indicates number of devices with inactive threat agent
func (m *DeviceProtectionOverview) SetInactiveThreatAgentDeviceCount(value *int32)() {
    m.inactiveThreatAgentDeviceCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceProtectionOverview) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPendingFullScanDeviceCount sets the pendingFullScanDeviceCount property value. Indicates number of devices pending full scan
func (m *DeviceProtectionOverview) SetPendingFullScanDeviceCount(value *int32)() {
    m.pendingFullScanDeviceCount = value
}
// SetPendingManualStepsDeviceCount sets the pendingManualStepsDeviceCount property value. Indicates number of devices with pending manual steps
func (m *DeviceProtectionOverview) SetPendingManualStepsDeviceCount(value *int32)() {
    m.pendingManualStepsDeviceCount = value
}
// SetPendingOfflineScanDeviceCount sets the pendingOfflineScanDeviceCount property value. Indicates number of pending offline scan devices
func (m *DeviceProtectionOverview) SetPendingOfflineScanDeviceCount(value *int32)() {
    m.pendingOfflineScanDeviceCount = value
}
// SetPendingQuickScanDeviceCount sets the pendingQuickScanDeviceCount property value. Indicates the number of devices that have a pending full scan. Valid values -2147483648 to 2147483647
func (m *DeviceProtectionOverview) SetPendingQuickScanDeviceCount(value *int32)() {
    m.pendingQuickScanDeviceCount = value
}
// SetPendingRestartDeviceCount sets the pendingRestartDeviceCount property value. Indicates number of devices pending restart
func (m *DeviceProtectionOverview) SetPendingRestartDeviceCount(value *int32)() {
    m.pendingRestartDeviceCount = value
}
// SetPendingSignatureUpdateDeviceCount sets the pendingSignatureUpdateDeviceCount property value. Indicates number of devices with an old signature
func (m *DeviceProtectionOverview) SetPendingSignatureUpdateDeviceCount(value *int32)() {
    m.pendingSignatureUpdateDeviceCount = value
}
// SetTotalReportedDeviceCount sets the totalReportedDeviceCount property value. Total device count.
func (m *DeviceProtectionOverview) SetTotalReportedDeviceCount(value *int32)() {
    m.totalReportedDeviceCount = value
}
// SetUnknownStateThreatAgentDeviceCount sets the unknownStateThreatAgentDeviceCount property value. Indicates number of devices with threat agent state as unknown
func (m *DeviceProtectionOverview) SetUnknownStateThreatAgentDeviceCount(value *int32)() {
    m.unknownStateThreatAgentDeviceCount = value
}
type DeviceProtectionOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCleanDeviceCount()(*int32)
    GetCriticalFailuresDeviceCount()(*int32)
    GetInactiveThreatAgentDeviceCount()(*int32)
    GetOdataType()(*string)
    GetPendingFullScanDeviceCount()(*int32)
    GetPendingManualStepsDeviceCount()(*int32)
    GetPendingOfflineScanDeviceCount()(*int32)
    GetPendingQuickScanDeviceCount()(*int32)
    GetPendingRestartDeviceCount()(*int32)
    GetPendingSignatureUpdateDeviceCount()(*int32)
    GetTotalReportedDeviceCount()(*int32)
    GetUnknownStateThreatAgentDeviceCount()(*int32)
    SetCleanDeviceCount(value *int32)()
    SetCriticalFailuresDeviceCount(value *int32)()
    SetInactiveThreatAgentDeviceCount(value *int32)()
    SetOdataType(value *string)()
    SetPendingFullScanDeviceCount(value *int32)()
    SetPendingManualStepsDeviceCount(value *int32)()
    SetPendingOfflineScanDeviceCount(value *int32)()
    SetPendingQuickScanDeviceCount(value *int32)()
    SetPendingRestartDeviceCount(value *int32)()
    SetPendingSignatureUpdateDeviceCount(value *int32)()
    SetTotalReportedDeviceCount(value *int32)()
    SetUnknownStateThreatAgentDeviceCount(value *int32)()
}
