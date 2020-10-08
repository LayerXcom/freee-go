# QuotationUpdateParamsQuotationContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 見積内容ID | [optional] 
**Order** | **int32** | 順序 | 
**Type** | **string** | 行の種類 &lt;ul&gt; &lt;li&gt;normal、discountを指定する場合、account_item_id,tax_codeとunit_priceが必須となります。&lt;/li&gt; &lt;li&gt;normalを指定した場合、qtyが必須となります。&lt;/li&gt; &lt;/ul&gt; | 
**Qty** | **float32** | 数量 | [optional] 
**Unit** | **string** | 単位 | [optional] 
**UnitPrice** | **float32** | 単価 (tax_entry_method: inclusiveの場合は税込価格、tax_entry_method: exclusiveの場合は税抜価格となります) | [optional] 
**Vat** | Pointer to **int32** | 消費税額 | [optional] 
**Description** | **string** | 備考 | [optional] 
**AccountItemId** | **int32** | 勘定科目ID | [optional] 
**TaxCode** | **int32** | 税区分コード | [optional] 
**ItemId** | **int32** | 品目ID | [optional] 
**SectionId** | **int32** | 部門ID | [optional] 
**TagIds** | **[]int32** |  | [optional] 
**Segment1TagId** | **int32** | セグメント１ID | [optional] 
**Segment2TagId** | **int32** | セグメント２ID | [optional] 
**Segment3TagId** | **int32** | セグメント３ID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


