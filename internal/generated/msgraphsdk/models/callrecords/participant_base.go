package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type ParticipantBase struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // List of administrativeUnitInfo objects for the call participant.
    administrativeUnitInfos []AdministrativeUnitInfoable
    // The identity of the call participant.
    identity ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CommunicationsIdentitySetable
}
// NewParticipantBase instantiates a new ParticipantBase and sets the default values.
func NewParticipantBase()(*ParticipantBase) {
    m := &ParticipantBase{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateParticipantBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParticipantBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.callRecords.organizer":
                        return NewOrganizer(), nil
                    case "#microsoft.graph.callRecords.participant":
                        return NewParticipant(), nil
                }
            }
        }
    }
    return NewParticipantBase(), nil
}
// GetAdministrativeUnitInfos gets the administrativeUnitInfos property value. List of administrativeUnitInfo objects for the call participant.
// returns a []AdministrativeUnitInfoable when successful
func (m *ParticipantBase) GetAdministrativeUnitInfos()([]AdministrativeUnitInfoable) {
    return m.administrativeUnitInfos
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParticipantBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnitInfos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAdministrativeUnitInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdministrativeUnitInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AdministrativeUnitInfoable)
                }
            }
            m.SetAdministrativeUnitInfos(res)
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateCommunicationsIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CommunicationsIdentitySetable))
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. The identity of the call participant.
// returns a CommunicationsIdentitySetable when successful
func (m *ParticipantBase) GetIdentity()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CommunicationsIdentitySetable) {
    return m.identity
}
// Serialize serializes information the current object
func (m *ParticipantBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdministrativeUnitInfos() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdministrativeUnitInfos()))
        for i, v := range m.GetAdministrativeUnitInfos() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("administrativeUnitInfos", cast)
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
    return nil
}
// SetAdministrativeUnitInfos sets the administrativeUnitInfos property value. List of administrativeUnitInfo objects for the call participant.
func (m *ParticipantBase) SetAdministrativeUnitInfos(value []AdministrativeUnitInfoable)() {
    m.administrativeUnitInfos = value
}
// SetIdentity sets the identity property value. The identity of the call participant.
func (m *ParticipantBase) SetIdentity(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CommunicationsIdentitySetable)() {
    m.identity = value
}
type ParticipantBaseable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdministrativeUnitInfos()([]AdministrativeUnitInfoable)
    GetIdentity()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CommunicationsIdentitySetable)
    SetAdministrativeUnitInfos(value []AdministrativeUnitInfoable)()
    SetIdentity(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CommunicationsIdentitySetable)()
}
