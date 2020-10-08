# TransferResponseTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 取引(振替)ID | 
**CompanyId** | **int32** | 事業所ID | 
**Amount** | **int32** | 金額 | 
**Date** | **string** | 振替日 (yyyy-mm-dd) | 
**FromWalletableType** | Pointer to **string** | 振替元口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**FromWalletableId** | **int32** | 振替元口座ID | 
**ToWalletableType** | Pointer to **string** | 振替先口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**ToWalletableId** | **int32** | 振替先口座ID | 
**Description** | **string** | 備考 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


