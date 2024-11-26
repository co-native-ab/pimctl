## pimctl group approval create

Approve or Deny a pending group assignment request

### Synopsis

Approve or Deny a pending group assignment request

```
pimctl group approval create [flags]
```

### Options

```
      --group-name string            The name of the group to grant or deny the assignment to
  -h, --help                         help for create
      --justification string         The justification for the approval or denial
      --review-result ReviewResult   Review result to use for pending assignment approval. Allowed: 'Approve', 'Deny' (default Approve)
      --user-principal-name string   The user principal name of the user to approve or deny the request for
```

### Options inherited from parent commands

```
  -p, --profile string   The cache profile name for pimctl. Enables running multiple sessions with different users at the same time. [env: PIMCTL_PROFILE] (default "default")
```

### SEE ALSO

* [pimctl group approval](pimctl_group_approval.md)	 - Manage Azure PIM group approvals

