module github.com/influxdata/telegraf

go 1.16

require (
	collectd.org v0.5.0
	github.com/BurntSushi/toml v0.3.1
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d
	github.com/antchfx/xmlquery v1.3.5
	github.com/antchfx/xpath v1.1.11
	github.com/aws/aws-sdk-go v1.34.34
	github.com/benbjohnson/clock v1.0.3
	github.com/bmatcuk/doublestar/v3 v3.0.0
	github.com/caio/go-tdigest v3.1.0+incompatible
	github.com/dimchansky/utfbom v1.1.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-logfmt/logfmt v0.5.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-ping/ping v0.0.0-20210201095549-52eed920f98c
	github.com/gobwas/glob v0.2.3
	github.com/gogo/protobuf v1.3.1
	github.com/golang/geo v0.0.0-20190916061304-5b978397cfec
	github.com/golang/snappy v0.0.1
	github.com/google/go-cmp v0.5.5
	github.com/gosnmp/gosnmp v1.30.0
	github.com/grafana/grafana-plugin-sdk-go v0.91.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/influxdata/tail v1.0.1-0.20200707181643-03a791b270e4
	github.com/influxdata/toml v0.0.0-20190415235208-270119a8ce65
	github.com/influxdata/wlog v0.0.0-20160411224016-7c63b0a71ef8
	github.com/kardianos/service v1.0.0
	github.com/karrick/godirwalk v1.16.1
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/kr/text v0.2.0 // indirect
	github.com/leesper/go_rng v0.0.0-20190531154944-a612b043e353 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/nsqio/go-nsq v1.0.8
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.20.0
	github.com/prometheus/prometheus v1.8.2-0.20200911110723-e83ef207b6c2
	github.com/shirou/gopsutil v3.20.11+incompatible
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.6.0
	github.com/vjeantet/grok v1.0.1
	go.starlark.net v0.0.0-20210312235212-74c10e2c17dc
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2
	golang.org/x/text v0.3.4
	gonum.org/v1/gonum v0.7.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/djherbis/times.v1 v1.2.0
)

// replaced due to https://github.com/satori/go.uuid/issues/73
replace github.com/satori/go.uuid => github.com/gofrs/uuid v3.2.0+incompatible
