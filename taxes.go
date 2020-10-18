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
	APIPathTaxes = "taxes"
)

type TaxCompanies struct {
	TaxCompanies []TaxCompany `json:"taxes"`
}

type TaxCompany struct {
	// 税区分コード
	Code int32 `json:"code"`
	// 税区分名
	Name string `json:"name"`
	// 税区分名（日本語表示用）
	NameJa string `json:"name_ja"`
	// 税区分の表示カテゴリ（tax_5: 5%表示の税区分、tax_8: 8%表示の税区分、tax_r8: 軽減税率8%表示の税区分、tax_10: 10%表示の税区分、null: 税率未設定税区分）
	DisplayCategory *string `json:"display_category"`
	// true: 使用する、false: 使用しない
	Available bool `json:"available"`
}

func (c *Client) GetTaxCompanies(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32,
) (*TaxCompanies, *oauth2.Token, error) {
	var result TaxCompanies

	v, err := query.Values(nil)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	tokenSource, err := c.call(ctx, path.Join(APIPathTaxes, "companies", fmt.Sprint(companyID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, token, nil
}
