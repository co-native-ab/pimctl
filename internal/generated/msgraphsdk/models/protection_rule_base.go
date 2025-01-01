package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtectionRuleBase struct {
    Entity
    // The identity of person who created the rule.
    createdBy IdentitySetable
    // The time of creation of the rule.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Contains error details if an operation on a rule fails.
    error PublicErrorable
    // true indicates that the protection rule is dynamic; false that it's static. Currently, only static rules are supported.
    isAutoApplyEnabled *bool
    // The identity of the person who last modified the rule.
    lastModifiedBy IdentitySetable
    // Timestamp of the last modification made to the rule.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status of the protection rule. The possible values are: draft, active, completed, completedWithErrors, unknownFutureValue. The draft member is currently unsupported.
    status *ProtectionRuleStatus
}
// NewProtectionRuleBase instantiates a new ProtectionRuleBase and sets the default values.
func NewProtectionRuleBase()(*ProtectionRuleBase) {
    m := &ProtectionRuleBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProtectionRuleBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectionRuleBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.driveProtectionRule":
                        return NewDriveProtectionRule(), nil
                    case "#microsoft.graph.mailboxProtectionRule":
                        return NewMailboxProtectionRule(), nil
                    case "#microsoft.graph.siteProtectionRule":
                        return NewSiteProtectionRule(), nil
                }
            }
        }
    }
    return NewProtectionRuleBase(), nil
}
// GetCreatedBy gets the createdBy property value. The identity of person who created the rule.
// returns a IdentitySetable when successful
func (m *ProtectionRuleBase) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The time of creation of the rule.
// returns a *Time when successful
func (m *ProtectionRuleBase) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetError gets the error property value. Contains error details if an operation on a rule fails.
// returns a PublicErrorable when successful
func (m *ProtectionRuleBase) GetError()(PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectionRuleBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
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
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(PublicErrorable))
        }
        return nil
    }
    res["isAutoApplyEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutoApplyEnabled(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtectionRuleStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ProtectionRuleStatus))
        }
        return nil
    }
    return res
}
// GetIsAutoApplyEnabled gets the isAutoApplyEnabled property value. true indicates that the protection rule is dynamic; false that it's static. Currently, only static rules are supported.
// returns a *bool when successful
func (m *ProtectionRuleBase) GetIsAutoApplyEnabled()(*bool) {
    return m.isAutoApplyEnabled
}
// GetLastModifiedBy gets the lastModifiedBy property value. The identity of the person who last modified the rule.
// returns a IdentitySetable when successful
func (m *ProtectionRuleBase) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Timestamp of the last modification made to the rule.
// returns a *Time when successful
func (m *ProtectionRuleBase) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetStatus gets the status property value. The status of the protection rule. The possible values are: draft, active, completed, completedWithErrors, unknownFutureValue. The draft member is currently unsupported.
// returns a *ProtectionRuleStatus when successful
func (m *ProtectionRuleBase) GetStatus()(*ProtectionRuleStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ProtectionRuleBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAutoApplyEnabled", m.GetIsAutoApplyEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The identity of person who created the rule.
func (m *ProtectionRuleBase) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The time of creation of the rule.
func (m *ProtectionRuleBase) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetError sets the error property value. Contains error details if an operation on a rule fails.
func (m *ProtectionRuleBase) SetError(value PublicErrorable)() {
    m.error = value
}
// SetIsAutoApplyEnabled sets the isAutoApplyEnabled property value. true indicates that the protection rule is dynamic; false that it's static. Currently, only static rules are supported.
func (m *ProtectionRuleBase) SetIsAutoApplyEnabled(value *bool)() {
    m.isAutoApplyEnabled = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The identity of the person who last modified the rule.
func (m *ProtectionRuleBase) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Timestamp of the last modification made to the rule.
func (m *ProtectionRuleBase) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetStatus sets the status property value. The status of the protection rule. The possible values are: draft, active, completed, completedWithErrors, unknownFutureValue. The draft member is currently unsupported.
func (m *ProtectionRuleBase) SetStatus(value *ProtectionRuleStatus)() {
    m.status = value
}
type ProtectionRuleBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetError()(PublicErrorable)
    GetIsAutoApplyEnabled()(*bool)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*ProtectionRuleStatus)
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetError(value PublicErrorable)()
    SetIsAutoApplyEnabled(value *bool)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *ProtectionRuleStatus)()
}
