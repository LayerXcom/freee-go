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
	APIPathTxns = "wallet_txns"

	TxnsTypeIncome  = "income"
	TxnsTypeExpense = "expense"
)

type WalletTxnsResponse struct {
	WalletTxns []WalletTxn `json:"wallet_txns"`
}

type WalletTxnResponse struct {
	WalletTxn WalletTxn `json:"wallet_txn"`
}

type WalletTxn struct {
	// 明細ID
	ID int64 `json:"id"`
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 取引日(yyyy-mm-dd)
	Date string `json:"date"`
	// 取引金額
	Amount int64 `json:"amount"`
	// 未決済金額
	DueAmount int64 `json:"due_amount"`
	// 残高(銀行口座等)
	Balance int64 `json:"balance"`
	// 入金／出金 (入金: income, 出金: expense)
	EntrySide string `json:"entry_side"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	WalletableType string `json:"walletable_type"`
	// 口座ID
	WalletableID int64 `json:"walletable_id"`
	// 取引内容
	Description string `json:"description"`
	// 明細のステータス（消込待ち: 1, 消込済み: 2, 無視: 3, 消込中: 4, 対象外: 6）
	Status int `json:"status"`
	// 登録時に<a href=\"https://support.freee.co.jp/hc/ja/articles/202848350-明細の自動登録ルールを設定する\" target=\"_blank\">自動登録ルールの設定</a>が適用され、登録処理が実行された場合、 trueになります。〜を推測する、〜の消込をするの条件の場合は一致してもfalseになります。
	RuleMatched bool `json:"rule_matched"`
}

type GetWalletTxnOpts struct {
	// walletable_type、walletable_idは同時に指定が必要です。
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	WalletableType string `url:"walletable_type,omitempty"`
	// 口座ID
	WalletableID int64 `url:"walletable_id,omitempty"`
	// 取引日で絞込：開始日 (yyyy-mm-dd)
	StartDate string `url:"start_date,omitempty"`
	// 取引日で絞込：終了日 (yyyy-mm-dd)
	EndDate string `url:"end_date,omitempty"`
	// 入金／出金 (入金: income, 出金: expense)
	EntrySide string `url:"entry_side,omitempty"`
	// 取得レコードのオフセット (デフォルト: 0)
	Offset int64 `url:"offset,omitempty"`
	// 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100)
	Limit int64 `url:"limit,omitempty"`
}

type WalletTxnCreateParams struct {
	// 入金／出金 (入金: income, 出金: expense)
	EntrySide string `json:"entry_side"`
	// 取引内容
	Description *string `json:"description,omitempty"`
	// 取引金額
	Amount int64 `json:"amount"`
	// 口座ID
	WalletableID int64 `json:"walletable_id"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	WalletableType string `json:"walletable_type"`
	// 取引日 (yyyy-mm-dd)
	Date string `json:"date"`
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 残高 (銀行口座等)
	Balance *int64 `json:"balance,omitempty"`
}

func (c *Client) GetWalletTransactions(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, opts GetWalletTxnOpts) (*WalletTxnsResponse, *oauth2.Token, error) {
	var result WalletTxnsResponse

	if (opts.WalletableType != "" && opts.WalletableID == 0) || (opts.WalletableID != 0 && opts.WalletableType == "") {
		return nil, oauth2Token, fmt.Errorf("either walletable_type or walletable_id is specified, then other value must be set")
	}

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}

	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathTxns), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) GetWalletTransaction(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, txnID int64, opts GetWalletTxnOpts,
) (*WalletTxn, *oauth2.Token, error) {
	var result WalletTxnResponse

	if (opts.WalletableType != "" && opts.WalletableID == 0) || (opts.WalletableID != 0 && opts.WalletableType == "") {
		return nil, oauth2Token, fmt.Errorf("either walletable_type or walletable_id is specified, then other value must be set")
	}

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}

	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathTxns, fmt.Sprint(txnID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result.WalletTxn, oauth2Token, nil
}

func (c *Client) CreateWalletTransaction(
	ctx context.Context, oauth2Token *oauth2.Token,
	params WalletTxnCreateParams,
) (*WalletTxnResponse, *oauth2.Token, error) {
	var result WalletTxnResponse
	oauth2Token, err := c.call(ctx, path.Join(APIPathTxns), http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}
