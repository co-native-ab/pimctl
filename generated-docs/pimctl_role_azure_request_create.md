## pimctl role azure request create

Create a new Azure role assignment request

### Synopsis

Create a new Azure role assignment request

```
pimctl role azure request create [flags]
```

### Options

```
      --azure-role-definition-id string   The id of the Azure role to request assignment for (can't be used with --azure-role-name)
      --azure-role-name string            The name of the Azure role to request assignment for (can't be used with --azure-role-definition-id and can't have two roles with the same name)
      --duration int                      The duration of the request in hours (default: configured maximum duration)
  -h, --help                              help for create
      --justification string              The justification for the request
      --scope string                      The scope to request assignment for
```

### Options inherited from parent commands

```
  -p, --profile string   The cache profile name for pimctl. Enables running multiple sessions with different users at the same time. [env: PIMCTL_PROFILE] (default "default")
```

### SEE ALSO

* [pimctl role azure request](pimctl_role_azure_request.md)	 - Manage Azure PIM role assignment requests

