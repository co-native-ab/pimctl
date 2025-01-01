package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GitHubRepoEvidence struct {
    AlertEvidence
    // The baseUrl property
    baseUrl *string
    // The login property
    login *string
    // The owner property
    owner *string
    // The ownerType property
    ownerType *string
    // The repoId property
    repoId *string
}
// NewGitHubRepoEvidence instantiates a new GitHubRepoEvidence and sets the default values.
func NewGitHubRepoEvidence()(*GitHubRepoEvidence) {
    m := &GitHubRepoEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.gitHubRepoEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGitHubRepoEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGitHubRepoEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitHubRepoEvidence(), nil
}
// GetBaseUrl gets the baseUrl property value. The baseUrl property
// returns a *string when successful
func (m *GitHubRepoEvidence) GetBaseUrl()(*string) {
    return m.baseUrl
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GitHubRepoEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["baseUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseUrl(val)
        }
        return nil
    }
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val)
        }
        return nil
    }
    res["repoId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepoId(val)
        }
        return nil
    }
    return res
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *GitHubRepoEvidence) GetLogin()(*string) {
    return m.login
}
// GetOwner gets the owner property value. The owner property
// returns a *string when successful
func (m *GitHubRepoEvidence) GetOwner()(*string) {
    return m.owner
}
// GetOwnerType gets the ownerType property value. The ownerType property
// returns a *string when successful
func (m *GitHubRepoEvidence) GetOwnerType()(*string) {
    return m.ownerType
}
// GetRepoId gets the repoId property value. The repoId property
// returns a *string when successful
func (m *GitHubRepoEvidence) GetRepoId()(*string) {
    return m.repoId
}
// Serialize serializes information the current object
func (m *GitHubRepoEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("baseUrl", m.GetBaseUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("login", m.GetLogin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerType", m.GetOwnerType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("repoId", m.GetRepoId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBaseUrl sets the baseUrl property value. The baseUrl property
func (m *GitHubRepoEvidence) SetBaseUrl(value *string)() {
    m.baseUrl = value
}
// SetLogin sets the login property value. The login property
func (m *GitHubRepoEvidence) SetLogin(value *string)() {
    m.login = value
}
// SetOwner sets the owner property value. The owner property
func (m *GitHubRepoEvidence) SetOwner(value *string)() {
    m.owner = value
}
// SetOwnerType sets the ownerType property value. The ownerType property
func (m *GitHubRepoEvidence) SetOwnerType(value *string)() {
    m.ownerType = value
}
// SetRepoId sets the repoId property value. The repoId property
func (m *GitHubRepoEvidence) SetRepoId(value *string)() {
    m.repoId = value
}
type GitHubRepoEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseUrl()(*string)
    GetLogin()(*string)
    GetOwner()(*string)
    GetOwnerType()(*string)
    GetRepoId()(*string)
    SetBaseUrl(value *string)()
    SetLogin(value *string)()
    SetOwner(value *string)()
    SetOwnerType(value *string)()
    SetRepoId(value *string)()
}
