package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SiteArchivalDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Represents the current archive status of the site collection. Returned only on $select. The possible values are: recentlyArchived, fullyArchived, reactivating, unknownFutureValue.
    archiveStatus *SiteArchiveStatus
    // The OdataType property
    odataType *string
}
// NewSiteArchivalDetails instantiates a new SiteArchivalDetails and sets the default values.
func NewSiteArchivalDetails()(*SiteArchivalDetails) {
    m := &SiteArchivalDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSiteArchivalDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSiteArchivalDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSiteArchivalDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SiteArchivalDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchiveStatus gets the archiveStatus property value. Represents the current archive status of the site collection. Returned only on $select. The possible values are: recentlyArchived, fullyArchived, reactivating, unknownFutureValue.
// returns a *SiteArchiveStatus when successful
func (m *SiteArchivalDetails) GetArchiveStatus()(*SiteArchiveStatus) {
    return m.archiveStatus
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SiteArchivalDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archiveStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSiteArchiveStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchiveStatus(val.(*SiteArchiveStatus))
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
func (m *SiteArchivalDetails) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SiteArchivalDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetArchiveStatus() != nil {
        cast := (*m.GetArchiveStatus()).String()
        err := writer.WriteStringValue("archiveStatus", &cast)
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
func (m *SiteArchivalDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchiveStatus sets the archiveStatus property value. Represents the current archive status of the site collection. Returned only on $select. The possible values are: recentlyArchived, fullyArchived, reactivating, unknownFutureValue.
func (m *SiteArchivalDetails) SetArchiveStatus(value *SiteArchiveStatus)() {
    m.archiveStatus = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SiteArchivalDetails) SetOdataType(value *string)() {
    m.odataType = value
}
type SiteArchivalDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchiveStatus()(*SiteArchiveStatus)
    GetOdataType()(*string)
    SetArchiveStatus(value *SiteArchiveStatus)()
    SetOdataType(value *string)()
}
