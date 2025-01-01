package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContainerEvidence struct {
    AlertEvidence
    // The list of arguments.
    args []string
    // The list of commands.
    command []string
    // The container ID.
    containerId *string
    // The image used to run the container.
    image ContainerImageEvidenceable
    // The privileged status.
    isPrivileged *bool
    // The container name.
    name *string
    // The pod this container belongs to.
    pod KubernetesPodEvidenceable
}
// NewContainerEvidence instantiates a new ContainerEvidence and sets the default values.
func NewContainerEvidence()(*ContainerEvidence) {
    m := &ContainerEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.containerEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateContainerEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContainerEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContainerEvidence(), nil
}
// GetArgs gets the args property value. The list of arguments.
// returns a []string when successful
func (m *ContainerEvidence) GetArgs()([]string) {
    return m.args
}
// GetCommand gets the command property value. The list of commands.
// returns a []string when successful
func (m *ContainerEvidence) GetCommand()([]string) {
    return m.command
}
// GetContainerId gets the containerId property value. The container ID.
// returns a *string when successful
func (m *ContainerEvidence) GetContainerId()(*string) {
    return m.containerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContainerEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["args"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetArgs(res)
        }
        return nil
    }
    res["command"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCommand(res)
        }
        return nil
    }
    res["containerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerId(val)
        }
        return nil
    }
    res["image"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContainerImageEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImage(val.(ContainerImageEvidenceable))
        }
        return nil
    }
    res["isPrivileged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrivileged(val)
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
    res["pod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKubernetesPodEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPod(val.(KubernetesPodEvidenceable))
        }
        return nil
    }
    return res
}
// GetImage gets the image property value. The image used to run the container.
// returns a ContainerImageEvidenceable when successful
func (m *ContainerEvidence) GetImage()(ContainerImageEvidenceable) {
    return m.image
}
// GetIsPrivileged gets the isPrivileged property value. The privileged status.
// returns a *bool when successful
func (m *ContainerEvidence) GetIsPrivileged()(*bool) {
    return m.isPrivileged
}
// GetName gets the name property value. The container name.
// returns a *string when successful
func (m *ContainerEvidence) GetName()(*string) {
    return m.name
}
// GetPod gets the pod property value. The pod this container belongs to.
// returns a KubernetesPodEvidenceable when successful
func (m *ContainerEvidence) GetPod()(KubernetesPodEvidenceable) {
    return m.pod
}
// Serialize serializes information the current object
func (m *ContainerEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetArgs() != nil {
        err = writer.WriteCollectionOfStringValues("args", m.GetArgs())
        if err != nil {
            return err
        }
    }
    if m.GetCommand() != nil {
        err = writer.WriteCollectionOfStringValues("command", m.GetCommand())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("containerId", m.GetContainerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("image", m.GetImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPrivileged", m.GetIsPrivileged())
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
        err = writer.WriteObjectValue("pod", m.GetPod())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArgs sets the args property value. The list of arguments.
func (m *ContainerEvidence) SetArgs(value []string)() {
    m.args = value
}
// SetCommand sets the command property value. The list of commands.
func (m *ContainerEvidence) SetCommand(value []string)() {
    m.command = value
}
// SetContainerId sets the containerId property value. The container ID.
func (m *ContainerEvidence) SetContainerId(value *string)() {
    m.containerId = value
}
// SetImage sets the image property value. The image used to run the container.
func (m *ContainerEvidence) SetImage(value ContainerImageEvidenceable)() {
    m.image = value
}
// SetIsPrivileged sets the isPrivileged property value. The privileged status.
func (m *ContainerEvidence) SetIsPrivileged(value *bool)() {
    m.isPrivileged = value
}
// SetName sets the name property value. The container name.
func (m *ContainerEvidence) SetName(value *string)() {
    m.name = value
}
// SetPod sets the pod property value. The pod this container belongs to.
func (m *ContainerEvidence) SetPod(value KubernetesPodEvidenceable)() {
    m.pod = value
}
type ContainerEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArgs()([]string)
    GetCommand()([]string)
    GetContainerId()(*string)
    GetImage()(ContainerImageEvidenceable)
    GetIsPrivileged()(*bool)
    GetName()(*string)
    GetPod()(KubernetesPodEvidenceable)
    SetArgs(value []string)()
    SetCommand(value []string)()
    SetContainerId(value *string)()
    SetImage(value ContainerImageEvidenceable)()
    SetIsPrivileged(value *bool)()
    SetName(value *string)()
    SetPod(value KubernetesPodEvidenceable)()
}
