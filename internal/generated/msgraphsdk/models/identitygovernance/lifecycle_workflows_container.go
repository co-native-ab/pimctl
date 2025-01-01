package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type LifecycleWorkflowsContainer struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The customTaskExtension instance.
    customTaskExtensions []CustomTaskExtensionable
    // Deleted workflows in your lifecycle workflows instance.
    deletedItems ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DeletedItemContainerable
    // The insight container holding workflow insight summaries for a tenant.
    insights Insightsable
    // The settings property
    settings LifecycleManagementSettingsable
    // The definition of tasks within the lifecycle workflows instance.
    taskDefinitions []TaskDefinitionable
    // The workflows in the lifecycle workflows instance.
    workflows []Workflowable
    // The workflow templates in the lifecycle workflow instance.
    workflowTemplates []WorkflowTemplateable
}
// NewLifecycleWorkflowsContainer instantiates a new LifecycleWorkflowsContainer and sets the default values.
func NewLifecycleWorkflowsContainer()(*LifecycleWorkflowsContainer) {
    m := &LifecycleWorkflowsContainer{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateLifecycleWorkflowsContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLifecycleWorkflowsContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleWorkflowsContainer(), nil
}
// GetCustomTaskExtensions gets the customTaskExtensions property value. The customTaskExtension instance.
// returns a []CustomTaskExtensionable when successful
func (m *LifecycleWorkflowsContainer) GetCustomTaskExtensions()([]CustomTaskExtensionable) {
    return m.customTaskExtensions
}
// GetDeletedItems gets the deletedItems property value. Deleted workflows in your lifecycle workflows instance.
// returns a DeletedItemContainerable when successful
func (m *LifecycleWorkflowsContainer) GetDeletedItems()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DeletedItemContainerable) {
    return m.deletedItems
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LifecycleWorkflowsContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customTaskExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomTaskExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomTaskExtensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomTaskExtensionable)
                }
            }
            m.SetCustomTaskExtensions(res)
        }
        return nil
    }
    res["deletedItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateDeletedItemContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedItems(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DeletedItemContainerable))
        }
        return nil
    }
    res["insights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInsightsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsights(val.(Insightsable))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLifecycleManagementSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(LifecycleManagementSettingsable))
        }
        return nil
    }
    res["taskDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaskDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaskDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TaskDefinitionable)
                }
            }
            m.SetTaskDefinitions(res)
        }
        return nil
    }
    res["workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkflowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Workflowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Workflowable)
                }
            }
            m.SetWorkflows(res)
        }
        return nil
    }
    res["workflowTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkflowTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkflowTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkflowTemplateable)
                }
            }
            m.SetWorkflowTemplates(res)
        }
        return nil
    }
    return res
}
// GetInsights gets the insights property value. The insight container holding workflow insight summaries for a tenant.
// returns a Insightsable when successful
func (m *LifecycleWorkflowsContainer) GetInsights()(Insightsable) {
    return m.insights
}
// GetSettings gets the settings property value. The settings property
// returns a LifecycleManagementSettingsable when successful
func (m *LifecycleWorkflowsContainer) GetSettings()(LifecycleManagementSettingsable) {
    return m.settings
}
// GetTaskDefinitions gets the taskDefinitions property value. The definition of tasks within the lifecycle workflows instance.
// returns a []TaskDefinitionable when successful
func (m *LifecycleWorkflowsContainer) GetTaskDefinitions()([]TaskDefinitionable) {
    return m.taskDefinitions
}
// GetWorkflows gets the workflows property value. The workflows in the lifecycle workflows instance.
// returns a []Workflowable when successful
func (m *LifecycleWorkflowsContainer) GetWorkflows()([]Workflowable) {
    return m.workflows
}
// GetWorkflowTemplates gets the workflowTemplates property value. The workflow templates in the lifecycle workflow instance.
// returns a []WorkflowTemplateable when successful
func (m *LifecycleWorkflowsContainer) GetWorkflowTemplates()([]WorkflowTemplateable) {
    return m.workflowTemplates
}
// Serialize serializes information the current object
func (m *LifecycleWorkflowsContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomTaskExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomTaskExtensions()))
        for i, v := range m.GetCustomTaskExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customTaskExtensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deletedItems", m.GetDeletedItems())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("insights", m.GetInsights())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    if m.GetTaskDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskDefinitions()))
        for i, v := range m.GetTaskDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("taskDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkflows()))
        for i, v := range m.GetWorkflows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("workflows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflowTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkflowTemplates()))
        for i, v := range m.GetWorkflowTemplates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("workflowTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomTaskExtensions sets the customTaskExtensions property value. The customTaskExtension instance.
func (m *LifecycleWorkflowsContainer) SetCustomTaskExtensions(value []CustomTaskExtensionable)() {
    m.customTaskExtensions = value
}
// SetDeletedItems sets the deletedItems property value. Deleted workflows in your lifecycle workflows instance.
func (m *LifecycleWorkflowsContainer) SetDeletedItems(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DeletedItemContainerable)() {
    m.deletedItems = value
}
// SetInsights sets the insights property value. The insight container holding workflow insight summaries for a tenant.
func (m *LifecycleWorkflowsContainer) SetInsights(value Insightsable)() {
    m.insights = value
}
// SetSettings sets the settings property value. The settings property
func (m *LifecycleWorkflowsContainer) SetSettings(value LifecycleManagementSettingsable)() {
    m.settings = value
}
// SetTaskDefinitions sets the taskDefinitions property value. The definition of tasks within the lifecycle workflows instance.
func (m *LifecycleWorkflowsContainer) SetTaskDefinitions(value []TaskDefinitionable)() {
    m.taskDefinitions = value
}
// SetWorkflows sets the workflows property value. The workflows in the lifecycle workflows instance.
func (m *LifecycleWorkflowsContainer) SetWorkflows(value []Workflowable)() {
    m.workflows = value
}
// SetWorkflowTemplates sets the workflowTemplates property value. The workflow templates in the lifecycle workflow instance.
func (m *LifecycleWorkflowsContainer) SetWorkflowTemplates(value []WorkflowTemplateable)() {
    m.workflowTemplates = value
}
type LifecycleWorkflowsContainerable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomTaskExtensions()([]CustomTaskExtensionable)
    GetDeletedItems()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DeletedItemContainerable)
    GetInsights()(Insightsable)
    GetSettings()(LifecycleManagementSettingsable)
    GetTaskDefinitions()([]TaskDefinitionable)
    GetWorkflows()([]Workflowable)
    GetWorkflowTemplates()([]WorkflowTemplateable)
    SetCustomTaskExtensions(value []CustomTaskExtensionable)()
    SetDeletedItems(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DeletedItemContainerable)()
    SetInsights(value Insightsable)()
    SetSettings(value LifecycleManagementSettingsable)()
    SetTaskDefinitions(value []TaskDefinitionable)()
    SetWorkflows(value []Workflowable)()
    SetWorkflowTemplates(value []WorkflowTemplateable)()
}
