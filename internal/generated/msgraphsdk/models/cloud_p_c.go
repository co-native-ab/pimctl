package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPC struct {
    Entity
    // The Microsoft Entra device ID for the Cloud PC, also known as the Azure Active Directory (Azure AD) device ID, that consists of 32 characters in a GUID format. Generated on a VM joined to Microsoft Entra ID. Read-only.
    aadDeviceId *string
    // The display name for the Cloud PC. Maximum length is 64 characters. Read-only. You can use the cloudPC: rename API to modify the Cloud PC name.
    displayName *string
    // The date and time when the grace period ends and reprovisioning or deprovisioning happen. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    gracePeriodEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the operating system image used for the Cloud PC. Maximum length is 50 characters. Only letters (A-Z, a-z), numbers (0-9), and special characters (-,,.) are allowed for this property. The property value can't begin or end with an underscore. Read-only.
    imageDisplayName *string
    // The last modified date and time of the Cloud PC. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Intune enrolled device ID for the Cloud PC that consists of 32 characters in a GUID format. The managedDeviceId property of Windows 365 Business Cloud PCs is always null as Windows 365 Business Cloud PCs aren't Intune-enrolled automatically by Windows 365. Read-only.
    managedDeviceId *string
    // The Intune enrolled device name for the Cloud PC. The managedDeviceName property of Windows 365 Business Cloud PCs is always null as Windows 365 Business Cloud PCs aren't Intune-enrolled automatically by Windows 365. Read-only.
    managedDeviceName *string
    // The on-premises connection that applied during the provisioning of Cloud PCs. Read-only.
    onPremisesConnectionName *string
    // The provisioning policy ID for the Cloud PC that consists of 32 characters in a GUID format. A policy defines the type of Cloud PC the user wants to create. Read-only.
    provisioningPolicyId *string
    // The provisioning policy that applied during the provisioning of Cloud PCs. Maximum length is 120 characters. Read-only.
    provisioningPolicyName *string
    // The type of licenses to be used when provisioning Cloud PCs using this policy. Possible values are: dedicated, shared, unknownFutureValue. The default value is dedicated.
    provisioningType *CloudPcProvisioningType
    // The service plan ID for the Cloud PC that consists of 32 characters in a GUID format. For more information about service plans, see Product names and service plan identifiers for licensing. Read-only.
    servicePlanId *string
    // The service plan name for the customer-facing Cloud PC entity. Read-only.
    servicePlanName *string
    // The user principal name (UPN) of the user assigned to the Cloud PC. Maximum length is 113 characters. For more information on username policies, see Password policies and account restrictions in Microsoft Entra ID. Read-only.
    userPrincipalName *string
}
// NewCloudPC instantiates a new CloudPC and sets the default values.
func NewCloudPC()(*CloudPC) {
    m := &CloudPC{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPCFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPCFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPC(), nil
}
// GetAadDeviceId gets the aadDeviceId property value. The Microsoft Entra device ID for the Cloud PC, also known as the Azure Active Directory (Azure AD) device ID, that consists of 32 characters in a GUID format. Generated on a VM joined to Microsoft Entra ID. Read-only.
// returns a *string when successful
func (m *CloudPC) GetAadDeviceId()(*string) {
    return m.aadDeviceId
}
// GetDisplayName gets the displayName property value. The display name for the Cloud PC. Maximum length is 64 characters. Read-only. You can use the cloudPC: rename API to modify the Cloud PC name.
// returns a *string when successful
func (m *CloudPC) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPC) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aadDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadDeviceId(val)
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
    res["gracePeriodEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodEndDateTime(val)
        }
        return nil
    }
    res["imageDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDisplayName(val)
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
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["onPremisesConnectionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionName(val)
        }
        return nil
    }
    res["provisioningPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyId(val)
        }
        return nil
    }
    res["provisioningPolicyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyName(val)
        }
        return nil
    }
    res["provisioningType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningType(val.(*CloudPcProvisioningType))
        }
        return nil
    }
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    res["servicePlanName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetGracePeriodEndDateTime gets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning or deprovisioning happen. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *CloudPC) GetGracePeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.gracePeriodEndDateTime
}
// GetImageDisplayName gets the imageDisplayName property value. The name of the operating system image used for the Cloud PC. Maximum length is 50 characters. Only letters (A-Z, a-z), numbers (0-9), and special characters (-,,.) are allowed for this property. The property value can't begin or end with an underscore. Read-only.
// returns a *string when successful
func (m *CloudPC) GetImageDisplayName()(*string) {
    return m.imageDisplayName
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last modified date and time of the Cloud PC. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *CloudPC) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetManagedDeviceId gets the managedDeviceId property value. The Intune enrolled device ID for the Cloud PC that consists of 32 characters in a GUID format. The managedDeviceId property of Windows 365 Business Cloud PCs is always null as Windows 365 Business Cloud PCs aren't Intune-enrolled automatically by Windows 365. Read-only.
// returns a *string when successful
func (m *CloudPC) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetManagedDeviceName gets the managedDeviceName property value. The Intune enrolled device name for the Cloud PC. The managedDeviceName property of Windows 365 Business Cloud PCs is always null as Windows 365 Business Cloud PCs aren't Intune-enrolled automatically by Windows 365. Read-only.
// returns a *string when successful
func (m *CloudPC) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetOnPremisesConnectionName gets the onPremisesConnectionName property value. The on-premises connection that applied during the provisioning of Cloud PCs. Read-only.
// returns a *string when successful
func (m *CloudPC) GetOnPremisesConnectionName()(*string) {
    return m.onPremisesConnectionName
}
// GetProvisioningPolicyId gets the provisioningPolicyId property value. The provisioning policy ID for the Cloud PC that consists of 32 characters in a GUID format. A policy defines the type of Cloud PC the user wants to create. Read-only.
// returns a *string when successful
func (m *CloudPC) GetProvisioningPolicyId()(*string) {
    return m.provisioningPolicyId
}
// GetProvisioningPolicyName gets the provisioningPolicyName property value. The provisioning policy that applied during the provisioning of Cloud PCs. Maximum length is 120 characters. Read-only.
// returns a *string when successful
func (m *CloudPC) GetProvisioningPolicyName()(*string) {
    return m.provisioningPolicyName
}
// GetProvisioningType gets the provisioningType property value. The type of licenses to be used when provisioning Cloud PCs using this policy. Possible values are: dedicated, shared, unknownFutureValue. The default value is dedicated.
// returns a *CloudPcProvisioningType when successful
func (m *CloudPC) GetProvisioningType()(*CloudPcProvisioningType) {
    return m.provisioningType
}
// GetServicePlanId gets the servicePlanId property value. The service plan ID for the Cloud PC that consists of 32 characters in a GUID format. For more information about service plans, see Product names and service plan identifiers for licensing. Read-only.
// returns a *string when successful
func (m *CloudPC) GetServicePlanId()(*string) {
    return m.servicePlanId
}
// GetServicePlanName gets the servicePlanName property value. The service plan name for the customer-facing Cloud PC entity. Read-only.
// returns a *string when successful
func (m *CloudPC) GetServicePlanName()(*string) {
    return m.servicePlanName
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC. Maximum length is 113 characters. For more information on username policies, see Password policies and account restrictions in Microsoft Entra ID. Read-only.
// returns a *string when successful
func (m *CloudPC) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *CloudPC) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("aadDeviceId", m.GetAadDeviceId())
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
        err = writer.WriteTimeValue("gracePeriodEndDateTime", m.GetGracePeriodEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imageDisplayName", m.GetImageDisplayName())
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
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesConnectionName", m.GetOnPremisesConnectionName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provisioningPolicyId", m.GetProvisioningPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provisioningPolicyName", m.GetProvisioningPolicyName())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningType() != nil {
        cast := (*m.GetProvisioningType()).String()
        err = writer.WriteStringValue("provisioningType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePlanName", m.GetServicePlanName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAadDeviceId sets the aadDeviceId property value. The Microsoft Entra device ID for the Cloud PC, also known as the Azure Active Directory (Azure AD) device ID, that consists of 32 characters in a GUID format. Generated on a VM joined to Microsoft Entra ID. Read-only.
func (m *CloudPC) SetAadDeviceId(value *string)() {
    m.aadDeviceId = value
}
// SetDisplayName sets the displayName property value. The display name for the Cloud PC. Maximum length is 64 characters. Read-only. You can use the cloudPC: rename API to modify the Cloud PC name.
func (m *CloudPC) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGracePeriodEndDateTime sets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning or deprovisioning happen. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) SetGracePeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.gracePeriodEndDateTime = value
}
// SetImageDisplayName sets the imageDisplayName property value. The name of the operating system image used for the Cloud PC. Maximum length is 50 characters. Only letters (A-Z, a-z), numbers (0-9), and special characters (-,,.) are allowed for this property. The property value can't begin or end with an underscore. Read-only.
func (m *CloudPC) SetImageDisplayName(value *string)() {
    m.imageDisplayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last modified date and time of the Cloud PC. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The Intune enrolled device ID for the Cloud PC that consists of 32 characters in a GUID format. The managedDeviceId property of Windows 365 Business Cloud PCs is always null as Windows 365 Business Cloud PCs aren't Intune-enrolled automatically by Windows 365. Read-only.
func (m *CloudPC) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetManagedDeviceName sets the managedDeviceName property value. The Intune enrolled device name for the Cloud PC. The managedDeviceName property of Windows 365 Business Cloud PCs is always null as Windows 365 Business Cloud PCs aren't Intune-enrolled automatically by Windows 365. Read-only.
func (m *CloudPC) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetOnPremisesConnectionName sets the onPremisesConnectionName property value. The on-premises connection that applied during the provisioning of Cloud PCs. Read-only.
func (m *CloudPC) SetOnPremisesConnectionName(value *string)() {
    m.onPremisesConnectionName = value
}
// SetProvisioningPolicyId sets the provisioningPolicyId property value. The provisioning policy ID for the Cloud PC that consists of 32 characters in a GUID format. A policy defines the type of Cloud PC the user wants to create. Read-only.
func (m *CloudPC) SetProvisioningPolicyId(value *string)() {
    m.provisioningPolicyId = value
}
// SetProvisioningPolicyName sets the provisioningPolicyName property value. The provisioning policy that applied during the provisioning of Cloud PCs. Maximum length is 120 characters. Read-only.
func (m *CloudPC) SetProvisioningPolicyName(value *string)() {
    m.provisioningPolicyName = value
}
// SetProvisioningType sets the provisioningType property value. The type of licenses to be used when provisioning Cloud PCs using this policy. Possible values are: dedicated, shared, unknownFutureValue. The default value is dedicated.
func (m *CloudPC) SetProvisioningType(value *CloudPcProvisioningType)() {
    m.provisioningType = value
}
// SetServicePlanId sets the servicePlanId property value. The service plan ID for the Cloud PC that consists of 32 characters in a GUID format. For more information about service plans, see Product names and service plan identifiers for licensing. Read-only.
func (m *CloudPC) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
// SetServicePlanName sets the servicePlanName property value. The service plan name for the customer-facing Cloud PC entity. Read-only.
func (m *CloudPC) SetServicePlanName(value *string)() {
    m.servicePlanName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC. Maximum length is 113 characters. For more information on username policies, see Password policies and account restrictions in Microsoft Entra ID. Read-only.
func (m *CloudPC) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
type CloudPCable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAadDeviceId()(*string)
    GetDisplayName()(*string)
    GetGracePeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetImageDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceId()(*string)
    GetManagedDeviceName()(*string)
    GetOnPremisesConnectionName()(*string)
    GetProvisioningPolicyId()(*string)
    GetProvisioningPolicyName()(*string)
    GetProvisioningType()(*CloudPcProvisioningType)
    GetServicePlanId()(*string)
    GetServicePlanName()(*string)
    GetUserPrincipalName()(*string)
    SetAadDeviceId(value *string)()
    SetDisplayName(value *string)()
    SetGracePeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetImageDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceId(value *string)()
    SetManagedDeviceName(value *string)()
    SetOnPremisesConnectionName(value *string)()
    SetProvisioningPolicyId(value *string)()
    SetProvisioningPolicyName(value *string)()
    SetProvisioningType(value *CloudPcProvisioningType)()
    SetServicePlanId(value *string)()
    SetServicePlanName(value *string)()
    SetUserPrincipalName(value *string)()
}
