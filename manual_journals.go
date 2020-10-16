package freee

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	APIPathManualJournals = "manual_journals"
)

type CreateManualJournalParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳）
	Adjustment                       bool                              `json:"adjustment,omitempty"`
	CreateManualJournalParamsDetails []CreateManualJournalParamsDetail `json:"details"`
}

type CreateManualJournalParamsDetail struct {
	// 貸借（貸方: credit, 借方: debit）
	EntrySide string `json:"entry_side"`
	// 税区分コード
	TaxCode int32 `json:"tax_code"`
	// 勘定科目ID
	AccountItemID int32 `json:"account_item_id"`
	// 取引金額（税込で指定してください）
	Amount uint64 `json:"amount"`
	// 消費税額（指定しない場合は自動で計算されます）
	Vat int32 `json:"vat"`
	// 取引先ID
	PartnerID int32 `json:"partner_id,omitempty"`
	// 取引先コード
	PartnerCode string `json:"partner_code,omitempty"`
	// 品目ID
	ItemID int32 `json:"item_id,omitempty"`
	// 部門ID
	SectionID int32 `json:"section_id,omitempty"`
	// メモタグID
	TagIds []int32 `json:"tag_ids,omitempty"`
	// セグメント１ID
	Segment1TagID uint64 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID uint64 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID uint64 `json:"segment_3_tag_id,omitempty"`
	// 備考
	Description string `json:"description,omitempty"`
}

func (c *Client) CreateManualJournal(
	ctx context.Context, oauth2Token *oauth2.Token,
	params CreateManualJournalParams,
) (*CreateManualJournalParamsDetail, *oauth2.Token, error) {
	var result CreateManualJournalParamsDetail

	tokenSource, err := c.call(ctx, APIPathManualJournals, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, token, nil
}
