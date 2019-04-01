module github.com/lopezator/sqlfmt

go 1.12

require (
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/cockroachdb/cockroach v2.2.0-alpha.20190211.0.20190220112258-1387b4fad485+incompatible
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.3
)

replace (
	github.com/cockroachdb/apd => github.com/cockroachdb/apd v1.1.0
	github.com/cockroachdb/cockroach => github.com/lopezator/cockroach v0.0.0-20190305125206-ee83b6ab8240
)
