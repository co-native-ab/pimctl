package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type LifecycleManagementSettings struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The emailSettings property
    emailSettings ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.EmailSettingsable
    // The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
    workflowScheduleIntervalInHours *int32
}
// NewLifecycleManagementSettings instantiates a new LifecycleManagementSettings and sets the default values.
func NewLifecycleManagementSettings()(*LifecycleManagementSettings) {
    m := &LifecycleManagementSettings{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateLifecycleManagementSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLifecycleManagementSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleManagementSettings(), nil
}
// GetEmailSettings gets the emailSettings property value. The emailSettings property
// returns a EmailSettingsable when successful
func (m *LifecycleManagementSettings) GetEmailSettings()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.EmailSettingsable) {
    return m.emailSettings
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LifecycleManagementSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateEmailSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSettings(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.EmailSettingsable))
        }
        return nil
    }
    res["workflowScheduleIntervalInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowScheduleIntervalInHours(val)
        }
        return nil
    }
    return res
}
// GetWorkflowScheduleIntervalInHours gets the workflowScheduleIntervalInHours property value. The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
// returns a *int32 when successful
func (m *LifecycleManagementSettings) GetWorkflowScheduleIntervalInHours()(*int32) {
    return m.workflowScheduleIntervalInHours
}
// Serialize serializes information the current object
func (m *LifecycleManagementSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("emailSettings", m.GetEmailSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workflowScheduleIntervalInHours", m.GetWorkflowScheduleIntervalInHours())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailSettings sets the emailSettings property value. The emailSettings property
func (m *LifecycleManagementSettings) SetEmailSettings(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.EmailSettingsable)() {
    m.emailSettings = value
}
// SetWorkflowScheduleIntervalInHours sets the workflowScheduleIntervalInHours property value. The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
func (m *LifecycleManagementSettings) SetWorkflowScheduleIntervalInHours(value *int32)() {
    m.workflowScheduleIntervalInHours = value
}
type LifecycleManagementSettingsable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailSettings()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.EmailSettingsable)
    GetWorkflowScheduleIntervalInHours()(*int32)
    SetEmailSettings(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.EmailSettingsable)()
    SetWorkflowScheduleIntervalInHours(value *int32)()
}
