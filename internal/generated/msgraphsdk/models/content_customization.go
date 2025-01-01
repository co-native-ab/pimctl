package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContentCustomization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Represents the content options of External Identities to be customized throughout the authentication flow for a tenant.
    attributeCollection []KeyValueable
    // A relative URL for the content options of External Identities to be customized throughout the authentication flow for a tenant.
    attributeCollectionRelativeUrl *string
    // The OdataType property
    odataType *string
    // Represents content options to customize during MFA proofup interruptions.
    registrationCampaign []KeyValueable
    // The relative URL of the content options to customize during MFA proofup interruptions.
    registrationCampaignRelativeUrl *string
}
// NewContentCustomization instantiates a new ContentCustomization and sets the default values.
func NewContentCustomization()(*ContentCustomization) {
    m := &ContentCustomization{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContentCustomizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContentCustomizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentCustomization(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContentCustomization) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttributeCollection gets the attributeCollection property value. Represents the content options of External Identities to be customized throughout the authentication flow for a tenant.
// returns a []KeyValueable when successful
func (m *ContentCustomization) GetAttributeCollection()([]KeyValueable) {
    return m.attributeCollection
}
// GetAttributeCollectionRelativeUrl gets the attributeCollectionRelativeUrl property value. A relative URL for the content options of External Identities to be customized throughout the authentication flow for a tenant.
// returns a *string when successful
func (m *ContentCustomization) GetAttributeCollectionRelativeUrl()(*string) {
    return m.attributeCollectionRelativeUrl
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContentCustomization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attributeCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValueable)
                }
            }
            m.SetAttributeCollection(res)
        }
        return nil
    }
    res["attributeCollectionRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeCollectionRelativeUrl(val)
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
    res["registrationCampaign"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValueable)
                }
            }
            m.SetRegistrationCampaign(res)
        }
        return nil
    }
    res["registrationCampaignRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationCampaignRelativeUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ContentCustomization) GetOdataType()(*string) {
    return m.odataType
}
// GetRegistrationCampaign gets the registrationCampaign property value. Represents content options to customize during MFA proofup interruptions.
// returns a []KeyValueable when successful
func (m *ContentCustomization) GetRegistrationCampaign()([]KeyValueable) {
    return m.registrationCampaign
}
// GetRegistrationCampaignRelativeUrl gets the registrationCampaignRelativeUrl property value. The relative URL of the content options to customize during MFA proofup interruptions.
// returns a *string when successful
func (m *ContentCustomization) GetRegistrationCampaignRelativeUrl()(*string) {
    return m.registrationCampaignRelativeUrl
}
// Serialize serializes information the current object
func (m *ContentCustomization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttributeCollection() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributeCollection()))
        for i, v := range m.GetAttributeCollection() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("attributeCollection", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("attributeCollectionRelativeUrl", m.GetAttributeCollectionRelativeUrl())
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
    if m.GetRegistrationCampaign() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegistrationCampaign()))
        for i, v := range m.GetRegistrationCampaign() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("registrationCampaign", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("registrationCampaignRelativeUrl", m.GetRegistrationCampaignRelativeUrl())
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
func (m *ContentCustomization) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttributeCollection sets the attributeCollection property value. Represents the content options of External Identities to be customized throughout the authentication flow for a tenant.
func (m *ContentCustomization) SetAttributeCollection(value []KeyValueable)() {
    m.attributeCollection = value
}
// SetAttributeCollectionRelativeUrl sets the attributeCollectionRelativeUrl property value. A relative URL for the content options of External Identities to be customized throughout the authentication flow for a tenant.
func (m *ContentCustomization) SetAttributeCollectionRelativeUrl(value *string)() {
    m.attributeCollectionRelativeUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContentCustomization) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRegistrationCampaign sets the registrationCampaign property value. Represents content options to customize during MFA proofup interruptions.
func (m *ContentCustomization) SetRegistrationCampaign(value []KeyValueable)() {
    m.registrationCampaign = value
}
// SetRegistrationCampaignRelativeUrl sets the registrationCampaignRelativeUrl property value. The relative URL of the content options to customize during MFA proofup interruptions.
func (m *ContentCustomization) SetRegistrationCampaignRelativeUrl(value *string)() {
    m.registrationCampaignRelativeUrl = value
}
type ContentCustomizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributeCollection()([]KeyValueable)
    GetAttributeCollectionRelativeUrl()(*string)
    GetOdataType()(*string)
    GetRegistrationCampaign()([]KeyValueable)
    GetRegistrationCampaignRelativeUrl()(*string)
    SetAttributeCollection(value []KeyValueable)()
    SetAttributeCollectionRelativeUrl(value *string)()
    SetOdataType(value *string)()
    SetRegistrationCampaign(value []KeyValueable)()
    SetRegistrationCampaignRelativeUrl(value *string)()
}
