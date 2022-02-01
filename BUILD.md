# Build and release

Until the path issues are fixed in CockroachDB this manual steps are required in order to make a successful build:

1. Clone/update: `https://github.com/cockroachdb/cockroach-gen` into `$(GOPATH)/src/github.com/cockroachdb/cockroach-gen`.
2. Modify the go.mod replace with a local one (without the `-gen` suffix): `github.com/cockroachdb/cockroach => ../../cockroachdb/cockroach`.
4. Add a new volume so the docker build images are aware of the local repo/replace directive (without `-gen` suffix): `-v $(GOPATH)/src/github.com/cockroachdb/cockroach-gen:/go/src/github.com/cockroachdb/cockroach \`.
5. `make build`