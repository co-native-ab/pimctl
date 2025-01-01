package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OnTokenIssuanceStartReturnClaim struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The identifier of the claim returned by an API that is to be add to a token being issued.
    claimIdInApiResponse *string
    // The OdataType property
    odataType *string
}
// NewOnTokenIssuanceStartReturnClaim instantiates a new OnTokenIssuanceStartReturnClaim and sets the default values.
func NewOnTokenIssuanceStartReturnClaim()(*OnTokenIssuanceStartReturnClaim) {
    m := &OnTokenIssuanceStartReturnClaim{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnTokenIssuanceStartReturnClaimFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnTokenIssuanceStartReturnClaimFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnTokenIssuanceStartReturnClaim(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OnTokenIssuanceStartReturnClaim) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClaimIdInApiResponse gets the claimIdInApiResponse property value. The identifier of the claim returned by an API that is to be add to a token being issued.
// returns a *string when successful
func (m *OnTokenIssuanceStartReturnClaim) GetClaimIdInApiResponse()(*string) {
    return m.claimIdInApiResponse
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnTokenIssuanceStartReturnClaim) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["claimIdInApiResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClaimIdInApiResponse(val)
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
func (m *OnTokenIssuanceStartReturnClaim) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OnTokenIssuanceStartReturnClaim) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("claimIdInApiResponse", m.GetClaimIdInApiResponse())
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
func (m *OnTokenIssuanceStartReturnClaim) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClaimIdInApiResponse sets the claimIdInApiResponse property value. The identifier of the claim returned by an API that is to be add to a token being issued.
func (m *OnTokenIssuanceStartReturnClaim) SetClaimIdInApiResponse(value *string)() {
    m.claimIdInApiResponse = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnTokenIssuanceStartReturnClaim) SetOdataType(value *string)() {
    m.odataType = value
}
type OnTokenIssuanceStartReturnClaimable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClaimIdInApiResponse()(*string)
    GetOdataType()(*string)
    SetClaimIdInApiResponse(value *string)()
    SetOdataType(value *string)()
}
