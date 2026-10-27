# Changelog

All notable changes to this project are documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
No version has been tagged yet; the first tagged release will be v1.0.0.

## [Unreleased]

### Added

- `go.mod`: module `github.com/go-ozzo/ozzo-log`, `go 1.21`. The import path is unchanged. Go 1.21 or newer is
  supported; in a module whose `go` line is older, `go get` or `go mod tidy` raises it to `go 1.21`. The module has
  no dependencies besides the standard library.
- `configtest/` — a separate module holding the ozzo-config integration test, so the library itself has no
  third-party dependencies, not even for its tests. It is never tagged and not meant to be imported.
- CI on GitHub Actions for both modules: tests with the race detector on Linux (oldstable, stable), Windows and
  macOS (stable) and on the minimum Go from `go.mod`, a linux/386 build, golangci-lint v2.15.0, gofmt and
  `go mod tidy` checks; govulncheck and coverage on Codecov for the library module. Travis CI is removed.
- golangci-lint configuration, SECURITY.md, CONTRIBUTING.md, CODEOWNERS, Dependabot for GitHub Actions,
  `.gitattributes` keeping LF line endings.

### Changed

- No change in library behaviour. The code changes are findings of the new linter configuration: explicit `_ =` for
  deliberately ignored write and close errors, `make(chan bool)` for `make(chan bool, 0)`, `WriteString` for
  `Write([]byte(s))`. The few findings that would change behaviour (error texts, log file mode, errors of log
  rotation) are marked with `//nolint` and a reason, to be decided before v1.0.0.
- Tests: `TestNetworkTarget` listens on a free port and closes the listener, so the tests pass with `-count` above 1.
  `TestFileTarget` writes its log file to a temporary directory instead of the working directory. `file_test.go`
  uses `os` instead of `io/ioutil`. `TestLoggerConfig` moved to `configtest/`.
- README badges: build status, coverage and API docs now point to GitHub Actions, Codecov and pkg.go.dev (were
  Travis CI, Coveralls/gocover.io and godoc.org).
