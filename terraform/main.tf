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

data "azuread_service_principal" "ms_graph" {
  client_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
}

locals {
  required_scopes = [
    "Group.Read.All",
    "PrivilegedAssignmentSchedule.ReadWrite.AzureADGroup",
    "PrivilegedEligibilitySchedule.Read.AzureADGroup",
    "RoleManagementPolicy.Read.AzureADGroup",
    "User.Read",
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
      for_each = local.required_scopes
      content {
        id   = data.azuread_service_principal.ms_graph.oauth2_permission_scope_ids[resource_access.value]
        type = "Scope"
      }
    }
  }

  tags = [
    "pimctl",
  ]
}

resource "azuread_service_principal" "this" {
  client_id = azuread_application.this.client_id
}

resource "azuread_service_principal_delegated_permission_grant" "this" {
  service_principal_object_id          = azuread_service_principal.this.object_id
  resource_service_principal_object_id = data.azuread_service_principal.ms_graph.object_id
  claim_values                         = local.required_scopes
}
