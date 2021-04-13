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
	APIPathSegments = "segments"
	SegmentID1      = uint32(1)
	SegmentID2      = uint32(2)
	SegmentID3      = uint32(3)
)

type SegmentTags struct {
	SegmentTags []SegmentTag `json:"segment_tags"`
}

type SegmentTagResponse struct {
	SegmentTag SegmentTag `json:"segment_tag"`
}

type SegmentTag struct {
	// セグメントタグID
	ID int32 `json:"id"`
	// セグメントタグ名
	Name string `json:"name"`
	// 備考
	Description *string `json:"description"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2"`
}

type SegmentTagParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// セグメントタグ名 (30文字以内)
	Name string `json:"name"`
	// 備考 (30文字以内)
	Description *string `json:"description,omitempty"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

type GetSegmentTagsOpts struct {
	Offset uint32 `url:"offset,omitempty"`
	Limit  uint32 `url:"limit,omitempty"`
}

func (c *Client) GetSegmentTags(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, segmentID uint32,
	opts GetSegmentTagsOpts,
) (*SegmentTags, *oauth2.Token, error) {
	var result SegmentTags

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathSegments, fmt.Sprint(segmentID), "tags"), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) CreateSegmentTag(
	ctx context.Context, oauth2Token *oauth2.Token,
	segmentID uint32, params SegmentTagParams,
) (*SegmentTag, *oauth2.Token, error) {
	var result SegmentTagResponse
	oauth2Token, err := c.call(ctx, path.Join(APIPathSegments, fmt.Sprint(segmentID), "tags"), http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.SegmentTag, oauth2Token, nil
}

func (c *Client) UpdateSegmentTag(
	ctx context.Context, oauth2Token *oauth2.Token,
	segmentID uint32, id uint32,
	params SegmentTagParams,
) (*SegmentTag, *oauth2.Token, error) {
	var result SegmentTagResponse
	oauth2Token, err := c.call(ctx, path.Join(APIPathSegments, fmt.Sprint(segmentID), "tags", fmt.Sprint(id)), http.MethodPut, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.SegmentTag, oauth2Token, nil
}

func (c *Client) DestroySegmentTag(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32,
	segmentID uint32, id uint32,
) (*oauth2.Token, error) {
	v, err := query.Values(nil)
	if err != nil {
		return oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathSegments, fmt.Sprint(segmentID), "tags", fmt.Sprint(id)), http.MethodDelete, oauth2Token, v, nil, nil)
	if err != nil {
		return oauth2Token, err
	}

	return oauth2Token, nil
}
