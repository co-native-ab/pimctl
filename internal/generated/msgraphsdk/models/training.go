package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Training struct {
    Entity
    // Training availability status. Possible values are: unknown, notAvailable, available, archive, delete, unknownFutureValue.
    availabilityStatus *TrainingAvailabilityStatus
    // Identity of the user who created the training.
    createdBy EmailIdentityable
    // Date and time when the training was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description for the training.
    description *string
    // The display name for the training.
    displayName *string
    // Training duration.
    durationInMinutes *int32
    // Indicates whether the training has any evaluation.
    hasEvaluation *bool
    // Language specific details on a training.
    languageDetails []TrainingLanguageDetailable
    // Identity of the user who last modified the training.
    lastModifiedBy EmailIdentityable
    // Date and time when the training was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Training content source. Possible values are: unknown, global, tenant, unknownFutureValue.
    source *SimulationContentSource
    // Supported locales for content for the associated training.
    supportedLocales []string
    // Training tags.
    tags []string
    // The type of training. Possible values are: unknown, phishing, unknownFutureValue.
    typeEscaped *TrainingType
}
// NewTraining instantiates a new Training and sets the default values.
func NewTraining()(*Training) {
    m := &Training{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrainingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTrainingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTraining(), nil
}
// GetAvailabilityStatus gets the availabilityStatus property value. Training availability status. Possible values are: unknown, notAvailable, available, archive, delete, unknownFutureValue.
// returns a *TrainingAvailabilityStatus when successful
func (m *Training) GetAvailabilityStatus()(*TrainingAvailabilityStatus) {
    return m.availabilityStatus
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the training.
// returns a EmailIdentityable when successful
func (m *Training) GetCreatedBy()(EmailIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time when the training was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *Training) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description for the training.
// returns a *string when successful
func (m *Training) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for the training.
// returns a *string when successful
func (m *Training) GetDisplayName()(*string) {
    return m.displayName
}
// GetDurationInMinutes gets the durationInMinutes property value. Training duration.
// returns a *int32 when successful
func (m *Training) GetDurationInMinutes()(*int32) {
    return m.durationInMinutes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Training) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["availabilityStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrainingAvailabilityStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityStatus(val.(*TrainingAvailabilityStatus))
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
    res["durationInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInMinutes(val)
        }
        return nil
    }
    res["hasEvaluation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasEvaluation(val)
        }
        return nil
    }
    res["languageDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrainingLanguageDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrainingLanguageDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TrainingLanguageDetailable)
                }
            }
            m.SetLanguageDetails(res)
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationContentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*SimulationContentSource))
        }
        return nil
    }
    res["supportedLocales"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSupportedLocales(res)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTags(res)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrainingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*TrainingType))
        }
        return nil
    }
    return res
}
// GetHasEvaluation gets the hasEvaluation property value. Indicates whether the training has any evaluation.
// returns a *bool when successful
func (m *Training) GetHasEvaluation()(*bool) {
    return m.hasEvaluation
}
// GetLanguageDetails gets the languageDetails property value. Language specific details on a training.
// returns a []TrainingLanguageDetailable when successful
func (m *Training) GetLanguageDetails()([]TrainingLanguageDetailable) {
    return m.languageDetails
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the training.
// returns a EmailIdentityable when successful
func (m *Training) GetLastModifiedBy()(EmailIdentityable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time when the training was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *Training) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetSource gets the source property value. Training content source. Possible values are: unknown, global, tenant, unknownFutureValue.
// returns a *SimulationContentSource when successful
func (m *Training) GetSource()(*SimulationContentSource) {
    return m.source
}
// GetSupportedLocales gets the supportedLocales property value. Supported locales for content for the associated training.
// returns a []string when successful
func (m *Training) GetSupportedLocales()([]string) {
    return m.supportedLocales
}
// GetTags gets the tags property value. Training tags.
// returns a []string when successful
func (m *Training) GetTags()([]string) {
    return m.tags
}
// GetTypeEscaped gets the type property value. The type of training. Possible values are: unknown, phishing, unknownFutureValue.
// returns a *TrainingType when successful
func (m *Training) GetTypeEscaped()(*TrainingType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Training) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAvailabilityStatus() != nil {
        cast := (*m.GetAvailabilityStatus()).String()
        err = writer.WriteStringValue("availabilityStatus", &cast)
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
        err = writer.WriteInt32Value("durationInMinutes", m.GetDurationInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasEvaluation", m.GetHasEvaluation())
        if err != nil {
            return err
        }
    }
    if m.GetLanguageDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLanguageDetails()))
        for i, v := range m.GetLanguageDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("languageDetails", cast)
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
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedLocales() != nil {
        err = writer.WriteCollectionOfStringValues("supportedLocales", m.GetSupportedLocales())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailabilityStatus sets the availabilityStatus property value. Training availability status. Possible values are: unknown, notAvailable, available, archive, delete, unknownFutureValue.
func (m *Training) SetAvailabilityStatus(value *TrainingAvailabilityStatus)() {
    m.availabilityStatus = value
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the training.
func (m *Training) SetCreatedBy(value EmailIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time when the training was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Training) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description for the training.
func (m *Training) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the training.
func (m *Training) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDurationInMinutes sets the durationInMinutes property value. Training duration.
func (m *Training) SetDurationInMinutes(value *int32)() {
    m.durationInMinutes = value
}
// SetHasEvaluation sets the hasEvaluation property value. Indicates whether the training has any evaluation.
func (m *Training) SetHasEvaluation(value *bool)() {
    m.hasEvaluation = value
}
// SetLanguageDetails sets the languageDetails property value. Language specific details on a training.
func (m *Training) SetLanguageDetails(value []TrainingLanguageDetailable)() {
    m.languageDetails = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the training.
func (m *Training) SetLastModifiedBy(value EmailIdentityable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time when the training was last modified. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Training) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetSource sets the source property value. Training content source. Possible values are: unknown, global, tenant, unknownFutureValue.
func (m *Training) SetSource(value *SimulationContentSource)() {
    m.source = value
}
// SetSupportedLocales sets the supportedLocales property value. Supported locales for content for the associated training.
func (m *Training) SetSupportedLocales(value []string)() {
    m.supportedLocales = value
}
// SetTags sets the tags property value. Training tags.
func (m *Training) SetTags(value []string)() {
    m.tags = value
}
// SetTypeEscaped sets the type property value. The type of training. Possible values are: unknown, phishing, unknownFutureValue.
func (m *Training) SetTypeEscaped(value *TrainingType)() {
    m.typeEscaped = value
}
type Trainingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityStatus()(*TrainingAvailabilityStatus)
    GetCreatedBy()(EmailIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDurationInMinutes()(*int32)
    GetHasEvaluation()(*bool)
    GetLanguageDetails()([]TrainingLanguageDetailable)
    GetLastModifiedBy()(EmailIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSource()(*SimulationContentSource)
    GetSupportedLocales()([]string)
    GetTags()([]string)
    GetTypeEscaped()(*TrainingType)
    SetAvailabilityStatus(value *TrainingAvailabilityStatus)()
    SetCreatedBy(value EmailIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDurationInMinutes(value *int32)()
    SetHasEvaluation(value *bool)()
    SetLanguageDetails(value []TrainingLanguageDetailable)()
    SetLastModifiedBy(value EmailIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSource(value *SimulationContentSource)()
    SetSupportedLocales(value []string)()
    SetTags(value []string)()
    SetTypeEscaped(value *TrainingType)()
}
