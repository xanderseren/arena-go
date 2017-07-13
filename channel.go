package arena

import "fmt"

type Channel struct {
	client            *Client
	ID                int    `json:"id"`
	Title             string `json:"title"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	AddedToAt         string `json:"added_to_at"`
	Published         bool   `json:"published"`
	Open              bool   `json:"open"`
	Collaboration     bool   `json:"collaboration"`
	CollaboratorCount int    `json:"collaborator_count"`
	Slug              string `json:"slug"`
	Length            int    `json:"length"`
	Kind              string `json:"kind"`
	Status            string `json:"status"`
	UserID            int    `json:"user_id"`
	BaseClass         string `json:"base_class"`
	Page              int    `json:"page"`
	Per               int    `json:"per"`
	FollowerCount     int    `json:"follower_count"`
}

// GetChannel returns a channel based on the ID
func (c *Client) GetChannel(channelID string, args Arguments) (channel *Channel, err error) {
	path := fmt.Sprintf("channels/%s", channelID)
	err = c.Get(path, args, &channel)
	if channel != nil {
		channel.client = c
	}
	return
}
