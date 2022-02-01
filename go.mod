module github.com/lopezator/sqlfmt

go 1.17

replace (
	github.com/abourget/teamcity => github.com/cockroachdb/teamcity v0.0.0-20180905144921-8ca25c33eb11
	github.com/cockroachdb/cockroach => github.com/cockroachdb/cockroach-gen v0.0.0-20220201140352-135da6111bcf
	vitess.io/vitess => github.com/cockroachdb/vitess v0.0.0-20210218160543-54524729cc82
)

require (
	github.com/cockroachdb/cockroach v0.0.0-20220201122851-c4f15d6c3d6e
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.2.1
)

require (
	github.com/DataDog/zstd v1.5.0 // indirect
	github.com/VividCortex/ewma v1.1.1 // indirect
	github.com/bazelbuild/rules_go v0.26.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/biogo/store v0.0.0-20160505134755-913427a1d5e8 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cockroachdb/apd/v3 v3.0.1 // indirect
	github.com/cockroachdb/errors v1.8.6 // indirect
	github.com/cockroachdb/gostdlib v1.13.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20211118104740-dabe8e521a4f // indirect
	github.com/cockroachdb/pebble v0.0.0-20220126162719-a5c1766b568a // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/cockroachdb/sentry-go v0.6.1-cockroachdb.2 // indirect
	github.com/cockroachdb/ttycolor v0.0.0-20210902133924-c7d7dcdde4e8 // indirect
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/go-kit/log v0.1.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gogo/status v1.1.0 // indirect
	github.com/golang/geo v0.0.0-20200319012246-673a6f80352d // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jaegertracing/jaeger v1.18.1 // indirect
	github.com/klauspost/compress v1.14.1 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/openzipkin/zipkin-go v0.2.5 // indirect
	github.com/petermattis/goid v0.0.0-20211229010228-4d14c490ee36 // indirect
	github.com/pierrre/geohash v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.1-0.20210607210712-147c58e9608a // indirect
	github.com/prometheus/common v0.30.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/prometheus/prometheus v1.8.2-0.20210914090109-37468d88dce8 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/twpayne/go-geom v1.4.1 // indirect
	github.com/twpayne/go-kml v1.5.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	go.etcd.io/etcd/raft/v3 v3.0.0-20210320072418-e51c697ec6e8 // indirect
	go.opentelemetry.io/otel v1.0.0-RC3 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC3 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.0.0-RC3.0.20210907151655-df2bdbbadb26 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.0.0-RC3 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.0.0-RC3 // indirect
	go.opentelemetry.io/otel/sdk v1.0.0-RC3 // indirect
	go.opentelemetry.io/otel/trace v1.0.0-RC3 // indirect
	go.opentelemetry.io/proto/otlp v0.9.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/exp v0.0.0-20220104160115-025e73f80486 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.8-0.20211029000441-d6a9af8af023 // indirect
	google.golang.org/genproto v0.0.0-20211222154725-9823f7ba7562 // indirect
	google.golang.org/grpc v1.43.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
