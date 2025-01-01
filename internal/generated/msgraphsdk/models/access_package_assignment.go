package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageAssignment struct {
    Entity
    // Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
    accessPackage AccessPackageable
    // Read-only. Supports $filter (eq) on the id property and $expand query parameters.
    assignmentPolicy AccessPackageAssignmentPolicyable
    // Information about all the custom extension calls that were made during the access package assignment workflow.
    customExtensionCalloutInstances []CustomExtensionCalloutInstanceable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    expiredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // When the access assignment is to be in place. Read-only.
    schedule EntitlementManagementScheduleable
    // The state of the access package assignment. The possible values are: delivering, partiallyDelivered, delivered, expired, deliveryFailed, unknownFutureValue. Read-only. Supports $filter (eq).
    state *AccessPackageAssignmentState
    // More information about the assignment lifecycle. Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered. Read-only.
    status *string
    // The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
    target AccessPackageSubjectable
}
// NewAccessPackageAssignment instantiates a new AccessPackageAssignment and sets the default values.
func NewAccessPackageAssignment()(*AccessPackageAssignment) {
    m := &AccessPackageAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignment(), nil
}
// GetAccessPackage gets the accessPackage property value. Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
// returns a AccessPackageable when successful
func (m *AccessPackageAssignment) GetAccessPackage()(AccessPackageable) {
    return m.accessPackage
}
// GetAssignmentPolicy gets the assignmentPolicy property value. Read-only. Supports $filter (eq) on the id property and $expand query parameters.
// returns a AccessPackageAssignmentPolicyable when successful
func (m *AccessPackageAssignment) GetAssignmentPolicy()(AccessPackageAssignmentPolicyable) {
    return m.assignmentPolicy
}
// GetCustomExtensionCalloutInstances gets the customExtensionCalloutInstances property value. Information about all the custom extension calls that were made during the access package assignment workflow.
// returns a []CustomExtensionCalloutInstanceable when successful
func (m *AccessPackageAssignment) GetCustomExtensionCalloutInstances()([]CustomExtensionCalloutInstanceable) {
    return m.customExtensionCalloutInstances
}
// GetExpiredDateTime gets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *AccessPackageAssignment) GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiredDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(AccessPackageable))
        }
        return nil
    }
    res["assignmentPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentPolicy(val.(AccessPackageAssignmentPolicyable))
        }
        return nil
    }
    res["customExtensionCalloutInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomExtensionCalloutInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomExtensionCalloutInstanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomExtensionCalloutInstanceable)
                }
            }
            m.SetCustomExtensionCalloutInstances(res)
        }
        return nil
    }
    res["expiredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiredDateTime(val)
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntitlementManagementScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(EntitlementManagementScheduleable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageAssignmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AccessPackageAssignmentState))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(AccessPackageSubjectable))
        }
        return nil
    }
    return res
}
// GetSchedule gets the schedule property value. When the access assignment is to be in place. Read-only.
// returns a EntitlementManagementScheduleable when successful
func (m *AccessPackageAssignment) GetSchedule()(EntitlementManagementScheduleable) {
    return m.schedule
}
// GetState gets the state property value. The state of the access package assignment. The possible values are: delivering, partiallyDelivered, delivered, expired, deliveryFailed, unknownFutureValue. Read-only. Supports $filter (eq).
// returns a *AccessPackageAssignmentState when successful
func (m *AccessPackageAssignment) GetState()(*AccessPackageAssignmentState) {
    return m.state
}
// GetStatus gets the status property value. More information about the assignment lifecycle. Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered. Read-only.
// returns a *string when successful
func (m *AccessPackageAssignment) GetStatus()(*string) {
    return m.status
}
// GetTarget gets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
// returns a AccessPackageSubjectable when successful
func (m *AccessPackageAssignment) GetTarget()(AccessPackageSubjectable) {
    return m.target
}
// Serialize serializes information the current object
func (m *AccessPackageAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackage", m.GetAccessPackage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignmentPolicy", m.GetAssignmentPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetCustomExtensionCalloutInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomExtensionCalloutInstances()))
        for i, v := range m.GetCustomExtensionCalloutInstances() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customExtensionCalloutInstances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expiredDateTime", m.GetExpiredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackage sets the accessPackage property value. Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageAssignment) SetAccessPackage(value AccessPackageable)() {
    m.accessPackage = value
}
// SetAssignmentPolicy sets the assignmentPolicy property value. Read-only. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageAssignment) SetAssignmentPolicy(value AccessPackageAssignmentPolicyable)() {
    m.assignmentPolicy = value
}
// SetCustomExtensionCalloutInstances sets the customExtensionCalloutInstances property value. Information about all the custom extension calls that were made during the access package assignment workflow.
func (m *AccessPackageAssignment) SetCustomExtensionCalloutInstances(value []CustomExtensionCalloutInstanceable)() {
    m.customExtensionCalloutInstances = value
}
// SetExpiredDateTime sets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignment) SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiredDateTime = value
}
// SetSchedule sets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) SetSchedule(value EntitlementManagementScheduleable)() {
    m.schedule = value
}
// SetState sets the state property value. The state of the access package assignment. The possible values are: delivering, partiallyDelivered, delivered, expired, deliveryFailed, unknownFutureValue. Read-only. Supports $filter (eq).
func (m *AccessPackageAssignment) SetState(value *AccessPackageAssignmentState)() {
    m.state = value
}
// SetStatus sets the status property value. More information about the assignment lifecycle. Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered. Read-only.
func (m *AccessPackageAssignment) SetStatus(value *string)() {
    m.status = value
}
// SetTarget sets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (m *AccessPackageAssignment) SetTarget(value AccessPackageSubjectable)() {
    m.target = value
}
type AccessPackageAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAssignmentPolicy()(AccessPackageAssignmentPolicyable)
    GetCustomExtensionCalloutInstances()([]CustomExtensionCalloutInstanceable)
    GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSchedule()(EntitlementManagementScheduleable)
    GetState()(*AccessPackageAssignmentState)
    GetStatus()(*string)
    GetTarget()(AccessPackageSubjectable)
    SetAccessPackage(value AccessPackageable)()
    SetAssignmentPolicy(value AccessPackageAssignmentPolicyable)()
    SetCustomExtensionCalloutInstances(value []CustomExtensionCalloutInstanceable)()
    SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSchedule(value EntitlementManagementScheduleable)()
    SetState(value *AccessPackageAssignmentState)()
    SetStatus(value *string)()
    SetTarget(value AccessPackageSubjectable)()
}
