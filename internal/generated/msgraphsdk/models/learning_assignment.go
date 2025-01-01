package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LearningAssignment struct {
    LearningCourseActivity
    // Assigned date for the course activity. Optional.
    assignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user ID of the assigner. Optional.
    assignerUserId *string
    // The assignmentType property
    assignmentType *AssignmentType
    // Due date for the course activity. Optional.
    dueDateTime DateTimeTimeZoneable
    // Notes for the course activity. Optional.
    notes ItemBodyable
}
// NewLearningAssignment instantiates a new LearningAssignment and sets the default values.
func NewLearningAssignment()(*LearningAssignment) {
    m := &LearningAssignment{
        LearningCourseActivity: *NewLearningCourseActivity(),
    }
    return m
}
// CreateLearningAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLearningAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLearningAssignment(), nil
}
// GetAssignedDateTime gets the assignedDateTime property value. Assigned date for the course activity. Optional.
// returns a *Time when successful
func (m *LearningAssignment) GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.assignedDateTime
}
// GetAssignerUserId gets the assignerUserId property value. The user ID of the assigner. Optional.
// returns a *string when successful
func (m *LearningAssignment) GetAssignerUserId()(*string) {
    return m.assignerUserId
}
// GetAssignmentType gets the assignmentType property value. The assignmentType property
// returns a *AssignmentType when successful
func (m *LearningAssignment) GetAssignmentType()(*AssignmentType) {
    return m.assignmentType
}
// GetDueDateTime gets the dueDateTime property value. Due date for the course activity. Optional.
// returns a DateTimeTimeZoneable when successful
func (m *LearningAssignment) GetDueDateTime()(DateTimeTimeZoneable) {
    return m.dueDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LearningAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LearningCourseActivity.GetFieldDeserializers()
    res["assignedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedDateTime(val)
        }
        return nil
    }
    res["assignerUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignerUserId(val)
        }
        return nil
    }
    res["assignmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentType(val.(*AssignmentType))
        }
        return nil
    }
    res["dueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(ItemBodyable))
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. Notes for the course activity. Optional.
// returns a ItemBodyable when successful
func (m *LearningAssignment) GetNotes()(ItemBodyable) {
    return m.notes
}
// Serialize serializes information the current object
func (m *LearningAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LearningCourseActivity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("assignedDateTime", m.GetAssignedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignerUserId", m.GetAssignerUserId())
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentType() != nil {
        cast := (*m.GetAssignmentType()).String()
        err = writer.WriteStringValue("assignmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedDateTime sets the assignedDateTime property value. Assigned date for the course activity. Optional.
func (m *LearningAssignment) SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignedDateTime = value
}
// SetAssignerUserId sets the assignerUserId property value. The user ID of the assigner. Optional.
func (m *LearningAssignment) SetAssignerUserId(value *string)() {
    m.assignerUserId = value
}
// SetAssignmentType sets the assignmentType property value. The assignmentType property
func (m *LearningAssignment) SetAssignmentType(value *AssignmentType)() {
    m.assignmentType = value
}
// SetDueDateTime sets the dueDateTime property value. Due date for the course activity. Optional.
func (m *LearningAssignment) SetDueDateTime(value DateTimeTimeZoneable)() {
    m.dueDateTime = value
}
// SetNotes sets the notes property value. Notes for the course activity. Optional.
func (m *LearningAssignment) SetNotes(value ItemBodyable)() {
    m.notes = value
}
type LearningAssignmentable interface {
    LearningCourseActivityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAssignerUserId()(*string)
    GetAssignmentType()(*AssignmentType)
    GetDueDateTime()(DateTimeTimeZoneable)
    GetNotes()(ItemBodyable)
    SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAssignerUserId(value *string)()
    SetAssignmentType(value *AssignmentType)()
    SetDueDateTime(value DateTimeTimeZoneable)()
    SetNotes(value ItemBodyable)()
}
