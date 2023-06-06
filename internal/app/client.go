package app

import (
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
}

func (c *Client) GetResponse(method string, u *url.URL, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, u.String(), body)
	resp, err := c.HTTPClient.Do(req)
	return resp, err
}
func (c *Client) GetResponseBytes(method string, u *url.URL, body io.Reader) ([]byte, error) {
	resp, err := c.GetResponse(method, u, nil)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}
