package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtectionUnitBase struct {
    Entity
    // The identity of person who created the protection unit.
    createdBy IdentitySetable
    // The time of creation of the protection unit.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Contains error details if an error occurred while creating a protection unit.
    error PublicErrorable
    // The identity of person who last modified the protection unit.
    lastModifiedBy IdentitySetable
    // Timestamp of the last modification of this protection unit.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the protection policy based on which protection unit was created.
    policyId *string
    // The status of the protection unit. The possible values are: protectRequested, protected, unprotectRequested, unprotected, removeRequested, unknownFutureValue.
    status *ProtectionUnitStatus
}
// NewProtectionUnitBase instantiates a new ProtectionUnitBase and sets the default values.
func NewProtectionUnitBase()(*ProtectionUnitBase) {
    m := &ProtectionUnitBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProtectionUnitBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectionUnitBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.driveProtectionUnit":
                        return NewDriveProtectionUnit(), nil
                    case "#microsoft.graph.mailboxProtectionUnit":
                        return NewMailboxProtectionUnit(), nil
                    case "#microsoft.graph.siteProtectionUnit":
                        return NewSiteProtectionUnit(), nil
                }
            }
        }
    }
    return NewProtectionUnitBase(), nil
}
// GetCreatedBy gets the createdBy property value. The identity of person who created the protection unit.
// returns a IdentitySetable when successful
func (m *ProtectionUnitBase) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The time of creation of the protection unit.
// returns a *Time when successful
func (m *ProtectionUnitBase) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetError gets the error property value. Contains error details if an error occurred while creating a protection unit.
// returns a PublicErrorable when successful
func (m *ProtectionUnitBase) GetError()(PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectionUnitBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["policyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtectionUnitStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ProtectionUnitStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The identity of person who last modified the protection unit.
// returns a IdentitySetable when successful
func (m *ProtectionUnitBase) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Timestamp of the last modification of this protection unit.
// returns a *Time when successful
func (m *ProtectionUnitBase) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPolicyId gets the policyId property value. The unique identifier of the protection policy based on which protection unit was created.
// returns a *string when successful
func (m *ProtectionUnitBase) GetPolicyId()(*string) {
    return m.policyId
}
// GetStatus gets the status property value. The status of the protection unit. The possible values are: protectRequested, protected, unprotectRequested, unprotected, removeRequested, unknownFutureValue.
// returns a *ProtectionUnitStatus when successful
func (m *ProtectionUnitBase) GetStatus()(*ProtectionUnitStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ProtectionUnitBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("policyId", m.GetPolicyId())
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
// SetCreatedBy sets the createdBy property value. The identity of person who created the protection unit.
func (m *ProtectionUnitBase) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The time of creation of the protection unit.
func (m *ProtectionUnitBase) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetError sets the error property value. Contains error details if an error occurred while creating a protection unit.
func (m *ProtectionUnitBase) SetError(value PublicErrorable)() {
    m.error = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The identity of person who last modified the protection unit.
func (m *ProtectionUnitBase) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Timestamp of the last modification of this protection unit.
func (m *ProtectionUnitBase) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPolicyId sets the policyId property value. The unique identifier of the protection policy based on which protection unit was created.
func (m *ProtectionUnitBase) SetPolicyId(value *string)() {
    m.policyId = value
}
// SetStatus sets the status property value. The status of the protection unit. The possible values are: protectRequested, protected, unprotectRequested, unprotected, removeRequested, unknownFutureValue.
func (m *ProtectionUnitBase) SetStatus(value *ProtectionUnitStatus)() {
    m.status = value
}
type ProtectionUnitBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetError()(PublicErrorable)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPolicyId()(*string)
    GetStatus()(*ProtectionUnitStatus)
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetError(value PublicErrorable)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPolicyId(value *string)()
    SetStatus(value *ProtectionUnitStatus)()
}
