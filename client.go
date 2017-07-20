package arena

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Client represents the HTTP client
// and any settings used to make requests
// to the arena API.
type Client struct {
	client  *http.Client
	Logger  *logrus.Logger
	BaseURL string
	Token   string

	Search   *SearchService
	Blocks   *BlocksService
	Channels *ChannelsService
	Users    *UsersService
}

// NewClient returns a Client configured with sane default
// values.
func NewClient(token string) *Client {
	logger := logrus.New()
	logger.Level = logrus.WarnLevel

	c := &Client{
		client:  http.DefaultClient,
		BaseURL: "https://api.are.na/v2",
		Logger:  logger,
		Token:   token,
	}

	c.Channels = &ChannelsService{
		client: c,
	}

	c.Blocks = &BlocksService{
		client: c,
	}

	c.Users = &UsersService{
		client: c,
	}

	c.Search = &SearchService{
		client: c,
	}

	return c
}

func (c *Client) Get(path string, args Arguments, target interface{}) error {

	params := args.ToURLValues()
	c.Logger.Debugf("GET request to %s?%s", path, params.Encode())

	if c.Token != "" {
		params.Set("access_token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		return errors.Wrapf(err, "Invalid GET request %s", url)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return makeHttpClientError(url, resp)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed on %s", url)
	}

	return nil
}

func (c *Client) Post(path string, args Arguments, data map[string][]string, target interface{}) error {
	params := args.ToURLValues()
	c.Logger.Debugf("GET request to %s?%s", path, params.Encode())

	if c.Token != "" {
		params.Set("access_token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	resp, err := http.PostForm(urlWithParams, data)
	if err != nil {
		return errors.Wrapf(err, "Invalid POST request %s", url)
	}

	// resp, err := c.client.Do(req)
	// if err != nil {
	// 	return errors.Wrapf(err, "HTTP request failure on %s", url)
	// }
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "HTTP Read error on response for %s", url)
	}

	fmt.Println(b)

	decoder := json.NewDecoder(bytes.NewBuffer(b))
	err = decoder.Decode(target) // This can't be target, really has to be Block struct but how?

	if err != nil {
		return errors.Wrapf(err, "JSON decode failed on %s:\n%s", url, string(b))
	}

	return nil
}
