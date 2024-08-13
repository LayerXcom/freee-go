package freee

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathInvoices = "invoices"

	InvoiceStatusDraft       = "draft"       // 下書き
	InvoiceStatusApplying    = "applying"    // 申請中
	InvoiceStatusRemanded    = "remanded"    // 差し戻し
	InvoiceStatusRejected    = "rejected"    // 却下
	InvoiceStatusApproved    = "approved"    // 承認済み
	InvoiceStatusUnsubmitted = "unsubmitted" // 送付待ち
	InvoiceStatusSubmitted   = "submitted"   // 送付済み

	PaymentStatusNone      = ""
	PaymentStatusUnsettled = "unsettled" // 入金待ち
	PaymentStatusSettled   = "settled"   // 入金済み

	PostingStatusNone              = ""
	PostingStatusUnrequested       = "unrequested"        // リクエスト前
	PostingStatusPreviewRegistered = "preview_registered" // プレビュー登録
	PostingStatusPreviewFailed     = "preview_failed"     // プレビュー登録失敗
	PostingStatusOrdered           = "ordered"            // 注文中
	PostingStatusOrderFailed       = "order_failed"       // 注文失敗
	PostingStatusPrinting          = "printing"           // 印刷中
	PostingStatusCanceled          = "canceled"           // キャンセル
	PostingStatusPosted            = "posted"             // 投函済み

	PaymentTypeTransfer    = "transfer"     // 振込
	PaymentTypeDirectDebit = "direct_debit" // 引き落とし

	InvoiceLayoutDefaultClassic                = "default_classic"                  // レイアウト１/クラシック (デフォルト)
	InvoiceLayoutStandardClassic               = "standard_classic"                 // レイアウト２/クラシック
	InvoiceLayoutEnvelopeClassic               = "envelope_classic"                 // 封筒１/クラシック
	InvoiceLayoutCarriedForwardStandardClassic = "carried_forward_standard_classic" // レイアウト３（繰越金額欄あり）/クラシック
	InvoiceLayoutCarriedForwardEnvelopeClassic = "carried_forward_envelope_classic" // 封筒２（繰越金額欄あり）/クラシック
	InvoiceLayoutDefaultModern                 = "default_modern"                   // レイアウト１/モダン
	InvoiceLayoutStandardModern                = "standard_modern"                  // レイアウト２/モダン
	InvoiceLayoutEnvelopeModern                = "envelope_modern"                  // 封筒/モダン

	TaxEntryMethodInclusive = "inclusive" // 内税
	TaxEntryMethodExclusive = "exclusive" // 外税

	InvoiceContentTypeNormal   = "normal"
	InvoiceContentTypeDiscount = "discount"
	InvoiceContentTypeText     = "text"

	UseVirtualTransferAccountNotUse = "not_use" // 利用しない
	UseVirtualTransferAccountUse    = "use"     // 利用する
)

type InvoicesResponse struct {
	Invoices []Invoice `json:"invoices"`
}

type InvoiceResponse struct {
	Invoice Invoice `json:"invoice"`
}

type GetInvoiceOpts struct {
	// 事業所ID
	CompanyID int64 `url:"company_id"`
	// 取引先IDで絞込
	PartnerID int64 `url:"partner_id,omitempty"`
	// 取引先コードで絞込
	PartnerCode string `url:"partner_code,omitempty"`
	// 請求日の開始日(yyyy-mm-dd)
	StartIssueDate string `url:"start_issue_date,omitempty"`
	// 請求日の終了日(yyyy-mm-dd)
	EndIssueDate string `url:"end_issue_date,omitempty"`
	// 期日の開始日(yyyy-mm-dd)
	StartDueDate string `url:"start_due_date,omitempty"`
	// 期日の終了日(yyyy-mm-dd)
	EndDueDate string `url:"end_due_date,omitempty"`
	// 請求書番号
	InvoiceNumber string `url:"invoice_number,omitempty"`
	// 概要
	Description string `url:"description,omitempty"`
	// 請求書ステータス (draft: 下書き, applying: 申請中, remanded: 差し戻し, rejected: 却下, approved: 承認済み, unsubmitted: 送付待ち, submitted: 送付済み)
	InvoiceStatus string `url:"invoice_status,omitempty"`
	// 入金ステータス (unsettled: 入金待ち, settled: 入金済み)
	PaymentStatus string `url:"payment_status,omitempty"`
	// 取得レコードのオフセット (デフォルト: 0)
	Offset int64 `url:"offset,omitempty"`
	// 取得レコードの件数 (デフォルト: 20, 最大: 100)
	Limit int64 `url:"limit,omitempty"`
}

type Invoice struct {
	// 請求書ID
	ID int64 `json:"id"`
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 請求日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 取引先ID
	PartnerID int64 `json:"partner_id"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 請求書番号
	InvoiceNumber string `json:"invoice_number"`
	// タイトル
	Title *string `json:"title,omitempty"`
	// 期日 (yyyy-mm-dd)
	DueDate *string `json:"due_date,omitempty"`
	// 合計金額
	TotalAmount int64 `json:"total_amount"`
	// 合計税額
	TotalVat int64 `json:"total_vat"`
	// 小計
	SubTotal int64 `json:"sub_total"`
	// 売上計上日
	BookingDate *string `json:"booking_date,omitempty"`
	// 概要
	Description *string `json:"description,omitempty"`
	// 請求書ステータス (draft: 下書き, applying: 申請中, remanded: 差し戻し, rejected: 却下, approved: 承認済み, submitted: 送付済み, unsubmitted: 請求書の承認フローが無効の場合のみ、unsubmitted（送付待ち）の値をとります)
	InvoiceStatus string `json:"invoice_status"`
	// 入金ステータス (unsettled: 入金待ち, settled: 入金済み)
	PaymentStatus string `json:"payment_status"`
	// 入金日 (yyyy-mm-dd)
	PaymentDate *string `json:"payment_date,omitempty"`
	// Web共有日時(最新)
	WebPublishedAt *string `json:"web_published_at,omitempty"`
	// Web共有ダウンロード日時(最新)
	WebDownloadedAt *string `json:"web_downloaded_at,omitempty"`
	// Web共有取引先確認日時(最新)
	WebConfirmedAt *string `json:"web_confirmed_at,omitempty"`
	// メール送信日時(最新)
	MailSentAt *string `json:"mail_sent_at,omitempty"`
	// 郵送ステータス(unrequested: リクエスト前, preview_registered: プレビュー登録, preview_failed: プレビュー登録失敗, ordered: 注文中, order_failed: 注文失敗, printing: 印刷中, canceled: キャンセル, posted: 投函済み)
	PostingStatus string `json:"posting_status"`
	// 取引先名
	PartnerName *string `json:"partner_name,omitempty"`
	// 請求書に表示する取引先名
	PartnerDisplayName *string `json:"partner_display_name,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	PartnerTitle *string `json:"partner_title,omitempty"`
	// 取引先郵便番号
	PartnerZipcode *string `json:"partner_zipcode,omitempty"`
	// 取引先都道府県コード（-1: 設定しない、0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	PartnerPrefectureCode *int64 `json:"partner_prefecture_code,omitempty"`
	// 取引先都道府県
	PartnerPrefectureName *string `json:"partner_prefecture_name,omitempty"`
	// 取引先市区町村・番地
	PartnerAddress1 *string `json:"partner_address1,omitempty"`
	// 取引先建物名・部屋番号など
	PartnerAddress2 *string `json:"partner_address2,omitempty"`
	// 取引先担当者名
	PartnerContactInfo *string `json:"partner_contact_info,omitempty"`
	// 事業所名
	CompanyName string `json:"company_name"`
	// 郵便番号
	CompanyZipcode *string `json:"company_zipcode,omitempty"`
	// 都道府県コード（-1: 設定しない、0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	CompanyPrefectureCode *int64 `json:"company_prefecture_code,omitempty"`
	// 都道府県
	CompanyPrefectureName *string `json:"company_prefecture_name,omitempty"`
	// 市区町村・番地
	CompanyAddress1 *string `json:"company_address1,omitempty"`
	// 建物名・部屋番号など
	CompanyAddress2 *string `json:"company_address2,omitempty"`
	// 事業所担当者名
	CompanyContactInfo *string `json:"company_contact_info,omitempty"`
	// 支払方法 (振込: transfer, 引き落とし: direct_debit)
	PaymentType string `json:"payment_type"`
	// 支払口座
	PaymentBankInfo *string `json:"payment_bank_info,omitempty"`
	// メッセージ
	Message *string `json:"message,omitempty"`
	// 備考
	Notes *string `json:"notes,omitempty"`
	// 請求書レイアウト
	// - default_classic - レイアウト１/クラシック (デフォルト)
	// - standard_classic - レイアウト２/クラシック
	// - envelope_classic - 封筒１/クラシック
	// - carried_forward_standard_classic - レイアウト３（繰越金額欄あり）/クラシック
	// - carried_forward_envelope_classic - 封筒２（繰越金額欄あり）/クラシック
	// - default_modern - レイアウト１/モダン
	// - standard_modern - レイアウト２/モダン
	// - envelope_modern - 封筒/モダン
	InvoiceLayout string `json:"invoice_layout"`
	// 請求書の消費税計算方法(inclusive: 内税, exclusive: 外税)
	TaxEntryMethod string `json:"tax_entry_method"`
	// 取引ID (invoice_statusがsubmitted, unsubmittedの時IDが表示されます)
	DealID *int64 `json:"deal_id,omitempty"`
	// 請求内容
	InvoiceContents       []InvoiceContent `json:"invoice_contents"`
	TotalAmountPerVatRate struct {
		// 税率5%の税込み金額合計
		Vat5 int64 `json:"vat_5"`
		// 税率8%の税込み金額合計
		Vat8 int64 `json:"vat_8"`
		// 軽減税率8%の税込み金額合計
		ReducedVat8 int64 `json:"reduced_vat_8"`
		// 税率10%の税込み金額合計
		Vat10 int64 `json:"vat_10"`
	} `json:"total_amount_per_vat_rate"`
	// 関連する見積書ID(配列)
	// 下記で作成したものが該当します。
	//
	// [見積書・納品書を納品書・請求書に変換する](https://support.freee.co.jp/hc/ja/articles/203318410#1-2)
	// [複数の見積書・納品書から合算請求書を作成する](https://support.freee.co.jp/hc/ja/articles/209076226)
	RelatedQuotationIDs []int64 `json:"related_quotation_ids,omitempty"`
}

type InvoiceContent struct {
	// 請求内容ID
	ID int64 `json:"id"`
	// 順序
	Order *int64 `json:"order,omitempty"`
	// 行の種類
	Type string `json:"type"`
	// 数量
	Qty json.Number `json:"qty"`
	// 単位
	Unit *string `json:"unit,omitempty"`
	// 単価
	UnitPrice json.Number `json:"unit_price"`
	// 内税/外税の判別とamountの税込み、税抜きについて
	//
	// - tax_entry_methodがexclusive (外税)の場合
	//   - amount: 消費税抜きの金額
	//   - vat: 消費税の金額
	// - tax_entry_methodがinclusive (内税)の場合
	//   - amount: 消費税込みの金額
	//   - vat: 消費税の金額
	Amount int64 `json:"amount"`
	// 消費税額
	Vat int64 `json:"vat"`
	// 軽減税率税区分（true: 対象、false: 対象外）
	ReducedVat bool `json:"reduced_vat"`
	// 備考
	Description *string `json:"description,omitempty"`
	// 勘定科目ID
	AccountItemID int64 `json:"account_item_id"`
	// 勘定科目名
	AccountItemName string `json:"account_item_name"`
	// 税区分コード
	TaxCode int64 `json:"tax_code"`
	// 品目ID
	ItemID *int64 `json:"item_id,omitempty"`
	// 品目
	ItemName *string `json:"item_name,omitempty"`
	// 部門ID
	SectionID *int64 `json:"section_id,omitempty"`
	// 部門
	SectionName *string `json:"section_name,omitempty"`
	// メモタグID
	TagIDs []int64 `json:"tag_ids"`
	// メモタグ
	TagNames []string `json:"tag_names"`
	// セグメント１ID
	Segment1TagID *int64 `json:"segment_1_tag_id,omitempty"`
	// セグメント１
	Segment1TagName *string `json:"segment_1_tag_name,omitempty"`
	// セグメント２ID
	Segment2TagID *int64 `json:"segment_2_tag_id,omitempty"`
	// セグメント２
	Segment2TagName *string `json:"segment_2_tag_name,omitempty"`
	// セグメント３ID
	Segment3TagID *int64 `json:"segment_3_tag_id,omitempty"`
	// セグメント３
	Segment3TagName *string `json:"segment_3_tag_name,omitempty"`
}

type InvoiceCreateParams struct {
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 請求日 (yyyy-mm-dd)
	IssueDate *string `json:"issue_date,omitempty"`
	// 取引先ID
	PartnerID *int64 `json:"partner_id,omitempty"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 請求書番号 (デフォルト: 自動採番されます)
	InvoiceNumber *string `json:"invoice_number,omitempty"`
	// タイトル (デフォルト: 請求書)
	Title *string `json:"title,omitempty"`
	// 期日 (yyyy-mm-dd)
	DueDate *string `json:"due_date,omitempty"`
	// 売上計上日
	BookingDate *string `json:"booking_date,omitempty"`
	// 概要
	Description *string `json:"description,omitempty"`
	// 請求書ステータス
	//
	// - draft: 下書き (デフォルト)
	// - (廃止予定) issue: 発行 (送付待ち (unsubmitted) と同じです。)
	// - unsubmitted: 送付待ち
	// - submitted: 送付済み
	//
	// issue, unsubmitted, submitted は請求書承認ワークフローを利用している場合は指定できません。
	InvoiceStatus *string `json:"invoice_status,omitempty"`
	// 請求書に表示する取引先名
	PartnerDisplayName string `json:"partner_display_name"`
	// 敬称（御中、様、(空白)の3つから選択
	PartnerTitle string `json:"partner_title"`
	// 取引先担当者名
	PartnerContactInfo *string `json:"partner_contact_info,omitempty"`
	// 取引先郵便番号 (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます)
	PartnerZipcode *string `json:"partner_zipcode,omitempty"`
	// 取引先都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄) (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます)
	PartnerPrefectureCode *int64 `json:"partner_prefecture_code,omitempty"`
	// 取引先市区町村・番地 (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます)
	PartnerAddress1 *string `json:"partner_address1,omitempty"`
	// 取引先建物名・部屋番号など (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます)
	PartnerAddress2 *string `json:"partner_address2,omitempty"`
	// 事業所名 (デフォルトは事業所設定情報が補完されます)
	CompanyName *string `json:"company_name,omitempty"`
	// 郵便番号 (デフォルトは事業所設定情報が補完されます)
	CompanyZipcode *string `json:"company_zipcode,omitempty"`
	// 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄) (デフォルトは事業所設定情報が補完されます)
	CompanyPrefectureCode *int64 `json:"company_prefecture_code,omitempty"`
	// 市区町村・番地 (デフォルトは事業所設定情報が補完されます)
	CompanyAddress1 *string `json:"company_address1,omitempty"`
	// 建物名・部屋番号など (デフォルトは事業所設定情報が補完されます)
	CompanyAddress2 *string `json:"company_address2,omitempty"`
	// 事業所担当者名 (デフォルトは請求書テンプレート情報が補完されます)
	CompanyContactInfo *string `json:"company_contact_info,omitempty"`
	// 支払方法 (振込: transfer, 引き落とし: direct_debit)
	PaymentType *string `json:"payment_type,omitempty"`
	// 支払口座
	PaymentBankInfo *string `json:"payment_bank_info,omitempty"`
	// 振込専用口座の利用(利用しない: not_use(デフォルト), 利用する: use)
	UseVirtualTransferAccount *string `json:"use_virtual_transfer_account,omitempty"`
	// メッセージ (デフォルト: 下記の通りご請求申し上げます。)
	Message *string `json:"message,omitempty"`
	// 備考
	Notes *string `json:"notes,omitempty"`
	// 請求書レイアウト
	// - default_classic - レイアウト１/クラシック (デフォルト)
	// - standard_classic - レイアウト２/クラシック
	// - envelope_classic - 封筒１/クラシック
	// - carried_forward_standard_classic - レイアウト３（繰越金額欄あり）/クラシック
	// - carried_forward_envelope_classic - 封筒２（繰越金額欄あり）/クラシック
	// - default_modern - レイアウト１/モダン
	// - standard_modern - レイアウト２/モダン
	// - envelope_modern - 封筒/モダン
	InvoiceLayout *string `json:"invoice_layout,omitempty"`
	// 請求書の消費税計算方法(inclusive: 内税表示, exclusive: 外税表示 (デフォルト))
	TaxEntryMethod *string `json:"tax_entry_method,omitempty"`
	// 請求内容
	InvoiceContents []InvoiceCreateParamsInvoiceContent `json:"invoice_contents"`
}

type InvoiceCreateParamsInvoiceContent struct {
	// 順序
	Order int64 `json:"order"`
	// 行の種類
	//
	// - normal、discountを指定する場合、account_item_id,tax_codeとunit_priceが必須となります。
	// - normalを指定した場合、qtyが必須となります。
	Type string `json:"type"`
	// 数量
	Qty json.Number `json:"qty"`
	// 単位
	Unit *string `json:"unit,omitempty"`
	// 単価 (tax_entry_method: inclusiveの場合は税込価格、tax_entry_method: exclusiveの場合は税抜価格となります)
	UnitPrice json.Number `json:"unit_price"`
	// 消費税額
	Vat *int64 `json:"vat,omitempty"`
	// 備考
	Description *string `json:"description,omitempty"`
	// 勘定科目ID
	AccountItemID int64 `json:"account_item_id"`
	// 税区分コード
	TaxCode int64 `json:"tax_code"`
	// 品目ID
	ItemID *int64 `json:"item_id,omitempty"`
	// 部門ID
	SectionID *int64 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs *[]int64 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID *int64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID *int64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID *int64 `json:"segment_3_tag_id,omitempty"`
}

type InvoiceUpdateParams struct {
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 請求日 (yyyy-mm-dd)
	IssueDate *string `json:"issue_date,omitempty"`
	// 取引先ID
	PartnerID *int64 `json:"partner_id,omitempty"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 請求書番号 (デフォルト: 自動採番されます)
	InvoiceNumber *string `json:"invoice_number,omitempty"`
	// タイトル (デフォルト: 請求書)
	Title *string `json:"title,omitempty"`
	// 期日 (yyyy-mm-dd)
	DueDate *string `json:"due_date,omitempty"`
	// 売上計上日
	BookingDate *string `json:"booking_date,omitempty"`
	// 概要
	Description *string `json:"description,omitempty"`
	// 請求書ステータス
	//
	// - draft: 下書き (デフォルト)
	// - (廃止予定) issue: 発行 (送付待ち (unsubmitted) と同じです。)
	// - unsubmitted: 送付待ち
	// - submitted: 送付済み
	//
	// issue, unsubmitted, submitted は請求書承認ワークフローを利用している場合は指定できません。
	InvoiceStatus *string `json:"invoice_status,omitempty"`
	// 請求書に表示する取引先名
	PartnerDisplayName string `json:"partner_display_name"`
	// 敬称（御中、様、(空白)の3つから選択）
	PartnerTitle string `json:"partner_title"`
	// 取引先担当者名
	PartnerContactInfo *string `json:"partner_contact_info,omitempty"`
	// 取引先郵便番号 (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます)
	PartnerZipcode *string `json:"partner_zipcode,omitempty"`
	// 取引先都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄) (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます)
	PartnerPrefectureCode *int64 `json:"partner_prefecture_code,omitempty"`
	// 取引先市区町村・番地 (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます)
	PartnerAddress1 *string `json:"partner_address1,omitempty"`
	// 取引先建物名・部屋番号など (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます)
	PartnerAddress2 *string `json:"partner_address2,omitempty"`
	// 事業所名 (デフォルトは事業所設定情報が補完されます)
	CompanyName *string `json:"company_name,omitempty"`
	// 郵便番号 (デフォルトは事業所設定情報が補完されます)
	CompanyZipcode *string `json:"company_zipcode,omitempty"`
	// 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄) (デフォルトは事業所設定情報が補完されます)
	CompanyPrefectureCode *int64 `json:"company_prefecture_code,omitempty"`
	// 市区町村・番地 (デフォルトは事業所設定情報が補完されます)
	CompanyAddress1 *string `json:"company_address1,omitempty"`
	// 建物名・部屋番号など (デフォルトは事業所設定情報が補完されます)
	CompanyAddress2 *string `json:"company_address2,omitempty"`
	// 事業所担当者名 (デフォルトは請求書テンプレート情報が補完されます)
	CompanyContactInfo *string `json:"company_contact_info,omitempty"`
	// 支払方法 (振込: transfer, 引き落とし: direct_debit)
	PaymentType *string `json:"payment_type,omitempty"`
	// 支払口座
	PaymentBankInfo *string `json:"payment_bank_info,omitempty"`
	// 振込専用口座の利用(利用しない: not_use(デフォルト), 利用する: use)
	UseVirtualTransferAccount *string `json:"use_virtual_transfer_account,omitempty"`
	// メッセージ (デフォルト: 下記の通りご請求申し上げます。)
	Message *string `json:"message,omitempty"`
	// 備考
	Notes *string `json:"notes,omitempty"`
	// 請求書レイアウト
	// - default_classic - レイアウト１/クラシック (デフォルト)
	// - standard_classic - レイアウト２/クラシック
	// - envelope_classic - 封筒１/クラシック
	// - carried_forward_standard_classic - レイアウト３（繰越金額欄あり）/クラシック
	// - carried_forward_envelope_classic - 封筒２（繰越金額欄あり）/クラシック
	// - default_modern - レイアウト１/モダン
	// - standard_modern - レイアウト２/モダン
	// - envelope_modern - 封筒/モダン
	InvoiceLayout *string `json:"invoice_layout,omitempty"`
	// 請求書の消費税計算方法(inclusive: 内税表示, exclusive: 外税表示 (デフォルト))
	TaxEntryMethod *string `json:"tax_entry_method,omitempty"`
	// 請求内容
	InvoiceContents []InvoiceUpdateParamsInvoiceContent `json:"invoice_contents"`
}

type InvoiceUpdateParamsInvoiceContent struct {
	// 請求内容ID
	ID *int64 `json:"id,omitempty"`
	// 順序
	Order int64 `json:"order"`
	// 行の種類
	//
	// - normal、discountを指定する場合、account_item_id,tax_codeとunit_priceが必須となります。
	// - normalを指定した場合、qtyが必須となります。
	Type string `json:"type"`
	// 数量
	Qty json.Number `json:"qty"`
	// 単位
	Unit *string `json:"unit,omitempty"`
	// 単価 (tax_entry_method: inclusiveの場合は税込価格、tax_entry_method: exclusiveの場合は税抜価格となります)
	UnitPrice json.Number `json:"unit_price"`
	// 消費税額
	Vat *int64 `json:"vat,omitempty"`
	// 備考
	Description *string `json:"description,omitempty"`
	// 勘定科目ID
	AccountItemID int64 `json:"account_item_id"`
	// 税区分コード
	TaxCode int64 `json:"tax_code"`
	// 品目ID
	ItemID *int64 `json:"item_id,omitempty"`
	// 部門ID
	SectionID *int64 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs *[]int64 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID *int64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID *int64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID *int64 `json:"segment_3_tag_id,omitempty"`
}

func (c *Client) GetInvoices(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, opts GetInvoiceOpts,
) (*InvoicesResponse, *oauth2.Token, error) {
	var result InvoicesResponse
	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathInvoices, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}

func (c *Client) GetInvoice(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, invoiceID int64, opts GetInvoiceOpts,
) (*Invoice, *oauth2.Token, error) {
	var result InvoiceResponse
	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathInvoices, fmt.Sprint(invoiceID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Invoice, oauth2Token, nil
}

func (c *Client) CreateInvoice(
	ctx context.Context, oauth2Token *oauth2.Token,
	params InvoiceCreateParams,
) (*Invoice, *oauth2.Token, error) {
	var result InvoiceResponse
	oauth2Token, err := c.call(ctx, APIPathInvoices, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Invoice, oauth2Token, nil
}

func (c *Client) UpdateInvoice(
	ctx context.Context, oauth2Token *oauth2.Token,
	invoiceID int64, params InvoiceUpdateParams,
) (*Invoice, *oauth2.Token, error) {
	var result InvoiceResponse
	oauth2Token, err := c.call(ctx, path.Join(APIPathInvoices, fmt.Sprint(invoiceID)), http.MethodPut, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result.Invoice, oauth2Token, nil
}

func (c *Client) DestroyInvoice(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, invoiceID int64,
) (*oauth2.Token, error) {
	v, err := query.Values(nil)
	if err != nil {
		return oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathInvoices, fmt.Sprint(invoiceID)), http.MethodDelete, oauth2Token, v, nil, nil)
	if err != nil {
		return oauth2Token, err
	}
	return oauth2Token, nil
}
