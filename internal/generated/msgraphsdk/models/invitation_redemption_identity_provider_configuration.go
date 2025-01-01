package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InvitationRedemptionIdentityProviderConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The fallback identity provider to be used in case no primary identity provider can be used for guest invitation redemption. Possible values are: defaultConfiguredIdp, emailOneTimePasscode, or microsoftAccount.
    fallbackIdentityProvider *B2bIdentityProvidersType
    // The OdataType property
    odataType *string
    // Collection of identity providers in priority order of preference to be used for guest invitation redemption. Possible values are: azureActiveDirectory, externalFederation, or socialIdentityProviders.
    primaryIdentityProviderPrecedenceOrder []B2bIdentityProvidersType
}
// NewInvitationRedemptionIdentityProviderConfiguration instantiates a new InvitationRedemptionIdentityProviderConfiguration and sets the default values.
func NewInvitationRedemptionIdentityProviderConfiguration()(*InvitationRedemptionIdentityProviderConfiguration) {
    m := &InvitationRedemptionIdentityProviderConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInvitationRedemptionIdentityProviderConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInvitationRedemptionIdentityProviderConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.defaultInvitationRedemptionIdentityProviderConfiguration":
                        return NewDefaultInvitationRedemptionIdentityProviderConfiguration(), nil
                }
            }
        }
    }
    return NewInvitationRedemptionIdentityProviderConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InvitationRedemptionIdentityProviderConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFallbackIdentityProvider gets the fallbackIdentityProvider property value. The fallback identity provider to be used in case no primary identity provider can be used for guest invitation redemption. Possible values are: defaultConfiguredIdp, emailOneTimePasscode, or microsoftAccount.
// returns a *B2bIdentityProvidersType when successful
func (m *InvitationRedemptionIdentityProviderConfiguration) GetFallbackIdentityProvider()(*B2bIdentityProvidersType) {
    return m.fallbackIdentityProvider
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InvitationRedemptionIdentityProviderConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fallbackIdentityProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseB2bIdentityProvidersType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFallbackIdentityProvider(val.(*B2bIdentityProvidersType))
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
    res["primaryIdentityProviderPrecedenceOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseB2bIdentityProvidersType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]B2bIdentityProvidersType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*B2bIdentityProvidersType))
                }
            }
            m.SetPrimaryIdentityProviderPrecedenceOrder(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *InvitationRedemptionIdentityProviderConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetPrimaryIdentityProviderPrecedenceOrder gets the primaryIdentityProviderPrecedenceOrder property value. Collection of identity providers in priority order of preference to be used for guest invitation redemption. Possible values are: azureActiveDirectory, externalFederation, or socialIdentityProviders.
// returns a []B2bIdentityProvidersType when successful
func (m *InvitationRedemptionIdentityProviderConfiguration) GetPrimaryIdentityProviderPrecedenceOrder()([]B2bIdentityProvidersType) {
    return m.primaryIdentityProviderPrecedenceOrder
}
// Serialize serializes information the current object
func (m *InvitationRedemptionIdentityProviderConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFallbackIdentityProvider() != nil {
        cast := (*m.GetFallbackIdentityProvider()).String()
        err := writer.WriteStringValue("fallbackIdentityProvider", &cast)
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
    if m.GetPrimaryIdentityProviderPrecedenceOrder() != nil {
        err := writer.WriteCollectionOfStringValues("primaryIdentityProviderPrecedenceOrder", SerializeB2bIdentityProvidersType(m.GetPrimaryIdentityProviderPrecedenceOrder()))
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
func (m *InvitationRedemptionIdentityProviderConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFallbackIdentityProvider sets the fallbackIdentityProvider property value. The fallback identity provider to be used in case no primary identity provider can be used for guest invitation redemption. Possible values are: defaultConfiguredIdp, emailOneTimePasscode, or microsoftAccount.
func (m *InvitationRedemptionIdentityProviderConfiguration) SetFallbackIdentityProvider(value *B2bIdentityProvidersType)() {
    m.fallbackIdentityProvider = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InvitationRedemptionIdentityProviderConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrimaryIdentityProviderPrecedenceOrder sets the primaryIdentityProviderPrecedenceOrder property value. Collection of identity providers in priority order of preference to be used for guest invitation redemption. Possible values are: azureActiveDirectory, externalFederation, or socialIdentityProviders.
func (m *InvitationRedemptionIdentityProviderConfiguration) SetPrimaryIdentityProviderPrecedenceOrder(value []B2bIdentityProvidersType)() {
    m.primaryIdentityProviderPrecedenceOrder = value
}
type InvitationRedemptionIdentityProviderConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFallbackIdentityProvider()(*B2bIdentityProvidersType)
    GetOdataType()(*string)
    GetPrimaryIdentityProviderPrecedenceOrder()([]B2bIdentityProvidersType)
    SetFallbackIdentityProvider(value *B2bIdentityProvidersType)()
    SetOdataType(value *string)()
    SetPrimaryIdentityProviderPrecedenceOrder(value []B2bIdentityProvidersType)()
}
