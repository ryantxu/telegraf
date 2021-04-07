package all

import (
	_ "github.com/influxdata/telegraf/plugins/inputs/replay"

	_ "github.com/influxdata/telegraf/plugins/inputs/cpu"
	_ "github.com/influxdata/telegraf/plugins/inputs/disk"
	_ "github.com/influxdata/telegraf/plugins/inputs/diskio"
	_ "github.com/influxdata/telegraf/plugins/inputs/file"
	_ "github.com/influxdata/telegraf/plugins/inputs/filecount"
	_ "github.com/influxdata/telegraf/plugins/inputs/filestat"
	_ "github.com/influxdata/telegraf/plugins/inputs/kernel"
	_ "github.com/influxdata/telegraf/plugins/inputs/mem"
	_ "github.com/influxdata/telegraf/plugins/inputs/net"
	_ "github.com/influxdata/telegraf/plugins/inputs/ping"
	_ "github.com/influxdata/telegraf/plugins/inputs/swap"
	_ "github.com/influxdata/telegraf/plugins/inputs/system"
)
