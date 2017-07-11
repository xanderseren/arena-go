package arena

import "fmt"

type Board struct {
	client *Client
	ID     string `json:"id"`
	Title  string `json:"title"`
}

/**
 * Channel retrieves a channel by its ID.
 */
func (c *Client) GetChannel(channelID string, args Arguments) (channel *Channel, err error) {
	path := fmt.Sprintf("channels/%s", channelID)
	err = c.Get(path, args, &channel)
	if channel != nil {
		channel.client = c
	}
	return
}
