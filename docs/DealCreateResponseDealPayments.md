# DealCreateResponseDealPayments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 取引行ID | 
**Date** | **string** | 支払日 | 
**FromWalletableType** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item) | [optional] 
**FromWalletableId** | **int32** | 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID） | [optional] 
**Amount** | **int32** | 支払金額 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


