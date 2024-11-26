# pimctl

pimctl is a tool to manage Entra PIM (Privileged Identity Management) activations and approvals from the terminal.

<p align="center">
  <img src="./assets/pimctl_menu_demo.gif">
</p>

## Background

We are building an environment where anything above read access requires developers ("users of Azure") to request assignment to become members of Entra Groups using PIM and also approve other people who have requested assignments. It can be a bit of an overhead leaving the IDE and going through the Azure portal to request and approve these assignments, and it would create less friction if it was possible to handle this from the terminal.

After looking into the different options out there, none of the ones we found matched our full needs, and the idea of pimctl was born. It's built for people who have a terminal in front of them most of the time and work in environments where they regularly need to request and approve membership assignments to Entra Groups and tries to streamline that workflow.

## Alpha Version

This current version of pimctl is in a very early stage. We want to do more, test more, and make it stable. The subcommands, flags, and everything in between will most likely change a couple of times, but before we do too much, we want to see what other people need.

As Moltke may have said, if he spoke English and wrote code:

> No CLI tool survives first contact with developers typing commands you never thought they would.

## Installing

### Windows

```powershell
$LatestTag = (Invoke-RestMethod -Uri "https://api.github.com/repos/co-native-ab/pimctl/releases" `
    | Where-Object { $_.prerelease -eq $true }) `
    | Select-Object -First 1 -ExpandProperty tag_name
$Version = $LatestTag -replace "^v", ""
$Arch = (Get-CimInstance Win32_ComputerSystem).SystemType
switch ($Arch) {
    "x64-based PC" { $Arch = "amd64" }
    "ARM64-based PC" { $Arch = "arm64" }
    default { throw "Unsupported architecture: $Arch" }
}
$TempDir = New-TemporaryFile | % { Remove-Item $_; New-Item -ItemType Directory -Path $_ }
Invoke-WebRequest "https://github.com/co-native-ab/pimctl/releases/download/v${Version}/pimctl_${Version}_windows_${Arch}.zip" -OutFile "${TempDir}\pimctl_${Version}_windows_${Arch}.zip"
Expand-Archive "${TempDir}\pimctl_${Version}_windows_${Arch}.zip" -DestinationPath "${TempDir}"
Move-Item "${TempDir}\pimctl.exe" "${ENV:LOCALAPPDATA}\Microsoft\WindowsApps\"
New-Item -Type Directory -Path "${ENV:USERPROFILE}\.IdentityService"
```

### Linux

```shell
LATEST_TAG=$(curl -s https://api.github.com/repos/co-native-ab/pimctl/releases | jq -r 'map_values(select(.prerelease == true)) | first(.[].tag_name)')
VERSION=${LATEST_TAG#"v"}
ARCH=$(uname -m | sed -e 's/x86_64/amd64/' -e 's/armv[0-9]*/&/' -e 's/aarch64/arm64/')
TEMP_DIR=$(mktemp -d)
curl -L "https://github.com/co-native-ab/pimctl/releases/download/v${VERSION}/pimctl_${VERSION}_linux_${ARCH}.tar.gz" -o "${TEMP_DIR}/pimctl_${VERSION}_linux_${ARCH}.tar.gz"
tar xzvf "${TEMP_DIR}/pimctl_${VERSION}_linux_${ARCH}.tar.gz" -C "${TEMP_DIR}"
sudo mv "${TEMP_DIR}/pimctl" /usr/local/bin/pimctl
```

### MacOS

Not yet tested, but available for download on the [release page](https://github.com/co-native-ab/pimctl/releases). Should be kind of the same way as with Linux.

## FAQ

### How do I use pimctl?

We haven't documented much of it as of right now but will in the future when it's not in alpha. Once you have installed it and set up the Entra Application, first log in using the Azure CLI (not a requirement but makes it easier to discover the client ID and tenant ID) and then run `pimctl login`. After that, you can see the subcommands by running `pimctl`. The two primary commands right now are `pimctl group eligible menu` to request assignment and `pimctl group approval menu` to approve requests from other people.

### Where are the sessions/tokens cached locally?

They are cached in `~/.IdentityService/pimctl_default.auth-record.json`, `~/.IdentityService/pimctl_default.cae`, and `~/.IdentityService/pimctl_default.config-cache.json`.

### Can I use multiple sessions with pimctl at the same time?

Yes, by using the environment variable `PIMCTL_PROFILE`. If you run `PIMCTL_PROFILE=foobar pimctl [subcommand]`, the cache will have the prefix `pimctl_foobar` instead of `pimctl_default`. This should be used together with `AZURE_CONFIG_DIR`, especially if the different sessions are in different Entra tenants (this could be skipped by manually providing `--client-id` and `--tenant-id` when running `pimctl login`).

### Will someone using pimctl get any extra permissions due to the delegated permissions in the Entra Application?

No, delegated permissions are used in the delegated access scenario. They're permissions that allow the application to act on a user's behalf. The application will never be able to access anything the signed-in user themselves couldn't access. See [Overview of permissions and consent in the Microsoft identity platform](https://learn.microsoft.com/entra/identity-platform/permissions-consent-overview#types-of-permissions) for more information.

### Why do I need something other than Azure CLI?

Azure CLI is usually the preferred way of handling authentication to Azure resources when in the terminal. Other applications (like Terraform and SOPS) can use the authenticated session from Azure CLI, piggybacking on it to not require another login. Unfortunately, the Entra Application used by Azure CLI doesn't include the required delegated permissions for working with Entra PIM Groups ([azure-cli #30149](https://github.com/Azure/azure-cli/issues/30149)). Because of this, we need to create our own Entra Application with the required delegated permissions and use it instead of the Azure CLI. In the future, when these permissions are included in Azure CLI, we may not need our own Entra Application anymore—and in the long run, all of the features of pimctl could be a native part of Azure CLI.

## Setting Up Entra Application

You need to create an Entra Application that will be used together with pimctl. The easiest way to configure it is by setting it up using the included Terraform module in `terraform/`. If you're not an administrator in Entra and can't grant admin approval, remove the resource `azuread_service_principal_delegated_permission_grant` and ask an admin to approve it manually.

If you want to set it up manually, the following is important:

- Enable public client
- Add redirect URI `http://localhost:45241`
- Through the metadata, add the tag `pimctl`
- Add the following (Delegated) Microsoft Graph scopes:
  - `Group.Read.All`
  - `PrivilegedAssignmentSchedule.ReadWrite.AzureADGroup`
  - `PrivilegedEligibilitySchedule.Read.AzureADGroup`
  - `RoleManagementPolicy.Read.AzureADGroup`
  - `User.Read`
- Grant admin approval to the scopes

**NOTE**: The tag `pimctl` is used during `pimctl login` to auto-discover the application if the user running the command has access to read the application as well as logged on using the Azure CLI. If the user doesn't have read access or isn't logged on using the Azure CLI, please use `pimctl login --client-id [entra-application-id] --tenant-id [entra-tenant-id]`. After the first login, `--client-id` and `--tenant-id` won't be necessary until `pimctl account clear` is run, since it will try to use the cached authentication record (stored in `~/.IdentityService/pimctl_[profile].auth-record.json`).

## How-To

### Use Multiple Profiles

**Session #1 ("default user"):**

```shell
az login
pimctl login
az account show
pimctl account show
```

**Session #2 ("user2"):**

```shell
export PIMCTL_PROFILE=user2
export AZURE_CONFIG_DIR=$(mktemp -d)
az login
pimctl login
az account show
pimctl account show
```

By changing `PIMCTL_PROFILE`, separate files for the cache will be used (in `$HOME/.IdentityService/`). By changing `AZURE_CONFIG_DIR`, the Azure CLI cache location will be changed (not required but used in pimctl discovery of Entra Application and Tenant ID).

## Alternatives

### Azure PIM CLI by [netr0m](https://github.com/netr0m)

[netr0m/az-pim-cli](https://github.com/netr0m/az-pim-cli) is a project written in Go, but when working with groups, it requires you to extract a token manually from the portal. In pimctl, Microsoft Graph is used for everything, but that also requires us (as of the time of writing) to set up an Entra Application with delegated scopes to handle it.

### Azure PIM CLI by [demoray](https://github.com/demoray)

[demoray/azure-pim-cli](https://github.com/demoray/azure-pim-cli) is a project written in Rust but doesn't support groups.

### Graph EasyPIM by [rakheshster](https://github.com/rakheshster)

[rakheshster/PowerShell-GraphEasyPIM](https://github.com/rakheshster/PowerShell-GraphEasyPIM) is a project written in PowerShell but doesn't support approvals.

## License

This project is licensed under the MIT License—see the [LICENSE](LICENSE) file for details.
