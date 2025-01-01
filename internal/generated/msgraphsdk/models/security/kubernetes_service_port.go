package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KubernetesServicePort struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The application protocol for this port.
    appProtocol *string
    // The name of this port within the service.
    name *string
    // The port on each node on which this service is exposed when the type is either NodePort or LoadBalancer.
    nodePort *int32
    // The OdataType property
    odataType *string
    // The port that this service exposes.
    port *int32
    // The protocol name. Possible values are: udp, tcp, sctp, unknownFutureValue.
    protocol *ContainerPortProtocol
    // The name or number of the port to access on the pods targeted by the service. The port number must be in the range 1 to 65535. The name must be an IANASVCNAME.
    targetPort *string
}
// NewKubernetesServicePort instantiates a new KubernetesServicePort and sets the default values.
func NewKubernetesServicePort()(*KubernetesServicePort) {
    m := &KubernetesServicePort{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateKubernetesServicePortFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKubernetesServicePortFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKubernetesServicePort(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *KubernetesServicePort) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppProtocol gets the appProtocol property value. The application protocol for this port.
// returns a *string when successful
func (m *KubernetesServicePort) GetAppProtocol()(*string) {
    return m.appProtocol
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KubernetesServicePort) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppProtocol(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["nodePort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodePort(val)
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
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContainerPortProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val.(*ContainerPortProtocol))
        }
        return nil
    }
    res["targetPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetPort(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of this port within the service.
// returns a *string when successful
func (m *KubernetesServicePort) GetName()(*string) {
    return m.name
}
// GetNodePort gets the nodePort property value. The port on each node on which this service is exposed when the type is either NodePort or LoadBalancer.
// returns a *int32 when successful
func (m *KubernetesServicePort) GetNodePort()(*int32) {
    return m.nodePort
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *KubernetesServicePort) GetOdataType()(*string) {
    return m.odataType
}
// GetPort gets the port property value. The port that this service exposes.
// returns a *int32 when successful
func (m *KubernetesServicePort) GetPort()(*int32) {
    return m.port
}
// GetProtocol gets the protocol property value. The protocol name. Possible values are: udp, tcp, sctp, unknownFutureValue.
// returns a *ContainerPortProtocol when successful
func (m *KubernetesServicePort) GetProtocol()(*ContainerPortProtocol) {
    return m.protocol
}
// GetTargetPort gets the targetPort property value. The name or number of the port to access on the pods targeted by the service. The port number must be in the range 1 to 65535. The name must be an IANASVCNAME.
// returns a *string when successful
func (m *KubernetesServicePort) GetTargetPort()(*string) {
    return m.targetPort
}
// Serialize serializes information the current object
func (m *KubernetesServicePort) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appProtocol", m.GetAppProtocol())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("nodePort", m.GetNodePort())
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
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    if m.GetProtocol() != nil {
        cast := (*m.GetProtocol()).String()
        err := writer.WriteStringValue("protocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetPort", m.GetTargetPort())
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
func (m *KubernetesServicePort) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppProtocol sets the appProtocol property value. The application protocol for this port.
func (m *KubernetesServicePort) SetAppProtocol(value *string)() {
    m.appProtocol = value
}
// SetName sets the name property value. The name of this port within the service.
func (m *KubernetesServicePort) SetName(value *string)() {
    m.name = value
}
// SetNodePort sets the nodePort property value. The port on each node on which this service is exposed when the type is either NodePort or LoadBalancer.
func (m *KubernetesServicePort) SetNodePort(value *int32)() {
    m.nodePort = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *KubernetesServicePort) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPort sets the port property value. The port that this service exposes.
func (m *KubernetesServicePort) SetPort(value *int32)() {
    m.port = value
}
// SetProtocol sets the protocol property value. The protocol name. Possible values are: udp, tcp, sctp, unknownFutureValue.
func (m *KubernetesServicePort) SetProtocol(value *ContainerPortProtocol)() {
    m.protocol = value
}
// SetTargetPort sets the targetPort property value. The name or number of the port to access on the pods targeted by the service. The port number must be in the range 1 to 65535. The name must be an IANASVCNAME.
func (m *KubernetesServicePort) SetTargetPort(value *string)() {
    m.targetPort = value
}
type KubernetesServicePortable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppProtocol()(*string)
    GetName()(*string)
    GetNodePort()(*int32)
    GetOdataType()(*string)
    GetPort()(*int32)
    GetProtocol()(*ContainerPortProtocol)
    GetTargetPort()(*string)
    SetAppProtocol(value *string)()
    SetName(value *string)()
    SetNodePort(value *int32)()
    SetOdataType(value *string)()
    SetPort(value *int32)()
    SetProtocol(value *ContainerPortProtocol)()
    SetTargetPort(value *string)()
}
