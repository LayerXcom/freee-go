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
	APIPathItems = "items"
)

type Items struct {
	Items []Item `json:"items"`
}

type ItemResponse struct {
	Item Item `json:"item"`
}

type Item struct {
	// 品目ID
	ID int64 `json:"id"`
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 品目名 (30文字以内)
	Name string `json:"name"`
	// 更新日(yyyy-mm-dd)
	UpdateDate string `json:"update_date"`
	// 品目の使用設定（true: 使用する、false: 使用しない）
	Available bool `json:"available"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
	// コード
	Code *string `json:"code,omitempty"`
}

type GetItemsOpts struct {
	Offset int64 `url:"offset,omitempty"`
	Limit  int64 `url:"limit,omitempty"`
}

type ItemParams struct {
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 品目名 (30文字以内)
	Name string `json:"name"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
	// コード（利用を有効にしている場合は必須）
	Code *string `json:"code,omitempty"`
}

func (c *Client) GetItems(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, opts GetItemsOpts,
) (*Items, *oauth2.Token, error) {
	var result Items

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathItems, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) CreateItem(
	ctx context.Context, oauth2Token *oauth2.Token,
	params ItemParams,
) (*Item, *oauth2.Token, error) {
	var result ItemResponse
	oauth2Token, err := c.call(ctx, APIPathItems, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Item, oauth2Token, nil
}

func (c *Client) UpdateItem(
	ctx context.Context, oauth2Token *oauth2.Token,
	params ItemParams,
	itemID int64,
) (*Item, *oauth2.Token, error) {
	var result ItemResponse
	oauth2Token, err := c.call(ctx, path.Join(APIPathItems, fmt.Sprint(itemID)), http.MethodPut, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Item, oauth2Token, nil
}

func (c *Client) DestroyItem(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, itemID int64,
) (*oauth2.Token, error) {
	v, err := query.Values(nil)
	if err != nil {
		return oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathItems, fmt.Sprint(itemID)), http.MethodDelete, oauth2Token, v, nil, nil)
	if err != nil {
		return oauth2Token, err
	}

	return oauth2Token, nil
}
