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
	APIPathDeals = "deals"

	DealTypeIncome  = "income"
	DealTypeExpense = "expense"
)

type DealResponse struct {
	Deal Deal `json:"deal"`
}

type GetDealOpts struct {
	// 取引の債権債務行の表示（without: 表示しない(デフォルト), with: 表示する）
	Accruals string `url:"accruals,omitempty"`
}

type Deal struct {
	// 取引ID
	ID uint64 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 支払期日 (yyyy-mm-dd)
	DueDate *string `json:"due_date,omitempty"`
	// 金額
	Amount int32 `json:"amount"`
	// 支払金額
	DueAmount *int32 `json:"due_amount,omitempty"`
	// 収支区分 (収入: income, 支出: expense)
	Type *string `json:"type,omitempty"`
	// 取引先ID
	PartnerID int32 `json:"partner_id"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 管理番号
	RefNumber *string `json:"ref_number,omitempty"`
	// 決済状況 (未決済: unsettled, 完了: settled)
	Status string `json:"status"`
	// 取引の明細行
	Details *[]DealDetails `json:"details,omitempty"`
	// 取引の+更新行
	Renews *[]DealRenews `json:"renews,omitempty"`
	// 取引の支払行
	Payments *[]DealPayments `json:"payments,omitempty"`
	// 証憑ファイル
	Receipts *[]DealReceipts `json:"receipts,omitempty"`
}

type DealDetails struct {
	ID uint64 `json:"id"`
	// 税区分コード
	TaxCode int32 `json:"tax_code"`
	// 勘定科目ID
	AccountItemID int32 `json:"account_item_id"`
	// 取引金額（税込で指定してください）
	Amount int32 `json:"amount"`
	// 品目ID
	ItemID *int32 `json:"item_id,omitempty"`
	// 部門ID
	SectionID *int32 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs *[]int32 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID *int32 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID *int32 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID *int32 `json:"segment_3_tag_id,omitempty"`
	// 備考
	Description *string `json:"description,omitempty"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int32 `json:"vat,omitempty"`
}

type DealRenews struct {
	// +更新行ID
	ID uint64 `json:"id"`
	// 更新日 (yyyy-mm-dd)
	UpdateDate string `json:"update_date"`
	// +更新の対象行ID
	RenewTargetId int64 `json:"renew_target_id"`
	// +更新の対象行タイプ
	RenewTargetType string `json:"renew_target_type"`
	// +更新の明細行一覧（配列）
	Details []DealDetails `json:"details"`
}

type DealPayments struct {
	// 取引行ID
	ID uint64 `json:"id"`
	// 支払日
	Date string `json:"date"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item)
	FromWalletableType string `json:"from_walletable_type,omitempty"`
	// 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID）
	FromWalletableID int32 `json:"from_walletable_id,omitempty"`
	// 支払金額
	Amount int32 `json:"amount"`
}

type DealReceipts struct {
	// 証憑ID
	ID int32 `json:"id"`
	// ステータス(unconfirmed:確認待ち、confirmed:確認済み、deleted:削除済み、ignored:無視)
	Status string `json:"status"`
	// メモ
	Description *string `json:"description,omitempty"`
	// MIMEタイプ
	MimeType string `json:"mime_type"`
	// 発生日
	IssueDate *string `json:"issue_date,omitempty"`
	// アップロード元種別
	Origin string `json:"origin"`
	// 作成日時（ISO8601形式）
	CreatedAt string `json:"created_at"`
	// ファイルのダウンロードURL（freeeにログインした状態でのみ閲覧可能です。） <br> <br> file_srcは廃止予定の属性になります。<br> file_srcに替わり、証憑ファイルのダウンロード APIをご利用ください。<br> 証憑ファイルのダウンロードAPIを利用することで、以下のようになります。 <ul>   <li>アプリケーション利用者はfreee APIアプリケーションにログインしていれば、証憑ダウンロード毎にfreeeに改めてログインすることなくファイルが参照できるようになります。</li> </ul>
	FileSrc string   `json:"file_src"`
	User    DealUser `json:"user"`
}

type DealUser struct {
	// ユーザーID
	ID int32 `json:"id"`
	// メールアドレス
	Email string `json:"email"`
	// 表示名
	DisplayName *string `json:"display_name,omitempty"`
}

type DealCreateParams struct {
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 収支区分 (収入: income, 支出: expense)
	Type string `json:"type"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 支払期日(yyyy-mm-dd)
	DueDate *string `json:"due_date,omitempty"`
	// 取引先ID
	PartnerID *int32 `json:"partner_id,omitempty"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 管理番号
	RefNumber *string                   `json:"ref_number,omitempty"`
	Details   []DealCreateParamsDetails `json:"details"`
	// 支払行一覧（配列）：未指定の場合、未決済の取引を作成します。
	Payments *[]DealCreateParamsPayments `json:"payments,omitempty"`
	// 証憑ファイルID（配列）
	ReceiptIDs *[]int32 `json:"receipt_ids,omitempty"`
}

type DealCreateParamsDetails struct {
	// 税区分コード
	TaxCode int32 `json:"tax_code"`
	// 勘定科目ID
	AccountItemID int32 `json:"account_item_id"`
	// 取引金額（税込で指定してください）
	Amount int32 `json:"amount"`
	// 品目ID
	ItemID *int32 `json:"item_id,omitempty"`
	// 部門ID
	SectionID *int32 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs *[]int32 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID *int32 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID *int32 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID *int32 `json:"segment_3_tag_id,omitempty"`
	// 備考
	Description *string `json:"description,omitempty"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int32 `json:"vat,omitempty"`
}

type DealCreateParamsPayments struct {
	// 支払金額：payments指定時は必須
	Amount int32 `json:"amount"`
	// 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID）：payments指定時は必須
	FromWalletableID int32 `json:"from_walletable_id"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item)：payments指定時は必須
	FromWalletableType string `json:"from_walletable_type"`
	// 支払日：payments指定時は必須
	Date string `json:"date"`
}

type DealUpdateParams struct {
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 収支区分 (収入: income, 支出: expense)
	Type string `json:"type"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 支払期日(yyyy-mm-dd)
	DueDate *string `json:"due_date,omitempty"`
	// 取引先ID
	PartnerID *int32 `json:"partner_id,omitempty"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 管理番号
	RefNumber *string                   `json:"ref_number,omitempty"`
	Details   []DealUpdateParamsDetails `json:"details"`
	// 証憑ファイルID（配列）
	ReceiptIDs []int32 `json:"receipt_ids,omitempty"`
}

type DealUpdateParamsDetails struct {
	// 取引行ID: 既存取引行を更新する場合に指定します。IDを指定しない取引行は、新規行として扱われ追加されます。また、detailsに含まれない既存の取引行は削除されます。更新後も残したい行は、必ず取引行IDを指定してdetailsに含めてください。
	ID *uint64 `json:"id,omitempty"`
	// 税区分コード
	TaxCode int32 `json:"tax_code"`
	// 勘定科目ID
	AccountItemID int32 `json:"account_item_id"`
	// 取引金額（税込で指定してください）
	Amount int32 `json:"amount"`
	// 品目ID
	ItemID *int32 `json:"item_id,omitempty"`
	// 部門ID
	SectionID *int32 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs *[]int32 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID *int32 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID *int32 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID *int32 `json:"segment_3_tag_id,omitempty"`
	// 備考
	Description *string `json:"description,omitempty"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int32 `json:"vat,omitempty"`
}

func (c *Client) GetDeal(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, dealID uint64, opts GetDealOpts,
) (*Deal, *oauth2.Token, error) {
	var result DealResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathDeals, fmt.Sprint(dealID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Deal, oauth2Token, nil
}

func (c *Client) CreateDeal(
	ctx context.Context, oauth2Token *oauth2.Token,
	params DealCreateParams,
) (*Deal, *oauth2.Token, error) {
	var result DealResponse
	oauth2Token, err := c.call(ctx, APIPathDeals, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Deal, oauth2Token, nil
}

func (c *Client) UpdateDeal(
	ctx context.Context, oauth2Token *oauth2.Token,
	dealID uint64, params DealUpdateParams,
) (*Deal, *oauth2.Token, error) {
	var result DealResponse
	oauth2Token, err := c.call(ctx, path.Join(APIPathDeals, fmt.Sprint(dealID)), http.MethodPut, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Deal, oauth2Token, nil
}

func (c *Client) DestroyDeal(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, dealID uint64,
) (*oauth2.Token, error) {
	v, err := query.Values(nil)
	if err != nil {
		return oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathDeals, fmt.Sprint(dealID)), http.MethodDelete, oauth2Token, v, nil, nil)
	if err != nil {
		return oauth2Token, err
	}

	return oauth2Token, nil
}
