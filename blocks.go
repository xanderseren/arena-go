package arena

import (
	"fmt"
	"net/url"
)

// BlocksService. MyAnimeList API docs: http://myanimelist.net/modules.php?go=api
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

// Get returns returns the full representation of a block based on the ID or Slug
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

// AddContent creates a new block and adds it to the specified channel.
// Entry is textual content that's rendered with Github Flavored Markdown.
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

// AddSource creates a new block and adds it to the specified channel.
// Entry is URL of content.
func (s *BlocksService) AddSource(channelID string, entry string) (block *Block, err error) {

	data := url.Values{
		"source": {entry},
	}
	path := "channels/" + channelID + "/blocks"
	err = s.client.Post(path, nil, data, &block)
	if block != nil {
		block.client = s.client
	}
	return
}

// EditTitle updates a block's attributes. title is the new title of the block
func (s *BlocksService) EditTitle(blockID string, title string) error {

	data := url.Values{
		"title": {title},
	}
	path := "blocks/" + blockID
	return s.client.Put(path, nil, data)
}

// EditDescription updates a block's attributes. description is the new
// description of the block. Markdown formatted text.
func (s *BlocksService) EditDescription(blockID string, description string) error {

	data := url.Values{
		"description": {description},
	}
	path := "blocks/" + blockID
	return s.client.Put(path, nil, data)
}

// EditDescription updates a block's attributes. content is the new
// content of the block. Markdown formatted text. For text block type only.
func (s *BlocksService) EditContent(blockID string, entry string) error {

	data := url.Values{
		"content": {entry},
	}
	path := "blocks/" + blockID
	return s.client.Put(path, nil, data)
}

// ListChannels earch returns a paginated list of channels the block exists in
// based on the search query q
func (s *BlocksService) ListChannels(blockID string, args Arguments) (searchStruct *SearchStruct, err error) {
	path := "blocks/" + blockID + "/channels"
	err = s.client.Get(path, args, &searchStruct)
	if searchStruct != nil {
		searchStruct.client = s.client
	}
	return
}
