package freee

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathTags = "tags"
)

type Tags struct {
	Tags []Tag `json:"tags"`
}

type TagResponse struct {
	Tag Tag `json:"tag"`
}

type Tag struct {
	// タグID
	ID int64 `json:"id"`
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 名前(30文字以内)
	Name *string `json:"name"`
	// ショートカット1 (255文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット2 (255文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

type TagParams struct {
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// メモタグ名 (30文字以内)
	Name string `json:"name"`
	// メモタグ検索用 (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// メモタグ検索用 (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

type GetTagsOpts struct {
	Offset int64 `url:"offset,omitempty"`
	Limit  int64 `url:"limit,omitempty"`
}

func (c *Client) GetTags(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, opts GetTagsOpts,
) (*Tags, *oauth2.Token, error) {
	var result Tags

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathTags, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) CreateTag(
	ctx context.Context, oauth2Token *oauth2.Token,
	params TagParams,
) (*Tag, *oauth2.Token, error) {
	var result TagResponse
	oauth2Token, err := c.call(ctx, APIPathTags, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Tag, oauth2Token, nil
}

func (c *Client) GetTag(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, tagID int64, opts GetTagsOpts,
) (*Tags, *oauth2.Token, error) {
	var result Tags

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathTags, fmt.Sprint(tagID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) UpdateTag(
	ctx context.Context, oauth2Token *oauth2.Token,
	tagID int64, params TagParams,
) (*Tag, *oauth2.Token, error) {
	var result TagResponse
	oauth2Token, err := c.call(ctx, path.Join(APIPathTags, fmt.Sprint(tagID)), http.MethodPut, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Tag, oauth2Token, nil
}

func (c *Client) DestroyTag(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, tagID int64,
) (*oauth2.Token, error) {
	v, err := query.Values(nil)
	if err != nil {
		return oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathTags, fmt.Sprint(tagID)), http.MethodDelete, oauth2Token, v, nil, nil)
	if err != nil {
		return oauth2Token, err
	}

	return oauth2Token, nil
}
