package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecurityVendorInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Specific provider (product/service - not vendor company); for example, WindowsDefenderATP.
    provider *string
    // Version of the provider or subprovider, if it exists, that generated the alert. Required
    providerVersion *string
    // Specific subprovider (under aggregating provider); for example, WindowsDefenderATP.SmartScreen.
    subProvider *string
    // Name of the alert vendor (for example, Microsoft, Dell, FireEye). Required
    vendorEscaped *string
}
// NewSecurityVendorInformation instantiates a new SecurityVendorInformation and sets the default values.
func NewSecurityVendorInformation()(*SecurityVendorInformation) {
    m := &SecurityVendorInformation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecurityVendorInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityVendorInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityVendorInformation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecurityVendorInformation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityVendorInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["provider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvider(val)
        }
        return nil
    }
    res["providerVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderVersion(val)
        }
        return nil
    }
    res["subProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubProvider(val)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorEscaped(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SecurityVendorInformation) GetOdataType()(*string) {
    return m.odataType
}
// GetProvider gets the provider property value. Specific provider (product/service - not vendor company); for example, WindowsDefenderATP.
// returns a *string when successful
func (m *SecurityVendorInformation) GetProvider()(*string) {
    return m.provider
}
// GetProviderVersion gets the providerVersion property value. Version of the provider or subprovider, if it exists, that generated the alert. Required
// returns a *string when successful
func (m *SecurityVendorInformation) GetProviderVersion()(*string) {
    return m.providerVersion
}
// GetSubProvider gets the subProvider property value. Specific subprovider (under aggregating provider); for example, WindowsDefenderATP.SmartScreen.
// returns a *string when successful
func (m *SecurityVendorInformation) GetSubProvider()(*string) {
    return m.subProvider
}
// GetVendorEscaped gets the vendor property value. Name of the alert vendor (for example, Microsoft, Dell, FireEye). Required
// returns a *string when successful
func (m *SecurityVendorInformation) GetVendorEscaped()(*string) {
    return m.vendorEscaped
}
// Serialize serializes information the current object
func (m *SecurityVendorInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("provider", m.GetProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("providerVersion", m.GetProviderVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subProvider", m.GetSubProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendor", m.GetVendorEscaped())
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
func (m *SecurityVendorInformation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SecurityVendorInformation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProvider sets the provider property value. Specific provider (product/service - not vendor company); for example, WindowsDefenderATP.
func (m *SecurityVendorInformation) SetProvider(value *string)() {
    m.provider = value
}
// SetProviderVersion sets the providerVersion property value. Version of the provider or subprovider, if it exists, that generated the alert. Required
func (m *SecurityVendorInformation) SetProviderVersion(value *string)() {
    m.providerVersion = value
}
// SetSubProvider sets the subProvider property value. Specific subprovider (under aggregating provider); for example, WindowsDefenderATP.SmartScreen.
func (m *SecurityVendorInformation) SetSubProvider(value *string)() {
    m.subProvider = value
}
// SetVendorEscaped sets the vendor property value. Name of the alert vendor (for example, Microsoft, Dell, FireEye). Required
func (m *SecurityVendorInformation) SetVendorEscaped(value *string)() {
    m.vendorEscaped = value
}
type SecurityVendorInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetProvider()(*string)
    GetProviderVersion()(*string)
    GetSubProvider()(*string)
    GetVendorEscaped()(*string)
    SetOdataType(value *string)()
    SetProvider(value *string)()
    SetProviderVersion(value *string)()
    SetSubProvider(value *string)()
    SetVendorEscaped(value *string)()
}
