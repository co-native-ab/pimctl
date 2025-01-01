package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SasTokenEvidence struct {
    AlertEvidence
    // The allowedIpAddresses property
    allowedIpAddresses *string
    // The allowedResourceTypes property
    allowedResourceTypes []string
    // The allowedServices property
    allowedServices []string
    // The expiryDateTime property
    expiryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The permissions property
    permissions []string
    // The protocol property
    protocol *string
    // The signatureHash property
    signatureHash *string
    // The signedWith property
    signedWith *string
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The storageResource property
    storageResource AzureResourceEvidenceable
}
// NewSasTokenEvidence instantiates a new SasTokenEvidence and sets the default values.
func NewSasTokenEvidence()(*SasTokenEvidence) {
    m := &SasTokenEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.sasTokenEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSasTokenEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSasTokenEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSasTokenEvidence(), nil
}
// GetAllowedIpAddresses gets the allowedIpAddresses property value. The allowedIpAddresses property
// returns a *string when successful
func (m *SasTokenEvidence) GetAllowedIpAddresses()(*string) {
    return m.allowedIpAddresses
}
// GetAllowedResourceTypes gets the allowedResourceTypes property value. The allowedResourceTypes property
// returns a []string when successful
func (m *SasTokenEvidence) GetAllowedResourceTypes()([]string) {
    return m.allowedResourceTypes
}
// GetAllowedServices gets the allowedServices property value. The allowedServices property
// returns a []string when successful
func (m *SasTokenEvidence) GetAllowedServices()([]string) {
    return m.allowedServices
}
// GetExpiryDateTime gets the expiryDateTime property value. The expiryDateTime property
// returns a *Time when successful
func (m *SasTokenEvidence) GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiryDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SasTokenEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["allowedIpAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedIpAddresses(val)
        }
        return nil
    }
    res["allowedResourceTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAllowedResourceTypes(res)
        }
        return nil
    }
    res["allowedServices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAllowedServices(res)
        }
        return nil
    }
    res["expiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryDateTime(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPermissions(res)
        }
        return nil
    }
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val)
        }
        return nil
    }
    res["signatureHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignatureHash(val)
        }
        return nil
    }
    res["signedWith"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignedWith(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["storageResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureResourceEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageResource(val.(AzureResourceEvidenceable))
        }
        return nil
    }
    return res
}
// GetPermissions gets the permissions property value. The permissions property
// returns a []string when successful
func (m *SasTokenEvidence) GetPermissions()([]string) {
    return m.permissions
}
// GetProtocol gets the protocol property value. The protocol property
// returns a *string when successful
func (m *SasTokenEvidence) GetProtocol()(*string) {
    return m.protocol
}
// GetSignatureHash gets the signatureHash property value. The signatureHash property
// returns a *string when successful
func (m *SasTokenEvidence) GetSignatureHash()(*string) {
    return m.signatureHash
}
// GetSignedWith gets the signedWith property value. The signedWith property
// returns a *string when successful
func (m *SasTokenEvidence) GetSignedWith()(*string) {
    return m.signedWith
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
// returns a *Time when successful
func (m *SasTokenEvidence) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStorageResource gets the storageResource property value. The storageResource property
// returns a AzureResourceEvidenceable when successful
func (m *SasTokenEvidence) GetStorageResource()(AzureResourceEvidenceable) {
    return m.storageResource
}
// Serialize serializes information the current object
func (m *SasTokenEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("allowedIpAddresses", m.GetAllowedIpAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedResourceTypes() != nil {
        err = writer.WriteCollectionOfStringValues("allowedResourceTypes", m.GetAllowedResourceTypes())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedServices() != nil {
        err = writer.WriteCollectionOfStringValues("allowedServices", m.GetAllowedServices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expiryDateTime", m.GetExpiryDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPermissions() != nil {
        err = writer.WriteCollectionOfStringValues("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("protocol", m.GetProtocol())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signatureHash", m.GetSignatureHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signedWith", m.GetSignedWith())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("storageResource", m.GetStorageResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedIpAddresses sets the allowedIpAddresses property value. The allowedIpAddresses property
func (m *SasTokenEvidence) SetAllowedIpAddresses(value *string)() {
    m.allowedIpAddresses = value
}
// SetAllowedResourceTypes sets the allowedResourceTypes property value. The allowedResourceTypes property
func (m *SasTokenEvidence) SetAllowedResourceTypes(value []string)() {
    m.allowedResourceTypes = value
}
// SetAllowedServices sets the allowedServices property value. The allowedServices property
func (m *SasTokenEvidence) SetAllowedServices(value []string)() {
    m.allowedServices = value
}
// SetExpiryDateTime sets the expiryDateTime property value. The expiryDateTime property
func (m *SasTokenEvidence) SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiryDateTime = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *SasTokenEvidence) SetPermissions(value []string)() {
    m.permissions = value
}
// SetProtocol sets the protocol property value. The protocol property
func (m *SasTokenEvidence) SetProtocol(value *string)() {
    m.protocol = value
}
// SetSignatureHash sets the signatureHash property value. The signatureHash property
func (m *SasTokenEvidence) SetSignatureHash(value *string)() {
    m.signatureHash = value
}
// SetSignedWith sets the signedWith property value. The signedWith property
func (m *SasTokenEvidence) SetSignedWith(value *string)() {
    m.signedWith = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *SasTokenEvidence) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStorageResource sets the storageResource property value. The storageResource property
func (m *SasTokenEvidence) SetStorageResource(value AzureResourceEvidenceable)() {
    m.storageResource = value
}
type SasTokenEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedIpAddresses()(*string)
    GetAllowedResourceTypes()([]string)
    GetAllowedServices()([]string)
    GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPermissions()([]string)
    GetProtocol()(*string)
    GetSignatureHash()(*string)
    GetSignedWith()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStorageResource()(AzureResourceEvidenceable)
    SetAllowedIpAddresses(value *string)()
    SetAllowedResourceTypes(value []string)()
    SetAllowedServices(value []string)()
    SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPermissions(value []string)()
    SetProtocol(value *string)()
    SetSignatureHash(value *string)()
    SetSignedWith(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStorageResource(value AzureResourceEvidenceable)()
}
