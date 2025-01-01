package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtectionPolicyBase struct {
    Entity
    // The identity of person who created the policy.
    createdBy IdentitySetable
    // The time of creation of the policy.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the policy to be created.
    displayName *string
    // The identity of the person who last modified the policy.
    lastModifiedBy IdentitySetable
    // The timestamp of the last modification of the policy.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Contains the retention setting details for the policy.
    retentionSettings []RetentionSettingable
    // The aggregated status of the protection units associated with the policy. The possible values are: inactive, activeWithErrors, updating, active, unknownFutureValue.
    status *ProtectionPolicyStatus
}
// NewProtectionPolicyBase instantiates a new ProtectionPolicyBase and sets the default values.
func NewProtectionPolicyBase()(*ProtectionPolicyBase) {
    m := &ProtectionPolicyBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProtectionPolicyBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectionPolicyBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.exchangeProtectionPolicy":
                        return NewExchangeProtectionPolicy(), nil
                    case "#microsoft.graph.oneDriveForBusinessProtectionPolicy":
                        return NewOneDriveForBusinessProtectionPolicy(), nil
                    case "#microsoft.graph.sharePointProtectionPolicy":
                        return NewSharePointProtectionPolicy(), nil
                }
            }
        }
    }
    return NewProtectionPolicyBase(), nil
}
// GetCreatedBy gets the createdBy property value. The identity of person who created the policy.
// returns a IdentitySetable when successful
func (m *ProtectionPolicyBase) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The time of creation of the policy.
// returns a *Time when successful
func (m *ProtectionPolicyBase) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The name of the policy to be created.
// returns a *string when successful
func (m *ProtectionPolicyBase) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectionPolicyBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["retentionSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRetentionSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RetentionSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RetentionSettingable)
                }
            }
            m.SetRetentionSettings(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtectionPolicyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ProtectionPolicyStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The identity of the person who last modified the policy.
// returns a IdentitySetable when successful
func (m *ProtectionPolicyBase) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The timestamp of the last modification of the policy.
// returns a *Time when successful
func (m *ProtectionPolicyBase) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRetentionSettings gets the retentionSettings property value. Contains the retention setting details for the policy.
// returns a []RetentionSettingable when successful
func (m *ProtectionPolicyBase) GetRetentionSettings()([]RetentionSettingable) {
    return m.retentionSettings
}
// GetStatus gets the status property value. The aggregated status of the protection units associated with the policy. The possible values are: inactive, activeWithErrors, updating, active, unknownFutureValue.
// returns a *ProtectionPolicyStatus when successful
func (m *ProtectionPolicyBase) GetStatus()(*ProtectionPolicyStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ProtectionPolicyBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetRetentionSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRetentionSettings()))
        for i, v := range m.GetRetentionSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("retentionSettings", cast)
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
// SetCreatedBy sets the createdBy property value. The identity of person who created the policy.
func (m *ProtectionPolicyBase) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The time of creation of the policy.
func (m *ProtectionPolicyBase) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The name of the policy to be created.
func (m *ProtectionPolicyBase) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The identity of the person who last modified the policy.
func (m *ProtectionPolicyBase) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The timestamp of the last modification of the policy.
func (m *ProtectionPolicyBase) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRetentionSettings sets the retentionSettings property value. Contains the retention setting details for the policy.
func (m *ProtectionPolicyBase) SetRetentionSettings(value []RetentionSettingable)() {
    m.retentionSettings = value
}
// SetStatus sets the status property value. The aggregated status of the protection units associated with the policy. The possible values are: inactive, activeWithErrors, updating, active, unknownFutureValue.
func (m *ProtectionPolicyBase) SetStatus(value *ProtectionPolicyStatus)() {
    m.status = value
}
type ProtectionPolicyBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRetentionSettings()([]RetentionSettingable)
    GetStatus()(*ProtectionPolicyStatus)
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRetentionSettings(value []RetentionSettingable)()
    SetStatus(value *ProtectionPolicyStatus)()
}
