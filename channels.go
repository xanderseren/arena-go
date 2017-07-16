package arena

import "fmt"

// MyAnimeList API docs: http://myanimelist.net/modules.php?go=api
type ChannelsService struct {
	client *Client
}

type ChannelContents struct {
	client   *Client
	Contents []Block `json:"contents"`
}

type Channel struct {
	client            *Client
	ID                int            `json:"id"`
	Title             string         `json:"title"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         string         `json:"updated_at"`
	AddedToAt         string         `json:"added_to_at"`
	Published         bool           `json:"published"`
	Open              bool           `json:"open"`
	Collaboration     bool           `json:"collaboration"`
	CollaboratorCount int            `json:"collaborator_count"`
	Slug              string         `json:"slug"`
	Length            int            `json:"length"`
	Kind              string         `json:"kind"`
	Status            string         `json:"status"`
	UserID            int            `json:"user_id"`
	Contents          []Block        `json:"contents"`
	BaseClass         string         `json:"base_class"`
	Page              int            `json:"page"`
	Per               int            `json:"per"`
	Collaborators     []UserStruct   `json:"collaborators"`
	FollowerCount     int            `json:"follower_count"`
	ShareLink         *string        `json:"share_link"`
	Metadata          MetadataStruct `json:"metadata"`
	ClassName         string         `json:"class_name"`
	CanIndex          bool           `json:"can_index"`
	User              UserStruct     `json:"user"`
}

// type ContentsStruct struct {
// }

type UserStruct struct {
	ID          int    `json:"id"`
	Slug        string `json:"slug"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	FullName    string `json:"full_name"`
	Avatar      string `json:"avatar"`
	AvatarImage struct {
		Thumb   string `json:"thumb"`
		Display string `json:"display"`
	} `json:"avatar_image"`
	ChannelCount   int            `json:"channel_count"`
	FollowingCount int            `json:"following_count"`
	ProfileId      int            `json:"profile_id"`
	FollowerCount  int            `json:"follower_count"`
	Initials       string         `json:"initials"`
	CanIndex       bool           `json:"can_index"`
	IsPremium      bool           `json:"is_premium"`
	Metadata       MetadataStruct `json:"metadata"`
	BaseClass      string         `json:"base_class"`
	Class          string         `json:"class"`
}

type MetadataStruct struct {
	Description string `json:"description"`
}

// GetChannel returns a channel based on the ID or slug
func (s *ChannelsService) Get(channelID string, args Arguments) (channel *Channel, err error) {
	path := fmt.Sprintf("channels/%s", channelID)
	err = s.client.Get(path, args, &channel)
	if channel != nil {
		channel.client = s.client
	}
	return
}

// GetChannel returns a channel based on the ID or slug
func (s *ChannelsService) Contents(channelID string, args Arguments) (blocks *ChannelContents, err error) {
	path := "channels/" + channelID + "/contents"
	err = s.client.Get(path, args, &blocks)
	if blocks != nil {
		blocks.client = s.client
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