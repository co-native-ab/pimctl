package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CrossTenantAccessPolicy struct {
    PolicyBase
    // Used to specify which Microsoft clouds an organization would like to collaborate with. By default, this value is empty. Supported values for this field are: microsoftonline.com, microsoftonline.us, and partner.microsoftonline.cn.
    allowedCloudEndpoints []string
    // Defines the default configuration for how your organization interacts with external Microsoft Entra organizations.
    defaultEscaped CrossTenantAccessPolicyConfigurationDefaultable
    // Defines partner-specific configurations for external Microsoft Entra organizations.
    partners []CrossTenantAccessPolicyConfigurationPartnerable
    // Represents the base policy in the directory for multitenant organization settings.
    templates PolicyTemplateable
}
// NewCrossTenantAccessPolicy instantiates a new CrossTenantAccessPolicy and sets the default values.
func NewCrossTenantAccessPolicy()(*CrossTenantAccessPolicy) {
    m := &CrossTenantAccessPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    odataTypeValue := "#microsoft.graph.crossTenantAccessPolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCrossTenantAccessPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCrossTenantAccessPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicy(), nil
}
// GetAllowedCloudEndpoints gets the allowedCloudEndpoints property value. Used to specify which Microsoft clouds an organization would like to collaborate with. By default, this value is empty. Supported values for this field are: microsoftonline.com, microsoftonline.us, and partner.microsoftonline.cn.
// returns a []string when successful
func (m *CrossTenantAccessPolicy) GetAllowedCloudEndpoints()([]string) {
    return m.allowedCloudEndpoints
}
// GetDefaultEscaped gets the default property value. Defines the default configuration for how your organization interacts with external Microsoft Entra organizations.
// returns a CrossTenantAccessPolicyConfigurationDefaultable when successful
func (m *CrossTenantAccessPolicy) GetDefaultEscaped()(CrossTenantAccessPolicyConfigurationDefaultable) {
    return m.defaultEscaped
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CrossTenantAccessPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["allowedCloudEndpoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAllowedCloudEndpoints(res)
        }
        return nil
    }
    res["default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyConfigurationDefaultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultEscaped(val.(CrossTenantAccessPolicyConfigurationDefaultable))
        }
        return nil
    }
    res["partners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCrossTenantAccessPolicyConfigurationPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CrossTenantAccessPolicyConfigurationPartnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CrossTenantAccessPolicyConfigurationPartnerable)
                }
            }
            m.SetPartners(res)
        }
        return nil
    }
    res["templates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePolicyTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplates(val.(PolicyTemplateable))
        }
        return nil
    }
    return res
}
// GetPartners gets the partners property value. Defines partner-specific configurations for external Microsoft Entra organizations.
// returns a []CrossTenantAccessPolicyConfigurationPartnerable when successful
func (m *CrossTenantAccessPolicy) GetPartners()([]CrossTenantAccessPolicyConfigurationPartnerable) {
    return m.partners
}
// GetTemplates gets the templates property value. Represents the base policy in the directory for multitenant organization settings.
// returns a PolicyTemplateable when successful
func (m *CrossTenantAccessPolicy) GetTemplates()(PolicyTemplateable) {
    return m.templates
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedCloudEndpoints() != nil {
        err = writer.WriteCollectionOfStringValues("allowedCloudEndpoints", m.GetAllowedCloudEndpoints())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("default", m.GetDefaultEscaped())
        if err != nil {
            return err
        }
    }
    if m.GetPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPartners()))
        for i, v := range m.GetPartners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("partners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("templates", m.GetTemplates())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedCloudEndpoints sets the allowedCloudEndpoints property value. Used to specify which Microsoft clouds an organization would like to collaborate with. By default, this value is empty. Supported values for this field are: microsoftonline.com, microsoftonline.us, and partner.microsoftonline.cn.
func (m *CrossTenantAccessPolicy) SetAllowedCloudEndpoints(value []string)() {
    m.allowedCloudEndpoints = value
}
// SetDefaultEscaped sets the default property value. Defines the default configuration for how your organization interacts with external Microsoft Entra organizations.
func (m *CrossTenantAccessPolicy) SetDefaultEscaped(value CrossTenantAccessPolicyConfigurationDefaultable)() {
    m.defaultEscaped = value
}
// SetPartners sets the partners property value. Defines partner-specific configurations for external Microsoft Entra organizations.
func (m *CrossTenantAccessPolicy) SetPartners(value []CrossTenantAccessPolicyConfigurationPartnerable)() {
    m.partners = value
}
// SetTemplates sets the templates property value. Represents the base policy in the directory for multitenant organization settings.
func (m *CrossTenantAccessPolicy) SetTemplates(value PolicyTemplateable)() {
    m.templates = value
}
type CrossTenantAccessPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyBaseable
    GetAllowedCloudEndpoints()([]string)
    GetDefaultEscaped()(CrossTenantAccessPolicyConfigurationDefaultable)
    GetPartners()([]CrossTenantAccessPolicyConfigurationPartnerable)
    GetTemplates()(PolicyTemplateable)
    SetAllowedCloudEndpoints(value []string)()
    SetDefaultEscaped(value CrossTenantAccessPolicyConfigurationDefaultable)()
    SetPartners(value []CrossTenantAccessPolicyConfigurationPartnerable)()
    SetTemplates(value PolicyTemplateable)()
}
