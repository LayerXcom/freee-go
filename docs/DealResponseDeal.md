# DealResponseDeal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 取引ID | 
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**DueDate** | **string** | 支払期日 (yyyy-mm-dd) | [optional] 
**Amount** | **int32** | 金額 | 
**DueAmount** | **int32** | 支払金額 | [optional] 
**Type** | **string** | 収支区分 (収入: income, 支出: expense) | [optional] 
**PartnerId** | **int32** | 取引先ID | 
**PartnerCode** | Pointer to **string** | 取引先コード | [optional] 
**RefNumber** | **string** | 管理番号 | [optional] 
**Status** | **string** | 決済状況 (未決済: unsettled, 完了: settled) | 
**Details** | [**[]DealCreateResponseDealDetails**](dealCreateResponse_deal_details.md) | 取引の明細行 | [optional] 
**Renews** | [**[]DealResponseDealRenews**](dealResponse_deal_renews.md) | 取引の+更新行 | [optional] 
**Payments** | [**[]DealCreateResponseDealPayments**](dealCreateResponse_deal_payments.md) | 取引の支払行 | [optional] 
**Receipts** | [**[]DealResponseDealReceipts**](dealResponse_deal_receipts.md) | 証憑ファイル | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


