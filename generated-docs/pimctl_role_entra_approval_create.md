## pimctl role entra approval create

Approve or Deny a pending Entra role assignment request

### Synopsis

Approve or Deny a pending Entra Role assignment request

```
pimctl role entra approval create [flags]
```

### Options

```
      --entra-role-name string       The name of the Entra role to grant or deny the assignment to
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

* [pimctl role entra approval](pimctl_role_entra_approval.md)	 - Manage Entra PIM role approvals

