package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GeoLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The city property
    city *string
    // The countryName property
    countryName *string
    // The latitude property
    latitude *float64
    // The longitude property
    longitude *float64
    // The OdataType property
    odataType *string
    // The state property
    state *string
}
// NewGeoLocation instantiates a new GeoLocation and sets the default values.
func NewGeoLocation()(*GeoLocation) {
    m := &GeoLocation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGeoLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGeoLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGeoLocation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GeoLocation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCity gets the city property value. The city property
// returns a *string when successful
func (m *GeoLocation) GetCity()(*string) {
    return m.city
}
// GetCountryName gets the countryName property value. The countryName property
// returns a *string when successful
func (m *GeoLocation) GetCountryName()(*string) {
    return m.countryName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GeoLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryName(val)
        }
        return nil
    }
    res["latitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatitude(val)
        }
        return nil
    }
    res["longitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongitude(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetLatitude gets the latitude property value. The latitude property
// returns a *float64 when successful
func (m *GeoLocation) GetLatitude()(*float64) {
    return m.latitude
}
// GetLongitude gets the longitude property value. The longitude property
// returns a *float64 when successful
func (m *GeoLocation) GetLongitude()(*float64) {
    return m.longitude
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *GeoLocation) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *GeoLocation) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *GeoLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("countryName", m.GetCountryName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("latitude", m.GetLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("longitude", m.GetLongitude())
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
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *GeoLocation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCity sets the city property value. The city property
func (m *GeoLocation) SetCity(value *string)() {
    m.city = value
}
// SetCountryName sets the countryName property value. The countryName property
func (m *GeoLocation) SetCountryName(value *string)() {
    m.countryName = value
}
// SetLatitude sets the latitude property value. The latitude property
func (m *GeoLocation) SetLatitude(value *float64)() {
    m.latitude = value
}
// SetLongitude sets the longitude property value. The longitude property
func (m *GeoLocation) SetLongitude(value *float64)() {
    m.longitude = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GeoLocation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. The state property
func (m *GeoLocation) SetState(value *string)() {
    m.state = value
}
type GeoLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCity()(*string)
    GetCountryName()(*string)
    GetLatitude()(*float64)
    GetLongitude()(*float64)
    GetOdataType()(*string)
    GetState()(*string)
    SetCity(value *string)()
    SetCountryName(value *string)()
    SetLatitude(value *float64)()
    SetLongitude(value *float64)()
    SetOdataType(value *string)()
    SetState(value *string)()
}
