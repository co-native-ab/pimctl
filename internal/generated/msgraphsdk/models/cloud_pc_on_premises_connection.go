package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcOnPremisesConnection struct {
    Entity
    // The fully qualified domain name (FQDN) of the Active Directory domain you want to join. Maximum length is 255. Optional.
    adDomainName *string
    // The password associated with the username of an Active Directory account (adDomainUsername).
    adDomainPassword *string
    // The username of an Active Directory account (user or service account) that has permission to create computer objects in Active Directory. Required format: admin@contoso.com. Optional.
    adDomainUsername *string
    // The interface URL of the partner service's resource that links to this Azure network connection. Returned only on $select.
    alternateResourceUrl *string
    // Specifies how the provisioned Cloud PC joins to Microsoft Entra. It includes different types, one is Microsoft Entra ID join, which means there's no on-premises Active Directory (AD) in the current tenant, and the Cloud PC device is joined by Microsoft Entra. Another one is hybridAzureADJoin, which means there's also an on-premises Active Directory (AD) in the current tenant and the Cloud PC device joins to on-premises Active Directory (AD) and Microsoft Entra. The type also determines which types of users can be assigned and can sign into a Cloud PC. The azureADJoin type indicates that cloud-only and hybrid users can be assigned and signed into the Cloud PC. hybridAzureADJoin indicates only hybrid users can be assigned and signed into the Cloud PC. The default value is hybridAzureADJoin.
    connectionType *CloudPcOnPremisesConnectionType
    // The display name for the Azure network connection.
    displayName *string
    // The healthCheckStatus property
    healthCheckStatus *CloudPcOnPremisesConnectionStatus
    // Indicates the results of health checks performed on the on-premises connection. Read-only. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetail. Read-only.
    healthCheckStatusDetail CloudPcOnPremisesConnectionStatusDetailable
    // When true, the Azure network connection is in use. When false, the connection isn't in use. You can't delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetail. Read-only.
    inUse *bool
    // The organizational unit (OU) in which the computer account is created. If left null, the OU configured as the default (a well-known computer object container) in the tenant's Active Directory domain (OU) is used. Optional.
    organizationalUnit *string
    // The unique identifier of the target resource group used associated with the on-premises network connectivity for Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}'
    resourceGroupId *string
    // The unique identifier of the target subnet used associated with the on-premises network connectivity for Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}'
    subnetId *string
    // The unique identifier of the Azure subscription associated with the tenant.
    subscriptionId *string
    // The name of the Azure subscription is used to create an Azure network connection. Read-only.
    subscriptionName *string
    // The unique identifier of the target virtual network used associated with the on-premises network connectivity for Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}'
    virtualNetworkId *string
    // Indicates the resource location of the target virtual network. For example, the location can be eastus2, westeurope, etc. Read-only (computed value).
    virtualNetworkLocation *string
}
// NewCloudPcOnPremisesConnection instantiates a new CloudPcOnPremisesConnection and sets the default values.
func NewCloudPcOnPremisesConnection()(*CloudPcOnPremisesConnection) {
    m := &CloudPcOnPremisesConnection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcOnPremisesConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcOnPremisesConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcOnPremisesConnection(), nil
}
// GetAdDomainName gets the adDomainName property value. The fully qualified domain name (FQDN) of the Active Directory domain you want to join. Maximum length is 255. Optional.
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetAdDomainName()(*string) {
    return m.adDomainName
}
// GetAdDomainPassword gets the adDomainPassword property value. The password associated with the username of an Active Directory account (adDomainUsername).
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetAdDomainPassword()(*string) {
    return m.adDomainPassword
}
// GetAdDomainUsername gets the adDomainUsername property value. The username of an Active Directory account (user or service account) that has permission to create computer objects in Active Directory. Required format: admin@contoso.com. Optional.
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetAdDomainUsername()(*string) {
    return m.adDomainUsername
}
// GetAlternateResourceUrl gets the alternateResourceUrl property value. The interface URL of the partner service's resource that links to this Azure network connection. Returned only on $select.
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetAlternateResourceUrl()(*string) {
    return m.alternateResourceUrl
}
// GetConnectionType gets the connectionType property value. Specifies how the provisioned Cloud PC joins to Microsoft Entra. It includes different types, one is Microsoft Entra ID join, which means there's no on-premises Active Directory (AD) in the current tenant, and the Cloud PC device is joined by Microsoft Entra. Another one is hybridAzureADJoin, which means there's also an on-premises Active Directory (AD) in the current tenant and the Cloud PC device joins to on-premises Active Directory (AD) and Microsoft Entra. The type also determines which types of users can be assigned and can sign into a Cloud PC. The azureADJoin type indicates that cloud-only and hybrid users can be assigned and signed into the Cloud PC. hybridAzureADJoin indicates only hybrid users can be assigned and signed into the Cloud PC. The default value is hybridAzureADJoin.
// returns a *CloudPcOnPremisesConnectionType when successful
func (m *CloudPcOnPremisesConnection) GetConnectionType()(*CloudPcOnPremisesConnectionType) {
    return m.connectionType
}
// GetDisplayName gets the displayName property value. The display name for the Azure network connection.
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcOnPremisesConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainName(val)
        }
        return nil
    }
    res["adDomainPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainPassword(val)
        }
        return nil
    }
    res["adDomainUsername"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainUsername(val)
        }
        return nil
    }
    res["alternateResourceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateResourceUrl(val)
        }
        return nil
    }
    res["connectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*CloudPcOnPremisesConnectionType))
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
    res["healthCheckStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthCheckStatus(val.(*CloudPcOnPremisesConnectionStatus))
        }
        return nil
    }
    res["healthCheckStatusDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcOnPremisesConnectionStatusDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthCheckStatusDetail(val.(CloudPcOnPremisesConnectionStatusDetailable))
        }
        return nil
    }
    res["inUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInUse(val)
        }
        return nil
    }
    res["organizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnit(val)
        }
        return nil
    }
    res["resourceGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceGroupId(val)
        }
        return nil
    }
    res["subnetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetId(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    res["subscriptionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionName(val)
        }
        return nil
    }
    res["virtualNetworkId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNetworkId(val)
        }
        return nil
    }
    res["virtualNetworkLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNetworkLocation(val)
        }
        return nil
    }
    return res
}
// GetHealthCheckStatus gets the healthCheckStatus property value. The healthCheckStatus property
// returns a *CloudPcOnPremisesConnectionStatus when successful
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus) {
    return m.healthCheckStatus
}
// GetHealthCheckStatusDetail gets the healthCheckStatusDetail property value. Indicates the results of health checks performed on the on-premises connection. Read-only. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetail. Read-only.
// returns a CloudPcOnPremisesConnectionStatusDetailable when successful
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatusDetail()(CloudPcOnPremisesConnectionStatusDetailable) {
    return m.healthCheckStatusDetail
}
// GetInUse gets the inUse property value. When true, the Azure network connection is in use. When false, the connection isn't in use. You can't delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetail. Read-only.
// returns a *bool when successful
func (m *CloudPcOnPremisesConnection) GetInUse()(*bool) {
    return m.inUse
}
// GetOrganizationalUnit gets the organizationalUnit property value. The organizational unit (OU) in which the computer account is created. If left null, the OU configured as the default (a well-known computer object container) in the tenant's Active Directory domain (OU) is used. Optional.
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetOrganizationalUnit()(*string) {
    return m.organizationalUnit
}
// GetResourceGroupId gets the resourceGroupId property value. The unique identifier of the target resource group used associated with the on-premises network connectivity for Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}'
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetResourceGroupId()(*string) {
    return m.resourceGroupId
}
// GetSubnetId gets the subnetId property value. The unique identifier of the target subnet used associated with the on-premises network connectivity for Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}'
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetSubnetId()(*string) {
    return m.subnetId
}
// GetSubscriptionId gets the subscriptionId property value. The unique identifier of the Azure subscription associated with the tenant.
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetSubscriptionId()(*string) {
    return m.subscriptionId
}
// GetSubscriptionName gets the subscriptionName property value. The name of the Azure subscription is used to create an Azure network connection. Read-only.
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetSubscriptionName()(*string) {
    return m.subscriptionName
}
// GetVirtualNetworkId gets the virtualNetworkId property value. The unique identifier of the target virtual network used associated with the on-premises network connectivity for Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}'
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetVirtualNetworkId()(*string) {
    return m.virtualNetworkId
}
// GetVirtualNetworkLocation gets the virtualNetworkLocation property value. Indicates the resource location of the target virtual network. For example, the location can be eastus2, westeurope, etc. Read-only (computed value).
// returns a *string when successful
func (m *CloudPcOnPremisesConnection) GetVirtualNetworkLocation()(*string) {
    return m.virtualNetworkLocation
}
// Serialize serializes information the current object
func (m *CloudPcOnPremisesConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("adDomainName", m.GetAdDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("adDomainPassword", m.GetAdDomainPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("adDomainUsername", m.GetAdDomainUsername())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alternateResourceUrl", m.GetAlternateResourceUrl())
        if err != nil {
            return err
        }
    }
    if m.GetConnectionType() != nil {
        cast := (*m.GetConnectionType()).String()
        err = writer.WriteStringValue("connectionType", &cast)
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
    if m.GetHealthCheckStatus() != nil {
        cast := (*m.GetHealthCheckStatus()).String()
        err = writer.WriteStringValue("healthCheckStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("healthCheckStatusDetail", m.GetHealthCheckStatusDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("inUse", m.GetInUse())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationalUnit", m.GetOrganizationalUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceGroupId", m.GetResourceGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subnetId", m.GetSubnetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriptionName", m.GetSubscriptionName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("virtualNetworkId", m.GetVirtualNetworkId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("virtualNetworkLocation", m.GetVirtualNetworkLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdDomainName sets the adDomainName property value. The fully qualified domain name (FQDN) of the Active Directory domain you want to join. Maximum length is 255. Optional.
func (m *CloudPcOnPremisesConnection) SetAdDomainName(value *string)() {
    m.adDomainName = value
}
// SetAdDomainPassword sets the adDomainPassword property value. The password associated with the username of an Active Directory account (adDomainUsername).
func (m *CloudPcOnPremisesConnection) SetAdDomainPassword(value *string)() {
    m.adDomainPassword = value
}
// SetAdDomainUsername sets the adDomainUsername property value. The username of an Active Directory account (user or service account) that has permission to create computer objects in Active Directory. Required format: admin@contoso.com. Optional.
func (m *CloudPcOnPremisesConnection) SetAdDomainUsername(value *string)() {
    m.adDomainUsername = value
}
// SetAlternateResourceUrl sets the alternateResourceUrl property value. The interface URL of the partner service's resource that links to this Azure network connection. Returned only on $select.
func (m *CloudPcOnPremisesConnection) SetAlternateResourceUrl(value *string)() {
    m.alternateResourceUrl = value
}
// SetConnectionType sets the connectionType property value. Specifies how the provisioned Cloud PC joins to Microsoft Entra. It includes different types, one is Microsoft Entra ID join, which means there's no on-premises Active Directory (AD) in the current tenant, and the Cloud PC device is joined by Microsoft Entra. Another one is hybridAzureADJoin, which means there's also an on-premises Active Directory (AD) in the current tenant and the Cloud PC device joins to on-premises Active Directory (AD) and Microsoft Entra. The type also determines which types of users can be assigned and can sign into a Cloud PC. The azureADJoin type indicates that cloud-only and hybrid users can be assigned and signed into the Cloud PC. hybridAzureADJoin indicates only hybrid users can be assigned and signed into the Cloud PC. The default value is hybridAzureADJoin.
func (m *CloudPcOnPremisesConnection) SetConnectionType(value *CloudPcOnPremisesConnectionType)() {
    m.connectionType = value
}
// SetDisplayName sets the displayName property value. The display name for the Azure network connection.
func (m *CloudPcOnPremisesConnection) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHealthCheckStatus sets the healthCheckStatus property value. The healthCheckStatus property
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)() {
    m.healthCheckStatus = value
}
// SetHealthCheckStatusDetail sets the healthCheckStatusDetail property value. Indicates the results of health checks performed on the on-premises connection. Read-only. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetail. Read-only.
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatusDetail(value CloudPcOnPremisesConnectionStatusDetailable)() {
    m.healthCheckStatusDetail = value
}
// SetInUse sets the inUse property value. When true, the Azure network connection is in use. When false, the connection isn't in use. You can't delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetail. Read-only.
func (m *CloudPcOnPremisesConnection) SetInUse(value *bool)() {
    m.inUse = value
}
// SetOrganizationalUnit sets the organizationalUnit property value. The organizational unit (OU) in which the computer account is created. If left null, the OU configured as the default (a well-known computer object container) in the tenant's Active Directory domain (OU) is used. Optional.
func (m *CloudPcOnPremisesConnection) SetOrganizationalUnit(value *string)() {
    m.organizationalUnit = value
}
// SetResourceGroupId sets the resourceGroupId property value. The unique identifier of the target resource group used associated with the on-premises network connectivity for Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}'
func (m *CloudPcOnPremisesConnection) SetResourceGroupId(value *string)() {
    m.resourceGroupId = value
}
// SetSubnetId sets the subnetId property value. The unique identifier of the target subnet used associated with the on-premises network connectivity for Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}'
func (m *CloudPcOnPremisesConnection) SetSubnetId(value *string)() {
    m.subnetId = value
}
// SetSubscriptionId sets the subscriptionId property value. The unique identifier of the Azure subscription associated with the tenant.
func (m *CloudPcOnPremisesConnection) SetSubscriptionId(value *string)() {
    m.subscriptionId = value
}
// SetSubscriptionName sets the subscriptionName property value. The name of the Azure subscription is used to create an Azure network connection. Read-only.
func (m *CloudPcOnPremisesConnection) SetSubscriptionName(value *string)() {
    m.subscriptionName = value
}
// SetVirtualNetworkId sets the virtualNetworkId property value. The unique identifier of the target virtual network used associated with the on-premises network connectivity for Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}'
func (m *CloudPcOnPremisesConnection) SetVirtualNetworkId(value *string)() {
    m.virtualNetworkId = value
}
// SetVirtualNetworkLocation sets the virtualNetworkLocation property value. Indicates the resource location of the target virtual network. For example, the location can be eastus2, westeurope, etc. Read-only (computed value).
func (m *CloudPcOnPremisesConnection) SetVirtualNetworkLocation(value *string)() {
    m.virtualNetworkLocation = value
}
type CloudPcOnPremisesConnectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdDomainName()(*string)
    GetAdDomainPassword()(*string)
    GetAdDomainUsername()(*string)
    GetAlternateResourceUrl()(*string)
    GetConnectionType()(*CloudPcOnPremisesConnectionType)
    GetDisplayName()(*string)
    GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus)
    GetHealthCheckStatusDetail()(CloudPcOnPremisesConnectionStatusDetailable)
    GetInUse()(*bool)
    GetOrganizationalUnit()(*string)
    GetResourceGroupId()(*string)
    GetSubnetId()(*string)
    GetSubscriptionId()(*string)
    GetSubscriptionName()(*string)
    GetVirtualNetworkId()(*string)
    GetVirtualNetworkLocation()(*string)
    SetAdDomainName(value *string)()
    SetAdDomainPassword(value *string)()
    SetAdDomainUsername(value *string)()
    SetAlternateResourceUrl(value *string)()
    SetConnectionType(value *CloudPcOnPremisesConnectionType)()
    SetDisplayName(value *string)()
    SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)()
    SetHealthCheckStatusDetail(value CloudPcOnPremisesConnectionStatusDetailable)()
    SetInUse(value *bool)()
    SetOrganizationalUnit(value *string)()
    SetResourceGroupId(value *string)()
    SetSubnetId(value *string)()
    SetSubscriptionId(value *string)()
    SetSubscriptionName(value *string)()
    SetVirtualNetworkId(value *string)()
    SetVirtualNetworkLocation(value *string)()
}
