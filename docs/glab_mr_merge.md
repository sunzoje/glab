## glab mr merge

Merge/Accept merge requests

### Synopsis

Merge/Accept merge requests

```
glab mr merge <id> [flags]
```

### Options

```
  -m, --message string           Get only closed merge requests
  -d, --remove-source-branch     Remove source branch on merge
      --sha string               Merge Commit sha
  -s, --squash                   Squash commits on merge
      --squash-message string    Squash commit message
      --when-pipeline-succeeds   Merge only when pipeline succeeds. Default to true (default true)
```

### Options inherited from parent commands

```
      --help          Show help for command
  -R, --repo string   Select another repository using the OWNER/REPO format or the project ID. Supports group namespaces
```

### SEE ALSO

* [glab mr](glab_mr.md)	 - Create, view and manage merge requests

###### Auto generated by spf13/cobra on 29-Sep-2020
