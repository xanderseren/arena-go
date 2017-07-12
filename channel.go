package arena

import (
	"encoding/json"
	"net/http"
)

type Channel struct {
	client    *Client
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	AddedToAt string `json:"added_to_at"`
	Slug      string `json:"slug"`
}

// doChannelRequest handles GET request for Channel
func (c *Client) doChannelRequest(channelID string) (Channel, error) {
	var channel Channel

	path := c.BaseURL + "channels/" + channelID
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return channel, err
	}

	body, err := c.do(req)
	if err != nil {
		return channel, err
	}
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&channel); err != nil {
		return channel, err
	}
	return channel, nil
}

// Channel returns the channel via the ID or Slug
func (c *Client) GetChannel(channelID string) (Channel, error) {
	return c.doChannelRequest(channelID)
}

// // GetChannel retrieves a channel by its ID or Slug.
// func (c *Client) GetChannel(channelID string) (channel *Channel, err error) {
// 	path := fmt.Sprintf("channels/%s", channelID)
// 	err = c.Get(path, &channel)
// 	if channel != nil {
// 		channel.client = c
// 	}
// 	return
// }
