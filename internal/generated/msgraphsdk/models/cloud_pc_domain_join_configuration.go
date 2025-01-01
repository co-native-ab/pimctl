package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcDomainJoinConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies the method by which the provisioned Cloud PC joins Microsoft Entra ID. If you choose the hybridAzureADJoin type, only provide a value for the onPremisesConnectionId property and leave the regionName property empty. If you choose the azureADJoin type, provide a value for either the onPremisesConnectionId or the regionName property. Possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
    domainJoinType *CloudPcDomainJoinType
    // The OdataType property
    odataType *string
    // The Azure network connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Microsoft Entra joined. If you enter an onPremisesConnectionId, leave the regionName property empty.
    onPremisesConnectionId *string
    // The logical geographic group this region belongs to. Multiple regions can belong to one region group. A customer can select a regionGroup when they provision a Cloud PC, and the Cloud PC is put in one of the regions in the group based on resource status. For example, the Europe region group contains the Northern Europe and Western Europe regions. Possible values are: default, australia, canada, usCentral, usEast, usWest, france, germany, europeUnion, unitedKingdom, japan, asia, india, southAmerica, euap, usGovernment, usGovernmentDOD, unknownFutureValue, norway, switzerland, southKorea. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: norway, switzerland, southKorea. Read-only.
    regionGroup *CloudPcRegionGroup
    // The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. Within this region, the Windows 365 service creates and manages the underlying virtual network. This option is available only when the IT admin selects Microsoft Entra joined as the domain join type. If you enter a regionName, leave the onPremisesConnectionId property empty.
    regionName *string
}
// NewCloudPcDomainJoinConfiguration instantiates a new CloudPcDomainJoinConfiguration and sets the default values.
func NewCloudPcDomainJoinConfiguration()(*CloudPcDomainJoinConfiguration) {
    m := &CloudPcDomainJoinConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcDomainJoinConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcDomainJoinConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDomainJoinType gets the domainJoinType property value. Specifies the method by which the provisioned Cloud PC joins Microsoft Entra ID. If you choose the hybridAzureADJoin type, only provide a value for the onPremisesConnectionId property and leave the regionName property empty. If you choose the azureADJoin type, provide a value for either the onPremisesConnectionId or the regionName property. Possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
// returns a *CloudPcDomainJoinType when successful
func (m *CloudPcDomainJoinConfiguration) GetDomainJoinType()(*CloudPcDomainJoinType) {
    return m.domainJoinType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcDomainJoinConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["domainJoinType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDomainJoinType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainJoinType(val.(*CloudPcDomainJoinType))
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
    res["onPremisesConnectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionId(val)
        }
        return nil
    }
    res["regionGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcRegionGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionGroup(val.(*CloudPcRegionGroup))
        }
        return nil
    }
    res["regionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcDomainJoinConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetOnPremisesConnectionId gets the onPremisesConnectionId property value. The Azure network connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Microsoft Entra joined. If you enter an onPremisesConnectionId, leave the regionName property empty.
// returns a *string when successful
func (m *CloudPcDomainJoinConfiguration) GetOnPremisesConnectionId()(*string) {
    return m.onPremisesConnectionId
}
// GetRegionGroup gets the regionGroup property value. The logical geographic group this region belongs to. Multiple regions can belong to one region group. A customer can select a regionGroup when they provision a Cloud PC, and the Cloud PC is put in one of the regions in the group based on resource status. For example, the Europe region group contains the Northern Europe and Western Europe regions. Possible values are: default, australia, canada, usCentral, usEast, usWest, france, germany, europeUnion, unitedKingdom, japan, asia, india, southAmerica, euap, usGovernment, usGovernmentDOD, unknownFutureValue, norway, switzerland, southKorea. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: norway, switzerland, southKorea. Read-only.
// returns a *CloudPcRegionGroup when successful
func (m *CloudPcDomainJoinConfiguration) GetRegionGroup()(*CloudPcRegionGroup) {
    return m.regionGroup
}
// GetRegionName gets the regionName property value. The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. Within this region, the Windows 365 service creates and manages the underlying virtual network. This option is available only when the IT admin selects Microsoft Entra joined as the domain join type. If you enter a regionName, leave the onPremisesConnectionId property empty.
// returns a *string when successful
func (m *CloudPcDomainJoinConfiguration) GetRegionName()(*string) {
    return m.regionName
}
// Serialize serializes information the current object
func (m *CloudPcDomainJoinConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDomainJoinType() != nil {
        cast := (*m.GetDomainJoinType()).String()
        err := writer.WriteStringValue("domainJoinType", &cast)
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
        err := writer.WriteStringValue("onPremisesConnectionId", m.GetOnPremisesConnectionId())
        if err != nil {
            return err
        }
    }
    if m.GetRegionGroup() != nil {
        cast := (*m.GetRegionGroup()).String()
        err := writer.WriteStringValue("regionGroup", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("regionName", m.GetRegionName())
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
func (m *CloudPcDomainJoinConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDomainJoinType sets the domainJoinType property value. Specifies the method by which the provisioned Cloud PC joins Microsoft Entra ID. If you choose the hybridAzureADJoin type, only provide a value for the onPremisesConnectionId property and leave the regionName property empty. If you choose the azureADJoin type, provide a value for either the onPremisesConnectionId or the regionName property. Possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
func (m *CloudPcDomainJoinConfiguration) SetDomainJoinType(value *CloudPcDomainJoinType)() {
    m.domainJoinType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcDomainJoinConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOnPremisesConnectionId sets the onPremisesConnectionId property value. The Azure network connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Microsoft Entra joined. If you enter an onPremisesConnectionId, leave the regionName property empty.
func (m *CloudPcDomainJoinConfiguration) SetOnPremisesConnectionId(value *string)() {
    m.onPremisesConnectionId = value
}
// SetRegionGroup sets the regionGroup property value. The logical geographic group this region belongs to. Multiple regions can belong to one region group. A customer can select a regionGroup when they provision a Cloud PC, and the Cloud PC is put in one of the regions in the group based on resource status. For example, the Europe region group contains the Northern Europe and Western Europe regions. Possible values are: default, australia, canada, usCentral, usEast, usWest, france, germany, europeUnion, unitedKingdom, japan, asia, india, southAmerica, euap, usGovernment, usGovernmentDOD, unknownFutureValue, norway, switzerland, southKorea. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: norway, switzerland, southKorea. Read-only.
func (m *CloudPcDomainJoinConfiguration) SetRegionGroup(value *CloudPcRegionGroup)() {
    m.regionGroup = value
}
// SetRegionName sets the regionName property value. The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. Within this region, the Windows 365 service creates and manages the underlying virtual network. This option is available only when the IT admin selects Microsoft Entra joined as the domain join type. If you enter a regionName, leave the onPremisesConnectionId property empty.
func (m *CloudPcDomainJoinConfiguration) SetRegionName(value *string)() {
    m.regionName = value
}
type CloudPcDomainJoinConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDomainJoinType()(*CloudPcDomainJoinType)
    GetOdataType()(*string)
    GetOnPremisesConnectionId()(*string)
    GetRegionGroup()(*CloudPcRegionGroup)
    GetRegionName()(*string)
    SetDomainJoinType(value *CloudPcDomainJoinType)()
    SetOdataType(value *string)()
    SetOnPremisesConnectionId(value *string)()
    SetRegionGroup(value *CloudPcRegionGroup)()
    SetRegionName(value *string)()
}
