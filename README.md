# SQLNull
![go-version-badge](https://img.shields.io/github/go-mod/go-version/gkits/sqlnull)
![test-pipeline-badge](https://github.com/gkits/sqlnull/actions/workflows/test.yml/badge.svg)
![test-coverage-badge](https://raw.githubusercontent.com/gKits/sqlnull/badges/.badges/main/coverage.svg)

A dead simple one stop solution to pass your database types directly to
your JSON REST-API.

---

This package provides wrapper types for all Null* types of the 
[`database/sql`](https://pkg.go.dev/database/sql) which implement the 
`Marshaler` and `Unmarshaler` interface of the 
[`encoding/json`](https://pkg.go.dev/encoding/json) package.

## Usage

1. Get the package
```bash
go get github.com/gkits/nullsql@latest
```

2. Import the package
```go
import "github.com/gkits/nullsql"
```

3. Start coding...

## Limitations

Due to the internal implementation of the 
[`encoding/json`](https://pkg.go.dev/encoding/json) package you can sadly
not omit the fields with value `null` using the `omitempty` tag.
