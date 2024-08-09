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
	ID int64 `json:"id"`
	// 勘定科目コード
	Code *string `json:"code,omitempty"`
	// 勘定科目名 (30文字以内)
	Name string `json:"name"`
	// ショートカット1 (20文字以内)
	Shortcut *string `json:"shortcut,omitempty"`
	// ショートカット2(勘定科目コード) (20文字以内)
	ShortcutNum *string `json:"shortcut_num,omitempty"`
	// 税区分コード
	TaxCode int64 `json:"tax_code"`
	// デフォルト設定がされている税区分ID
	DefaultTaxID int64 `json:"default_tax_id,omitempty"`
	// デフォルト設定がされている税区分コード
	DefaultTaxCode int64 `json:"default_tax_code"`
	// 勘定科目カテゴリー
	AccountCategory string `json:"account_category"`
	// 勘定科目のカテゴリーID
	AccountCategoryID int64    `json:"account_category_id"`
	Categories        []string `json:"categories"`
	// 勘定科目の使用設定（true: 使用する、false: 使用しない）
	Available bool `json:"available"`
	// 口座ID
	WalletableID *int64 `json:"walletable_id"`
	// 決算書表示名（小カテゴリー）
	GroupName *string `json:"group_name,omitempty"`
	// 収入取引相手勘定科目名
	CorrespondingIncomeName *string `json:"corresponding_income_name,omitempty"`
	// 収入取引相手勘定科目ID
	CorrespondingIncomeID *int64 `json:"corresponding_income_id,omitempty"`
	// 支出取引相手勘定科目名
	CorrespondingExpenseName *string `json:"corresponding_expense_name,omitempty"`
	// 支出取引相手勘定科目ID
	CorrespondingExpenseID *int64 `json:"corresponding_expense_id,omitempty"`
}

type AccountItemCreateParams struct {
	// 事業所ID
	CompanyID   int64                              `json:"company_id"`
	AccountItem AccountItemCreateParamsAccountItem `json:"account_item"`
}

type AccountItemCreateParamsAccountItem struct {
	// 勘定科目カテゴリーID Selectablesフォーム用選択項目情報エンドポイント(account_groups.account_category_id)で取得可能です
	AccountCategoryID int64 `json:"account_category_id"`
	// 減価償却累計額勘定科目ID（法人のみ利用可能）
	AccumulatedDepAccountItemID *int64 `json:"accumulated_dep_account_item_id,omitempty"`
	// 勘定科目コード
	Code *string `json:"code,omitempty"`
	// 支出取引相手勘定科目ID
	CorrespondingExpenseID int64 `json:"corresponding_expense_id"`
	// 収入取引相手勘定科目ID
	CorrespondingIncomeID int64 `json:"corresponding_income_id"`
	// 決算書表示名（小カテゴリー） Selectablesフォーム用選択項目情報エンドポイント(account_groups.name)で取得可能です
	GroupName string `json:"group_name"`
	// 品目
	Items *[]AccountItemCreateParamsItems `json:"items,omitempty"`
	// 勘定科目名 (30文字以内)
	Name string `json:"name"`
	// 取引先
	Partners *[]AccountItemCreateParamsPartners `json:"partners,omitempty"`
	// 検索可能:2, 検索不可：3(登録時未指定の場合は2で登録されます。更新時未指定の場合はsearchableは変更されません。)
	Searchable *int64 `json:"searchable,omitempty"`
	// ショートカット1 (20文字以内)
	Shortcut *string `json:"shortcut,omitempty"`
	// ショートカット2 (20文字以内)
	ShortcutNum *string `json:"shortcut_num,omitempty"`
	// 税区分コード 指定できるコードは本APIの注意点をご確認ください。
	TaxCode int64 `json:"tax_code"`
}

type AccountItemCreateParamsItems struct {
	ID *int64 `json:"id,omitempty"`
}

type AccountItemCreateParamsPartners struct {
	ID *int64 `json:"id,omitempty"`
}

type AccountItemCreateResponse struct {
	AccountItem AccountItemDetail `json:"account_item"`
}

// AccountItemDetail はPOST /api/1/account_items のレスポンス。
type AccountItemDetail struct {
	// 勘定科目ID
	ID int64 `json:"id"`
	// 勘定科目名 (30文字以内)
	Name string `json:"name"`
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 税区分コード
	TaxCode int64 `json:"tax_code"`
	// 勘定科目カテゴリー
	AccountCategory string `json:"account_category"`
	// 勘定科目のカテゴリーID
	AccountCategoryID int64 `json:"account_category_id"`
	// ショートカット1 (20文字以内)
	Shortcut *string `json:"shortcut,omitempty"`
	// ショートカット2 (20文字以内)
	ShortcutNum *string `json:"shortcut_num,omitempty"`
	// 勘定科目コード (20文字以内)
	Code *string `json:"code"`
	// 検索可能:2, 検索不可：3
	Searchable int64 `json:"searchable"`
	// 減価償却累計額勘定科目（法人のみ利用可能）
	AccumulatedDepAccountItemName *string `json:"accumulated_dep_account_item_name,omitempty"`
	// 減価償却累計額勘定科目ID（法人のみ利用可能）
	AccumulatedDepAccountItemID *int64 `json:"accumulated_dep_account_item_id"`
	// 品目
	Items *[]AccountItemDetalItems `json:"items,omitempty"`
	// 取引先
	Partners *[]AccountItemDetailPartners `json:"partners,omitempty"`
	// 勘定科目の使用設定（true: 使用する、false: 使用しない）
	Available bool `json:"available"`
	// 口座ID
	WalletableID *int64 `json:"walletable_id"`
	// 決算書表示名（小カテゴリー）
	GroupName *string `json:"group_name,omitempty"`
	// 決算書表示名ID（小カテゴリー）
	GroupID *int64 `json:"group_id"`
	// 収入取引相手勘定科目名
	CorrespondingIncomeName *string `json:"corresponding_income_name,omitempty"`
	// 収入取引相手勘定科目ID
	CorrespondingIncomeID *int64 `json:"corresponding_income_id,omitempty"`
	// 支出取引相手勘定科目名
	CorrespondingExpenseName *string `json:"corresponding_expense_name,omitempty"`
	// 支出取引相手勘定科目ID
	CorrespondingExpenseID *int64 `json:"corresponding_expense_id,omitempty"`
}

type AccountItemDetalItems struct {
	// 品目ID
	ID *int64 `json:"id,omitempty"`
	// 品目名
	Name string `json:"name"`
}

type AccountItemDetailPartners struct {
	// 取引先ID
	ID *int64 `json:"id,omitempty"`
	// 取引先名
	Name string `json:"name"`
}

func (c *Client) GetAccountItems(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, opts GetAccountItemsOpts,
) (*AccountItems, *oauth2.Token, error) {
	var result AccountItems

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathAccountItems, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) CreateAccountItem(
	ctx context.Context, oauth2Token *oauth2.Token,
	params AccountItemCreateParams,
) (*AccountItemDetail, *oauth2.Token, error) {
	var result AccountItemCreateResponse
	oauth2Token, err := c.call(ctx, APIPathAccountItems, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.AccountItem, oauth2Token, nil
}
