package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type Relation struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
    fromTerm Termable
    // The type of relation. Possible values are: pin, reuse.
    relationship *RelationType
    // The [set] in which the relation is relevant.
    set Setable
    // The to [term] of the relation. The term to which the relationship is defined.
    toTerm Termable
}
// NewRelation instantiates a new Relation and sets the default values.
func NewRelation()(*Relation) {
    m := &Relation{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateRelationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Relation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fromTerm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromTerm(val.(Termable))
        }
        return nil
    }
    res["relationship"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRelationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelationship(val.(*RelationType))
        }
        return nil
    }
    res["set"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSet(val.(Setable))
        }
        return nil
    }
    res["toTerm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToTerm(val.(Termable))
        }
        return nil
    }
    return res
}
// GetFromTerm gets the fromTerm property value. The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
// returns a Termable when successful
func (m *Relation) GetFromTerm()(Termable) {
    return m.fromTerm
}
// GetRelationship gets the relationship property value. The type of relation. Possible values are: pin, reuse.
// returns a *RelationType when successful
func (m *Relation) GetRelationship()(*RelationType) {
    return m.relationship
}
// GetSet gets the set property value. The [set] in which the relation is relevant.
// returns a Setable when successful
func (m *Relation) GetSet()(Setable) {
    return m.set
}
// GetToTerm gets the toTerm property value. The to [term] of the relation. The term to which the relationship is defined.
// returns a Termable when successful
func (m *Relation) GetToTerm()(Termable) {
    return m.toTerm
}
// Serialize serializes information the current object
func (m *Relation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fromTerm", m.GetFromTerm())
        if err != nil {
            return err
        }
    }
    if m.GetRelationship() != nil {
        cast := (*m.GetRelationship()).String()
        err = writer.WriteStringValue("relationship", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("set", m.GetSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("toTerm", m.GetToTerm())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFromTerm sets the fromTerm property value. The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
func (m *Relation) SetFromTerm(value Termable)() {
    m.fromTerm = value
}
// SetRelationship sets the relationship property value. The type of relation. Possible values are: pin, reuse.
func (m *Relation) SetRelationship(value *RelationType)() {
    m.relationship = value
}
// SetSet sets the set property value. The [set] in which the relation is relevant.
func (m *Relation) SetSet(value Setable)() {
    m.set = value
}
// SetToTerm sets the toTerm property value. The to [term] of the relation. The term to which the relationship is defined.
func (m *Relation) SetToTerm(value Termable)() {
    m.toTerm = value
}
type Relationable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFromTerm()(Termable)
    GetRelationship()(*RelationType)
    GetSet()(Setable)
    GetToTerm()(Termable)
    SetFromTerm(value Termable)()
    SetRelationship(value *RelationType)()
    SetSet(value Setable)()
    SetToTerm(value Termable)()
}
