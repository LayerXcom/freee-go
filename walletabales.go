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
	APIPathWalletables = "walletables"

	// 口座種別
	WalletTypeBankAccount = "bank_account" // 銀行口座
	WalletTypeCreditCard  = "credit_card"  // クレジットカード
	WalletTypeWallet      = "wallet"       //現金, その他の決済口座
)

type WalletablesResponse struct {
	Walletables []Walletable `json:"walletables"`
	Meta        Meta         `json:"meta"`
}

type WalletableResponse struct {
	Walletable Walletable `json:"walletable"`
	Meta       Meta       `json:"meta"`
}

type Meta struct {
	UpToDate bool `json:"up_to_date"`
}

type GetWalletablesOpts struct {
	// 残高情報を含める
	WithBalance bool `url:"with_balance,omitempty"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, その他の決済口座: wallet)
	Type string `url:"type,omitempty"`
}

type Walletable struct {
	// 口座ID
	ID uint64 `json:"id"`
	// 口座名 (255文字以内)
	Name string `json:"name"`
	// サービスID
	BankID uint64 `json:"bank_id"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	Type string `json:"type"`
	// 同期残高
	LastBalance int64 `json:"last_balance,omitempty"`
	// 登録残高
	WalletableBalance int64 `json:"walletable_balance,omitempty"`
}

func (c *Client) GetWalletables(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, opts GetWalletablesOpts,
) (*WalletablesResponse, *oauth2.Token, error) {
	var result WalletablesResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}

	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathWalletables), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) GetWalletable(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, walletType string, walletableID uint64,
) (*Walletable, *oauth2.Token, error) {
	var result WalletableResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, oauth2Token, err
	}

	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathWalletables, walletType, fmt.Sprint(walletableID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result.Walletable, oauth2Token, nil
}
