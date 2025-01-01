package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type TriggerAndScopeBasedConditions struct {
    WorkflowExecutionConditions
    // Defines who the workflow runs for.
    scope ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.SubjectSetable
    // What triggers a workflow to run.
    trigger WorkflowExecutionTriggerable
}
// NewTriggerAndScopeBasedConditions instantiates a new TriggerAndScopeBasedConditions and sets the default values.
func NewTriggerAndScopeBasedConditions()(*TriggerAndScopeBasedConditions) {
    m := &TriggerAndScopeBasedConditions{
        WorkflowExecutionConditions: *NewWorkflowExecutionConditions(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.triggerAndScopeBasedConditions"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTriggerAndScopeBasedConditionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTriggerAndScopeBasedConditionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTriggerAndScopeBasedConditions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TriggerAndScopeBasedConditions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowExecutionConditions.GetFieldDeserializers()
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.SubjectSetable))
        }
        return nil
    }
    res["trigger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowExecutionTriggerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrigger(val.(WorkflowExecutionTriggerable))
        }
        return nil
    }
    return res
}
// GetScope gets the scope property value. Defines who the workflow runs for.
// returns a SubjectSetable when successful
func (m *TriggerAndScopeBasedConditions) GetScope()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.SubjectSetable) {
    return m.scope
}
// GetTrigger gets the trigger property value. What triggers a workflow to run.
// returns a WorkflowExecutionTriggerable when successful
func (m *TriggerAndScopeBasedConditions) GetTrigger()(WorkflowExecutionTriggerable) {
    return m.trigger
}
// Serialize serializes information the current object
func (m *TriggerAndScopeBasedConditions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowExecutionConditions.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("trigger", m.GetTrigger())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetScope sets the scope property value. Defines who the workflow runs for.
func (m *TriggerAndScopeBasedConditions) SetScope(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.SubjectSetable)() {
    m.scope = value
}
// SetTrigger sets the trigger property value. What triggers a workflow to run.
func (m *TriggerAndScopeBasedConditions) SetTrigger(value WorkflowExecutionTriggerable)() {
    m.trigger = value
}
type TriggerAndScopeBasedConditionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WorkflowExecutionConditionsable
    GetScope()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.SubjectSetable)
    GetTrigger()(WorkflowExecutionTriggerable)
    SetScope(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.SubjectSetable)()
    SetTrigger(value WorkflowExecutionTriggerable)()
}
