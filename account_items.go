package freee

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathAccountItems = "account_items"
)

type GetAccountItemsOpts struct {
	BaseDate string `url:"base_date,omitempty"`
}

type AccountItems struct {
	AccountItems []AccountItem `json:"account_items"`
}

type AccountItem struct {
	// 勘定科目ID
	ID int32 `json:"id"`
	// 勘定科目名 (30文字以内)
	Name string `json:"name"`
	// ショートカット1 (20文字以内)
	Shortcut *string `json:"shortcut,omitempty"`
	// ショートカット2(勘定科目コード) (20文字以内)
	ShortcutNum *string `json:"shortcut_num,omitempty"`
	// デフォルト設定がされている税区分ID
	DefaultTaxID int32 `json:"default_tax_id,omitempty"`
	// デフォルト設定がされている税区分コード
	DefaultTaxCode int32 `json:"default_tax_code"`
	// 勘定科目カテゴリー
	AccountCategory string `json:"account_category"`
	// 勘定科目のカテゴリーID
	AccountCategoryID int32    `json:"account_category_id"`
	Categories        []string `json:"categories"`
	// 勘定科目の使用設定（true: 使用する、false: 使用しない）
	Available bool `json:"available"`
	// 口座ID
	WalletableID *int32 `json:"walletable_id"`
	// 決算書表示名（小カテゴリー）
	GroupName *string `json:"group_name,omitempty"`
	// 収入取引相手勘定科目名
	CorrespondingIncomeName *string `json:"corresponding_income_name,omitempty"`
	// 収入取引相手勘定科目ID
	CorrespondingIncomeID *int32 `json:"corresponding_income_id,omitempty"`
	// 支出取引相手勘定科目名
	CorrespondingExpenseName *string `json:"corresponding_expense_name,omitempty"`
	// 支出取引相手勘定科目ID
	CorrespondingExpenseID *int32 `json:"corresponding_expense_id,omitempty"`
}

func (c *Client) GetAccountItems(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, opts GetAccountItemsOpts,
) (*AccountItems, *oauth2.Token, error) {
	var result AccountItems

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	tokenSource, err := c.call(ctx, APIPathAccountItems, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, token, nil
}
