package arena

import (
	"fmt"
	"net/url"
	"sort"
)

// ChannelsService. MyAnimeList API docs: http://myanimelist.net/modules.php?go=api
type ChannelsService struct {
	client *Client
}

type ChannelContents struct {
	client   *Client
	Contents []Block `json:"contents"`
}

type Connections struct {
	client       *Client
	Length       int       `json:"length"`
	TotalPages   int       `json:"total_pages"`
	CurrentPage  int       `json:"current_page"`
	Per          int       `json:"per"`
	ChannelTitle string    `json:"channel_title"`
	BaseClass    string    `json:"base_class"`
	Channels     []Channel `json:"channels"`
}

type Collaborators struct {
	client       *Client
	Length       int          `json:"length"`
	TotalPages   int          `json:"total_pages"`
	CurrentPage  int          `json:"current_page"`
	Per          int          `json:"per"`
	ChannelTitle string       `json:"channel_title"`
	Users        []UserStruct `json:"users"`
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

type MetadataStruct struct {
	Description string `json:"description"`
}

// Get Returns a complete representation of a channel based on the ID or
// slug. Channel contents can be paginated.
func (s *ChannelsService) Get(channelID string, args Arguments) (channel *Channel, err error) {
	path := fmt.Sprintf("channels/%s", channelID)
	err = s.client.Get(path, args, &channel)
	if channel != nil {
		channel.client = s.client
	}

	sort.Slice(channel.Contents[:], func(i, j int) bool {
		return channel.Contents[i].Position > channel.Contents[j].Position
	})

	return
}

// Contents returns all the contents of a channel. Nothing else
// (collaborators, etc.) based on the ID or slug
func (s *ChannelsService) Contents(channelID string, args Arguments) (blocks *ChannelContents, err error) {
	path := "channels/" + channelID + "/contents"
	err = s.client.Get(path, args, &blocks)
	if blocks != nil {
		blocks.client = s.client
	}
	return
}

// Connections returns all of the channels connected to blocks in the channel
// between channels on the ID or slug. For instance if one wanted to spider
// the connections.
func (s *ChannelsService) Connections(channelID string, args Arguments) (connections *Connections, err error) {
	path := "channels/" + channelID + "/connections"
	err = s.client.Get(path, args, &connections)
	if connections != nil {
		connections.client = s.client
	}
	return
}

// Collaborators returns all members that are part of a channel based on
// the ID or slug. Does not return channel owner.
func (s *ChannelsService) Collaborators(channelID string, args Arguments) (collaborators *Collaborators, err error) {
	path := "channels/" + channelID + "/collaborators"
	err = s.client.Get(path, args, &collaborators)
	if collaborators != nil {
		collaborators.client = s.client
	}
	return
}

// Add creates a new channel. Title is title of the channel. Status sets the
// visibility of the channel. Can be: ["public", "closed", "private"]
func (s *ChannelsService) Add(title string, status string) (channel *Channel, err error) {
	data := url.Values{
		"title":  {title},
		"status": {status},
	}
	path := "channels/"
	err = s.client.Post(path, nil, data, &channel)
	if channel != nil {
		channel.client = s.client
	}
	return
}

// Connect creates a new connection. connectChannelID is the ID of the channel
// to connect, recipientChannelID is the ID of the channel to connect to
func (s *ChannelsService) Connect(connectChannelID string, recipientChannelID string) (channel *Channel, err error) {
	data := url.Values{
		"connectable_id":   {connectChannelID},
		"connectable_type": {"Channel"},
	}
	path := "channels/" + recipientChannelID + "/connections"
	fmt.Println(path)
	err = s.client.Post(path, nil, data, &channel)
	if channel != nil {
		channel.client = s.client
	}
	return
}

// Search returns a selection of channels based on the search query q
func (s *ChannelsService) Search(args Arguments) (searchStruct *SearchStruct, err error) {
	path := "search/channels"
	err = s.client.Get(path, args, &searchStruct)
	if searchStruct != nil {
		searchStruct.client = s.client
	}
	return
}
