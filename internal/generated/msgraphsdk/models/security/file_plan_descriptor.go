package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type FilePlanDescriptor struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Represents the file plan descriptor of type authority applied to a particular retention label.
    authority FilePlanAuthorityable
    // Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
    authorityTemplate AuthorityTemplateable
    // The category property
    category FilePlanAppliedCategoryable
    // Specifies a group of similar types of content in a particular department.
    categoryTemplate CategoryTemplateable
    // Represents the file plan descriptor of type citation applied to a particular retention label.
    citation FilePlanCitationable
    // The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
    citationTemplate CitationTemplateable
    // Represents the file plan descriptor of type department applied to a particular retention label.
    department FilePlanDepartmentable
    // Specifies the  department or business unit of an organization to which a label belongs.
    departmentTemplate DepartmentTemplateable
    // Represents the file plan descriptor of type filePlanReference applied to a particular retention label.
    filePlanReference FilePlanReferenceable
    // Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
    filePlanReferenceTemplate FilePlanReferenceTemplateable
}
// NewFilePlanDescriptor instantiates a new FilePlanDescriptor and sets the default values.
func NewFilePlanDescriptor()(*FilePlanDescriptor) {
    m := &FilePlanDescriptor{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateFilePlanDescriptorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilePlanDescriptorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilePlanDescriptor(), nil
}
// GetAuthority gets the authority property value. Represents the file plan descriptor of type authority applied to a particular retention label.
// returns a FilePlanAuthorityable when successful
func (m *FilePlanDescriptor) GetAuthority()(FilePlanAuthorityable) {
    return m.authority
}
// GetAuthorityTemplate gets the authorityTemplate property value. Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
// returns a AuthorityTemplateable when successful
func (m *FilePlanDescriptor) GetAuthorityTemplate()(AuthorityTemplateable) {
    return m.authorityTemplate
}
// GetCategory gets the category property value. The category property
// returns a FilePlanAppliedCategoryable when successful
func (m *FilePlanDescriptor) GetCategory()(FilePlanAppliedCategoryable) {
    return m.category
}
// GetCategoryTemplate gets the categoryTemplate property value. Specifies a group of similar types of content in a particular department.
// returns a CategoryTemplateable when successful
func (m *FilePlanDescriptor) GetCategoryTemplate()(CategoryTemplateable) {
    return m.categoryTemplate
}
// GetCitation gets the citation property value. Represents the file plan descriptor of type citation applied to a particular retention label.
// returns a FilePlanCitationable when successful
func (m *FilePlanDescriptor) GetCitation()(FilePlanCitationable) {
    return m.citation
}
// GetCitationTemplate gets the citationTemplate property value. The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
// returns a CitationTemplateable when successful
func (m *FilePlanDescriptor) GetCitationTemplate()(CitationTemplateable) {
    return m.citationTemplate
}
// GetDepartment gets the department property value. Represents the file plan descriptor of type department applied to a particular retention label.
// returns a FilePlanDepartmentable when successful
func (m *FilePlanDescriptor) GetDepartment()(FilePlanDepartmentable) {
    return m.department
}
// GetDepartmentTemplate gets the departmentTemplate property value. Specifies the  department or business unit of an organization to which a label belongs.
// returns a DepartmentTemplateable when successful
func (m *FilePlanDescriptor) GetDepartmentTemplate()(DepartmentTemplateable) {
    return m.departmentTemplate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FilePlanDescriptor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFilePlanAuthorityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthority(val.(FilePlanAuthorityable))
        }
        return nil
    }
    res["authorityTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorityTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorityTemplate(val.(AuthorityTemplateable))
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFilePlanAppliedCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(FilePlanAppliedCategoryable))
        }
        return nil
    }
    res["categoryTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCategoryTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryTemplate(val.(CategoryTemplateable))
        }
        return nil
    }
    res["citation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFilePlanCitationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCitation(val.(FilePlanCitationable))
        }
        return nil
    }
    res["citationTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCitationTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCitationTemplate(val.(CitationTemplateable))
        }
        return nil
    }
    res["department"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFilePlanDepartmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val.(FilePlanDepartmentable))
        }
        return nil
    }
    res["departmentTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDepartmentTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartmentTemplate(val.(DepartmentTemplateable))
        }
        return nil
    }
    res["filePlanReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFilePlanReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePlanReference(val.(FilePlanReferenceable))
        }
        return nil
    }
    res["filePlanReferenceTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFilePlanReferenceTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePlanReferenceTemplate(val.(FilePlanReferenceTemplateable))
        }
        return nil
    }
    return res
}
// GetFilePlanReference gets the filePlanReference property value. Represents the file plan descriptor of type filePlanReference applied to a particular retention label.
// returns a FilePlanReferenceable when successful
func (m *FilePlanDescriptor) GetFilePlanReference()(FilePlanReferenceable) {
    return m.filePlanReference
}
// GetFilePlanReferenceTemplate gets the filePlanReferenceTemplate property value. Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
// returns a FilePlanReferenceTemplateable when successful
func (m *FilePlanDescriptor) GetFilePlanReferenceTemplate()(FilePlanReferenceTemplateable) {
    return m.filePlanReferenceTemplate
}
// Serialize serializes information the current object
func (m *FilePlanDescriptor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authority", m.GetAuthority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorityTemplate", m.GetAuthorityTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("categoryTemplate", m.GetCategoryTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("citation", m.GetCitation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("citationTemplate", m.GetCitationTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("departmentTemplate", m.GetDepartmentTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("filePlanReference", m.GetFilePlanReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("filePlanReferenceTemplate", m.GetFilePlanReferenceTemplate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthority sets the authority property value. Represents the file plan descriptor of type authority applied to a particular retention label.
func (m *FilePlanDescriptor) SetAuthority(value FilePlanAuthorityable)() {
    m.authority = value
}
// SetAuthorityTemplate sets the authorityTemplate property value. Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
func (m *FilePlanDescriptor) SetAuthorityTemplate(value AuthorityTemplateable)() {
    m.authorityTemplate = value
}
// SetCategory sets the category property value. The category property
func (m *FilePlanDescriptor) SetCategory(value FilePlanAppliedCategoryable)() {
    m.category = value
}
// SetCategoryTemplate sets the categoryTemplate property value. Specifies a group of similar types of content in a particular department.
func (m *FilePlanDescriptor) SetCategoryTemplate(value CategoryTemplateable)() {
    m.categoryTemplate = value
}
// SetCitation sets the citation property value. Represents the file plan descriptor of type citation applied to a particular retention label.
func (m *FilePlanDescriptor) SetCitation(value FilePlanCitationable)() {
    m.citation = value
}
// SetCitationTemplate sets the citationTemplate property value. The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
func (m *FilePlanDescriptor) SetCitationTemplate(value CitationTemplateable)() {
    m.citationTemplate = value
}
// SetDepartment sets the department property value. Represents the file plan descriptor of type department applied to a particular retention label.
func (m *FilePlanDescriptor) SetDepartment(value FilePlanDepartmentable)() {
    m.department = value
}
// SetDepartmentTemplate sets the departmentTemplate property value. Specifies the  department or business unit of an organization to which a label belongs.
func (m *FilePlanDescriptor) SetDepartmentTemplate(value DepartmentTemplateable)() {
    m.departmentTemplate = value
}
// SetFilePlanReference sets the filePlanReference property value. Represents the file plan descriptor of type filePlanReference applied to a particular retention label.
func (m *FilePlanDescriptor) SetFilePlanReference(value FilePlanReferenceable)() {
    m.filePlanReference = value
}
// SetFilePlanReferenceTemplate sets the filePlanReferenceTemplate property value. Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
func (m *FilePlanDescriptor) SetFilePlanReferenceTemplate(value FilePlanReferenceTemplateable)() {
    m.filePlanReferenceTemplate = value
}
type FilePlanDescriptorable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthority()(FilePlanAuthorityable)
    GetAuthorityTemplate()(AuthorityTemplateable)
    GetCategory()(FilePlanAppliedCategoryable)
    GetCategoryTemplate()(CategoryTemplateable)
    GetCitation()(FilePlanCitationable)
    GetCitationTemplate()(CitationTemplateable)
    GetDepartment()(FilePlanDepartmentable)
    GetDepartmentTemplate()(DepartmentTemplateable)
    GetFilePlanReference()(FilePlanReferenceable)
    GetFilePlanReferenceTemplate()(FilePlanReferenceTemplateable)
    SetAuthority(value FilePlanAuthorityable)()
    SetAuthorityTemplate(value AuthorityTemplateable)()
    SetCategory(value FilePlanAppliedCategoryable)()
    SetCategoryTemplate(value CategoryTemplateable)()
    SetCitation(value FilePlanCitationable)()
    SetCitationTemplate(value CitationTemplateable)()
    SetDepartment(value FilePlanDepartmentable)()
    SetDepartmentTemplate(value DepartmentTemplateable)()
    SetFilePlanReference(value FilePlanReferenceable)()
    SetFilePlanReferenceTemplate(value FilePlanReferenceTemplateable)()
}
