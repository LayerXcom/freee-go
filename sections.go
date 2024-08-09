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
	APIPathSections = "sections"
)

type Sections struct {
	Sections []Section `json:"sections"`
}

type SectionResponse struct {
	Section Section `json:"section"`
}

type Section struct {
	// 部門ID
	ID int64 `json:"id"`
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 部門コード
	Code *string `json:"code,omitempty"`
	// 部門名 (30文字以内)
	Name string `json:"name"`
	// 部門の使用設定（true: 使用する、false: 使用しない）
	Available bool `json:"available"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

type SectionParams struct {
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 部門コード（利用を有効にしている場合は必須）
	Code *string `json:"code,omitempty"`
	// 部門名 (30文字以内)
	Name string `json:"name"`
	// 正式名称 (255文字以内)
	LongName *string `json:"long_name,omitempty"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
	// 親部門ID (ビジネスプラン以上)
	ParentID *int64 `json:"parent_id,omitempty"`
}

func (c *Client) GetSections(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64,
) (*Sections, *oauth2.Token, error) {
	var result Sections

	v, err := query.Values(nil)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathSections, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) CreateSection(
	ctx context.Context, oauth2Token *oauth2.Token,
	params SectionParams,
) (*Section, *oauth2.Token, error) {
	var result SectionResponse
	oauth2Token, err := c.call(ctx, APIPathSections, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Section, oauth2Token, nil
}

func (c *Client) UpdateSection(
	ctx context.Context, oauth2Token *oauth2.Token,
	sectionID int64, params SectionParams,
) (*Section, *oauth2.Token, error) {
	var result SectionResponse
	oauth2Token, err := c.call(ctx, path.Join(APIPathSections, fmt.Sprint(sectionID)), http.MethodPut, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Section, oauth2Token, nil
}

func (c *Client) DestroySection(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, sectionID int64,
) (*oauth2.Token, error) {
	v, err := query.Values(nil)
	if err != nil {
		return oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathSections, fmt.Sprint(sectionID)), http.MethodDelete, oauth2Token, v, nil, nil)
	if err != nil {
		return oauth2Token, err
	}

	return oauth2Token, nil
}
