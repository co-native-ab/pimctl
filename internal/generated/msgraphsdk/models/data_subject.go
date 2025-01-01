package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DataSubject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Email of the data subject.
    email *string
    // First name of the data subject.
    firstName *string
    // Last Name of the data subject.
    lastName *string
    // The OdataType property
    odataType *string
    // The country/region of residency. The residency information is uesed only for internal reporting but not for the content search.
    residency *string
}
// NewDataSubject instantiates a new DataSubject and sets the default values.
func NewDataSubject()(*DataSubject) {
    m := &DataSubject{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDataSubjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDataSubjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataSubject(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DataSubject) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. Email of the data subject.
// returns a *string when successful
func (m *DataSubject) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DataSubject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["firstName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["lastName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
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
    res["residency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResidency(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the firstName property value. First name of the data subject.
// returns a *string when successful
func (m *DataSubject) GetFirstName()(*string) {
    return m.firstName
}
// GetLastName gets the lastName property value. Last Name of the data subject.
// returns a *string when successful
func (m *DataSubject) GetLastName()(*string) {
    return m.lastName
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DataSubject) GetOdataType()(*string) {
    return m.odataType
}
// GetResidency gets the residency property value. The country/region of residency. The residency information is uesed only for internal reporting but not for the content search.
// returns a *string when successful
func (m *DataSubject) GetResidency()(*string) {
    return m.residency
}
// Serialize serializes information the current object
func (m *DataSubject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firstName", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastName", m.GetLastName())
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
        err := writer.WriteStringValue("residency", m.GetResidency())
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
func (m *DataSubject) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. Email of the data subject.
func (m *DataSubject) SetEmail(value *string)() {
    m.email = value
}
// SetFirstName sets the firstName property value. First name of the data subject.
func (m *DataSubject) SetFirstName(value *string)() {
    m.firstName = value
}
// SetLastName sets the lastName property value. Last Name of the data subject.
func (m *DataSubject) SetLastName(value *string)() {
    m.lastName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DataSubject) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResidency sets the residency property value. The country/region of residency. The residency information is uesed only for internal reporting but not for the content search.
func (m *DataSubject) SetResidency(value *string)() {
    m.residency = value
}
type DataSubjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetFirstName()(*string)
    GetLastName()(*string)
    GetOdataType()(*string)
    GetResidency()(*string)
    SetEmail(value *string)()
    SetFirstName(value *string)()
    SetLastName(value *string)()
    SetOdataType(value *string)()
    SetResidency(value *string)()
}
