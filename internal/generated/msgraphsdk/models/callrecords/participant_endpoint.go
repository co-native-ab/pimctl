package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type ParticipantEndpoint struct {
    Endpoint
    // Identity associated with the endpoint.
    associatedIdentity ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Identityable
    // CPU number of cores used by the media endpoint.
    cpuCoresCount *int32
    // CPU name used by the media endpoint.
    cpuName *string
    // CPU processor speed used by the media endpoint.
    cpuProcessorSpeedInMhz *int32
    // The feedback provided by the user of this endpoint about the quality of the session.
    feedback UserFeedbackable
    // Identity associated with the endpoint. The identity property is deprecated and will stop returning data on June 30, 2026. Going forward, use the associatedIdentity property.
    identity ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable
    // Name of the device used by the media endpoint.
    name *string
}
// NewParticipantEndpoint instantiates a new ParticipantEndpoint and sets the default values.
func NewParticipantEndpoint()(*ParticipantEndpoint) {
    m := &ParticipantEndpoint{
        Endpoint: *NewEndpoint(),
    }
    odataTypeValue := "#microsoft.graph.callRecords.participantEndpoint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateParticipantEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParticipantEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipantEndpoint(), nil
}
// GetAssociatedIdentity gets the associatedIdentity property value. Identity associated with the endpoint.
// returns a Identityable when successful
func (m *ParticipantEndpoint) GetAssociatedIdentity()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Identityable) {
    return m.associatedIdentity
}
// GetCpuCoresCount gets the cpuCoresCount property value. CPU number of cores used by the media endpoint.
// returns a *int32 when successful
func (m *ParticipantEndpoint) GetCpuCoresCount()(*int32) {
    return m.cpuCoresCount
}
// GetCpuName gets the cpuName property value. CPU name used by the media endpoint.
// returns a *string when successful
func (m *ParticipantEndpoint) GetCpuName()(*string) {
    return m.cpuName
}
// GetCpuProcessorSpeedInMhz gets the cpuProcessorSpeedInMhz property value. CPU processor speed used by the media endpoint.
// returns a *int32 when successful
func (m *ParticipantEndpoint) GetCpuProcessorSpeedInMhz()(*int32) {
    return m.cpuProcessorSpeedInMhz
}
// GetFeedback gets the feedback property value. The feedback provided by the user of this endpoint about the quality of the session.
// returns a UserFeedbackable when successful
func (m *ParticipantEndpoint) GetFeedback()(UserFeedbackable) {
    return m.feedback
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParticipantEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Endpoint.GetFieldDeserializers()
    res["associatedIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssociatedIdentity(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Identityable))
        }
        return nil
    }
    res["cpuCoresCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuCoresCount(val)
        }
        return nil
    }
    res["cpuName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuName(val)
        }
        return nil
    }
    res["cpuProcessorSpeedInMhz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuProcessorSpeedInMhz(val)
        }
        return nil
    }
    res["feedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFeedbackFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedback(val.(UserFeedbackable))
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable))
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
    return res
}
// GetIdentity gets the identity property value. Identity associated with the endpoint. The identity property is deprecated and will stop returning data on June 30, 2026. Going forward, use the associatedIdentity property.
// returns a IdentitySetable when successful
func (m *ParticipantEndpoint) GetIdentity()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable) {
    return m.identity
}
// GetName gets the name property value. Name of the device used by the media endpoint.
// returns a *string when successful
func (m *ParticipantEndpoint) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ParticipantEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Endpoint.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("associatedIdentity", m.GetAssociatedIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cpuCoresCount", m.GetCpuCoresCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cpuName", m.GetCpuName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cpuProcessorSpeedInMhz", m.GetCpuProcessorSpeedInMhz())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("feedback", m.GetFeedback())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identity", m.GetIdentity())
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
    return nil
}
// SetAssociatedIdentity sets the associatedIdentity property value. Identity associated with the endpoint.
func (m *ParticipantEndpoint) SetAssociatedIdentity(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Identityable)() {
    m.associatedIdentity = value
}
// SetCpuCoresCount sets the cpuCoresCount property value. CPU number of cores used by the media endpoint.
func (m *ParticipantEndpoint) SetCpuCoresCount(value *int32)() {
    m.cpuCoresCount = value
}
// SetCpuName sets the cpuName property value. CPU name used by the media endpoint.
func (m *ParticipantEndpoint) SetCpuName(value *string)() {
    m.cpuName = value
}
// SetCpuProcessorSpeedInMhz sets the cpuProcessorSpeedInMhz property value. CPU processor speed used by the media endpoint.
func (m *ParticipantEndpoint) SetCpuProcessorSpeedInMhz(value *int32)() {
    m.cpuProcessorSpeedInMhz = value
}
// SetFeedback sets the feedback property value. The feedback provided by the user of this endpoint about the quality of the session.
func (m *ParticipantEndpoint) SetFeedback(value UserFeedbackable)() {
    m.feedback = value
}
// SetIdentity sets the identity property value. Identity associated with the endpoint. The identity property is deprecated and will stop returning data on June 30, 2026. Going forward, use the associatedIdentity property.
func (m *ParticipantEndpoint) SetIdentity(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)() {
    m.identity = value
}
// SetName sets the name property value. Name of the device used by the media endpoint.
func (m *ParticipantEndpoint) SetName(value *string)() {
    m.name = value
}
type ParticipantEndpointable interface {
    Endpointable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssociatedIdentity()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Identityable)
    GetCpuCoresCount()(*int32)
    GetCpuName()(*string)
    GetCpuProcessorSpeedInMhz()(*int32)
    GetFeedback()(UserFeedbackable)
    GetIdentity()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)
    GetName()(*string)
    SetAssociatedIdentity(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Identityable)()
    SetCpuCoresCount(value *int32)()
    SetCpuName(value *string)()
    SetCpuProcessorSpeedInMhz(value *int32)()
    SetFeedback(value UserFeedbackable)()
    SetIdentity(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)()
    SetName(value *string)()
}
