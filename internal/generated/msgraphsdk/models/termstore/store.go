package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type Store struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Default language of the term store.
    defaultLanguageTag *string
    // Collection of all groups available in the term store.
    groups []Groupable
    // List of languages for the term store.
    languageTags []string
    // Collection of all sets available in the term store. This relationship can only be used to load a specific term set.
    sets []Setable
}
// NewStore instantiates a new Store and sets the default values.
func NewStore()(*Store) {
    m := &Store{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateStoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStore(), nil
}
// GetDefaultLanguageTag gets the defaultLanguageTag property value. Default language of the term store.
// returns a *string when successful
func (m *Store) GetDefaultLanguageTag()(*string) {
    return m.defaultLanguageTag
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Store) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultLanguageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguageTag(val)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Groupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Groupable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["languageTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetLanguageTags(res)
        }
        return nil
    }
    res["sets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Setable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Setable)
                }
            }
            m.SetSets(res)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Collection of all groups available in the term store.
// returns a []Groupable when successful
func (m *Store) GetGroups()([]Groupable) {
    return m.groups
}
// GetLanguageTags gets the languageTags property value. List of languages for the term store.
// returns a []string when successful
func (m *Store) GetLanguageTags()([]string) {
    return m.languageTags
}
// GetSets gets the sets property value. Collection of all sets available in the term store. This relationship can only be used to load a specific term set.
// returns a []Setable when successful
func (m *Store) GetSets()([]Setable) {
    return m.sets
}
// Serialize serializes information the current object
func (m *Store) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("defaultLanguageTag", m.GetDefaultLanguageTag())
        if err != nil {
            return err
        }
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
    if m.GetLanguageTags() != nil {
        err = writer.WriteCollectionOfStringValues("languageTags", m.GetLanguageTags())
        if err != nil {
            return err
        }
    }
    if m.GetSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSets()))
        for i, v := range m.GetSets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultLanguageTag sets the defaultLanguageTag property value. Default language of the term store.
func (m *Store) SetDefaultLanguageTag(value *string)() {
    m.defaultLanguageTag = value
}
// SetGroups sets the groups property value. Collection of all groups available in the term store.
func (m *Store) SetGroups(value []Groupable)() {
    m.groups = value
}
// SetLanguageTags sets the languageTags property value. List of languages for the term store.
func (m *Store) SetLanguageTags(value []string)() {
    m.languageTags = value
}
// SetSets sets the sets property value. Collection of all sets available in the term store. This relationship can only be used to load a specific term set.
func (m *Store) SetSets(value []Setable)() {
    m.sets = value
}
type Storeable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultLanguageTag()(*string)
    GetGroups()([]Groupable)
    GetLanguageTags()([]string)
    GetSets()([]Setable)
    SetDefaultLanguageTag(value *string)()
    SetGroups(value []Groupable)()
    SetLanguageTags(value []string)()
    SetSets(value []Setable)()
}
