package arena

import "fmt"

type Block struct {
	client    *Client
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Slug      string `json:"slug"`
}

// GetBlock returns a channel based on the ID
func (c *Client) GetBlock(blockID string, args Arguments) (block *Block, err error) {
	path := fmt.Sprintf("blocks/%s", blockID)
	err = c.Get(path, args, &block)
	if block != nil {
		block.client = c
	}
	return
}
