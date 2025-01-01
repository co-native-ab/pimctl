package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type TaskProcessingResult struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The date time when taskProcessingResult execution ended. Value is null if task execution is still in progress.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date time when the taskProcessingResult was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Describes why the taskProcessingResult has failed.
    failureReason *string
    // The processingStatus property
    processingStatus *LifecycleWorkflowProcessingStatus
    // The date time when taskProcessingResult execution started. Value is null if task execution has not yet started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject property
    subject ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable
    // The task property
    task Taskable
}
// NewTaskProcessingResult instantiates a new TaskProcessingResult and sets the default values.
func NewTaskProcessingResult()(*TaskProcessingResult) {
    m := &TaskProcessingResult{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateTaskProcessingResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTaskProcessingResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskProcessingResult(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. The date time when taskProcessingResult execution ended. Value is null if task execution is still in progress.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
// returns a *Time when successful
func (m *TaskProcessingResult) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCreatedDateTime gets the createdDateTime property value. The date time when the taskProcessingResult was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
// returns a *Time when successful
func (m *TaskProcessingResult) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFailureReason gets the failureReason property value. Describes why the taskProcessingResult has failed.
// returns a *string when successful
func (m *TaskProcessingResult) GetFailureReason()(*string) {
    return m.failureReason
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TaskProcessingResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["failureReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureReason(val)
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
    res["task"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTask(val.(Taskable))
        }
        return nil
    }
    return res
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
// returns a *LifecycleWorkflowProcessingStatus when successful
func (m *TaskProcessingResult) GetProcessingStatus()(*LifecycleWorkflowProcessingStatus) {
    return m.processingStatus
}
// GetStartedDateTime gets the startedDateTime property value. The date time when taskProcessingResult execution started. Value is null if task execution has not yet started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
// returns a *Time when successful
func (m *TaskProcessingResult) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetSubject gets the subject property value. The subject property
// returns a Userable when successful
func (m *TaskProcessingResult) GetSubject()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable) {
    return m.subject
}
// GetTask gets the task property value. The task property
// returns a Taskable when successful
func (m *TaskProcessingResult) GetTask()(Taskable) {
    return m.task
}
// Serialize serializes information the current object
func (m *TaskProcessingResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("failureReason", m.GetFailureReason())
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
    {
        err = writer.WriteObjectValue("task", m.GetTask())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. The date time when taskProcessingResult execution ended. Value is null if task execution is still in progress.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskProcessingResult) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date time when the taskProcessingResult was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskProcessingResult) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetFailureReason sets the failureReason property value. Describes why the taskProcessingResult has failed.
func (m *TaskProcessingResult) SetFailureReason(value *string)() {
    m.failureReason = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *TaskProcessingResult) SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)() {
    m.processingStatus = value
}
// SetStartedDateTime sets the startedDateTime property value. The date time when taskProcessingResult execution started. Value is null if task execution has not yet started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskProcessingResult) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetSubject sets the subject property value. The subject property
func (m *TaskProcessingResult) SetSubject(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)() {
    m.subject = value
}
// SetTask sets the task property value. The task property
func (m *TaskProcessingResult) SetTask(value Taskable)() {
    m.task = value
}
type TaskProcessingResultable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailureReason()(*string)
    GetProcessingStatus()(*LifecycleWorkflowProcessingStatus)
    GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)
    GetTask()(Taskable)
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailureReason(value *string)()
    SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)()
    SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)()
    SetTask(value Taskable)()
}
