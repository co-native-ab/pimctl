package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KubernetesPodEvidence struct {
    AlertEvidence
    // The list of pod containers which are not init or ephemeral containers.
    containers []ContainerEvidenceable
    // The pod controller.
    controller KubernetesControllerEvidenceable
    // The list of pod ephemeral containers.
    ephemeralContainers []ContainerEvidenceable
    // The list of pod init containers.
    initContainers []ContainerEvidenceable
    // The pod labels.
    labels Dictionaryable
    // The pod name.
    name *string
    // The pod namespace.
    namespace KubernetesNamespaceEvidenceable
    // The pod IP.
    podIp IpEvidenceable
    // The pod service account.
    serviceAccount KubernetesServiceAccountEvidenceable
}
// NewKubernetesPodEvidence instantiates a new KubernetesPodEvidence and sets the default values.
func NewKubernetesPodEvidence()(*KubernetesPodEvidence) {
    m := &KubernetesPodEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.kubernetesPodEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateKubernetesPodEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKubernetesPodEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKubernetesPodEvidence(), nil
}
// GetContainers gets the containers property value. The list of pod containers which are not init or ephemeral containers.
// returns a []ContainerEvidenceable when successful
func (m *KubernetesPodEvidence) GetContainers()([]ContainerEvidenceable) {
    return m.containers
}
// GetController gets the controller property value. The pod controller.
// returns a KubernetesControllerEvidenceable when successful
func (m *KubernetesPodEvidence) GetController()(KubernetesControllerEvidenceable) {
    return m.controller
}
// GetEphemeralContainers gets the ephemeralContainers property value. The list of pod ephemeral containers.
// returns a []ContainerEvidenceable when successful
func (m *KubernetesPodEvidence) GetEphemeralContainers()([]ContainerEvidenceable) {
    return m.ephemeralContainers
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KubernetesPodEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["containers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContainerEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContainerEvidenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ContainerEvidenceable)
                }
            }
            m.SetContainers(res)
        }
        return nil
    }
    res["controller"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKubernetesControllerEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetController(val.(KubernetesControllerEvidenceable))
        }
        return nil
    }
    res["ephemeralContainers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContainerEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContainerEvidenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ContainerEvidenceable)
                }
            }
            m.SetEphemeralContainers(res)
        }
        return nil
    }
    res["initContainers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContainerEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContainerEvidenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ContainerEvidenceable)
                }
            }
            m.SetInitContainers(res)
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
    res["podIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIpEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPodIp(val.(IpEvidenceable))
        }
        return nil
    }
    res["serviceAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKubernetesServiceAccountEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccount(val.(KubernetesServiceAccountEvidenceable))
        }
        return nil
    }
    return res
}
// GetInitContainers gets the initContainers property value. The list of pod init containers.
// returns a []ContainerEvidenceable when successful
func (m *KubernetesPodEvidence) GetInitContainers()([]ContainerEvidenceable) {
    return m.initContainers
}
// GetLabels gets the labels property value. The pod labels.
// returns a Dictionaryable when successful
func (m *KubernetesPodEvidence) GetLabels()(Dictionaryable) {
    return m.labels
}
// GetName gets the name property value. The pod name.
// returns a *string when successful
func (m *KubernetesPodEvidence) GetName()(*string) {
    return m.name
}
// GetNamespace gets the namespace property value. The pod namespace.
// returns a KubernetesNamespaceEvidenceable when successful
func (m *KubernetesPodEvidence) GetNamespace()(KubernetesNamespaceEvidenceable) {
    return m.namespace
}
// GetPodIp gets the podIp property value. The pod IP.
// returns a IpEvidenceable when successful
func (m *KubernetesPodEvidence) GetPodIp()(IpEvidenceable) {
    return m.podIp
}
// GetServiceAccount gets the serviceAccount property value. The pod service account.
// returns a KubernetesServiceAccountEvidenceable when successful
func (m *KubernetesPodEvidence) GetServiceAccount()(KubernetesServiceAccountEvidenceable) {
    return m.serviceAccount
}
// Serialize serializes information the current object
func (m *KubernetesPodEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContainers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContainers()))
        for i, v := range m.GetContainers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("containers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("controller", m.GetController())
        if err != nil {
            return err
        }
    }
    if m.GetEphemeralContainers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEphemeralContainers()))
        for i, v := range m.GetEphemeralContainers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("ephemeralContainers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInitContainers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInitContainers()))
        for i, v := range m.GetInitContainers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("initContainers", cast)
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
        err = writer.WriteObjectValue("podIp", m.GetPodIp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("serviceAccount", m.GetServiceAccount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContainers sets the containers property value. The list of pod containers which are not init or ephemeral containers.
func (m *KubernetesPodEvidence) SetContainers(value []ContainerEvidenceable)() {
    m.containers = value
}
// SetController sets the controller property value. The pod controller.
func (m *KubernetesPodEvidence) SetController(value KubernetesControllerEvidenceable)() {
    m.controller = value
}
// SetEphemeralContainers sets the ephemeralContainers property value. The list of pod ephemeral containers.
func (m *KubernetesPodEvidence) SetEphemeralContainers(value []ContainerEvidenceable)() {
    m.ephemeralContainers = value
}
// SetInitContainers sets the initContainers property value. The list of pod init containers.
func (m *KubernetesPodEvidence) SetInitContainers(value []ContainerEvidenceable)() {
    m.initContainers = value
}
// SetLabels sets the labels property value. The pod labels.
func (m *KubernetesPodEvidence) SetLabels(value Dictionaryable)() {
    m.labels = value
}
// SetName sets the name property value. The pod name.
func (m *KubernetesPodEvidence) SetName(value *string)() {
    m.name = value
}
// SetNamespace sets the namespace property value. The pod namespace.
func (m *KubernetesPodEvidence) SetNamespace(value KubernetesNamespaceEvidenceable)() {
    m.namespace = value
}
// SetPodIp sets the podIp property value. The pod IP.
func (m *KubernetesPodEvidence) SetPodIp(value IpEvidenceable)() {
    m.podIp = value
}
// SetServiceAccount sets the serviceAccount property value. The pod service account.
func (m *KubernetesPodEvidence) SetServiceAccount(value KubernetesServiceAccountEvidenceable)() {
    m.serviceAccount = value
}
type KubernetesPodEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainers()([]ContainerEvidenceable)
    GetController()(KubernetesControllerEvidenceable)
    GetEphemeralContainers()([]ContainerEvidenceable)
    GetInitContainers()([]ContainerEvidenceable)
    GetLabels()(Dictionaryable)
    GetName()(*string)
    GetNamespace()(KubernetesNamespaceEvidenceable)
    GetPodIp()(IpEvidenceable)
    GetServiceAccount()(KubernetesServiceAccountEvidenceable)
    SetContainers(value []ContainerEvidenceable)()
    SetController(value KubernetesControllerEvidenceable)()
    SetEphemeralContainers(value []ContainerEvidenceable)()
    SetInitContainers(value []ContainerEvidenceable)()
    SetLabels(value Dictionaryable)()
    SetName(value *string)()
    SetNamespace(value KubernetesNamespaceEvidenceable)()
    SetPodIp(value IpEvidenceable)()
    SetServiceAccount(value KubernetesServiceAccountEvidenceable)()
}
