## pimctl role entra request create

Create a new Entra role assignment request

### Synopsis

Create a new Entra role assignment request

```
pimctl role entra request create [flags]
```

### Options

```
      --duration int             The duration of the request in hours (default: configured maximum duration)
      --entra-role-id string     The id of the Entra role to request assignment for (can't be used with --entra-role-name)
      --entra-role-name string   The name of the Entra role to request assignment for (can't be used with --entra-role-id and can't have two roles with the same name)
  -h, --help                     help for create
      --justification string     The justification for the request
      --scope-id string          The scope id to request the role assignment for (default "/")
```

### Options inherited from parent commands

```
  -p, --profile string   The cache profile name for pimctl. Enables running multiple sessions with different users at the same time. [env: PIMCTL_PROFILE] (default "default")
```

### SEE ALSO

* [pimctl role entra request](pimctl_role_entra_request.md)	 - Manage Entra PIM role assignment requests

