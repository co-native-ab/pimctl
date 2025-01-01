package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WebPartData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Data version of the web part. The value is defined by the web part developer. Different dataVersions usually refers to a different property structure.
    dataVersion *string
    // Description of the web part.
    description *string
    // The OdataType property
    odataType *string
    // Properties bag of the web part.
    properties i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Contains collections of data that can be processed by server side services like search index and link fixup.
    serverProcessedContent ServerProcessedContentable
    // Title of the web part.
    title *string
}
// NewWebPartData instantiates a new WebPartData and sets the default values.
func NewWebPartData()(*WebPartData) {
    m := &WebPartData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebPartDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWebPartDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebPartData(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WebPartData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDataVersion gets the dataVersion property value. Data version of the web part. The value is defined by the web part developer. Different dataVersions usually refers to a different property structure.
// returns a *string when successful
func (m *WebPartData) GetDataVersion()(*string) {
    return m.dataVersion
}
// GetDescription gets the description property value. Description of the web part.
// returns a *string when successful
func (m *WebPartData) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WebPartData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dataVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataVersion(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperties(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["serverProcessedContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerProcessedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerProcessedContent(val.(ServerProcessedContentable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WebPartData) GetOdataType()(*string) {
    return m.odataType
}
// GetProperties gets the properties property value. Properties bag of the web part.
// returns a UntypedNodeable when successful
func (m *WebPartData) GetProperties()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.properties
}
// GetServerProcessedContent gets the serverProcessedContent property value. Contains collections of data that can be processed by server side services like search index and link fixup.
// returns a ServerProcessedContentable when successful
func (m *WebPartData) GetServerProcessedContent()(ServerProcessedContentable) {
    return m.serverProcessedContent
}
// GetTitle gets the title property value. Title of the web part.
// returns a *string when successful
func (m *WebPartData) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *WebPartData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dataVersion", m.GetDataVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteObjectValue("properties", m.GetProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("serverProcessedContent", m.GetServerProcessedContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *WebPartData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDataVersion sets the dataVersion property value. Data version of the web part. The value is defined by the web part developer. Different dataVersions usually refers to a different property structure.
func (m *WebPartData) SetDataVersion(value *string)() {
    m.dataVersion = value
}
// SetDescription sets the description property value. Description of the web part.
func (m *WebPartData) SetDescription(value *string)() {
    m.description = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WebPartData) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProperties sets the properties property value. Properties bag of the web part.
func (m *WebPartData) SetProperties(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.properties = value
}
// SetServerProcessedContent sets the serverProcessedContent property value. Contains collections of data that can be processed by server side services like search index and link fixup.
func (m *WebPartData) SetServerProcessedContent(value ServerProcessedContentable)() {
    m.serverProcessedContent = value
}
// SetTitle sets the title property value. Title of the web part.
func (m *WebPartData) SetTitle(value *string)() {
    m.title = value
}
type WebPartDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataVersion()(*string)
    GetDescription()(*string)
    GetOdataType()(*string)
    GetProperties()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetServerProcessedContent()(ServerProcessedContentable)
    GetTitle()(*string)
    SetDataVersion(value *string)()
    SetDescription(value *string)()
    SetOdataType(value *string)()
    SetProperties(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetServerProcessedContent(value ServerProcessedContentable)()
    SetTitle(value *string)()
}
