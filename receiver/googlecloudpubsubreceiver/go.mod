module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver

go 1.23.0

require (
	cloud.google.com/go/logging v1.13.0
	cloud.google.com/go/pubsub v1.47.0
	github.com/google/go-cmp v0.6.0
	github.com/googleapis/gax-go/v2 v2.14.1
	github.com/iancoleman/strcase v0.3.0
	github.com/json-iterator/go v1.1.12
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding v0.119.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/component/componenttest v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/confmap v1.25.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/confmap/xconfmap v0.0.0-20250210155359-76f44e1e21d1
	go.opentelemetry.io/collector/consumer v1.25.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/consumer/consumertest v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/exporter v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/pdata v1.25.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/receiver v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/collector/receiver/receivertest v0.119.1-0.20250210123122-44b3eeda354c
	go.opentelemetry.io/otel v1.34.0
	go.opentelemetry.io/otel/metric v1.34.0
	go.opentelemetry.io/otel/sdk/metric v1.34.0
	go.opentelemetry.io/otel/trace v1.34.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	google.golang.org/api v0.220.0
	google.golang.org/genproto v0.0.0-20250122153221-138b5a5a4fd4
	google.golang.org/genproto/googleapis/api v0.0.0-20250124145028-65684f501c47
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.5
)

require (
	cloud.google.com/go v0.118.1 // indirect
	cloud.google.com/go/auth v0.14.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.7 // indirect
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	cloud.google.com/go/iam v1.3.1 // indirect
	cloud.google.com/go/longrunning v0.6.4 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	go.einride.tech/aip v0.68.1 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/config/configretry v1.25.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/extension v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/extension/xextension v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/featuregate v1.25.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/pipeline v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.119.1-0.20250210123122-44b3eeda354c // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.58.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.58.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/oauth2 v0.25.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250127172529-29210b9bc287 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding => ../../extension/encoding
