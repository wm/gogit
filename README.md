WIP - result will output open pulls something like:

```
Open Pulls

| Snowflake | Comments | Passing | :octocatted: | created_by | last_comment_by | title |
| #199      | 3        | ✓       | ✓            |
| #201      | 1        | x       | x            |
| #198      | 0        | ✓       | x            |
```

CLI App

```
$ go build cmd/gogit
$ gogit
```
