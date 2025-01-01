package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServicePrincipalEvidence struct {
    AlertEvidence
    // The appId property
    appId *string
    // The appOwnerTenantId property
    appOwnerTenantId *string
    // The servicePrincipalName property
    servicePrincipalName *string
    // The servicePrincipalObjectId property
    servicePrincipalObjectId *string
    // The servicePrincipalType property
    servicePrincipalType *ServicePrincipalType
    // The tenantId property
    tenantId *string
}
// NewServicePrincipalEvidence instantiates a new ServicePrincipalEvidence and sets the default values.
func NewServicePrincipalEvidence()(*ServicePrincipalEvidence) {
    m := &ServicePrincipalEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.servicePrincipalEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServicePrincipalEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServicePrincipalEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalEvidence(), nil
}
// GetAppId gets the appId property value. The appId property
// returns a *string when successful
func (m *ServicePrincipalEvidence) GetAppId()(*string) {
    return m.appId
}
// GetAppOwnerTenantId gets the appOwnerTenantId property value. The appOwnerTenantId property
// returns a *string when successful
func (m *ServicePrincipalEvidence) GetAppOwnerTenantId()(*string) {
    return m.appOwnerTenantId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServicePrincipalEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["appOwnerTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppOwnerTenantId(val)
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
    res["servicePrincipalObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalObjectId(val)
        }
        return nil
    }
    res["servicePrincipalType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServicePrincipalType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalType(val.(*ServicePrincipalType))
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetServicePrincipalName gets the servicePrincipalName property value. The servicePrincipalName property
// returns a *string when successful
func (m *ServicePrincipalEvidence) GetServicePrincipalName()(*string) {
    return m.servicePrincipalName
}
// GetServicePrincipalObjectId gets the servicePrincipalObjectId property value. The servicePrincipalObjectId property
// returns a *string when successful
func (m *ServicePrincipalEvidence) GetServicePrincipalObjectId()(*string) {
    return m.servicePrincipalObjectId
}
// GetServicePrincipalType gets the servicePrincipalType property value. The servicePrincipalType property
// returns a *ServicePrincipalType when successful
func (m *ServicePrincipalEvidence) GetServicePrincipalType()(*ServicePrincipalType) {
    return m.servicePrincipalType
}
// GetTenantId gets the tenantId property value. The tenantId property
// returns a *string when successful
func (m *ServicePrincipalEvidence) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ServicePrincipalEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appOwnerTenantId", m.GetAppOwnerTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalName", m.GetServicePrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalObjectId", m.GetServicePrincipalObjectId())
        if err != nil {
            return err
        }
    }
    if m.GetServicePrincipalType() != nil {
        cast := (*m.GetServicePrincipalType()).String()
        err = writer.WriteStringValue("servicePrincipalType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The appId property
func (m *ServicePrincipalEvidence) SetAppId(value *string)() {
    m.appId = value
}
// SetAppOwnerTenantId sets the appOwnerTenantId property value. The appOwnerTenantId property
func (m *ServicePrincipalEvidence) SetAppOwnerTenantId(value *string)() {
    m.appOwnerTenantId = value
}
// SetServicePrincipalName sets the servicePrincipalName property value. The servicePrincipalName property
func (m *ServicePrincipalEvidence) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
// SetServicePrincipalObjectId sets the servicePrincipalObjectId property value. The servicePrincipalObjectId property
func (m *ServicePrincipalEvidence) SetServicePrincipalObjectId(value *string)() {
    m.servicePrincipalObjectId = value
}
// SetServicePrincipalType sets the servicePrincipalType property value. The servicePrincipalType property
func (m *ServicePrincipalEvidence) SetServicePrincipalType(value *ServicePrincipalType)() {
    m.servicePrincipalType = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ServicePrincipalEvidence) SetTenantId(value *string)() {
    m.tenantId = value
}
type ServicePrincipalEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetAppOwnerTenantId()(*string)
    GetServicePrincipalName()(*string)
    GetServicePrincipalObjectId()(*string)
    GetServicePrincipalType()(*ServicePrincipalType)
    GetTenantId()(*string)
    SetAppId(value *string)()
    SetAppOwnerTenantId(value *string)()
    SetServicePrincipalName(value *string)()
    SetServicePrincipalObjectId(value *string)()
    SetServicePrincipalType(value *ServicePrincipalType)()
    SetTenantId(value *string)()
}
