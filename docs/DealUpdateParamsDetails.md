# DealUpdateParamsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 取引行ID: 既存取引行を更新する場合に指定します。IDを指定しない取引行は、新規行として扱われ追加されます。また、detailsに含まれない既存の取引行は削除されます。更新後も残したい行は、必ず取引行IDを指定してdetailsに含めてください。 | [optional] 
**TaxCode** | **int32** | 税区分コード | 
**AccountItemId** | **int32** | 勘定科目ID | 
**Amount** | **int32** | 取引金額（税込で指定してください） | 
**ItemId** | **int32** | 品目ID | [optional] 
**SectionId** | **int32** | 部門ID | [optional] 
**TagIds** | **[]int32** | メモタグID | [optional] 
**Segment1TagId** | **int32** | セグメント１ID | [optional] 
**Segment2TagId** | **int32** | セグメント２ID | [optional] 
**Segment3TagId** | **int32** | セグメント３ID | [optional] 
**Description** | **string** | 備考 | [optional] 
**Vat** | **int32** | 消費税額（指定しない場合は自動で計算されます） | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


