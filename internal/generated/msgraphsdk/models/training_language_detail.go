package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TrainingLanguageDetail struct {
    Entity
    // Language specific content for the training.
    content *string
    // Identity of the user who created the language details.
    createdBy EmailIdentityable
    // Date and time when the language details were created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description as defined by the user.
    description *string
    // Display name as defined by the user.
    displayName *string
    // Indicates whether the training has a default language.
    isDefaultLangauge *bool
    // Identity of the user who last modified the details.
    lastModifiedBy EmailIdentityable
    // Date and time when the trainingLanguageDetail was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Content locale for the training detail.
    locale *string
}
// NewTrainingLanguageDetail instantiates a new TrainingLanguageDetail and sets the default values.
func NewTrainingLanguageDetail()(*TrainingLanguageDetail) {
    m := &TrainingLanguageDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrainingLanguageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTrainingLanguageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrainingLanguageDetail(), nil
}
// GetContent gets the content property value. Language specific content for the training.
// returns a *string when successful
func (m *TrainingLanguageDetail) GetContent()(*string) {
    return m.content
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the language details.
// returns a EmailIdentityable when successful
func (m *TrainingLanguageDetail) GetCreatedBy()(EmailIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time when the language details were created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *TrainingLanguageDetail) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description as defined by the user.
// returns a *string when successful
func (m *TrainingLanguageDetail) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name as defined by the user.
// returns a *string when successful
func (m *TrainingLanguageDetail) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TrainingLanguageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailIdentityable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isDefaultLangauge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultLangauge(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(EmailIdentityable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
        }
        return nil
    }
    return res
}
// GetIsDefaultLangauge gets the isDefaultLangauge property value. Indicates whether the training has a default language.
// returns a *bool when successful
func (m *TrainingLanguageDetail) GetIsDefaultLangauge()(*bool) {
    return m.isDefaultLangauge
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the details.
// returns a EmailIdentityable when successful
func (m *TrainingLanguageDetail) GetLastModifiedBy()(EmailIdentityable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time when the trainingLanguageDetail was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *TrainingLanguageDetail) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLocale gets the locale property value. Content locale for the training detail.
// returns a *string when successful
func (m *TrainingLanguageDetail) GetLocale()(*string) {
    return m.locale
}
// Serialize serializes information the current object
func (m *TrainingLanguageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultLangauge", m.GetIsDefaultLangauge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("locale", m.GetLocale())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. Language specific content for the training.
func (m *TrainingLanguageDetail) SetContent(value *string)() {
    m.content = value
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the language details.
func (m *TrainingLanguageDetail) SetCreatedBy(value EmailIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time when the language details were created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *TrainingLanguageDetail) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description as defined by the user.
func (m *TrainingLanguageDetail) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name as defined by the user.
func (m *TrainingLanguageDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsDefaultLangauge sets the isDefaultLangauge property value. Indicates whether the training has a default language.
func (m *TrainingLanguageDetail) SetIsDefaultLangauge(value *bool)() {
    m.isDefaultLangauge = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the details.
func (m *TrainingLanguageDetail) SetLastModifiedBy(value EmailIdentityable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time when the trainingLanguageDetail was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *TrainingLanguageDetail) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLocale sets the locale property value. Content locale for the training detail.
func (m *TrainingLanguageDetail) SetLocale(value *string)() {
    m.locale = value
}
type TrainingLanguageDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetCreatedBy()(EmailIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsDefaultLangauge()(*bool)
    GetLastModifiedBy()(EmailIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocale()(*string)
    SetContent(value *string)()
    SetCreatedBy(value EmailIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsDefaultLangauge(value *bool)()
    SetLastModifiedBy(value EmailIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocale(value *string)()
}
