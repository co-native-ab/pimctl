package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServerProcessedContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A key-value map where keys are string identifiers and values are rich text with HTML format. SharePoint servers treat the values as HTML content and run services like safety checks, search index and link fixup on them.
    htmlStrings []MetaDataKeyStringPairable
    // A key-value map where keys are string identifiers and values are image sources. SharePoint servers treat the values as image sources and run services like search index and link fixup on them.
    imageSources []MetaDataKeyStringPairable
    // A key-value map where keys are string identifiers and values are links. SharePoint servers treat the values as links and run services like link fixup on them.
    links []MetaDataKeyStringPairable
    // The OdataType property
    odataType *string
    // A key-value map where keys are string identifiers and values are strings that should be search indexed.
    searchablePlainTexts []MetaDataKeyStringPairable
}
// NewServerProcessedContent instantiates a new ServerProcessedContent and sets the default values.
func NewServerProcessedContent()(*ServerProcessedContent) {
    m := &ServerProcessedContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServerProcessedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServerProcessedContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerProcessedContent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServerProcessedContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServerProcessedContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["htmlStrings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyStringPairable)
                }
            }
            m.SetHtmlStrings(res)
        }
        return nil
    }
    res["imageSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyStringPairable)
                }
            }
            m.SetImageSources(res)
        }
        return nil
    }
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyStringPairable)
                }
            }
            m.SetLinks(res)
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
    res["searchablePlainTexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MetaDataKeyStringPairable)
                }
            }
            m.SetSearchablePlainTexts(res)
        }
        return nil
    }
    return res
}
// GetHtmlStrings gets the htmlStrings property value. A key-value map where keys are string identifiers and values are rich text with HTML format. SharePoint servers treat the values as HTML content and run services like safety checks, search index and link fixup on them.
// returns a []MetaDataKeyStringPairable when successful
func (m *ServerProcessedContent) GetHtmlStrings()([]MetaDataKeyStringPairable) {
    return m.htmlStrings
}
// GetImageSources gets the imageSources property value. A key-value map where keys are string identifiers and values are image sources. SharePoint servers treat the values as image sources and run services like search index and link fixup on them.
// returns a []MetaDataKeyStringPairable when successful
func (m *ServerProcessedContent) GetImageSources()([]MetaDataKeyStringPairable) {
    return m.imageSources
}
// GetLinks gets the links property value. A key-value map where keys are string identifiers and values are links. SharePoint servers treat the values as links and run services like link fixup on them.
// returns a []MetaDataKeyStringPairable when successful
func (m *ServerProcessedContent) GetLinks()([]MetaDataKeyStringPairable) {
    return m.links
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ServerProcessedContent) GetOdataType()(*string) {
    return m.odataType
}
// GetSearchablePlainTexts gets the searchablePlainTexts property value. A key-value map where keys are string identifiers and values are strings that should be search indexed.
// returns a []MetaDataKeyStringPairable when successful
func (m *ServerProcessedContent) GetSearchablePlainTexts()([]MetaDataKeyStringPairable) {
    return m.searchablePlainTexts
}
// Serialize serializes information the current object
func (m *ServerProcessedContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHtmlStrings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHtmlStrings()))
        for i, v := range m.GetHtmlStrings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("htmlStrings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImageSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImageSources()))
        for i, v := range m.GetImageSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("imageSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
        for i, v := range m.GetLinks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("links", cast)
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
    if m.GetSearchablePlainTexts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSearchablePlainTexts()))
        for i, v := range m.GetSearchablePlainTexts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("searchablePlainTexts", cast)
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
func (m *ServerProcessedContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHtmlStrings sets the htmlStrings property value. A key-value map where keys are string identifiers and values are rich text with HTML format. SharePoint servers treat the values as HTML content and run services like safety checks, search index and link fixup on them.
func (m *ServerProcessedContent) SetHtmlStrings(value []MetaDataKeyStringPairable)() {
    m.htmlStrings = value
}
// SetImageSources sets the imageSources property value. A key-value map where keys are string identifiers and values are image sources. SharePoint servers treat the values as image sources and run services like search index and link fixup on them.
func (m *ServerProcessedContent) SetImageSources(value []MetaDataKeyStringPairable)() {
    m.imageSources = value
}
// SetLinks sets the links property value. A key-value map where keys are string identifiers and values are links. SharePoint servers treat the values as links and run services like link fixup on them.
func (m *ServerProcessedContent) SetLinks(value []MetaDataKeyStringPairable)() {
    m.links = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ServerProcessedContent) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSearchablePlainTexts sets the searchablePlainTexts property value. A key-value map where keys are string identifiers and values are strings that should be search indexed.
func (m *ServerProcessedContent) SetSearchablePlainTexts(value []MetaDataKeyStringPairable)() {
    m.searchablePlainTexts = value
}
type ServerProcessedContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtmlStrings()([]MetaDataKeyStringPairable)
    GetImageSources()([]MetaDataKeyStringPairable)
    GetLinks()([]MetaDataKeyStringPairable)
    GetOdataType()(*string)
    GetSearchablePlainTexts()([]MetaDataKeyStringPairable)
    SetHtmlStrings(value []MetaDataKeyStringPairable)()
    SetImageSources(value []MetaDataKeyStringPairable)()
    SetLinks(value []MetaDataKeyStringPairable)()
    SetOdataType(value *string)()
    SetSearchablePlainTexts(value []MetaDataKeyStringPairable)()
}
