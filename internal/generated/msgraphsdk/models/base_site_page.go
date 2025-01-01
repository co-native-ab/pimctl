package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BaseSitePage struct {
    BaseItem
    // The name of the page layout of the page. The possible values are: microsoftReserved, article, home, unknownFutureValue.
    pageLayout *PageLayoutType
    // The publishing status and the MM.mm version of the page.
    publishingState PublicationFacetable
    // Title of the sitePage.
    title *string
}
// NewBaseSitePage instantiates a new BaseSitePage and sets the default values.
func NewBaseSitePage()(*BaseSitePage) {
    m := &BaseSitePage{
        BaseItem: *NewBaseItem(),
    }
    odataTypeValue := "#microsoft.graph.baseSitePage"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBaseSitePageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBaseSitePageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.sitePage":
                        return NewSitePage(), nil
                }
            }
        }
    }
    return NewBaseSitePage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BaseSitePage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["pageLayout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePageLayoutType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageLayout(val.(*PageLayoutType))
        }
        return nil
    }
    res["publishingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicationFacetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingState(val.(PublicationFacetable))
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
// GetPageLayout gets the pageLayout property value. The name of the page layout of the page. The possible values are: microsoftReserved, article, home, unknownFutureValue.
// returns a *PageLayoutType when successful
func (m *BaseSitePage) GetPageLayout()(*PageLayoutType) {
    return m.pageLayout
}
// GetPublishingState gets the publishingState property value. The publishing status and the MM.mm version of the page.
// returns a PublicationFacetable when successful
func (m *BaseSitePage) GetPublishingState()(PublicationFacetable) {
    return m.publishingState
}
// GetTitle gets the title property value. Title of the sitePage.
// returns a *string when successful
func (m *BaseSitePage) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *BaseSitePage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPageLayout() != nil {
        cast := (*m.GetPageLayout()).String()
        err = writer.WriteStringValue("pageLayout", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publishingState", m.GetPublishingState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPageLayout sets the pageLayout property value. The name of the page layout of the page. The possible values are: microsoftReserved, article, home, unknownFutureValue.
func (m *BaseSitePage) SetPageLayout(value *PageLayoutType)() {
    m.pageLayout = value
}
// SetPublishingState sets the publishingState property value. The publishing status and the MM.mm version of the page.
func (m *BaseSitePage) SetPublishingState(value PublicationFacetable)() {
    m.publishingState = value
}
// SetTitle sets the title property value. Title of the sitePage.
func (m *BaseSitePage) SetTitle(value *string)() {
    m.title = value
}
type BaseSitePageable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPageLayout()(*PageLayoutType)
    GetPublishingState()(PublicationFacetable)
    GetTitle()(*string)
    SetPageLayout(value *PageLayoutType)()
    SetPublishingState(value PublicationFacetable)()
    SetTitle(value *string)()
}
