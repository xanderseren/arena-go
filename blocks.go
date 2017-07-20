package arena

import (
	"fmt"
	"net/url"
)

// MyAnimeList API docs: http://myanimelist.net/modules.php?go=api
type BlocksService struct {
	client *Client
}

type BlockEntry struct {
	client  *Client
	Content string
}

type Block struct {
	client          *Client
	ID              int              `json:"id"`
	Title           string           `json:"title"`
	UpdatedAt       string           `json:"updated_at"`
	CreatedAt       string           `json:"created_at"`
	State           string           `json:"state"`
	CommentCount    int              `json:"comment_count"`
	GeneratedTitle  string           `json:"generated_title"`
	ContentHtml     *string          `json:"content_html"`
	DescriptionHtml *string          `json:"description_html"`
	Visibility      string           `json:"visibility"`
	Content         *string          `json:"content"`
	Description     *string          `json:"description"`
	Source          SourceStruct     `json:"source"`
	Image           ImageStruct      `json:"image"`
	Embed           EmbedStruct      `json:"embed"`
	Attachment      AttachmentStruct `json:"attachment"`
	Metadata        MetadataStruct   `json:"metadata"`
	BaseClass       string           `json:"base_class"`
	Class           string           `json:"class"`
	User            UserStruct       `json:"user"`
}

type SourceStruct *struct {
	Url      string  `json:"url"`
	Title    *string `json:"title"`
	Provider struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"provider"`
}

type ImageStruct *struct {
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	UpdatedAt   string `json:"updated_at"`
	Thumb       struct {
		Url string `json:"url"`
	} `json:"thumb"`
	Square struct {
		Url string `json:"url"`
	} `json:"square"`
	Display struct {
		Url string `json:"url"`
	} `json:"display"`
	Large struct {
		Url string `json:"url"`
	} `json:"large"`
	Original struct {
		Url             string `json:"url"`
		FileSize        int    `json:"file_size"`
		FileSizeDisplay string `json:"file_size:display"`
	} `json:"original"`
}

type EmbedStruct *struct {
	Url           *string `json:"url"`
	Type          *string `json:"type"`
	Title         *string `json:"title"`
	AuthorName    *string `json:"author_name"`
	AuthorUrl     *string `json:"author_url"`
	SourceUrl     *string `json:"source_url"`
	Thumbnail_url *string `json:"thumbnail_url"`
	Width         int     `json:"width"`
	Height        int     `json:"height"`
	Html          string  `json:"html"`
}

type AttachmentStruct *struct {
	FileName        string `json:"file_name"`
	FileSize        int    `json:"file_size"`
	FileSizeDisplay string `json:"file_size_display"`
	ContentType     string `json:"content_type"`
	Extension       string `json:"extension"`
	Url             string `json:"url"`
}

// Get returns a block based on the ID or Slug
func (s *BlocksService) Get(blockID string, args Arguments) (block *Block, err error) {
	path := fmt.Sprintf("blocks/%s", blockID)
	err = s.client.Get(path, args, &block)
	if block != nil {
		block.client = s.client
	}
	return
}

// Search returns a selection of blocks based on the search query q
func (s *BlocksService) Search(args Arguments) (searchStruct *SearchStruct, err error) {
	path := "search/blocks"
	err = s.client.Get(path, args, &searchStruct)
	if searchStruct != nil {
		searchStruct.client = s.client
	}
	return
}

func (s *BlocksService) AddContent(channelID string, entry string) (block *Block, err error) {

	data := url.Values{
		"content": {entry},
	}
	fmt.Println(data)
	path := "channels/" + channelID + "/blocks"
	err = s.client.Post(path, nil, data, &block)
	if block != nil {
		block.client = s.client
	}
	return
}

func (s *BlocksService) AddSource(channelID string, entry string) (block *Block, err error) {

	data := url.Values{
		"source": {entry},
	}
	fmt.Println(data)
	path := "channels/" + channelID + "/blocks"
	err = s.client.Post(path, nil, data, &block)
	if block != nil {
		block.client = s.client
	}
	return
}

func (s *BlocksService) EditTitle(blockID string, entry string) (block *Block, err error) {

	data := url.Values{
		"title": {entry},
	}
	path := "blocks/" + blockID
	err = s.client.Put(path, nil, data)
	if block != nil {
		block.client = s.client
	}
	return
}

// ListChannels earch returns a selection of blocks based on the search query q
func (s *BlocksService) ListChannels(blockID string, args Arguments) (searchStruct *SearchStruct, err error) {
	path := "blocks/" + blockID + "/channels"
	err = s.client.Get(path, args, &searchStruct)
	if searchStruct != nil {
		searchStruct.client = s.client
	}
	return
}

//
// // EditBlock returns a channel based on the ID
// func (c *Client) EditBlock(blockID string, args Arguments) (block *Block, err error) {
// 	path := fmt.Sprintf("blocks/%s", blockID)
// 	err = c.Put(path, args, &block)
// 	if block != nil {
// 		block.client = c
// 	}
// 	return
// }
