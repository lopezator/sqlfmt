# Build and release

1. Update `cockroachdb/cockroach-gen` to the latest version:

```bash
go get github.com/cockroachdb/cockroach-gen@main
```

2. Update `cockroachdb/cockroach` to the latest version:

```bash
go get github.com/cockroachdb/cockroach@master
```

3. Tidy up the `go.mod` file:

```bash
go mod tidy
```

4. Update the version in the `Makefile` accordingly.

5. Generate the binaries:

```bash
make build
```

6. Release a new version:

Tag the new version and upload the resulting binaries in the `bin/` folder to github.
