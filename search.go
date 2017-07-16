package arena

type Search struct {
	client        *Client
	Term          string            `json:"term"`
	Per           int               `json:"per"`
	CurrentPage   int               `json:"current_page"`
	TotalPages    int               `json:"total_pages"`
	Length        int               `json:"length"`
	Authenticated bool              `json:"authenticated"`
	Blocks        []SearchedBlock   `json:"blocks"`
	Channels      []SearchedChannel `json:"channels"`
	Users         []UserStruct      `json:"users"`
}

type SearchedBlock struct {
	client          *Client
	ID              string              `json:"id"`
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

type SearchedChannel struct {
	client            *Client
	ID                string         `json:"id"`
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

// Search allows an authenticated user to search anime titles. Upon failure it
// will return ErrNoContent.
func (c *Client) Search(args Arguments) (search *Search, err error) {
	path := "search"
	err = c.Get(path, args, &search)
	if search != nil {
		search.client = c
	}
	return
}

// Search allows an authenticated user to search anime titles. Upon failure it
// will return ErrNoContent.
func (c *Client) SearchBlocks(args Arguments) (search *Search, err error) {
	path := "search/blocks"
	err = c.Get(path, args, &search)
	if search != nil {
		search.client = c
	}
	return
}

// Search allows an authenticated user to search anime titles. Upon failure it
// will return ErrNoContent.
func (c *Client) SearchChannels(args Arguments) (search *Search, err error) {
	path := "search/channels"
	err = c.Get(path, args, &search)
	if search != nil {
		search.client = c
	}
	return
}
