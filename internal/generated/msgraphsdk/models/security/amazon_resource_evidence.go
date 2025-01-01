package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AmazonResourceEvidence struct {
    AlertEvidence
    // The unique identifier for the Amazon account.
    amazonAccountId *string
    // The Amazon resource identifier (ARN) for the cloud resource.
    amazonResourceId *string
    // The name of the resource.
    resourceName *string
    // The type of the resource.
    resourceType *string
}
// NewAmazonResourceEvidence instantiates a new AmazonResourceEvidence and sets the default values.
func NewAmazonResourceEvidence()(*AmazonResourceEvidence) {
    m := &AmazonResourceEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.amazonResourceEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAmazonResourceEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAmazonResourceEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAmazonResourceEvidence(), nil
}
// GetAmazonAccountId gets the amazonAccountId property value. The unique identifier for the Amazon account.
// returns a *string when successful
func (m *AmazonResourceEvidence) GetAmazonAccountId()(*string) {
    return m.amazonAccountId
}
// GetAmazonResourceId gets the amazonResourceId property value. The Amazon resource identifier (ARN) for the cloud resource.
// returns a *string when successful
func (m *AmazonResourceEvidence) GetAmazonResourceId()(*string) {
    return m.amazonResourceId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AmazonResourceEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["amazonAccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmazonAccountId(val)
        }
        return nil
    }
    res["amazonResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmazonResourceId(val)
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
// GetResourceName gets the resourceName property value. The name of the resource.
// returns a *string when successful
func (m *AmazonResourceEvidence) GetResourceName()(*string) {
    return m.resourceName
}
// GetResourceType gets the resourceType property value. The type of the resource.
// returns a *string when successful
func (m *AmazonResourceEvidence) GetResourceType()(*string) {
    return m.resourceType
}
// Serialize serializes information the current object
func (m *AmazonResourceEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("amazonAccountId", m.GetAmazonAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("amazonResourceId", m.GetAmazonResourceId())
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
// SetAmazonAccountId sets the amazonAccountId property value. The unique identifier for the Amazon account.
func (m *AmazonResourceEvidence) SetAmazonAccountId(value *string)() {
    m.amazonAccountId = value
}
// SetAmazonResourceId sets the amazonResourceId property value. The Amazon resource identifier (ARN) for the cloud resource.
func (m *AmazonResourceEvidence) SetAmazonResourceId(value *string)() {
    m.amazonResourceId = value
}
// SetResourceName sets the resourceName property value. The name of the resource.
func (m *AmazonResourceEvidence) SetResourceName(value *string)() {
    m.resourceName = value
}
// SetResourceType sets the resourceType property value. The type of the resource.
func (m *AmazonResourceEvidence) SetResourceType(value *string)() {
    m.resourceType = value
}
type AmazonResourceEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmazonAccountId()(*string)
    GetAmazonResourceId()(*string)
    GetResourceName()(*string)
    GetResourceType()(*string)
    SetAmazonAccountId(value *string)()
    SetAmazonResourceId(value *string)()
    SetResourceName(value *string)()
    SetResourceType(value *string)()
}
