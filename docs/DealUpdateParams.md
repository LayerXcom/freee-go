# DealUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**Type** | **string** | 収支区分 (収入: income, 支出: expense) | 
**CompanyId** | **int32** | 事業所ID | 
**DueDate** | **string** | 支払期日(yyyy-mm-dd) | [optional] 
**PartnerId** | **int32** | 取引先ID | [optional] 
**PartnerCode** | **string** | 取引先コード | [optional] 
**RefNumber** | **string** | 管理番号 | [optional] 
**Details** | [**[]DealUpdateParamsDetails**](dealUpdateParams_details.md) |  | 
**ReceiptIds** | Pointer to **[]int32** | 証憑ファイルID（配列） | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


