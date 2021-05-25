module github.com/lopezator/sqlfmt

go 1.16

replace (
	github.com/abourget/teamcity => github.com/cockroachdb/teamcity v0.0.0-20180905144921-8ca25c33eb11
	github.com/cockroachdb/cockroach v0.0.0-20210524102351-3eeb35f3ea83 => github.com/cockroachdb/cockroach-gen v0.0.0-20210525085710-84c172399a67
	vitess.io/vitess => github.com/cockroachdb/vitess v0.0.0-20210218160543-54524729cc82
)

require (
	github.com/cockroachdb/cockroach v0.0.0-20210524102351-3eeb35f3ea83
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
)
