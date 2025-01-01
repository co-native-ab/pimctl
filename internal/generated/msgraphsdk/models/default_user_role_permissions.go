package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DefaultUserRolePermissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether the default user role can create applications. This setting corresponds to the Users can register applications setting in the User settings menu in the Microsoft Entra admin center.
    allowedToCreateApps *bool
    // Indicates whether the default user role can create security groups. This setting corresponds to the following menus in the Microsoft Entra admin center:  The Users can create security groups in Microsoft Entra admin centers, API or PowerShell setting in the Group settings menu.  Users can create security groups setting in the User settings menu.
    allowedToCreateSecurityGroups *bool
    // Indicates whether the default user role can create tenants. This setting corresponds to the Restrict non-admin users from creating tenants setting in the User settings menu in the Microsoft Entra admin center.  When this setting is false, users assigned the Tenant Creator role can still create tenants.
    allowedToCreateTenants *bool
    // Indicates whether the registered owners of a device can read their own BitLocker recovery keys with default user role.
    allowedToReadBitlockerKeysForOwnedDevice *bool
    // Indicates whether the default user role can read other users. DO NOT SET THIS VALUE TO false.
    allowedToReadOtherUsers *bool
    // The OdataType property
    odataType *string
    // Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
    permissionGrantPoliciesAssigned []string
}
// NewDefaultUserRolePermissions instantiates a new DefaultUserRolePermissions and sets the default values.
func NewDefaultUserRolePermissions()(*DefaultUserRolePermissions) {
    m := &DefaultUserRolePermissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDefaultUserRolePermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDefaultUserRolePermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultUserRolePermissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DefaultUserRolePermissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedToCreateApps gets the allowedToCreateApps property value. Indicates whether the default user role can create applications. This setting corresponds to the Users can register applications setting in the User settings menu in the Microsoft Entra admin center.
// returns a *bool when successful
func (m *DefaultUserRolePermissions) GetAllowedToCreateApps()(*bool) {
    return m.allowedToCreateApps
}
// GetAllowedToCreateSecurityGroups gets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups. This setting corresponds to the following menus in the Microsoft Entra admin center:  The Users can create security groups in Microsoft Entra admin centers, API or PowerShell setting in the Group settings menu.  Users can create security groups setting in the User settings menu.
// returns a *bool when successful
func (m *DefaultUserRolePermissions) GetAllowedToCreateSecurityGroups()(*bool) {
    return m.allowedToCreateSecurityGroups
}
// GetAllowedToCreateTenants gets the allowedToCreateTenants property value. Indicates whether the default user role can create tenants. This setting corresponds to the Restrict non-admin users from creating tenants setting in the User settings menu in the Microsoft Entra admin center.  When this setting is false, users assigned the Tenant Creator role can still create tenants.
// returns a *bool when successful
func (m *DefaultUserRolePermissions) GetAllowedToCreateTenants()(*bool) {
    return m.allowedToCreateTenants
}
// GetAllowedToReadBitlockerKeysForOwnedDevice gets the allowedToReadBitlockerKeysForOwnedDevice property value. Indicates whether the registered owners of a device can read their own BitLocker recovery keys with default user role.
// returns a *bool when successful
func (m *DefaultUserRolePermissions) GetAllowedToReadBitlockerKeysForOwnedDevice()(*bool) {
    return m.allowedToReadBitlockerKeysForOwnedDevice
}
// GetAllowedToReadOtherUsers gets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users. DO NOT SET THIS VALUE TO false.
// returns a *bool when successful
func (m *DefaultUserRolePermissions) GetAllowedToReadOtherUsers()(*bool) {
    return m.allowedToReadOtherUsers
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DefaultUserRolePermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedToCreateApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToCreateApps(val)
        }
        return nil
    }
    res["allowedToCreateSecurityGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToCreateSecurityGroups(val)
        }
        return nil
    }
    res["allowedToCreateTenants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToCreateTenants(val)
        }
        return nil
    }
    res["allowedToReadBitlockerKeysForOwnedDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToReadBitlockerKeysForOwnedDevice(val)
        }
        return nil
    }
    res["allowedToReadOtherUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToReadOtherUsers(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["permissionGrantPoliciesAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPermissionGrantPoliciesAssigned(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DefaultUserRolePermissions) GetOdataType()(*string) {
    return m.odataType
}
// GetPermissionGrantPoliciesAssigned gets the permissionGrantPoliciesAssigned property value. Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
// returns a []string when successful
func (m *DefaultUserRolePermissions) GetPermissionGrantPoliciesAssigned()([]string) {
    return m.permissionGrantPoliciesAssigned
}
// Serialize serializes information the current object
func (m *DefaultUserRolePermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowedToCreateApps", m.GetAllowedToCreateApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToCreateSecurityGroups", m.GetAllowedToCreateSecurityGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToCreateTenants", m.GetAllowedToCreateTenants())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToReadBitlockerKeysForOwnedDevice", m.GetAllowedToReadBitlockerKeysForOwnedDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToReadOtherUsers", m.GetAllowedToReadOtherUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionGrantPoliciesAssigned() != nil {
        err := writer.WriteCollectionOfStringValues("permissionGrantPoliciesAssigned", m.GetPermissionGrantPoliciesAssigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DefaultUserRolePermissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedToCreateApps sets the allowedToCreateApps property value. Indicates whether the default user role can create applications. This setting corresponds to the Users can register applications setting in the User settings menu in the Microsoft Entra admin center.
func (m *DefaultUserRolePermissions) SetAllowedToCreateApps(value *bool)() {
    m.allowedToCreateApps = value
}
// SetAllowedToCreateSecurityGroups sets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups. This setting corresponds to the following menus in the Microsoft Entra admin center:  The Users can create security groups in Microsoft Entra admin centers, API or PowerShell setting in the Group settings menu.  Users can create security groups setting in the User settings menu.
func (m *DefaultUserRolePermissions) SetAllowedToCreateSecurityGroups(value *bool)() {
    m.allowedToCreateSecurityGroups = value
}
// SetAllowedToCreateTenants sets the allowedToCreateTenants property value. Indicates whether the default user role can create tenants. This setting corresponds to the Restrict non-admin users from creating tenants setting in the User settings menu in the Microsoft Entra admin center.  When this setting is false, users assigned the Tenant Creator role can still create tenants.
func (m *DefaultUserRolePermissions) SetAllowedToCreateTenants(value *bool)() {
    m.allowedToCreateTenants = value
}
// SetAllowedToReadBitlockerKeysForOwnedDevice sets the allowedToReadBitlockerKeysForOwnedDevice property value. Indicates whether the registered owners of a device can read their own BitLocker recovery keys with default user role.
func (m *DefaultUserRolePermissions) SetAllowedToReadBitlockerKeysForOwnedDevice(value *bool)() {
    m.allowedToReadBitlockerKeysForOwnedDevice = value
}
// SetAllowedToReadOtherUsers sets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users. DO NOT SET THIS VALUE TO false.
func (m *DefaultUserRolePermissions) SetAllowedToReadOtherUsers(value *bool)() {
    m.allowedToReadOtherUsers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DefaultUserRolePermissions) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPermissionGrantPoliciesAssigned sets the permissionGrantPoliciesAssigned property value. Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
func (m *DefaultUserRolePermissions) SetPermissionGrantPoliciesAssigned(value []string)() {
    m.permissionGrantPoliciesAssigned = value
}
type DefaultUserRolePermissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedToCreateApps()(*bool)
    GetAllowedToCreateSecurityGroups()(*bool)
    GetAllowedToCreateTenants()(*bool)
    GetAllowedToReadBitlockerKeysForOwnedDevice()(*bool)
    GetAllowedToReadOtherUsers()(*bool)
    GetOdataType()(*string)
    GetPermissionGrantPoliciesAssigned()([]string)
    SetAllowedToCreateApps(value *bool)()
    SetAllowedToCreateSecurityGroups(value *bool)()
    SetAllowedToCreateTenants(value *bool)()
    SetAllowedToReadBitlockerKeysForOwnedDevice(value *bool)()
    SetAllowedToReadOtherUsers(value *bool)()
    SetOdataType(value *string)()
    SetPermissionGrantPoliciesAssigned(value []string)()
}
