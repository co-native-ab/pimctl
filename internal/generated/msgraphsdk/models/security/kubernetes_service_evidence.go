package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KubernetesServiceEvidence struct {
    AlertEvidence
    // The service cluster IP.
    clusterIP IpEvidenceable
    // The service external IPs.
    externalIPs []IpEvidenceable
    // The service labels.
    labels Dictionaryable
    // The service name.
    name *string
    // The service namespace.
    namespace KubernetesNamespaceEvidenceable
    // The service selector.
    selector Dictionaryable
    // The list of service ports.
    servicePorts []KubernetesServicePortable
    // The serviceType property
    serviceType *KubernetesServiceType
}
// NewKubernetesServiceEvidence instantiates a new KubernetesServiceEvidence and sets the default values.
func NewKubernetesServiceEvidence()(*KubernetesServiceEvidence) {
    m := &KubernetesServiceEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.kubernetesServiceEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateKubernetesServiceEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKubernetesServiceEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKubernetesServiceEvidence(), nil
}
// GetClusterIP gets the clusterIP property value. The service cluster IP.
// returns a IpEvidenceable when successful
func (m *KubernetesServiceEvidence) GetClusterIP()(IpEvidenceable) {
    return m.clusterIP
}
// GetExternalIPs gets the externalIPs property value. The service external IPs.
// returns a []IpEvidenceable when successful
func (m *KubernetesServiceEvidence) GetExternalIPs()([]IpEvidenceable) {
    return m.externalIPs
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KubernetesServiceEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["clusterIP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIpEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterIP(val.(IpEvidenceable))
        }
        return nil
    }
    res["externalIPs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpEvidenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IpEvidenceable)
                }
            }
            m.SetExternalIPs(res)
        }
        return nil
    }
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabels(val.(Dictionaryable))
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
    res["namespace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKubernetesNamespaceEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamespace(val.(KubernetesNamespaceEvidenceable))
        }
        return nil
    }
    res["selector"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelector(val.(Dictionaryable))
        }
        return nil
    }
    res["servicePorts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKubernetesServicePortFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KubernetesServicePortable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KubernetesServicePortable)
                }
            }
            m.SetServicePorts(res)
        }
        return nil
    }
    res["serviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKubernetesServiceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceType(val.(*KubernetesServiceType))
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. The service labels.
// returns a Dictionaryable when successful
func (m *KubernetesServiceEvidence) GetLabels()(Dictionaryable) {
    return m.labels
}
// GetName gets the name property value. The service name.
// returns a *string when successful
func (m *KubernetesServiceEvidence) GetName()(*string) {
    return m.name
}
// GetNamespace gets the namespace property value. The service namespace.
// returns a KubernetesNamespaceEvidenceable when successful
func (m *KubernetesServiceEvidence) GetNamespace()(KubernetesNamespaceEvidenceable) {
    return m.namespace
}
// GetSelector gets the selector property value. The service selector.
// returns a Dictionaryable when successful
func (m *KubernetesServiceEvidence) GetSelector()(Dictionaryable) {
    return m.selector
}
// GetServicePorts gets the servicePorts property value. The list of service ports.
// returns a []KubernetesServicePortable when successful
func (m *KubernetesServiceEvidence) GetServicePorts()([]KubernetesServicePortable) {
    return m.servicePorts
}
// GetServiceType gets the serviceType property value. The serviceType property
// returns a *KubernetesServiceType when successful
func (m *KubernetesServiceEvidence) GetServiceType()(*KubernetesServiceType) {
    return m.serviceType
}
// Serialize serializes information the current object
func (m *KubernetesServiceEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("clusterIP", m.GetClusterIP())
        if err != nil {
            return err
        }
    }
    if m.GetExternalIPs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExternalIPs()))
        for i, v := range m.GetExternalIPs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("externalIPs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("labels", m.GetLabels())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("namespace", m.GetNamespace())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("selector", m.GetSelector())
        if err != nil {
            return err
        }
    }
    if m.GetServicePorts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicePorts()))
        for i, v := range m.GetServicePorts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("servicePorts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetServiceType() != nil {
        cast := (*m.GetServiceType()).String()
        err = writer.WriteStringValue("serviceType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClusterIP sets the clusterIP property value. The service cluster IP.
func (m *KubernetesServiceEvidence) SetClusterIP(value IpEvidenceable)() {
    m.clusterIP = value
}
// SetExternalIPs sets the externalIPs property value. The service external IPs.
func (m *KubernetesServiceEvidence) SetExternalIPs(value []IpEvidenceable)() {
    m.externalIPs = value
}
// SetLabels sets the labels property value. The service labels.
func (m *KubernetesServiceEvidence) SetLabels(value Dictionaryable)() {
    m.labels = value
}
// SetName sets the name property value. The service name.
func (m *KubernetesServiceEvidence) SetName(value *string)() {
    m.name = value
}
// SetNamespace sets the namespace property value. The service namespace.
func (m *KubernetesServiceEvidence) SetNamespace(value KubernetesNamespaceEvidenceable)() {
    m.namespace = value
}
// SetSelector sets the selector property value. The service selector.
func (m *KubernetesServiceEvidence) SetSelector(value Dictionaryable)() {
    m.selector = value
}
// SetServicePorts sets the servicePorts property value. The list of service ports.
func (m *KubernetesServiceEvidence) SetServicePorts(value []KubernetesServicePortable)() {
    m.servicePorts = value
}
// SetServiceType sets the serviceType property value. The serviceType property
func (m *KubernetesServiceEvidence) SetServiceType(value *KubernetesServiceType)() {
    m.serviceType = value
}
type KubernetesServiceEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClusterIP()(IpEvidenceable)
    GetExternalIPs()([]IpEvidenceable)
    GetLabels()(Dictionaryable)
    GetName()(*string)
    GetNamespace()(KubernetesNamespaceEvidenceable)
    GetSelector()(Dictionaryable)
    GetServicePorts()([]KubernetesServicePortable)
    GetServiceType()(*KubernetesServiceType)
    SetClusterIP(value IpEvidenceable)()
    SetExternalIPs(value []IpEvidenceable)()
    SetLabels(value Dictionaryable)()
    SetName(value *string)()
    SetNamespace(value KubernetesNamespaceEvidenceable)()
    SetSelector(value Dictionaryable)()
    SetServicePorts(value []KubernetesServicePortable)()
    SetServiceType(value *KubernetesServiceType)()
}
