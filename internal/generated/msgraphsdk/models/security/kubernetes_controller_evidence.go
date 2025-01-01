package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KubernetesControllerEvidence struct {
    AlertEvidence
    // The labels for the Kubernetes pod.
    labels Dictionaryable
    // The controller name.
    name *string
    // The service account namespace.
    namespace KubernetesNamespaceEvidenceable
    // The controller type.
    typeEscaped *string
}
// NewKubernetesControllerEvidence instantiates a new KubernetesControllerEvidence and sets the default values.
func NewKubernetesControllerEvidence()(*KubernetesControllerEvidence) {
    m := &KubernetesControllerEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.kubernetesControllerEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateKubernetesControllerEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKubernetesControllerEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKubernetesControllerEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KubernetesControllerEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. The labels for the Kubernetes pod.
// returns a Dictionaryable when successful
func (m *KubernetesControllerEvidence) GetLabels()(Dictionaryable) {
    return m.labels
}
// GetName gets the name property value. The controller name.
// returns a *string when successful
func (m *KubernetesControllerEvidence) GetName()(*string) {
    return m.name
}
// GetNamespace gets the namespace property value. The service account namespace.
// returns a KubernetesNamespaceEvidenceable when successful
func (m *KubernetesControllerEvidence) GetNamespace()(KubernetesNamespaceEvidenceable) {
    return m.namespace
}
// GetTypeEscaped gets the type property value. The controller type.
// returns a *string when successful
func (m *KubernetesControllerEvidence) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *KubernetesControllerEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLabels sets the labels property value. The labels for the Kubernetes pod.
func (m *KubernetesControllerEvidence) SetLabels(value Dictionaryable)() {
    m.labels = value
}
// SetName sets the name property value. The controller name.
func (m *KubernetesControllerEvidence) SetName(value *string)() {
    m.name = value
}
// SetNamespace sets the namespace property value. The service account namespace.
func (m *KubernetesControllerEvidence) SetNamespace(value KubernetesNamespaceEvidenceable)() {
    m.namespace = value
}
// SetTypeEscaped sets the type property value. The controller type.
func (m *KubernetesControllerEvidence) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type KubernetesControllerEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabels()(Dictionaryable)
    GetName()(*string)
    GetNamespace()(KubernetesNamespaceEvidenceable)
    GetTypeEscaped()(*string)
    SetLabels(value Dictionaryable)()
    SetName(value *string)()
    SetNamespace(value KubernetesNamespaceEvidenceable)()
    SetTypeEscaped(value *string)()
}
