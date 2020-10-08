# AccountItemResponseAccountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 勘定科目ID | 
**Name** | **string** | 勘定科目名 (30文字以内) | 
**CompanyId** | **int32** | 事業所ID | 
**TaxCode** | **int32** | 税区分コード | 
**AccountCategory** | **string** | 勘定科目カテゴリー | 
**AccountCategoryId** | **int32** | 勘定科目のカテゴリーID | 
**Shortcut** | **string** | ショートカット1 (20文字以内) | [optional] 
**ShortcutNum** | **string** | ショートカット2(勘定科目コード) (20文字以内) | [optional] 
**Searchable** | **int32** | 検索可能:2, 検索不可：3 | 
**AccumulatedDepAccountItemName** | **string** | 減価償却累計額勘定科目（法人のみ利用可能） | [optional] 
**AccumulatedDepAccountItemId** | Pointer to **int32** | 減価償却累計額勘定科目ID（法人のみ利用可能） | [optional] 
**Items** | [**[]AccountItemResponseAccountItemItems**](accountItemResponse_account_item_items.md) |  | [optional] 
**Partners** | [**[]AccountItemResponseAccountItemPartners**](accountItemResponse_account_item_partners.md) |  | [optional] 
**Available** | **bool** | 勘定科目の使用設定（true: 使用する、false: 使用しない） | 
**WalletableId** | Pointer to **int32** | 口座ID | 
**GroupName** | Pointer to **string** | 決算書表示名（小カテゴリー） | [optional] 
**CorrespondingIncomeName** | Pointer to **string** | 収入取引相手勘定科目名 | [optional] 
**CorrespondingIncomeId** | Pointer to **int32** | 収入取引相手勘定科目ID | [optional] 
**CorrespondingExpenseName** | Pointer to **string** | 支出取引相手勘定科目名 | [optional] 
**CorrespondingExpenseId** | Pointer to **int32** | 支出取引相手勘定科目ID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


