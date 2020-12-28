# lets-go project

From https://lets-go.alexedwards.net/

## TODO

- [ ] create postgresql role/user with schema privileges.

## Tips

### Project structure

- https://peter.bourgon.org/go-best-practices-2016/#repository-structure
- https://github.com/thockin/go-build-template

## Gotchas

### git: show older file version

`git show a23ffe79ca:main.go`

### html linter

[tidy](https://www.html-tidy.org/)

`http POST :8080/ | tidy -errors -quiet`

### interval in prepared pg statement:

Use `$1::interval`. See [superior query](https://dba.stackexchange.com/questions/208580/passing-value-of-datatype-interval-in-parametrized-query)

Or use [How I use Postgres with Go](https://jbrandhorst.com/post/postgres/)
