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

	// tax_5: 5%表示の税区分
	TaxRate5 = "tax_5"
	// tax_8: 8%表示の税区分
	TaxRate8 = "tax_8"
	// tax_r8: 軽減税率8%表示の税区分
	TaxRateR8 = "tax_r8"
	// tax_10: 10%表示の税区分
	TaxRate10 = "tax_10"
	// null: 税率未設定税区分
)

type TaxCompanies struct {
	TaxCompanies []TaxCompany `json:"taxes"`
}

type TaxCompany struct {
	// 税区分コード
	Code int64 `json:"code"`
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
	companyID int64,
) (*TaxCompanies, *oauth2.Token, error) {
	var result TaxCompanies

	v, err := query.Values(nil)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathTaxes, "companies", fmt.Sprint(companyID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}
