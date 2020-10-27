package freee

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathPartners = "partners"
)

type Partners struct {
	Partners []Partner `json:"partners"`
}

type PartnerResponse struct {
	Partner Partner `json:"partner"`
}

type Partner struct {
	// 取引先ID
	ID int32 `json:"id"`
	// 取引先コード
	Code *string `json:"code,omitempty"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 取引先名
	Name string `json:"name"`
	// ショートカット1 (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット2 (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
	// 事業所種別（null: 未設定、1: 法人、2: 個人）
	OrgCode *int32 `json:"org_code,omitempty"`
	// 地域（JP: 国内、ZZ:国外）
	CountryCode string `json:"country_code,omitempty"`
	// 正式名称（255文字以内）
	LongName *string `json:"long_name,omitempty"`
	// カナ名称（255文字以内）
	NameKana *string `json:"name_kana,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	DefaultTitle *string `json:"default_title,omitempty"`
	// 電話番号
	Phone *string `json:"phone,omitempty"`
	// 担当者 氏名
	ContactName *string `json:"contact_name,omitempty"`
	// 担当者 メールアドレス
	Email *string `json:"email,omitempty"`
	// 振込元口座ID（一括振込ファイル用）:（未設定の場合は、nullです。）
	PayerWalletableID *int32 `json:"payer_walletable_id,omitempty"`
	// 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee)
	TransferFeeHandlingSide string `json:"transfer_fee_handling_side,omitempty"`
	// 郵便番号
	AddressAttributesZipcode *string `json:"address_attributes[zipcode],omitempty"`
	// 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	AddressAttributesPrefectureCode int32 `json:"address_attributes[prefecture_code],omitempty"`
	// 市区町村・番地
	AddressAttributesStreetName1 *string `json:"address_attributes[street_name1],omitempty"`
	// 建物名・部屋番号など
	AddressAttributesStreetName2 *string `json:"address_attributes[street_name2],omitempty"`
	// 請求書送付方法(mail:メール、posting:郵送、mail_and_posting:メールと郵送)
	PartnerDocSettingAttributesSendingMethod *string `json:"partner_doc_setting_attributes[sending_method],omitempty"`
	// 銀行名
	PartnerBankAccountAttributesBankName *string `json:"partner_bank_account_attributes[bank_name],omitempty"`
	// 銀行名（カナ）
	PartnerBankAccountAttributesBankNameKana *string `json:"partner_bank_account_attributes[bank_name_kana],omitempty"`
	// 銀行番号
	PartnerBankAccountAttributesBankCode *string `json:"partner_bank_account_attributes[bank_code],omitempty"`
	// 支店名
	PartnerBankAccountAttributesBranchName *string `json:"partner_bank_account_attributes[branch_name],omitempty"`
	// 支店名（カナ）
	PartnerBankAccountAttributesBranchKana *string `json:"partner_bank_account_attributes[branch_kana],omitempty"`
	// 支店番号
	PartnerBankAccountAttributesBranchCode *string `json:"partner_bank_account_attributes[branch_code],omitempty"`
	// 口座種別(ordinary:普通、checking:当座、earmarked:納税準備預金、savings:貯蓄、other:その他)
	PartnerBankAccountAttributesAccountType *string `json:"partner_bank_account_attributes[account_type],omitempty"`
	// 口座番号
	PartnerBankAccountAttributesAccountNumber *string `json:"partner_bank_account_attributes[account_number],omitempty"`
	// 受取人名（カナ）
	PartnerBankAccountAttributesAccountName *string `json:"partner_bank_account_attributes[account_name],omitempty"`
	// 受取人名
	PartnerBankAccountAttributesLongAccountName *string `json:"partner_bank_account_attributes[long_account_name],omitempty"`
}

type CreatePartnerParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 取引先名 (255文字以内)
	Name string `json:"name"`
	// 取引先コード（取引先コードの利用を有効にしている場合は、codeの指定は必須です。）
	Code string `json:"code,omitempty"`
	// ショートカット１ (255文字以内)
	Shortcut1 string `json:"shortcut1,omitempty"`
	// ショートカット２ (255文字以内)
	Shortcut2 string `json:"shortcut2,omitempty"`
	// 事業所種別（null: 未設定、1: 法人、2: 個人）
	OrgCode *int32 `json:"org_code,omitempty"`
	// 地域（JP: 国内、ZZ:国外）
	CountryCode string `json:"country_code,omitempty"`
	// 正式名称（255文字以内）
	LongName string `json:"long_name,omitempty"`
	// カナ名称（255文字以内）
	NameKana string `json:"name_kana,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	DefaultTitle string `json:"default_title,omitempty"`
	// 電話番号
	Phone string `json:"phone,omitempty"`
	// 担当者 氏名 (255文字以内)
	ContactName string `json:"contact_name,omitempty"`
	// 担当者 メールアドレス (255文字以内)
	Email string `json:"email,omitempty"`
	// 振込元口座ID（一括振込ファイル用）:（walletableのtypeが'bank_account'のidのみ指定できます。また、未設定にする場合は、nullを指定してください。）
	PayerWalletableID *int32 `json:"payer_walletable_id,omitempty"`
	// 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee)
	TransferFeeHandlingSide      string                                          `json:"transfer_fee_handling_side,omitempty"`
	AddressAttributes            CreatePartnerParamsAddressAttributes            `json:"address_attributes,omitempty"`
	PartnerDocSettingAttributes  CreatePartnerParamsPartnerDocSettingAttributes  `json:"partner_doc_setting_attributes,omitempty"`
	PartnerBankAccountAttributes CreatePartnerParamsPartnerBankAccountAttributes `json:"partner_bank_account_attributes,omitempty"`
	PaymentTermAttributes        CreatePartnerParamsPaymentTermAttributes        `json:"payment_term_attributes,omitempty"`
	InvoicePaymentTermAttributes CreatePartnerParamsPaymentTermAttributes        `json:"invoice_payment_term_attributes,omitempty"`
}

type CreatePartnerParamsAddressAttributes struct {
	// 郵便番号（8文字以内）
	Zipcode string `json:"zipcode,omitempty"`
	// 都道府県コード（0: 北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	PrefectureCode int32 `json:"prefecture_code,omitempty"`
	// 市区町村・番地（255文字以内）
	StreetName1 string `json:"street_name1,omitempty"`
	// 建物名・部屋番号など（255文字以内）
	StreetName2 string `json:"street_name2,omitempty"`
}

type CreatePartnerParamsPartnerDocSettingAttributes struct {
	// 請求書送付方法(email:メール、posting:郵送、email_and_posting:メールと郵送)
	SendingMethod string `json:"sending_method,omitempty"`
}

type CreatePartnerParamsPartnerBankAccountAttributes struct {
	// 銀行名
	BankName string `json:"bank_name,omitempty"`
	// 銀行名（カナ）
	BankNameKana string `json:"bank_name_kana,omitempty"`
	// 銀行番号
	BankCode string `json:"bank_code,omitempty"`
	// 支店名
	BranchName string `json:"branch_name,omitempty"`
	// 支店名（カナ）
	BranchKana string `json:"branch_kana,omitempty"`
	// 支店番号
	BranchCode string `json:"branch_code,omitempty"`
	// 口座種別(ordinary:普通、checking：当座、earmarked：納税準備預金、savings：貯蓄、other:その他)
	AccountType string `json:"account_type,omitempty"`
	// 口座番号
	AccountNumber string `json:"account_number,omitempty"`
	// 受取人名
	LongAccountName string `json:"long_account_name,omitempty"`
	// 受取人名（カナ）
	AccountName string `json:"account_name,omitempty"`
}

type CreatePartnerParamsPaymentTermAttributes struct {
	// 締め日（29, 30, 31日の末日を指定する場合は、32を指定してください。）
	CutoffDay int32 `json:"cutoff_day,omitempty"`
	// 支払月
	AdditionalMonths int32 `json:"additional_months,omitempty"`
	// 支払日（29, 30, 31日の末日を指定する場合は、32を指定してください。）
	FixedDay int32 `json:"fixed_day,omitempty"`
}

func (c *Client) CreatePartner(
	ctx context.Context, oauth2Token *oauth2.Token,
	params CreatePartnerParams,
) (*Partner, *oauth2.Token, error) {
	var result PartnerResponse

	tokenSource, err := c.call(ctx, APIPathPartners, http.MethodPost, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, oauth2Token, err
	}
	if &result == nil || &result.Partner == nil {
		return nil, token, errors.New("failed to parse response")
	}
	return &result.Partner, token, nil
}

type UpdatePartnerParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 取引先名 (255文字以内)
	Name string `json:"name"`
	// ショートカット１ (255文字以内)
	Shortcut1 string `json:"shortcut1,omitempty"`
	// ショートカット２ (255文字以内)
	Shortcut2 string `json:"shortcut2,omitempty"`
	// 事業所種別（null: 未設定、1: 法人、2: 個人）
	OrgCode *int32 `json:"org_code,omitempty"`
	// 地域（JP: 国内、ZZ:国外）
	CountryCode string `json:"country_code,omitempty"`
	// 正式名称（255文字以内）
	LongName string `json:"long_name,omitempty"`
	// カナ名称（255文字以内）
	NameKana string `json:"name_kana,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	DefaultTitle string `json:"default_title,omitempty"`
	// 電話番号
	Phone string `json:"phone,omitempty"`
	// 担当者 氏名 (255文字以内)
	ContactName string `json:"contact_name,omitempty"`
	// 担当者 メールアドレス (255文字以内)
	Email string `json:"email,omitempty"`
	// 振込元口座ID（一括振込ファイル用）:（walletableのtypeが'bank_account'のidのみ指定できます。また、未設定にする場合は、nullを指定してください。）
	PayerWalletableID *int32 `json:"payer_walletable_id,omitempty"`
	// 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee)
	TransferFeeHandlingSide      string                                          `json:"transfer_fee_handling_side,omitempty"`
	AddressAttributes            CreatePartnerParamsAddressAttributes            `json:"address_attributes,omitempty"`
	PartnerDocSettingAttributes  CreatePartnerParamsPartnerDocSettingAttributes  `json:"partner_doc_setting_attributes,omitempty"`
	PartnerBankAccountAttributes CreatePartnerParamsPartnerBankAccountAttributes `json:"partner_bank_account_attributes,omitempty"`
	PaymentTermAttributes        CreatePartnerParamsPaymentTermAttributes        `json:"payment_term_attributes,omitempty"`
	InvoicePaymentTermAttributes CreatePartnerParamsPaymentTermAttributes        `json:"invoice_payment_term_attributes,omitempty"`
}

func (c *Client) UpdatePartner(
	ctx context.Context, oauth2Token *oauth2.Token,
	partnerID uint32, params UpdatePartnerParams,
) (*Partner, *oauth2.Token, error) {
	var result PartnerResponse

	tokenSource, err := c.call(ctx, path.Join(APIPathPartners, fmt.Sprint(partnerID)), http.MethodPut, oauth2Token, nil, params, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, oauth2Token, err
	}
	if &result == nil || &result.Partner == nil {
		return nil, token, errors.New("failed to parse response")
	}
	return &result.Partner, token, nil
}

type GetPartnersOpts struct {
	Offset  uint32 `url:"offset,omitempty"`
	Limit   uint32 `url:"limit,omitempty"`
	Keyword string `url:"keyword,omitempty"`
}

func (c *Client) GetPartners(
	ctx context.Context, oauth2Token *oauth2.Token,
	companyID uint32, opts GetPartnersOpts,
) (*Partners, *oauth2.Token, error) {
	var result Partners

	v, err := query.Values(opts)
	if err != nil {
		return nil, oauth2Token, err
	}
	SetCompanyID(&v, companyID)
	tokenSource, err := c.call(ctx, APIPathPartners, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, oauth2Token, err
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, oauth2Token, err
	}
	return &result, token, nil
}
