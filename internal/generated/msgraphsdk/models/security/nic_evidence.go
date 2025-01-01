package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NicEvidence struct {
    AlertEvidence
    // The current IP address of the NIC.
    ipAddress IpEvidenceable
    // The MAC address of the NIC.
    macAddress *string
    // The current virtual local area networks of the NIC.
    vlans []string
}
// NewNicEvidence instantiates a new NicEvidence and sets the default values.
func NewNicEvidence()(*NicEvidence) {
    m := &NicEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.nicEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateNicEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNicEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNicEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NicEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
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
    res["vlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetVlans(res)
        }
        return nil
    }
    return res
}
// GetIpAddress gets the ipAddress property value. The current IP address of the NIC.
// returns a IpEvidenceable when successful
func (m *NicEvidence) GetIpAddress()(IpEvidenceable) {
    return m.ipAddress
}
// GetMacAddress gets the macAddress property value. The MAC address of the NIC.
// returns a *string when successful
func (m *NicEvidence) GetMacAddress()(*string) {
    return m.macAddress
}
// GetVlans gets the vlans property value. The current virtual local area networks of the NIC.
// returns a []string when successful
func (m *NicEvidence) GetVlans()([]string) {
    return m.vlans
}
// Serialize serializes information the current object
func (m *NicEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("ipAddress", m.GetIpAddress())
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
    if m.GetVlans() != nil {
        err = writer.WriteCollectionOfStringValues("vlans", m.GetVlans())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIpAddress sets the ipAddress property value. The current IP address of the NIC.
func (m *NicEvidence) SetIpAddress(value IpEvidenceable)() {
    m.ipAddress = value
}
// SetMacAddress sets the macAddress property value. The MAC address of the NIC.
func (m *NicEvidence) SetMacAddress(value *string)() {
    m.macAddress = value
}
// SetVlans sets the vlans property value. The current virtual local area networks of the NIC.
func (m *NicEvidence) SetVlans(value []string)() {
    m.vlans = value
}
type NicEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIpAddress()(IpEvidenceable)
    GetMacAddress()(*string)
    GetVlans()([]string)
    SetIpAddress(value IpEvidenceable)()
    SetMacAddress(value *string)()
    SetVlans(value []string)()
}
