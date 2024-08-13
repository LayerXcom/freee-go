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
	APIPathReceipts = "receipts"
)

type CreateReceiptParams struct {
	// 事業所ID
	CompanyID int64 `json:"company_id"`
	// メモ (255文字以内)
	Description string `json:"description,omitempty"`
	// 取引日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 証憑ファイル
	Receipt []byte `json:"receipt"`
}

type Receipts struct {
	Receipts []Receipt `json:"receipts"`
}

type ReceiptResponse struct {
	Receipt Receipt `json:"receipt"`
}

type Receipt struct {
	// 証憑ID
	ID int64 `json:"id"`
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

type GetReceiptOpts struct {
	StartDate        string `url:"start_date"`
	EndDate          string `url:"end_date"`
	UserName         string `url:"user_name,omitempty"`
	Number           int64  `url:"number,omitempty"`
	CommentType      string `url:"comment_type,omitempty"`
	CommentImportant bool   `url:"comment_important,omitempty"`
	Category         string `url:"category,omitempty"`
	Offset           int64  `url:"offset,omitempty"`
	Limit            int64  `url:"limit,omitempty"`
}

type UserCreatedReceipt struct {
	// ユーザーID
	ID int64 `json:"id"`
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
	postBody := map[string]string{
		"company_id":  fmt.Sprint(params.CompanyID),
		"description": params.Description,
		"issue_date":  params.IssueDate,
	}
	var result ReceiptResponse
	oauth2Token, err := c.postFiles(ctx, APIPathReceipts, http.MethodPost, oauth2Token, nil, postBody, receiptName, params.Receipt, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}

func (c *Client) GetReceipt(
	ctx context.Context, oauth2Token *oauth2.Token, companyID int64, receiptID int64,
) (*ReceiptResponse, *oauth2.Token, error) {
	var result ReceiptResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	oauth2Token, err = c.call(ctx, path.Join(APIPathReceipts, fmt.Sprint(receiptID)), http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, oauth2Token, nil
}
