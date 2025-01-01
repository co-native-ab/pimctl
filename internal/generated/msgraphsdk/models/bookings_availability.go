package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BookingsAvailability struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The availabilityType property
    availabilityType *BookingsServiceAvailabilityType
    // The hours of operation in a week. The business hours value is set to null if the availability type isn't customWeeklyHours.
    businessHours []BookingWorkHoursable
    // The OdataType property
    odataType *string
}
// NewBookingsAvailability instantiates a new BookingsAvailability and sets the default values.
func NewBookingsAvailability()(*BookingsAvailability) {
    m := &BookingsAvailability{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBookingsAvailabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBookingsAvailabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.bookingsAvailabilityWindow":
                        return NewBookingsAvailabilityWindow(), nil
                }
            }
        }
    }
    return NewBookingsAvailability(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BookingsAvailability) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailabilityType gets the availabilityType property value. The availabilityType property
// returns a *BookingsServiceAvailabilityType when successful
func (m *BookingsAvailability) GetAvailabilityType()(*BookingsServiceAvailabilityType) {
    return m.availabilityType
}
// GetBusinessHours gets the businessHours property value. The hours of operation in a week. The business hours value is set to null if the availability type isn't customWeeklyHours.
// returns a []BookingWorkHoursable when successful
func (m *BookingsAvailability) GetBusinessHours()([]BookingWorkHoursable) {
    return m.businessHours
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BookingsAvailability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availabilityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingsServiceAvailabilityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityType(val.(*BookingsServiceAvailabilityType))
        }
        return nil
    }
    res["businessHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingWorkHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkHoursable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingWorkHoursable)
                }
            }
            m.SetBusinessHours(res)
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
func (m *BookingsAvailability) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *BookingsAvailability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAvailabilityType() != nil {
        cast := (*m.GetAvailabilityType()).String()
        err := writer.WriteStringValue("availabilityType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessHours() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBusinessHours()))
        for i, v := range m.GetBusinessHours() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("businessHours", cast)
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
func (m *BookingsAvailability) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailabilityType sets the availabilityType property value. The availabilityType property
func (m *BookingsAvailability) SetAvailabilityType(value *BookingsServiceAvailabilityType)() {
    m.availabilityType = value
}
// SetBusinessHours sets the businessHours property value. The hours of operation in a week. The business hours value is set to null if the availability type isn't customWeeklyHours.
func (m *BookingsAvailability) SetBusinessHours(value []BookingWorkHoursable)() {
    m.businessHours = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BookingsAvailability) SetOdataType(value *string)() {
    m.odataType = value
}
type BookingsAvailabilityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityType()(*BookingsServiceAvailabilityType)
    GetBusinessHours()([]BookingWorkHoursable)
    GetOdataType()(*string)
    SetAvailabilityType(value *BookingsServiceAvailabilityType)()
    SetBusinessHours(value []BookingWorkHoursable)()
    SetOdataType(value *string)()
}
