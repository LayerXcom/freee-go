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
