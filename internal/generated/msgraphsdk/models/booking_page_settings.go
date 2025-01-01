package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BookingPageSettings struct {
    // The accessControl property
    accessControl *BookingPageAccessControl
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Custom color for the booking page. The value should be in Hex format. For example, #123456.
    bookingPageColorCode *string
    // The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
    businessTimeZone *string
    // The personal data collection and usage consent message in the booking page.
    customerConsentMessage *string
    // Determines whether the one-time password is required to create an appointment. The default value is false.
    enforceOneTimePassword *bool
    // Indicates whether the business logo is displayed on the booking page. The default value is false.
    isBusinessLogoDisplayEnabled *bool
    // Enables personal data collection and the usage consent toggle on the booking page. The default value is false.
    isCustomerConsentEnabled *bool
    // Indicates whether web crawlers index this page. The defaults value is false.
    isSearchEngineIndexabilityDisabled *bool
    // Indicates whether the time zone of the time slot is set to the time zone of the business. The default value is false.
    isTimeSlotTimeZoneSetToBusinessTimeZone *bool
    // The OdataType property
    odataType *string
    // URL of a webpage that provides the terms and conditions of the business. If a privacy policy isn't included, the following text appears on the booking page as default: 'The policies and practices of {bookingbusinessname} apply to the use of your data.'
    privacyPolicyWebUrl *string
    // URL of a webpage that provides the terms and conditions of the business.
    termsAndConditionsWebUrl *string
}
// NewBookingPageSettings instantiates a new BookingPageSettings and sets the default values.
func NewBookingPageSettings()(*BookingPageSettings) {
    m := &BookingPageSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBookingPageSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBookingPageSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingPageSettings(), nil
}
// GetAccessControl gets the accessControl property value. The accessControl property
// returns a *BookingPageAccessControl when successful
func (m *BookingPageSettings) GetAccessControl()(*BookingPageAccessControl) {
    return m.accessControl
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BookingPageSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBookingPageColorCode gets the bookingPageColorCode property value. Custom color for the booking page. The value should be in Hex format. For example, #123456.
// returns a *string when successful
func (m *BookingPageSettings) GetBookingPageColorCode()(*string) {
    return m.bookingPageColorCode
}
// GetBusinessTimeZone gets the businessTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
// returns a *string when successful
func (m *BookingPageSettings) GetBusinessTimeZone()(*string) {
    return m.businessTimeZone
}
// GetCustomerConsentMessage gets the customerConsentMessage property value. The personal data collection and usage consent message in the booking page.
// returns a *string when successful
func (m *BookingPageSettings) GetCustomerConsentMessage()(*string) {
    return m.customerConsentMessage
}
// GetEnforceOneTimePassword gets the enforceOneTimePassword property value. Determines whether the one-time password is required to create an appointment. The default value is false.
// returns a *bool when successful
func (m *BookingPageSettings) GetEnforceOneTimePassword()(*bool) {
    return m.enforceOneTimePassword
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BookingPageSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessControl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingPageAccessControl)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessControl(val.(*BookingPageAccessControl))
        }
        return nil
    }
    res["bookingPageColorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBookingPageColorCode(val)
        }
        return nil
    }
    res["businessTimeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessTimeZone(val)
        }
        return nil
    }
    res["customerConsentMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerConsentMessage(val)
        }
        return nil
    }
    res["enforceOneTimePassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceOneTimePassword(val)
        }
        return nil
    }
    res["isBusinessLogoDisplayEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBusinessLogoDisplayEnabled(val)
        }
        return nil
    }
    res["isCustomerConsentEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomerConsentEnabled(val)
        }
        return nil
    }
    res["isSearchEngineIndexabilityDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchEngineIndexabilityDisabled(val)
        }
        return nil
    }
    res["isTimeSlotTimeZoneSetToBusinessTimeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTimeSlotTimeZoneSetToBusinessTimeZone(val)
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
    res["privacyPolicyWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyPolicyWebUrl(val)
        }
        return nil
    }
    res["termsAndConditionsWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsAndConditionsWebUrl(val)
        }
        return nil
    }
    return res
}
// GetIsBusinessLogoDisplayEnabled gets the isBusinessLogoDisplayEnabled property value. Indicates whether the business logo is displayed on the booking page. The default value is false.
// returns a *bool when successful
func (m *BookingPageSettings) GetIsBusinessLogoDisplayEnabled()(*bool) {
    return m.isBusinessLogoDisplayEnabled
}
// GetIsCustomerConsentEnabled gets the isCustomerConsentEnabled property value. Enables personal data collection and the usage consent toggle on the booking page. The default value is false.
// returns a *bool when successful
func (m *BookingPageSettings) GetIsCustomerConsentEnabled()(*bool) {
    return m.isCustomerConsentEnabled
}
// GetIsSearchEngineIndexabilityDisabled gets the isSearchEngineIndexabilityDisabled property value. Indicates whether web crawlers index this page. The defaults value is false.
// returns a *bool when successful
func (m *BookingPageSettings) GetIsSearchEngineIndexabilityDisabled()(*bool) {
    return m.isSearchEngineIndexabilityDisabled
}
// GetIsTimeSlotTimeZoneSetToBusinessTimeZone gets the isTimeSlotTimeZoneSetToBusinessTimeZone property value. Indicates whether the time zone of the time slot is set to the time zone of the business. The default value is false.
// returns a *bool when successful
func (m *BookingPageSettings) GetIsTimeSlotTimeZoneSetToBusinessTimeZone()(*bool) {
    return m.isTimeSlotTimeZoneSetToBusinessTimeZone
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *BookingPageSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetPrivacyPolicyWebUrl gets the privacyPolicyWebUrl property value. URL of a webpage that provides the terms and conditions of the business. If a privacy policy isn't included, the following text appears on the booking page as default: 'The policies and practices of {bookingbusinessname} apply to the use of your data.'
// returns a *string when successful
func (m *BookingPageSettings) GetPrivacyPolicyWebUrl()(*string) {
    return m.privacyPolicyWebUrl
}
// GetTermsAndConditionsWebUrl gets the termsAndConditionsWebUrl property value. URL of a webpage that provides the terms and conditions of the business.
// returns a *string when successful
func (m *BookingPageSettings) GetTermsAndConditionsWebUrl()(*string) {
    return m.termsAndConditionsWebUrl
}
// Serialize serializes information the current object
func (m *BookingPageSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessControl() != nil {
        cast := (*m.GetAccessControl()).String()
        err := writer.WriteStringValue("accessControl", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bookingPageColorCode", m.GetBookingPageColorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("businessTimeZone", m.GetBusinessTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customerConsentMessage", m.GetCustomerConsentMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enforceOneTimePassword", m.GetEnforceOneTimePassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isBusinessLogoDisplayEnabled", m.GetIsBusinessLogoDisplayEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCustomerConsentEnabled", m.GetIsCustomerConsentEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSearchEngineIndexabilityDisabled", m.GetIsSearchEngineIndexabilityDisabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isTimeSlotTimeZoneSetToBusinessTimeZone", m.GetIsTimeSlotTimeZoneSetToBusinessTimeZone())
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
        err := writer.WriteStringValue("privacyPolicyWebUrl", m.GetPrivacyPolicyWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("termsAndConditionsWebUrl", m.GetTermsAndConditionsWebUrl())
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
// SetAccessControl sets the accessControl property value. The accessControl property
func (m *BookingPageSettings) SetAccessControl(value *BookingPageAccessControl)() {
    m.accessControl = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingPageSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBookingPageColorCode sets the bookingPageColorCode property value. Custom color for the booking page. The value should be in Hex format. For example, #123456.
func (m *BookingPageSettings) SetBookingPageColorCode(value *string)() {
    m.bookingPageColorCode = value
}
// SetBusinessTimeZone sets the businessTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
func (m *BookingPageSettings) SetBusinessTimeZone(value *string)() {
    m.businessTimeZone = value
}
// SetCustomerConsentMessage sets the customerConsentMessage property value. The personal data collection and usage consent message in the booking page.
func (m *BookingPageSettings) SetCustomerConsentMessage(value *string)() {
    m.customerConsentMessage = value
}
// SetEnforceOneTimePassword sets the enforceOneTimePassword property value. Determines whether the one-time password is required to create an appointment. The default value is false.
func (m *BookingPageSettings) SetEnforceOneTimePassword(value *bool)() {
    m.enforceOneTimePassword = value
}
// SetIsBusinessLogoDisplayEnabled sets the isBusinessLogoDisplayEnabled property value. Indicates whether the business logo is displayed on the booking page. The default value is false.
func (m *BookingPageSettings) SetIsBusinessLogoDisplayEnabled(value *bool)() {
    m.isBusinessLogoDisplayEnabled = value
}
// SetIsCustomerConsentEnabled sets the isCustomerConsentEnabled property value. Enables personal data collection and the usage consent toggle on the booking page. The default value is false.
func (m *BookingPageSettings) SetIsCustomerConsentEnabled(value *bool)() {
    m.isCustomerConsentEnabled = value
}
// SetIsSearchEngineIndexabilityDisabled sets the isSearchEngineIndexabilityDisabled property value. Indicates whether web crawlers index this page. The defaults value is false.
func (m *BookingPageSettings) SetIsSearchEngineIndexabilityDisabled(value *bool)() {
    m.isSearchEngineIndexabilityDisabled = value
}
// SetIsTimeSlotTimeZoneSetToBusinessTimeZone sets the isTimeSlotTimeZoneSetToBusinessTimeZone property value. Indicates whether the time zone of the time slot is set to the time zone of the business. The default value is false.
func (m *BookingPageSettings) SetIsTimeSlotTimeZoneSetToBusinessTimeZone(value *bool)() {
    m.isTimeSlotTimeZoneSetToBusinessTimeZone = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BookingPageSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrivacyPolicyWebUrl sets the privacyPolicyWebUrl property value. URL of a webpage that provides the terms and conditions of the business. If a privacy policy isn't included, the following text appears on the booking page as default: 'The policies and practices of {bookingbusinessname} apply to the use of your data.'
func (m *BookingPageSettings) SetPrivacyPolicyWebUrl(value *string)() {
    m.privacyPolicyWebUrl = value
}
// SetTermsAndConditionsWebUrl sets the termsAndConditionsWebUrl property value. URL of a webpage that provides the terms and conditions of the business.
func (m *BookingPageSettings) SetTermsAndConditionsWebUrl(value *string)() {
    m.termsAndConditionsWebUrl = value
}
type BookingPageSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessControl()(*BookingPageAccessControl)
    GetBookingPageColorCode()(*string)
    GetBusinessTimeZone()(*string)
    GetCustomerConsentMessage()(*string)
    GetEnforceOneTimePassword()(*bool)
    GetIsBusinessLogoDisplayEnabled()(*bool)
    GetIsCustomerConsentEnabled()(*bool)
    GetIsSearchEngineIndexabilityDisabled()(*bool)
    GetIsTimeSlotTimeZoneSetToBusinessTimeZone()(*bool)
    GetOdataType()(*string)
    GetPrivacyPolicyWebUrl()(*string)
    GetTermsAndConditionsWebUrl()(*string)
    SetAccessControl(value *BookingPageAccessControl)()
    SetBookingPageColorCode(value *string)()
    SetBusinessTimeZone(value *string)()
    SetCustomerConsentMessage(value *string)()
    SetEnforceOneTimePassword(value *bool)()
    SetIsBusinessLogoDisplayEnabled(value *bool)()
    SetIsCustomerConsentEnabled(value *bool)()
    SetIsSearchEngineIndexabilityDisabled(value *bool)()
    SetIsTimeSlotTimeZoneSetToBusinessTimeZone(value *bool)()
    SetOdataType(value *string)()
    SetPrivacyPolicyWebUrl(value *string)()
    SetTermsAndConditionsWebUrl(value *string)()
}
