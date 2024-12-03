module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter

go 1.22.0

require (
	cloud.google.com/go/pubsub v1.45.1
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/component/componenttest v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/config/configretry v1.20.1-0.20241202231142-b9ff1bc54c99
	go.opentelemetry.io/collector/confmap v1.20.1-0.20241202231142-b9ff1bc54c99
	go.opentelemetry.io/collector/consumer v0.114.1-0.20241202231142-b9ff1bc54c99
	go.opentelemetry.io/collector/exporter v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/exporter/exportertest v0.114.1-0.20241130171227-c52d625647be
	go.opentelemetry.io/collector/pdata v1.20.1-0.20241202231142-b9ff1bc54c99
	go.uber.org/zap v1.27.0
	google.golang.org/api v0.205.0
	google.golang.org/grpc v1.67.1
)

require (
	cloud.google.com/go v0.116.0 // indirect
	cloud.google.com/go/auth v0.10.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.5 // indirect
	cloud.google.com/go/compute/metadata v0.5.2 // indirect
	cloud.google.com/go/iam v1.2.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.13.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	go.einride.tech/aip v0.68.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/exporter/exporterprofiles v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/extension v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/pipeline v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/receiver v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/receiver/receiverprofiles v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.114.1-0.20241130171227-c52d625647be // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/oauth2 v0.23.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20241021214115-324edc3d5d38 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241021214115-324edc3d5d38 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
