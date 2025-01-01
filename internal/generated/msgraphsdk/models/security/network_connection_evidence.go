package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NetworkConnectionEvidence struct {
    AlertEvidence
    // The destinationAddress property
    destinationAddress IpEvidenceable
    // The destinationPort property
    destinationPort *int32
    // The protocol property
    protocol *ProtocolType
    // The sourceAddress property
    sourceAddress IpEvidenceable
    // The sourcePort property
    sourcePort *int32
}
// NewNetworkConnectionEvidence instantiates a new NetworkConnectionEvidence and sets the default values.
func NewNetworkConnectionEvidence()(*NetworkConnectionEvidence) {
    m := &NetworkConnectionEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.networkConnectionEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateNetworkConnectionEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkConnectionEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkConnectionEvidence(), nil
}
// GetDestinationAddress gets the destinationAddress property value. The destinationAddress property
// returns a IpEvidenceable when successful
func (m *NetworkConnectionEvidence) GetDestinationAddress()(IpEvidenceable) {
    return m.destinationAddress
}
// GetDestinationPort gets the destinationPort property value. The destinationPort property
// returns a *int32 when successful
func (m *NetworkConnectionEvidence) GetDestinationPort()(*int32) {
    return m.destinationPort
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkConnectionEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["destinationAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIpEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationAddress(val.(IpEvidenceable))
        }
        return nil
    }
    res["destinationPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationPort(val)
        }
        return nil
    }
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtocolType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val.(*ProtocolType))
        }
        return nil
    }
    res["sourceAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIpEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceAddress(val.(IpEvidenceable))
        }
        return nil
    }
    res["sourcePort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourcePort(val)
        }
        return nil
    }
    return res
}
// GetProtocol gets the protocol property value. The protocol property
// returns a *ProtocolType when successful
func (m *NetworkConnectionEvidence) GetProtocol()(*ProtocolType) {
    return m.protocol
}
// GetSourceAddress gets the sourceAddress property value. The sourceAddress property
// returns a IpEvidenceable when successful
func (m *NetworkConnectionEvidence) GetSourceAddress()(IpEvidenceable) {
    return m.sourceAddress
}
// GetSourcePort gets the sourcePort property value. The sourcePort property
// returns a *int32 when successful
func (m *NetworkConnectionEvidence) GetSourcePort()(*int32) {
    return m.sourcePort
}
// Serialize serializes information the current object
func (m *NetworkConnectionEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("destinationAddress", m.GetDestinationAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("destinationPort", m.GetDestinationPort())
        if err != nil {
            return err
        }
    }
    if m.GetProtocol() != nil {
        cast := (*m.GetProtocol()).String()
        err = writer.WriteStringValue("protocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceAddress", m.GetSourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sourcePort", m.GetSourcePort())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDestinationAddress sets the destinationAddress property value. The destinationAddress property
func (m *NetworkConnectionEvidence) SetDestinationAddress(value IpEvidenceable)() {
    m.destinationAddress = value
}
// SetDestinationPort sets the destinationPort property value. The destinationPort property
func (m *NetworkConnectionEvidence) SetDestinationPort(value *int32)() {
    m.destinationPort = value
}
// SetProtocol sets the protocol property value. The protocol property
func (m *NetworkConnectionEvidence) SetProtocol(value *ProtocolType)() {
    m.protocol = value
}
// SetSourceAddress sets the sourceAddress property value. The sourceAddress property
func (m *NetworkConnectionEvidence) SetSourceAddress(value IpEvidenceable)() {
    m.sourceAddress = value
}
// SetSourcePort sets the sourcePort property value. The sourcePort property
func (m *NetworkConnectionEvidence) SetSourcePort(value *int32)() {
    m.sourcePort = value
}
type NetworkConnectionEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestinationAddress()(IpEvidenceable)
    GetDestinationPort()(*int32)
    GetProtocol()(*ProtocolType)
    GetSourceAddress()(IpEvidenceable)
    GetSourcePort()(*int32)
    SetDestinationAddress(value IpEvidenceable)()
    SetDestinationPort(value *int32)()
    SetProtocol(value *ProtocolType)()
    SetSourceAddress(value IpEvidenceable)()
    SetSourcePort(value *int32)()
}
