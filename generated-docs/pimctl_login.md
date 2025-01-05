## pimctl login

Login to pimctl

### Synopsis

Login to pimctl. This command will authenticate you to Azure PIM

```
pimctl login [flags]
```

### Options

```
      --client-id string                     Entra Application Client ID for pimctl
      --credential-method CredentialMethod   Credential method to use for pimctl. Allowed: 'DeviceCode', 'InteractiveBrowser' (default InteractiveBrowser)
  -h, --help                                 help for login
      --tenant-id string                     Entra Tenant ID to use for login (default "organizations")
```

### Options inherited from parent commands

```
  -p, --profile string   The cache profile name for pimctl. Enables running multiple sessions with different users at the same time. [env: PIMCTL_PROFILE] (default "default")
```

### SEE ALSO

* [pimctl](pimctl.md)	 - CLI to manage Azure PIM roles and assignments

