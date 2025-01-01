package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SensorSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Description of the sensor.
    description *string
    // DNS names for the domain controller
    domainControllerDnsNames []string
    // Indicates whether to delay updates for the sensor.
    isDelayedDeploymentEnabled *bool
    // The networkAdapters property
    networkAdapters []NetworkAdapterable
    // The OdataType property
    odataType *string
}
// NewSensorSettings instantiates a new SensorSettings and sets the default values.
func NewSensorSettings()(*SensorSettings) {
    m := &SensorSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSensorSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSensorSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensorSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SensorSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Description of the sensor.
// returns a *string when successful
func (m *SensorSettings) GetDescription()(*string) {
    return m.description
}
// GetDomainControllerDnsNames gets the domainControllerDnsNames property value. DNS names for the domain controller
// returns a []string when successful
func (m *SensorSettings) GetDomainControllerDnsNames()([]string) {
    return m.domainControllerDnsNames
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SensorSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["domainControllerDnsNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDomainControllerDnsNames(res)
        }
        return nil
    }
    res["isDelayedDeploymentEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDelayedDeploymentEnabled(val)
        }
        return nil
    }
    res["networkAdapters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNetworkAdapterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NetworkAdapterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NetworkAdapterable)
                }
            }
            m.SetNetworkAdapters(res)
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
    return res
}
// GetIsDelayedDeploymentEnabled gets the isDelayedDeploymentEnabled property value. Indicates whether to delay updates for the sensor.
// returns a *bool when successful
func (m *SensorSettings) GetIsDelayedDeploymentEnabled()(*bool) {
    return m.isDelayedDeploymentEnabled
}
// GetNetworkAdapters gets the networkAdapters property value. The networkAdapters property
// returns a []NetworkAdapterable when successful
func (m *SensorSettings) GetNetworkAdapters()([]NetworkAdapterable) {
    return m.networkAdapters
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SensorSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SensorSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDomainControllerDnsNames() != nil {
        err := writer.WriteCollectionOfStringValues("domainControllerDnsNames", m.GetDomainControllerDnsNames())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDelayedDeploymentEnabled", m.GetIsDelayedDeploymentEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkAdapters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkAdapters()))
        for i, v := range m.GetNetworkAdapters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("networkAdapters", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensorSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Description of the sensor.
func (m *SensorSettings) SetDescription(value *string)() {
    m.description = value
}
// SetDomainControllerDnsNames sets the domainControllerDnsNames property value. DNS names for the domain controller
func (m *SensorSettings) SetDomainControllerDnsNames(value []string)() {
    m.domainControllerDnsNames = value
}
// SetIsDelayedDeploymentEnabled sets the isDelayedDeploymentEnabled property value. Indicates whether to delay updates for the sensor.
func (m *SensorSettings) SetIsDelayedDeploymentEnabled(value *bool)() {
    m.isDelayedDeploymentEnabled = value
}
// SetNetworkAdapters sets the networkAdapters property value. The networkAdapters property
func (m *SensorSettings) SetNetworkAdapters(value []NetworkAdapterable)() {
    m.networkAdapters = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SensorSettings) SetOdataType(value *string)() {
    m.odataType = value
}
type SensorSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDomainControllerDnsNames()([]string)
    GetIsDelayedDeploymentEnabled()(*bool)
    GetNetworkAdapters()([]NetworkAdapterable)
    GetOdataType()(*string)
    SetDescription(value *string)()
    SetDomainControllerDnsNames(value []string)()
    SetIsDelayedDeploymentEnabled(value *bool)()
    SetNetworkAdapters(value []NetworkAdapterable)()
    SetOdataType(value *string)()
}
