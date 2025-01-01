package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Image struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Optional. Height of the image, in pixels. Read-only.
    height *int32
    // The OdataType property
    odataType *string
    // Optional. Width of the image, in pixels. Read-only.
    width *int32
}
// NewImage instantiates a new Image and sets the default values.
func NewImage()(*Image) {
    m := &Image{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Image) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Image) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["height"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
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
    res["width"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
// GetHeight gets the height property value. Optional. Height of the image, in pixels. Read-only.
// returns a *int32 when successful
func (m *Image) GetHeight()(*int32) {
    return m.height
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *Image) GetOdataType()(*string) {
    return m.odataType
}
// GetWidth gets the width property value. Optional. Width of the image, in pixels. Read-only.
// returns a *int32 when successful
func (m *Image) GetWidth()(*int32) {
    return m.width
}
// Serialize serializes information the current object
func (m *Image) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("height", m.GetHeight())
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
        err := writer.WriteInt32Value("width", m.GetWidth())
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
func (m *Image) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHeight sets the height property value. Optional. Height of the image, in pixels. Read-only.
func (m *Image) SetHeight(value *int32)() {
    m.height = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Image) SetOdataType(value *string)() {
    m.odataType = value
}
// SetWidth sets the width property value. Optional. Width of the image, in pixels. Read-only.
func (m *Image) SetWidth(value *int32)() {
    m.width = value
}
type Imageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeight()(*int32)
    GetOdataType()(*string)
    GetWidth()(*int32)
    SetHeight(value *int32)()
    SetOdataType(value *string)()
    SetWidth(value *int32)()
}
