# PartnerResponsePartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 取引先ID | 
**Code** | Pointer to **string** | 取引先コード | 
**CompanyId** | **int32** | 事業所ID | 
**Name** | **string** | 取引先名 | 
**Shortcut1** | Pointer to **string** | ショートカット1 (20文字以内) | [optional] 
**Shortcut2** | Pointer to **string** | ショートカット2 (20文字以内) | [optional] 
**OrgCode** | Pointer to **int32** | 事業所種別（null: 未設定、1: 法人、2: 個人） | [optional] 
**CountryCode** | **string** | 地域（JP: 国内、ZZ:国外） | [optional] 
**LongName** | Pointer to **string** | 正式名称（255文字以内） | [optional] 
**NameKana** | Pointer to **string** | カナ名称（255文字以内） | [optional] 
**DefaultTitle** | Pointer to **string** | 敬称（御中、様、(空白)の3つから選択） | [optional] 
**Phone** | Pointer to **string** | 電話番号 | [optional] 
**ContactName** | Pointer to **string** | 担当者 氏名 | [optional] 
**Email** | Pointer to **string** | 担当者 メールアドレス | [optional] 
**PayerWalletableId** | Pointer to **int32** | 振込元口座ID（一括振込ファイル用）:（未設定の場合は、nullです。） | [optional] 
**TransferFeeHandlingSide** | **string** | 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee) | [optional] 
**AddressAttributesZipcode** | Pointer to **string** | 郵便番号 | [optional] 
**AddressAttributesPrefectureCode** | **int32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**AddressAttributesStreetName1** | Pointer to **string** | 市区町村・番地 | [optional] 
**AddressAttributesStreetName2** | Pointer to **string** | 建物名・部屋番号など | [optional] 
**PartnerDocSettingAttributesSendingMethod** | Pointer to **string** | 請求書送付方法(email:メール、posting:郵送、email_and_posting:メールと郵送) | [optional] 
**PartnerBankAccountAttributesBankName** | Pointer to **string** | 銀行名 | [optional] 
**PartnerBankAccountAttributesBankNameKana** | Pointer to **string** | 銀行名（カナ） | [optional] 
**PartnerBankAccountAttributesBankCode** | Pointer to **string** | 銀行番号 | [optional] 
**PartnerBankAccountAttributesBranchName** | Pointer to **string** | 支店名 | [optional] 
**PartnerBankAccountAttributesBranchKana** | Pointer to **string** | 支店名（カナ） | [optional] 
**PartnerBankAccountAttributesBranchCode** | Pointer to **string** | 支店番号 | [optional] 
**PartnerBankAccountAttributesAccountType** | Pointer to **string** | 口座種別(ordinary:普通、checking:当座、earmarked:納税準備預金、savings:貯蓄、other:その他) | [optional] 
**PartnerBankAccountAttributesAccountNumber** | Pointer to **string** | 口座番号 | [optional] 
**PartnerBankAccountAttributesAccountName** | Pointer to **string** | 受取人名（カナ） | [optional] 
**PartnerBankAccountAttributesLongAccountName** | Pointer to **string** | 受取人名 | [optional] 
**PaymentTermAttributesCutoffDay** | Pointer to **int32** | 締め日（29, 30, 31日の末日を指定する場合は、32。） | [optional] 
**PaymentTermAttributesAdditionalMonths** | Pointer to **int32** | 支払月 | [optional] 
**PaymentTermAttributesFixedDay** | Pointer to **int32** | 支払日（29, 30, 31日の末日を指定する場合は、32。） | [optional] 
**InvoicePaymentTermAttributesCutoffDay** | Pointer to **int32** | 締め日（29, 30, 31日の末日を指定する場合は、32。） | [optional] 
**InvoicePaymentTermAttributesAdditionalMonths** | Pointer to **int32** | 支払月 | [optional] 
**InvoicePaymentTermAttributesFixedDay** | Pointer to **int32** | 支払日（29, 30, 31日の末日を指定する場合は、32。） | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


