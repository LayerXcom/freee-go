# QuotationResponseQuotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 見積書ID | 
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 見積日 (yyyy-mm-dd) | 
**PartnerId** | Pointer to **int32** | 取引先ID | 
**PartnerCode** | Pointer to **string** | 取引先コード | [optional] 
**QuotationNumber** | **string** | 見積書番号 | 
**Title** | Pointer to **string** | タイトル | [optional] 
**TotalAmount** | **int32** | 合計金額 | 
**TotalVat** | **int32** | 消費税 | [optional] 
**SubTotal** | **int32** | 小計 | [optional] 
**Description** | Pointer to **string** | 概要 | [optional] 
**QuotationStatus** | **string** | 見積書ステータス  (unsubmitted: 送付待ち, submitted: 送付済み, all: 全て) | 
**WebPublishedAt** | Pointer to **string** | Web共有日時(最新) | [optional] 
**WebDownloadedAt** | Pointer to **string** | Web共有ダウンロード日時(最新) | [optional] 
**WebConfirmedAt** | Pointer to **string** | Web共有取引先確認日時(最新) | [optional] 
**MailSentAt** | Pointer to **string** | メール送信日時(最新) | [optional] 
**PartnerName** | Pointer to **string** | 取引先名 | [optional] 
**PartnerDisplayName** | Pointer to **string** | 見積書に表示する取引先名 | [optional] 
**PartnerTitle** | Pointer to **string** | 敬称（御中、様、(空白)の3つから選択） | 
**PartnerZipcode** | Pointer to **string** | 郵便番号 | [optional] 
**PartnerPrefectureCode** | Pointer to **int32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**PartnerPrefectureName** | Pointer to **string** | 都道府県 | [optional] 
**PartnerAddress1** | Pointer to **string** | 市区町村・番地 | [optional] 
**PartnerAddress2** | Pointer to **string** | 建物名・部屋番号など | [optional] 
**PartnerContactInfo** | Pointer to **string** | 取引先担当者名 | [optional] 
**CompanyName** | **string** | 事業所名 | 
**CompanyZipcode** | Pointer to **string** | 郵便番号 | [optional] 
**CompanyPrefectureCode** | Pointer to **int32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**CompanyPrefectureName** | Pointer to **string** | 都道府県 | [optional] 
**CompanyAddress1** | Pointer to **string** | 市区町村・番地 | [optional] 
**CompanyAddress2** | Pointer to **string** | 建物名・部屋番号など | [optional] 
**CompanyContactInfo** | Pointer to **string** | 事業所担当者名 | [optional] 
**Message** | Pointer to **string** | メッセージ | [optional] 
**Notes** | Pointer to **string** | 備考 | [optional] 
**QuotationLayout** | **string** | 見積書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | 
**TaxEntryMethod** | **string** | 見積書の消費税計算方法(inclusive: 内税, exclusive: 外税) | 
**QuotationContents** | [**[]QuotationResponseQuotationQuotationContents**](quotationResponse_quotation_quotation_contents.md) | 見積内容 | [optional] 
**TotalAmountPerVatRate** | [**InvoiceResponseInvoiceTotalAmountPerVatRate**](invoiceResponse_invoice_total_amount_per_vat_rate.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


