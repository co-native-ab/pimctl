package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageAutomaticRequestSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The duration for which access must be retained before the target's access is revoked once they leave the allowed target scope.
    gracePeriodBeforeAccessRemoval *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The OdataType property
    odataType *string
    // Indicates whether automatic assignment must be removed for targets who move out of the allowed target scope.
    removeAccessWhenTargetLeavesAllowedTargets *bool
    // If set to true, automatic assignments will be created for targets in the allowed target scope.
    requestAccessForAllowedTargets *bool
}
// NewAccessPackageAutomaticRequestSettings instantiates a new AccessPackageAutomaticRequestSettings and sets the default values.
func NewAccessPackageAutomaticRequestSettings()(*AccessPackageAutomaticRequestSettings) {
    m := &AccessPackageAutomaticRequestSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessPackageAutomaticRequestSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageAutomaticRequestSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAutomaticRequestSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccessPackageAutomaticRequestSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageAutomaticRequestSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["gracePeriodBeforeAccessRemoval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodBeforeAccessRemoval(val)
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
    res["removeAccessWhenTargetLeavesAllowedTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveAccessWhenTargetLeavesAllowedTargets(val)
        }
        return nil
    }
    res["requestAccessForAllowedTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestAccessForAllowedTargets(val)
        }
        return nil
    }
    return res
}
// GetGracePeriodBeforeAccessRemoval gets the gracePeriodBeforeAccessRemoval property value. The duration for which access must be retained before the target's access is revoked once they leave the allowed target scope.
// returns a *ISODuration when successful
func (m *AccessPackageAutomaticRequestSettings) GetGracePeriodBeforeAccessRemoval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.gracePeriodBeforeAccessRemoval
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AccessPackageAutomaticRequestSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetRemoveAccessWhenTargetLeavesAllowedTargets gets the removeAccessWhenTargetLeavesAllowedTargets property value. Indicates whether automatic assignment must be removed for targets who move out of the allowed target scope.
// returns a *bool when successful
func (m *AccessPackageAutomaticRequestSettings) GetRemoveAccessWhenTargetLeavesAllowedTargets()(*bool) {
    return m.removeAccessWhenTargetLeavesAllowedTargets
}
// GetRequestAccessForAllowedTargets gets the requestAccessForAllowedTargets property value. If set to true, automatic assignments will be created for targets in the allowed target scope.
// returns a *bool when successful
func (m *AccessPackageAutomaticRequestSettings) GetRequestAccessForAllowedTargets()(*bool) {
    return m.requestAccessForAllowedTargets
}
// Serialize serializes information the current object
func (m *AccessPackageAutomaticRequestSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("gracePeriodBeforeAccessRemoval", m.GetGracePeriodBeforeAccessRemoval())
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
        err := writer.WriteBoolValue("removeAccessWhenTargetLeavesAllowedTargets", m.GetRemoveAccessWhenTargetLeavesAllowedTargets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requestAccessForAllowedTargets", m.GetRequestAccessForAllowedTargets())
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
func (m *AccessPackageAutomaticRequestSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGracePeriodBeforeAccessRemoval sets the gracePeriodBeforeAccessRemoval property value. The duration for which access must be retained before the target's access is revoked once they leave the allowed target scope.
func (m *AccessPackageAutomaticRequestSettings) SetGracePeriodBeforeAccessRemoval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.gracePeriodBeforeAccessRemoval = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageAutomaticRequestSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemoveAccessWhenTargetLeavesAllowedTargets sets the removeAccessWhenTargetLeavesAllowedTargets property value. Indicates whether automatic assignment must be removed for targets who move out of the allowed target scope.
func (m *AccessPackageAutomaticRequestSettings) SetRemoveAccessWhenTargetLeavesAllowedTargets(value *bool)() {
    m.removeAccessWhenTargetLeavesAllowedTargets = value
}
// SetRequestAccessForAllowedTargets sets the requestAccessForAllowedTargets property value. If set to true, automatic assignments will be created for targets in the allowed target scope.
func (m *AccessPackageAutomaticRequestSettings) SetRequestAccessForAllowedTargets(value *bool)() {
    m.requestAccessForAllowedTargets = value
}
type AccessPackageAutomaticRequestSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGracePeriodBeforeAccessRemoval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetOdataType()(*string)
    GetRemoveAccessWhenTargetLeavesAllowedTargets()(*bool)
    GetRequestAccessForAllowedTargets()(*bool)
    SetGracePeriodBeforeAccessRemoval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetOdataType(value *string)()
    SetRemoveAccessWhenTargetLeavesAllowedTargets(value *bool)()
    SetRequestAccessForAllowedTargets(value *bool)()
}
