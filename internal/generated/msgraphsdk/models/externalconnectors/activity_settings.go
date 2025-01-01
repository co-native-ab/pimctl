package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActivitySettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Specifies configurations to identify an externalItem based on a shared URL.
    urlToItemResolvers []UrlToItemResolverBaseable
}
// NewActivitySettings instantiates a new ActivitySettings and sets the default values.
func NewActivitySettings()(*ActivitySettings) {
    m := &ActivitySettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActivitySettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActivitySettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActivitySettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActivitySettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActivitySettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["urlToItemResolvers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUrlToItemResolverBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UrlToItemResolverBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UrlToItemResolverBaseable)
                }
            }
            m.SetUrlToItemResolvers(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ActivitySettings) GetOdataType()(*string) {
    return m.odataType
}
// GetUrlToItemResolvers gets the urlToItemResolvers property value. Specifies configurations to identify an externalItem based on a shared URL.
// returns a []UrlToItemResolverBaseable when successful
func (m *ActivitySettings) GetUrlToItemResolvers()([]UrlToItemResolverBaseable) {
    return m.urlToItemResolvers
}
// Serialize serializes information the current object
func (m *ActivitySettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetUrlToItemResolvers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUrlToItemResolvers()))
        for i, v := range m.GetUrlToItemResolvers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("urlToItemResolvers", cast)
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
func (m *ActivitySettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ActivitySettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUrlToItemResolvers sets the urlToItemResolvers property value. Specifies configurations to identify an externalItem based on a shared URL.
func (m *ActivitySettings) SetUrlToItemResolvers(value []UrlToItemResolverBaseable)() {
    m.urlToItemResolvers = value
}
type ActivitySettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetUrlToItemResolvers()([]UrlToItemResolverBaseable)
    SetOdataType(value *string)()
    SetUrlToItemResolvers(value []UrlToItemResolverBaseable)()
}
