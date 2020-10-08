# WalletTxnResponseWalletTxn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 明細ID | 
**CompanyId** | **int32** | 事業所ID | 
**Date** | **string** | 取引日(yyyy-mm-dd) | 
**Amount** | **int32** | 取引金額 | 
**DueAmount** | **int32** | 未決済金額 | 
**Balance** | **int32** | 残高(銀行口座等) | 
**EntrySide** | **string** | 入金／出金 (入金: income, 出金: expense) | 
**WalletableType** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**WalletableId** | **int32** | 口座ID | 
**Description** | **string** | 取引内容 | 
**Status** | **int32** | 明細のステータス（消込待ち: 1, 消込済み: 2, 無視: 3, 消込中: 4） | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


