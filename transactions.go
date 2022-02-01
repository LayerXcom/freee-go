package freee

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
	"net/http"
	"path"
)

const (
	APIPathTxns = "wallet_txns"

	WalletTypeBankAccount = "bank_account"
	WalletTypeCreditCard = "credit_card"
	WalletTypeWallet = "wallet" //現金

	TxnsTypeIncome      = "income"
	TxnsTypeExpense      = "expense"
)

type WalletTxnsResponse struct {
	WalletTxns []WalletTxn `json:"wallet_txns"`
}

type WalletTxnResponse struct {
	WalletTxn WalletTxn `json:"wallet_txn"`
}

type GetWalletTxnOpts struct {
	// walletable_type、walletable_idは同時に指定が必要です。
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	WalletableType string `url:"walletable_type,omitempty"`
	// 口座ID
	WalletableID   uint64  `url:"walletable_id,omitempty"`

	// 取引日で絞込：開始日 (yyyy-mm-dd)
	StartDate 	 string `url:"start_date,omitempty"`
	// 取引日で絞込：終了日 (yyyy-mm-dd)
	EndDate 	 string `url:"end_date,omitempty"`
	// 入金／出金 (入金: income, 出金: expense)
	EntrySide string `url:"entry_side,omitempty"`
	// 取得レコードのオフセット (デフォルト: 0)
	Offset   uint32 `url:"offset,omitempty"`
	// 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100)
	Limit    uint32 `url:"limit,omitempty"`
}

type WalletTxn struct {
	// 明細ID
	ID uint64 `json:"id"`
	// 事業所ID
	CompanyID uint32 `json:"company_id"`
	// 取引日（yyyy-mm-dd）
	Date string `json:"date"`
	// 取引金額
	Amount int32 `json:"amount"`
	// 未決済金額
	DueAmount int32 `json:"due_amount"`
	// 残高
	Balance int32 `json:"balance"`
	// 入金/出勤（入金: income, 出勤: expense）
	EntrySide string `json:"entry_side"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	WalletableType string `json:"walletable_type"`
	// 口座ID
	WalletableID uint64 `json:"walletable_id"`
	// 取引内容
	Description string `json:"description"`
	// 明細のステータス（消込待ち: 1, 消込済み: 2, 無視: 3, 消込中: 4）
	Status uint `json:"status"`
}

func (c *Client) GetWalletTransactions(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, opts GetWalletTxnOpts) (*WalletTxnsResponse,  *oauth2.Token, error) {
	var result WalletTxnsResponse

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
	companyID uint32, txnID uint64, opts GetWalletTxnOpts,
) (*WalletTxn,  *oauth2.Token, error) {
	var result WalletTxnResponse

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
