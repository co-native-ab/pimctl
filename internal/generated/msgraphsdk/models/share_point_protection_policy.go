package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SharePointProtectionPolicy struct {
    ProtectionPolicyBase
    // The rules associated with the SharePoint Protection policy.
    siteInclusionRules []SiteProtectionRuleable
    // The protection units (sites) that are protected under the site protection policy.
    siteProtectionUnits []SiteProtectionUnitable
}
// NewSharePointProtectionPolicy instantiates a new SharePointProtectionPolicy and sets the default values.
func NewSharePointProtectionPolicy()(*SharePointProtectionPolicy) {
    m := &SharePointProtectionPolicy{
        ProtectionPolicyBase: *NewProtectionPolicyBase(),
    }
    odataTypeValue := "#microsoft.graph.sharePointProtectionPolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSharePointProtectionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSharePointProtectionPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharePointProtectionPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SharePointProtectionPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectionPolicyBase.GetFieldDeserializers()
    res["siteInclusionRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSiteProtectionRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SiteProtectionRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SiteProtectionRuleable)
                }
            }
            m.SetSiteInclusionRules(res)
        }
        return nil
    }
    res["siteProtectionUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSiteProtectionUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SiteProtectionUnitable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SiteProtectionUnitable)
                }
            }
            m.SetSiteProtectionUnits(res)
        }
        return nil
    }
    return res
}
// GetSiteInclusionRules gets the siteInclusionRules property value. The rules associated with the SharePoint Protection policy.
// returns a []SiteProtectionRuleable when successful
func (m *SharePointProtectionPolicy) GetSiteInclusionRules()([]SiteProtectionRuleable) {
    return m.siteInclusionRules
}
// GetSiteProtectionUnits gets the siteProtectionUnits property value. The protection units (sites) that are protected under the site protection policy.
// returns a []SiteProtectionUnitable when successful
func (m *SharePointProtectionPolicy) GetSiteProtectionUnits()([]SiteProtectionUnitable) {
    return m.siteProtectionUnits
}
// Serialize serializes information the current object
func (m *SharePointProtectionPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectionPolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSiteInclusionRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSiteInclusionRules()))
        for i, v := range m.GetSiteInclusionRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("siteInclusionRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSiteProtectionUnits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSiteProtectionUnits()))
        for i, v := range m.GetSiteProtectionUnits() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("siteProtectionUnits", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSiteInclusionRules sets the siteInclusionRules property value. The rules associated with the SharePoint Protection policy.
func (m *SharePointProtectionPolicy) SetSiteInclusionRules(value []SiteProtectionRuleable)() {
    m.siteInclusionRules = value
}
// SetSiteProtectionUnits sets the siteProtectionUnits property value. The protection units (sites) that are protected under the site protection policy.
func (m *SharePointProtectionPolicy) SetSiteProtectionUnits(value []SiteProtectionUnitable)() {
    m.siteProtectionUnits = value
}
type SharePointProtectionPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectionPolicyBaseable
    GetSiteInclusionRules()([]SiteProtectionRuleable)
    GetSiteProtectionUnits()([]SiteProtectionUnitable)
    SetSiteInclusionRules(value []SiteProtectionRuleable)()
    SetSiteProtectionUnits(value []SiteProtectionUnitable)()
}
