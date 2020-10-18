package freee

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathTags = "tags"
)

type Tags struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	// タグID
	ID int32 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 名前(30文字以内)
	Name *string `json:"name"`
	// ショートカット1 (255文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット2 (255文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

type GetTagsOpts struct {
	Offset uint32 `url:"offset,omitempty"`
	Limit  uint32 `url:"limit,omitempty"`
}

func (c *Client) GetTags(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, opts GetTagsOpts,
) (*Tags, *oauth2.Token, error) {
	var result Tags

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	tokenSource, err := c.call(ctx, APIPathTags, http.MethodGet, oauth2Token, v, opts, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, token, nil
}
