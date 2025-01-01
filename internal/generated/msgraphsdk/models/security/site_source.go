package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type SiteSource struct {
    DataSource
    // The site property
    site ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Siteable
}
// NewSiteSource instantiates a new SiteSource and sets the default values.
func NewSiteSource()(*SiteSource) {
    m := &SiteSource{
        DataSource: *NewDataSource(),
    }
    odataTypeValue := "#microsoft.graph.security.siteSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSiteSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSiteSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSiteSource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SiteSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSource.GetFieldDeserializers()
    res["site"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSite(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Siteable))
        }
        return nil
    }
    return res
}
// GetSite gets the site property value. The site property
// returns a Siteable when successful
func (m *SiteSource) GetSite()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Siteable) {
    return m.site
}
// Serialize serializes information the current object
func (m *SiteSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("site", m.GetSite())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSite sets the site property value. The site property
func (m *SiteSource) SetSite(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Siteable)() {
    m.site = value
}
type SiteSourceable interface {
    DataSourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSite()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Siteable)
    SetSite(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Siteable)()
}
