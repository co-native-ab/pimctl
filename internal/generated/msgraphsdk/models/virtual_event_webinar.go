package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEventWebinar struct {
    VirtualEvent
    // To whom the webinar is visible. Possible values are: everyone, organization, and unknownFutureValue.
    audience *MeetingAudience
    // Identity information of coorganizers of the webinar.
    coOrganizers []CommunicationsUserIdentityable
    // Registration configuration of the webinar.
    registrationConfiguration VirtualEventWebinarRegistrationConfigurationable
    // Registration records of the webinar.
    registrations []VirtualEventRegistrationable
}
// NewVirtualEventWebinar instantiates a new VirtualEventWebinar and sets the default values.
func NewVirtualEventWebinar()(*VirtualEventWebinar) {
    m := &VirtualEventWebinar{
        VirtualEvent: *NewVirtualEvent(),
    }
    odataTypeValue := "#microsoft.graph.virtualEventWebinar"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateVirtualEventWebinarFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventWebinarFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventWebinar(), nil
}
// GetAudience gets the audience property value. To whom the webinar is visible. Possible values are: everyone, organization, and unknownFutureValue.
// returns a *MeetingAudience when successful
func (m *VirtualEventWebinar) GetAudience()(*MeetingAudience) {
    return m.audience
}
// GetCoOrganizers gets the coOrganizers property value. Identity information of coorganizers of the webinar.
// returns a []CommunicationsUserIdentityable when successful
func (m *VirtualEventWebinar) GetCoOrganizers()([]CommunicationsUserIdentityable) {
    return m.coOrganizers
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEventWebinar) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VirtualEvent.GetFieldDeserializers()
    res["audience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingAudience)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudience(val.(*MeetingAudience))
        }
        return nil
    }
    res["coOrganizers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCommunicationsUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CommunicationsUserIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CommunicationsUserIdentityable)
                }
            }
            m.SetCoOrganizers(res)
        }
        return nil
    }
    res["registrationConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEventWebinarRegistrationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationConfiguration(val.(VirtualEventWebinarRegistrationConfigurationable))
        }
        return nil
    }
    res["registrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventRegistrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventRegistrationable)
                }
            }
            m.SetRegistrations(res)
        }
        return nil
    }
    return res
}
// GetRegistrationConfiguration gets the registrationConfiguration property value. Registration configuration of the webinar.
// returns a VirtualEventWebinarRegistrationConfigurationable when successful
func (m *VirtualEventWebinar) GetRegistrationConfiguration()(VirtualEventWebinarRegistrationConfigurationable) {
    return m.registrationConfiguration
}
// GetRegistrations gets the registrations property value. Registration records of the webinar.
// returns a []VirtualEventRegistrationable when successful
func (m *VirtualEventWebinar) GetRegistrations()([]VirtualEventRegistrationable) {
    return m.registrations
}
// Serialize serializes information the current object
func (m *VirtualEventWebinar) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VirtualEvent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAudience() != nil {
        cast := (*m.GetAudience()).String()
        err = writer.WriteStringValue("audience", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCoOrganizers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCoOrganizers()))
        for i, v := range m.GetCoOrganizers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("coOrganizers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("registrationConfiguration", m.GetRegistrationConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetRegistrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegistrations()))
        for i, v := range m.GetRegistrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("registrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudience sets the audience property value. To whom the webinar is visible. Possible values are: everyone, organization, and unknownFutureValue.
func (m *VirtualEventWebinar) SetAudience(value *MeetingAudience)() {
    m.audience = value
}
// SetCoOrganizers sets the coOrganizers property value. Identity information of coorganizers of the webinar.
func (m *VirtualEventWebinar) SetCoOrganizers(value []CommunicationsUserIdentityable)() {
    m.coOrganizers = value
}
// SetRegistrationConfiguration sets the registrationConfiguration property value. Registration configuration of the webinar.
func (m *VirtualEventWebinar) SetRegistrationConfiguration(value VirtualEventWebinarRegistrationConfigurationable)() {
    m.registrationConfiguration = value
}
// SetRegistrations sets the registrations property value. Registration records of the webinar.
func (m *VirtualEventWebinar) SetRegistrations(value []VirtualEventRegistrationable)() {
    m.registrations = value
}
type VirtualEventWebinarable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEventable
    GetAudience()(*MeetingAudience)
    GetCoOrganizers()([]CommunicationsUserIdentityable)
    GetRegistrationConfiguration()(VirtualEventWebinarRegistrationConfigurationable)
    GetRegistrations()([]VirtualEventRegistrationable)
    SetAudience(value *MeetingAudience)()
    SetCoOrganizers(value []CommunicationsUserIdentityable)()
    SetRegistrationConfiguration(value VirtualEventWebinarRegistrationConfigurationable)()
    SetRegistrations(value []VirtualEventRegistrationable)()
}
