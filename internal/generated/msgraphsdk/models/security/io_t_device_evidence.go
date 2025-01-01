package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IoTDeviceEvidence struct {
    AlertEvidence
    // The device ID.
    deviceId *string
    // The friendly name of the device.
    deviceName *string
    // The URL to the device page in the IoT Defender portal.
    devicePageLink *string
    // The device subtype.
    deviceSubType *string
    // The type of the device. For example, 'temperature sensor,' 'freezer,' 'wind turbine,' and so on.
    deviceType *string
    // The importance level for the IoT device. Possible values are low, normal, high, and unknownFutureValue.
    importance *IoTDeviceImportanceType
    // The azureResourceEvidence entity that represents the IoT Hub that the device belongs to.
    ioTHub AzureResourceEvidenceable
    // The ID of the Azure Security Center for the IoT agent that is running on the device.
    ioTSecurityAgentId *string
    // The current IP address of the device.
    ipAddress IpEvidenceable
    // Indicates whether the device classified as an authorized device.
    isAuthorized *bool
    // Indicates whether the device classified as a programming device.
    isProgramming *bool
    // Indicates whether the device classified as a scanner.
    isScanner *bool
    // The MAC address of the device.
    macAddress *string
    // The manufacturer of the device.
    manufacturer *string
    // The model of the device.
    model *string
    // The current network interface controllers on the device.
    nics []NicEvidenceable
    // The operating system the device is running.
    operatingSystem *string
    // The owners for the device.
    owners []string
    // The list of protocols that the device supports.
    protocols []string
    // The Purdue Layer of the device.
    purdueLayer *string
    // The sensor that monitors the device.
    sensor *string
    // The serial number of the device.
    serialNumber *string
    // The site location of the device.
    site *string
    // The source (microsoft/vendor) of the device entity.
    source *string
    // A URL reference to the source item where the device is managed.
    sourceRef UrlEvidenceable
    // The zone location of the device within a site.
    zone *string
}
// NewIoTDeviceEvidence instantiates a new IoTDeviceEvidence and sets the default values.
func NewIoTDeviceEvidence()(*IoTDeviceEvidence) {
    m := &IoTDeviceEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.ioTDeviceEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIoTDeviceEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIoTDeviceEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIoTDeviceEvidence(), nil
}
// GetDeviceId gets the deviceId property value. The device ID.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. The friendly name of the device.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDevicePageLink gets the devicePageLink property value. The URL to the device page in the IoT Defender portal.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetDevicePageLink()(*string) {
    return m.devicePageLink
}
// GetDeviceSubType gets the deviceSubType property value. The device subtype.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetDeviceSubType()(*string) {
    return m.deviceSubType
}
// GetDeviceType gets the deviceType property value. The type of the device. For example, 'temperature sensor,' 'freezer,' 'wind turbine,' and so on.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetDeviceType()(*string) {
    return m.deviceType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IoTDeviceEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["devicePageLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePageLink(val)
        }
        return nil
    }
    res["deviceSubType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceSubType(val)
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val)
        }
        return nil
    }
    res["importance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIoTDeviceImportanceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*IoTDeviceImportanceType))
        }
        return nil
    }
    res["ioTHub"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureResourceEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIoTHub(val.(AzureResourceEvidenceable))
        }
        return nil
    }
    res["ioTSecurityAgentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIoTSecurityAgentId(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIpEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val.(IpEvidenceable))
        }
        return nil
    }
    res["isAuthorized"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAuthorized(val)
        }
        return nil
    }
    res["isProgramming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsProgramming(val)
        }
        return nil
    }
    res["isScanner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScanner(val)
        }
        return nil
    }
    res["macAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddress(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["nics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNicEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NicEvidenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NicEvidenceable)
                }
            }
            m.SetNics(res)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetOwners(res)
        }
        return nil
    }
    res["protocols"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetProtocols(res)
        }
        return nil
    }
    res["purdueLayer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurdueLayer(val)
        }
        return nil
    }
    res["sensor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensor(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["site"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSite(val)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    res["sourceRef"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUrlEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceRef(val.(UrlEvidenceable))
        }
        return nil
    }
    res["zone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZone(val)
        }
        return nil
    }
    return res
}
// GetImportance gets the importance property value. The importance level for the IoT device. Possible values are low, normal, high, and unknownFutureValue.
// returns a *IoTDeviceImportanceType when successful
func (m *IoTDeviceEvidence) GetImportance()(*IoTDeviceImportanceType) {
    return m.importance
}
// GetIoTHub gets the ioTHub property value. The azureResourceEvidence entity that represents the IoT Hub that the device belongs to.
// returns a AzureResourceEvidenceable when successful
func (m *IoTDeviceEvidence) GetIoTHub()(AzureResourceEvidenceable) {
    return m.ioTHub
}
// GetIoTSecurityAgentId gets the ioTSecurityAgentId property value. The ID of the Azure Security Center for the IoT agent that is running on the device.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetIoTSecurityAgentId()(*string) {
    return m.ioTSecurityAgentId
}
// GetIpAddress gets the ipAddress property value. The current IP address of the device.
// returns a IpEvidenceable when successful
func (m *IoTDeviceEvidence) GetIpAddress()(IpEvidenceable) {
    return m.ipAddress
}
// GetIsAuthorized gets the isAuthorized property value. Indicates whether the device classified as an authorized device.
// returns a *bool when successful
func (m *IoTDeviceEvidence) GetIsAuthorized()(*bool) {
    return m.isAuthorized
}
// GetIsProgramming gets the isProgramming property value. Indicates whether the device classified as a programming device.
// returns a *bool when successful
func (m *IoTDeviceEvidence) GetIsProgramming()(*bool) {
    return m.isProgramming
}
// GetIsScanner gets the isScanner property value. Indicates whether the device classified as a scanner.
// returns a *bool when successful
func (m *IoTDeviceEvidence) GetIsScanner()(*bool) {
    return m.isScanner
}
// GetMacAddress gets the macAddress property value. The MAC address of the device.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetMacAddress()(*string) {
    return m.macAddress
}
// GetManufacturer gets the manufacturer property value. The manufacturer of the device.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The model of the device.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetModel()(*string) {
    return m.model
}
// GetNics gets the nics property value. The current network interface controllers on the device.
// returns a []NicEvidenceable when successful
func (m *IoTDeviceEvidence) GetNics()([]NicEvidenceable) {
    return m.nics
}
// GetOperatingSystem gets the operatingSystem property value. The operating system the device is running.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOwners gets the owners property value. The owners for the device.
// returns a []string when successful
func (m *IoTDeviceEvidence) GetOwners()([]string) {
    return m.owners
}
// GetProtocols gets the protocols property value. The list of protocols that the device supports.
// returns a []string when successful
func (m *IoTDeviceEvidence) GetProtocols()([]string) {
    return m.protocols
}
// GetPurdueLayer gets the purdueLayer property value. The Purdue Layer of the device.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetPurdueLayer()(*string) {
    return m.purdueLayer
}
// GetSensor gets the sensor property value. The sensor that monitors the device.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetSensor()(*string) {
    return m.sensor
}
// GetSerialNumber gets the serialNumber property value. The serial number of the device.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSite gets the site property value. The site location of the device.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetSite()(*string) {
    return m.site
}
// GetSource gets the source property value. The source (microsoft/vendor) of the device entity.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetSource()(*string) {
    return m.source
}
// GetSourceRef gets the sourceRef property value. A URL reference to the source item where the device is managed.
// returns a UrlEvidenceable when successful
func (m *IoTDeviceEvidence) GetSourceRef()(UrlEvidenceable) {
    return m.sourceRef
}
// GetZone gets the zone property value. The zone location of the device within a site.
// returns a *string when successful
func (m *IoTDeviceEvidence) GetZone()(*string) {
    return m.zone
}
// Serialize serializes information the current object
func (m *IoTDeviceEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("devicePageLink", m.GetDevicePageLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceSubType", m.GetDeviceSubType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceType", m.GetDeviceType())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ioTHub", m.GetIoTHub())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ioTSecurityAgentId", m.GetIoTSecurityAgentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAuthorized", m.GetIsAuthorized())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isProgramming", m.GetIsProgramming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isScanner", m.GetIsScanner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    if m.GetNics() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNics()))
        for i, v := range m.GetNics() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("nics", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    if m.GetOwners() != nil {
        err = writer.WriteCollectionOfStringValues("owners", m.GetOwners())
        if err != nil {
            return err
        }
    }
    if m.GetProtocols() != nil {
        err = writer.WriteCollectionOfStringValues("protocols", m.GetProtocols())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("purdueLayer", m.GetPurdueLayer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sensor", m.GetSensor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("site", m.GetSite())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceRef", m.GetSourceRef())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("zone", m.GetZone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceId sets the deviceId property value. The device ID.
func (m *IoTDeviceEvidence) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. The friendly name of the device.
func (m *IoTDeviceEvidence) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDevicePageLink sets the devicePageLink property value. The URL to the device page in the IoT Defender portal.
func (m *IoTDeviceEvidence) SetDevicePageLink(value *string)() {
    m.devicePageLink = value
}
// SetDeviceSubType sets the deviceSubType property value. The device subtype.
func (m *IoTDeviceEvidence) SetDeviceSubType(value *string)() {
    m.deviceSubType = value
}
// SetDeviceType sets the deviceType property value. The type of the device. For example, 'temperature sensor,' 'freezer,' 'wind turbine,' and so on.
func (m *IoTDeviceEvidence) SetDeviceType(value *string)() {
    m.deviceType = value
}
// SetImportance sets the importance property value. The importance level for the IoT device. Possible values are low, normal, high, and unknownFutureValue.
func (m *IoTDeviceEvidence) SetImportance(value *IoTDeviceImportanceType)() {
    m.importance = value
}
// SetIoTHub sets the ioTHub property value. The azureResourceEvidence entity that represents the IoT Hub that the device belongs to.
func (m *IoTDeviceEvidence) SetIoTHub(value AzureResourceEvidenceable)() {
    m.ioTHub = value
}
// SetIoTSecurityAgentId sets the ioTSecurityAgentId property value. The ID of the Azure Security Center for the IoT agent that is running on the device.
func (m *IoTDeviceEvidence) SetIoTSecurityAgentId(value *string)() {
    m.ioTSecurityAgentId = value
}
// SetIpAddress sets the ipAddress property value. The current IP address of the device.
func (m *IoTDeviceEvidence) SetIpAddress(value IpEvidenceable)() {
    m.ipAddress = value
}
// SetIsAuthorized sets the isAuthorized property value. Indicates whether the device classified as an authorized device.
func (m *IoTDeviceEvidence) SetIsAuthorized(value *bool)() {
    m.isAuthorized = value
}
// SetIsProgramming sets the isProgramming property value. Indicates whether the device classified as a programming device.
func (m *IoTDeviceEvidence) SetIsProgramming(value *bool)() {
    m.isProgramming = value
}
// SetIsScanner sets the isScanner property value. Indicates whether the device classified as a scanner.
func (m *IoTDeviceEvidence) SetIsScanner(value *bool)() {
    m.isScanner = value
}
// SetMacAddress sets the macAddress property value. The MAC address of the device.
func (m *IoTDeviceEvidence) SetMacAddress(value *string)() {
    m.macAddress = value
}
// SetManufacturer sets the manufacturer property value. The manufacturer of the device.
func (m *IoTDeviceEvidence) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The model of the device.
func (m *IoTDeviceEvidence) SetModel(value *string)() {
    m.model = value
}
// SetNics sets the nics property value. The current network interface controllers on the device.
func (m *IoTDeviceEvidence) SetNics(value []NicEvidenceable)() {
    m.nics = value
}
// SetOperatingSystem sets the operatingSystem property value. The operating system the device is running.
func (m *IoTDeviceEvidence) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOwners sets the owners property value. The owners for the device.
func (m *IoTDeviceEvidence) SetOwners(value []string)() {
    m.owners = value
}
// SetProtocols sets the protocols property value. The list of protocols that the device supports.
func (m *IoTDeviceEvidence) SetProtocols(value []string)() {
    m.protocols = value
}
// SetPurdueLayer sets the purdueLayer property value. The Purdue Layer of the device.
func (m *IoTDeviceEvidence) SetPurdueLayer(value *string)() {
    m.purdueLayer = value
}
// SetSensor sets the sensor property value. The sensor that monitors the device.
func (m *IoTDeviceEvidence) SetSensor(value *string)() {
    m.sensor = value
}
// SetSerialNumber sets the serialNumber property value. The serial number of the device.
func (m *IoTDeviceEvidence) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSite sets the site property value. The site location of the device.
func (m *IoTDeviceEvidence) SetSite(value *string)() {
    m.site = value
}
// SetSource sets the source property value. The source (microsoft/vendor) of the device entity.
func (m *IoTDeviceEvidence) SetSource(value *string)() {
    m.source = value
}
// SetSourceRef sets the sourceRef property value. A URL reference to the source item where the device is managed.
func (m *IoTDeviceEvidence) SetSourceRef(value UrlEvidenceable)() {
    m.sourceRef = value
}
// SetZone sets the zone property value. The zone location of the device within a site.
func (m *IoTDeviceEvidence) SetZone(value *string)() {
    m.zone = value
}
type IoTDeviceEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDevicePageLink()(*string)
    GetDeviceSubType()(*string)
    GetDeviceType()(*string)
    GetImportance()(*IoTDeviceImportanceType)
    GetIoTHub()(AzureResourceEvidenceable)
    GetIoTSecurityAgentId()(*string)
    GetIpAddress()(IpEvidenceable)
    GetIsAuthorized()(*bool)
    GetIsProgramming()(*bool)
    GetIsScanner()(*bool)
    GetMacAddress()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetNics()([]NicEvidenceable)
    GetOperatingSystem()(*string)
    GetOwners()([]string)
    GetProtocols()([]string)
    GetPurdueLayer()(*string)
    GetSensor()(*string)
    GetSerialNumber()(*string)
    GetSite()(*string)
    GetSource()(*string)
    GetSourceRef()(UrlEvidenceable)
    GetZone()(*string)
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDevicePageLink(value *string)()
    SetDeviceSubType(value *string)()
    SetDeviceType(value *string)()
    SetImportance(value *IoTDeviceImportanceType)()
    SetIoTHub(value AzureResourceEvidenceable)()
    SetIoTSecurityAgentId(value *string)()
    SetIpAddress(value IpEvidenceable)()
    SetIsAuthorized(value *bool)()
    SetIsProgramming(value *bool)()
    SetIsScanner(value *bool)()
    SetMacAddress(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetNics(value []NicEvidenceable)()
    SetOperatingSystem(value *string)()
    SetOwners(value []string)()
    SetProtocols(value []string)()
    SetPurdueLayer(value *string)()
    SetSensor(value *string)()
    SetSerialNumber(value *string)()
    SetSite(value *string)()
    SetSource(value *string)()
    SetSourceRef(value UrlEvidenceable)()
    SetZone(value *string)()
}
