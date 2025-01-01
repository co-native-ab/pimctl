package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageAssignmentRequestWorkflowExtension struct {
    CustomCalloutExtension
    // The callback configuration for a custom extension.
    callbackConfiguration CustomExtensionCallbackConfigurationable
    // The userPrincipalName of the user or identity of the subject that created this resource. Read-only.
    createdBy *string
    // When the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The userPrincipalName of the identity that last modified the object.
    lastModifiedBy *string
    // When the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewAccessPackageAssignmentRequestWorkflowExtension instantiates a new AccessPackageAssignmentRequestWorkflowExtension and sets the default values.
func NewAccessPackageAssignmentRequestWorkflowExtension()(*AccessPackageAssignmentRequestWorkflowExtension) {
    m := &AccessPackageAssignmentRequestWorkflowExtension{
        CustomCalloutExtension: *NewCustomCalloutExtension(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageAssignmentRequestWorkflowExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAccessPackageAssignmentRequestWorkflowExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageAssignmentRequestWorkflowExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentRequestWorkflowExtension(), nil
}
// GetCallbackConfiguration gets the callbackConfiguration property value. The callback configuration for a custom extension.
// returns a CustomExtensionCallbackConfigurationable when successful
func (m *AccessPackageAssignmentRequestWorkflowExtension) GetCallbackConfiguration()(CustomExtensionCallbackConfigurationable) {
    return m.callbackConfiguration
}
// GetCreatedBy gets the createdBy property value. The userPrincipalName of the user or identity of the subject that created this resource. Read-only.
// returns a *string when successful
func (m *AccessPackageAssignmentRequestWorkflowExtension) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. When the object was created.
// returns a *Time when successful
func (m *AccessPackageAssignmentRequestWorkflowExtension) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageAssignmentRequestWorkflowExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomCalloutExtension.GetFieldDeserializers()
    res["callbackConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomExtensionCallbackConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackConfiguration(val.(CustomExtensionCallbackConfigurationable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
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
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The userPrincipalName of the identity that last modified the object.
// returns a *string when successful
func (m *AccessPackageAssignmentRequestWorkflowExtension) GetLastModifiedBy()(*string) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the object was last modified.
// returns a *Time when successful
func (m *AccessPackageAssignmentRequestWorkflowExtension) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequestWorkflowExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomCalloutExtension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("callbackConfiguration", m.GetCallbackConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallbackConfiguration sets the callbackConfiguration property value. The callback configuration for a custom extension.
func (m *AccessPackageAssignmentRequestWorkflowExtension) SetCallbackConfiguration(value CustomExtensionCallbackConfigurationable)() {
    m.callbackConfiguration = value
}
// SetCreatedBy sets the createdBy property value. The userPrincipalName of the user or identity of the subject that created this resource. Read-only.
func (m *AccessPackageAssignmentRequestWorkflowExtension) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. When the object was created.
func (m *AccessPackageAssignmentRequestWorkflowExtension) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The userPrincipalName of the identity that last modified the object.
func (m *AccessPackageAssignmentRequestWorkflowExtension) SetLastModifiedBy(value *string)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the object was last modified.
func (m *AccessPackageAssignmentRequestWorkflowExtension) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
type AccessPackageAssignmentRequestWorkflowExtensionable interface {
    CustomCalloutExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallbackConfiguration()(CustomExtensionCallbackConfigurationable)
    GetCreatedBy()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCallbackConfiguration(value CustomExtensionCallbackConfigurationable)()
    SetCreatedBy(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
