package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KubernetesClusterEvidence struct {
    AlertEvidence
    // The cloud identifier of the cluster. Can be either an amazonResourceEvidence, azureResourceEvidence, or googleCloudResourceEvidence object.
    cloudResource AlertEvidenceable
    // The distribution type of the cluster.
    distribution *string
    // The cluster name.
    name *string
    // The platform the cluster runs on. Possible values are: unknown, aks, eks, gke, arc, unknownFutureValue.
    platform *KubernetesPlatform
    // The kubernetes version of the cluster.
    version *string
}
// NewKubernetesClusterEvidence instantiates a new KubernetesClusterEvidence and sets the default values.
func NewKubernetesClusterEvidence()(*KubernetesClusterEvidence) {
    m := &KubernetesClusterEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.kubernetesClusterEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateKubernetesClusterEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKubernetesClusterEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKubernetesClusterEvidence(), nil
}
// GetCloudResource gets the cloudResource property value. The cloud identifier of the cluster. Can be either an amazonResourceEvidence, azureResourceEvidence, or googleCloudResourceEvidence object.
// returns a AlertEvidenceable when successful
func (m *KubernetesClusterEvidence) GetCloudResource()(AlertEvidenceable) {
    return m.cloudResource
}
// GetDistribution gets the distribution property value. The distribution type of the cluster.
// returns a *string when successful
func (m *KubernetesClusterEvidence) GetDistribution()(*string) {
    return m.distribution
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KubernetesClusterEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["cloudResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlertEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudResource(val.(AlertEvidenceable))
        }
        return nil
    }
    res["distribution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDistribution(val)
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
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKubernetesPlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*KubernetesPlatform))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The cluster name.
// returns a *string when successful
func (m *KubernetesClusterEvidence) GetName()(*string) {
    return m.name
}
// GetPlatform gets the platform property value. The platform the cluster runs on. Possible values are: unknown, aks, eks, gke, arc, unknownFutureValue.
// returns a *KubernetesPlatform when successful
func (m *KubernetesClusterEvidence) GetPlatform()(*KubernetesPlatform) {
    return m.platform
}
// GetVersion gets the version property value. The kubernetes version of the cluster.
// returns a *string when successful
func (m *KubernetesClusterEvidence) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *KubernetesClusterEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("cloudResource", m.GetCloudResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("distribution", m.GetDistribution())
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
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCloudResource sets the cloudResource property value. The cloud identifier of the cluster. Can be either an amazonResourceEvidence, azureResourceEvidence, or googleCloudResourceEvidence object.
func (m *KubernetesClusterEvidence) SetCloudResource(value AlertEvidenceable)() {
    m.cloudResource = value
}
// SetDistribution sets the distribution property value. The distribution type of the cluster.
func (m *KubernetesClusterEvidence) SetDistribution(value *string)() {
    m.distribution = value
}
// SetName sets the name property value. The cluster name.
func (m *KubernetesClusterEvidence) SetName(value *string)() {
    m.name = value
}
// SetPlatform sets the platform property value. The platform the cluster runs on. Possible values are: unknown, aks, eks, gke, arc, unknownFutureValue.
func (m *KubernetesClusterEvidence) SetPlatform(value *KubernetesPlatform)() {
    m.platform = value
}
// SetVersion sets the version property value. The kubernetes version of the cluster.
func (m *KubernetesClusterEvidence) SetVersion(value *string)() {
    m.version = value
}
type KubernetesClusterEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudResource()(AlertEvidenceable)
    GetDistribution()(*string)
    GetName()(*string)
    GetPlatform()(*KubernetesPlatform)
    GetVersion()(*string)
    SetCloudResource(value AlertEvidenceable)()
    SetDistribution(value *string)()
    SetName(value *string)()
    SetPlatform(value *KubernetesPlatform)()
    SetVersion(value *string)()
}
