# ManualJournalCreateParamsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntrySide** | **string** | 貸借（貸方: credit, 借方: debit） | 
**TaxCode** | **int32** | 税区分コード | 
**AccountItemId** | **int32** | 勘定科目ID | 
**Amount** | **int32** | 取引金額（税込で指定してください） | 
**Vat** | **int32** | 消費税額（指定しない場合は自動で計算されます） | [optional] 
**PartnerId** | **int32** | 取引先ID | [optional] 
**PartnerCode** | **string** | 取引先コード | [optional] 
**ItemId** | **int32** | 品目ID | [optional] 
**SectionId** | **int32** | 部門ID | [optional] 
**TagIds** | **[]int32** | メモタグID | [optional] 
**Segment1TagId** | **int32** | セグメント１ID | [optional] 
**Segment2TagId** | **int32** | セグメント２ID | [optional] 
**Segment3TagId** | **int32** | セグメント３ID | [optional] 
**Description** | **string** | 備考 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


