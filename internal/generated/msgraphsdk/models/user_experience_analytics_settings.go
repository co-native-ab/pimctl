package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsSettings the user experience analytics insight is the recomendation to improve the user experience analytics score.
type UserExperienceAnalyticsSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // When TRUE, indicates Tenant attach is configured properly and System Center Configuration Manager (SCCM) tenant attached devices will show up in endpoint analytics reporting. When FALSE, indicates Tenant attach is not configured. FALSE by default.
    configurationManagerDataConnectorConfigured *bool
    // The OdataType property
    odataType *string
}
// NewUserExperienceAnalyticsSettings instantiates a new UserExperienceAnalyticsSettings and sets the default values.
func NewUserExperienceAnalyticsSettings()(*UserExperienceAnalyticsSettings) {
    m := &UserExperienceAnalyticsSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserExperienceAnalyticsSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserExperienceAnalyticsSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfigurationManagerDataConnectorConfigured gets the configurationManagerDataConnectorConfigured property value. When TRUE, indicates Tenant attach is configured properly and System Center Configuration Manager (SCCM) tenant attached devices will show up in endpoint analytics reporting. When FALSE, indicates Tenant attach is not configured. FALSE by default.
// returns a *bool when successful
func (m *UserExperienceAnalyticsSettings) GetConfigurationManagerDataConnectorConfigured()(*bool) {
    return m.configurationManagerDataConnectorConfigured
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configurationManagerDataConnectorConfigured"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerDataConnectorConfigured(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *UserExperienceAnalyticsSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("configurationManagerDataConnectorConfigured", m.GetConfigurationManagerDataConnectorConfigured())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfigurationManagerDataConnectorConfigured sets the configurationManagerDataConnectorConfigured property value. When TRUE, indicates Tenant attach is configured properly and System Center Configuration Manager (SCCM) tenant attached devices will show up in endpoint analytics reporting. When FALSE, indicates Tenant attach is not configured. FALSE by default.
func (m *UserExperienceAnalyticsSettings) SetConfigurationManagerDataConnectorConfigured(value *bool)() {
    m.configurationManagerDataConnectorConfigured = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsSettings) SetOdataType(value *string)() {
    m.odataType = value
}
type UserExperienceAnalyticsSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurationManagerDataConnectorConfigured()(*bool)
    GetOdataType()(*string)
    SetConfigurationManagerDataConnectorConfigured(value *bool)()
    SetOdataType(value *string)()
}
