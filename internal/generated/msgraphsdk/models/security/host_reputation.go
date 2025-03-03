package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type HostReputation struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The classification property
    classification *HostReputationClassification
    // A collection of rules that have been used to calculate the classification and score.
    rules []HostReputationRuleable
    // The calculated score (0-100) of the requested host. A higher value indicates that this host is more likely to be suspicious or malicious.
    score *int32
}
// NewHostReputation instantiates a new HostReputation and sets the default values.
func NewHostReputation()(*HostReputation) {
    m := &HostReputation{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateHostReputationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHostReputationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostReputation(), nil
}
// GetClassification gets the classification property value. The classification property
// returns a *HostReputationClassification when successful
func (m *HostReputation) GetClassification()(*HostReputationClassification) {
    return m.classification
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HostReputation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHostReputationClassification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(*HostReputationClassification))
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostReputationRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostReputationRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostReputationRuleable)
                }
            }
            m.SetRules(res)
        }
        return nil
    }
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    return res
}
// GetRules gets the rules property value. A collection of rules that have been used to calculate the classification and score.
// returns a []HostReputationRuleable when successful
func (m *HostReputation) GetRules()([]HostReputationRuleable) {
    return m.rules
}
// GetScore gets the score property value. The calculated score (0-100) of the requested host. A higher value indicates that this host is more likely to be suspicious or malicious.
// returns a *int32 when successful
func (m *HostReputation) GetScore()(*int32) {
    return m.score
}
// Serialize serializes information the current object
func (m *HostReputation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := (*m.GetClassification()).String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("score", m.GetScore())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassification sets the classification property value. The classification property
func (m *HostReputation) SetClassification(value *HostReputationClassification)() {
    m.classification = value
}
// SetRules sets the rules property value. A collection of rules that have been used to calculate the classification and score.
func (m *HostReputation) SetRules(value []HostReputationRuleable)() {
    m.rules = value
}
// SetScore sets the score property value. The calculated score (0-100) of the requested host. A higher value indicates that this host is more likely to be suspicious or malicious.
func (m *HostReputation) SetScore(value *int32)() {
    m.score = value
}
type HostReputationable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassification()(*HostReputationClassification)
    GetRules()([]HostReputationRuleable)
    GetScore()(*int32)
    SetClassification(value *HostReputationClassification)()
    SetRules(value []HostReputationRuleable)()
    SetScore(value *int32)()
}
