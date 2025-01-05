terraform {
  required_version = "1.9.8"
  required_providers {
    azuread = {
      source  = "hashicorp/azuread"
      version = "3.0.2"
    }
  }
}

data "azuread_application_published_app_ids" "well_known" {}

data "azuread_service_principal" "msgraph" {
  client_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
}

data "azuread_service_principal" "arm" {
  client_id = data.azuread_application_published_app_ids.well_known.result.AzureResourceManager
}

locals {
  required_msgraph_scopes = [
    # Required for token exchange to ARM
    "offline_access",
    # Required for both Entra Groups and Roles
    "User.Read",
    "Group.Read.All",
    # Required for Entra Groups
    "PrivilegedAssignmentSchedule.ReadWrite.AzureADGroup",
    "PrivilegedEligibilitySchedule.Read.AzureADGroup",
    "RoleManagementPolicy.Read.AzureADGroup",
    # Required for Entra Roles
    "RoleEligibilitySchedule.Read.Directory",
    "RoleAssignmentSchedule.ReadWrite.Directory",
    "RoleManagementPolicy.Read.Directory",
    "PrivilegedAccess.ReadWrite.AzureAD",
  ]
}

resource "azuread_application" "this" {
  display_name                   = "pimctl"
  fallback_public_client_enabled = true

  public_client {
    redirect_uris = [
      "http://localhost:45241",
    ]
  }

  api {
    requested_access_token_version = 2
  }

  required_resource_access {
    resource_app_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph

    dynamic "resource_access" {
      for_each = local.required_msgraph_scopes
      content {
        id   = data.azuread_service_principal.msgraph.oauth2_permission_scope_ids[resource_access.value]
        type = "Scope"
      }
    }
  }

  required_resource_access {
    resource_app_id = data.azuread_application_published_app_ids.well_known.result.AzureResourceManager

    resource_access {
      id   = data.azuread_service_principal.arm.oauth2_permission_scope_ids["user_impersonation"]
      type = "Scope"
    }
  }

  tags = [
    "pimctl",
  ]
}

resource "azuread_service_principal" "this" {
  client_id = azuread_application.this.client_id
}

resource "null_resource" "required_msgraph_scopes" {
  triggers = {
    required_scopes = join(",", local.required_msgraph_scopes)
  }
}

resource "azuread_service_principal_delegated_permission_grant" "msgraph" {
  service_principal_object_id          = azuread_service_principal.this.object_id
  resource_service_principal_object_id = data.azuread_service_principal.msgraph.object_id
  claim_values                         = local.required_msgraph_scopes

  lifecycle {
    replace_triggered_by = [null_resource.required_msgraph_scopes]
  }
}

resource "azuread_service_principal_delegated_permission_grant" "arm" {
  service_principal_object_id          = azuread_service_principal.this.object_id
  resource_service_principal_object_id = data.azuread_service_principal.arm.object_id
  claim_values                         = ["user_impersonation"]
}
