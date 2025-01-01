package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConditionalAccessRoot struct {
    Entity
    // Read-only. Nullable. Returns a collection of the specified authentication context class references.
    authenticationContextClassReferences []AuthenticationContextClassReferenceable
    // The authenticationStrength property
    authenticationStrength AuthenticationStrengthRootable
    // Read-only. Nullable. Returns a collection of the specified named locations.
    namedLocations []NamedLocationable
    // Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies.
    policies []ConditionalAccessPolicyable
    // Read-only. Nullable. Returns a collection of the specified Conditional Access templates.
    templates []ConditionalAccessTemplateable
}
// NewConditionalAccessRoot instantiates a new ConditionalAccessRoot and sets the default values.
func NewConditionalAccessRoot()(*ConditionalAccessRoot) {
    m := &ConditionalAccessRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateConditionalAccessRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalAccessRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessRoot(), nil
}
// GetAuthenticationContextClassReferences gets the authenticationContextClassReferences property value. Read-only. Nullable. Returns a collection of the specified authentication context class references.
// returns a []AuthenticationContextClassReferenceable when successful
func (m *ConditionalAccessRoot) GetAuthenticationContextClassReferences()([]AuthenticationContextClassReferenceable) {
    return m.authenticationContextClassReferences
}
// GetAuthenticationStrength gets the authenticationStrength property value. The authenticationStrength property
// returns a AuthenticationStrengthRootable when successful
func (m *ConditionalAccessRoot) GetAuthenticationStrength()(AuthenticationStrengthRootable) {
    return m.authenticationStrength
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionalAccessRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationContextClassReferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationContextClassReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationContextClassReferenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationContextClassReferenceable)
                }
            }
            m.SetAuthenticationContextClassReferences(res)
        }
        return nil
    }
    res["authenticationStrength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationStrengthRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStrength(val.(AuthenticationStrengthRootable))
        }
        return nil
    }
    res["namedLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNamedLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NamedLocationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NamedLocationable)
                }
            }
            m.SetNamedLocations(res)
        }
        return nil
    }
    res["policies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConditionalAccessPolicyable)
                }
            }
            m.SetPolicies(res)
        }
        return nil
    }
    res["templates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConditionalAccessTemplateable)
                }
            }
            m.SetTemplates(res)
        }
        return nil
    }
    return res
}
// GetNamedLocations gets the namedLocations property value. Read-only. Nullable. Returns a collection of the specified named locations.
// returns a []NamedLocationable when successful
func (m *ConditionalAccessRoot) GetNamedLocations()([]NamedLocationable) {
    return m.namedLocations
}
// GetPolicies gets the policies property value. Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies.
// returns a []ConditionalAccessPolicyable when successful
func (m *ConditionalAccessRoot) GetPolicies()([]ConditionalAccessPolicyable) {
    return m.policies
}
// GetTemplates gets the templates property value. Read-only. Nullable. Returns a collection of the specified Conditional Access templates.
// returns a []ConditionalAccessTemplateable when successful
func (m *ConditionalAccessRoot) GetTemplates()([]ConditionalAccessTemplateable) {
    return m.templates
}
// Serialize serializes information the current object
func (m *ConditionalAccessRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationContextClassReferences() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationContextClassReferences()))
        for i, v := range m.GetAuthenticationContextClassReferences() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("authenticationContextClassReferences", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationStrength", m.GetAuthenticationStrength())
        if err != nil {
            return err
        }
    }
    if m.GetNamedLocations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNamedLocations()))
        for i, v := range m.GetNamedLocations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("namedLocations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicies()))
        for i, v := range m.GetPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("policies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplates()))
        for i, v := range m.GetTemplates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("templates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationContextClassReferences sets the authenticationContextClassReferences property value. Read-only. Nullable. Returns a collection of the specified authentication context class references.
func (m *ConditionalAccessRoot) SetAuthenticationContextClassReferences(value []AuthenticationContextClassReferenceable)() {
    m.authenticationContextClassReferences = value
}
// SetAuthenticationStrength sets the authenticationStrength property value. The authenticationStrength property
func (m *ConditionalAccessRoot) SetAuthenticationStrength(value AuthenticationStrengthRootable)() {
    m.authenticationStrength = value
}
// SetNamedLocations sets the namedLocations property value. Read-only. Nullable. Returns a collection of the specified named locations.
func (m *ConditionalAccessRoot) SetNamedLocations(value []NamedLocationable)() {
    m.namedLocations = value
}
// SetPolicies sets the policies property value. Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies.
func (m *ConditionalAccessRoot) SetPolicies(value []ConditionalAccessPolicyable)() {
    m.policies = value
}
// SetTemplates sets the templates property value. Read-only. Nullable. Returns a collection of the specified Conditional Access templates.
func (m *ConditionalAccessRoot) SetTemplates(value []ConditionalAccessTemplateable)() {
    m.templates = value
}
type ConditionalAccessRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationContextClassReferences()([]AuthenticationContextClassReferenceable)
    GetAuthenticationStrength()(AuthenticationStrengthRootable)
    GetNamedLocations()([]NamedLocationable)
    GetPolicies()([]ConditionalAccessPolicyable)
    GetTemplates()([]ConditionalAccessTemplateable)
    SetAuthenticationContextClassReferences(value []AuthenticationContextClassReferenceable)()
    SetAuthenticationStrength(value AuthenticationStrengthRootable)()
    SetNamedLocations(value []NamedLocationable)()
    SetPolicies(value []ConditionalAccessPolicyable)()
    SetTemplates(value []ConditionalAccessTemplateable)()
}
