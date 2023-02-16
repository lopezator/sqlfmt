# Build and release

1. `go get ...` returns an error because the `go.mod` file has a replace directive. To get around this, edit the `go.mod` file.

```go
github.com/cockroachdb/cockroach => github.com/cockroachdb/cockroach-gen <commitSHA>

github.com/cockroachdb/cockroach <commitSHA>
```

Note: Make sure that the Go version is the same as the one used in cockroachdb and the base docker image to make the build.

2. Tidy up the `go.mod` file:

```bash
go mod tidy
```

3. Update the version in the `Makefile` accordingly.

4. Generate the binaries:

```bash
make build
```

5. Release a new version:

Tag the new version and upload the resulting binaries in the `bin/` folder to github.
