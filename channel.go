package arena

import "fmt"

type Channel struct {
	client *Client
	ID     string `json:"id"`
	Title  string `json:"title"`
}

/**
 * Channel retrieves a channel by its ID.
 */
func (c *Client) GetChannel(channelID string) (channel *Channel, err error) {
	path := fmt.Sprintf("channels/%s", channelID)
	err = c.Get(path, &channel)
	if channel != nil {
		channel.client = c
	}
	return
}
