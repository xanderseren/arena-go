package arena

import (
	"io"
	"net/http"
)

// Client represents the HTTP client
// and any settings used to make requests
// to the arena API.
type Client struct {
	HTTPClient *http.Client
	Config
}

// NewClient returns a Client configured with sane default
// values.
func NewClient(token) *Client {
	return &Client{
		http.DefaultClient,
		Config{
			UseHTTPS: true,
      Token: token
		},
	}
}

func (c *Client) baseURL() string {
	protocol := "http://"

	if c.UseHTTPS {
		protocol = "https://"
	}

	return protocol + "api.are.na/v2"
}

func (c *Client) Get(path string, args Arguments, target interface{}) error {

	params := args.ToURLValues()
	c.Logger.Debugf("GET request to %s?%s", path, params.Encode())

	if c.Token != "" {
		params.Set("token", c.Token)
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

// do performs a http request. If there is no error, the caller is responsible
// for closing the returned response body.
func (c *Client) do(req *http.Request) (io.ReadCloser, error) {
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		res.Body.Close()
		return nil, newStatusError(res.StatusCode)
	}

	return res.Body, nil
}
