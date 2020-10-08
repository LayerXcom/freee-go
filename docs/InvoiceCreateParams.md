# InvoiceCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 請求日 (yyyy-mm-dd) | [optional] 
**PartnerId** | Pointer to **int32** | 取引先ID | [optional] 
**PartnerCode** | Pointer to **string** | 取引先コード | [optional] 
**InvoiceNumber** | **string** | 請求書番号 (デフォルト: 自動採番されます) | [optional] 
**Title** | **string** | タイトル (デフォルト: 請求書) | [optional] 
**DueDate** | **string** | 期日 (yyyy-mm-dd) | [optional] 
**BookingDate** | **string** | 売上計上日 | [optional] 
**Description** | **string** | 概要 | [optional] 
**InvoiceStatus** | **string** | 請求書ステータス&lt;br&gt; &lt;ul&gt;   &lt;li&gt;draft: 下書き (デフォルト)&lt;/li&gt;   &lt;li&gt;(廃止予定) issue: 発行 (送付待ち (unsubmitted) と同じです。)&lt;/li&gt;   &lt;li&gt;unsubmitted: 送付待ち&lt;/li&gt;   &lt;li&gt;submitted: 送付済み&lt;/li&gt; &lt;/ul&gt; issue, unsubmitted, submitted は請求書承認ワークフローを利用している場合は指定できません。  | [optional] 
**PartnerDisplayName** | **string** | 請求書に表示する取引先名 | 
**PartnerTitle** | Pointer to **string** | 敬称（御中、様、(空白)の3つから選択） | 
**PartnerContactInfo** | Pointer to **string** | 取引先担当者名 | [optional] 
**PartnerZipcode** | Pointer to **string** | 取引先郵便番号 (デフォルトはpartner_idもしくははpartner_codeで指定された取引先設定情報が補完されます) | [optional] 
**PartnerPrefectureCode** | Pointer to **int32** | 取引先都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄) (デフォルトはpartner_idもしくははpartner_codeで指定された取引先設定情報が補完されます) | [optional] 
**PartnerAddress1** | Pointer to **string** | 取引先市区町村・番地 (デフォルトはpartner_idもしくははpartner_codeで指定された取引先設定情報が補完されます) | [optional] 
**PartnerAddress2** | Pointer to **string** | 取引先建物名・部屋番号など (デフォルトはpartner_idもしくははpartner_codeで指定された取引先設定情報が補完されます) | [optional] 
**CompanyName** | **string** | 事業所名 (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyZipcode** | **string** | 郵便番号 (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyPrefectureCode** | **int32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄) (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyAddress1** | **string** | 市区町村・番地 (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyAddress2** | **string** | 建物名・部屋番号など (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyContactInfo** | **string** | 事業所担当者名 (デフォルトは事業所設定情報が補完されます) | [optional] 
**PaymentType** | **string** | 支払方法 (振込: transfer, 引き落とし: direct_debit) | [optional] 
**PaymentBankInfo** | **string** | 支払口座 | [optional] 
**UseVirtualTransferAccount** | **string** | 振込専用口座の利用(利用しない: not_use(デフォルト), 利用する: use) | [optional] 
**Message** | **string** | メッセージ (デフォルト: 下記の通りご請求申し上げます。) | [optional] 
**Notes** | **string** | 備考 | [optional] 
**InvoiceLayout** | **string** | 請求書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;carried_forward_standard_classic&#x60; - レイアウト３（繰越金額欄あり）/クラシック  * &#x60;carried_forward_envelope_classic&#x60; - 封筒２（繰越金額欄あり）/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | [optional] 
**TaxEntryMethod** | **string** | 請求書の消費税計算方法(inclusive: 内税表示, exclusive: 外税表示 (デフォルト)) | [optional] 
**InvoiceContents** | [**[]InvoiceCreateParamsInvoiceContents**](invoiceCreateParams_invoice_contents.md) | 請求内容 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


