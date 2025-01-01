package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LicenseUnitsDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of units that are enabled for the active subscription of the service SKU.
    enabled *int32
    // The number of units that are locked out because the customer canceled their subscription of the service SKU.
    lockedOut *int32
    // The OdataType property
    odataType *string
    // The number of units that are suspended because the subscription of the service SKU has been canceled. The units can't be assigned but can still be reactivated before they're deleted.
    suspended *int32
    // The number of units that are in warning status. When the subscription of the service SKU has expired, the customer has a grace period to renew their subscription before it's canceled (moved to a suspended state).
    warning *int32
}
// NewLicenseUnitsDetail instantiates a new LicenseUnitsDetail and sets the default values.
func NewLicenseUnitsDetail()(*LicenseUnitsDetail) {
    m := &LicenseUnitsDetail{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLicenseUnitsDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLicenseUnitsDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLicenseUnitsDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LicenseUnitsDetail) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The number of units that are enabled for the active subscription of the service SKU.
// returns a *int32 when successful
func (m *LicenseUnitsDetail) GetEnabled()(*int32) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LicenseUnitsDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["lockedOut"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockedOut(val)
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
    res["suspended"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuspended(val)
        }
        return nil
    }
    res["warning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarning(val)
        }
        return nil
    }
    return res
}
// GetLockedOut gets the lockedOut property value. The number of units that are locked out because the customer canceled their subscription of the service SKU.
// returns a *int32 when successful
func (m *LicenseUnitsDetail) GetLockedOut()(*int32) {
    return m.lockedOut
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *LicenseUnitsDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetSuspended gets the suspended property value. The number of units that are suspended because the subscription of the service SKU has been canceled. The units can't be assigned but can still be reactivated before they're deleted.
// returns a *int32 when successful
func (m *LicenseUnitsDetail) GetSuspended()(*int32) {
    return m.suspended
}
// GetWarning gets the warning property value. The number of units that are in warning status. When the subscription of the service SKU has expired, the customer has a grace period to renew their subscription before it's canceled (moved to a suspended state).
// returns a *int32 when successful
func (m *LicenseUnitsDetail) GetWarning()(*int32) {
    return m.warning
}
// Serialize serializes information the current object
func (m *LicenseUnitsDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("lockedOut", m.GetLockedOut())
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
        err := writer.WriteInt32Value("suspended", m.GetSuspended())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("warning", m.GetWarning())
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
func (m *LicenseUnitsDetail) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The number of units that are enabled for the active subscription of the service SKU.
func (m *LicenseUnitsDetail) SetEnabled(value *int32)() {
    m.enabled = value
}
// SetLockedOut sets the lockedOut property value. The number of units that are locked out because the customer canceled their subscription of the service SKU.
func (m *LicenseUnitsDetail) SetLockedOut(value *int32)() {
    m.lockedOut = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LicenseUnitsDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSuspended sets the suspended property value. The number of units that are suspended because the subscription of the service SKU has been canceled. The units can't be assigned but can still be reactivated before they're deleted.
func (m *LicenseUnitsDetail) SetSuspended(value *int32)() {
    m.suspended = value
}
// SetWarning sets the warning property value. The number of units that are in warning status. When the subscription of the service SKU has expired, the customer has a grace period to renew their subscription before it's canceled (moved to a suspended state).
func (m *LicenseUnitsDetail) SetWarning(value *int32)() {
    m.warning = value
}
type LicenseUnitsDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*int32)
    GetLockedOut()(*int32)
    GetOdataType()(*string)
    GetSuspended()(*int32)
    GetWarning()(*int32)
    SetEnabled(value *int32)()
    SetLockedOut(value *int32)()
    SetOdataType(value *string)()
    SetSuspended(value *int32)()
    SetWarning(value *int32)()
}
