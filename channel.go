package arena

import "fmt"

type Channel struct {
	client            *Client
	ID                int                   `json:"id"`
	Title             string                `json:"title"`
	CreatedAt         string                `json:"created_at"`
	UpdatedAt         string                `json:"updated_at"`
	AddedToAt         string                `json:"added_to_at"`
	Published         bool                  `json:"published"`
	Open              bool                  `json:"open"`
	Collaboration     bool                  `json:"collaboration"`
	CollaboratorCount int                   `json:"collaborator_count"`
	Slug              string                `json:"slug"`
	Length            int                   `json:"length"`
	Kind              string                `json:"kind"`
	Status            string                `json:"status"`
	UserID            int                   `json:"user_id"`
	Contents          []ContentsStruct      `json:"contents"`
	BaseClass         string                `json:"base_class"`
	Page              int                   `json:"page"`
	Per               int                   `json:"per"`
	Collaborators     []CollaboratorsStruct `json:"collaborators"`
	FollowerCount     int                   `json:"follower_count"`
	ShareLink         *string               `json:"share_link"`
	Metadata          MetadataStruct        `json:"metadata"`
	ClassName         string                `json:"class_name"`
	CanIndex          bool                  `json:"can_index"`
	User              UserStruct            `json:"user"`
}

type ContentsStruct struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Slug      string `json:"slug"`
}

type CollaboratorsStruct struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Avatar    string `json:"avatar"`
}

type UserStruct struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Avatar    string `json:"avatar"`
}

type MetadataStruct struct {
	Description string `json:"description"`
}

// GetChannel returns a channel based on the ID or slug
func (c *Client) GetChannel(channelID string, args Arguments) (channel *Channel, err error) {
	path := fmt.Sprintf("channels/%s", channelID)
	err = c.Get(path, args, &channel)
	if channel != nil {
		channel.client = c
	}
	return
}

// GetChannelContents returns a channels contents based on the ID r slug
// func (c *Client) GetChannelContents(channelID string, args Arguments) (channel *Channel, err error) {
// 	path := fmt.Sprintf("channels/%s", channelID, "/%scontents")
// 	err = c.Get(path, args, &channel)
// 	if channel != nil {
// 		channel.client = c
// 	}
// 	return
// }
