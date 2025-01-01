package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GitHubOrganizationEvidence struct {
    AlertEvidence
    // The company property
    company *string
    // The displayName property
    displayName *string
    // The email property
    email *string
    // The login property
    login *string
    // The orgId property
    orgId *string
    // The webUrl property
    webUrl *string
}
// NewGitHubOrganizationEvidence instantiates a new GitHubOrganizationEvidence and sets the default values.
func NewGitHubOrganizationEvidence()(*GitHubOrganizationEvidence) {
    m := &GitHubOrganizationEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.gitHubOrganizationEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGitHubOrganizationEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGitHubOrganizationEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitHubOrganizationEvidence(), nil
}
// GetCompany gets the company property value. The company property
// returns a *string when successful
func (m *GitHubOrganizationEvidence) GetCompany()(*string) {
    return m.company
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *GitHubOrganizationEvidence) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *GitHubOrganizationEvidence) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GitHubOrganizationEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
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
    res["orgId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrgId(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *GitHubOrganizationEvidence) GetLogin()(*string) {
    return m.login
}
// GetOrgId gets the orgId property value. The orgId property
// returns a *string when successful
func (m *GitHubOrganizationEvidence) GetOrgId()(*string) {
    return m.orgId
}
// GetWebUrl gets the webUrl property value. The webUrl property
// returns a *string when successful
func (m *GitHubOrganizationEvidence) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *GitHubOrganizationEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("company", m.GetCompany())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
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
        err = writer.WriteStringValue("orgId", m.GetOrgId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompany sets the company property value. The company property
func (m *GitHubOrganizationEvidence) SetCompany(value *string)() {
    m.company = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *GitHubOrganizationEvidence) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email property
func (m *GitHubOrganizationEvidence) SetEmail(value *string)() {
    m.email = value
}
// SetLogin sets the login property value. The login property
func (m *GitHubOrganizationEvidence) SetLogin(value *string)() {
    m.login = value
}
// SetOrgId sets the orgId property value. The orgId property
func (m *GitHubOrganizationEvidence) SetOrgId(value *string)() {
    m.orgId = value
}
// SetWebUrl sets the webUrl property value. The webUrl property
func (m *GitHubOrganizationEvidence) SetWebUrl(value *string)() {
    m.webUrl = value
}
type GitHubOrganizationEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompany()(*string)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetLogin()(*string)
    GetOrgId()(*string)
    GetWebUrl()(*string)
    SetCompany(value *string)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetLogin(value *string)()
    SetOrgId(value *string)()
    SetWebUrl(value *string)()
}
