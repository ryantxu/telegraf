package grafanalive

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/outputs"
	"github.com/influxdata/telegraf/plugins/serializers"
)

// GrafanaLive connects to grafana server
type GrafanaLive struct {
	URL    string          `toml:"url"`
	Stream string          `toml:"stream"`
	Log    telegraf.Logger `toml:"-"`

	client     *Client
	serializer serializers.Serializer
}

var sampleConfig = `
[[outputs.grafana]]
  # The address of the local grafana instance
  url = "http://localhost:3000"
  # The channel to write data into grafana with
  stream = "telegraf"
`

func (g *GrafanaLive) SetSerializer(serializer serializers.Serializer) {
	g.serializer = serializer
}

func (g *GrafanaLive) Connect() error {
	var err error

	g.Log.Infof("Connecting to grafana live: %s", g.URL)
	g.client, err = NewClient(g.URL)
	if err != nil {
		return err
	}
	g.Log.Info("Connected... waiting for data")
	return err
}

func (g *GrafanaLive) Close() error {
	return nil
}

func (g *GrafanaLive) SampleConfig() string {
	return sampleConfig
}

func (g *GrafanaLive) Description() string {
	return "Send telegraf metrics to a grafana live stream"
}

func (g *GrafanaLive) Write(metrics []telegraf.Metric) error {
	reqBody, err := g.serializer.SerializeBatch(metrics)
	if err != nil {
		return err
	}
	return g.client.Publish("push/"+g.Stream, reqBody)
}

func init() {
	outputs.Add("grafana-live", func() telegraf.Output {
		return &GrafanaLive{}
	})
}
