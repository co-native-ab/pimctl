package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type Indicator struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The artifact property
    artifact Artifactable
    // The source property
    source *IndicatorSource
}
// NewIndicator instantiates a new Indicator and sets the default values.
func NewIndicator()(*Indicator) {
    m := &Indicator{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateIndicatorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIndicatorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.security.articleIndicator":
                        return NewArticleIndicator(), nil
                    case "#microsoft.graph.security.intelligenceProfileIndicator":
                        return NewIntelligenceProfileIndicator(), nil
                }
            }
        }
    }
    return NewIndicator(), nil
}
// GetArtifact gets the artifact property value. The artifact property
// returns a Artifactable when successful
func (m *Indicator) GetArtifact()(Artifactable) {
    return m.artifact
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Indicator) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["artifact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateArtifactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtifact(val.(Artifactable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIndicatorSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*IndicatorSource))
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. The source property
// returns a *IndicatorSource when successful
func (m *Indicator) GetSource()(*IndicatorSource) {
    return m.source
}
// Serialize serializes information the current object
func (m *Indicator) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("artifact", m.GetArtifact())
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArtifact sets the artifact property value. The artifact property
func (m *Indicator) SetArtifact(value Artifactable)() {
    m.artifact = value
}
// SetSource sets the source property value. The source property
func (m *Indicator) SetSource(value *IndicatorSource)() {
    m.source = value
}
type Indicatorable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArtifact()(Artifactable)
    GetSource()(*IndicatorSource)
    SetArtifact(value Artifactable)()
    SetSource(value *IndicatorSource)()
}
