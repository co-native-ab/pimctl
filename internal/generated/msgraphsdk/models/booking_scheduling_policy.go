package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingSchedulingPolicy this type represents the set of policies that dictate how bookings can be created in a Booking Calendar.
type BookingSchedulingPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // True to allow customers to choose a specific person for the booking.
    allowStaffSelection *bool
    // Custom availability of the service in a given time frame.
    customAvailabilities []BookingsAvailabilityWindowable
    // General availability of the service defined by the scheduling policy.
    generalAvailability BookingsAvailabilityable
    // Indicates whether the meeting invite is sent to the customers. The default value is false.
    isMeetingInviteToCustomersEnabled *bool
    // Maximum number of days in advance that a booking can be made. It follows the ISO 8601 format.
    maximumAdvance *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The minimum amount of time before which bookings and cancellations must be made. It follows the ISO 8601 format.
    minimumLeadTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The OdataType property
    odataType *string
    // True to notify the business via email when a booking is created or changed. Use the email address specified in the email property of the bookingBusiness entity for the business.
    sendConfirmationsToOwner *bool
    // Duration of each time slot, denoted in ISO 8601 format.
    timeSlotInterval *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewBookingSchedulingPolicy instantiates a new BookingSchedulingPolicy and sets the default values.
func NewBookingSchedulingPolicy()(*BookingSchedulingPolicy) {
    m := &BookingSchedulingPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBookingSchedulingPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBookingSchedulingPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingSchedulingPolicy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BookingSchedulingPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowStaffSelection gets the allowStaffSelection property value. True to allow customers to choose a specific person for the booking.
// returns a *bool when successful
func (m *BookingSchedulingPolicy) GetAllowStaffSelection()(*bool) {
    return m.allowStaffSelection
}
// GetCustomAvailabilities gets the customAvailabilities property value. Custom availability of the service in a given time frame.
// returns a []BookingsAvailabilityWindowable when successful
func (m *BookingSchedulingPolicy) GetCustomAvailabilities()([]BookingsAvailabilityWindowable) {
    return m.customAvailabilities
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BookingSchedulingPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowStaffSelection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowStaffSelection(val)
        }
        return nil
    }
    res["customAvailabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingsAvailabilityWindowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingsAvailabilityWindowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingsAvailabilityWindowable)
                }
            }
            m.SetCustomAvailabilities(res)
        }
        return nil
    }
    res["generalAvailability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBookingsAvailabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneralAvailability(val.(BookingsAvailabilityable))
        }
        return nil
    }
    res["isMeetingInviteToCustomersEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMeetingInviteToCustomersEnabled(val)
        }
        return nil
    }
    res["maximumAdvance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAdvance(val)
        }
        return nil
    }
    res["minimumLeadTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumLeadTime(val)
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
    res["sendConfirmationsToOwner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendConfirmationsToOwner(val)
        }
        return nil
    }
    res["timeSlotInterval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeSlotInterval(val)
        }
        return nil
    }
    return res
}
// GetGeneralAvailability gets the generalAvailability property value. General availability of the service defined by the scheduling policy.
// returns a BookingsAvailabilityable when successful
func (m *BookingSchedulingPolicy) GetGeneralAvailability()(BookingsAvailabilityable) {
    return m.generalAvailability
}
// GetIsMeetingInviteToCustomersEnabled gets the isMeetingInviteToCustomersEnabled property value. Indicates whether the meeting invite is sent to the customers. The default value is false.
// returns a *bool when successful
func (m *BookingSchedulingPolicy) GetIsMeetingInviteToCustomersEnabled()(*bool) {
    return m.isMeetingInviteToCustomersEnabled
}
// GetMaximumAdvance gets the maximumAdvance property value. Maximum number of days in advance that a booking can be made. It follows the ISO 8601 format.
// returns a *ISODuration when successful
func (m *BookingSchedulingPolicy) GetMaximumAdvance()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.maximumAdvance
}
// GetMinimumLeadTime gets the minimumLeadTime property value. The minimum amount of time before which bookings and cancellations must be made. It follows the ISO 8601 format.
// returns a *ISODuration when successful
func (m *BookingSchedulingPolicy) GetMinimumLeadTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.minimumLeadTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *BookingSchedulingPolicy) GetOdataType()(*string) {
    return m.odataType
}
// GetSendConfirmationsToOwner gets the sendConfirmationsToOwner property value. True to notify the business via email when a booking is created or changed. Use the email address specified in the email property of the bookingBusiness entity for the business.
// returns a *bool when successful
func (m *BookingSchedulingPolicy) GetSendConfirmationsToOwner()(*bool) {
    return m.sendConfirmationsToOwner
}
// GetTimeSlotInterval gets the timeSlotInterval property value. Duration of each time slot, denoted in ISO 8601 format.
// returns a *ISODuration when successful
func (m *BookingSchedulingPolicy) GetTimeSlotInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.timeSlotInterval
}
// Serialize serializes information the current object
func (m *BookingSchedulingPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowStaffSelection", m.GetAllowStaffSelection())
        if err != nil {
            return err
        }
    }
    if m.GetCustomAvailabilities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomAvailabilities()))
        for i, v := range m.GetCustomAvailabilities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("customAvailabilities", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("generalAvailability", m.GetGeneralAvailability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMeetingInviteToCustomersEnabled", m.GetIsMeetingInviteToCustomersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maximumAdvance", m.GetMaximumAdvance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("minimumLeadTime", m.GetMinimumLeadTime())
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
        err := writer.WriteBoolValue("sendConfirmationsToOwner", m.GetSendConfirmationsToOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("timeSlotInterval", m.GetTimeSlotInterval())
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
func (m *BookingSchedulingPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowStaffSelection sets the allowStaffSelection property value. True to allow customers to choose a specific person for the booking.
func (m *BookingSchedulingPolicy) SetAllowStaffSelection(value *bool)() {
    m.allowStaffSelection = value
}
// SetCustomAvailabilities sets the customAvailabilities property value. Custom availability of the service in a given time frame.
func (m *BookingSchedulingPolicy) SetCustomAvailabilities(value []BookingsAvailabilityWindowable)() {
    m.customAvailabilities = value
}
// SetGeneralAvailability sets the generalAvailability property value. General availability of the service defined by the scheduling policy.
func (m *BookingSchedulingPolicy) SetGeneralAvailability(value BookingsAvailabilityable)() {
    m.generalAvailability = value
}
// SetIsMeetingInviteToCustomersEnabled sets the isMeetingInviteToCustomersEnabled property value. Indicates whether the meeting invite is sent to the customers. The default value is false.
func (m *BookingSchedulingPolicy) SetIsMeetingInviteToCustomersEnabled(value *bool)() {
    m.isMeetingInviteToCustomersEnabled = value
}
// SetMaximumAdvance sets the maximumAdvance property value. Maximum number of days in advance that a booking can be made. It follows the ISO 8601 format.
func (m *BookingSchedulingPolicy) SetMaximumAdvance(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.maximumAdvance = value
}
// SetMinimumLeadTime sets the minimumLeadTime property value. The minimum amount of time before which bookings and cancellations must be made. It follows the ISO 8601 format.
func (m *BookingSchedulingPolicy) SetMinimumLeadTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.minimumLeadTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BookingSchedulingPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSendConfirmationsToOwner sets the sendConfirmationsToOwner property value. True to notify the business via email when a booking is created or changed. Use the email address specified in the email property of the bookingBusiness entity for the business.
func (m *BookingSchedulingPolicy) SetSendConfirmationsToOwner(value *bool)() {
    m.sendConfirmationsToOwner = value
}
// SetTimeSlotInterval sets the timeSlotInterval property value. Duration of each time slot, denoted in ISO 8601 format.
func (m *BookingSchedulingPolicy) SetTimeSlotInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.timeSlotInterval = value
}
type BookingSchedulingPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowStaffSelection()(*bool)
    GetCustomAvailabilities()([]BookingsAvailabilityWindowable)
    GetGeneralAvailability()(BookingsAvailabilityable)
    GetIsMeetingInviteToCustomersEnabled()(*bool)
    GetMaximumAdvance()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMinimumLeadTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetOdataType()(*string)
    GetSendConfirmationsToOwner()(*bool)
    GetTimeSlotInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    SetAllowStaffSelection(value *bool)()
    SetCustomAvailabilities(value []BookingsAvailabilityWindowable)()
    SetGeneralAvailability(value BookingsAvailabilityable)()
    SetIsMeetingInviteToCustomersEnabled(value *bool)()
    SetMaximumAdvance(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMinimumLeadTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetOdataType(value *string)()
    SetSendConfirmationsToOwner(value *bool)()
    SetTimeSlotInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
}
