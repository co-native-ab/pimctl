## pimctl group request create

Create a new group assignment request

### Synopsis

Create a new group assignment request

```
pimctl group request create [flags]
```

### Options

```
      --duration int           The duration of the request in hours (default: configured maximum duration)
      --group-id string        The id of the group to request assignment for (can't be used with --group-name)
      --group-name string      The name of the group to request assignment for (can't be used with --group-id and can't have two groups with the same name)
  -h, --help                   help for create
      --justification string   The justification for the request
```

### Options inherited from parent commands

```
  -p, --profile string   The cache profile name for pimctl. Enables running multiple sessions with different users at the same time. [env: PIMCTL_PROFILE] (default "default")
```

### SEE ALSO

* [pimctl group request](pimctl_group_request.md)	 - Manage Azure PIM group assignment requests

