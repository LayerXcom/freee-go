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

type GetSegmentTagsOpts struct {
	Offset uint32 `url:"offset,omitempty"`
	Limit  uint32 `url:"limit,omitempty"`
}

func (c *Client) GetSegmentTags(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, tagID uint32, opts GetSegmentTagsOpts,
) (*SegmentTags, *oauth2.Token, error) {
	var result SegmentTags

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathSegments, fmt.Sprint(tagID), "tags"), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}
