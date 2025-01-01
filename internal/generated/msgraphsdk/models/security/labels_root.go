package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type LabelsRoot struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
    authorities []AuthorityTemplateable
    // Specifies a group of similar types of content in a particular department.
    categories []CategoryTemplateable
    // The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
    citations []CitationTemplateable
    // Specifies the department or business unit of an organization to which a label belongs.
    departments []DepartmentTemplateable
    // Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
    filePlanReferences []FilePlanReferenceTemplateable
    // Represents how customers can manage their data, whether and for how long to retain or delete it.
    retentionLabels []RetentionLabelable
}
// NewLabelsRoot instantiates a new LabelsRoot and sets the default values.
func NewLabelsRoot()(*LabelsRoot) {
    m := &LabelsRoot{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateLabelsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelsRoot(), nil
}
// GetAuthorities gets the authorities property value. Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
// returns a []AuthorityTemplateable when successful
func (m *LabelsRoot) GetAuthorities()([]AuthorityTemplateable) {
    return m.authorities
}
// GetCategories gets the categories property value. Specifies a group of similar types of content in a particular department.
// returns a []CategoryTemplateable when successful
func (m *LabelsRoot) GetCategories()([]CategoryTemplateable) {
    return m.categories
}
// GetCitations gets the citations property value. The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
// returns a []CitationTemplateable when successful
func (m *LabelsRoot) GetCitations()([]CitationTemplateable) {
    return m.citations
}
// GetDepartments gets the departments property value. Specifies the department or business unit of an organization to which a label belongs.
// returns a []DepartmentTemplateable when successful
func (m *LabelsRoot) GetDepartments()([]DepartmentTemplateable) {
    return m.departments
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authorities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorityTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorityTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthorityTemplateable)
                }
            }
            m.SetAuthorities(res)
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCategoryTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CategoryTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CategoryTemplateable)
                }
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["citations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCitationTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CitationTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CitationTemplateable)
                }
            }
            m.SetCitations(res)
        }
        return nil
    }
    res["departments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDepartmentTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DepartmentTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DepartmentTemplateable)
                }
            }
            m.SetDepartments(res)
        }
        return nil
    }
    res["filePlanReferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilePlanReferenceTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilePlanReferenceTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FilePlanReferenceTemplateable)
                }
            }
            m.SetFilePlanReferences(res)
        }
        return nil
    }
    res["retentionLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRetentionLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RetentionLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RetentionLabelable)
                }
            }
            m.SetRetentionLabels(res)
        }
        return nil
    }
    return res
}
// GetFilePlanReferences gets the filePlanReferences property value. Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
// returns a []FilePlanReferenceTemplateable when successful
func (m *LabelsRoot) GetFilePlanReferences()([]FilePlanReferenceTemplateable) {
    return m.filePlanReferences
}
// GetRetentionLabels gets the retentionLabels property value. Represents how customers can manage their data, whether and for how long to retain or delete it.
// returns a []RetentionLabelable when successful
func (m *LabelsRoot) GetRetentionLabels()([]RetentionLabelable) {
    return m.retentionLabels
}
// Serialize serializes information the current object
func (m *LabelsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthorities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthorities()))
        for i, v := range m.GetAuthorities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("authorities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCitations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCitations()))
        for i, v := range m.GetCitations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("citations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDepartments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDepartments()))
        for i, v := range m.GetDepartments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("departments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFilePlanReferences() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFilePlanReferences()))
        for i, v := range m.GetFilePlanReferences() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("filePlanReferences", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRetentionLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRetentionLabels()))
        for i, v := range m.GetRetentionLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("retentionLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorities sets the authorities property value. Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
func (m *LabelsRoot) SetAuthorities(value []AuthorityTemplateable)() {
    m.authorities = value
}
// SetCategories sets the categories property value. Specifies a group of similar types of content in a particular department.
func (m *LabelsRoot) SetCategories(value []CategoryTemplateable)() {
    m.categories = value
}
// SetCitations sets the citations property value. The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
func (m *LabelsRoot) SetCitations(value []CitationTemplateable)() {
    m.citations = value
}
// SetDepartments sets the departments property value. Specifies the department or business unit of an organization to which a label belongs.
func (m *LabelsRoot) SetDepartments(value []DepartmentTemplateable)() {
    m.departments = value
}
// SetFilePlanReferences sets the filePlanReferences property value. Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
func (m *LabelsRoot) SetFilePlanReferences(value []FilePlanReferenceTemplateable)() {
    m.filePlanReferences = value
}
// SetRetentionLabels sets the retentionLabels property value. Represents how customers can manage their data, whether and for how long to retain or delete it.
func (m *LabelsRoot) SetRetentionLabels(value []RetentionLabelable)() {
    m.retentionLabels = value
}
type LabelsRootable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorities()([]AuthorityTemplateable)
    GetCategories()([]CategoryTemplateable)
    GetCitations()([]CitationTemplateable)
    GetDepartments()([]DepartmentTemplateable)
    GetFilePlanReferences()([]FilePlanReferenceTemplateable)
    GetRetentionLabels()([]RetentionLabelable)
    SetAuthorities(value []AuthorityTemplateable)()
    SetCategories(value []CategoryTemplateable)()
    SetCitations(value []CitationTemplateable)()
    SetDepartments(value []DepartmentTemplateable)()
    SetFilePlanReferences(value []FilePlanReferenceTemplateable)()
    SetRetentionLabels(value []RetentionLabelable)()
}
