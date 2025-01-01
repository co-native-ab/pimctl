package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type AnswerVariant struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The answer variation description that is shown on the search results page.
    description *string
    // The answer variation name that is displayed in search results.
    displayName *string
    // The country or region that can view this answer variation.
    languageTag *string
    // The OdataType property
    odataType *string
    // The device or operating system that can view this answer variation. Possible values are: android, androidForWork, ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, unknown, androidASOP, androidMobileApplicationManagement, iOSMobileApplicationManagement, unknownFutureValue.
    platform *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DevicePlatformType
    // The URL link for the answer variation. When users select this answer variation from the search results, they're directed to the specified URL.
    webUrl *string
}
// NewAnswerVariant instantiates a new AnswerVariant and sets the default values.
func NewAnswerVariant()(*AnswerVariant) {
    m := &AnswerVariant{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAnswerVariantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAnswerVariantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnswerVariant(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AnswerVariant) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The answer variation description that is shown on the search results page.
// returns a *string when successful
func (m *AnswerVariant) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The answer variation name that is displayed in search results.
// returns a *string when successful
func (m *AnswerVariant) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AnswerVariant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
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
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DevicePlatformType))
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetLanguageTag gets the languageTag property value. The country or region that can view this answer variation.
// returns a *string when successful
func (m *AnswerVariant) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AnswerVariant) GetOdataType()(*string) {
    return m.odataType
}
// GetPlatform gets the platform property value. The device or operating system that can view this answer variation. Possible values are: android, androidForWork, ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, unknown, androidASOP, androidMobileApplicationManagement, iOSMobileApplicationManagement, unknownFutureValue.
// returns a *DevicePlatformType when successful
func (m *AnswerVariant) GetPlatform()(*ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DevicePlatformType) {
    return m.platform
}
// GetWebUrl gets the webUrl property value. The URL link for the answer variation. When users select this answer variation from the search results, they're directed to the specified URL.
// returns a *string when successful
func (m *AnswerVariant) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *AnswerVariant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
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
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *AnswerVariant) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The answer variation description that is shown on the search results page.
func (m *AnswerVariant) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The answer variation name that is displayed in search results.
func (m *AnswerVariant) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLanguageTag sets the languageTag property value. The country or region that can view this answer variation.
func (m *AnswerVariant) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AnswerVariant) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPlatform sets the platform property value. The device or operating system that can view this answer variation. Possible values are: android, androidForWork, ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, unknown, androidASOP, androidMobileApplicationManagement, iOSMobileApplicationManagement, unknownFutureValue.
func (m *AnswerVariant) SetPlatform(value *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DevicePlatformType)() {
    m.platform = value
}
// SetWebUrl sets the webUrl property value. The URL link for the answer variation. When users select this answer variation from the search results, they're directed to the specified URL.
func (m *AnswerVariant) SetWebUrl(value *string)() {
    m.webUrl = value
}
type AnswerVariantable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLanguageTag()(*string)
    GetOdataType()(*string)
    GetPlatform()(*ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DevicePlatformType)
    GetWebUrl()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLanguageTag(value *string)()
    SetOdataType(value *string)()
    SetPlatform(value *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.DevicePlatformType)()
    SetWebUrl(value *string)()
}
