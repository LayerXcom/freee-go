# PartnerCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Name** | **string** | 取引先名 (255文字以内) | 
**Code** | **string** | 取引先コード（取引先コードの利用を有効にしている場合は、codeの指定は必須です。） | [optional] 
**Shortcut1** | **string** | ショートカット１ (255文字以内) | [optional] 
**Shortcut2** | **string** | ショートカット２ (255文字以内) | [optional] 
**OrgCode** | Pointer to **int32** | 事業所種別（null: 未設定、1: 法人、2: 個人） | [optional] 
**CountryCode** | **string** | 地域（JP: 国内、ZZ:国外） | [optional] 
**LongName** | **string** | 正式名称（255文字以内） | [optional] 
**NameKana** | **string** | カナ名称（255文字以内） | [optional] 
**DefaultTitle** | **string** | 敬称（御中、様、(空白)の3つから選択） | [optional] 
**Phone** | **string** | 電話番号 | [optional] 
**ContactName** | **string** | 担当者 氏名 (255文字以内) | [optional] 
**Email** | **string** | 担当者 メールアドレス (255文字以内) | [optional] 
**PayerWalletableId** | Pointer to **int32** | 振込元口座ID（一括振込ファイル用）:（walletableのtypeが&#39;bank_account&#39;のidのみ指定できます。また、未設定にする場合は、nullを指定してください。） | [optional] 
**TransferFeeHandlingSide** | **string** | 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee) | [optional] 
**AddressAttributes** | [**PartnerCreateParamsAddressAttributes**](partnerCreateParams_address_attributes.md) |  | [optional] 
**PartnerDocSettingAttributes** | [**PartnerCreateParamsPartnerDocSettingAttributes**](partnerCreateParams_partner_doc_setting_attributes.md) |  | [optional] 
**PartnerBankAccountAttributes** | [**PartnerCreateParamsPartnerBankAccountAttributes**](partnerCreateParams_partner_bank_account_attributes.md) |  | [optional] 
**PaymentTermAttributes** | [**PartnerCreateParamsPaymentTermAttributes**](partnerCreateParams_payment_term_attributes.md) |  | [optional] 
**InvoicePaymentTermAttributes** | [**PartnerCreateParamsPaymentTermAttributes**](partnerCreateParams_payment_term_attributes.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


