package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type UserProcessingResult struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The date time that the workflow execution for a user completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of tasks that failed in the workflow execution.
    failedTasksCount *int32
    // The processingStatus property
    processingStatus *LifecycleWorkflowProcessingStatus
    // The date time that the workflow is scheduled to be executed for a user.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    scheduledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date time that the workflow execution started. Value is null if the workflow execution has not started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject property
    subject ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable
    // The associated individual task execution.
    taskProcessingResults []TaskProcessingResultable
    // The total number of tasks that in the workflow execution.
    totalTasksCount *int32
    // The total number of unprocessed tasks for the workflow.
    totalUnprocessedTasksCount *int32
    // The workflowExecutionType property
    workflowExecutionType *WorkflowExecutionType
    // The version of the workflow that was executed.
    workflowVersion *int32
}
// NewUserProcessingResult instantiates a new UserProcessingResult and sets the default values.
func NewUserProcessingResult()(*UserProcessingResult) {
    m := &UserProcessingResult{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateUserProcessingResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserProcessingResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserProcessingResult(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. The date time that the workflow execution for a user completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
// returns a *Time when successful
func (m *UserProcessingResult) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetFailedTasksCount gets the failedTasksCount property value. The number of tasks that failed in the workflow execution.
// returns a *int32 when successful
func (m *UserProcessingResult) GetFailedTasksCount()(*int32) {
    return m.failedTasksCount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserProcessingResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["failedTasksCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedTasksCount(val)
        }
        return nil
    }
    res["processingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLifecycleWorkflowProcessingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingStatus(val.(*LifecycleWorkflowProcessingStatus))
        }
        return nil
    }
    res["scheduledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledDateTime(val)
        }
        return nil
    }
    res["startedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartedDateTime(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable))
        }
        return nil
    }
    res["taskProcessingResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaskProcessingResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaskProcessingResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TaskProcessingResultable)
                }
            }
            m.SetTaskProcessingResults(res)
        }
        return nil
    }
    res["totalTasksCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTasksCount(val)
        }
        return nil
    }
    res["totalUnprocessedTasksCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUnprocessedTasksCount(val)
        }
        return nil
    }
    res["workflowExecutionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkflowExecutionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowExecutionType(val.(*WorkflowExecutionType))
        }
        return nil
    }
    res["workflowVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowVersion(val)
        }
        return nil
    }
    return res
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
// returns a *LifecycleWorkflowProcessingStatus when successful
func (m *UserProcessingResult) GetProcessingStatus()(*LifecycleWorkflowProcessingStatus) {
    return m.processingStatus
}
// GetScheduledDateTime gets the scheduledDateTime property value. The date time that the workflow is scheduled to be executed for a user.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
// returns a *Time when successful
func (m *UserProcessingResult) GetScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.scheduledDateTime
}
// GetStartedDateTime gets the startedDateTime property value. The date time that the workflow execution started. Value is null if the workflow execution has not started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
// returns a *Time when successful
func (m *UserProcessingResult) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetSubject gets the subject property value. The subject property
// returns a Userable when successful
func (m *UserProcessingResult) GetSubject()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable) {
    return m.subject
}
// GetTaskProcessingResults gets the taskProcessingResults property value. The associated individual task execution.
// returns a []TaskProcessingResultable when successful
func (m *UserProcessingResult) GetTaskProcessingResults()([]TaskProcessingResultable) {
    return m.taskProcessingResults
}
// GetTotalTasksCount gets the totalTasksCount property value. The total number of tasks that in the workflow execution.
// returns a *int32 when successful
func (m *UserProcessingResult) GetTotalTasksCount()(*int32) {
    return m.totalTasksCount
}
// GetTotalUnprocessedTasksCount gets the totalUnprocessedTasksCount property value. The total number of unprocessed tasks for the workflow.
// returns a *int32 when successful
func (m *UserProcessingResult) GetTotalUnprocessedTasksCount()(*int32) {
    return m.totalUnprocessedTasksCount
}
// GetWorkflowExecutionType gets the workflowExecutionType property value. The workflowExecutionType property
// returns a *WorkflowExecutionType when successful
func (m *UserProcessingResult) GetWorkflowExecutionType()(*WorkflowExecutionType) {
    return m.workflowExecutionType
}
// GetWorkflowVersion gets the workflowVersion property value. The version of the workflow that was executed.
// returns a *int32 when successful
func (m *UserProcessingResult) GetWorkflowVersion()(*int32) {
    return m.workflowVersion
}
// Serialize serializes information the current object
func (m *UserProcessingResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedTasksCount", m.GetFailedTasksCount())
        if err != nil {
            return err
        }
    }
    if m.GetProcessingStatus() != nil {
        cast := (*m.GetProcessingStatus()).String()
        err = writer.WriteStringValue("processingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("scheduledDateTime", m.GetScheduledDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startedDateTime", m.GetStartedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    if m.GetTaskProcessingResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskProcessingResults()))
        for i, v := range m.GetTaskProcessingResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("taskProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalTasksCount", m.GetTotalTasksCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUnprocessedTasksCount", m.GetTotalUnprocessedTasksCount())
        if err != nil {
            return err
        }
    }
    if m.GetWorkflowExecutionType() != nil {
        cast := (*m.GetWorkflowExecutionType()).String()
        err = writer.WriteStringValue("workflowExecutionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workflowVersion", m.GetWorkflowVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. The date time that the workflow execution for a user completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *UserProcessingResult) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetFailedTasksCount sets the failedTasksCount property value. The number of tasks that failed in the workflow execution.
func (m *UserProcessingResult) SetFailedTasksCount(value *int32)() {
    m.failedTasksCount = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *UserProcessingResult) SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)() {
    m.processingStatus = value
}
// SetScheduledDateTime sets the scheduledDateTime property value. The date time that the workflow is scheduled to be executed for a user.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *UserProcessingResult) SetScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.scheduledDateTime = value
}
// SetStartedDateTime sets the startedDateTime property value. The date time that the workflow execution started. Value is null if the workflow execution has not started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *UserProcessingResult) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetSubject sets the subject property value. The subject property
func (m *UserProcessingResult) SetSubject(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)() {
    m.subject = value
}
// SetTaskProcessingResults sets the taskProcessingResults property value. The associated individual task execution.
func (m *UserProcessingResult) SetTaskProcessingResults(value []TaskProcessingResultable)() {
    m.taskProcessingResults = value
}
// SetTotalTasksCount sets the totalTasksCount property value. The total number of tasks that in the workflow execution.
func (m *UserProcessingResult) SetTotalTasksCount(value *int32)() {
    m.totalTasksCount = value
}
// SetTotalUnprocessedTasksCount sets the totalUnprocessedTasksCount property value. The total number of unprocessed tasks for the workflow.
func (m *UserProcessingResult) SetTotalUnprocessedTasksCount(value *int32)() {
    m.totalUnprocessedTasksCount = value
}
// SetWorkflowExecutionType sets the workflowExecutionType property value. The workflowExecutionType property
func (m *UserProcessingResult) SetWorkflowExecutionType(value *WorkflowExecutionType)() {
    m.workflowExecutionType = value
}
// SetWorkflowVersion sets the workflowVersion property value. The version of the workflow that was executed.
func (m *UserProcessingResult) SetWorkflowVersion(value *int32)() {
    m.workflowVersion = value
}
type UserProcessingResultable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailedTasksCount()(*int32)
    GetProcessingStatus()(*LifecycleWorkflowProcessingStatus)
    GetScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)
    GetTaskProcessingResults()([]TaskProcessingResultable)
    GetTotalTasksCount()(*int32)
    GetTotalUnprocessedTasksCount()(*int32)
    GetWorkflowExecutionType()(*WorkflowExecutionType)
    GetWorkflowVersion()(*int32)
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailedTasksCount(value *int32)()
    SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)()
    SetScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)()
    SetTaskProcessingResults(value []TaskProcessingResultable)()
    SetTotalTasksCount(value *int32)()
    SetTotalUnprocessedTasksCount(value *int32)()
    SetWorkflowExecutionType(value *WorkflowExecutionType)()
    SetWorkflowVersion(value *int32)()
}
