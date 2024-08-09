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
	APIPathManualJournals = "manual_journals"

	ManualJournalEntrySideCredit = "credit"
	ManualJournalEntrySideDebit  = "debit"
)

type ManualJournalsResponse struct {
	ManualJournals []ManualJournal `json:"manual_journals"`
}

type ManualJournalResponse struct {
	ManualJournal ManualJournal `json:"manual_journal"`
}

type ManualJournal struct {
	// 振替伝票ID
	ID int64 `json:"id"`
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳）
	Adjustment bool `json:"adjustment"`
	// 仕訳番号
	TxnNumber *string `json:"txn_number"`
	// 貸借行一覧（配列）: 貸借合わせて100行まで登録できます。
	Details []ManualJournalDetails `json:"details"`
	// 証憑ファイルID（ファイルボックスのファイルID）（配列）
	ReceiptIDs []int64 `json:"receipt_ids"`
}

type ManualJournalDetails struct {
	// 貸借行ID
	ID int64 `json:"id"`
	// 貸借(貸方: credit, 借方: debit)
	EntrySide string `json:"entry_side"`
	// 勘定科目ID
	AccountItemID int64 `json:"account_item_id"`
	// 税区分コード
	TaxCode int64 `json:"tax_code"`
	// 取引先ID
	PartnerID *int64 `json:"partner_id"`
	// 取引先名
	PartnerName *string `json:"partner_name"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 正式名称（255文字以内）
	PartnerLongName *string `json:"partner_long_name"`
	// 品目ID
	ItemID *int64 `json:"item_id"`
	// 品目
	ItemName *string `json:"item_name"`
	// 部門ID
	SectionID *int64 `json:"section_id"`
	// 部門
	SectionName *string  `json:"section_name"`
	TagIDs      []int64  `json:"tag_ids"`
	TagNames    []string `json:"tag_names"`
	// セグメント１ID
	Segment1TagID int64 `json:"segment_1_tag_id,omitempty"`
	// セグメント１ID
	Segment1TagName *string `json:"segment_1_tag_name,omitempty"`
	// セグメント２ID
	Segment2TagID int64 `json:"segment_2_tag_id,omitempty"`
	// セグメント２
	Segment2TagName *string `json:"segment_2_tag_name,omitempty"`
	// セグメント３ID
	Segment3TagID int64 `json:"segment_3_tag_id,omitempty"`
	// セグメント３
	Segment3TagName *string `json:"segment_3_tag_name,omitempty"`
	// 金額（税込で指定してください）
	Amount int64 `json:"amount"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int64 `json:"vat,omitempty"`
	// 備考
	Description string `json:"description"`
}

type CreateManualJournalParams struct {
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 仕訳番号
	TxnNumber string `json:"txn_number,omitempty"`
	// 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳）
	Adjustment                       bool                              `json:"adjustment,omitempty"`
	CreateManualJournalParamsDetails []CreateManualJournalParamsDetail `json:"details"`
	// 証憑ファイルID（ファイルボックスのファイルID）（配列）
	ReceiptIDs []int64 `json:"receipt_ids,omitempty"`
}

type CreateManualJournalParamsDetail struct {
	// 貸借（貸方: credit, 借方: debit）
	EntrySide string `json:"entry_side"`
	// 税区分コード
	TaxCode int64 `json:"tax_code"`
	// 勘定科目ID
	AccountItemID int64 `json:"account_item_id"`
	// 取引金額（税込で指定してください）
	Amount int64 `json:"amount"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int64 `json:"vat,omitempty"`
	// 取引先ID
	PartnerID int64 `json:"partner_id,omitempty"`
	// 取引先コード
	PartnerCode string `json:"partner_code,omitempty"`
	// 品目ID
	ItemID int64 `json:"item_id,omitempty"`
	// 部門ID
	SectionID int64 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs []int64 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID int64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID int64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID int64 `json:"segment_3_tag_id,omitempty"`
	// 備考
	Description string `json:"description,omitempty"`
}

type UpdateManualJournalParams struct {
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳）
	Adjustment bool                               `json:"adjustment,omitempty"`
	Details    []UpdateManualJournalParamsDetails `json:"details"`
	// 証憑ファイルID（ファイルボックスのファイルID）（配列）
	ReceiptIDs []int64 `json:"receipt_ids,omitempty"`
}

// ManualJournalUpdateParamsDetails 貸借行一覧（配列）: 貸借合わせて100行まで登録できます。
type UpdateManualJournalParamsDetails struct {
	// 貸借行ID: 既存貸借行を更新または削除する場合に指定します。IDを指定しない貸借行は、新規行として扱われ追加されます。
	ID int64 `json:"id,omitempty"`
	// 貸借（貸方: credit, 借方: debit）
	EntrySide string `json:"entry_side"`
	// 税区分コード
	TaxCode int64 `json:"tax_code"`
	// 勘定科目ID
	AccountItemID int64 `json:"account_item_id"`
	// 取引金額（税込で指定してください）
	Amount int64 `json:"amount"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat *int64 `json:"vat,omitempty"`
	// 取引先ID
	PartnerID int64 `json:"partner_id,omitempty"`
	// 取引先コード
	PartnerCode string `json:"partner_code,omitempty"`
	// 品目ID
	ItemID int64 `json:"item_id,omitempty"`
	// 部門ID
	SectionID int64 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs []int64 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID int64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID int64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID int64 `json:"segment_3_tag_id,omitempty"`
	// 備考
	Description string `json:"description,omitempty"`
}

type GetManualJournalsOpts struct {
	// 発生日で絞込：開始日(yyyy-mm-dd)
	StartIssueDate string `url:"start_issue_date,omitempty"`
	// 発生日で絞込：終了日(yyyy-mm-dd)
	EndIssueDate string `url:"end_issue_date,omitempty"`
	// 貸借で絞込 (貸方: credit, 借方: debit)
	EntrySide string `url:"entry_side,omitempty"`
	Offset    int64  `url:"offset,omitempty"`
	Limit     int64  `url:"limit,omitempty"`
}

func (c *Client) CreateManualJournal(
	ctx context.Context, oauth2Token *oauth2.Token,
	params CreateManualJournalParams,
) (*ManualJournalResponse, *oauth2.Token, error) {
	var result ManualJournalResponse

	oauth2Token, err := c.call(ctx, APIPathManualJournals, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) UpdateManualJournal(
	ctx context.Context, oauth2Token *oauth2.Token,
	journalID int64, params UpdateManualJournalParams,
) (*ManualJournalResponse, *oauth2.Token, error) {
	var result ManualJournalResponse

	oauth2Token, err := c.call(ctx, path.Join(APIPathManualJournals, fmt.Sprint(journalID)), http.MethodPut, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) DestroyManualJournal(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, journalID int64,
) (*oauth2.Token, error) {
	v, err := query.Values(nil)
	if err != nil {
		return oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathManualJournals, fmt.Sprint(journalID)), http.MethodDelete, oauth2Token, v, nil, nil)
	if err != nil {
		return oauth2Token, err
	}

	return oauth2Token, nil
}

func (c *Client) GetManualJournal(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, journalID int64, opts GetManualJournalsOpts,
) (*ManualJournalResponse, *oauth2.Token, error) {
	var result ManualJournalResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}

	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathManualJournals, fmt.Sprint(journalID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}

func (c *Client) GetManualJournals(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID int64, opts GetManualJournalsOpts,
) (*ManualJournalsResponse, *oauth2.Token, error) {
	var result ManualJournalsResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}

	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathManualJournals), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}
