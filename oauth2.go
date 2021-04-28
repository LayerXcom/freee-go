package freee

import (
	"context"

	"golang.org/x/oauth2"
)

const (
	Oauth2TokenURL = "https://accounts.secure.freee.co.jp/public_api/token"
	Oauth2AuthURL  = "https://accounts.secure.freee.co.jp/public_api/authorize"
)

func (c *Client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return c.config.Oauth2.AuthCodeURL(state, opts...)
}

func (c *Client) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return c.config.Oauth2.Exchange(ctx, code, opts...)
}
