package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Thumbnail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content stream for the thumbnail.
    content []byte
    // The height of the thumbnail, in pixels.
    height *int32
    // The OdataType property
    odataType *string
    // The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested.
    sourceItemId *string
    // The URL used to fetch the thumbnail content.
    url *string
    // The width of the thumbnail, in pixels.
    width *int32
}
// NewThumbnail instantiates a new Thumbnail and sets the default values.
func NewThumbnail()(*Thumbnail) {
    m := &Thumbnail{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateThumbnailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateThumbnailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThumbnail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Thumbnail) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content stream for the thumbnail.
// returns a []byte when successful
func (m *Thumbnail) GetContent()([]byte) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Thumbnail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
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
    res["sourceItemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceItemId(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
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
// GetHeight gets the height property value. The height of the thumbnail, in pixels.
// returns a *int32 when successful
func (m *Thumbnail) GetHeight()(*int32) {
    return m.height
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *Thumbnail) GetOdataType()(*string) {
    return m.odataType
}
// GetSourceItemId gets the sourceItemId property value. The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested.
// returns a *string when successful
func (m *Thumbnail) GetSourceItemId()(*string) {
    return m.sourceItemId
}
// GetUrl gets the url property value. The URL used to fetch the thumbnail content.
// returns a *string when successful
func (m *Thumbnail) GetUrl()(*string) {
    return m.url
}
// GetWidth gets the width property value. The width of the thumbnail, in pixels.
// returns a *int32 when successful
func (m *Thumbnail) GetWidth()(*int32) {
    return m.width
}
// Serialize serializes information the current object
func (m *Thumbnail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteStringValue("sourceItemId", m.GetSourceItemId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *Thumbnail) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content stream for the thumbnail.
func (m *Thumbnail) SetContent(value []byte)() {
    m.content = value
}
// SetHeight sets the height property value. The height of the thumbnail, in pixels.
func (m *Thumbnail) SetHeight(value *int32)() {
    m.height = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Thumbnail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSourceItemId sets the sourceItemId property value. The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested.
func (m *Thumbnail) SetSourceItemId(value *string)() {
    m.sourceItemId = value
}
// SetUrl sets the url property value. The URL used to fetch the thumbnail content.
func (m *Thumbnail) SetUrl(value *string)() {
    m.url = value
}
// SetWidth sets the width property value. The width of the thumbnail, in pixels.
func (m *Thumbnail) SetWidth(value *int32)() {
    m.width = value
}
type Thumbnailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetHeight()(*int32)
    GetOdataType()(*string)
    GetSourceItemId()(*string)
    GetUrl()(*string)
    GetWidth()(*int32)
    SetContent(value []byte)()
    SetHeight(value *int32)()
    SetOdataType(value *string)()
    SetSourceItemId(value *string)()
    SetUrl(value *string)()
    SetWidth(value *int32)()
}
