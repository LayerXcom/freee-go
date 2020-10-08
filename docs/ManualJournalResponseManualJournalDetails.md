# ManualJournalResponseManualJournalDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 貸借行ID | 
**EntrySide** | **string** | 貸借(貸方: credit, 借方: debit) | 
**AccountItemId** | **int32** | 勘定科目ID | 
**TaxCode** | **int32** | 税区分コード | 
**PartnerId** | Pointer to **int32** | 取引先ID | 
**PartnerName** | Pointer to **string** | 取引先名 | 
**PartnerCode** | Pointer to **string** | 取引先コード | [optional] 
**PartnerLongName** | Pointer to **string** | 正式名称（255文字以内） | 
**ItemId** | Pointer to **int32** | 品目ID | 
**ItemName** | Pointer to **string** | 品目 | 
**SectionId** | Pointer to **int32** | 部門ID | 
**SectionName** | Pointer to **string** | 部門 | 
**TagIds** | **[]int32** |  | 
**TagNames** | **[]string** |  | 
**Segment1TagId** | **int32** | セグメント１ID | [optional] 
**Segment1TagName** | **int32** | セグメント１ID | [optional] 
**Segment2TagId** | **int32** | セグメント２ID | [optional] 
**Segment2TagName** | **int32** | セグメント２ | [optional] 
**Segment3TagId** | **int32** | セグメント３ID | [optional] 
**Segment3TagName** | **int32** | セグメント３ | [optional] 
**Amount** | **int32** | 金額（税込で指定してください） | 
**Vat** | **int32** | 消費税額（指定しない場合は自動で計算されます） | 
**Description** | **string** | 備考 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


