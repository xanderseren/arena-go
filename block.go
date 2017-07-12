package arena

import (
	"encoding/json"
	"net/http"
)

type Block struct {
	client    *Client
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Slug      string `json:"slug"`
}

// doBlockRequest handles GET request for Block
func (c *Client) doBlockRequest(blockID string) (Block, error) {
	var block Block

	path := c.BaseURL + "blocks/" + blockID
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return block, err
	}

	body, err := c.do(req)
	if err != nil {
		return block, err
	}
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&block); err != nil {
		return block, err
	}
	return block, nil
}

// Block returns the block via the ID or Slug
func (c *Client) GetBlock(blockID string) (Block, error) {
	return c.doBlockRequest(blockID)
}

// // GetBlock retrieves a block by its ID or Slug.
// func (c *Client) GetBlock(blockID string) (block *Block, err error) {
// 	path := fmt.Sprintf("blocks/%s", blockID)
// 	err = c.Get(path, &block)
// 	if block != nil {
// 		block.client = c
// 	}
// 	return
// }
