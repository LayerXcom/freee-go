package freee

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathJournals = "journals"
)

type JournalResponse struct {
	Journal Journal `json:"journals"`
}

type Journal struct {
	// 受け付けID
	ID uint64 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// ダウンロード形式
	DownloadType string `json:"download_type"`
	// 取得開始日 (yyyy-mm-dd)
	StartDate string `json:"start_date"`
	// 取得終了日 (yyyy-mm-dd)
	EndDate string `json:"end_date"`
	// 補助科目やコメントとして出力する項目
	VisibleTags []string `json:"visible_tags"`
	// 追加出力するID項目
	VisibleIDs []string `json:"visible_ids"`
	// ステータス確認用URL
	StatusURL string `json:"status_url"`
	// 集計結果が最新かどうか
	UpToDate bool `json:"up_to_date"`
	// 集計が最新でない場合の要因情報
	UpToDateReasons []JournalUpToDateReasons `json:"up_to_date_reasons"`
}

type JournalUpToDateReasons struct {
	// コード
	Code string `json:"code"`
	// 集計が最新でない理由
	Message string `json:"message"`
}

type GetJournalsOpts struct {
	// ダウンロード形式
	DownloadType string `json:"download_type"`
	// 補助科目やコメントとして出力する項目
	VisibleTags []string `json:"visible_tags,omitempty"`
	// 追加出力するID項目
	VisibleIDs []string `json:"visible_ids,omitempty"`
	// 取得開始日 (yyyy-mm-dd)
	StartDate string `json:"start_date,omitempty"`
	// 取得終了日 (yyyy-mm-dd)
	EndDate string `json:"end_date,omitempty"`
}

func (c *Client) GetJournals(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, opts GetJournalsOpts,
) (*JournalResponse, *oauth2.Token, error) {
	var result JournalResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, APIPathJournals, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	return &result, oauth2Token, nil
}
