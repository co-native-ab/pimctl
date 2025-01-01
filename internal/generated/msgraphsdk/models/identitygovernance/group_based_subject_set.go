package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type GroupBasedSubjectSet struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.SubjectSet
    // The groups property
    groups []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Groupable
}
// NewGroupBasedSubjectSet instantiates a new GroupBasedSubjectSet and sets the default values.
func NewGroupBasedSubjectSet()(*GroupBasedSubjectSet) {
    m := &GroupBasedSubjectSet{
        SubjectSet: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewSubjectSet(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.groupBasedSubjectSet"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupBasedSubjectSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupBasedSubjectSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupBasedSubjectSet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupBasedSubjectSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectSet.GetFieldDeserializers()
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Groupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Groupable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The groups property
// returns a []Groupable when successful
func (m *GroupBasedSubjectSet) GetGroups()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Groupable) {
    return m.groups
}
// Serialize serializes information the current object
func (m *GroupBasedSubjectSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectSet.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroups sets the groups property value. The groups property
func (m *GroupBasedSubjectSet) SetGroups(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Groupable)() {
    m.groups = value
}
type GroupBasedSubjectSetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.SubjectSetable
    GetGroups()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Groupable)
    SetGroups(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Groupable)()
}
