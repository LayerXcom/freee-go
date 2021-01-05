package freee

import (
	"context"
	"golang.org/x/oauth2"
	"strconv"
)

const (
	APIPathReceipts = "receipts"
)

type CreateReceiptParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// メモ (255文字以内)
	Description string `json:"description,omitempty"`
	// 取引日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 証憑ファイル
	Receipt []byte `json:"receipt"`
}

// NewReceiptUpdateParams instantiates a new ReceiptUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReceiptParams(companyID int32, description string, issueDate string, receipt []byte) *CreateReceiptParams {
	this := CreateReceiptParams{}
	this.CompanyID = companyID
	this.IssueDate = issueDate
	this.Description = description
	this.Receipt = receipt
	return &this
}

type ReceiptResponse struct {
	Receipt Receipt `json:"receipt"`
}

type Receipt struct {
	// 証憑ID
	Id int32 `json:"id"`
	// ステータス(unconfirmed:確認待ち、confirmed:確認済み、deleted:削除済み、ignored:無視)
	Status string `json:"status"`
	// メモ
	Description string `json:"description,omitempty"`
	// MIMEタイプ
	MimeType string `json:"mime_type"`
	// 発生日
	IssueDate string `json:"issue_date,omitempty"`
	// アップロード元種別
	Origin string `json:"origin"`
	// 作成日時（ISO8601形式）
	CreatedAt string `json:"created_at"`
	// ファイルのダウンロードURL（freeeにログインした状態でのみ閲覧可能です。） <br> <br> file_srcは廃止予定の属性になります。<br> file_srcに替わり、証憑ファイルのダウンロード APIをご利用ください。<br> 証憑ファイルのダウンロードAPIを利用することで、以下のようになります。 <ul>   <li>アプリケーション利用者はfreee APIアプリケーションにログインしていれば、証憑ダウンロード毎にfreeeに改めてログインすることなくファイルが参照できるようになります。</li> </ul>
	FileSrc string             `json:"file_src"`
	User    UserCreatedReceipt `json:"user"`
}

type UserCreatedReceipt struct {
	// ユーザーID
	Id int32 `json:"id"`
	// メールアドレス
	Email string `json:"email"`
	// 表示名
	DisplayName string `json:"display_name,omitempty"`
}

func (c *Client) CreateReceipt(
	ctx context.Context, oauth2Token *oauth2.Token,
	params CreateReceiptParams,
	receiptName string,
) (*ReceiptResponse, *oauth2.Token, error) {
	request := map[string]string{}
	request["company_id"] = strconv.Itoa(int(params.CompanyID))
	request["description"] = params.Description
	request["issue_date"] = params.IssueDate
	var result ReceiptResponse
	oauth2Token, err := c.upload(ctx, APIPathReceipts, oauth2Token, nil, request, receiptName, params.Receipt, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}

//type ReceiptCreateResponse struct {
//	Deal DealCreateResponseDeal `json:"deal"`
//}
//
//// ReceiptCreateResponseDeal struct for DealCreateResponseDeal
//type ReceiptCreateResponseDeal struct {
//	// 取引ID
//	ID int32 `json:"id"`
//	// 事業所ID
//	CompanyID int32 `json:"company_id"`
//	// 発生日 (yyyy-mm-dd)
//	IssueDate string `json:"issue_date"`
//	// 支払期日 (yyyy-mm-dd)
//	DueDate string `json:"due_date,omitempty"`
//	// 金額
//	Amount int32 `json:"amount"`
//	// 支払金額
//	DueAmount int32 `json:"due_amount,omitempty"`
//	// 収支区分 (収入: income, 支出: expense)
//	Type string `json:"type,omitempty"`
//	// 取引先ID
//	PartnerID int32 `json:"partner_id"`
//	// 取引先コード
//	PartnerCode string `json:"partner_code,omitempty"`
//	// 管理番号
//	RefNumber string `json:"ref_number,omitempty"`
//	// 決済状況 (未決済: unsettled, 完了: settled)
//	Status string `json:"status"`
//	// 取引の明細行
//	Details *[]DealCreateResponseDealDetails `json:"details,omitempty"`
//	// 取引の支払行
//	Payments *[]DealCreateResponseDealPayments `json:"payments,omitempty"`
//}
//
//// DealCreateResponseDealDetails struct for DealCreateResponseDealDetails
//type DealCreateResponseDealDetails struct {
//	// 取引行ID
//	ID int32 `json:"id"`
//	// 勘定科目ID
//	AccountItemID int32 `json:"account_item_id"`
//	// 税区分コード
//	TaxCode int32 `json:"tax_code"`
//	// 品目ID
//	ItemID int32 `json:"item_id,omitempty"`
//	// 部門ID
//	SectionID int32 `json:"section_id,omitempty"`
//	// メモタグID
//	TagIDs []int32 `json:"tag_ids,omitempty"`
//	// セグメント１ID
//	Segment1TagID int32 `json:"segment_1_tag_id,omitempty"`
//	// セグメント２ID
//	Segment2TagID int32 `json:"segment_2_tag_id,omitempty"`
//	// セグメント３ID
//	Segment3TagID int32 `json:"segment_3_tag_id,omitempty"`
//	// 取引金額
//	Amount int32 `json:"amount"`
//	// 消費税額
//	Vat int32 `json:"vat"`
//	// 備考
//	Description string `json:"description,omitempty"`
//	// 貸借（貸方: credit, 借方: debit）
//	EntrySide string `json:"entry_side"`
//}
//
//// DealCreateResponseDealPayments struct for DealCreateResponseDealPayments
//type DealCreateResponseDealPayments struct {
//	// 取引行ID
//	ID int32 `json:"id"`
//	// 支払日
//	Date string `json:"date"`
//	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item)
//	FromWalletableType string `json:"from_walletable_type,omitempty"`
//	// 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID）
//	FromWalletableID int32 `json:"from_walletable_id,omitempty"`
//	// 支払金額
//	Amount int32 `json:"amount"`
//}
//
//// DealCreateParams struct for DealCreateParams
//type DealCreateParams struct {
//	// 発生日 (yyyy-mm-dd)
//	IssueDate string `json:"issue_date"`
//	// 収支区分 (収入: income, 支出: expense)
//	Type string `json:"type"`
//	// 事業所ID
//	CompanyID int32 `json:"company_id"`
//	// 支払期日(yyyy-mm-dd)
//	DueDate string `json:"due_date,omitempty"`
//	// 取引先ID
//	PartnerID int32 `json:"partner_id,omitempty"`
//	// 取引先コード
//	PartnerCode string `json:"partner_code,omitempty"`
//	// 管理番号
//	RefNumber string                     `json:"ref_number,omitempty"`
//	Details   *[]DealCreateParamsDetails `json:"details"`
//	// 支払行一覧（配列）：未指定の場合、未決済の取引を作成します。
//	Payments *[]DealCreateParamsPayments `json:"payments,omitempty"`
//	// 証憑ファイルID（配列）
//	ReceiptIDs []int32 `json:"receipt_ids,omitempty"`
//}
//
//// DealCreateParamsDetails struct for DealCreateParamsDetails
//type DealCreateParamsDetails struct {
//	// 税区分コード
//	TaxCode int32 `json:"tax_code"`
//	// 勘定科目ID
//	AccountItemID int32 `json:"account_item_id"`
//	// 取引金額（税込で指定してください）
//	Amount int32 `json:"amount"`
//	// 品目ID
//	ItemID int32 `json:"item_id,omitempty"`
//	// 部門ID
//	SectionID int32 `json:"section_id,omitempty"`
//	// メモタグID
//	TagIDs []int32 `json:"tag_ids,omitempty"`
//	// セグメント１ID
//	Segment1TagID int32 `json:"segment_1_tag_id,omitempty"`
//	// セグメント２ID
//	Segment2TagID int32 `json:"segment_2_tag_id,omitempty"`
//	// セグメント３ID
//	Segment3TagID int32 `json:"segment_3_tag_id,omitempty"`
//	// 備考
//	Description string `json:"description,omitempty"`
//	// 消費税額（指定しない場合は自動で計算されます）
//	Vat int32 `json:"vat,omitempty"`
//}
//
//// DealCreateParamsPayments struct for DealCreateParamsPayments
//type DealCreateParamsPayments struct {
//	// 支払金額：payments指定時は必須
//	Amount int32 `json:"amount"`
//	// 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID）：payments指定時は必須
//	FromWalletableID int32 `json:"from_walletable_id"`
//	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item)：payments指定時は必須
//	FromWalletableType string `json:"from_walletable_type"`
//	// 支払日：payments指定時は必須
//	Date string `json:"date"`
//}
//
