package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AzureResourceEvidence struct {
    AlertEvidence
    // The unique identifier for the Azure resource.
    resourceId *string
    // The name of the resource.
    resourceName *string
    // The type of the resource.
    resourceType *string
}
// NewAzureResourceEvidence instantiates a new AzureResourceEvidence and sets the default values.
func NewAzureResourceEvidence()(*AzureResourceEvidence) {
    m := &AzureResourceEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.azureResourceEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureResourceEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureResourceEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureResourceEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureResourceEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["resourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val)
        }
        return nil
    }
    return res
}
// GetResourceId gets the resourceId property value. The unique identifier for the Azure resource.
// returns a *string when successful
func (m *AzureResourceEvidence) GetResourceId()(*string) {
    return m.resourceId
}
// GetResourceName gets the resourceName property value. The name of the resource.
// returns a *string when successful
func (m *AzureResourceEvidence) GetResourceName()(*string) {
    return m.resourceName
}
// GetResourceType gets the resourceType property value. The type of the resource.
// returns a *string when successful
func (m *AzureResourceEvidence) GetResourceType()(*string) {
    return m.resourceType
}
// Serialize serializes information the current object
func (m *AzureResourceEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceType", m.GetResourceType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceId sets the resourceId property value. The unique identifier for the Azure resource.
func (m *AzureResourceEvidence) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetResourceName sets the resourceName property value. The name of the resource.
func (m *AzureResourceEvidence) SetResourceName(value *string)() {
    m.resourceName = value
}
// SetResourceType sets the resourceType property value. The type of the resource.
func (m *AzureResourceEvidence) SetResourceType(value *string)() {
    m.resourceType = value
}
type AzureResourceEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceId()(*string)
    GetResourceName()(*string)
    GetResourceType()(*string)
    SetResourceId(value *string)()
    SetResourceName(value *string)()
    SetResourceType(value *string)()
}
