package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcAuditActor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the application.
    applicationDisplayName *string
    // Microsoft Entra application ID.
    applicationId *string
    // IP address.
    ipAddress *string
    // The OdataType property
    odataType *string
    // The delegated partner tenant ID.
    remoteTenantId *string
    // The delegated partner user ID.
    remoteUserId *string
    // Service Principal Name (SPN).
    servicePrincipalName *string
    // Microsoft Entra user ID.
    userId *string
    // List of user permissions and application permissions when the audit event was performed.
    userPermissions []string
    // User Principal Name (UPN).
    userPrincipalName *string
    // List of role scope tags.
    userRoleScopeTags []CloudPcUserRoleScopeTagInfoable
}
// NewCloudPcAuditActor instantiates a new CloudPcAuditActor and sets the default values.
func NewCloudPcAuditActor()(*CloudPcAuditActor) {
    m := &CloudPcAuditActor{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcAuditActorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcAuditActorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcAuditActor(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcAuditActor) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApplicationDisplayName gets the applicationDisplayName property value. Name of the application.
// returns a *string when successful
func (m *CloudPcAuditActor) GetApplicationDisplayName()(*string) {
    return m.applicationDisplayName
}
// GetApplicationId gets the applicationId property value. Microsoft Entra application ID.
// returns a *string when successful
func (m *CloudPcAuditActor) GetApplicationId()(*string) {
    return m.applicationId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcAuditActor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationDisplayName(val)
        }
        return nil
    }
    res["applicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
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
    res["remoteTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteTenantId(val)
        }
        return nil
    }
    res["remoteUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteUserId(val)
        }
        return nil
    }
    res["servicePrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetUserPermissions(res)
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
    res["userRoleScopeTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcUserRoleScopeTagInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcUserRoleScopeTagInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcUserRoleScopeTagInfoable)
                }
            }
            m.SetUserRoleScopeTags(res)
        }
        return nil
    }
    return res
}
// GetIpAddress gets the ipAddress property value. IP address.
// returns a *string when successful
func (m *CloudPcAuditActor) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcAuditActor) GetOdataType()(*string) {
    return m.odataType
}
// GetRemoteTenantId gets the remoteTenantId property value. The delegated partner tenant ID.
// returns a *string when successful
func (m *CloudPcAuditActor) GetRemoteTenantId()(*string) {
    return m.remoteTenantId
}
// GetRemoteUserId gets the remoteUserId property value. The delegated partner user ID.
// returns a *string when successful
func (m *CloudPcAuditActor) GetRemoteUserId()(*string) {
    return m.remoteUserId
}
// GetServicePrincipalName gets the servicePrincipalName property value. Service Principal Name (SPN).
// returns a *string when successful
func (m *CloudPcAuditActor) GetServicePrincipalName()(*string) {
    return m.servicePrincipalName
}
// GetUserId gets the userId property value. Microsoft Entra user ID.
// returns a *string when successful
func (m *CloudPcAuditActor) GetUserId()(*string) {
    return m.userId
}
// GetUserPermissions gets the userPermissions property value. List of user permissions and application permissions when the audit event was performed.
// returns a []string when successful
func (m *CloudPcAuditActor) GetUserPermissions()([]string) {
    return m.userPermissions
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name (UPN).
// returns a *string when successful
func (m *CloudPcAuditActor) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUserRoleScopeTags gets the userRoleScopeTags property value. List of role scope tags.
// returns a []CloudPcUserRoleScopeTagInfoable when successful
func (m *CloudPcAuditActor) GetUserRoleScopeTags()([]CloudPcUserRoleScopeTagInfoable) {
    return m.userRoleScopeTags
}
// Serialize serializes information the current object
func (m *CloudPcAuditActor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationDisplayName", m.GetApplicationDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
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
    {
        err := writer.WriteStringValue("remoteTenantId", m.GetRemoteTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remoteUserId", m.GetRemoteUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePrincipalName", m.GetServicePrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    if m.GetUserPermissions() != nil {
        err := writer.WriteCollectionOfStringValues("userPermissions", m.GetUserPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetUserRoleScopeTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserRoleScopeTags()))
        for i, v := range m.GetUserRoleScopeTags() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("userRoleScopeTags", cast)
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
func (m *CloudPcAuditActor) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApplicationDisplayName sets the applicationDisplayName property value. Name of the application.
func (m *CloudPcAuditActor) SetApplicationDisplayName(value *string)() {
    m.applicationDisplayName = value
}
// SetApplicationId sets the applicationId property value. Microsoft Entra application ID.
func (m *CloudPcAuditActor) SetApplicationId(value *string)() {
    m.applicationId = value
}
// SetIpAddress sets the ipAddress property value. IP address.
func (m *CloudPcAuditActor) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcAuditActor) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemoteTenantId sets the remoteTenantId property value. The delegated partner tenant ID.
func (m *CloudPcAuditActor) SetRemoteTenantId(value *string)() {
    m.remoteTenantId = value
}
// SetRemoteUserId sets the remoteUserId property value. The delegated partner user ID.
func (m *CloudPcAuditActor) SetRemoteUserId(value *string)() {
    m.remoteUserId = value
}
// SetServicePrincipalName sets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *CloudPcAuditActor) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
// SetUserId sets the userId property value. Microsoft Entra user ID.
func (m *CloudPcAuditActor) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPermissions sets the userPermissions property value. List of user permissions and application permissions when the audit event was performed.
func (m *CloudPcAuditActor) SetUserPermissions(value []string)() {
    m.userPermissions = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name (UPN).
func (m *CloudPcAuditActor) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserRoleScopeTags sets the userRoleScopeTags property value. List of role scope tags.
func (m *CloudPcAuditActor) SetUserRoleScopeTags(value []CloudPcUserRoleScopeTagInfoable)() {
    m.userRoleScopeTags = value
}
type CloudPcAuditActorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationDisplayName()(*string)
    GetApplicationId()(*string)
    GetIpAddress()(*string)
    GetOdataType()(*string)
    GetRemoteTenantId()(*string)
    GetRemoteUserId()(*string)
    GetServicePrincipalName()(*string)
    GetUserId()(*string)
    GetUserPermissions()([]string)
    GetUserPrincipalName()(*string)
    GetUserRoleScopeTags()([]CloudPcUserRoleScopeTagInfoable)
    SetApplicationDisplayName(value *string)()
    SetApplicationId(value *string)()
    SetIpAddress(value *string)()
    SetOdataType(value *string)()
    SetRemoteTenantId(value *string)()
    SetRemoteUserId(value *string)()
    SetServicePrincipalName(value *string)()
    SetUserId(value *string)()
    SetUserPermissions(value []string)()
    SetUserPrincipalName(value *string)()
    SetUserRoleScopeTags(value []CloudPcUserRoleScopeTagInfoable)()
}
