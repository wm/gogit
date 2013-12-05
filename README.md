WIP - result will output open pulls something like:

```
Open Pulls

| Pull | Comments | Passing | :octocatted: |
| 2070 |        0 |         |        false |
| 2072 |        0 |         |        false |
| 2071 |        0 |         |        false |
| 2073 |        0 |         |        false |
| 2011 |        3 |         |         true |
| 2058 |        3 |         |         true |
```

CLI App

You need to set some ENV variables


```
$ go build cmd/gogit

$ export GOGIT_GH_TOKEN=<your github API token>
$ export GOGIT_GH_OWNER='wm'
$ export GOGIT_GH_REPO='gogit'

$ gogit
```
