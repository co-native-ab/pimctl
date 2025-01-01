package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VmMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cloudProvider property
    cloudProvider *VmCloudProvider
    // The OdataType property
    odataType *string
    // Unique identifier of the Azure resource.
    resourceId *string
    // Unique identifier of the Azure subscription the customer tenant belongs to.
    subscriptionId *string
    // Unique identifier of the virtual machine instance.
    vmId *string
}
// NewVmMetadata instantiates a new VmMetadata and sets the default values.
func NewVmMetadata()(*VmMetadata) {
    m := &VmMetadata{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVmMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVmMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVmMetadata(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VmMetadata) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCloudProvider gets the cloudProvider property value. The cloudProvider property
// returns a *VmCloudProvider when successful
func (m *VmMetadata) GetCloudProvider()(*VmCloudProvider) {
    return m.cloudProvider
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VmMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVmCloudProvider)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudProvider(val.(*VmCloudProvider))
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
    res["vmId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *VmMetadata) GetOdataType()(*string) {
    return m.odataType
}
// GetResourceId gets the resourceId property value. Unique identifier of the Azure resource.
// returns a *string when successful
func (m *VmMetadata) GetResourceId()(*string) {
    return m.resourceId
}
// GetSubscriptionId gets the subscriptionId property value. Unique identifier of the Azure subscription the customer tenant belongs to.
// returns a *string when successful
func (m *VmMetadata) GetSubscriptionId()(*string) {
    return m.subscriptionId
}
// GetVmId gets the vmId property value. Unique identifier of the virtual machine instance.
// returns a *string when successful
func (m *VmMetadata) GetVmId()(*string) {
    return m.vmId
}
// Serialize serializes information the current object
func (m *VmMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCloudProvider() != nil {
        cast := (*m.GetCloudProvider()).String()
        err := writer.WriteStringValue("cloudProvider", &cast)
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
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vmId", m.GetVmId())
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
func (m *VmMetadata) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCloudProvider sets the cloudProvider property value. The cloudProvider property
func (m *VmMetadata) SetCloudProvider(value *VmCloudProvider)() {
    m.cloudProvider = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VmMetadata) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResourceId sets the resourceId property value. Unique identifier of the Azure resource.
func (m *VmMetadata) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetSubscriptionId sets the subscriptionId property value. Unique identifier of the Azure subscription the customer tenant belongs to.
func (m *VmMetadata) SetSubscriptionId(value *string)() {
    m.subscriptionId = value
}
// SetVmId sets the vmId property value. Unique identifier of the virtual machine instance.
func (m *VmMetadata) SetVmId(value *string)() {
    m.vmId = value
}
type VmMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudProvider()(*VmCloudProvider)
    GetOdataType()(*string)
    GetResourceId()(*string)
    GetSubscriptionId()(*string)
    GetVmId()(*string)
    SetCloudProvider(value *VmCloudProvider)()
    SetOdataType(value *string)()
    SetResourceId(value *string)()
    SetSubscriptionId(value *string)()
    SetVmId(value *string)()
}
