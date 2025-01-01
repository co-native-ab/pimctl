package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEventPresenterDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Bio of the presenter.
    bio ItemBodyable
    // The presenter's company name.
    company *string
    // The presenter's job title.
    jobTitle *string
    // The presenter's LinkedIn profile URL.
    linkedInProfileWebUrl *string
    // The OdataType property
    odataType *string
    // The presenter's personal website URL.
    personalSiteWebUrl *string
    // The content stream of the presenter's photo.
    photo []byte
    // The presenter's Twitter profile URL.
    twitterProfileWebUrl *string
}
// NewVirtualEventPresenterDetails instantiates a new VirtualEventPresenterDetails and sets the default values.
func NewVirtualEventPresenterDetails()(*VirtualEventPresenterDetails) {
    m := &VirtualEventPresenterDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVirtualEventPresenterDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventPresenterDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventPresenterDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VirtualEventPresenterDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBio gets the bio property value. Bio of the presenter.
// returns a ItemBodyable when successful
func (m *VirtualEventPresenterDetails) GetBio()(ItemBodyable) {
    return m.bio
}
// GetCompany gets the company property value. The presenter's company name.
// returns a *string when successful
func (m *VirtualEventPresenterDetails) GetCompany()(*string) {
    return m.company
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEventPresenterDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBio(val.(ItemBodyable))
        }
        return nil
    }
    res["company"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompany(val)
        }
        return nil
    }
    res["jobTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["linkedInProfileWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkedInProfileWebUrl(val)
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
    res["personalSiteWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalSiteWebUrl(val)
        }
        return nil
    }
    res["photo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val)
        }
        return nil
    }
    res["twitterProfileWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwitterProfileWebUrl(val)
        }
        return nil
    }
    return res
}
// GetJobTitle gets the jobTitle property value. The presenter's job title.
// returns a *string when successful
func (m *VirtualEventPresenterDetails) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetLinkedInProfileWebUrl gets the linkedInProfileWebUrl property value. The presenter's LinkedIn profile URL.
// returns a *string when successful
func (m *VirtualEventPresenterDetails) GetLinkedInProfileWebUrl()(*string) {
    return m.linkedInProfileWebUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *VirtualEventPresenterDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetPersonalSiteWebUrl gets the personalSiteWebUrl property value. The presenter's personal website URL.
// returns a *string when successful
func (m *VirtualEventPresenterDetails) GetPersonalSiteWebUrl()(*string) {
    return m.personalSiteWebUrl
}
// GetPhoto gets the photo property value. The content stream of the presenter's photo.
// returns a []byte when successful
func (m *VirtualEventPresenterDetails) GetPhoto()([]byte) {
    return m.photo
}
// GetTwitterProfileWebUrl gets the twitterProfileWebUrl property value. The presenter's Twitter profile URL.
// returns a *string when successful
func (m *VirtualEventPresenterDetails) GetTwitterProfileWebUrl()(*string) {
    return m.twitterProfileWebUrl
}
// Serialize serializes information the current object
func (m *VirtualEventPresenterDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bio", m.GetBio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("company", m.GetCompany())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("linkedInProfileWebUrl", m.GetLinkedInProfileWebUrl())
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
        err := writer.WriteStringValue("personalSiteWebUrl", m.GetPersonalSiteWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("twitterProfileWebUrl", m.GetTwitterProfileWebUrl())
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
func (m *VirtualEventPresenterDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBio sets the bio property value. Bio of the presenter.
func (m *VirtualEventPresenterDetails) SetBio(value ItemBodyable)() {
    m.bio = value
}
// SetCompany sets the company property value. The presenter's company name.
func (m *VirtualEventPresenterDetails) SetCompany(value *string)() {
    m.company = value
}
// SetJobTitle sets the jobTitle property value. The presenter's job title.
func (m *VirtualEventPresenterDetails) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetLinkedInProfileWebUrl sets the linkedInProfileWebUrl property value. The presenter's LinkedIn profile URL.
func (m *VirtualEventPresenterDetails) SetLinkedInProfileWebUrl(value *string)() {
    m.linkedInProfileWebUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VirtualEventPresenterDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPersonalSiteWebUrl sets the personalSiteWebUrl property value. The presenter's personal website URL.
func (m *VirtualEventPresenterDetails) SetPersonalSiteWebUrl(value *string)() {
    m.personalSiteWebUrl = value
}
// SetPhoto sets the photo property value. The content stream of the presenter's photo.
func (m *VirtualEventPresenterDetails) SetPhoto(value []byte)() {
    m.photo = value
}
// SetTwitterProfileWebUrl sets the twitterProfileWebUrl property value. The presenter's Twitter profile URL.
func (m *VirtualEventPresenterDetails) SetTwitterProfileWebUrl(value *string)() {
    m.twitterProfileWebUrl = value
}
type VirtualEventPresenterDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBio()(ItemBodyable)
    GetCompany()(*string)
    GetJobTitle()(*string)
    GetLinkedInProfileWebUrl()(*string)
    GetOdataType()(*string)
    GetPersonalSiteWebUrl()(*string)
    GetPhoto()([]byte)
    GetTwitterProfileWebUrl()(*string)
    SetBio(value ItemBodyable)()
    SetCompany(value *string)()
    SetJobTitle(value *string)()
    SetLinkedInProfileWebUrl(value *string)()
    SetOdataType(value *string)()
    SetPersonalSiteWebUrl(value *string)()
    SetPhoto(value []byte)()
    SetTwitterProfileWebUrl(value *string)()
}
