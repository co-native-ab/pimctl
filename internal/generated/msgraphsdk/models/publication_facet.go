package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PublicationFacet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The user who checked out the file.
    checkedOutBy IdentitySetable
    // The state of publication for this document. Either published or checkout. Read-only.
    level *string
    // The OdataType property
    odataType *string
    // The unique identifier for the version that is visible to the current caller. Read-only.
    versionId *string
}
// NewPublicationFacet instantiates a new PublicationFacet and sets the default values.
func NewPublicationFacet()(*PublicationFacet) {
    m := &PublicationFacet{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePublicationFacetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePublicationFacetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublicationFacet(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PublicationFacet) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckedOutBy gets the checkedOutBy property value. The user who checked out the file.
// returns a IdentitySetable when successful
func (m *PublicationFacet) GetCheckedOutBy()(IdentitySetable) {
    return m.checkedOutBy
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PublicationFacet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checkedOutBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckedOutBy(val.(IdentitySetable))
        }
        return nil
    }
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
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
    res["versionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionId(val)
        }
        return nil
    }
    return res
}
// GetLevel gets the level property value. The state of publication for this document. Either published or checkout. Read-only.
// returns a *string when successful
func (m *PublicationFacet) GetLevel()(*string) {
    return m.level
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *PublicationFacet) GetOdataType()(*string) {
    return m.odataType
}
// GetVersionId gets the versionId property value. The unique identifier for the version that is visible to the current caller. Read-only.
// returns a *string when successful
func (m *PublicationFacet) GetVersionId()(*string) {
    return m.versionId
}
// Serialize serializes information the current object
func (m *PublicationFacet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("checkedOutBy", m.GetCheckedOutBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("level", m.GetLevel())
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
        err := writer.WriteStringValue("versionId", m.GetVersionId())
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
func (m *PublicationFacet) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckedOutBy sets the checkedOutBy property value. The user who checked out the file.
func (m *PublicationFacet) SetCheckedOutBy(value IdentitySetable)() {
    m.checkedOutBy = value
}
// SetLevel sets the level property value. The state of publication for this document. Either published or checkout. Read-only.
func (m *PublicationFacet) SetLevel(value *string)() {
    m.level = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PublicationFacet) SetOdataType(value *string)() {
    m.odataType = value
}
// SetVersionId sets the versionId property value. The unique identifier for the version that is visible to the current caller. Read-only.
func (m *PublicationFacet) SetVersionId(value *string)() {
    m.versionId = value
}
type PublicationFacetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckedOutBy()(IdentitySetable)
    GetLevel()(*string)
    GetOdataType()(*string)
    GetVersionId()(*string)
    SetCheckedOutBy(value IdentitySetable)()
    SetLevel(value *string)()
    SetOdataType(value *string)()
    SetVersionId(value *string)()
}
