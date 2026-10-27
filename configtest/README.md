# configtest

Integration test of ozzo-log with [ozzo-config](https://github.com/go-ozzo/ozzo-config): it configures a `Logger`
and its targets from JSON, the way the main README describes.

It is a separate Go module so that `github.com/go-ozzo/ozzo-log` itself depends on the standard library only:
projects that import ozzo-log download nothing for this test. The module uses the ozzo-log code of this
repository (`replace github.com/go-ozzo/ozzo-log => ../`). It is never tagged and is not meant to be imported.

```bash
cd configtest
go test -race ./...
```
