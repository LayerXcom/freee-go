# Go API client for freee

会計 freee API のクライアントライブラリ

[![Go Reference](https://pkg.go.dev/badge/github.com/layerxcom/freee-go.svg)](https://pkg.go.dev/github.com/layerxcom/freee-go)
[![test](https://github.com/LayerXcom/freee-go/actions/workflows/tests.yaml/badge.svg)](https://github.com/LayerXcom/freee-go/actions/workflows/tests.yaml)

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/LayerXcom/freee-go"
	"golang.org/x/oauth2"
)

func main() {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")
	conf := freee.NewConfig(clientID, clientSecret, redirectURL)
	conf.Log = log.New(os.Stdout, "", log.LstdFlags)
	client := freee.NewClient(conf)

	ctx := context.Background()
	token := &oauth2.Token{
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
	}
	me, token, err := client.GetUsersMe(ctx, token, freee.GetUsersMeOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", token)
}
```

## References

- [会計 API リファレンス](https://developer.freee.co.jp/docs/accounting/reference#/)
- [freee API のGoクライアントライブラリを公開しました - LayerX エンジニアブログ](https://tech.layerx.co.jp/entry/2021/05/12/110000)

## APIs

### 勘定科目

- [ ] GET /api/1/account_items 勘定科目一覧の取得
- [ ] POST /api/1/account_items 勘定科目の作成
- [x] GET /api/1/account_items/{id} 勘定科目の詳細情報の取得
- [ ] PUT /api/1/account_items/{id} 勘定科目の更新
- [ ] DELETE /api/1/account_items/{id} 勘定科目の削除

### 申請経路

- [ ] GET /api/1/approval_flow_routes 申請経路一覧の取得
- [ ] GET /api/1/approval_flow_routes/{id} 申請経路の取得

### 各種申請

- [ ] GET /api/1/approval_requests 各種申請の一覧
- [ ] POST /api/1/approval_requests 各種申請の作成
- [ ] GET /api/1/approval_requests/{id} 各種申請の取得
- [ ] PUT /api/1/approval_requests/{id} 各種申請の更新
- [ ] DELETE /api/1/approval_requests/{id} 各種申請の削除
- [ ] GET /api/1/approval_requests/forms 各種申請の申請フォーム一覧の取得
- [ ] GET /api/1/approval_requests/forms/{id} 各種申請の申請フォームの取得
- [ ] POST /api/1/approval_requests/{id}/actions 各種申請の承認操作

### 連携サービス

- [ ] GET /api/1/banks 連携サービス一覧の取得
- [ ] GET /api/1/banks/{id} 連携サービスの取得

### 事業所

- [ ] GET /api/1/companies 事業所一覧の取得
- [ ] GET /api/1/companies/{id} 事業所の詳細情報の取得

### 取引

- [ ] GET /api/1/deals 取引（収入／支出）一覧の取得
- [x] POST /api/1/deals 取引（収入／支出）の作成
- [x] GET /api/1/deals/{id} 取引（収入／支出）の取得
- [x] PUT /api/1/deals/{id} 取引（収入／支出）の更新
- [x] DELETE /api/1/deals/{id} 取引（収入／支出）の削除

### 経費科目

- [ ] GET /api/1/expense_application_line_templates 経費科目一覧の取得
- [ ] POST /api/1/expense_application_line_templates 経費科目の作成
- [ ] GET /api/1/expense_application_line_templates/{id} 経費科目の取得
- [ ] PUT /api/1/expense_application_line_templates/{id} 経費科目の更新
- [ ] DELETE /api/1/expense_application_line_templates/{id} 経費科目の削除

### 経費精算

- [ ] GET /api/1/expense_applications 経費申請一覧の取得
- [ ] POST /api/1/expense_applications 経費申請の作成
- [ ] GET /api/1/expense_applications/{id} 経費申請詳細の取得
- [ ] PUT /api/1/expense_applications/{id} 経費申請の更新
- [ ] DELETE /api/1/expense_applications/{id} 経費申請の削除
- [ ] POST /api/1/expense_applications/{id}/actions 経費申請の承認操作

### 請求書

- [ ] GET /api/1/invoices 請求書一覧の取得
- [ ] POST /api/1/invoices 請求書の作成
- [ ] GET /api/1/invoices/{id} 請求書の取得
- [ ] PUT /api/1/invoices/{id} 請求書の更新
- [ ] DELETE /api/1/invoices/{id} 請求書の削除

### 品目

- [ ] GET /api/1/items 品目一覧の取得
- [x] POST /api/1/items 品目の作成
- [x] GET /api/1/items/{id} 品目の取得
- [x] PUT /api/1/items/{id} 品目の更新
- [x] DELETE /api/1/items/{id} 品目の削除

### 仕訳帳

- [ ] GET /api/1/journals ダウンロード要求
- [ ] GET /api/1/journals/reports/{id}/status ステータス確認
- [ ] GET /api/1/journals/reports/{id}/download ダウンロード実行

### 振替伝票

- [ ] GET /api/1/manual_journals 振替伝票一覧の取得
- [ ] POST /api/1/manual_journals 振替伝票の作成
- [x] GET /api/1/manual_journals/{id} 振替伝票の取得
- [x] PUT /api/1/manual_journals/{id} 振替伝票の更新
- [x] DELETE /api/1/manual_journals/{id} 振替伝票の削除

### 取引先

- [x] GET /api/1/partners 取引先一覧の取得
- [x] POST /api/1/partners 取引先の作成
- [ ] GET /api/1/partners/{id} 取引先の取得
- [x] PUT /api/1/partners/{id} 取引先の更新
- [x] DELETE /api/1/partners/{id} 取引先の削除
- [ ] PUT /api/1/partners/code/{code} 取引先の更新

### 取引の支払行

- [ ] POST /api/1/deals/{id}/payments 取引（収入／支出）の支払行作成
- [ ] PUT /api/1/deals/{id}/payments/{payment_id} 取引（収入／支出）の支払行更新
- [ ] DELETE /api/1/deals/{id}/payments/{payment_id} 取引（収入／支出）の支払行削除

### 見積書

- [ ] GET /api/1/quotations 見積書一覧の取得
- [ ] POST /api/1/quotations 見積書の作成
- [ ] GET /api/1/quotations/{id} 見積書の取得
- [ ] PUT /api/1/quotations/{id} 見積書の更新
- [ ] DELETE /api/1/quotations/{id} 見積書の削除

### ファイルボックス

- [ ] GET /api/1/receipts ファイルボックス 証憑ファイル一覧の取得
- [x] POST /api/1/receipts ファイルボックス 証憑ファイルアップロード
- [x] GET /api/1/receipts/{id} ファイルボックス 証憑ファイルの取得
- [ ] PUT /api/1/receipts/{id} ファイルボックス 証憑ファイル情報更新
- [ ] DELETE /api/1/receipts/{id} ファイルボックス 証憑ファイルを削除する
- [ ] GET /api/1/receipts/{id}/download ファイルボックス 証憑ファイルのダウンロード

### 取引の+更新

- [ ] POST /api/1/deals/{id}/renews 取引（収入／支出）に対する+更新の作成
- [ ] PUT /api/1/deals/{id}/renews/{renew_id} 取引（収入／支出）の+更新の更新
- [ ] DELETE /api/1/deals/{id}/renews/{renew_id} 取引（収入／支出）の+更新の削除

### 部門

- [x] GET /api/1/sections 部門一覧の取得
- [x] POST /api/1/sections 部門の作成
- [ ] GET /api/1/sections/{id} 部門の取得
- [x] PUT /api/1/sections/{id} 部門の更新
- [ ] DELETE /api/1/sections/{id} 部門の削除

### セグメントタグ

- [x] GET /api/1/segments/{segment_id}/tags セグメントタグ一覧の取得
- [x] POST /api/1/segments/{segment_id}/tags セグメントの作成
- [x] PUT /api/1/segments/{segment_id}/tags/{id} セグメントタグの更新
- [x] DELETE /api/1/segments/{segment_id}/tags/{id} セグメントタグの削除

### フォーム用選択項目情報

- [ ] GET /api/1/forms/selectables フォーム用選択項目情報の取得

### メモタグ

- [x] GET /api/1/tags メモタグ一覧の取得
- [x] POST /api/1/tags メモタグの作成
- [ ] GET /api/1/tags/{id} メモタグの詳細情報の取得
- [x] PUT /api/1/tags/{id} メモタグの更新
- [x] DELETE /api/1/tags/{id} メモタグの削除

### 税区分

- [ ] GET /api/1/taxes/codes 税区分コード一覧の取得
- [ ] GET /api/1/taxes/codes/{code} 税区分コードの取得
- [x] GET /api/1/taxes/companies/{company_id} 税区分コード詳細一覧の取得

### 取引（振替）

- [ ] GET /api/1/transfers 取引（振替）一覧の取得
- [ ] POST /api/1/transfers 取引（振替）の作成
- [ ] GET /api/1/transfers/{id} 取引（振替）の取得
- [ ] PUT /api/1/transfers/{id} 取引（振替）の更新
- [ ] DELETE /api/1/transfers/{id} 取引（振替）の削除する

### 試算表

- [ ] GET /api/1/reports/trial_bs 貸借対照表の取得
- [ ] GET /api/1/reports/trial_bs_two_years 貸借対照表(前年比較)の取得
- [ ] GET /api/1/reports/trial_bs_three_years 貸借対照表(３期間比較)の取得
- [ ] GET /api/1/reports/trial_pl 損益計算書の取得
- [ ] GET /api/1/reports/trial_pl_two_years 損益計算書(前年比較)の取得
- [ ] GET /api/1/reports/trial_pl_three_years 損益計算書(３期間比較)の取得
- [ ] GET /api/1/reports/trial_pl_sections 損益計算書(部門比較)の取得

### ユーザー

- [ ] GET /api/1/users 事業所に所属するユーザー一覧の取得
- [ ] GET /api/1/users/capabilities ログインユーザーの権限の取得
- [x] GET /api/1/users/me ログインユーザー情報の取得
- [ ] PUT /api/1/users/me ユーザー情報の更新

### 明細

- [x] GET /api/1/wallet_txns 明細一覧の取得
- [ ] POST /api/1/wallet_txns 明細の作成
- [x] GET /api/1/wallet_txns/{id} 明細の取得
- [ ] DELETE /api/1/wallet_txns/{id} 明細の削除

### 口座

- [ ] GET /api/1/walletables 口座一覧の取得
- [ ] POST /api/1/walletables 口座の作成
- [ ] GET /api/1/walletables/{type}/{id} 口座情報の取得
- [ ] PUT /api/1/walletables/{type}/{id} 口座の更新
- [ ] DELETE /api/1/walletables/{type}/{id} 口座の削除
