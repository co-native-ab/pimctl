package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EducationAssignmentDefaults struct {
    Entity
    // Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen.
    addedStudentAction *EducationAddedStudentAction
    // Optional field to control adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. You must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: studentsOnly. The default value is none.
    addToCalendarAction *EducationAddToCalendarOptions
    // Class-level default value for due time field. Default value is 23:59:00.
    dueTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Default Teams channel to which notifications are sent. Default value is null.
    notificationChannelUrl *string
}
// NewEducationAssignmentDefaults instantiates a new EducationAssignmentDefaults and sets the default values.
func NewEducationAssignmentDefaults()(*EducationAssignmentDefaults) {
    m := &EducationAssignmentDefaults{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationAssignmentDefaultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationAssignmentDefaultsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignmentDefaults(), nil
}
// GetAddedStudentAction gets the addedStudentAction property value. Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen.
// returns a *EducationAddedStudentAction when successful
func (m *EducationAssignmentDefaults) GetAddedStudentAction()(*EducationAddedStudentAction) {
    return m.addedStudentAction
}
// GetAddToCalendarAction gets the addToCalendarAction property value. Optional field to control adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. You must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: studentsOnly. The default value is none.
// returns a *EducationAddToCalendarOptions when successful
func (m *EducationAssignmentDefaults) GetAddToCalendarAction()(*EducationAddToCalendarOptions) {
    return m.addToCalendarAction
}
// GetDueTime gets the dueTime property value. Class-level default value for due time field. Default value is 23:59:00.
// returns a *TimeOnly when successful
func (m *EducationAssignmentDefaults) GetDueTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.dueTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationAssignmentDefaults) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedStudentAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddedStudentAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedStudentAction(val.(*EducationAddedStudentAction))
        }
        return nil
    }
    res["addToCalendarAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddToCalendarOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddToCalendarAction(val.(*EducationAddToCalendarOptions))
        }
        return nil
    }
    res["dueTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueTime(val)
        }
        return nil
    }
    res["notificationChannelUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationChannelUrl(val)
        }
        return nil
    }
    return res
}
// GetNotificationChannelUrl gets the notificationChannelUrl property value. Default Teams channel to which notifications are sent. Default value is null.
// returns a *string when successful
func (m *EducationAssignmentDefaults) GetNotificationChannelUrl()(*string) {
    return m.notificationChannelUrl
}
// Serialize serializes information the current object
func (m *EducationAssignmentDefaults) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddedStudentAction() != nil {
        cast := (*m.GetAddedStudentAction()).String()
        err = writer.WriteStringValue("addedStudentAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddToCalendarAction() != nil {
        cast := (*m.GetAddToCalendarAction()).String()
        err = writer.WriteStringValue("addToCalendarAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("dueTime", m.GetDueTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationChannelUrl", m.GetNotificationChannelUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedStudentAction sets the addedStudentAction property value. Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen.
func (m *EducationAssignmentDefaults) SetAddedStudentAction(value *EducationAddedStudentAction)() {
    m.addedStudentAction = value
}
// SetAddToCalendarAction sets the addToCalendarAction property value. Optional field to control adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. You must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: studentsOnly. The default value is none.
func (m *EducationAssignmentDefaults) SetAddToCalendarAction(value *EducationAddToCalendarOptions)() {
    m.addToCalendarAction = value
}
// SetDueTime sets the dueTime property value. Class-level default value for due time field. Default value is 23:59:00.
func (m *EducationAssignmentDefaults) SetDueTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.dueTime = value
}
// SetNotificationChannelUrl sets the notificationChannelUrl property value. Default Teams channel to which notifications are sent. Default value is null.
func (m *EducationAssignmentDefaults) SetNotificationChannelUrl(value *string)() {
    m.notificationChannelUrl = value
}
type EducationAssignmentDefaultsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedStudentAction()(*EducationAddedStudentAction)
    GetAddToCalendarAction()(*EducationAddToCalendarOptions)
    GetDueTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetNotificationChannelUrl()(*string)
    SetAddedStudentAction(value *EducationAddedStudentAction)()
    SetAddToCalendarAction(value *EducationAddToCalendarOptions)()
    SetDueTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetNotificationChannelUrl(value *string)()
}
