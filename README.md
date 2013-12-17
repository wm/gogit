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
$ cd cmd ; go install cmd/gogit
```

```
$ export GOGIT_GH_TOKEN=<your github API token>
$ export GOGIT_OWNER=<the organization or owner of repositories>

$ gogit
```

or

```
$ gogit -token 7c5f06367ffea77071c84e32f02a505304248097 github API token> -owner wm
```
