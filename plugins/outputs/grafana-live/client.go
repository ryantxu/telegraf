package grafanalive

import (
	"log"
	"net/url"
	"path"
	"time"

	"github.com/centrifugal/centrifuge-go"
)

// Client connects to the GrafanaLive server
type Client struct {
	connected bool
	client    *centrifuge.Client
	lastWarn  time.Time
	Log       log.Logger
}

// Publish to channel.
func (c *Client) Publish(channel string, data []byte) error {
	_, err := c.client.Publish(channel, data)
	return err
}

type liveClientHandler struct {
	client *Client
}

func (h *liveClientHandler) OnConnect(c *centrifuge.Client, e centrifuge.ConnectEvent) {
	h.client.connected = true
}

func (h *liveClientHandler) OnError(c *centrifuge.Client, e centrifuge.ErrorEvent) {
}

func (h *liveClientHandler) OnDisconnect(c *centrifuge.Client, e centrifuge.DisconnectEvent) {
	h.client.connected = false
}

// NewClient creates new Client.
func NewClient(connectionURL string) (*Client, error) {
	u, err := url.Parse(connectionURL)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "https" {
		u.Scheme = "wss"
	} else {
		u.Scheme = "ws"
	}
	u.Path = path.Join(u.Path, "live", "ws")
	values := u.Query()
	values.Set("format", "protobuf")
	u.RawQuery = values.Encode()

	c := centrifuge.New(u.String(), centrifuge.DefaultConfig())

	glc := &Client{
		client: c,
	}
	handler := &liveClientHandler{
		client: glc,
	}
	c.OnConnect(handler)
	c.OnError(handler)
	c.OnDisconnect(handler)

	err = c.Connect()
	if err != nil {
		return nil, err
	}
	return glc, nil
}
