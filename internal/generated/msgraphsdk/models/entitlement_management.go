package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EntitlementManagement struct {
    Entity
    // Approval stages for decisions associated with access package assignment requests.
    accessPackageAssignmentApprovals []Approvalable
    // Access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
    accessPackages []AccessPackageable
    // Access package assignment policies govern which subjects can request or be assigned an access package via an access package assignment.
    assignmentPolicies []AccessPackageAssignmentPolicyable
    // Access package assignment requests created by or on behalf of a subject.
    assignmentRequests []AccessPackageAssignmentRequestable
    // The assignment of an access package to a subject for a period of time.
    assignments []AccessPackageAssignmentable
    // A container for access packages.
    catalogs []AccessPackageCatalogable
    // References to a directory or domain of another organization whose users can request access.
    connectedOrganizations []ConnectedOrganizationable
    // A reference to the geolocation environments in which a resource is located.
    resourceEnvironments []AccessPackageResourceEnvironmentable
    // Represents a request to add or remove a resource to or from a catalog respectively.
    resourceRequests []AccessPackageResourceRequestable
    // The resourceRoleScopes property
    resourceRoleScopes []AccessPackageResourceRoleScopeable
    // The resources associated with the catalogs.
    resources []AccessPackageResourceable
    // The settings that control the behavior of Microsoft Entra entitlement management.
    settings EntitlementManagementSettingsable
}
// NewEntitlementManagement instantiates a new EntitlementManagement and sets the default values.
func NewEntitlementManagement()(*EntitlementManagement) {
    m := &EntitlementManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEntitlementManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagement(), nil
}
// GetAccessPackageAssignmentApprovals gets the accessPackageAssignmentApprovals property value. Approval stages for decisions associated with access package assignment requests.
// returns a []Approvalable when successful
func (m *EntitlementManagement) GetAccessPackageAssignmentApprovals()([]Approvalable) {
    return m.accessPackageAssignmentApprovals
}
// GetAccessPackages gets the accessPackages property value. Access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
// returns a []AccessPackageable when successful
func (m *EntitlementManagement) GetAccessPackages()([]AccessPackageable) {
    return m.accessPackages
}
// GetAssignmentPolicies gets the assignmentPolicies property value. Access package assignment policies govern which subjects can request or be assigned an access package via an access package assignment.
// returns a []AccessPackageAssignmentPolicyable when successful
func (m *EntitlementManagement) GetAssignmentPolicies()([]AccessPackageAssignmentPolicyable) {
    return m.assignmentPolicies
}
// GetAssignmentRequests gets the assignmentRequests property value. Access package assignment requests created by or on behalf of a subject.
// returns a []AccessPackageAssignmentRequestable when successful
func (m *EntitlementManagement) GetAssignmentRequests()([]AccessPackageAssignmentRequestable) {
    return m.assignmentRequests
}
// GetAssignments gets the assignments property value. The assignment of an access package to a subject for a period of time.
// returns a []AccessPackageAssignmentable when successful
func (m *EntitlementManagement) GetAssignments()([]AccessPackageAssignmentable) {
    return m.assignments
}
// GetCatalogs gets the catalogs property value. A container for access packages.
// returns a []AccessPackageCatalogable when successful
func (m *EntitlementManagement) GetCatalogs()([]AccessPackageCatalogable) {
    return m.catalogs
}
// GetConnectedOrganizations gets the connectedOrganizations property value. References to a directory or domain of another organization whose users can request access.
// returns a []ConnectedOrganizationable when successful
func (m *EntitlementManagement) GetConnectedOrganizations()([]ConnectedOrganizationable) {
    return m.connectedOrganizations
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EntitlementManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignmentApprovals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Approvalable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Approvalable)
                }
            }
            m.SetAccessPackageAssignmentApprovals(res)
        }
        return nil
    }
    res["accessPackages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageable)
                }
            }
            m.SetAccessPackages(res)
        }
        return nil
    }
    res["assignmentPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageAssignmentPolicyable)
                }
            }
            m.SetAssignmentPolicies(res)
        }
        return nil
    }
    res["assignmentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageAssignmentRequestable)
                }
            }
            m.SetAssignmentRequests(res)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["catalogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageCatalogable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageCatalogable)
                }
            }
            m.SetCatalogs(res)
        }
        return nil
    }
    res["connectedOrganizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectedOrganizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectedOrganizationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConnectedOrganizationable)
                }
            }
            m.SetConnectedOrganizations(res)
        }
        return nil
    }
    res["resourceEnvironments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceEnvironmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceEnvironmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceEnvironmentable)
                }
            }
            m.SetResourceEnvironments(res)
        }
        return nil
    }
    res["resourceRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceRequestable)
                }
            }
            m.SetResourceRequests(res)
        }
        return nil
    }
    res["resourceRoleScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceRoleScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRoleScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceRoleScopeable)
                }
            }
            m.SetResourceRoleScopes(res)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntitlementManagementSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(EntitlementManagementSettingsable))
        }
        return nil
    }
    return res
}
// GetResourceEnvironments gets the resourceEnvironments property value. A reference to the geolocation environments in which a resource is located.
// returns a []AccessPackageResourceEnvironmentable when successful
func (m *EntitlementManagement) GetResourceEnvironments()([]AccessPackageResourceEnvironmentable) {
    return m.resourceEnvironments
}
// GetResourceRequests gets the resourceRequests property value. Represents a request to add or remove a resource to or from a catalog respectively.
// returns a []AccessPackageResourceRequestable when successful
func (m *EntitlementManagement) GetResourceRequests()([]AccessPackageResourceRequestable) {
    return m.resourceRequests
}
// GetResourceRoleScopes gets the resourceRoleScopes property value. The resourceRoleScopes property
// returns a []AccessPackageResourceRoleScopeable when successful
func (m *EntitlementManagement) GetResourceRoleScopes()([]AccessPackageResourceRoleScopeable) {
    return m.resourceRoleScopes
}
// GetResources gets the resources property value. The resources associated with the catalogs.
// returns a []AccessPackageResourceable when successful
func (m *EntitlementManagement) GetResources()([]AccessPackageResourceable) {
    return m.resources
}
// GetSettings gets the settings property value. The settings that control the behavior of Microsoft Entra entitlement management.
// returns a EntitlementManagementSettingsable when successful
func (m *EntitlementManagement) GetSettings()(EntitlementManagementSettingsable) {
    return m.settings
}
// Serialize serializes information the current object
func (m *EntitlementManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageAssignmentApprovals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignmentApprovals()))
        for i, v := range m.GetAccessPackageAssignmentApprovals() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentApprovals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackages()))
        for i, v := range m.GetAccessPackages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentPolicies()))
        for i, v := range m.GetAssignmentPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignmentPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentRequests()))
        for i, v := range m.GetAssignmentRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCatalogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCatalogs()))
        for i, v := range m.GetCatalogs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("catalogs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectedOrganizations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectedOrganizations()))
        for i, v := range m.GetConnectedOrganizations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("connectedOrganizations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceEnvironments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceEnvironments()))
        for i, v := range m.GetResourceEnvironments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceEnvironments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceRequests()))
        for i, v := range m.GetResourceRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceRoleScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceRoleScopes()))
        for i, v := range m.GetResourceRoleScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceRoleScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackageAssignmentApprovals sets the accessPackageAssignmentApprovals property value. Approval stages for decisions associated with access package assignment requests.
func (m *EntitlementManagement) SetAccessPackageAssignmentApprovals(value []Approvalable)() {
    m.accessPackageAssignmentApprovals = value
}
// SetAccessPackages sets the accessPackages property value. Access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
func (m *EntitlementManagement) SetAccessPackages(value []AccessPackageable)() {
    m.accessPackages = value
}
// SetAssignmentPolicies sets the assignmentPolicies property value. Access package assignment policies govern which subjects can request or be assigned an access package via an access package assignment.
func (m *EntitlementManagement) SetAssignmentPolicies(value []AccessPackageAssignmentPolicyable)() {
    m.assignmentPolicies = value
}
// SetAssignmentRequests sets the assignmentRequests property value. Access package assignment requests created by or on behalf of a subject.
func (m *EntitlementManagement) SetAssignmentRequests(value []AccessPackageAssignmentRequestable)() {
    m.assignmentRequests = value
}
// SetAssignments sets the assignments property value. The assignment of an access package to a subject for a period of time.
func (m *EntitlementManagement) SetAssignments(value []AccessPackageAssignmentable)() {
    m.assignments = value
}
// SetCatalogs sets the catalogs property value. A container for access packages.
func (m *EntitlementManagement) SetCatalogs(value []AccessPackageCatalogable)() {
    m.catalogs = value
}
// SetConnectedOrganizations sets the connectedOrganizations property value. References to a directory or domain of another organization whose users can request access.
func (m *EntitlementManagement) SetConnectedOrganizations(value []ConnectedOrganizationable)() {
    m.connectedOrganizations = value
}
// SetResourceEnvironments sets the resourceEnvironments property value. A reference to the geolocation environments in which a resource is located.
func (m *EntitlementManagement) SetResourceEnvironments(value []AccessPackageResourceEnvironmentable)() {
    m.resourceEnvironments = value
}
// SetResourceRequests sets the resourceRequests property value. Represents a request to add or remove a resource to or from a catalog respectively.
func (m *EntitlementManagement) SetResourceRequests(value []AccessPackageResourceRequestable)() {
    m.resourceRequests = value
}
// SetResourceRoleScopes sets the resourceRoleScopes property value. The resourceRoleScopes property
func (m *EntitlementManagement) SetResourceRoleScopes(value []AccessPackageResourceRoleScopeable)() {
    m.resourceRoleScopes = value
}
// SetResources sets the resources property value. The resources associated with the catalogs.
func (m *EntitlementManagement) SetResources(value []AccessPackageResourceable)() {
    m.resources = value
}
// SetSettings sets the settings property value. The settings that control the behavior of Microsoft Entra entitlement management.
func (m *EntitlementManagement) SetSettings(value EntitlementManagementSettingsable)() {
    m.settings = value
}
type EntitlementManagementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageAssignmentApprovals()([]Approvalable)
    GetAccessPackages()([]AccessPackageable)
    GetAssignmentPolicies()([]AccessPackageAssignmentPolicyable)
    GetAssignmentRequests()([]AccessPackageAssignmentRequestable)
    GetAssignments()([]AccessPackageAssignmentable)
    GetCatalogs()([]AccessPackageCatalogable)
    GetConnectedOrganizations()([]ConnectedOrganizationable)
    GetResourceEnvironments()([]AccessPackageResourceEnvironmentable)
    GetResourceRequests()([]AccessPackageResourceRequestable)
    GetResourceRoleScopes()([]AccessPackageResourceRoleScopeable)
    GetResources()([]AccessPackageResourceable)
    GetSettings()(EntitlementManagementSettingsable)
    SetAccessPackageAssignmentApprovals(value []Approvalable)()
    SetAccessPackages(value []AccessPackageable)()
    SetAssignmentPolicies(value []AccessPackageAssignmentPolicyable)()
    SetAssignmentRequests(value []AccessPackageAssignmentRequestable)()
    SetAssignments(value []AccessPackageAssignmentable)()
    SetCatalogs(value []AccessPackageCatalogable)()
    SetConnectedOrganizations(value []ConnectedOrganizationable)()
    SetResourceEnvironments(value []AccessPackageResourceEnvironmentable)()
    SetResourceRequests(value []AccessPackageResourceRequestable)()
    SetResourceRoleScopes(value []AccessPackageResourceRoleScopeable)()
    SetResources(value []AccessPackageResourceable)()
    SetSettings(value EntitlementManagementSettingsable)()
}
