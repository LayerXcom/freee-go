package freee

import (
	"golang.org/x/oauth2"
)

func (c *Client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return c.config.Oauth2.AuthCodeURL(state, opts...)
}
