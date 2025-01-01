package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DelegatedAdminRelationshipCustomerParticipant struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display name of the customer tenant as set by Microsoft Entra ID. Read-only
    displayName *string
    // The OdataType property
    odataType *string
    // The Microsoft Entra ID-assigned tenant ID of the customer tenant.
    tenantId *string
}
// NewDelegatedAdminRelationshipCustomerParticipant instantiates a new DelegatedAdminRelationshipCustomerParticipant and sets the default values.
func NewDelegatedAdminRelationshipCustomerParticipant()(*DelegatedAdminRelationshipCustomerParticipant) {
    m := &DelegatedAdminRelationshipCustomerParticipant{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDelegatedAdminRelationshipCustomerParticipantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDelegatedAdminRelationshipCustomerParticipantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminRelationshipCustomerParticipant(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DelegatedAdminRelationshipCustomerParticipant) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The display name of the customer tenant as set by Microsoft Entra ID. Read-only
// returns a *string when successful
func (m *DelegatedAdminRelationshipCustomerParticipant) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DelegatedAdminRelationshipCustomerParticipant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DelegatedAdminRelationshipCustomerParticipant) GetOdataType()(*string) {
    return m.odataType
}
// GetTenantId gets the tenantId property value. The Microsoft Entra ID-assigned tenant ID of the customer tenant.
// returns a *string when successful
func (m *DelegatedAdminRelationshipCustomerParticipant) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *DelegatedAdminRelationshipCustomerParticipant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *DelegatedAdminRelationshipCustomerParticipant) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The display name of the customer tenant as set by Microsoft Entra ID. Read-only
func (m *DelegatedAdminRelationshipCustomerParticipant) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DelegatedAdminRelationshipCustomerParticipant) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTenantId sets the tenantId property value. The Microsoft Entra ID-assigned tenant ID of the customer tenant.
func (m *DelegatedAdminRelationshipCustomerParticipant) SetTenantId(value *string)() {
    m.tenantId = value
}
type DelegatedAdminRelationshipCustomerParticipantable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetOdataType()(*string)
    GetTenantId()(*string)
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
    SetTenantId(value *string)()
}
